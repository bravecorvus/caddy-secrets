# caddy secrets

> This is a direct copy of the code that was originally hosted at [https://github.com/startsmartlabs/caddy-secrets](https://github.com/startsmartlabs/caddy-secrets). To the people who wrote this originally, I can not longer build Caddy server from source because of this dependency which is required by `freman/caddy-reauth/backends/refresh/auth.go`. If you feel that you need more attribution (I cannot see the individual members of `smartsmartlabs`), or protest to have `caddy-secrets` be archived in this way, please contact me at [me@gilgameshskytrooper.io](me@gilgameshskytrooper.io).

This is a plugin for [caddy server](https://caddyserver.com/), it reads a yaml file to make static secrets available to the middleware in a package level yaml.MapSlice

## Build

To use it you need to compile your own version of caddy with this plugin. First fetch the code

- `go get -u github.com/mholt/caddy/...`
- **(No Longer a valid git repo)**`go get -u github.com/startsmartlabs/caddy-secrets`
- `go get -u github.com/gilgameshskytrooper/caddy-secrets`

Update the file in `$GOPATH/src/github.com/mholt/caddy/caddy/caddymain/run.go` and import `_ "github.com/gilgameshskytrooper/caddy-secrets"`.
Then update the file in `$GOPATH/src/github.com/mholt/caddy/caddyhttp/httpserver/plugin.go` and add "secrets" at the start of the `var directives` list. If there are plugins that come before secrets on the list, they won't have access to the values read from the yaml file.

And finally build caddy with:

- `cd $GOPATH/src/github.com/mholt/caddy/caddy`
- `./build.bash`

This will produce the caddy binary in that same folder. For more information about how plugins work read [this doc](https://github.com/mholt/caddy/wiki/Writing-a-Plugin:-Directives). 

## Usage

Example minimal usage in `Caddyfile`

```
caddy.test {
    secrets secrets.yml
}
```

This module leverages [go-getter](https://github.com/hashicorp/go-getter); so long as the file path fits the formats they handle it can get files from relative path, Git, Mercurial, HTTP, or Amazon S3 (if the AWS credentials are setup). The map slice that gets populated with the secret values is called SecretsMap and will be availabe in modules that import caddy-secrets.

## Dependencies

- Currently, the secrets file has to be in yaml format and the map slice exported is of type [yaml.MapSlice](https://godoc.org/gopkg.in/yaml.v2#MapSlice)
