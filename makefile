# Makefile for building and running Go program with arguments

# Variables
NAME := main
CORE_CMD := cmd

# Build the program
build:
	@echo "Building $(PROGRAM_NAME)..."
	cd cmd && go build -o ${NAME}

# Run the program with arguments
run: build
	@echo "Running $(NAME) with arguments: $(ARGS)"
	cd cmd && ./$(NAME) $(ARGS)

test:
	@echo "Running tests..."
	go test -v ./...