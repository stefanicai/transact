# Transact
Enables storing a financial transaction in USD and retrieving it in a different currency

## Includes
### 1. Swagger definition for API
We generate the handler and client/server models using [Ogen](https://ogen.dev/).

**Note**: I've only used OpenAPI/Swagger because I assumed it has to work with older front end clients.
I'd prefer GRPC otherwise for speed and lower bandwidth. Particularly if it's used as an internal microservice. We can use [buf](https://buf.build/) to describe that and generate go models/handler, quite the same as Ogen.

### 2. Source code and unit tests
I use testify for most tests. I added `ginkgo` for the `internal/transaction` package, but I'm faster with testify, so used that for everything else.
To run tests you can use:

```bash
make test
#or
make test-coverage
```

NOTE: There a configuration file in `config/local.yaml` which is used to configure the application. You can change it as you see fit.

### 3. Docker image for deployment
Docker image is in `Dockerfile` file. We build with a golang image, then build the "production" image from a distroless image.

There is a `docker-compose.yaml` file as well, which will start mongo and our application, running on port 8080.

### 4. Kubernetes deployment files
Created a draft of a kubernetes deployment helm charts, in `deployment/helm`. The configuration file for the specific environment are passed in to the container as config maps.

## How to run
The application can be run in a few ways.

### 1. Run from command line
In the root folder of the repository, run

```
make run
```

## 2. Run using docker-compose
In the root folder of the repo, run

```
docker-compose build
docker-compose up
```

That will ramp up mongodb and the application.

## 3. Kubernetes
This was more in the note of `production ready`. I haven't tested it, but the deployment should work, possibly with minor changes. It would deploy a specific image in a namespace `transact-dev`, and pass on the specific config file via config maps.

It's missing the NetworkPolicy and potentially the ServiceEntry if using Istio. They usually don't sit with the code (for security reasons)

## Testing the app
The application will run at http://localhost:8080 or http://127.0.0.1:8080

Sample (working) Store call:
```bash
curl --request POST \
  --url http://127.0.0.1:8080/create \
  --header 'Content-Type: application/json' \
  --data '{
	"description": "A transaction",
	"amount": "10.3",
	"date": "2023-12-05"
}'
```

Sample (working) Get call below. Just replace the ID sent with the one in the response of the Store call.
```bash
curl --request GET \
  --url http://127.0.0.1:8080/get \
  --header 'Content-Type: application/json' \
  --data '{
	"id":"5c545996-6417-440b-a3a2-51d5c9d220b2",
	"country": "Australia"
}'
```

## To do for a production deliverable
- implement OTEL
- add graceful shutdown to the server
- restrict CORS
- wrap errors in handler so internal errors don't go through to the client. Only validation errors should generally go through for security reasons.
- Add better handling of http return error status codes - it's pretty basic now
- a few more tests, I didn't cover quite all the cases
- fix the issue with big.Rat and MongoDB. It seems the issue is not a json marshalling issue (that works ok for pointer fields), but a Mongo specific, might need more reading, potentially changing the type used for holding currency.
- use the context - I haven't used it at all really. E.g. I don't pass the context to the http call, so if context gets cancelled by the client, the request still goes through.
- continuous integration, depending on what's used - e.g. github workflows if hosted on github