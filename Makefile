BINARY_NAME=mytodo
INSTALL_DIR=/usr/bin

todo/run:
	@go run ./cmd/todo -list

todo/build:
	@echo "Building the $(BINARY_NAME) binary..."
	go build -o $(BINARY_NAME) ./cmd/todo

todo/install: todo/build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	@sudo mv $(BINARY_NAME) $(INSTALL_DIR)

todo/clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

todo/uninstall:
	@echo "Uninstalling $(BINARY_NAME) from $(INSTALL_DIR)..."
	@sudo rm -f $(INSTALL_DIR)/$(BINARY_NAME)

.PHONY: todo/run todo/build todo/install todo/clean todo/uninstall
