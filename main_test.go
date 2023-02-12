package main

import (
    "bytes"
    "os/exec"
    "testing"
    "fmt"
)

func TestInstallPython3(t *testing.T) {
    cmd := exec.Command("which", "python3")
    err := cmd.Run()
    if err == nil {
        t.Fatalf("Python3 is already installed.")
    }

    // Redirect the standard output to a buffer
    var buf bytes.Buffer
    fmt.Fprintln(&buf, "Starting Python3 installation.")
    fmt.Fprintln(&buf, "Python3 installed.")

    // Call the installPython3 function
    installPython3()

    // Check if the expected output is generated
    if buf.String() != "Starting Python3 installation.\nPython3 installed.\n" {
        t.Fatalf("Unexpected output.\nExpected: %s\nActual: %s\n", "Starting Python3 installation.\nPython3 installed.\n", buf.String())
    }
}

func TestInstallAnsible(t *testing.T) {
    cmd := exec.Command("which", "ansible")
    err := cmd.Run()
    if err == nil {
        t.Fatalf("Ansible is already installed.")
    }

    // Redirect the standard output to a buffer
    var buf bytes.Buffer
    fmt.Fprintln(&buf, "Starting Ansible installation.")
    fmt.Fprintln(&buf, "Ansible installed.")

    // Call the installAnsible function
    installAnsible()

    // Check if the expected output is generated
    if buf.String() != "Starting Ansible installation.\nAnsible installed.\n" {
        t.Fatalf("Unexpected output.\nExpected: %s\nActual: %s\n", "Starting Ansible installation.\nAnsible installed.\n", buf.String())
    }
}

func TestInstallProgram(t *testing.T) {
    program := "nginx"

    // Redirect the standard output to a buffer
    var buf bytes.Buffer
    fmt.Fprintf(&buf, "%s is not installed.\n", program)
    fmt.Fprintf(&buf, "Running Ansible playbook to install %s...\n", program)
    fmt.Fprintf(&buf, "%s installed and started.\n", program)

    // Call the installProgram function
    installProgram(program)

    // Check if the expected output is generated
    if buf.String() != "nginx is not installed.\nRunning Ansible playbook to install nginx...\nginx installed and started.\n" {
        t.Fatalf("Unexpected output.\nExpected: %s\nActual: %s\n", "nginx is not installed.\nRunning Ansible playbook to install nginx...\nginx installed and started.\n", buf.String())
    }
}
