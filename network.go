package evengsdk

import (
	"context"
	"encoding/json"
	"maps"
	"net/url"
	"slices"
	"strconv"
	"strings"
)

type NetworkService struct {
	client *Client
}

type Network struct {
	Id         int         `json:"id"`
	Count      int         `json:"count"`
	Left       int         `json:"left"`
	Name       string      `json:"name"`
	Top        int         `json:"top"`
	Type       string      `json:"type"`
	Visibility json.Number `json:"visibility"`
	Icon       string      `json:"icon"`
}

// GetNetworks returns all networks in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NetworkService) GetNetworks(path string) (map[string]Network, error) {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	eve, _, err := s.client.Do(context.Background(), "GET", "api/labs/"+path+url.QueryEscape(name)+"/networks", nil)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(eve.Data)
	if err != nil {
		return nil, err
	}
	var networks map[string]Network
	err = json.Unmarshal(data, &networks)
	if err != nil {
		return nil, err
	}
	return networks, nil
}

// GetNetwork returns the network with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NetworkService) GetNetwork(path string, id int) (*Network, error) {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	eve, _, err := s.client.Do(context.Background(), "GET", "api/labs/"+path+url.QueryEscape(name)+"/networks/"+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(eve.Data)
	if err != nil {
		return nil, err
	}
	var network Network
	err = json.Unmarshal(data, &network)
	if err != nil {
		return nil, err
	}
	return &network, nil
}

// CreateNetwork creates a new network in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
// The network parameter should be a pointer to a Network struct. The Id field will be set to the id of the new network.
func (s *NetworkService) CreateNetwork(path string, network *Network) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	data, err := json.Marshal(network)
	if err != nil {
		return err
	}
	eve, _, err := s.client.Do(context.Background(), "POST", "api/labs/"+path+url.QueryEscape(name)+"/networks", data)
	if err != nil {
		return err
	}
	network.Id = int(eve.Data.(map[string]interface{})["id"].(float64))
	return nil
}

// UpdateNetwork updates the network with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NetworkService) UpdateNetwork(path string, network *Network) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	data, err := json.Marshal(network)
	if err != nil {
		return err
	}
	_, _, err = s.client.Do(context.Background(), "PUT", "api/labs/"+path+url.QueryEscape(name)+"/networks/"+strconv.Itoa(network.Id), data)
	return err
}

// DeleteNetwork deletes the network with the specified id in the specified path.
// The path should be the full path to the lab file, including the extension (e.g. /path/to/labfile.unl).
func (s *NetworkService) DeleteNetwork(path string, id int) error {
	name := path[strings.LastIndex(path, "/")+1:]
	path = path[:strings.LastIndex(path, "/")+1]
	_, _, err := s.client.Do(context.Background(), "DELETE", "api/labs/"+path+url.QueryEscape(name)+"/networks/"+strconv.Itoa(id), nil)
	return err
}

// GetNetworksList returns a list of all networks.
// This is a convenience method that returns the names of all networks available in the system.
func (s *NetworkService) GetNetworksList() ([]string, error) {
	eve, _, err := s.client.Do(context.Background(), "GET", "api/list/networks", nil)
	if err != nil {
		return nil, err
	}
	vals := slices.Collect(maps.Keys(eve.Data.(map[string]interface{})))
	return vals, nil
}
