package main_tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	brasilapi "api/busca-cep/services"

	"github.com/stretchr/testify/assert"
)

func TestFetchAddressFromAPI(t *testing.T) {
	// Simulating a local test server that replicates the behavior of the external service
	server := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Handling the request for a specific CEP
		if req.URL.Path == "/api/cep/v2/01001000" {
			// Returning a sample JSON response for the specified CEP
			response := brasilapi.Address{
				Cep:          "01001000",
				State:        "SP",
				City:         "São Paulo",
				Neighborhood: "Sé",
				Street:       "Praça da Sé",
				Service:      "TestService",
				Location: brasilapi.Location{
					Type: "TestType",
					Coordinates: brasilapi.Coordinates{
						Longitude: "123.456",
						Latitude:  "-45.678",
					},
				},
			}

			// Marshal the response to JSON
			responseJSON, _ := json.Marshal(response)
			rw.WriteHeader(http.StatusOK)
			rw.Write(responseJSON)
		}
	})
	testServer := httptest.NewServer(server)
	defer testServer.Close()

	// Using the service function to make a request to the local test server
	response, err := http.Get(testServer.URL + "/api/cep/v2/01001000")
	assert.NoError(t, err)

	// Reading the response body
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	// Unmarshaling the response body into an Address struct
	var address brasilapi.Address
	err = json.Unmarshal(body, &address)
	assert.NoError(t, err)

	// Asserting the values from the real API against the expected values
	assert.Equal(t, "01001000", address.Cep)
	assert.Equal(t, "SP", address.State)
	assert.Equal(t, "São Paulo", address.City)
	assert.Equal(t, "Sé", address.Neighborhood)
	assert.Equal(t, "Praça da Sé", address.Street)
	// Add assertions for other fields as needed
}
