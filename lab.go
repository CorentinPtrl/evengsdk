package evengsdk

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"
)

type LabService struct {
	client *Client
}

type Lab struct {
	Path        string      `json:"path,omitempty"`
	Author      string      `json:"author"`
	Body        string      `json:"body"`
	Description string      `json:"description"`
	Filename    string      `json:"filename"`
	Id          string      `json:"id,omitempty"`
	Name        string      `json:"name"`
	Version     json.Number `json:"version"`
}

// GetLab returns the lab with the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *LabService) GetLab(path string) (*Lab, error) {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	eve, _, err := s.client.Do(context.Background(), "GET", "api/labs"+path+url.QueryEscape(name), nil)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(eve.Data)
	if err != nil {
		return nil, err
	}
	var lab Lab
	err = json.Unmarshal(data, &lab)
	if err != nil {
		return nil, err
	}
	return &lab, nil
}

// CreateLab creates a new lab in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl). or just the path to the folder.
func (s *LabService) CreateLab(path string, lab Lab) error {
	name := lab.Name
	if strings.Contains(path[strings.LastIndex(path, "/")+1:], ".unl") {
		name = path[strings.LastIndex(path, "/")+1:]
		name = name[:strings.LastIndex(name, ".")]
	}
	path = path[:strings.LastIndex(path, "/")+1]
	lab.Path = path
	lab.Name = name
	body, err := json.Marshal(lab)
	if err != nil {
		return err
	}
	_, _, err = s.client.Do(context.Background(), "POST", "api/labs", body)
	if err != nil {
		return err
	}
	return nil
}

// UpdateLab updates the lab with the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl). or just the path to the folder.
func (s *LabService) UpdateLab(path string, lab Lab) error {
	name := lab.Name
	if strings.Contains(path[strings.LastIndex(path, "/")+1:], ".unl") {
		name = path[strings.LastIndex(path, "/")+1:]
		name = name[:strings.LastIndex(name, ".")]
	}
	path = path[:strings.LastIndex(path, "/")+1]
	lab.Path = path
	body, err := json.Marshal(lab)
	if err != nil {
		return err
	}
	_, _, err = s.client.Do(context.Background(), "PUT", "api/labs"+path+url.QueryEscape(name)+".unl", body)
	if err != nil {
		return err
	}
	return nil
}

// DeleteLab deletes the lab with the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *LabService) DeleteLab(path string) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	_, _, err := s.client.Do(context.Background(), "DELETE", "api/labs"+path+url.QueryEscape(name), nil)
	if err != nil {
		return err
	}
	return nil
}

// MoveLab moves the lab with the specified path to the new path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
// The newPath should be the full path to the new location of the lab file, including the extension (e.g. /path/to/labfile.unl) or just the path to the folder.
func (s *LabService) MoveLab(path string, newPath string) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	if strings.Contains(newPath[strings.LastIndex(newPath, "/")+1:], ".unl") {
		newPath = newPath[:strings.LastIndex(newPath, "/")+1]
	}
	_, _, err := s.client.Do(context.Background(), "PUT", "api/labs"+path+url.QueryEscape(name)+"/move", []byte(`{"path":"`+newPath+`"}`))
	if err != nil {
		return err
	}
	return nil
}

// LockLab locks the lab with the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *LabService) LockLab(path string) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	_, _, err := s.client.Do(context.Background(), "PUT", "api/labs"+path+url.QueryEscape(name)+"/Lock", nil)
	if err != nil {
		return err
	}
	return nil
}

// UnlockLab unlocks the lab with the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *LabService) UnlockLab(path string) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	_, _, err := s.client.Do(context.Background(), "PUT", "api/labs"+path+url.QueryEscape(name)+"/Unlock", nil)
	if err != nil {
		return err
	}
	return nil
}

// GetTopology returns the topology of the lab with the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *LabService) GetTopology(path string) ([]map[string]interface{}, error) {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	eve, _, err := s.client.Do(context.Background(), "GET", "api/labs"+path+url.QueryEscape(name)+"/topology", nil)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(eve.Data)
	if err != nil {
		return nil, err
	}
	var topology []map[string]interface{}
	err = json.Unmarshal(data, &topology)
	if err != nil {
		return nil, err
	}
	return topology, nil
}

// CloseLab closes the lab for the current user.
func (s *LabService) CloseLab() error {
	_, _, err := s.client.Do(context.Background(), "DELETE", "api/labs/close", nil)
	if err != nil {
		return err
	}
	return nil
}
