package iapi

import (
	"testing"
)

func TestCreatePackage(t *testing.T) {

	package_name := "_api"

	pkg, err := Icinga2_Server.GetPackage(package_name)
	if err != nil {
		t.Error(err)
	}
	if pkg.ActiveStage == "" {
		t.Error("Failed to get ActiveStage from package")
	}
	if pkg.Name == "" {
		t.Error("Failed to get Name from package")
	}
}

// func TestCreatePackage(t *testing.T) {
//
// 	name := "test-package"
// 	host := "127.0.0.2"
// 	port := 5665
//
// 	_, err := Icinga2_Server.Create(name, host, port, log_duration)
//
// 	if err != nil {
// 		t.Error(err)
// 	}
// }
//
// func TestDeletePackage(t *testing.T) {
//
// 	name := "go-icinga2-api-1"
//
// 	err := Icinga2_Server.DeleteEndpoint(name)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }
