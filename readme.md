# hellorestserver

A simple golang REST API application using

- [Fiber](https://docs.gofiber.io/) ; an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for golang.
- [GORM](https://gorm.io) ; The fantastic ORM library for Golang

 This application 

 - initializes a MySQL database
 - creates a Fiber app
 - sets up the routes
 - starts the Fiber app on a given port

# Dependencies

- [hellorestdatabase](https://github.com/dremond71/hellorestdatabase) ; contains a global variable `var DB *gorm.DB` used by other modules
- [hellorestbook](https://github.com/dremond71/hellorestbook) ; contains the `Book` struct and its CRUD REST API functions

# Using docker-compose

## Build hellorestserver:localdev container

```sh
docker-compose build
```

## Starting hellorestserver and mysqldb

```sh
my-docker-compose-up.bat
```

## Stopping and Starting the services separately

```sh
my-docker-compose-down.bat
```

# Call REST API with Postman

Import `hello-rest.postman_collection.json` into Postman.

Play with the various REST API endpoints.

# Useful References

- https://tutorialedge.net/golang/basic-rest-api-go-fiber/
- https://gorm.io/id_ID/docs/connecting_to_the_database.html
- https://gorm.io/docs/conventions.html#NamingStrategy
- https://medium.com/@pliutau/docker-and-go-modules-4265894f9fc
- https://iamvickyav.medium.com/mysql-init-script-on-docker-compose-e53677102e48
- https://towardsdatascience.com/connect-to-mysql-running-in-docker-container-from-a-local-machine-6d996c574e55
- https://dev.to/kinshukbasu/setting-up-docker-databases-with-mounted-volumes-kjg
- https://hub.docker.com/r/ubuntu/mysql
