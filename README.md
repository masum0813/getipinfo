# Get IP Address with Golang

[![Go](https://github.com/masum0813/getipinfo/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/masum0813/getipinfo/actions/workflows/go.yml)

Get IP address(es) with go-lang is a simple command line tool to get your IP address vpn, internal, external, etc.

## Usage

```bash
ipinfo -h
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  ipinfo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  external    Get external ip information
  help        Help about any command
  internal    Get internal ip information
  version     Show ipinfo version information
  vpn         Get vpn ip information

Flags:
  -h, --help     help for ipinfo
  -t, --toggle   Help message for toggle

Use "ipinfo [command] --help" for more information about a command.
```

External IP:

```bash
ipinfo external
External ip address: xx.xx.xx.xx
```

Internal IP:

```bash
ipinfo internal
VPN ip address: 192.168.1.49
```

VPN IP:

```bash
ipinfo vpn
VPN ip address: xx.xx.xx.xx
```

* How to install ipinfo

```bash
go install github.com/masum0813/ipinfo@latest
```

* Install Cobra[https://github.com/spf13/cobra]

```bash
go get -u github.com/spf13/cobra@latest
```

```bash
go install github.com/spf13/cobra-cli@latest
```

## Cobra Generator

<https://github.com/spf13/cobra-cli/blob/main/README.md>

<https://cobra.dev>

## cobra-cli usage

```bash
cobra-cli init
go run main.go
cobra-cli add internal
```

### References

<https://github.com/GlenDC/go-external-ip>

<https://dev.to/koddr/github-action-for-release-your-go-projects-as-fast-and-easily-as-possible-20a2>
