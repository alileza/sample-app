package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type API struct {
	httpClient *http.Client
	baseURL    string
}

func New(baseURL string) *API {
	return &API{
		httpClient: new(http.Client),
		baseURL:    baseURL,
	}
}

type Customer struct {
	CustomerID int64  `json:"customer_id"`
	Email      string `json:"email"`
	Status     string `json:"status"`
}

func (a *API) url(path string, vals ...interface{}) string {
	return a.baseURL + fmt.Sprintf(path, vals...)
}

func (a *API) Get(customerID int64) (*Customer, error) {
	resp, err := a.httpClient.Get(a.url("/customers/%d", customerID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cust Customer
	if err := json.NewDecoder(resp.Body).Decode(&cust); err != nil {
		return nil, err
	}

	return &cust, nil
}
