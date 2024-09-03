package iapi

import (
	"testing"
	"time"
)

var packageName string
var packageStageName string
var configFilePath string
var configData string

func TestCreatePackageStage(t *testing.T) {

	packageName = "test-stage-package"
	configFilePath = "conf.d/test-host.conf"
	configData = "object Host \"local-host\" { address = \"127.0.0.1\", check_command = \"hostalive\" }"

	// Create package for stage testing
	_, err := Icinga2_Server.CreatePackage(packageName)
	if err != nil {
		t.Error(err)
	}

	pkgStageResult, err := Icinga2_Server.CreatePackageStage(packageName, configFilePath, configData)
	if err != nil {
		t.Error(err)
	}
	packageStageName = pkgStageResult[0].Stage

	// Sleep to allow time for icinga to reload
	time.Sleep(15 * time.Second)

	pkg, err := Icinga2_Server.GetPackage(packageName)
	if err != nil {
		t.Error(err)
	}

	if pkgStageResult[0].Stage != pkg.ActiveStage {
		t.Error("New Package is not the current stage")
	}
}
func TestCreatePackageStageError(t *testing.T) {

	packageNameError := "test-stage-package-error"
	configFilePathError := "conf.d/test-host.conf"
	configDataError := "objec Host \"local-host\" { address = \"127.0.0.1\", check_command = \"hostalive\" }"

	// Create package for stage testing
	_, err := Icinga2_Server.CreatePackage(packageNameError)
	if err != nil {
		t.Error(err)
	}

	pkgStageResult, err := Icinga2_Server.CreatePackageStage(packageNameError, configFilePathError, configDataError)
	if err != nil {
		t.Error(err)
	}

	// Sleep to allow time for icinga to reload
	time.Sleep(15 * time.Second)

	pkg, err := Icinga2_Server.GetPackage(packageNameError)
	if err != nil {
		t.Error(err)
	}

	if pkgStageResult[0].Stage == pkg.ActiveStage {
		t.Error("New Package the current stage")
	}
}

func TestGetPackageStage(t *testing.T) {

	pkgStage, err := Icinga2_Server.GetPackageStage(packageName, packageStageName)
	if err != nil {
		t.Error(err)
	}
	if len(pkgStage) < 1 {
		t.Error("No objects in slice")
	}

	var pathExists bool
	pathExists = false
	for _, file := range pkgStage {
		if file.Name == configFilePath {
			pathExists = true
		}
	}
	if !pathExists {
		t.Error("Couldnt find config file path in specified stage")
	}

}
func TestGetPackageStageBadStage(t *testing.T) {

	pkgStage, err := Icinga2_Server.GetPackageStage(packageName, "bad-package-stage")
	if err == nil {
		t.Error(err)
	}
	if len(pkgStage) > 1 {
		t.Error("Too many objects in slice")
	}
}
func TestGetPackageStageInvalidCharacter(t *testing.T) {

	pkgStage, err := Icinga2_Server.GetPackageStage(packageName, "stage&&$$!!__")
	if err == nil {
		t.Error(err)
	}
	if len(pkgStage) > 1 {
		t.Error("Too many objects in slice")
	}
}
