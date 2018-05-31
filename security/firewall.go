package security

import (
	"os/exec"
	"runtime"
	"strings"

	"github.com/senorprogrammer/wtf/wtf"
)

const osxFirewallCmd = "/usr/libexec/ApplicationFirewall/socketfilterfw"

/* -------------------- Exported Functions -------------------- */

func firewallStateLinux() string {
	return "[red]NA[white]"
}

func firewallStateMacOS() string {
	cmd := exec.Command(osxFirewallCmd, "--getglobalstate")
	str := wtf.ExecuteCommand(cmd)

	return status(str)
}

func FirewallState() string {
	switch runtime.GOOS {
	case "linux":
		return firewallStateLinux()
	case "macos":
		return firewallStateMacOS()
	default:
		return ""
	}
}

func firewallStealthStateLinux() string {
	return "[red]NA[white]"
}

func firewallStealthStateMacOS() string {
	cmd := exec.Command(osxFirewallCmd, "--getstealthmode")
	str := wtf.ExecuteCommand(cmd)

	return status(str)
}

func FirewallStealthState() string {
	switch runtime.GOOS {
	case "linux":
		return firewallStealthStateLinux()
	case "macos":
		return firewallStealthStateMacOS()
	default:
		return ""
	}
}

/* -------------------- Unexported Functions -------------------- */

func status(str string) string {
	icon := "[red]off[white]"

	if strings.Contains(str, "enabled") {
		icon = "[green]on[white]"
	}

	return icon
}
