# Go miniURL

A URL minifier REST API written in Go, with Gin framework

## Stack

- [Go](https://go.dev)
- [Gin Framework](https://pkg.go.dev/github.com/gin-gonic/gin)
- [MongoDB](https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo)

## Packages

- [JWT](https://pkg.go.dev/github.com/golang-jwt/jwt/v5)
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [Cors](https://pkg.go.dev/github.com/gin-contrib/cors)

## Environment Variables

- `APP_MODE=development | production`
- `PORT=4000`
- `DATABASE_URI=<your mongodb uri>`
- `DATABASE_NAME=<your mongodb collection name>`
- `JWT_ISSUER=`
- `JWT_SECRET=`
- `EMAIL_VERIFICATION_TOKEN_SECRET=`
- `HASH_COST=10`

> larger HASH_COST will result in slower but more secure hashing

## Getting Started

- Clone the repo and run `go get` to install all the dependencies.
- Create a `.env` file at the root of the project and add the above mentioned variables.
- Run `go run .` to start the project in `debug` mode. It is a good idea to set `APP_MODE=development`.
- Request the following endpoints from either your own UI or Postman.

## Endpoints

> Common prefix – `/api`

### Authentication

> Common prefix – `/auth`

#### Login

#### Register

#### Reset Password

- Get link
- Set new password

#### Validate user token

### Minify

> Common prefix – `/minify`

#### Get all minified urls

**Request**

```
curl --location --request GET 'http://localhost:4000/api/minify'
```

**Response**

```json
[
  {
    "_id": "6598609d7d1a137466f0f8eb",
    "label": "Official Go Website",
    "active": true
  }
]
```

#### Generate new minified URL

**Request**

Payload

```json
{
  "link": "https://go.dev",
  "label": "Official Go Website"
}
```

```
curl --location --request POST 'http://localhost:4000/api/minify' \
--header 'Content-Type: application/json' \
--data '{
	"link": "https://go.dev",
	"label": "Official Go Website"
}'
```

#### Get single minified url

Sending a request to this endpoint will redirect to the original url that was minified.

```
curl --location --request GET 'http://localhost:4000/api/minify/65985472c4fd59a24436a281'
```

#### Update a url

#### Delete a url
