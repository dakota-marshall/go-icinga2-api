// This should create a package object on icinga
// This should be defined in order to deploy packages and zones
package iapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

const packageStageEndpoint = "/config/stages"
const packageStageFileEndpoint = "/config/files"

func (server *Server) GetPackageStage(packageName string, packageStageName string) ([]PackageStageFile, error) {

	var packageData []PackageStageFile

	results, err := server.NewAPIRequest("GET", packageStageEndpoint+"/"+packageName+"/"+packageStageName, nil)
	if err != nil {
		return []PackageStageFile{}, err
	}
	if results.Code != 200 {
		return []PackageStageFile{}, errors.New(results.ErrorString)
	}

	// Contents of the results is an interface object. Need to convert it to json first.
	jsonStr, marshalErr := json.Marshal(results.Results)
	if marshalErr != nil {
		return []PackageStageFile{}, marshalErr
	}

	// then the JSON can be pushed into the appropriate struct.
	// Note : Results is a slice so much push into a slice.

	unmarshalErr := json.Unmarshal(jsonStr, &packageData)
	if unmarshalErr != nil {
		return []PackageStageFile{}, unmarshalErr
	}

	return packageData, err

}

// Create Package ...
func (server *Server) CreatePackageStage(pkgName string, configFilePath string, configData string) ([]PackageStageCreateResult, error) {

	newPackageStage := map[string]map[string]string{
		"files": {
			configFilePath: configData,
		},
	}
	var packageResult []PackageStageCreateResult

	// Create JSON from completed struct
	payloadJSON, marshalErr := json.Marshal(newPackageStage)
	if marshalErr != nil {
		return nil, marshalErr
	}

	// Make the API request to upload the config.
	results, err := server.NewAPIRequest("POST", packageStageEndpoint+"/"+pkgName, []byte(payloadJSON))
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
	if unmarshalErr := json.Unmarshal(jsonStr, &packageResult); unmarshalErr != nil {
		return nil, unmarshalErr
	}

	if results.Code == 200 {
		return packageResult, err
	}

	return packageResult, fmt.Errorf("%s", results.ErrorString)

}

func (server *Server) GetPackageStageFile(packageName string, packageStageName string, filePath string) (string, error) {

	results, err := server.NewFileRequest("GET", packageStageFileEndpoint+"/"+packageName+"/"+packageStageName+"/"+filePath, nil)
	if err != nil {
		return "", err
	}
	if results.Code != 200 {
		return "", errors.New(results.ErrorString)
	}

	return results.Result, err

}
