/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/masum0813/getipinfo/networkdata/deviceinfo"

	"github.com/spf13/cobra"
)

var (
	strInternalDeviceFlag string
	// internalCmd represents the internal command
	internalCmd = &cobra.Command{
		Use:   "internal",
		Short: "Get internal ip information",
		Long: `Get internal ip information on your computer/systems if vpn adapter starts with "en"
For example:

getipinfo internal
getipinfo internal -d en0
`,
		Run: func(cmd *cobra.Command, args []string) {

			// var di deviceinfo.DeviceInfo
			// data := di.GetIpAddressFromInterfaceName()

			var retVal string

			di := deviceinfo.DeviceInfo{
				DeviceName: strInternalDeviceFlag,
			}

			datas := di.GetIpAddressFromInterfaceName()

			// Print internal ip address
			if len(datas) == 0 {
				retVal = fmt.Sprintf("Are you connected to the internet or internal adapter not starts with \"%s\" ?", strInternalDeviceFlag)
			} else {
				for _, data := range datas {
					// fmt.Println(data)
					var tmpData deviceinfo.AdapterDetails = deviceinfo.AdapterDetails(data)
					retVal += fmt.Sprintf("Ip address (%s): %s \n", tmpData.IfaceName, tmpData.IpAddress )	
				}
				
			}
			println(retVal)
		},
	}
)

func init() {
	rootCmd.AddCommand(internalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// internalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// internalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	internalCmd.Flags().StringVarP(&strInternalDeviceFlag, "device", "d", "en0", "Device name")

}
