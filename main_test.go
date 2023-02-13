package main

import (
    "bytes"
    "os"
    "testing"
)

var tests = []struct {
    program string
    python3 bool
    ansible bool
    output  string
}{
    {"python3", true, false, "Python3 is already installed."},
    {"ansible", false, true, "Ansible is already installed."},
    {"nano", false, false, "Error updating package cache."},
    {"nano", false, true, "nano installed and started."},
}

func TestMain(t *testing.T) {
    for _, test := range tests {
        if test.python3 {
            os.Setenv("PATH", os.Getenv("PATH")+":/usr/bin/python3")
        }
        if test.ansible {
            os.Setenv("PATH", os.Getenv("PATH")+":/usr/bin/ansible")
        }

        // Capture stdout
        stdout := os.Stdout
        os.Stdout, _ = os.Pipe()

        // Call main
        os.Args = []string{"main", test.program}
        main()

        // Restore stdout
        os.Stdout = stdout

        // Check output
        output := bytes.TrimSpace(os.Stdout.Bytes())
        if output != test.output {
            t.Errorf("Unexpected output for program %q: got %q, want %q", test.program, output, test.output)
        }
    }
}
