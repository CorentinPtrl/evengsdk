package test

import (
	"github.com/CorentinPtrl/evengsdk"
	"os"
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
