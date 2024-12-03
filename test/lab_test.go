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
	err = client.Lab.UpdateLab("/", evengsdk.Lab{
		Name:        time.Format("15-04-05"),
		Description: "Updated Description",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestLabService_UpdateLabWithExtension(t *testing.T) {
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

func TestLabService_RenameLab(t *testing.T) {
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
	err = client.Lab.UpdateLab("/"+curtime.Format("15-04-05")+".unl", evengsdk.Lab{
		Name:        curtime.Format("15-04-05") + "-updated",
		Description: "Updated Description",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer client.Lab.DeleteLab("/" + curtime.Format("15-04-05") + "-updated" + ".unl")
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

func TestLabService_MoveLab(t *testing.T) {
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
	err = client.Folder.CreateFolder("/move-" + curtime.Format("15-04-05"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Folder.DeleteFolder("/move-" + curtime.Format("15-04-05"))
	err = client.Lab.MoveLab("/"+curtime.Format("15-04-05")+".unl", "/move-"+curtime.Format("15-04-05"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestLabService_MoveLabWithExtension(t *testing.T) {
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
	err = client.Folder.CreateFolder("/move-" + curtime.Format("15-04-05"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Folder.DeleteFolder("/move-" + curtime.Format("15-04-05"))
	err = client.Lab.MoveLab("/"+curtime.Format("15-04-05")+".unl", "/move-"+curtime.Format("15-04-05")+"/"+curtime.Format("15-04-05")+".unl")
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
