package iapi

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// func TestGetValidEndpoint(t *testing.T) {
//
// 	name := "c1-mysql-1"
//
// 	_, err := Icinga2_Server.GetEndpoint(name)
//
// 	if err != nil {
// 		t.Error(err)
// 	}
// }
//
// func TestGetInvalidEndpoint(t *testing.T) {
//
// 	name := "c2-mysql-1"
// 	_, err := Icinga2_Server.GetEndpoint(name)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

func TestCreateEndpointRawDataDefaultValues(t *testing.T) {

	name := "endpoint-raw-data-default"
	host := "127.0.0.2"
	port := 5665
	// packageName := "endpoint-raw-data-default-endpoint"
	logDuration := "1d"
	expectedRawData := fmt.Sprintf("object Endpoint \"%s\" { host = \"%s\", port = \"%s\", log_duration = \"%s\" }",
		name,
		host,
		strconv.Itoa(port),
		logDuration,
	)

	endpointResult, err := Icinga2_Server.CreateEndpoint(name, host, port, "", "")
	if err != nil {
		t.Error(err)
	}

	// Allow icinga time to reload
	time.Sleep(15 * time.Second)

	// Check Raw Data
	if endpointResult.RawData != expectedRawData {
		fmt.Println(endpointResult.RawData)
		fmt.Println(expectedRawData)
		t.Error("Expected rawData not correct")
	}
}
func TestCreateEndpointRawDataCustomValues(t *testing.T) {

	name := "endpoint-raw-data-custom-values"
	host := "127.0.0.2"
	port := 5665
	packageName := "endpoint-raw-data-custom-test"
	logDuration := "3d"
	expectedRawData := fmt.Sprintf("object Endpoint \"%s\" { host = \"%s\", port = \"%s\", log_duration = \"%s\" }",
		name,
		host,
		strconv.Itoa(port),
		logDuration,
	)

	endpointResult, err := Icinga2_Server.CreateEndpoint(name, host, port, logDuration, packageName)
	if err != nil {
		t.Error(err)
	}

	// Allow icinga time to reload
	time.Sleep(15 * time.Second)

	// Check Raw Data
	if endpointResult.RawData != expectedRawData {
		fmt.Println(endpointResult.RawData)
		fmt.Println(expectedRawData)
		t.Error("Expected rawData not correct")
	}
}
func TestCreateEndpointVerifyStage(t *testing.T) {

	name := "endpoint-verify-stage"
	host := "127.0.0.2"
	port := 5665
	packageName := "test-endpoint-package2"
	log_duration := ""

	endpointResult, err := Icinga2_Server.CreateEndpoint(name, host, port, log_duration, packageName)
	if err != nil {
		t.Error(err)
	}

	// Allow icinga time to reload
	time.Sleep(15 * time.Second)

	// Get package and ensure the stage is the expected stage
	pkgResult, err := Icinga2_Server.GetPackage(packageName)
	if err != nil {
		t.Error(err)
	}

	if pkgResult.ActiveStage != endpointResult.Package.ActiveStage {
		t.Error(err)
	}

}

//
// func TestDeleteEndpoint(t *testing.T) {
//
// 	name := "go-icinga2-api-1"
//
// 	err := Icinga2_Server.DeleteEndpoint(name)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }
//
// func TestDeleteEndpointDNE(t *testing.T) {
// 	name := "go-icinga2-api-1"
// 	err := Icinga2_Server.DeleteEndpoint(name)
// 	if err.Error() != "No objects found." {
// 		t.Error(err)
// 	}
// }
