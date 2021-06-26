# Deployment

## Deployment through a docker container

The following commands will create a *docker image* and create a *docker container*:

```sh
$ docker build -t earaujoassis/space .
$ docker run -d -p 8080:8080 earaujoassis/space
```

The project also provides a *docker-compose* setup, which could be configured through:

```sh
$ docker-compose up --build
```

As describe in the *docker-compose* file, it relies on a Redis instance and a PostgreSQL database.

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
