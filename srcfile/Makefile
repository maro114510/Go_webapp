# Copyright (c) Nohira Atsuki. All rights reserved.

# ###########################################
TARGET = $(shell basename `pwd`)

.PHONY: build format test clean ppwd cp init tidy

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



# End of Make