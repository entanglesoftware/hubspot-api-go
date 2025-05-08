package users_test

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/users"

	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/discovery/crm"
)

func TestGetUsers(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	// Initialize a variable of type Users
	user := users.GetUsersParams{
		Properties: &[]string{"hs_job_title", "hs_availability_status", "hs_working_hours"},
	}

	ct := crm.Users()
	response, err := ct.GetUsersWithResponse(context.Background(), &user)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Results == nil {
			t.Fatalf("Response contains no results")
		}

		for _, result := range response.JSON200.Results {
			t.Logf("%+v\n", result)
			t.Log("-----")
			if result.Properties != nil {
				for key, value := range result.Properties {
					t.Logf("Key: %s, Value: %+v\n", key, value)
				}
			} else {
				t.Log("No properties found.")
			}
			t.Log("-----")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

func TestGetUser(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	userId := "389641148436"

	// Initialize a variable of type Users
	user := users.GetUserByIdParams{
		Properties: &[]string{"hs_job_title", "hs_availability_status", "hs_working_hours"},
	}

	ct := crm.Users()
	response, err := ct.GetUserByIdWithResponse(context.Background(), userId, &user)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Id == "" {
			t.Fatalf("Response contains no results")
		}

		if response.JSON200.Properties != nil {
			t.Logf("Properties: %s", response.JSON200.Properties)
		} else {
			t.Log("No properties found.")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

func TestSaveUsers(t *testing.T) {
	// Fetch the access token from the environment
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	ct := crm.Users()

	body := users.CreateUserJSONBody{
		Properties: map[string]string{
			"hs_job_title":           "CEO",
			"hs_availability_status": "available",
			"hs_working_hours":       "[{\"days\":\"MONDAY_TO_FRIDAY\",\"startMinute\":540,\"endMinute\":1020}]",
			"hs_email":               "mehul.engt@gmail.com",
			"hs_standard_time_zone":  "America/New_York",
		},
	}
	body1, err := json.Marshal(body)
	response, err := ct.CreateUserWithBodyWithResponse(context.Background(), "application/json", bytes.NewReader(body1))
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 201 {
		if response.JSON201 == nil || response.JSON201.Id == "" {
			t.Fatalf("Response contains no results")
		}

		if response.JSON201.Properties != nil {
			t.Logf("Properties: %s", response.JSON201.Properties)
		} else {
			t.Log("No properties found.")
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %s", response.StatusCode(), response.Body)
	}
}

func TestUpdateUser(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	userId := "389641148436"

	// Initialize a variable of type Users
	user := users.UpdateUserJSONRequestBody{
		Properties: map[string]string{
			"hs_job_title":           "CEO",
			"hs_availability_status": "away",
			"hs_working_hours":       "[{\"days\":\"SATURDAY_SUNDAY\",\"startMinute\":540,\"endMinute\":1020}]",
		},
	}

	ct := crm.Users()
	response, err := ct.UpdateUserWithResponse(context.Background(), userId, user)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Properties == nil {
			t.Fatalf("Response contains no results")
		}

		for key, value := range response.JSON200.Properties {
			t.Logf("Key: %s, Value: %+v\n", key, value)
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}

func TestDeleteUser(t *testing.T) {
	config := configuration.Configuration{
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}
	crm := crm.NewCrmDiscovery(&config)

	userId := "427637542790"

	ct := crm.Users()
	response, err := ct.DeleteUserByIdWithResponse(context.Background(), userId)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 204 {
		t.Logf("User Deleted")
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
