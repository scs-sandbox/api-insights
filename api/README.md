# API-Insights-Backend

## Prerequisites

* Golang 1.18+
* Install ruleset in the api folder
```
npm install @cisco-developer/api-insights-openapi-rulesets
```
* Install Java and `openapi-diff-cli-2.1.0-beta.3-all.jar`, and put it to some folder.
```
curl -OL https://repo1.maven.org/maven2/org/openapitools/openapidiff/openapi-diff-cli/2.1.0-beta.3/openapi-diff-cli-2.1.0-beta.3-all.jar
```

## Get started locally

* Start the dependencies in this repo's root folder. Run the command:
```
docker-compose up mysql
```
* Start the backend service
```
OPENAPI_DIFF_JAR_FILE=/some-dir/openapi-diff-cli-2.1.0-beta.3-all.jar go run cmd/api-insights/main.go serve
```
* Test your service is running on port `8081`.
```
curl -v localhost:8081/v1/healthz
```

## API docs

API docs is here: [docs](https://developer.cisco.com/docs/api-insights/#!overview)
