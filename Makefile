GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=npuzzle
GOPATH := ${PWD}

SRC=src/lib/ src/goals src/heuristics src/node src/npuzzle src/utils\
	src/multiprocessing/ src/displaying main.go

export GOPATH

all: $(BINARY_NAME)

$(BINARY_NAME): $(SRC)
	$(GOBUILD) -o $(BINARY_NAME) -v

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
