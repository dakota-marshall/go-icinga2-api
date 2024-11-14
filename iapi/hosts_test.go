package iapi

import "testing"

func TestGetValidHost(t *testing.T) {

	hostname := "c1-mysql-1"

	_, err := Icinga2_Server.GetHost(hostname)

	if err != nil {
		t.Error(err)
	}
}

func TestGetInvalidHost(t *testing.T) {

	hostname := "c2-mysql-1"
	_, err := Icinga2_Server.GetHost(hostname)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateSimpleHost(t *testing.T) {

	hostname := "go-icinga2-api-1"
	IPAddress := "127.0.0.2"
	CheckCommand := "hostalive"

	_, err := Icinga2_Server.CreateHost(hostname, IPAddress, "", CheckCommand, nil, nil, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestCreateSimpleIPv6Host(t *testing.T) {

	hostname := "go-icinga2-api-3"
	IPAddress := "127.0.0.2"
	IP6Address := "::1"
	CheckCommand := "hostalive"

	_, err := Icinga2_Server.CreateHost(hostname, IPAddress, IP6Address, CheckCommand, nil, nil, nil)

	if err != nil {
		t.Error(err)
	}
}

func TestCreateHostWithVariables(t *testing.T) {

	hostname := "go-icinga2-api-2"
	IPAddress := "127.0.0.3"
	CheckCommand := "hostalive"

	variables := make(map[string]interface{})

	variables["vars.os"] = "Linux"
	variables["vars.creator"] = "Terraform"
	variables["vars.urls"] = []string{"test-url1.example.com", "test-url2.example.com"}

	_, err := Icinga2_Server.CreateHost(hostname, IPAddress, "", CheckCommand, variables, nil, nil)
	if err != nil {
		t.Error(err)
	}

	// Delete host after creating it.
	deleteErr := Icinga2_Server.DeleteHost(hostname)
	if deleteErr != nil {
		t.Error(err)
	}
}

func TestUpdateHost(t *testing.T) {

	hostname := "go-icinga2-api-5"
	IPAddress := "127.0.0.4"
	CheckCommand := "hostalive"

	variables := make(map[string]interface{})

	variables["vars.os"] = "Linux"
	variables["vars.creator"] = "Terraform"
	variables["vars.urls"] = []string{"test-url2.example.com", "test-url3.example.com"}

	_, err := Icinga2_Server.CreateHost(hostname, IPAddress, "", CheckCommand, variables, nil, nil)
	if err != nil {
		t.Error(err)
	}

	variables["vars.urls"] = []string{"test-url4.example.com", "test-url5.example.com"}

	updateHost := HostStruct{
		Name: hostname,
		Type: "Host",
		Attrs: HostAttrs{
			ActionURL:    "",
			Address:      "127.0.0.1",
			Address6:     "::1",
			CheckCommand: "hostalive",
			DisplayName:  "go-icinga2-api-5",
			Groups:       []string{},
			Notes:        "",
			NotesURL:     "",
			Templates:    []string{},
			Vars:         variables,
		},
	}

	_, errUpdate := Icinga2_Server.UpdateHost(updateHost)
	if errUpdate != nil {
		t.Error(errUpdate)
	}

	// Delete host after creating it.
	deleteErr := Icinga2_Server.DeleteHost(hostname)
	if deleteErr != nil {
		t.Error(deleteErr)
	}

}

func TestCreateHostWithTemplates(t *testing.T) {
	hostname := "go-icinga2-api-2"
	IPAddress := "127.0.0.3"
	CheckCommand := "hostalive"

	templates := []string{"template1", "template2"}

	_, err := Icinga2_Server.CreateHost(hostname, IPAddress, "", CheckCommand, nil, templates, nil)
	if err != nil {
		t.Error(err)
	}

	// Delete host after creating it.
	deleteErr := Icinga2_Server.DeleteHost(hostname)
	if deleteErr != nil {
		t.Error(err)
	}
}

func TestCreateHostWithGroup(t *testing.T) {
	hostname := "go-icinga2-api-2"
	IPAddress := "127.0.0.3"
	CheckCommand := "hostalive"
	Group := []string{"linux-servers"}

	_, err := Icinga2_Server.CreateHost(hostname, IPAddress, "", CheckCommand, nil, nil, Group)
	if err != nil {
		t.Error(err)
	}

	// Delete host after creating it.
	deleteErr := Icinga2_Server.DeleteHost(hostname)
	if deleteErr != nil {
		t.Error(err)
	}
}
func TestDeleteHost(t *testing.T) {

	hostname := "go-icinga2-api-1"

	err := Icinga2_Server.DeleteHost(hostname)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteHostDNE(t *testing.T) {
	hostname := "go-icinga2-api-1"
	err := Icinga2_Server.DeleteHost(hostname)
	if err.Error() != "No objects found." {
		t.Error(err)
	}
}
