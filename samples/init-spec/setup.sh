#! /bin/bash

set -e
# set -ex

export $(grep -v '^#' .env | xargs)

host=${host:-http://localhost:8081}

echo "You are using host: $host"

services=("cart" "catalogue" "payment" "user" "order")
# services=("cart" )
product_tag="DevRel Store"

base_spec_version=${1:-v0.0-rev1}
updated_spec_version=${2:-v0.0-rev2}
updated_spec_version2=${3:-v0.1-rev1}

function ensure_exist() {
  [ -d "./$1" ] || (echo "$1" does not exist && exit 1)
}

function upload() {
  echo "## Upload $1"
  rev=${1#*rev}
  for service in ${services[*]}; do
    echo upload spec rev "$rev" for "$service"
    api-insights-cli -H "$host" service uploadspec "$1/$service".json -s "$service" --revision "$rev"
    sleep 1
  done
}

function recreate_service() {
  service=$1
  service_title="$(tr '[:lower:]' '[:upper:]' <<< ${service:0:1})${service:1}"

  echo delete "$service"
  api-insights-cli -H "$host" service delete "$service" --debug || true
  echo create "$service"

  printf -v payload '{ "organization_id": "DevNet", "product_tag": "%s", "name_id": "%s", "title": "%s Demo API", "description": "%s microservice for %s demo application", "contact": {"name": "Engineering Team", "email": "engineering@merchandiseshop.com", "url": "https://app-8081-apiregistry1.devenv-int.ap-ne-1.devnetcloud.com/"}, "analyzers_configs": {"drift": {"service_name_id": "%s.sock-shop"}} }' "$product_tag" "$service" "$service_title" "$service_title" "$product_tag" "$service"

  api-insights-cli -H "$host" service create --data "$payload" || true
}

function main {
  ensure_exist "$base_spec_version" || exit 1
  ensure_exist "$updated_spec_version" || exit 1

  for service in ${services[*]}; do
    recreate_service "$service"
  done

  api-insights-cli -H "$host" service list

  upload $base_spec_version
  sleep 5
  upload $updated_spec_version
  sleep 5
  upload $updated_spec_version2

  echo
  echo Done! Check the result in api-insights-ui
}

main "$@"
