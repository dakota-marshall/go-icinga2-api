// This should create a package object on icinga
// This should be defined in order to deploy packages and zones
package iapi

import (
	"encoding/json"
	"fmt"
)

const packageEndpoint = "/config/packages"

func (server *Server) GetPackage(packageName string) (PackageStruct, error) {

	var allPackages []PackageStruct
	var correctPackage PackageStruct

	results, err := server.NewAPIRequest("GET", packageEndpoint+"/", nil)
	if err != nil {
		return PackageStruct{}, err
	}

	// Contents of the results is an interface object. Need to convert it to json first.
	jsonStr, marshalErr := json.Marshal(results.Results)
	if marshalErr != nil {
		return PackageStruct{}, marshalErr
	}

	// then the JSON can be pushed into the appropriate struct.
	// Note : Results is a slice so much push into a slice.

	unmarshalErr := json.Unmarshal(jsonStr, &allPackages)
	if unmarshalErr != nil {
		return PackageStruct{}, unmarshalErr
	}

	// Endpoint only ever returns all packages, so get the package we actually need
	for _, pkg := range allPackages {
		if pkg.Name == packageName {
			correctPackage = pkg
		}
	}

	return correctPackage, err

}
