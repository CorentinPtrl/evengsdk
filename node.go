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

// InterfaceEntry can handle both slice and map structures.
type InterfaceEntry map[int]Interface

type Interfaces struct {
	Ethernet InterfaceEntry `json:"ethernet"`
	Serial   InterfaceEntry `json:"serial"`
}

type Style struct {
	Style           string      `json:"style"`
	Color           string      `json:"color"`
	Srcpos          float32     `json:"srcpos"`
	Dstpos          float32     `json:"dstpos"`
	Linkstyle       string      `json:"linkstyle"`
	Width           json.Number `json:"width"`
	Label           string      `json:"label"`
	Labelpos        float32     `json:"labelpos"`
	Stub            json.Number `json:"stub"`
	Curviness       json.Number `json:"curviness"`
	Beziercurviness json.Number `json:"beziercurviness"`
	Round           json.Number `json:"round"`
	Midpoint        float32     `json:"midpoint"`
	Id              string      `json:"id"`
	Node            string      `json:"node"`
	InterfaceId     string      `json:"interface_id"`
	Type            string      `json:"type"`
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

func (s *NodeService) startNodesCommunity(path string) error {
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

func (s *NodeService) startNodesPro(path string) error {
	nodes, err := s.GetNodes(path)
	if err != nil {
		return err
	}
	for _, node := range nodes {
		err = s.StartNode(path, node.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

// StartNodes starts all nodes in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) StartNodes(path string) error {
	if s.client.isPro {
		return s.startNodesPro(path)
	}
	return s.startNodesCommunity(path)
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

// UpdateNodeInterfaceStyle updates the style of the interface with the specified id of the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
// The style parameter should be a Style struct. The attributes Node and Type will be set automatically.
func (s *NodeService) UpdateNodeInterfaceStyle(path string, node int, style Style) error {
	if !s.client.isPro {
		return errors.New("This function is only available in the Pro version")
	}
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	style.Node = strconv.Itoa(node)
	style.Type = "ethernet"
	data, err := json.Marshal(style)
	if err != nil {
		return err
	}
	_, _, err = s.client.Do(context.Background(), "PUT", "api/labs/"+path+url.QueryEscape(name)+"/nodes/"+strconv.Itoa(node)+"/style", data)
	if err != nil {
		return err
	}
	return nil
}

// UpdateNodeInterfaceStyleByName updates the style of the interface with the specified name of the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
// The name should be the name of the interface (e.g. Gi0/0).
// The style parameter should be a Style struct. The attributes Node, Type, InterfaceId and Id will be set automatically.
func (s *NodeService) UpdateNodeInterfaceStyleByName(path string, node int, intf string, style Style) error {
	index, inter, err := s.GetNodeInterface(path, node, intf)
	if err != nil {
		return err
	}
	style.InterfaceId = strconv.Itoa(index)
	style.Id = "network_id:" + strconv.Itoa(inter.NetworkId)
	return s.UpdateNodeInterfaceStyle(path, node, style)
}

// GetNodeInterface returns the interface with the specified name of the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
// The name should be the name of the interface (e.g. Gi0/0).
// returns the index of the interface, the interface and an error.
func (s *NodeService) GetNodeInterface(path string, node int, intf string) (int, Interface, error) {
	interfaces, err := s.GetNodeInterfaces(path, node)
	if err != nil {
		return 0, Interface{}, err
	}
	for index, eth := range interfaces.Ethernet {
		if eth.Name == intf {
			return index, eth, nil
		}
	}
	return 0, Interface{}, errors.New("Interface not found")
}

// UpdateNodeInterfaceName updates the interface with the specified name of the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
// The name should be the name of the interface (e.g. Gi0/0).
func (s *NodeService) UpdateNodeInterfaceName(path string, node int, intf string, network int) error {
	index, _, err := s.GetNodeInterface(path, node, intf)
	if err != nil {
		return err
	}
	return s.UpdateNodeInterface(path, node, index, network)
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
	var eve *Response
	var err error
	if s.client.isPro {
		eve, _, err = s.client.Do(context.Background(), "POST", "api/labs/"+path+url.QueryEscape(name)+"/configs/"+strconv.Itoa(node), []byte("{\"cfsid\":\"default\"}"))
	} else {
		eve, _, err = s.client.Do(context.Background(), "GET", "api/labs/"+path+url.QueryEscape(name)+"/configs/"+strconv.Itoa(node), nil)
	}
	if err != nil {
		return "", err
	}
	if eve.Data == nil {
		return "", nil
	}
	return eve.Data.(map[string]interface{})["data"].(string), nil
}

// UpdateNodeConfig updates the config of the node with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NodeService) UpdateNodeConfig(path string, node int, config string) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	payload := map[string]string{"data": config}
	if s.client.isPro {
		payload["cfsid"] = "default"
	}
	data, err := json.Marshal(payload)
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

// some nodes have ethernet as map[string]interface{} instead of []interface{}, namely the IOL nodes
// this function will convert the map to a slice
func (e *InterfaceEntry) UnmarshalJSON(data []byte) error {
	var slice []Interface
	if err := json.Unmarshal(data, &slice); err == nil {
		*e = make(map[int]Interface, len(slice))
		for i, entry := range slice {
			(*e)[i] = entry
		}
		return nil
	}
	var m map[int]Interface
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	*e = m
	return nil
}
