package objects_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/objects"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

func TestUpdateObject(t *testing.T) {
	// Fetch the access token from the environment
	token := os.Getenv("HS_ACCESS_TOKEN")

	if token == "" {
		t.Skip("HS_ACCESS_TOKEN is not set. Skipping test.")
	}

	// Correctly initialize the struct with the proper syntax
	config := configuration.Configuration{
		AccessToken:            token,
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}

	hsClient := hubspot.NewClient(config)

	// Initialize the client
	hsClient.SetAccessToken(token)

	// Make the API call
	objectParams := objects.UpdateObjectParams{}

	objectBody := objects.UpdateObjectJSONBody{
		Properties: map[string]string{
			"firstname": "John",
			"lastname":  "Doe",
			"email":     "johndoe57hy@example.com",
		},
	}

	// Serialize the ticket properties to JSON
	body, err := json.Marshal(objectBody)
	if err != nil {
		log.Fatalf("Error serializing ticket properties: %v", err)
	}

	ct := hsClient.Crm().Objects()

	objectType := "contacts"

	objectId := "87481797267"

	contentType := "application/json"

	response, err := ct.UpdateObjectWithBodyWithResponse(context.Background(), objectType, objectId, &objectParams, contentType, bytes.NewReader(body))
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		if response.JSON200 == nil || response.JSON200.Properties == nil {
			fmt.Println("No data found.")
			return
		}

		properties := response.JSON200.Properties
		if properties == nil {
			fmt.Println("No properties found.")
			return
		}
		for key, value := range properties {
			fmt.Printf("%s: %s\n", key, value)
		}
	} else {
		t.Fatalf("Test Failed with status code %d: %v", response.StatusCode(), response)
	}
}
