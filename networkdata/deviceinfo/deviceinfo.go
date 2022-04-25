package deviceinfo

import (
	"fmt"
	"net"
	"strings"
)

type DeviceInfo struct {
	DeviceName string
}

func (deviceInfo DeviceInfo) GetIpAddressFromInterfaceName() string {

	// Get internal ip address
	var retVal string = ""
	// var deviceinfo DeviceInfo
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			println(err)
		}

		if strings.HasPrefix(iface.Name, deviceInfo.DeviceName) {

			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				if ip == nil || ip.IsLoopback() {
					continue
				}
				ip = ip.To4()
				if ip == nil {
					continue // not an ipv4 address
				}

				retVal = fmt.Sprintf("%s", ip.String())
			}
		}
	}

	// deviceinfo.ReturnValue = retVal
	return retVal
}
