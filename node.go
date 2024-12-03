package evengsdk

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type NodeService struct {
	client *Client
}

type Node struct {
	Console  string      `json:"console"`
	Delay    int         `json:"delay"`
	Id       int         `json:"id"`
	Left     int         `json:"left"`
	Icon     string      `json:"icon"`
	Image    string      `json:"image"`
	Name     string      `json:"name"`
	Ram      int         `json:"ram"`
	Status   int         `json:"status"`
	Template string      `json:"template"`
	Type     string      `json:"type"`
	Top      int         `json:"top"`
	Url      string      `json:"url"`
	Config   json.Number `json:"config"`
	Cpu      int         `json:"cpu"`
	Ethernet int         `json:"ethernet"`
	Uuid     string      `json:"uuid"`
}
type Interface struct {
	Name      string `json:"name"`
	NetworkId int    `json:"network_id"`
}

type Interfaces struct {
	Ethernet []Interface `json:"ethernet"`
	Serial   []Interface `json:"serial"`
}

// GetNodes returns all nodes in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) GetNodes(path string) (map[string]Node, error) {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	eve, _, err := s.client.Do(context.Background(), "GET", "api/labs/"+path+url.QueryEscape(name)+"/nodes", nil)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(eve.Data)
	if err != nil {
		return nil, err
	}
	var nodes map[string]Node
	err = json.Unmarshal(data, &nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

// GetNode returns the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) GetNode(path string, node int) (*Node, error) {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	eve, _, err := s.client.Do(context.Background(), "GET", "api/labs/"+path+url.QueryEscape(name)+"/nodes/"+strconv.Itoa(node), nil)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(eve.Data)
	if err != nil {
		return nil, err
	}
	var nodeConfig Node
	err = json.Unmarshal(data, &nodeConfig)
	if err != nil {
		return nil, err
	}
	nodeConfig.Id = node
	return &nodeConfig, nil
}

// CreateNode creates a new node in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
// The node should be a pointer to a Node struct. The Id field will be set to the id of the new node.
func (s *NodeService) CreateNode(path string, node *Node) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	body, err := json.Marshal(node)
	if err != nil {
		return err
	}
	resp, _, err := s.client.Do(context.Background(), "POST", "api/labs/"+path+url.QueryEscape(name)+"/nodes", body)
	if err != nil {
		return err
	}
	node.Id = int(resp.Data.(map[string]interface{})["id"].(float64))
	return nil
}

// UpdateNode updates the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) UpdateNode(path string, node *Node) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	body, err := json.Marshal(node)
	if err != nil {
		return err
	}
	_, _, err = s.client.Do(context.Background(), "PUT", "api/labs/"+path+url.QueryEscape(name)+"/nodes/"+strconv.Itoa(node.Id), body)
	if err != nil {
		return err
	}
	return nil
}

// DeleteNode deletes the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) DeleteNode(path string, nodeId int) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	_, _, err := s.client.Do(context.Background(), "DELETE", "api/labs/"+path+url.QueryEscape(name)+"/nodes/"+strconv.Itoa(nodeId), nil)
	if err != nil {
		return err
	}
	return nil
}

// StartNodes starts all nodes in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) StartNodes(path string) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	evengresp, _, err := s.client.Do(context.Background(), "GET", "api/labs/"+path[1:]+url.QueryEscape(name)+"/nodes/start", nil)
	if err != nil {
		return err
	}
	if evengresp.Status != "success" {
		return errors.New(evengresp.Message)
	}
	return nil
}

// GetNodeInterfaces returns all interfaces of the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) GetNodeInterfaces(path string, node int) (*Interfaces, error) {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	eve, _, err := s.client.Do(context.Background(), "GET", "api/labs/"+path+url.QueryEscape(name)+"/nodes/"+strconv.Itoa(node)+"/interfaces", nil)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(eve.Data)
	if err != nil {
		return nil, err
	}
	var nodeConfig Interfaces
	err = json.Unmarshal(data, &nodeConfig)
	if err != nil {
		return nil, err
	}
	return &nodeConfig, nil
}

// UpdateNodeInterface updates the interface with the specified id of the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) UpdateNodeInterface(path string, node int, intf int, network int) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	data, err := json.Marshal(map[string]interface{}{strconv.Itoa(intf): network})
	if network == 0 {
		data, err = json.Marshal(map[string]interface{}{strconv.Itoa(intf): ""})
	}
	if err != nil {
		return err
	}
	_, _, err = s.client.Do(context.Background(), "PUT", "api/labs/"+path+url.QueryEscape(name)+"/nodes/"+strconv.Itoa(node)+"/interfaces", data)
	if err != nil {
		return err
	}
	return nil
}

// StartNode starts the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) StartNode(path string, node int) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	evengresp, _, err := s.client.Do(context.Background(), "GET", "api/labs/"+path+url.QueryEscape(name)+"/nodes/"+strconv.Itoa(node)+"/start", nil)
	if err != nil {
		return err
	}
	if evengresp.Status != "success" {
		return errors.New(evengresp.Message)
	}
	return nil
}

// StopNodes stops all nodes in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) StopNodes(path string) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	evengresp, _, err := s.client.Do(context.Background(), "GET", "api/labs/"+path+url.QueryEscape(name)+"/nodes/stop", nil)
	if err != nil {
		return err
	}
	if evengresp.Status != "success" {
		return errors.New(evengresp.Message)
	}
	return nil
}

// StopNode stops the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) StopNode(path string, node int) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	evengresp, _, err := s.client.Do(context.Background(), "GET", "api/labs/"+path+url.QueryEscape(name)+"/nodes/"+strconv.Itoa(node)+"/stop", nil)
	if err != nil {
		return err
	}
	if evengresp.Status != "success" {
		return errors.New(evengresp.Message)
	}
	return nil
}

// GetNodeConfig returns the config of the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) GetNodeConfig(path string, node int) (string, error) {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	eve, _, err := s.client.Do(context.Background(), "GET", "api/labs/"+path+url.QueryEscape(name)+"/configs/"+strconv.Itoa(node), nil)
	if err != nil {
		return "", err
	}
	return eve.Data.(map[string]interface{})["data"].(string), nil
}

// UpdateNodeConfig updates the config of the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) UpdateNodeConfig(path string, node int, config string) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	data, err := json.Marshal(map[string]string{"data": config})
	if err != nil {
		return err
	}
	_, _, err = s.client.Do(context.Background(), "PUT", "api/labs/"+path+url.QueryEscape(name)+"/configs/"+strconv.Itoa(node), data)
	if err != nil {
		return err
	}
	return nil
}

// GetTemplates returns all templates.
func (s *NodeService) GetTemplates() (map[string]string, error) {
	eve, _, err := s.client.Do(context.Background(), "GET", "api/list/templates/", nil)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(eve.Data)
	if err != nil {
		return nil, err
	}
	var templates map[string]string
	err = json.Unmarshal(data, &templates)
	if err != nil {
		return nil, err
	}
	return templates, nil
}

func (s *NodeService) GetTemplate(name string) (map[string]interface{}, error) {
	eve, _, err := s.client.Do(context.Background(), "GET", "api/list/templates/"+name, nil)
	if err != nil {
		return nil, err
	}
	return eve.Data.(map[string]interface{}), nil
}
