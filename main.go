package main

import (
	"encoding/json"
	"log"
	"net/http"

	customerapi "github.com/alileza/orderapp/api/customer"
	"github.com/alileza/orderapp/config"
	"github.com/alileza/orderapp/pkg/order"
	"github.com/jmoiron/sqlx"
	"github.com/streadway/amqp"
)

func main() {
	cfg := config.Retrieve()

	db, err := sqlx.Open("postgres", cfg.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	queueConn, err := amqp.Dial(cfg.QueueDSN)
	if err != nil {
		panic(err)
	}
	customerAPI := customerapi.New(cfg.CustomerAppBaseURL)

	orderClient := order.New(db, queueConn, customerAPI)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var requestBody struct {
				CustomerID int64 `json:"customer_id"`
				ProductID  int64 `json:"product_id"`
			}
			if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}

			_, err := orderClient.Create(&order.Order{
				CustomerID: requestBody.CustomerID,
				ProductID:  requestBody.ProductID,
			})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusCreated)
		default:
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
	})

	if err := http.ListenAndServe(":9000", mux); err != nil {
		log.Fatal(err)
	}
}
