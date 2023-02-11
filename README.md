# Easy Package Installer/Updater for Linux (Ubuntu + Ansible)

This is a simple CLI utility that helps install and keep specific packages up-to-date on Ubuntu systems. The utility checks for updates to the specified package and automates the process of updating it to the latest version.

This is a simple utility written in Go that can be used to install or update packages on a system. It uses Ansible as the underlying technology for installing packages. If Ansible is not already installed on the system, the utility provides the option to install it.

## Note

- This utility is designed now to run on Ubuntu distributives only.
- This utility install the LATEST versions of packages.

## Features

- Installs the specified package if it's not already installed
- Checks for updates to the specified package
- Automates the process of updating the package to the latest version
- Installs Ansible if it's not already installed

## Requirements

- Go 1.13+
- Ubuntu

## Getting started

```shell
curl -o easy-package-installer https://github.com/sv222/easy-installer-updater-cli/releases/download/v0.1.0/easy-package-installer && mv easy-package-installer /usr/local/bin
```

## Usage

To install or update a package, simply run the utility with the name of the package as the last argument.

For example, to install the package `nginx`:

```shell
easy-package-installer nginx
```

### Build from source

1.Clone the repository:

```shell
https://github.com/sv222/easy-installer-updater-cli.git
```

2.Build & Install the binary using the Makefile:

```shell
make build && make install
```

Note:
You can specify cli name in Makefile
BINARY_NAME=easy-package-installer

3.Now you can run the utility to install or update a package.

## Docker Usage

### Build the Docker Image

```shell
docker build -t app .
```

### Run the Container

```shell
docker run app [package-name]
```

## Contributing

Feel free to contribute to this project by submitting pull requests or reporting issues.

## License

This project is licensed under the MIT License.
