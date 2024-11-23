package http

import "github.com/entanglesoftware/hubspot-api-go/constants"

// IAPIType defines an interface for API interactions
type IAPIType interface {
	GetPage(limit int, after string, properties []string, propertiesWithHistory []string, associations []string, archived bool) (ICollectionType, error)
}

// INextPage represents the next page in pagination
type INextPage struct {
	After string
}

// IForwardPaging represents forward paging information
type IForwardPaging struct {
	Next *INextPage
}

// ICollectionType represents a collection type with results and paging
type ICollectionType struct {
	Results []interface{}
	Paging  *IForwardPaging
}

// getAll fetches all pages from the API and returns the combined results
func getAll(api IAPIType, limit *int, after string, properties []string, propertiesWithHistory []string, associations []string, archived bool) ([]interface{}, error) {
	limitInternal := constants.DefaultObjectsLimit
	if limit != nil {
		limitInternal = *limit
	}

	var result []interface{}
	afterInternal := after

	for {
		response, err := api.GetPage(limitInternal, afterInternal, properties, propertiesWithHistory, associations, archived)
		if err != nil {
			return nil, err
		}

		result = append(result, response.Results...)

		if response.Paging == nil || response.Paging.Next == nil {
			break
		}

		afterInternal = response.Paging.Next.After
	}

	return result, nil
}

type MockAPI struct{}

func (api *MockAPI) GetPage(limit int, after string, properties []string, propertiesWithHistory []string, associations []string, archived bool) (ICollectionType, error) {
	// Mock implementation for demonstration purposes
	results := []interface{}{"item1", "item2"}
	paging := &IForwardPaging{Next: &INextPage{After: ""}} // No more pages

	return ICollectionType{
		Results: results,
		Paging:  paging,
	}, nil
}
