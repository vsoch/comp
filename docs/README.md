# Compenv

Compenv is a tool for comparing container environments, meaning between containers,
or between the host and container. Eventually I'd like to define a more controlled
interface for setting envars on the fly, etc.

## Setup

By default, `compenv` will use docker as a container backend, however you can set
your preferred container backend by doing:

## TODO set container backend via environment?

## Commands

Compenv provides basic commands for interacting with containers - e.g., at face value
it serves as a wrapper. Here are some examples:

```bash
# shell into the docker.io/ubuntu container
$ compenv shell ubuntu

# run the container
$ compenv run vanessa/salad

# inspect the environment of the container
$ compenv env vanessa/salad
```
