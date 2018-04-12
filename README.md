# caddy secrets

This is a plugin for [caddy server](https://caddyserver.com/), it reads a yaml file to make static secrets available to the middleware in a package level yaml.MapSlice

## Build

To use it you need to compile your own version of caddy with this plugin. First fetch the code

- `go get -u github.com/mholt/caddy/...`
- `go get -u github.com/startsmartlabs/caddy-secrets`

Update the file in `$GOPATH/src/github.com/mholt/caddy/caddy/caddymain/run.go` and import `_ "github.com/startsmartlabs/caddy-secrets"`.
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
