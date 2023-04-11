# order-processor
order-processor is a simple order processing service that can create, list and get orders by their ID. The service is written in go. It has quick validation for its using GORM.
Data for this service is store in MySQL which can be set up using the dockerfile. It has a small amount of testing where SQLite was selected not create a separate server just for few endpoints.
Chi router was chosen as a lightweight solution with easy to use features to handle the HTTP services.

## Requirements
* GO 1.19+
* Docker

## Running the service
1. Clone the repository 
2. Navigate to the cloned project
3. Build the Docker image using the following command:
   `docker build -t my-mysql-image .`
4. Run the Docker container from the built image using the following command:
`docker run --name my-mysql-container -p 3306:3306 -d my-mysql-image`
5. Once the service is running you might need to create a mysql account with the credentials specified in the `schema.sql` file
6. After that is done you can run the service by running `go run main.go` while in the main project directory.

## Specifications
* Service is listening to port `:3000`
* MySQL is running on port `3306`

## Endpoints

### GET HEALTH
* Simple GET request to check if the service is running
  `http://localhost:3000/api/health`

### POST ORDER
* POST request that creates an order, checks if it is valid using GORM and returns the newly created order.
  `http://localhost:3000/api/order`
* Sample body:
```
{
    "proxy_count": 50,
    "name": "Example Order"
}
```

### GET ORDER LIST
* GET request that retrieves all orders from the database.
  `http://localhost:3000/api/order`

### GET ORDER BY ID
* GET request that retrieves a specific order by the orderID.
  `http://localhost:3000/api/order/{orderID}`

## Testing
Testing is only done in order directory for the two handler functions. It uses SQLite to not launch a separate sandbox MySQL server just for testing of a couple handlers.
To test it you can simply run `go test -v ./...` from the root of directory of the project.
Testing could be enhanced though.

## TODO
* Enhance logging
* Increase test coverage
* Add config