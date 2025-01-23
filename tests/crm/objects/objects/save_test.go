package objects_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/objects"
	"github.com/entanglesoftware/hubspot-api-go/configuration"
	"github.com/entanglesoftware/hubspot-api-go/hubspot"
)

func TestSaveObject(t *testing.T) {
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
	associationTypeId := int32(580)
	associationCategory := "HUBSPOT_DEFINED"
	id := "28106025611"

	// Make the API call
	reqBody := objects.CreateObjectJSONRequestBody{
		Associations: &[]struct {
			To *struct {
				Id *string "json:\"id,omitempty\""
			} "json:\"to,omitempty\""
			Types *[]struct {
				AssociationCategory *objects.CreateObjectJSONBodyAssociationsTypesAssociationCategory "json:\"associationCategory,omitempty\""
				AssociationTypeId   *int32                                                            "json:\"associationTypeId,omitempty\""
			} "json:\"types,omitempty\""
		}{
			{
				To: &struct {
					Id *string "json:\"id,omitempty\""
				}{
					Id: &id, // Replace with actual ID
				},
				Types: &[]struct {
					AssociationCategory *objects.CreateObjectJSONBodyAssociationsTypesAssociationCategory "json:\"associationCategory,omitempty\""
					AssociationTypeId   *int32                                                            "json:\"associationTypeId,omitempty\""
				}{
					{
						AssociationCategory: (*objects.CreateObjectJSONBodyAssociationsTypesAssociationCategory)(&associationCategory), // Replace with actual category
						AssociationTypeId:   &associationTypeId,                                                                        // Replace with actual association type ID
					},
				},
			},
		},
		Properties: &map[string]string{},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("API call failed: %v", err)
	}

	ct := hsClient.Crm().Objects()

	objectType := "leads"

	contentType := "application/json"

	response, err := ct.CreateObjectWithBodyWithResponse(context.Background(), objectType, contentType, bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error serializing lead properties: %v", err)
	}

	if response.StatusCode() == 201 {
		if response.JSON201 == nil || response.JSON201.Id == "" {
			fmt.Println("No data found.")
			return
		}

		properties := response.JSON201.Properties
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
