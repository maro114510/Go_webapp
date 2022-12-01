# Copyright (c) Nohira Atsuki. All rights reserved.

# ###########################################
TARGET = $(shell basename `pwd`)

.PHONY: build format test clean ppwd cp init tidy

.PHONY: help build-local up down logs ps

SRCDIR  = ./cmd
BINDIR  = ./bin
TESTDIR = ./test

BIN = $(BINDIR)/$(TARGET)

GO    = go
BUILD = build
RUN   = run
TEST  = test
MOD   = mod
TIDY  = tidy

# ###########################################


ppwd:
	@echo $(SRCDIR)

# if you sync github.com you have not use this command
init:
	@$(GO) $(MOD) init $(TARGET)

tidy:
	@$(GO) $(MOD) $(TIDY)

build: ppwd
	@$(GO) $(BUILD) -o $(BIN) $(SRCDIR)

run: ppwd
	@$(GO) $(RUN) $(SRCDIR)

clean:
	@rm -r $(BIN)

test:
	@$(GO) $(TEST) $(TESTDIR) -v

format:
	@gofmt -w -d $(SRCDIR)

cp:
	@$(BIN)

#############################################################################

DOCKER_TAG := latest

build:
	docker build -t srcfile:${DOCKER_TAG} \
		--target deploy ./

build-local:
	docker-compose build --no-cache

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

ps:
	docker-compose ps




#############################################################################


# End of Make