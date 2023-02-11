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
    // Update package cache
    cmd := exec.Command("ansible", "localhost", "-m", "apt", "-a", "update_cache=yes")
    err := cmd.Run()
    if err != nil {
        fmt.Printf("Error updating package cache.\n")
        os.Exit(1)
    }

    // Install the program
    cmd = exec.Command("ansible", "localhost", "-m", "ansible.builtin.apt", "-a", fmt.Sprintf("name=%s state=latest", program))
    err = cmd.Run()
    if err != nil {
        fmt.Printf("Error installing %s.\n", program)
        os.Exit(1)
    }

    // Start the program
    cmd = exec.Command("ansible", "localhost", "-m", "service", "-a", fmt.Sprintf("name=%s state=started", program))
    err = cmd.Run()
    if err != nil {
        fmt.Printf("Error starting %s.\n", program)
        os.Exit(1)
    }

    fmt.Printf("%s installed and started.\n", program)
}
