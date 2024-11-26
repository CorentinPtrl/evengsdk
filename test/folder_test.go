package test

import (
	"github.com/CorentinPtrl/evengsdk"
	"os"
	"testing"
	"time"
)

func TestFolderService_GetFolder(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Folder.GetFolder("/")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFolderService_CreateFolder(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	currentTime := time.Now()
	err = client.Folder.CreateFolder("/" + currentTime.Format("15-04-05"))
	if err != nil {
		t.Fatal(err)
	}
	client.Folder.DeleteFolder("/" + currentTime.Format("15-04-05"))
}

func TestFolderService_UpdateFolder(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	currentTime := time.Now()
	err = client.Folder.CreateFolder("/" + currentTime.Format("15-04-05"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Folder.DeleteFolder("/Updated")
	err = client.Folder.UpdateFolder("/"+currentTime.Format("15-04-05"), evengsdk.Folder{
		Name: "",
		Path: "/Updated",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestFolderService_DeleteFolder(t *testing.T) {
	client, err := evengsdk.NewBasicAuthClient(os.Getenv("EVE_USER"), os.Getenv("EVE_PASSWORD"), "0", os.Getenv("EVE_HOST"))
	if err != nil {
		t.Fatal(err)
	}
	currentTime := time.Now()
	err = client.Folder.CreateFolder("/" + currentTime.Format("15-04-05"))
	if err != nil {
		t.Fatal(err)
	}
	err = client.Folder.DeleteFolder("/" + currentTime.Format("15-04-05"))
	if err != nil {
		t.Fatal(err)
	}
}
