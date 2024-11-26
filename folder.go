package evengsdk

import (
	"context"
	"encoding/json"
	"strings"
)

type FolderService struct {
	client *Client
}

type Folders struct {
	Folders []Folder `json:"folders"`
	Labs    []Lab    `json:"labs"`
}

type Folder struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type LabFolder struct {
	File   string `json:"file"`
	Path   string `json:"path"`
	Umtime int64  `json:"umtime"`
	Mtime  string `json:"mtime"`
}

// GetFolder returns a list of folders and labs in the specified folder.
// The root path is "/".
func (s *FolderService) GetFolder(path string) (*Folders, error) {
	eve, _, err := s.client.Do(context.Background(), "GET", "api/folders"+path, nil)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(eve.Data)
	if err != nil {
		return nil, err
	}
	var folder Folders
	err = json.Unmarshal(data, &folder)
	if err != nil {
		return nil, err
	}
	return &folder, nil
}

// CreateFolder creates a new folder in the specified path.
func (s *FolderService) CreateFolder(path string) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	folders := Folder{
		Name: name,
		Path: path,
	}
	body, err := json.Marshal(folders)
	if err != nil {
		return err
	}
	_, _, err = s.client.Do(context.Background(), "POST", "api/folders", body)
	if err != nil {
		return err
	}
	return nil
}

// UpdateFolder updates the specified folder.
// Keep in mind that only the path is required in the folder struct.
func (s *FolderService) UpdateFolder(path string, folder Folder) error {
	body, err := json.Marshal(folder)
	if err != nil {
		return err
	}
	_, _, err = s.client.Do(context.Background(), "PUT", "api/folders/"+path, body)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFolder deletes the specified folder.
func (s *FolderService) DeleteFolder(path string) error {
	_, _, err := s.client.Do(context.Background(), "DELETE", "api/folders"+path, nil)
	if err != nil {
		return err
	}
	return nil
}
