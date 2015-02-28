.PHONY: version all run release clean

APP_NAME := $(shell basename "$$PWD")
VER := $(shell head -n 1 .version)

SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))

ifeq ($(BRANCH), master)
	VERSION := $(VER)
else
	VERSION := $(VER)-$(BRANCH)
endif

BUILD := $(SHA)-$(BRANCH)

EXECUTABLE := bin/$(APP_NAME)
PACKAGE := release/$(APP_NAME)-$(VERSION).tar.gz

# Build Binaries setting main.version and main.build vars
LDFLAGS :=-ldflags "-X main.version $(VERSION) -X main.build $(BUILD)"

# Package target

.DEFAULT: all

all: | $(EXECUTABLE)

release: | $(PACKAGE)

run: bin/$(APP_NAME)
	bin/$(APP_NAME)

test:
	go test -r -cover

clean:
	@echo Cleaning Workspace...
	rm -dRf bin
	rm -dRf release

info:
	@printf "%-9s %-8s \n" "Name:" $(APP_NAME)
	@printf "%-9s %-8s \n" "Version:" $(VERSION)
	@printf "%-9s %-8s \n" "Package:" $(PACKAGE)
	@printf "%-9s \n"  "Changelog:"
	@cat .version | sed 1d

$(EXECUTABLE):
	go build $(LDFLAGS) -o $@ $<

$(PACKAGE): all
ifeq ("$(wildcard $(PACKAGE))","")
	@echo "Packaging Binaries..."
	@mkdir -p tmp
	@cp -R $(EXECUTABLE) tmp
	@mkdir -p release
	@echo "Creating new release";
	@tar -zcf $@ -C tmp .;
	@echo "## $(VERSION)" >> CHANGELOG.md;
	@cat .version | sed 1d >> CHANGELOG.md;
	@echo "\n" >> CHANGELOG.md;
	@rm -rf tmp
	@echo "Done, created: $@"
else
	@echo "Release already exists, skipping.";
endif

