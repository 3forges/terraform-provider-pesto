#!/bin/bash

# --- # --- # --- 
# --- # --- 
# --- 
# This script installs 'golangci-lint' on
# GNU/Linux Alpine.
# --- 
# --- # --- 
# --- # --- # --- 

export GOLANGCI_LINT_VERSION=${GOLANGCI_LINT_VERSION:-'v1.62.0'}
# binary will be $(go env GOPATH)/bin/golangci-lint
# ---
# In Alpine Linux (as it does not come with curl by default)
wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s "${GOLANGCI_LINT_VERSION}"

# ---
# In stdout, you should find:
# golangci/golangci-lint info checking GitHub for tag 'v1.62.0'
# golangci/golangci-lint info found version: 1.62.0 for v1.62.0/xxx/amd64

golangci-lint --version

# ---
# In stdout, you should find:
# golangci-lint has version 1.62.0 built with go1.23.2 from 22b58c9b on 2024-11-10T19:09:02Z
# 





