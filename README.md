# mailbase

[![Go 1.18](https://img.shields.io/github/go-mod/go-version/illiafox/mailbase.svg)](https://go.dev/learn/)
[![GoReportCard example](https://goreportcard.com/badge/github.com/illiafox/mailbase)](https://goreportcard.com/report/github.com/illiafox/mailbase)

[![Go](https://github.com/illiafox/mailbase/actions/workflows/go.yml/badge.svg)](https://github.com/illiafox/mailbase/actions/workflows/go.yml)
[![Docker](https://github.com/illiafox/mailbase/actions/workflows/docker-image.yml/badge.svg)](https://github.com/illiafox/mailbase/actions/workflows/docker-image.yml)
[![CodeQL](https://github.com/illiafox/mailbase/actions/workflows/codeql.yml/badge.svg)](https://github.com/illiafox/mailbase/actions/workflows/codeql.yml)

## login/register/logout/delete server with mail checking using go, redis and mysql



# Requirements

* **redis-server:** `6.2.0` (with `GetDel` support)
* **Mysql:** `8.0.28` (don't know accurate version, so use my as definitive)
* **Go:** `1.18`
* **packages:** `reffered to go.mod`

# Running

### HTTPS is _required_ for [jwt](https://github.com/golang-jwt/jwt) cookies

---
The **config template** is located in [cmd](https://github.com/illiafox/mailbase/blob/master/cmd/config.toml) folder

The [config parser](https://github.com/illiafox/mailbase/blob/master/util/config/config.go) supports small variety of formats (you can implement new)


``` go
go run . -conf conf.toml -type toml

go run . -conf conf.json -type json

go run . -conf conf.yaml -type yaml
```
---
Logs are saved in `log.txt`

With **[Multiwriter](https://github.com/illiafox/mailbase/blob/master/util/multiwriter/writer.go)** outputs can be expanded to an unlimited count (**[io.Writer](https://pkg.go.dev/io#Writer)** implemented)
```go
log.SetOutput( multiwriter.NewMultiWriter(os.Stderr,fileWriter,otherWriter) )
```
## HTTP mode:

**[jwt](https://github.com/golang-jwt/jwt)** works bad without `https` (depending on the browser)

Although, you can force it in the config file 
``` toml
[Host]
HTTP = true
```


# Mail links
In the [mails](https://github.com/illiafox/mailbase/tree/master/shared/templates/mails) folder you ought to change mail message links

For instance, the default url is `https://localhost:8080/api/verify?key=` 

With an unique site it would look like `https://yoursite.com/api/verify?key=`

# Docker

Image connects to the local databases, `--net=host` is required

To add execution arguments use the `$ARGS` environment variable


# [nojwt](https://github.com/illiafox/mailbase/tree/nojwt) branch
The old **unsecured** server version works well with `http`, but all cookies can be stolen easily

Another solution is to use **[ngrok](https://ngrok.com/)** and the like services to create http tunnel, which allow you to choose [newer version](https://github.com/illiafox/mailbase) with `jwt`

#### ngrok example for regions:
```shell
ngrok http -region=eu 8080

ngrok http -region=us 8080
```
---

# Site map

`/` home page

`/register` register page

`/login` login page

`/api/logout` logout page

`/api/register` parses register form

`/api/login` parses login form

`/api/verify` verifies key, which is sent in email message

`/api/forgot` sends mail with recover password link

`/api/reset` recover form

