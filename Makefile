
GO=go
CMD_DIR=./cmd
BACKEND_DIR=$(CMD_DIR)/backend


MQTT_URL=localhost
MQTT_PORT=1883

.PHONY: run
run: build
	MQTT_URL=$(MQTT_URL) MQTT_PORT=$(MQTT_PORT) ./backend

.PHONY: build
build:
	$(GO) build $(BACKEND_DIR)