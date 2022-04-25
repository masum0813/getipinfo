/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import "ipinfo/cmd"

var (
	// VERSION is set during build
	VERSION = "0.0.5"
)

func main() {
	cmd.Execute(VERSION)
}
