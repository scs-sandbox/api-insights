# Copyright 2022 Cisco Systems, Inc. and its affiliates.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# SPDX-License-Identifier: Apache-2.0

FROM library/golang:1.18-alpine AS builder

WORKDIR /go/src/github.com/cisco-developer/api-insights/api
COPY . .

ENV CGO_ENABLED=0
ENV GOPATH=/go
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
RUN go build -a -tags musl -ldflags '-w' -o ./bin/api-insights github.com/cisco-developer/api-insights/api/cmd/api-insights

FROM ghcr.io/cisco-developer/api-insights-tools:v0.1.1
ARG BUILDER_DIR=/go/src/github.com/cisco-developer/api-insights/api

LABEL Description="API Insights"
LABEL org.opencontainers.image.source="https://github.com/cisco-developer/api-insights"

RUN apk update && \
    apk upgrade && \
    apk add \
        bash \
        mailcap \
        ca-certificates \
        go \
    && rm -rf /var/cache/apk/*

COPY --from=builder ${BUILDER_DIR}/*.json /config/
COPY --from=builder ${BUILDER_DIR}/bin/api-insights /usr/local/bin/api-insights

WORKDIR /config

COPY ./*.json /config/
COPY ./internal/data/ /config/internal/data/

ENTRYPOINT ["/usr/local/bin/api-insights", "serve"]
