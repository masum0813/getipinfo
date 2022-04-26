/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/masum0813/ipinfo/networkdata/deviceinfo"

	"github.com/spf13/cobra"
)

// internalCmd represents the internal command
var internalCmd = &cobra.Command{
	Use:   "internal",
	Short: "Get internal ip information",
	Long: `Get internal ip information on your computer/systems if vpn adapter starts with "en"
For example:

ipinfo internal`,
	Run: func(cmd *cobra.Command, args []string) {

		// var di deviceinfo.DeviceInfo
		// data := di.GetIpAddressFromInterfaceName()
		var retVal string
		di := deviceinfo.DeviceInfo{
			DeviceName: "en0",
		}
		data := di.GetIpAddressFromInterfaceName()

		// Print internal ip address
		if data == "" {
			retVal = "Are you connected to the internet or internal adapter not starts with \"en\" ?"
		} else {
			retVal = fmt.Sprintf("VPN ip address: %s", data)
		}
		println(retVal)
	},
}

func init() {
	rootCmd.AddCommand(internalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// internalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// internalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
