/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import "github.com/masum0813/getipinfo/cmd"

var (
	// VERSION is set during build
	VERSION = "0.1.0"
)

func main() {
	cmd.Execute(VERSION)
}
