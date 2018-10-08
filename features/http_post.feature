Feature: handling http post request

  Scenario: Send an invalid request
    Given "http-client" send request to "POST /"
    Then "http-client" response code should be 400

  Scenario: Send a valid request on Customer App unavailable
    Given "http-client" send request to "POST /" with body
      """
        {
          "customer_id": 123,
          "product_id": 334
        }
      """
    Then "http-client" response code should be 500

  Scenario: Send a valid request
    Given listen message from "mq" target "orders:created"
    Given set "customer-app" with path "/customers/123" response code to 200 and response body
      """
        {
          "customer_id": 123,
          "email": "alirezayahya@gmail.com",
          "status": "active"
        }
      """
    Given "http-client" send request to "POST /" with body
      """
        {
          "customer_id": 123,
          "product_id": 334
        }
      """
    Then "http-client" response code should be 201
    Then "db1" table "orders" should look like
      | customer_id | product_id |
      | 123         | 334        |
    Then message from "mq" target "orders:created" count should be 1
    Then message from "mq" target "orders:created" should look like
      """
        {
          "order": {
            "order_id": "*",
            "customer_id": 123,
            "product_id": 334,
            "created_at": "*"
          },
          "customer": {
            "customer_id": 123,
            "email": "alirezayahya@gmail.com",
            "status": "active"
          }
        }
      """
