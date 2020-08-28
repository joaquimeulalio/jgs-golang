package main

import (
	//	"bytes"
	//	"encoding/json"
	"fmt"
	//	"log"
	//	"os/exec"
	//	"strings"

	udev "github.com/jochenvg/go-udev"
)

func test() {
	fmt.Printf(" testing: %v\n", " v1.00")
	u := udev.Udev{}
	e := u.NewEnumerate()

	e.AddMatchSubsystem("block")
	e.AddMatchIsInitialized()

	n := 0

	alldevices, _ := e.Devices()
	for j := range alldevices {
		sdevice := alldevices[j]
		if len(sdevice.PropertyValue("ID_VENDOR_ID")) > 0 ||
			len(sdevice.PropertyValue("ID_MODEL")) > 0 ||
			len(sdevice.PropertyValue("ID_SERIAL_SHORT")) > 0 ||
			len(sdevice.PropertyValue("ID_REVISION")) > 0 {

			fmt.Println(n)
			n++
			fmt.Println("sdevice[ID_VENDOR_ID]", sdevice.PropertyValue("ID_VENDOR_ID"))
			fmt.Println("sdevice[ID_MODEL]", sdevice.PropertyValue("ID_MODEL"))
			fmt.Println("sdevice[ID_SERIAL_SHORT]", sdevice.PropertyValue("ID_SERIAL_SHORT"))
			fmt.Println("sdevice[ID_REVISION]", sdevice.PropertyValue("ID_REVISION"))
			fmt.Println("")
		}
	}

}

func main() {
	test()
	return
}
