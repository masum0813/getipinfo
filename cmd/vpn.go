/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/masum0813/getipinfo/networkdata/deviceinfo"

	"github.com/spf13/cobra"
)

// vpnCmd represents the vpn command
var (
	strVpnDeviceFlag string
	vpnCmd           = &cobra.Command{
		Use:   "vpn",
		Short: "Get vpn ip information",
		Long: `Get vpn ip information on your computer/systems if vpn adapter starts with "ppp"
For example:

getipinfo vpn
getipinfo vpn -d ppp0
`,
		Run: func(cmd *cobra.Command, args []string) {

			// var di deviceinfo.DeviceInfo
			// data := di.GetIpAddressFromInterfaceName()
			var retVal string
			di := deviceinfo.DeviceInfo{
				DeviceName: strVpnDeviceFlag,
			}
			datas := di.GetIpAddressFromInterfaceName()

			// Print internal ip address
			if len(datas) == 0 {
				retVal = fmt.Sprintf("Are you connected to the internet or internal adapter not starts with \"%s\" ?", strInternalDeviceFlag)
			} else {
				for _, data := range datas {
					// fmt.Println(data)
					var tmpData deviceinfo.AdapterDetails = deviceinfo.AdapterDetails(data)
					retVal += fmt.Sprintf("Ip address (%s): %s \n", tmpData.IfaceName, tmpData.IpAddress)
				}
			}
			println(retVal)
		},
	}
)

func init() {
	rootCmd.AddCommand(vpnCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vpnCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vpnCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	vpnCmd.Flags().StringVarP(&strVpnDeviceFlag, "device", "d", "ppp", "Device name")
}
