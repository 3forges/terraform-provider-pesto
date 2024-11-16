#!/bin/bash

# ---
# installs latest release version of 
# cosign which can be bult with go1.21 : cosign version v2.2.4

# ---
# Note: installing cosign from source, by building it from source, is long, but it is something justified: If you
uninstallPreviouslyInstalled () {
export INSTALLED_COSIGN_BIN_PATH=$(which goreleaser)

if [ "x${INSTALLED_COSIGN_BIN_PATH}" == "x" ]; then
  echo "GoReleaser is not installed"
else
  echo "GoReleaser is installed at [${INSTALLED_COSIGN_BIN_PATH}]"
  ls -alh ${INSTALLED_COSIGN_BIN_PATH}
  export  INSTALLED_COSIGN_LINT_VERSION=$(${INSTALLED_COSIGN_BIN_PATH} version)
  echo "Installed version of GoReleaser is [${INSTALLED_COSIGN_LINT_VERSION}]"
  # ----
  echo "Now uninstaling GoReleaser version [${INSTALLED_COSIGN_LINT_VERSION}]"
  rm -f "${INSTALLED_COSIGN_BIN_PATH}"
  export UNINSTALLATION_EXIT_CODE=$?
  if [ "${UNINSTALLATION_EXIT_CODE}" == "0" ]; then
    echo "GoReleaser version [${INSTALLED_COSIGN_LINT_VERSION}] successfully uninstalled."
  else
    echo "An Error occured uninstalling GoReleaser version [${INSTALLED_COSIGN_LINT_VERSION}]."
  fi;
fi;

}

uninstallPreviouslyInstalled


export DESIRED_COSIGN_VERSION=${DESIRED_COSIGN_VERSION:-'v2.2.4'}
git clone https://github.com/sigstore/cosign
cd cosign
git checkout "${DESIRED_COSIGN_VERSION}"
go install ./cmd/cosign

$(go env GOPATH)/bin/cosign version

cosign version


exit 0

$ $(go env GOPATH)/bin/cosign version
  ______   ______        _______. __    _______ .__   __.
 /      | /  __  \      /       ||  |  /  _____||  \ |  |
|  ,----'|  |  |  |    |   (----`|  | |  |  __  |   \|  |
|  |     |  |  |  |     \   \    |  | |  | |_ | |  . `  |
|  `----.|  `--'  | .----)   |   |  | |  |__| | |  |\   |
 \______| \______/  |_______/    |__|  \______| |__| \__|
cosign: A tool for Container Signing, Verification and Storage in an OCI registry.

GitVersion:    devel
GitCommit:     fb651b4ddd8176bd81756fca2d988dd8611f514d
GitTreeState:  clean
BuildDate:     2024-04-10T21:57:27
GoVersion:     go1.22.5
Compiler:      gc
Platform:      windows/amd64
