package iapi

import (
	"encoding/json"
	"fmt"
)

////////////////////////////////////////////////////////
// Endpoints are not exposed via the IcingaAPI.       //
// This will likely require utilization of the        //
// Packages endpoint to deploy these via config files //
////////////////////////////////////////////////////////

const endpointEndpoint = "/objects/endpoint"

func (server *Server) GetEndpoint(endpoint string) ([]EndpointStruct, error) {

	var endpoints []EndpointStruct

	results, err := server.NewAPIRequest("GET", endpointEndpoint+"/"+endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Contents of the results is an interface object. Need to convert it to json first.
	jsonStr, marshalErr := json.Marshal(results.Results)
	if marshalErr != nil {
		return nil, marshalErr
	}

	// then the JSON can be pushed into the appropriate struct.
	// Note : Results is a slice so much push into a slice.

	if unmarshalErr := json.Unmarshal(jsonStr, &endpoints); unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return endpoints, err

}

// Create Endpoint ...
func (server *Server) CreateEndpoint(name string, host string, port int, log_duration string) ([]EndpointStruct, error) {

	var newAttrs EndpointAttrs
	newAttrs.Host = host
	newAttrs.Port = port
	newAttrs.LogDuration = log_duration

	// LogDuration is optional, set to default of 1d if blank string is passed
	if log_duration == "" {
		newAttrs.LogDuration = "1d"
	} else {
		newAttrs.LogDuration = log_duration
	}

	var newEndpoint EndpointStruct
	newEndpoint.Name = name
	newEndpoint.Type = "Endpoint"
	newEndpoint.Attrs = newAttrs

	// Create JSON from completed struct
	payloadJSON, marshalErr := json.Marshal(newEndpoint)
	if marshalErr != nil {
		return nil, marshalErr
	}

	// Make the API request to create the endpoint.
	results, err := server.NewAPIRequest("PUT", endpointEndpoint+"/"+name, []byte(payloadJSON))
	if err != nil {
		return nil, err
	}

	if results.Code == 200 {
		hosts, err := server.GetEndpoint(name)
		return hosts, err
	}

	return nil, fmt.Errorf("%s", results.ErrorString)

}

// DeleteHost ...
func (server *Server) DeleteEndpoint(name string) error {
	results, err := server.NewAPIRequest("DELETE", endpointEndpoint+"/"+name+"?cascade=1", nil)
	if err != nil {
		return err
	}

	if results.Code == 200 {
		return nil
	}

	return fmt.Errorf("%s", results.ErrorString)
}
