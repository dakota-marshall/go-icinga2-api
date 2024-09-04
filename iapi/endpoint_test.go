package iapi

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateEndpointRawDataDefaultValues(t *testing.T) {

	name := "endpoint-raw-data-default"
	host := "127.0.0.2"
	port := 5665
	// packageName := "endpoint-raw-data-default-endpoint"
	logDuration := "1d"
	expectedRawData := fmt.Sprintf("object Zone \"%s\" { parent = \"master\", endpoints = [ \"%s\" ] }\n object Endpoint \"%s\" { host = \"%s\", port = \"%s\", log_duration = \"%s\" }",
		name,
		name,
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
	logDuration := "40000"
	expectedRawData := fmt.Sprintf("object Zone \"%s\" { parent = \"master\", endpoints = [ \"%s\" ] }\n object Endpoint \"%s\" { host = \"%s\", port = \"%s\", log_duration = \"%s\" }",
		name,
		name,
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
func TestGetValidEndpoint(t *testing.T) {

	name := "endpoint-raw-data-custom-values"
	packageName := "endpoint-raw-data-custom-test"

	_, err := Icinga2_Server.GetEndpoint(name, packageName)
	if err != nil {
		t.Error(err)
	}
}

func TestGetInvalidEndpoint(t *testing.T) {

	name := "c2-mysql-1"
	packageName := "bad-package"

	_, err := Icinga2_Server.GetEndpoint(name, packageName)
	if err == nil {
		t.Error(err)
	}
}

func TestDeleteEndpoint(t *testing.T) {

	packageName := "test-endpoint-package2"

	err := Icinga2_Server.DeletePackage(packageName)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(15 * time.Second)
}
