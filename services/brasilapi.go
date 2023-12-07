package brasilapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Coordinates struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

type Location struct {
	Type        string      `json:"type"`
	Coordinates Coordinates `json:"coordinates"`
}

type Address struct {
	Cep          string   `json:"cep"`
	State        string   `json:"state"`
	City         string   `json:"city"`
	Neighborhood string   `json:"neighborhood"`
	Street       string   `json:"street"`
	Service      string   `json:"service"`
	Location     Location `json:"location"`
}

func GetAddressByCEP(req *gin.Context) {
	cep := req.Param("cep")
	address, err := FetchAddressFromAPI(cep)

	if err != nil {
		req.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar os dados"})
		return
	}

	req.JSON(http.StatusOK, address)
}

func FetchAddressFromAPI(cep string) (Address, error) {
	var address Address

	brasilAPI := "https://brasilapi.com.br/api/cep/v2/" + cep
	response, err := http.Get(brasilAPI)

	if err != nil {
		return address, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return address, err
	}

	err = json.Unmarshal(body, &address)
	if err != nil {
		return address, err
	}

	return address, nil
}
