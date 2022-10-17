# API-Insights-Backend

## Prerequisites

* Golang 1.18+

## Get started locally

* Start the dependencies in this repo's root folder. Run the command:
```
docker-compose up mysql
```
* Start the backend service
```
go run cmd/api-insights/main.go serve
```
* Test your service is running on port `8081`.
```
curl -v localhost:8081/v1/healthz
```

## API docs

API docs is here: [docs](https://developer.cisco.com/docs/api-insights/#!overview)
