package test

import (
	"github.com/CorentinPtrl/evengsdk"
	"os"
	"testing"
	"time"
)

func TestLabService_CreateLab(t *testing.T) {
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
	client.Lab.DeleteLab("/" + time.Format("15-04-05") + ".unl")
}

func TestLabService_GetLab(t *testing.T) {
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
	_, err = client.Lab.GetLab("/" + time.Format("15-04-05") + ".unl")
	if err != nil {
		t.Fatal(err)
	}
}

func TestLabService_UpdateLab(t *testing.T) {
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
	err = client.Lab.UpdateLab("/"+time.Format("15-04-05")+".unl", evengsdk.Lab{
		Description: "Updated Description",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestLabService_DeleteLab(t *testing.T) {
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
	err = client.Lab.DeleteLab("/" + time.Format("15-04-05") + ".unl")
	if err != nil {
		t.Fatal(err)
	}
}

func TestLabService_LockLab(t *testing.T) {
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
	err = client.Lab.LockLab("/" + time.Format("15-04-05") + ".unl")
	if err != nil {
		t.Fatal(err)
	}
}

func TestLabService_UnlockLab(t *testing.T) {
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
	err = client.Lab.LockLab("/" + time.Format("15-04-05") + ".unl")
	if err != nil {
		t.Fatal(err)
	}
	err = client.Lab.UnlockLab("/" + time.Format("15-04-05") + ".unl")
	if err != nil {
		t.Fatal(err)
	}
	client.Lab.DeleteLab("/" + time.Format("15-04-05") + ".unl")
}

func TestLabService_GetTopology(t *testing.T) {
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
	//TODO: Should populate the lab
	_, err = client.Lab.GetTopology("/" + time.Format("15-04-05") + ".unl")
	if err != nil {
		t.Fatal(err)
	}
}
