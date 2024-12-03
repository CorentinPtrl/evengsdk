package test

import (
	"github.com/CorentinPtrl/evengsdk"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestNetworkService_CreateNetwork(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	time := time.Now()
	err = client.Lab.CreateLab("/"+time.Format("15-04-05")+".unl", evengsdk.Lab{
		Description: "Unit Test",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Lab.DeleteLab("/" + time.Format("15-04-05") + ".unl")
	net := &evengsdk.Network{
		Left:       0,
		Top:        0,
		Name:       "Test",
		Type:       "bridge",
		Visibility: "1",
		Icon:       "lan.png",
	}
	err = client.Network.CreateNetwork("/"+time.Format("15-04-05")+".unl", net)
	if err != nil {
		t.Fatal(err)
	}
	err = client.Network.DeleteNetwork("/"+time.Format("15-04-05")+".unl", net.Id)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Network.GetNetwork("/"+time.Format("15-04-05")+".unl", net.Id)
	if err == nil {
		t.Fatal("Network was not deleted")
	}
}

func TestNetworkService_CreateManyNetworks(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	curtime := time.Now()
	err = client.Lab.CreateLab("/"+curtime.Format("15-04-05")+".unl", evengsdk.Lab{
		Description: "Unit Test",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Lab.DeleteLab("/" + curtime.Format("15-04-05") + ".unl")
	var networks []*evengsdk.Network
	for i := 0; i < 100; i++ {
		net := &evengsdk.Network{
			Left:       0,
			Top:        50 + i*50,
			Name:       "test_network_" + strconv.Itoa(i),
			Type:       "bridge",
			Visibility: "1",
			Icon:       "01-Cloud-Default.svg",
		}
		err = client.Network.CreateNetwork("/"+curtime.Format("15-04-05")+".unl", net)
		if err != nil {
			t.Fatal(err)
		}
		networks = append(networks, net)
	}
	for _, net := range networks {
		err = client.Network.DeleteNetwork("/"+curtime.Format("15-04-05")+".unl", net.Id)
		if err != nil {
			t.Fatal(err)
		}
		_, err = client.Network.GetNetwork("/"+curtime.Format("15-04-05")+".unl", net.Id)
		if err == nil {
			t.Fatal("Network was not deleted")
		}
	}
}

func TestNetworkService_NetworkVisibilityNoNodes(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	curtime := time.Now()
	err = client.Lab.CreateLab("/"+curtime.Format("15-04-05")+".unl", evengsdk.Lab{
		Description: "Unit Test",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Lab.DeleteLab("/" + curtime.Format("15-04-05") + ".unl")
	net := &evengsdk.Network{
		Left:       0,
		Top:        0,
		Name:       "Test",
		Type:       "bridge",
		Visibility: "1",
		Icon:       "lan.png",
	}
	err = client.Network.CreateNetwork("/"+curtime.Format("15-04-05")+".unl", net)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Network.GetNetwork("/"+curtime.Format("15-04-05")+".unl", net.Id)
	if err != nil {
		t.Fatal(err)
	}
	net.Visibility = "0"
	err = client.Network.UpdateNetwork("/"+curtime.Format("15-04-05")+".unl", net)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Network.GetNetwork("/"+curtime.Format("15-04-05")+".unl", net.Id)
	if err == nil {
		t.Fatal("Network was not deleted")
	}
}

func TestNetworkService_NetworkVisibilityWithNodes(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	curtime := time.Now()
	err = client.Lab.CreateLab("/"+curtime.Format("15-04-05")+".unl", evengsdk.Lab{
		Description: "Unit Test",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Lab.DeleteLab("/" + curtime.Format("15-04-05") + ".unl")
	net := &evengsdk.Network{
		Left:       0,
		Top:        0,
		Name:       "Test",
		Type:       "bridge",
		Visibility: "1",
		Icon:       "lan.png",
	}
	err = client.Network.CreateNetwork("/"+curtime.Format("15-04-05")+".unl", net)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Network.GetNetwork("/"+curtime.Format("15-04-05")+".unl", net.Id)
	if err != nil {
		t.Fatal(err)
	}
	node := &evengsdk.Node{
		Ethernet: 2,
		Name:     "vpc",
		Template: "vpcs",
		Type:     "qemu",
	}
	err = client.Node.CreateNode("/"+curtime.Format("15-04-05")+".unl", node)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Node.DeleteNode("/"+curtime.Format("15-04-05")+".unl", node.Id)
	_, err = client.Node.GetNode("/"+curtime.Format("15-04-05")+".unl", node.Id)
	if err != nil {
		t.Fatal(err)
	}
	err = client.Node.UpdateNodeInterface("/"+curtime.Format("15-04-05")+".unl", node.Id, 1, net.Id)
	if err != nil {
		t.Fatal(err)
	}
	net.Visibility = "0"
	err = client.Network.UpdateNetwork("/"+curtime.Format("15-04-05")+".unl", net)
	if err != nil {
		t.Fatal(err)
	}
	gnet, err := client.Network.GetNetwork("/"+curtime.Format("15-04-05")+".unl", net.Id)
	if err != nil {
		t.Fatal("Network was not deleted")
	}
	if gnet.Count != 1 {
		t.Fatal("Network count is not 1")
	}
	err = client.Network.DeleteNetwork("/"+curtime.Format("15-04-05")+".unl", net.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNetworkService_GetNetwork(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	time := time.Now()
	err = client.Lab.CreateLab("/"+time.Format("15-04-05")+".unl", evengsdk.Lab{
		Description: "Unit Test",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Lab.DeleteLab("/" + time.Format("15-04-05") + ".unl")
	net := &evengsdk.Network{
		Left:       0,
		Top:        0,
		Name:       "Test",
		Type:       "bridge",
		Visibility: "1",
		Icon:       "lan.png",
	}
	err = client.Network.CreateNetwork("/"+time.Format("15-04-05")+".unl", net)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Network.GetNetwork("/"+time.Format("15-04-05")+".unl", net.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNetworkService_GetNetworks(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	time := time.Now()
	err = client.Lab.CreateLab("/"+time.Format("15-04-05")+".unl", evengsdk.Lab{
		Description: "Unit Test",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Lab.DeleteLab("/" + time.Format("15-04-05") + ".unl")
	net := &evengsdk.Network{
		Left:       0,
		Top:        0,
		Name:       "Test",
		Type:       "bridge",
		Visibility: "1",
		Icon:       "lan.png",
	}
	err = client.Network.CreateNetwork("/"+time.Format("15-04-05")+".unl", net)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Network.GetNetworks("/" + time.Format("15-04-05") + ".unl")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNetworkService_UpdateNetwork(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	time := time.Now()
	err = client.Lab.CreateLab("/"+time.Format("15-04-05")+".unl", evengsdk.Lab{
		Description: "Unit Test",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Lab.DeleteLab("/" + time.Format("15-04-05") + ".unl")
	net := &evengsdk.Network{
		Left:       0,
		Top:        0,
		Name:       "Test",
		Type:       "bridge",
		Visibility: "1",
		Icon:       "lan.png",
	}
	err = client.Network.CreateNetwork("/"+time.Format("15-04-05")+".unl", net)
	if err != nil {
		t.Fatal(err)
	}
	net.Visibility = "0"
	err = client.Network.UpdateNetwork("/"+time.Format("15-04-05")+".unl", net)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNetworkService_DeleteNetwork(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	time := time.Now()
	err = client.Lab.CreateLab("/"+time.Format("15-04-05")+".unl", evengsdk.Lab{
		Description: "Unit Test",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Lab.DeleteLab("/" + time.Format("15-04-05") + ".unl")
	net := &evengsdk.Network{
		Left:       0,
		Top:        0,
		Name:       "Test",
		Type:       "bridge",
		Visibility: "1",
		Icon:       "lan.png",
	}
	err = client.Network.CreateNetwork("/"+time.Format("15-04-05")+".unl", net)
	if err != nil {
		t.Fatal(err)
	}
	err = client.Network.DeleteNetwork("/"+time.Format("15-04-05")+".unl", net.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNetworkService_GetNetworksList(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Network.GetNetworksList()
	if err != nil {
		t.Fatal(err)
	}
}
