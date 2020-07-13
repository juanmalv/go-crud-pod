# CRUD with Go and Package Oriented Design (POD)

Simple implementation of a basic API with CRUD following Package Oriented Design guidelines.

This implementation uses MongoDB as Database and GinGonic as a Router.

Dependencies will be downloaded with `go install`

In order to run this project, cd into `./scripts` and run `docker-compose up`. 
This will build a docker container with MongoDB.

## Operations

Supported requests:

 - GET `http://localhost:8080/income`
 - POST `http://localhost:8080/income`
 - PUT `http://localhost:8080/income`
 - DELETE `http://localhost:8080/income`

Request Body:

> {
> "amount" : 23
> }
