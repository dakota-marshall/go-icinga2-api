package iapi

import (
	"fmt"
	"strconv"
)

// Endpoints are created via the Package API, these are configured on the server via static config files

func (server *Server) GetEndpoint(endpointName string, packageName string) (EndpointStruct, error) {

	var endpoint EndpointStruct
	endpoint.Name = endpointName
	endpoint.Path = "conf.d/" + endpointName + ".conf"
	// endpoint.Attrs = ""

	// Get package data
	pkgResult, err := server.GetPackage(packageName)
	if err != nil {
		return EndpointStruct{}, err
	}

	endpoint.Package = pkgResult
	endpoint.Stage = endpoint.Package.ActiveStage

	// Check if path is in package active stage
	stageResult, err := server.GetPackageStage(packageName, endpoint.Stage)
	if err != nil {
		return EndpointStruct{}, err
	}
	var pathExists bool
	for _, file := range stageResult {
		if endpoint.Path == file.Name {
			pathExists = true
		}
	}
	if !pathExists {
		return EndpointStruct{}, fmt.Errorf("path " + endpoint.Path + " does not exist in provided package")
	}

	// Get rawdata from package stage
	fileResult, err := server.GetPackageStageFile(packageName, endpoint.Package.ActiveStage, endpoint.Path)
	if err != nil {
		return EndpointStruct{}, err
	}
	endpoint.RawData = fileResult

	return endpoint, err

}

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
	newEndpoint.RawData = fmt.Sprintf("object Zone \"%s\" { parent = \"master\", endpoints = [ \"%s\" ] }\n object Endpoint \"%s\" { host = \"%s\", port = \"%s\", log_duration = \"%s\" }",
		newEndpoint.Name,
		newEndpoint.Name,
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
