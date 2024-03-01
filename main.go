package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	os_release, _ := ioutil.ReadFile("/etc/os-release")

	strs := string(os_release)
	ubuntu := strings.Contains(strs, "ubuntu")
	if ubuntu {
		fmt.Println("ubuntu")
	} else {
		fmt.Println("not found")
	}
}
