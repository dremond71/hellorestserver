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

# MySQL running in a Docker Container

Got MySQL running in a Docker container by following https://github.com/emonddr/setup_mysql_workbench_docker_mysql.
(host: localhost, port: 3306)

# Useful References

- https://tutorialedge.net/golang/basic-rest-api-go-fiber/
- https://gorm.io/id_ID/docs/connecting_to_the_database.html
- https://gorm.io/docs/conventions.html#NamingStrategy
- https://github.com/emonddr/setup_mysql_workbench_docker_mysql
- https://medium.com/@pliutau/docker-and-go-modules-4265894f9fc

# Misc

Turns out the Docker command `docker-machine ip default` to find the IP Address of the Docker daemon in Windows 10 is no longer shipped with Docker Desktop. In MySQL WorkBench, I can use `localhost` instead of an IP Address when trying to connect to MySQL running in a Docker container.
