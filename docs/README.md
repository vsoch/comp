# Comp

Comp is a tool for comparing container environments, meaning between containers,
or between the host and container. Eventually I'd like to define a more controlled
interface for setting envars on the fly, etc.

## Setup

By default, `comp` will use docker as a container backend, however you can set
your preferred container backend by doing:

## TODO set container backend via environment?

## Commands

Compenv provides basic commands for interacting with containers - e.g., at face value
it serves as a wrapper. Here are some examples:

```bash
# shell into the docker.io/ubuntu container
$ comp shell ubuntu

# run the container
$ comp run vanessa/salad

# inspect the environment of the container
$ comp env vanessa/salad
```

### Environment

The most basic thing you can do is inspect the environment. With a container this
looks like:

```bash
$ comp env vanessa/salad
/usr/bin/podman run -it --entrypoint env vanessa/salad
NAME    	VALUE
--------	-----
HOME		/root
HOSTNAME	7a9f4092d7e1
PATH		/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
TERM		xterm
container	podman
```

and to inspect the host, you can leave out the container name (or provide an explicit `.` to indicate the present working directory.

```bash
$ comp env
$ comp env .
...
```

### Diff

The basic of a diff is to show changes. This means that we eliminiate variables
that are the same, and only leave behind those that are added or removed from the
first to the second.

```bash
# Compare container vanessa/salad to local environment
# E.g., how did the local environment change from the container?
$ comp diff vanessa/salad .
```

## TODO

 - move env into main library
 - differ should be able to break apart changes to PATH/LD_LIBRARY_PATH etc
 - create search command to search the environment
 - create json output option
 - create diff command to compare envs (json output, or colored diff)
 - create releases for different platforms

