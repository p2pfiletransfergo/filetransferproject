package main

import (
	"fmt"
	"os/exec"
	"strings"
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

	// If rule not found, Windows prints this:
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
