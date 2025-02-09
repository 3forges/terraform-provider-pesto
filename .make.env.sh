#!/bin/bash

#### >>>>>>>>>> #### >>>>>>>>>> #### >>>>>>>>>> 
#### >>>>>>>>>> #### >>>>>>>>>> #### >>>>>>>>>> 
#### >>> Make env. GNU/Linux: bash 
#### >>>>>>>>>> #### >>>>>>>>>> #### >>>>>>>>>> 
#### >>>>>>>>>> #### >>>>>>>>>> #### >>>>>>>>>>

export WHERE_GO_INSTALL_PUTS_EXE="$(go env GOPATH)/bin"

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
export TF_DEV_TARGET="linux_amd64" # note here: since I am running 
                            # tofu (terraform) from Git bash 
                            # for windows, the TARGET (cpu
                            # arch) is 'windows_386', even
                            # if my physical machine 
                            # actually is an amd64 cpu arch
