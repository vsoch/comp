# Config
##### $HOME/.pak/conf/backends.conf
```toml
[local]
type = "builtin"
backups = "5"

[hpc]
type = "ssh"
host = "ssh-host.university.edu"
port = "22"
username = "user"
ssh-key-path = "/path/to/ssh/key"

[jetstream]
type = "openstack"
id = "bc257241e21747768c83fb9806af392d"
project_id = "e99b6f4b9bf84a9da27e20c9cbfe887a"
secret = "securesecret"
endpoint = "iu.jetstream-cloud.org"
```

##### $HOME/.pak/conf/paks/spack-dev.conf
```toml
name = "spack-dev"
image = "ghcr.io/spack/ubuntu-bionic"
createdAt = "2021-12-20"

[[mounts]]
path = "/external/path:/internal/path"

[[mounts]]
path = "/external/path/two:/internal/path/two"

```
