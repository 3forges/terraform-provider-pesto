#!/bin/bash

# --- # --- # --- 
# --- # --- 
# --- 
# This script installs 'goreleaser' on
# Microsoft Windows, using Git Bash for Windows.
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
# export VERSION="${GORELEASER_VERSION}"
# 
export WHERE_AM_I=$(pwd)
export TMP_INSTALL_OPS_HOME=$(mktemp -d -t INSTALL_GORELEASER_XXXXXX)

cd ${TMP_INSTALL_OPS_HOME}
echo " > TMP_INSTALL_OPS_HOME=[${TMP_INSTALL_OPS_HOME}]"
ls -alh ${TMP_INSTALL_OPS_HOME}
curl -LO https://github.com/goreleaser/goreleaser/releases/download/${GORELEASER_VERSION}/goreleaser_Windows_x86_64.zip
curl -LO https://github.com/goreleaser/goreleaser/releases/download/${GORELEASER_VERSION}/checksums.txt
curl -LO https://github.com/goreleaser/goreleaser/releases/download/${GORELEASER_VERSION}/checksums.txt.pem
curl -LO https://github.com/goreleaser/goreleaser/releases/download/${GORELEASER_VERSION}/checksums.txt.sig


ls -alh ./goreleaser_Windows_x86_64.zip


verifyIntegrity() {

cosign version
export IS_COSIGN_INSTALLED=$?

if [ "x${IS_COSIGN_INSTALLED}" == "0" ]; then
  echo "cosign is installed"
else
  echo "[verifyIntegrity()] ERROR: cosign is not installed. Install cosign before invoking [$0#verifyIntegrity()]."
  exit 5
fi;

export VERSION_TO_VERIFY=$1
export RELEASES_URL="https://github.com/goreleaser/goreleaser/releases"

if [ "x${VERSION_TO_VERIFY}" == "x" ]; then
  echo "[verifyIntegrity()] ERROR: No version to verify was provided as argument."
  exit 7
fi;

echo "Verifying checksums..."
sha256sum --ignore-missing --quiet --check checksums.txt
sha256sum --ignore-missing --check checksums.txt

if command -v cosign >/dev/null 2>&1; then
        echo "Verifying signatures..."
        cosign verify-blob \
                --certificate-identity-regexp "https://github.com/goreleaser/goreleaser.*/.github/workflows/.*.yml@refs/tags/$VERSION_TO_VERIFY" \
                --certificate-oidc-issuer 'https://token.actions.githubusercontent.com' \
                --cert "$RELEASES_URL/download/$VERSION_TO_VERIFY/checksums.txt.pem" \
                --signature "$RELEASES_URL/download/$VERSION_TO_VERIFY/checksums.txt.sig" \
                checksums.txt
else
        echo "Could not verify signatures, cosign is not installed."
fi
}

verifyIntegrity ${GORELEASER_VERSION}

export IS_GORELEASER_INTEGRITY_OK=$?

if [ "x${IS_COSIGN_INSTALLED}" == "0" ]; then
  echo "cosign is installed"
else
  echo "[$0] ERROR: goreleaser download [${TMP_INSTALL_OPS_HOME}/goreleaser_Windows_x86_64.zip] integrity check failed. Stopping installation process."
  exit 11
fi;

mkdir -p ./deflated/
unzip ./goreleaser_Windows_x86_64.zip -d ./deflated/


ls -alh ${TMP_INSTALL_OPS_HOME}
ls -alh ${TMP_INSTALL_OPS_HOME}/deflated/

ls -alh ${TMP_INSTALL_OPS_HOME}/deflated/goreleaser

mkdir -p $(go env GOPATH)/bin/

cp ${TMP_INSTALL_OPS_HOME}/deflated/goreleaser.exe $(go env GOPATH)/bin/

mkdir -p $(go env GOPATH)/bin/goreleaser.${GORELEASER_VERSION}/

cp -fR ${TMP_INSTALL_OPS_HOME}/deflated/* $(go env GOPATH)/bin/goreleaser.${GORELEASER_VERSION}/

rm -fr ${TMP_INSTALL_OPS_HOME}

installAutoCompletions () {

# ---
# Install auto-completions

# - bash
if [ -f ~/.local/share/bash-completion/completions/goreleaser.bash ]; then
  rm ~/.local/share/bash-completion/completions/goreleaser.bash
fi;

mkdir -p ~/.local/share/bash-completion/completions/
cp $(go env GOPATH)/bin/goreleaser.${GORELEASER_VERSION}/completions/goreleaser.bash ~/.local/share/bash-completion/completions
export RC_FILE_CONTAINS_SRC_CMD=$(cat ~/.bashrc | grep 'source ~/.local/share/bash-completion/completions/goreleaser.bash')
if [ "x${RC_FILE_CONTAINS_SRC_CMD}" == "x" ]; then
  echo 'source ~/.local/share/bash-completion/completions/goreleaser.bash' | tee -a ~/.bashrc
fi;

# - fish
if [ -f ~/.local/share/bash-completion/completions/goreleaser.fish ]; then
  rm ~/.local/share/bash-completion/completions/goreleaser.fish
fi;

mkdir -p ~/.local/share/fish-completion/completions/
cp $(go env GOPATH)/bin/goreleaser.${GORELEASER_VERSION}/completions/goreleaser.fish ~/.local/share/fish-completion/completions
export RC_FILE_CONTAINS_SRC_CMD=$(cat ~/.bashrc | grep 'source ~/.local/share/bash-completion/completions/goreleaser.fish')
if [ "x${RC_FILE_CONTAINS_SRC_CMD}" == "x" ]; then
  echo 'source ~/.local/share/fish-completion/completions/goreleaser.fish' | tee -a ~/.fishrc
fi;

# - zsh
if [ -f ~/.local/share/bash-completion/completions/goreleaser.zsh ]; then
  rm ~/.local/share/bash-completion/completions/goreleaser.zsh
fi;

mkdir -p ~/.local/share/zsh-completion/completions/
cp $(go env GOPATH)/bin/goreleaser.${GORELEASER_VERSION}/completions/goreleaser.zsh ~/.local/share/zsh-completion/completions
export RC_FILE_CONTAINS_SRC_CMD=$(cat ~/.bashrc | grep 'source ~/.local/share/bash-completion/completions/goreleaser.zshrc')
if [ "x${RC_FILE_CONTAINS_SRC_CMD}" == "x" ]; then
  echo 'source ~/.local/share/zsh-completion/completions/goreleaser.zsh' | tee -a ~/.zshrc
fi;
}

installAutoCompletions


cd ${WHERE_AM_I}

goreleaser  --version
