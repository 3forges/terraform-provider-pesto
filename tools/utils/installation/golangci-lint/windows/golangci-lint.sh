#!/bin/bash

# --- # --- # --- 
# --- # --- 
# --- 
# This script installs 'golangci-lint' on
# Microsoft Windows, using Git Bash for Windows.
# --- 
# --- # --- 
# --- # --- # --- 
uninstallPreviouslyInstalled () {
export INSTALLED_GOLANGCI_BIN_PATH=$(which golangci-lint)

if [ "x${INSTALLED_GOLANGCI_BIN_PATH}" == "x" ]; then
  echo "GolangCI Lint is not installed"
else
  echo "GolangCI Lint is installed at [${INSTALLED_GOLANGCI_BIN_PATH}]"
  ls -alh ${INSTALLED_GOLANGCI_BIN_PATH}
  export  INSTALLED_GOLANGCI_LINT_VERSION=$(${INSTALLED_GOLANGCI_BIN_PATH} version)
  echo "Installed version of GolangCI Lint is [${INSTALLED_GOLANGCI_LINT_VERSION}]"
  # ----
  echo "Now uninstaling GolangCI Lint version [${INSTALLED_GOLANGCI_LINT_VERSION}]"
  rm -f "${INSTALLED_GOLANGCI_BIN_PATH}"
  export UNINSTALLATION_EXIT_CODE=$?
  if [ "${UNINSTALLATION_EXIT_CODE}" == "0" ]; then
    echo "GolangCI Lint version [${INSTALLED_GOLANGCI_LINT_VERSION}] successfully uninstalled."
  else
    echo "An Error occured uninstalling GolangCI Lint version [${INSTALLED_GOLANGCI_LINT_VERSION}]."
  fi;
fi;

}

uninstallPreviouslyInstalled

export GOLANGCI_LINT_VERSION=${GOLANGCI_LINT_VERSION:-'v1.62.0'}
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin "${GOLANGCI_LINT_VERSION}"

# ---
# In stdout, you should find:
# golangci/golangci-lint info checking GitHub for tag 'v1.62.0'
# golangci/golangci-lint info found version: 1.62.0 for v1.62.0/windows/amd64

golangci-lint --version

# ---
# In stdout, you should find:
# golangci-lint has version 1.62.0 built with go1.23.2 from 22b58c9b on 2024-11-10T19:09:02Z
# 


# ---
# References
# 
# --- [https://golangci-lint.run]
# --- [https://golangci-lint.run/welcome/install/]
# --- [https://golangci-lint.run/welcome/quick-start/]
# 


