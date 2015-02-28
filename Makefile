.PHONY: build release run package changelog test clean info

NAME := $(shell basename "$$PWD")
VER  := $(shell head -n 1 .version)

ARCH   := $(shell uname -m)
SHA    := $(shell git rev-parse --short HEAD)
BRANCH := $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
BUILD  := $(SHA)-$(BRANCH)

VER_PRESENT := $(shell grep "$(VER)" CHANGELOG.md)

ifeq ($(BRANCH), master)
	VERSION := $(VER)
else
	VERSION := $(VER)-$(BRANCH)
endif

LDFLAGS :=-ldflags "-X main.version $(VERSION) -X main.build $(BUILD)"

.DEFAULT: build

build:
	@echo "Creating Binaries..."
	@mkdir -p bin/Linux  && GOS=linux go build $(LDFLAGS) -o bin/Linux/$(NAME)
	@mkdir -p bin/Darwin && GOOS=darwin go build $(LDFLAGS) -o bin/Darwin/$(NAME)

ifeq ($(VER_PRESENT),)
release: package changelog
else
release: package
	@echo "Version already exists, skipping CHANGELOG.md";
endif

run: build
	@bin/$(shell uname)/$(NAME)

package: build
	@echo "Packaging Binaries..."
	@mkdir -p release
	@tar -zcf release/$(NAME)_$(VERSION)_linux_$(ARCH).tgz  -C bin/Linux  $(NAME)
	@tar -zcf release/$(NAME)_$(VERSION)_darwin_$(ARCH).tgz -C bin/Darwin $(NAME)

changelog:
	@echo "Updating CHANGELOG.md";
	@echo "## $(VERSION)" >> CHANGELOG.md;
	@cat .version | sed 1d >> CHANGELOG.md;
	@echo "\n" >> CHANGELOG.md;
	@echo "Done, check ./release"

test:
	go test -r -cover

clean:
	@echo Cleaning Workspace...
	rm -dRf bin
	rm -dRf release

info:
	@printf "%-9s %-8s \n" "Name:" $(NAME)
	@printf "%-9s %-8s \n" "Version:" $(VERSION)
	@printf "%-9s \n"  "Changelog:"
	@cat .version | sed 1d

