package bridge

import (
	"os/exec"
)

func CreateBridge(name string) error {
	cmd := exec.Command("ip", "link", "add", "name", name, "type", "bridge")
	return cmd.Run()
}

func RemoveBridge(name string) error {
	cmd := exec.Command("ip", "link", "delete", name, "type", "bridge")
	return cmd.Run()
}

func StartBridge(name string) error {
	cmd := exec.Command("ip", "link", "set", "dev", name, "up")
	return cmd.Run()
}

func StopBridge(name string) error {
	cmd := exec.Command("ip", "link", "set", "dev", name, "down")
	return cmd.Run()
}

func IsBridgeExist(name string) bool {
	cmd := exec.Command("ip", "link", "show", "dev", name)
	err := cmd.Run()
	if err == nil {
		return true
	}
	return false
}

func AssignIpToBridge(ip, bridgeName string) error {
	cmd := exec.Command("ip", "addr", "add", "dev", bridgeName, ip)
	return cmd.Run()
}

func RemoveIpFromBridge(ip, bridgeName string) error {
	cmd := exec.Command("ip", "addr", "del", "dev", bridgeName, ip)
	return cmd.Run()
}
