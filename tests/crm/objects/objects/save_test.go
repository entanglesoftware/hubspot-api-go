package objects_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/entanglesoftware/hubspot-api-go/codegen/crm/objects/objects"

	"github.com/entanglesoftware/hubspot-api-go/tests/crm"
)

func TestSaveObject(t *testing.T) {
	hsClient := crm.GetTestHubSpotClient(t)
	associationTypeId := int32(580)
	associationCategory := "HUBSPOT_DEFINED"
	id := "28106025611"

	ct := hsClient.Crm().Objects()

	objectType := "leads"

	body := objects.CreateObjectJSONRequestBody{
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
	}

	response, err := ct.CreateObjectWithResponse(context.Background(), objectType, body)
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
