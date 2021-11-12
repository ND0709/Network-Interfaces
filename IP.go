package main

import(
	"fmt"
	"log"
	"net"
	"strings"
)
func Ip(int){

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddress: %v\n", err.Error()))
		return
	}

	for _, i := range ifaces {
		addrs,err := i.Addrs()
		if err != nil{
			log.Print(fmt.Errorf("localAddress: %v\n", err.Error()))
			continue
		}
		for _, a := range addrs {
			log.Printf("Index:%d MTU:%d InterfaceName:%v Hardwareaddress:%v %v\n", i.Index, i.MTU,i.Name, i.HardwareAddr,a)
			if strings.Contains(i.Flags.String(),"up") {
				fmt.Print("Flag UP")
				fmt.Print("loopback")
			} else {
				fmt.Println("Flag:Down")
			}
		}
	}
}
func main() {
	Ip(1)
}

