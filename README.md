# mailbase

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/illiafox/mailbase.svg)](https://go.dev/learn/)
[![Go](https://github.com/illiafox/mailbase/actions/workflows/go.yml/badge.svg)](https://github.com/illiafox/mailbase/actions/workflows/go.yml)
[![GoReportCard example](https://goreportcard.com/badge/github.com/illiafox/mailbase)](https://goreportcard.com/report/github.com/illiafox/mailbase)

## simple login/register server with mail checking using go, redis and mysql



# Requirements

* **redis-server:** `6.2.0` (with `GetDel` support)
* **Mysql:** `8.0.28` (don't know accurate version, so use my as definitive)
* **Go:** `1.18` (also you can rename `any` `interface{}` to run project on `1.17`, lower haven't been tested)
* **packages:** `reffered to go.mod`

# Running

You can find **config template** in [cmd](https://github.com/illiafox/mailbase/blob/master/cmd/config.toml) folder

Config [parser](https://github.com/illiafox/mailbase/blob/master/util/config/config.go) supports small variety of formats (you can implement new)


```go
go run . -conf conf.toml -type toml

go run . -conf conf.json -type json

go run . -conf conf.yaml -type yaml
```

# MailVerify link
In [mail/index.html](https://github.com/illiafox/mailbase/blob/9157a8c3b058879b87655a4b2e1bc7ef31c03234/shared/templates/mail/index.html#L18) (line 18, marked) you ought to change mail message link

For instance, default url is `https://localhost:8080/api/verify?key=`.  With unique site it would look like `https://yoursite.com/api/verify?key=`


# Site map

`/` home page

`/register` register page

`/login` login page

`/api/register` parses register form

`/api/login` parses login form

`/api/verify` verifies key, which is sent in email message


# Screenshots 

![Screenshot from 2022-03-20 21-43-56](https://user-images.githubusercontent.com/61962654/159179952-01cefdbf-08ca-401a-adf9-5f3a35c13d1c.png)
![Screenshot from 2022-03-20 21-45-25](https://user-images.githubusercontent.com/61962654/159180004-d8f089b6-e30c-487e-b61b-9d99af345792.png)
![Screenshot from 2022-03-20 21-45-15](https://user-images.githubusercontent.com/61962654/159180007-edacfd64-bee8-4f49-8b02-b61de7f12501.png)

# Todo
1. comments and explanations (soon)
2. password recovery
