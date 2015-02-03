package bridge

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	netctlConfigsPath = "/etc/netctl/"
)

func CreateBridge(name, bridgeInterface, description, ipType, ipAddr string,
	interfacesBindsTo, customOptions []string) error {

	if ipType != "dhcp" && ipType != "static" && ipType != "no" {
		return fmt.Errorf("Wrong IP type")
	}

	file, err := os.Create(netctlConfigsPath + name)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Fprintf(writer, "Description=\"%s\"\n", description)
	fmt.Fprintln(writer, "Interface=", bridgeInterface)
	fmt.Fprintln(writer, "Connection=bridge")
	fmt.Fprintf(writer, "BindsToInterfaces=(%s)\n",
		strings.Join(interfacesBindsTo, " "))
	fmt.Fprintln(writer, "IP=", ipType)
	if ipType == "static" {
		fmt.Fprintln(writer, "Address=", ipAddr)
	}
	for _, option := range customOptions {
		fmt.Fprintln(writer, option)
	}

	return writer.Flush()
}

func RemoveBridge(name string) error {
	return os.Remove(netctlConfigsPath + name)
}

func StartBridge(name string) error {
	cmd := exec.Command("netctl", "start", name)
	return cmd.Run()
}

func StopBridge(name string) error {
	cmd := exec.Command("netctl", "stop", name)
	return cmd.Run()
}

func IsBridgeExist(name string) bool {
	_, err := os.Stat(netctlConfigsPath + name)
	if err == nil {
		return true
	}
	return false
}

func AssignIpToBridge(ip, bridgeInterface string) error {
	cmd := exec.Command("ip", "addr", "add", "dev", bridgeInterface, ip)
	return cmd.Run()
}

func RemoveIpFromBridge(ip, bridgeInterface string) error {
	cmd := exec.Command("ip", "addr", "del", "dev", bridgeInterface, ip)
	return cmd.Run()
}
