/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/masum0813/ipinfo/networkdata/deviceinfo"

	"github.com/spf13/cobra"
)

// vpnCmd represents the vpn command
var vpnCmd = &cobra.Command{
	Use:   "vpn",
	Short: "Get vpn ip information",
	Long: `Get vpn ip information on your computer/systems if vpn adapter starts with "ppp"
For example:

ipinfo vpn`,
	Run: func(cmd *cobra.Command, args []string) {

		// var di deviceinfo.DeviceInfo
		// data := di.GetIpAddressFromInterfaceName()
		var retVal string
		di := deviceinfo.DeviceInfo{
			DeviceName: "ppp0",
		}
		data := di.GetIpAddressFromInterfaceName()

		// Print internal ip address
		if data == "" {
			retVal = "Are you connected to the VPN or Vpn adapter not starts with \"ppp\" ?"
		} else {
			retVal = fmt.Sprintf("VPN ip address: %s", data)
		}
		println(retVal)
	},
}

func init() {
	rootCmd.AddCommand(vpnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vpnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vpnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
