# Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.

# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.

# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

APP="blinky"
EXE="blinky"

SHELL = /bin/bash

DIR = $(shell pwd)

DOCKER = docker

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

VERSION=$(shell \
        grep "const Version" src/github.com/nlamirault/blinky/version/version.go \
        |awk -F'=' '{print $$2}' \
        |sed -e "s/[^0-9.]//g" \
	|sed -e "s/ //g")

PACKAGE=$(APP)-$(VERSION)
ARCHIVE=$(PACKAGE).tar

all: help

help:
	@echo -e "$(OK_COLOR)==== $(APP) [$(VERSION)] ====$(NO_COLOR)"
	@echo -e "$(WARN_COLOR)init$(NO_COLOR)   :  Install requirements"
	@echo -e "$(WARN_COLOR)deps$(NO_COLOR)   :  Install dependencies"
	@echo -e "$(WARN_COLOR)build$(NO_COLOR)  :  Make all binaries"
	@echo -e "$(WARN_COLOR)clean$(NO_COLOR)  :  Cleanup"
	@echo -e "$(WARN_COLOR)reset$(NO_COLOR)  :  Remove all dependencies"

clean:
	@echo -e "$(OK_COLOR)[$(APP)] Cleanup$(NO_COLOR)"
	@rm -f $(EXE) $(EXE)_* $(APP)-*.tar.gz coverage.out gover.coverprofile

.PHONY: init
init:
	@echo -e "$(OK_COLOR)[$(APP)] Install requirements$(NO_COLOR)"
	@go get github.com/constabulary/gb/...

build:
	@echo -e "$(OK_COLOR)[$(APP)] Build $(NO_COLOR)"
	@$(GB) build





doc:
	@GOPATH=$(GO_PATH) godoc -http=:6060 -index

fmt:
	@echo -e "$(OK_COLOR)[$(APP)] Launch fmt $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go fmt github.com/nlamirault/$(APP)/...

errcheck:
	@echo -e "$(OK_COLOR)[$(APP)] Launch errcheck $(NO_COLOR)"
	@GOPATH=$(GO_PATH) $(ERRCHECK) github.com/nlamirault/$(APP)/...

vet:
	@echo -e "$(OK_COLOR)[$(APP)] Launch vet $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go vet github.com/nlamirault/$(APP)/...

lint:
	@echo -e "$(OK_COLOR)[$(APP)] Launch golint $(NO_COLOR)"
	@GOPATH=$(GO_PATH) $(GOLINT) github.com/nlamirault/$(APP)/...

style: fmt vet lint

test:
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go test -v github.com/nlamirault/$(APP)/...

race:
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests race $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go test -race github.com/nlamirault/$(APP)/...

coverage:
	@echo -e "$(OK_COLOR)[$(APP)] Launch code coverage $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go test github.com/nlamirault/$(APP)/... -cover

covoutput:
	@echo -e "$(OK_COLOR)[$(APP)] Launch code coverage $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go test ${pkg} -covermode=count -coverprofile=coverage.out
#	@GOPATH=$(GO_PATH) go tool cover -html=coverage.out
	@GOPATH=$(GO_PATH) go tool cover -func=coverage.out

release: clean build
	@echo -e "$(OK_COLOR)[$(APP)] Make archive $(VERSION) $(NO_COLOR)"
	@rm -fr $(PACKAGE) && mkdir $(PACKAGE)
	@cp -r $(EXE) $(PACKAGE)
	@tar cf $(ARCHIVE) $(PACKAGE)
	@gzip $(ARCHIVE)
	@rm -fr $(PACKAGE)
	@addons/github.sh $(VERSION)

# for go-projectile
gopath:
	@echo ${GOPATH}
