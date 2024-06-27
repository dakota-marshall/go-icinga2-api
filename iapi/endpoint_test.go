package iapi

import "testing"

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

func TestCreateEndpoint(t *testing.T) {

	name := "go-icinga2-api-1"
	host := "127.0.0.2"
	port := 5665
	log_duration := ""

	_, err := Icinga2_Server.CreateEndpoint(name, host, port, log_duration)

	if err != nil {
		t.Error(err)
	}
}

func TestDeleteEndpoint(t *testing.T) {

	name := "go-icinga2-api-1"

	err := Icinga2_Server.DeleteEndpoint(name)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteEndpointDNE(t *testing.T) {
	name := "go-icinga2-api-1"
	err := Icinga2_Server.DeleteEndpoint(name)
	if err.Error() != "No objects found." {
		t.Error(err)
	}
}
