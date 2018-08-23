# Sample App

This orderapp, is a sample app. Not for production purpose.

![untitled diagram](https://user-images.githubusercontent.com/1962129/44516810-58e5f100-a6c6-11e8-9bda-cb8e58e7d35d.png)

here how `sample-app` diagram looks like.

### Requirement
To be able to start sample-app you need to pass these environment variables:

- DATABASE_DSN          : postgresql datasource name
- QUEUE_DSN             : rabbitmq datasource name
- CUSTOMER_APP_BASE_URL : http base url

### Default configuration

`sample-app` would expose on port `:9000`, and only accept on 1 endpoint which is `POST /create`

