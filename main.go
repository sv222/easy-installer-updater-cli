// Easy Go Installer/Updater for any Linux packages

package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {

    // Check if Python3 is installed
    cmd := exec.Command("which", "python3")
    err := cmd.Run()
    if err != nil {
        // Check Python3 and install it.
        fmt.Println("Python3 is not installed.")
        fmt.Println("Starting Python3 installation.")
        installPython3()
        fmt.Println("Python3 is already installed.")
    }

    // Check if Ansible is installed
    cmd = exec.Command("which", "ansible")
    err = cmd.Run()
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


func installPython3() {
    cmd := exec.Command("apt", "update")
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error updating packages.")
        os.Exit(1)
    }
    cmd = exec.Command("apt", "install", "-y", "python3")
    err = cmd.Run()
    if err != nil {
        fmt.Println("Error installing Python3.")
        os.Exit(1)
    }
    fmt.Println("Python3 installed.")
}

func installAnsible() {
    cmd := exec.Command("apt", "update")
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error updating packages.")
        os.Exit(1)
    }
    cmd = exec.Command("apt", "install", "-y", "ansible", "-qq")
    err = cmd.Run()
    if err != nil {
        fmt.Println("Error installing Ansible.")
        os.Exit(1)
    }
    fmt.Println("Ansible installed.")
}

func installProgram(program string) {
    // Update package cache
    cmd := exec.Command("ansible", "localhost", "-m", "apt", "-a", "update_cache=yes", "-vvvv")
    err := cmd.Run()
    if err != nil {
        fmt.Printf("Error updating package cache.\n")
        os.Exit(1)
    }

    // Install the program
    cmd = exec.Command("ansible", "localhost", "-m", "ansible.builtin.apt", "-a", fmt.Sprintf("name=%s state=latest", program), "-vvvv")
    err = cmd.Run()
    if err != nil {
        fmt.Printf("Error installing %s.\n", program)
        os.Exit(1)
    }

    // Start the program
    cmd = exec.Command("ansible", "localhost", "-m", "service", "-a", fmt.Sprintf("name=%s state=started", program), "-vvvv")
    err = cmd.Run()
    if err != nil {
        fmt.Printf("Error starting %s.\n", program)
        os.Exit(1)
    }

    fmt.Printf("%s installed and started.\n", program)
}
