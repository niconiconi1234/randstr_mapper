package main

import (
	"fmt"
	"github.com/kubeedge/mappers-go/mapper-sdk-go/pkg/models"
	"github.com/kubeedge/mappers-go/mapper-sdk-go/pkg/service"
	"github.com/niconiconi1234/randstr_mapper/device"
)

func main() {
	fmt.Println("randstr mapper start")
	var randStrDevice models.ProtocolDriver = &device.RandStrDevice{}
	service.Bootstrap("randstr-virtual-protocol", randStrDevice)
}
