package schemas_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/schemas"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

// TestUpdateSchema tests the creation of a schema in HubSpot CRM
func TestUpdateSchema(t *testing.T) {
	// Fetch the access token from the environment
	token := os.Getenv("HS_ACCESS_TOKEN")

	if token == "" {
		t.Skip("HS_ACCESS_TOKEN is not set. Skipping test.")
	}

	// Initialize the configuration
	config := configuration.Configuration{
		AccessToken:            token,
		BasePath:               configuration.BaseURL,
		NumberOfAPICallRetries: 3,
	}

	hsClient := hubspot.NewClient(config)

	// Define the schema payload
	pluralLabel := "My New1 objects"
	singularLabel := "My New1 object"

	schema := schemas.UpdateSchemaJSONRequestBody{
		Labels: &struct {
			Plural   *string "json:\"plural,omitempty\""
			Singular *string "json:\"singular,omitempty\""
		}{
			Plural:   &pluralLabel,
			Singular: &singularLabel,
		},
	}

	// Serialize the product properties to JSON
	body, err := json.Marshal(schema)
	if err != nil {
		log.Fatalf("Error serializing product properties: %v", err)
	}

	contentType := "application/json"

	objectType := "2-39502275"

	// Make the API call to create the schema
	response, err := hsClient.Crm().SchemaItems().UpdateSchemaWithBodyWithResponse(context.Background(), objectType, contentType, bytes.NewReader(body))
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	if response.StatusCode() == 200 {
		t.Logf("Schema update successfully")
		responseString := string(response.Body)
		fmt.Println("Response as string:")
		fmt.Println(responseString)
	} else {
		t.Fatalf("Test failed with status code %d: %v", response.StatusCode(), response)
	}
}
