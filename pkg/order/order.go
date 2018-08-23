package order

import (
	"encoding/json"
	"time"

	customer "github.com/alileza/sample-app/api/customer"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

type Order struct {
	OrderID    int64     `json:"order_id"`
	CustomerID int64     `json:"customer_id"`
	ProductID  int64     `json:"product_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type Client struct {
	db          *sqlx.DB
	queue       *amqp.Connection
	customerAPI *customer.API
}

func New(db *sqlx.DB, queue *amqp.Connection, customerAPI *customer.API) *Client {
	return &Client{db, queue, customerAPI}
}

func (c *Client) Create(o *Order) (*Order, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow("INSERT INTO orders (customer_id, product_id) VALUES ($1, $2) RETURNING order_id", o.CustomerID, o.ProductID).
		Scan(&o.OrderID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	cust, err := c.customerAPI.Get(o.CustomerID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	ch, err := c.queue.Channel()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(map[string]interface{}{
		"order":    o,
		"customer": cust,
	})
	if err != nil {
		return nil, err
	}

	err = ch.Publish("orders", "created", true, true, amqp.Publishing{
		Body: body,
	})

	return o, nil
}
