// Easy Go Installer/Updater for any Linux packages

package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    // Check if Ansible is installed
    cmd := exec.Command("which", "ansible")
    err := cmd.Run()
    if err != nil {
        // Check Ansible and install it.
        fmt.Println("Ansible is not installed.")
        fmt.Println("Starting Ansible installation.")
        installAnsible()
        fmt.Println("Ansible is already installed.")
        }

    // Get the name of the program to install
    program := os.Args[len(os.Args)-1]

    // Check if the program is already installed
    cmd = exec.Command("which", program)
    err = cmd.Run()
    if err != nil {
        fmt.Printf("%s is not installed.\n", program)
        fmt.Printf("Running Ansible playbook to install %s...\n", program)
        installProgram(program)
    } else {
        fmt.Printf("%s is already installed.\n", program)
    }
}

func installAnsible() {
    cmd := exec.Command("apt", "update")
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error updating packages.")
        os.Exit(1)
    }
    cmd = exec.Command("apt", "install", "-y", "ansible")
    err = cmd.Run()
    if err != nil {
        fmt.Println("Error installing Ansible.")
        os.Exit(1)
    }
    fmt.Println("Ansible installed.")
}

func installProgram(program string) {
    // Run Ansible playbook to install the program
    cmd := exec.Command("ansible-playbook", "install-program.yml", "-e", fmt.Sprintf("program=%s", program))
    err := cmd.Run()
    if err != nil {
        fmt.Printf("Error installing %s.\n", program)
        os.Exit(1)
    }
    fmt.Printf("%s installed.\n", program)
}
