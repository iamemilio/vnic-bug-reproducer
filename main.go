package main

import (
	"fmt"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/portsbinding"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/gophercloud/utils/openstack/clientconfig"
)

func main() {

	client, err := clientconfig.NewServiceClient("network", nil)
	if err != nil {
		panic(err)
	}

	portCreateOpts := ports.CreateOpts{
		Name:      "portsbindingtest00",
		NetworkID: "149fd037-f905-455c-bbc2-7ec14341679e",
	}

	createOpts := portsbinding.CreateOptsExt{
		CreateOptsBuilder: portCreateOpts,
		VNICType:          "direct",
	}

	port, err := ports.Create(client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Println(port)
}
