package main

//commands in this files work windows only for now but can i exspanded for linux and mac as well
import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

func ruleExists(name string) bool {

	cmd := exec.Command(
		"netsh",
		"advfirewall",
		"firewall",
		"show",
		"rule",
		"name="+name,
	)

	out, _ := cmd.CombinedOutput()

	if strings.Contains(string(out), "No rules match") {
		return false
	}

	return true
}
func rules() {
	if ruleExists("UDP8080") {
		fmt.Println("Exists")
	} else {
		cmd := exec.Command("netsh",
			"advfirewall",
			"firewall",
			"add",
			"rule",
			`name=UDP8080`,
			"dir=in",
			"action=allow",
			"protocol=UDP",
			"localport=8080")
		op, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}
		fmt.Println(op)
	}
	if ruleExists("UDP8080OUT") {
		fmt.Println("Exists")
		return
	}
	cmd2 := exec.Command("netsh",
		"advfirewall",
		"firewall",
		"add",
		"rule",
		`name=UDP8080OUT`,
		"dir=out",
		"action=allow",
		"protocol=UDP",
		"localport=8080")
	op1, err := cmd2.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(op1)
}

var (
	kernel32                    = syscall.NewLazyDLL("kernel32.dll")
	setThreadExecutionStateProc = kernel32.NewProc("SetThreadExecutionState")
)

const (
	ES_CONTINUOUS      = 0x80000000
	ES_SYSTEM_REQUIRED = 0x00000001
)

func PreventSleep() {
	setThreadExecutionStateProc.Call(
		uintptr(ES_CONTINUOUS | ES_SYSTEM_REQUIRED),
	)
}

func AllowSleep() {
	setThreadExecutionStateProc.Call(
		uintptr(ES_CONTINUOUS),
	)
}
