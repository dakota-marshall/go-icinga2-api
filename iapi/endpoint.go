package iapi

import (
	"fmt"
	"strconv"
)

// func (server *Server) GetEndpoint(endpoint string) ([]EndpointStruct, error) {
//
// 	var endpoints []EndpointStruct
//
// 	// Contents of the results is an interface object. Need to convert it to json first.
// 	jsonStr, marshalErr := json.Marshal(results.Results)
// 	if marshalErr != nil {
// 		return nil, marshalErr
// 	}
//
// 	// then the JSON can be pushed into the appropriate struct.
// 	// Note : Results is a slice so much push into a slice.
//
// 	if unmarshalErr := json.Unmarshal(jsonStr, &endpoints); unmarshalErr != nil {
// 		return nil, unmarshalErr
// 	}
//
// 	return endpoints, err
//
// }

// Create Endpoint ...
func (server *Server) CreateEndpoint(name string, host string, port int, logDuration string, packageName string) (EndpointStruct, error) {

	// Endpoint Config
	var newAttrs EndpointAttrs
	newAttrs.Host = host
	newAttrs.Port = port

	// LogDuration is optional, set to default of 1d if blank string is passed
	if logDuration == "" {
		newAttrs.LogDuration = "1d"
	} else {
		newAttrs.LogDuration = logDuration
	}

	// Package Name is optional, set to default if blank
	if packageName == "" {
		packageName = name + "-endpoint"
	}

	var newEndpoint EndpointStruct
	newEndpoint.Name = name
	newEndpoint.Attrs = newAttrs
	newEndpoint.Path = "conf.d/" + name + ".conf"

	// Create config file from attrs
	newEndpoint.RawData = fmt.Sprintf("object Endpoint \"%s\" { host = \"%s\", port = \"%s\", log_duration = \"%s\" }",
		newEndpoint.Name,
		newEndpoint.Attrs.Host,
		strconv.Itoa(newEndpoint.Attrs.Port),
		newEndpoint.Attrs.LogDuration,
	)

	// Create package for endpoint config
	pkgResult, err := server.CreatePackage(packageName)
	if err != nil {
		return EndpointStruct{}, err
	}
	newEndpoint.Package = pkgResult

	// Create config
	stageResult, err := server.CreatePackageStage(packageName, newEndpoint.Path, newEndpoint.RawData)
	if err != nil {
		return EndpointStruct{}, err
	}
	if stageResult[0].Code == 200 {
		// Update Package
		newEndpoint.Stage = stageResult[0].Stage
		newEndpoint.Package, err = server.GetPackage(packageName)
		if err != nil {
			return EndpointStruct{}, err
		}

		return newEndpoint, nil
	}

	return EndpointStruct{}, fmt.Errorf("%s", stageResult[0].Status)

}

// DeleteHost ...
// func (server *Server) DeleteEndpoint(name string) error {
// 	results, err := server.NewAPIRequest("DELETE", endpointEndpoint+"/"+name+"?cascade=1", nil)
// 	if err != nil {
// 		return err
// 	}
//
// 	if results.Code == 200 {
// 		return nil
// 	}
//
// 	return fmt.Errorf("%s", results.ErrorString)
// }
