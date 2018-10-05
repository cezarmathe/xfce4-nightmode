package main

import (
	// formatting
	"fmt"
	// os
	"os"
	// command execution
	"os/exec"
)

var NightTheme string = "Arc-Dark"
var NightWmTheme string = "Arc-Dark"

var DayTheme string = "Arc"
var DayWmTheme string = "Arc"

func main() {

	nm := os.Args[1]

	if nm == "on" {

		cmd1 := exec.Command("xfconf-query", "-c", "xsettings", "-p", "/Net/ThemeName", "-s", NightTheme)
		err1 := cmd1.Run()
		if err1 != nil {
			os.Stderr.WriteString(err1.Error())
			panic(err1)
		}

		cmd2 := exec.Command("xfconf-query", "-c", "xfwm4", "-p", "/general/theme", "-s", NightWmTheme)
		err2 := cmd2.Run()
		if err2 != nil {
			os.Stderr.WriteString(err2.Error())
			panic(err2)
		}

		fmt.Println("Night mode on.")
		os.Exit(0)
	}

	if nm == "off" {

		cmd1 := exec.Command("xfconf-query", "-c", "xsettings", "-p", "/Net/ThemeName", "-s", DayTheme)
		err1 := cmd1.Run()
		if err1 != nil {
			os.Stderr.WriteString(err1.Error())
			panic(err1)
		}

		cmd2 := exec.Command("xfconf-query", "-c", "xfwm4", "-p", "/general/theme", "-s", DayWmTheme)
		err2 := cmd2.Run()
		if err2 != nil {
			os.Stderr.WriteString(err2.Error())
			panic(err2)
		}

		fmt.Println("Night mode on.")
		os.Exit(0)
	}

	fmt.Println("Unknown option: " + nm)
	os.Exit(1)
}