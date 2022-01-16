# User Guide

Comp is a tool for comparing container environments, meaning between containers,
or between the host and container. Eventually I'd like to define a more controlled
interface for setting envars on the fly, etc.

## Setup

By default, `comp` will use docker as a container backend, however you can set
your preferred container backend by doing:

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

By default the output will print to the screen in a table. You can also ask for json, pretty json,
or writing to file:

```bash
$ comp env vanessa/salad --json
$ comp env vanessa/salad --json --pretty
$ comp env vanessa/salad --outfile examples/vanessa-salad-env.json
```

By default, saving to file does not suppress output. To do that:

```bash
$ comp env vanessa/salad --outfile examples/vanessa-salad-env.json --quiet
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

The above will print a colored diff to your screen:
 - red is removed from the first to the second
 - green is added
 - yellow is changed (with details in diff)
 
You can also print to json, pretty json, or output to file.

```bash
$ go run main.go diff . vanessa/salad --json
```
```bash
$ go run main.go diff . vanessa/salad --json --pretty
/usr/bin/podman run -it --entrypoint env vanessa/salad
{
    "added": {
        "HOSTNAME": "dbcec8f50013",
        "container": "podman"
    },
    "removed": {
        "COLORTERM": "truecolor",
        "CONDA_DEFAULT_ENV": "base",
...
        "_": "/usr/local/go/bin/go",
        "_CE_CONDA": "",
        "_CE_M": ""
    },
    "unchanged": {},
    "changed": {
        "HOME": {
            "name": "HOME",
            "original": "/home/dinosaur",
            "new": "/root"
        },
        "PATH": {
            "name": "PATH",
            "original": "/home/vanessa/anaconda3/bin:/home/vanessa/anaconda3/condabin:/home/vanessa/.rbenv/plugins/ruby-build/bin:/home/vanessa/.rbenv/shims:/home/vanessa/.rbenv/bin:/home/vanessa/.cargo/bin:/home/vanessa/anaconda3/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/usr/local/go/bin",
            "new": "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
        },
        "TERM": {
            "name": "TERM",
            "original": "xterm-256color",
            "new": "xterm"
        }
    }
}
```
```bash
$ go run main.go diff busybox vanessa/salad --outfile examples/vanessa-salad-busybox-diff.json
```

Or don't print the table to the screen:

```bash
$ go run main.go diff . vanessa/salad --outfile examples/vanessa-salad-diff.json --quite
```
