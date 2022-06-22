.PHONY: install.codegen
install.codegen:
	@$(GO) install ${ROOT_DIR}/tools/codegen/codegen.go