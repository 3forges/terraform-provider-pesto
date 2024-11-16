#!/bin/bash

# --- # --- # --- 
# --- # --- 
# --- 
# This script installs 'goreleaser' on
# GNU/Linux Debian, using bash.
# --- 
# --- # --- 
# --- # --- # --- 

uninstallPreviouslyInstalled () {
export INSTALLED_GORELEASER_BIN_PATH=$(which goreleaser)

if [ "x${INSTALLED_GORELEASER_BIN_PATH}" == "x" ]; then
  echo "GoReleaser is not installed"
else
  echo "GoReleaser is installed at [${INSTALLED_GORELEASER_BIN_PATH}]"
  ls -alh ${INSTALLED_GORELEASER_BIN_PATH}
  export  INSTALLED_GORELEASER_LINT_VERSION=$(${INSTALLED_GORELEASER_BIN_PATH} version)
  echo "Installed version of GoReleaser is [${INSTALLED_GORELEASER_LINT_VERSION}]"
  # ----
  echo "Now uninstaling GoReleaser version [${INSTALLED_GORELEASER_LINT_VERSION}]"
  rm -f "${INSTALLED_GORELEASER_BIN_PATH}"
  export UNINSTALLATION_EXIT_CODE=$?
  if [ "${UNINSTALLATION_EXIT_CODE}" == "0" ]; then
    echo "GoReleaser version [${INSTALLED_GORELEASER_LINT_VERSION}] successfully uninstalled."
  else
    echo "An Error occured uninstalling GoReleaser version [${INSTALLED_GORELEASER_LINT_VERSION}]."
  fi;
fi;

}

uninstallPreviouslyInstalled


export GORELEASER_VERSION=${GORELEASER_VERSION:-'v2.4.5'}

# --- 
# 
# echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | sudo tee /etc/apt/sources.list.d/goreleaser.list
# sudo apt update
# sudo apt install goreleaser


# echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | sudo tee /etc/apt/sources.list.d/goreleaser.list
# sudo apt-get update -y
# sudo apt-get install -y goreleaser

curl -sfL https://goreleaser.com/static/run | bash VERSION=${GORELEASER_VERSION} -s -- check



goreleaser version