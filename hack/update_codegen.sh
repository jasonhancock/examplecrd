#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

source "${CODEGEN_PKG}/kube_codegen.sh"

kube::codegen::gen_client \
    --with-watch \
    --input-pkg-root github.com/jasonhancock/examplecrd/pkg/apis \
    --output-pkg-root github.com/jasonhancock/examplecrd/pkg/generated \
    --output-base "/tmp/go/src" \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt"
