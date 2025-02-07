# ARG BUILDPLATFORM="linux/amd64"
ARG BUILDERIMAGE="golang:1.21"
# ARG BASEIMAGE="gcr.io/distroless/static:nonroot"
# ARG BASEIMAGE="cgr.dev/chainguard/static:latest-glibc"

# FROM --platform=$BUILDPLATFORM $BUILDERIMAGE as builder
FROM $BUILDERIMAGE as builder
ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT=""
ARG LDFLAGS

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=${TARGETOS}
ENV GOARCH=${TARGETARCH}
ENV GOARM=${TARGETVARIANT}

WORKDIR /go/app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o provider cmd/provider/main.go



FROM alpine:3.16.2
RUN apk --no-cache add ca-certificates bash

ARG TARGETOS
ARG TARGETARCH

COPY --from=builder /go/app/provider /usr/local/bin/provider

ENV USER_ID=65532

# Setup Terraform environment

## Provider-dependent configuration
ARG TERRAFORM_VERSION
ARG TERRAFORM_PROVIDER_SOURCE
ARG TERRAFORM_PROVIDER_VERSION
ARG TERRAFORM_PROVIDER_DOWNLOAD_NAME
ARG TERRAFORM_NATIVE_PROVIDER_BINARY
ARG TERRAFORM_CUSTOM_REGISTRY_DOWNLOAD_URL
## End of - Provider-dependent configuration

# Provider controller needs these environment variable at runtime
ENV TERRAFORM_VERSION=${TERRAFORM_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE=${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION=${TERRAFORM_PROVIDER_VERSION}
ENV TERRAFORM_CUSTOM_REGISTRY_DOWNLOAD_URL=${TERRAFORM_CUSTOM_REGISTRY_DOWNLOAD_URL}

ENV PLUGIN_DIR /terraform/provider-mirror/registry.terraform.io/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${TARGETOS}_${TARGETARCH}
# ENV PLUGIN_DIR=/terraform/provider-mirror/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${TARGETOS}_${TARGETARCH}
ENV TERRAFORM_NATIVE_PROVIDER_PATH=${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}
ENV TF_CLI_CONFIG_FILE=/terraform/.terraformrc
ENV TF_FORK=0

RUN mkdir -p ${PLUGIN_DIR}

ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp
ADD https://${TERRAFORM_CUSTOM_REGISTRY_DOWNLOAD_URL}/${TERRAFORM_PROVIDER_VERSION}/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp
# ADD https://releases.hashicorp.com/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}/${TERRAFORM_PROVIDER_VERSION}/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp
ADD cluster/images/provider-grafana/terraformrc.hcl ${TF_CLI_CONFIG_FILE}

RUN unzip /tmp/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip -d /usr/local/bin \
  && chmod +x /usr/local/bin/terraform \
  && rm /tmp/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip \
  && unzip /tmp/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip -d ${PLUGIN_DIR} \
  && chmod +x ${PLUGIN_DIR}/* \
  && rm /tmp/${TERRAFORM_PROVIDER_DOWNLOAD_NAME}_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip \
  && chown -R ${USER_ID}:${USER_ID} /terraform
# End of - Setup Terraform environment

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
