# mailbase

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/illiafox/mailbase.svg)](https://go.dev/learn/)
[![Go](https://github.com/illiafox/mailbase/actions/workflows/go.yml/badge.svg)](https://github.com/illiafox/mailbase/actions/workflows/go.yml)
[![GoReportCard example](https://goreportcard.com/badge/github.com/illiafox/mailbase)](https://goreportcard.com/report/github.com/illiafox/mailbase)

## login/register/logout/delete server with mail checking using go, redis and mysql



# Requirements

* **redis-server:** `6.2.0` (with `GetDel` support)
* **Mysql:** `8.0.28` (don't know accurate version, so use my as definitive)
* **Go:** `1.18`
* **packages:** `reffered to go.mod`

# Running

### HTTPS is _required_ for [jwt](https://github.com/golang-jwt/jwt) cookies

---
You can find **config template** in [cmd](https://github.com/illiafox/mailbase/blob/master/cmd/config.toml) folder

Config [parser](https://github.com/illiafox/mailbase/blob/master/util/config/config.go) supports small variety of formats (you can implement new)


``` go
go run . -conf conf.toml -type toml

go run . -conf conf.json -type json

go run . -conf conf.yaml -type yaml
```

## HTTP mode:

**[jwt](https://github.com/golang-jwt/jwt)** works bad without `https` (depending on browser)

Although, you can force it in config file 
``` toml
HTTP = true
```


# Mail links
In [mails](https://github.com/illiafox/mailbase/tree/master/shared/templates/mails) folder you ought to change mail message links

For instance, default url is `https://localhost:8080/api/verify?key=` 

With unique site it would look like `https://yoursite.com/api/verify?key=`

# Docker

Image connects to local databases, `--net=host` is obvious

To add execution arguments use `$ARGS` environment variable


# [nojwt](https://github.com/illiafox/mailbase/tree/nojwt) branch
Old **unsecured** server version works well with `http`, but all cookies can be stolen in few steps

Another solution is use **[ngrok](https://ngrok.com/)** and the like services to create http tunnel, which allow you to choose [newer version](https://github.com/illiafox/mailbase) with `jwt`

### ngrok example:
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

