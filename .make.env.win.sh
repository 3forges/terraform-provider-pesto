#!/bin/bash

#### >>>>>>>>>> #### >>>>>>>>>> #### >>>>>>>>>>> 
#### >>>>>>>>>> #### >>>>>>>>>> #### >>>>>>>>>>> 
#### >>> Make env. Windows: Git bash for Windows 
#### >>>>>>>>>> #### >>>>>>>>>> #### >>>>>>>>>>> 
#### >>>>>>>>>> #### >>>>>>>>>> #### >>>>>>>>>>> 
# /!\: Does not work for powershell!
# 

export TF_RC_FILE_PATH=$(echo $APPDATA | sed 's#\\#/#g' | sed 's#C:#/c#g')/terraform.rc

export MY_GO_BIN_FOLDER=$(go env GOBIN)
export MY_GO_PATH_FOLDER=$(go env GOPATH)

export MY_NONDEFAULT_GO_BIN_FOLDER=$(echo "$MY_GO_BIN_FOLDER/bin" | sed 's#\\#/#g' | sed 's#C:##g')
export MY_DEFAULT_GO_BIN_FOLDER=$(echo "$MY_GO_PATH_FOLDER/bin" | sed 's#\\#/#g' | sed 's#C:#/c#g')


# note:
# > if GOBIN is not empty, then the golang binary will be generated inside the bin subfolder of the 'GOBIN' folder.
# > if GOBIN is empty, then the golang binary will be generated inside the bin subfolder of the 'GOPATH' folder.

if [ "x${MY_GO_BIN_FOLDER}" == "x" ]; then
  echo "GO BIN is empty, so will use default [${MY_DEFAULT_GO_BIN_FOLDER}]"
  export PATH_FOR_DEV_OVERRIDES="${MY_DEFAULT_GO_BIN_FOLDER}"
else
  echo "GO BIN is not empty"
  export PATH_FOR_DEV_OVERRIDES="${MY_NONDEFAULT_GO_BIN_FOLDER}"
fi;

# echo "MY_NONDEFAULT_GO_BIN_FOLDER=[${MY_NONDEFAULT_GO_BIN_FOLDER}]"
# echo "MY_GO_BIN_FOLDER=[${MY_GO_BIN_FOLDER}]"
# echo "MY_GO_PATH_FOLDER=[${MY_GO_PATH_FOLDER}]"
# echo "MY_GO_PATH_FOLDER/bin=[${MY_GO_PATH_FOLDER}/bin]"
# echo "MY_DEFAULT_GO_BIN_FOLDER=[${MY_DEFAULT_GO_BIN_FOLDER}]"
# echo "TF_RC_FILE_PATH=[${TF_RC_FILE_PATH}]"
echo "  PATH_FOR_DEV_OVERRIDES=[${PATH_FOR_DEV_OVERRIDES}]"



###############

export WHERE_GO_INSTALL_PUTS_EXE="$(go env GOPATH | sed 's#C:#/c#g' | sed 's#\\#/#g')/bin"
# export WHERE_GO_INSTALL_PUTS_EXE="${PATH_FOR_DEV_OVERRIDES}"
# echo "  PATH_FOR_DEV_OVERRIDES=[${PATH_FOR_DEV_OVERRIDES}]"
echo "  WHERE_GO_INSTALL_PUTS_EXE=[${WHERE_GO_INSTALL_PUTS_EXE}]"

# ---
# According:
# - https://developer.hashicorp.com/terraform/cli/commands/init#plugin-installation
# and
# - https://developer.hashicorp.com/terraform/cli/commands/init#plugin-dir-path
# and
# - https://developer.hashicorp.com/terraform/cli/config/config-file#explicit-installation-method-configuration
# 
# when I use the [-plugin-dir=<some folder path>] 
# tofu init command option, the 
# tofu init command looks up for the provider executable in
# a subfolder of the folder passed to the
# [-plugin-dir=<some folder path>] option, and that subfolder 
# must comply with the below layout :
# 
#  ${HOSTNAME}/${NAMESPACE}/${TYPE}/${VERSION}/${TARGET}/terraform-provider-${TYPE}_v${VERSION}_${TARGET}.exe
# 
# and in my case:
# 
#  export HOSTNAME="pesto-io.io"
#  export NAMESPACE="terraform"
#  export TYPE="pesto" # (it could be export TYPE="aws" )
#  export VERSION="0.0.1" # this value needs only to be a
#                         # semver compliant version number, of 
#                         # the form "x.y.z", with x, y, and z 
#                         # being integers (and nothing else).
#                         # 
#                         # As I am in development mode, I
#                         # just chose 0.0.1, but it could be #                         # any version number, it would 
#                         # work just as well
#  export TARGET="windows_386" # note here: since I am running 
#                              # tofu (terraform) from Git bash 
#                              # for windows, the TARGET (cpu
#                              # arch) is 'windows_386', even
#                              # if my physical machine 
#                              # actually is an amd64 cpu arch
# 
#  # export TARGET="windows_amd64"

export TF_DEV_HOSTNAME="pesto-io.io"
export TF_DEV_NAMESPACE="terraform"
export TF_DEV_TYPE="pesto" # (it could be export TYPE="aws" )
export TF_DEV_VERSION=${TF_DEV_VERSION:-"0.0.1"} # this value needs only to be a
                       # semver compliant version number, of 
                       # the form "x.y.z", with x, y, and z 
                       # being integers (and nothing else).
                       # 
                       # As I am in development mode, I
                       # just chose 0.0.1, but it could be #                         # any version number, it would 
                       # work just as well
export TF_DEV_TARGET="windows_386" # note here: since I am running 
                            # tofu (terraform) from Git bash 
                            # for windows, the TARGET (cpu
                            # arch) is 'windows_386', even
                            # if my physical machine 
                            # actually is an amd64 cpu arch
