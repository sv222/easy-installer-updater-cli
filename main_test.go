package main

import (
    "bytes"
    "os/exec"
    "testing"
)

func TestInstallProgram(t *testing.T) {
    tests := []struct {
        name     string
        program  string
        output   string
        hasError bool
    }{
        {
            name:    "install nginx",
            program: "nginx",
            output:  "nginx installed and started.\n",
        },
        {
            name:    "install non-existent program",
            program: "nonexistent",
            output:  "Error installing nonexistent.\n",
            hasError: true,
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            original := exec.Command
            defer func() { exec.Command = original }()
            exec.Command = func(command string, args ...string) *exec.Cmd {
                return exec.Command("echo", "mocked output")
            }

            r, w, _ := os.Pipe()
            defer r.Close()
            defer w.Close()

            os.Stdout = w

            var b bytes.Buffer
            defer func() { os.Stdout = os.Stdout }()
            os.Stdout = &b

            installProgram(test.program)

            w.Close()
            output := b.String()

            if test.hasError {
                if output == test.output {
                    t.Errorf("expected output to be %q, got %q", test.output, output)
                }
                return
            }

            if output != test.output {
                t.Errorf("expected output to be %q, got %q", test.output, output)
            }
        })
    }
}
