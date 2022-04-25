/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	externalip "github.com/glendc/go-external-ip"
	"github.com/spf13/cobra"
)

// CLI Flags
var (
	timeout  = flag.Duration("t", time.Second*5, "consensus's voting timeout")
	verbose  = flag.Bool("v", false, "log errors to STDERR, when defined")
	protocol = flag.Uint("p", 0, "IP Protocol to be used (0, 4, or 6)")
)

// externalCmd represents the external command
var externalCmd = &cobra.Command{
	Use:   "external",
	Short: "Get external ip information",
	Long: `Get internal ip information on your computer/systems.
For example:

ipinfo external`,
	Run: func(cmd *cobra.Command, args []string) {

		var retVal string

		// configure the consensus
		cfg := externalip.DefaultConsensusConfig()
		if timeout != nil {
			cfg.WithTimeout(*timeout)
		}

		// optionally create the logger,
		// if no logger is defined, all logs will be discarded.
		var logger *log.Logger
		if verbose != nil && *verbose {
			logger = externalip.NewLogger(os.Stderr)
		}

		// create the consensus
		consensus := externalip.DefaultConsensus(cfg, logger)
		err := consensus.UseIPProtocol(*protocol)
		errCheck(err)

		// retrieve the external ip
		ip, err := consensus.ExternalIP()
		errCheck(err)

		// success, simply output the IP in string format
		// fmt.Println(ip.String())
		// Print internal ip address
		if ip.String() == "" {
			retVal = "Are you connected to the internet or internal adapter not starts with \"en\" ?"
		} else {
			retVal = fmt.Sprintf("External ip address: %s", ip.String())
		}
		println(retVal)
	},
}

func errCheck(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(externalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// externalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// externalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
