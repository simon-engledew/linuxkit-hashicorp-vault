MAKEFLAGS += --no-builtin-rules
.SUFFIXES:

FILES := $(shell git ls-files)
FILES += $(shell git ls-files --others --exclude-standard)
FILES := $(filter-out $(shell git ls-files --deleted), $(FILES))

FILES += pkg/init/root/usr/bin/console

define search
$(filter $(1),$(FILES))
endef

PKGS := $(patsubst %/build.yml,%,$(call search,pkg/%/build.yml))
OPTIONS := -network -org simon-engledew -hash latest

pkg/init/root/usr/bin/console: go/cmd/console $(call search,go/cmd/console/%)
	(cd $<; GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(abspath $@) .)

.SECONDEXPANSION:
$(PKGS): $$(call search,$$@/%)
	linuxkit pkg build $(OPTIONS) $@
	@touch $@

out/vault.vmdk: vault.yml $(PKGS)
	(linuxkit build -dir $(abspath out/) -format vmdk vault.yml)

.PHONY: vmware
vmware: out/vault.vmdk
	@(cd out; linuxkit run vmware -disk size=5G -mem 4096 vault)
