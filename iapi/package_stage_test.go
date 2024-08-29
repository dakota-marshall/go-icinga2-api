package iapi

import (
	"testing"
)

func TestGetPackageStage(t *testing.T) {

	// pkgName := "test-package-stage"

	// pkg, err := Icinga2_Server.GetPackage(pkgName)
	// if err != nil {
	// 	t.Error(err)
	// }
	// if pkg.ActiveStage == "" {
	// 	t.Error("Failed to get ActiveStage from package")
	// }
	// if pkg.Name == "" {
	// 	t.Error("Failed to get Name from package")
	// }
}

func TestCreatePackageStage(t *testing.T) {

	pkgName := "test-package-stage"
	configFilePath := "conf.d/test-host.conf"
	configData := "object Host \"local-host\" { address = \"127.0.0.1\", check_command = \"hostalive\" }"

	// Create package for stage testing
	_, err := Icinga2_Server.CreatePackage(pkgName)
	if err != nil {
		t.Error(err)
	}

	_, err = Icinga2_Server.CreatePackageStage(pkgName, configFilePath, configData)
	if err != nil {
		t.Error(err)
	}

	_, err = Icinga2_Server.GetPackage(pkgName)
	if err != nil {
		t.Error(err)
	}
	// TODO: Need to figure out how to ensure the config we upload becomes the active stage
	// if config[0].Stage != pkg.ActiveStage {
	// 	t.Error("New Package is not the current stage")
	// }
}

func TestDeletePackageStage(t *testing.T) {

	pkgName := "test-package-stage"

	// Delete test package
	err := Icinga2_Server.DeletePackage(pkgName)
	if err != nil {
		t.Error(err)
	}
}
