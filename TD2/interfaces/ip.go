package main

import (
	"fmt"
	//"strconv"
)
type IPAddr [4]byte

func (hosts IPAddr) String() string {
	var a string
	for i := range hosts {
		//a += fmt.Sprintf("%d.", hosts[i])
		//a += strconv.Itoa("%d.", hosts[i])
	}
	return fmt.Sprintf(a[:len(a)-1])
}


func main() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}