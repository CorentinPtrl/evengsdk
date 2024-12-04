package test

import (
	"github.com/CorentinPtrl/evengsdk"
	"os"
	"testing"
	"time"
)

func TestNodeService_CreateNode(t *testing.T) {
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
	node := &evengsdk.Node{
		Cpu:      1,
		Delay:    0,
		Ethernet: 8,
		Image:    "viosl2-adventerprisek9-m.03.2017",
		Left:     0,
		Name:     "Switch_Test",
		Ram:      1024,
		Template: "viosl2",
		Top:      0,
		Config:   "1",
		Type:     "qemu",
	}
	err = client.Node.CreateNode("/"+time.Format("15-04-05")+".unl", node)
	if err != nil {
		t.Fatal(err)
	}
	client.Node.DeleteNode("/"+time.Format("15-04-05")+".unl", node.Id)
	client.Lab.DeleteLab("/" + time.Format("15-04-05") + ".unl")
}

func TestNodeService_CreateVPCNode(t *testing.T) {
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
	node := &evengsdk.Node{
		Name:     "vpc",
		Template: "vpcs",
		Type:     "qemu",
	}
	err = client.Node.CreateNode("/"+time.Format("15-04-05")+".unl", node)
	if err != nil {
		t.Fatal(err)
	}
	client.Node.DeleteNode("/"+time.Format("15-04-05")+".unl", node.Id)
	client.Lab.DeleteLab("/" + time.Format("15-04-05") + ".unl")
}

