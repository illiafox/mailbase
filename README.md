# mailbase

[![Go 1.18](https://img.shields.io/github/go-mod/go-version/illiafox/mailbase.svg)](https://go.dev/learn/)
[![GoReportCard example](https://goreportcard.com/badge/github.com/illiafox/mailbase)](https://goreportcard.com/report/github.com/illiafox/mailbase)

[![Go](https://github.com/illiafox/mailbase/actions/workflows/go.yml/badge.svg)](https://github.com/illiafox/mailbase/actions/workflows/go.yml)
[![Docker](https://github.com/illiafox/mailbase/actions/workflows/docker-image.yml/badge.svg)](https://github.com/illiafox/mailbase/actions/workflows/docker-image.yml)
[![CodeQL](https://github.com/illiafox/mailbase/actions/workflows/codeql.yml/badge.svg)](https://github.com/illiafox/mailbase/actions/workflows/codeql.yml)

## login/register/logout/delete server with mail checking using go, redis and mysql

# [Images (must-view) ⬇️](https://github.com/illiafox/mailbase/edit/master/README.md#images)


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

# Images
---


Session not found -> register/login

![Screenshot from 2022-04-11 15-07-03](https://user-images.githubusercontent.com/61962654/162739574-d8b73673-7ead-45e9-b9fc-61bd4fff276e.png)
![Screenshot from 2022-04-11 15-07-08](https://user-images.githubusercontent.com/61962654/162739595-062d0f37-4305-4f65-9b55-8112343f289c.png)

---

Forgot password? Wait for mail

![Screenshot from 2022-04-11 15-07-15](https://user-images.githubusercontent.com/61962654/162739755-7263ec9e-1af1-4825-ad4e-47753782ccec.png)

![Screenshot from 2022-04-11 15-07-37](https://user-images.githubusercontent.com/61962654/162739754-32b0fd95-84ed-4341-a860-69fe96218517.png)


---

Then Reset

![Screenshot from 2022-04-11 15-07-59](https://user-images.githubusercontent.com/61962654/162739749-063e2642-8b38-4269-be6e-e201f1ddf3ba.png)

--- 

Welcome to main page

![Screenshot from 2022-04-11 15-08-54](https://user-images.githubusercontent.com/61962654/162740039-1f48b708-94c8-4ae0-8bc4-74b8366af99c.png)

---

If you're admin, advanced link will appear

![Screenshot from 2022-04-11 15-08-58](https://user-images.githubusercontent.com/61962654/162740118-bc432379-680a-4049-99d4-ba8bad64f493.png)

---

Let's write report `HTML SUPPORTED`

![Screenshot from 2022-04-11 15-11-51](https://user-images.githubusercontent.com/61962654/162740218-71f35ba8-1cac-4bed-a6c3-9bac498f1bff.png)

![Screenshot from 2022-04-11 15-11-51](https://user-images.githubusercontent.com/61962654/162740346-c77e3cfe-58cb-49d8-b308-6fde5cb438d3.png)

---

Admin panel

![Screenshot from 2022-04-11 15-12-02](https://user-images.githubusercontent.com/61962654/162740519-87241b45-cc5b-46de-99b4-780b309f8645.png)

---

View reports:

![Screenshot from 2022-04-11 15-21-12](https://user-images.githubusercontent.com/61962654/162740677-c1b64410-32cf-4f30-819f-c48a3001ee85.png)
![Screenshot from 2022-04-11 15-22-34](https://user-images.githubusercontent.com/61962654/162740694-558ba176-c9b7-4559-b039-666c19eb72e6.png)

---

Read and answer `HTML SUPPORTED`:

![Screenshot from 2022-04-11 15-14-08](https://user-images.githubusercontent.com/61962654/162740767-2c9759da-635d-4e93-a8f7-c5e1ae30f57b.png)

---

Reporter will receive mail:

![Screenshot from 2022-04-11 15-38-46](https://user-images.githubusercontent.com/61962654/162741014-a8afd7ff-2f90-4223-818a-22f019116d02.png)

---

Status of report will be updated:

![Screenshot from 2022-04-11 15-14-23](https://user-images.githubusercontent.com/61962654/162741114-0def69d4-5aee-4c87-b7ff-c2644031798f.png)

---

Let's disable server (admin panel wiil be available):

![Screenshot from 2022-04-11 15-16-33](https://user-images.githubusercontent.com/61962654/162741216-6861a850-4895-481a-8d6d-8ef2ef4e385a.png)

---

Ooops...

![Screenshot from 2022-04-11 15-16-44](https://user-images.githubusercontent.com/61962654/162741294-acb9c3c4-170e-4fc0-8f20-6a2d065edcf9.png)

---

It would be better to fix this...

![Screenshot from 2022-04-11 15-16-55](https://user-images.githubusercontent.com/61962654/162741382-4dab7d12-52cc-4f06-81bc-d1cc4831cc83.png)

---

Superadmins who are granted from console can add/delete admins be theirs email or id

![Screenshot from 2022-04-11 15-42-06](https://user-images.githubusercontent.com/61962654/162741497-7abc3689-8cc2-4e13-808f-7cb6aefe6e23.png)




# [nojwt](https://github.com/illiafox/mailbase/tree/nojwt) branch
The old **unsecured** server version works well with `http`, but all cookies can be stolen easily

Another solution is to use **[ngrok](https://ngrok.com/)** and the like services to create http tunnel, which allow you to choose [newer version](https://github.com/illiafox/mailbase) with `jwt`

#### ngrok example for regions:
```shell
ngrok http -region=eu 8080

ngrok http -region=us 8080
```


