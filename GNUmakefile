# SHELL := /bin/bash
# I use ONESHELL: to be able to generate files with cat EOF
# - 
.ONESHELL:
# --- 
# default: testacc
default: dev.win

tf_dev_hostname = $(TF_DEV_HOSTNAME)
tf_dev_namespace = $(TF_DEV_NAMESPACE)
tf_dev_type = $(TF_DEV_TYPE)
tf_dev_version = $(TF_DEV_VERSION)
tf_dev_target = $(TF_DEV_TARGET)
where_go_install_puts_exe = $(WHERE_GO_INSTALL_PUTS_EXE)
path_for_dev_overrides = $(PATH_FOR_DEV_OVERRIDES)
tf_rc_file_path = $(TF_RC_FILE_PATH)


# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m


###########################
### Dev Tools

.PHONY: dev.win.terraformrc
dev.win.terraformrc:
	echo 'Windows Dev Terraformrc'
	echo " tf_rc_file_path = $(tf_rc_file_path)"
	echo " path_for_dev_overrides = $(path_for_dev_overrides)"
	# ---
	# backup previous terraformrc if exists
	rm $(tf_rc_file_path).bckup || true
	cp $(tf_rc_file_path) $(tf_rc_file_path).bckup
	# cat <<EOF >./.generated.tf_rc_file.rc
	# ---
	# regenerate the terraformrc
	cat <<EOF >$(tf_rc_file_path)
	provider_installation {

	  dev_overrides {
	    # "pesto-io.io/terraform/pesto" = "<PATH>"
	    "$(tf_dev_hostname)/$(tf_dev_namespace)/$(tf_dev_type)" = "$(path_for_dev_overrides)"
	  }

	  # For all other providers, install them directly from their origin provider
	  # registries as normal. If you omit this, Terraform will _only_ use
	  # the dev_overrides block, and so no other providers will be available.
	  direct {}
	}
	EOF

.PHONY: dev.win
dev.win:
	echo 'Windows Dev Build'
	go install .
	echo " where_go_install_puts_exe = $(where_go_install_puts_exe)"
	# ls -alh "$(where_go_install_puts_exe)"
	echo " path_for_dev_overrides = $(path_for_dev_overrides)"
	# ls -alh "$(path_for_dev_overrides)"
	mkdir -p $(where_go_install_puts_exe)/$(tf_dev_hostname)/$(tf_dev_namespace)/$(tf_dev_type)/$(tf_dev_version)/$(tf_dev_target)/
	cp "$(where_go_install_puts_exe)/terraform-provider-$(tf_dev_type).exe" "$(where_go_install_puts_exe)/$(tf_dev_hostname)/$(tf_dev_namespace)/$(tf_dev_type)/$(tf_dev_version)/$(tf_dev_target)/terraform-provider-$(tf_dev_type)_v$(tf_dev_version)_$(tf_dev_target).exe"
	rm "$(where_go_install_puts_exe)/terraform-provider-$(tf_dev_type).exe"
	# terraform-provider-$(tf_dev_type)_v$(tf_dev_version)_$(tf_dev_target).exe