func TestNodeService_GetNode(t *testing.T) {
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
	node := &evengsdk.Node{
		Cpu:      1,
		Delay:    0,
		Ethernet: 8,
		Image:    "viosl2-adventerprisek9-m.03.2017",
		Left:     0,
		Name:     "Switch_Test",
		Ram:      1024,
		Template: "viosl2",
		Top:      0,
		Config:   "1",
		Type:     "qemu",
	}
	err = client.Node.CreateNode("/"+time.Format("15-04-05")+".unl", node)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Node.GetNode("/"+time.Format("15-04-05")+".unl", node.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNodeService_UpdateNode(t *testing.T) {
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
	node := &evengsdk.Node{
		Cpu:      1,
		Delay:    0,
		Ethernet: 8,
		Image:    "viosl2-adventerprisek9-m.03.2017",
		Left:     0,
		Name:     "Switch_Test",
		Ram:      1024,
		Template: "viosl2",
		Top:      0,
		Config:   "1",
		Type:     "qemu",
	}
	err = client.Node.CreateNode("/"+time.Format("15-04-05")+".unl", node)
	if err != nil {
		t.Fatal(err)
	}
	node.Name = "Switch_Test_Updated"
	err = client.Node.UpdateNode("/"+time.Format("15-04-05")+".unl", node)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNodeService_DeleteNode(t *testing.T) {
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
	node := &evengsdk.Node{
		Cpu:      1,
		Delay:    0,
		Ethernet: 8,
		Image:    "viosl2-adventerprisek9-m.03.2017",
		Left:     0,
		Name:     "Switch_Test",
		Ram:      1024,
		Template: "viosl2",
		Top:      0,
		Config:   "1",
		Type:     "qemu",
	}
	err = client.Node.CreateNode("/"+time.Format("15-04-05")+".unl", node)
	if err != nil {
		t.Fatal(err)
	}
	err = client.Node.DeleteNode("/"+time.Format("15-04-05")+".unl", node.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNodeService_GetNodeConfig(t *testing.T) {
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
	node := &evengsdk.Node{
		Cpu:      1,
		Delay:    0,
		Ethernet: 8,
		Image:    "viosl2-adventerprisek9-m.03.2017",
		Left:     0,
		Name:     "Switch_Test",
		Ram:      1024,
		Template: "viosl2",
		Top:      0,
		Config:   "1",
		Type:     "qemu",
	}
	err = client.Node.CreateNode("/"+time.Format("15-04-05")+".unl", node)
	if err != nil {
		t.Fatal(err)
	}
	err = client.Node.UpdateNodeConfig("/"+time.Format("15-04-05")+".unl", node.Id, "hostname Switch_Test")
	if err != nil {
		t.Fatal(err)
	}
	config, err := client.Node.GetNodeConfig("/"+time.Format("15-04-05")+".unl", node.Id)
	if err != nil {
		t.Fatal(err)
	}
	if config != "hostname Switch_Test" {
		t.Fatal("Config is not correct")
	}
	client.Node.DeleteNode("/"+time.Format("15-04-05")+".unl", node.Id)
}

func TestNodeService_UpdateNodeConfig(t *testing.T) {
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
	node := &evengsdk.Node{
		Cpu:      1,
		Delay:    0,
		Ethernet: 8,
		Image:    "viosl2-adventerprisek9-m.03.2017",
		Left:     0,
		Name:     "Switch_Test",
		Ram:      1024,
		Template: "viosl2",
		Top:      0,
		Config:   "1",
		Type:     "qemu",
	}
	err = client.Node.CreateNode("/"+time.Format("15-04-05")+".unl", node)
	if err != nil {
		t.Fatal(err)
	}
	err = client.Node.UpdateNodeConfig("/"+time.Format("15-04-05")+".unl", node.Id, "hostname Switch_Test")
	if err != nil {
		t.Fatal(err)
	}
	client.Node.DeleteNode("/"+time.Format("15-04-05")+".unl", node.Id)
}

func TestNodeService_GetInvalidNodeConfig(t *testing.T) {
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
	_, err = client.Node.GetNodeConfig("/"+time.Format("15-04-05")+".unl", 0)
	if err == nil {
		t.Fatal("Should have failed")
	}
}
func TestNodeService_GetNodeInterfaces(t *testing.T) {
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
	node := &evengsdk.Node{
		Cpu:      1,
		Delay:    0,
		Ethernet: 8,
		Image:    "viosl2-adventerprisek9-m.03.2017",
		Left:     0,
		Name:     "Switch_Test",
		Ram:      1024,
		Template: "viosl2",
		Top:      0,
		Config:   "1",
		Type:     "qemu",
	}
	err = client.Node.CreateNode("/"+time.Format("15-04-05")+".unl", node)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Node.GetNodeInterfaces("/"+time.Format("15-04-05")+".unl", node.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNodeService_UpdateNodeInterface(t *testing.T) {
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
	node := &evengsdk.Node{
		Cpu:      1,
		Delay:    0,
		Ethernet: 8,
		Image:    "viosl2-adventerprisek9-m.03.2017",
		Left:     0,
		Name:     "Switch_Test",
		Ram:      1024,
		Template: "viosl2",
		Top:      0,
		Config:   "1",
		Type:     "qemu",
	}
	err = client.Node.CreateNode("/"+time.Format("15-04-05")+".unl", node)
	if err != nil {
		t.Fatal(err)
	}
	network := &evengsdk.Network{
		Left:       0,
		Top:        0,
		Name:       "Test",
		Type:       "bridge",
		Visibility: "1",
		Icon:       "lan.png",
	}
	err = client.Network.CreateNetwork("/"+time.Format("15-04-05")+".unl", network)
	if err != nil {
		t.Fatal(err)
	}
	err = client.Node.UpdateNodeInterface("/"+time.Format("15-04-05")+".unl", node.Id, 1, network.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNodeService_GetNodes(t *testing.T) {
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
	node := &evengsdk.Node{
		Cpu:      1,
		Delay:    0,
		Ethernet: 8,
		Image:    "viosl2-adventerprisek9-m.03.2017",
		Left:     0,
		Name:     "Switch_Test",
		Ram:      1024,
		Template: "viosl2",
		Top:      0,
		Config:   "1",
		Type:     "qemu",
	}
	err = client.Node.CreateNode("/"+time.Format("15-04-05")+".unl", node)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Node.GetNodes("/" + time.Format("15-04-05") + ".unl")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNodeService_GetTemplates(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Node.GetTemplates()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNodeService_GetTemplate(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Node.GetTemplate("vpcs")
	if err != nil {
		t.Fatal(err)
	}
}
