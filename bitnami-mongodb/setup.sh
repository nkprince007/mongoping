#!/usr/bin/env bash

CHART_VERSION="${CHART_VERSION:-13.16.4}"
HELM_RELEASE="${HELM_RELEASE:-xcloud}"
NAMESPACE="${NAMESPACE:-xcloud-mongo}"

helm upgrade \
    --install "$HELM_RELEASE" \
    -n "$NAMESPACE" \
    --create-namespace \
    --version "$CHART_VERSION" \
    --values values.yaml \
    oci://registry-1.docker.io/bitnamicharts/mongodb

