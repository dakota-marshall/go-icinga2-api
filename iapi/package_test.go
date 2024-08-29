package iapi

import (
	"testing"
)

func TestGetPackage(t *testing.T) {

	pkgName := "_api"

	pkg, err := Icinga2_Server.GetPackage(pkgName)
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

func TestCreatePackage(t *testing.T) {

	pkgName := "test-package"

	pkg, err := Icinga2_Server.CreatePackage(pkgName)
	if err != nil {
		t.Error(err)
	}

	if pkg.Name != pkgName {
		t.Error("Package name does not match requested package")
	}
}

func TestDeletePackage(t *testing.T) {

	pkgName := "test-package"

	err := Icinga2_Server.DeletePackage(pkgName)
	if err != nil {
		t.Error(err)
	}
}
