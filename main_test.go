package main

import (
    "bytes"
    "fmt"
    "os"
    "os/exec"
    "testing"
)

func TestMain(t *testing.T) {
    tests := []struct {
        name     string
        args     []string
        output   string
        wantErr  bool
        mockFunc func()
    }{
        {
            name:   "Case 1: Python3 is not installed",
            args:   []string{"program_to_install"},
            output: "Python3 is not installed.\nStarting Python3 installation.\nPython3 installed.\nAnsible is not installed.\nStarting Ansible installation.\nAnsible installed.\nprogram_to_install is not installed.\nRunning Ansible playbook to install program_to_install...\nprogram_to_install installed and started.\n",
            mockFunc: func() {
                cmd := exec.Command("which", "python3")
                err := cmd.Run()
                if err == nil {
                    t.Errorf("Expected error, but got nil")
                }
            },
        },
        {
            name:   "Case 2: Python3 and Ansible are installed",
            args:   []string{"program_to_install"},
            output: "program_to_install is not installed.\nRunning Ansible playbook to install program_to_install...\nprogram_to_install installed and started.\n",
            mockFunc: func() {
                cmd := exec.Command("which", "python3")
                err := cmd.Run()
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }

                cmd = exec.Command("which", "ansible")
                err = cmd.Run()
                if err != nil {
                    t.Errorf("Unexpected error: %v", err)
                }
            },
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            defer func() {
                if r := recover(); r != nil {
                    t.Errorf("Unexpected panic: %v", r)
                }
            }()

            test.mockFunc()

            oldStdout := os.Stdout
            defer func() { os.Stdout = oldStdout }()
            b := &bytes.Buffer{}
            os.Stdout = b

            os.Args = test.args
            main()

            if b.String() != test.output {
                t.Errorf("Expected %q, but got %q", test.output, b.String())
            }
        })
    }
}
