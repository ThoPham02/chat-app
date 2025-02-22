ROOT_DIR=$(shell pwd)
OUT_DIR=$(ROOT_DIR)/out
LOGS_DIR=$(ROOT_DIR)/logs

MODULE_NAME=chat_app
SERVICE=service
SERVICE_DIR=$(ROOT_DIR)/$(SERVICE)
API_DIR=$(ROOT_DIR)/api
SCHEMA_DIR=$(ROOT_DIR)/schema

# auth service
AUTH_DIR=$(SERVICE_DIR)/auth
AUTH_API_DIR=$(AUTH_DIR)/api
AUTH_MODEL_DIR=$(AUTH_DIR)/model

# chat service
CHAT_DIR=$(SERVICE_DIR)/chat
CHAT_API_DIR=$(CHAT_DIR)/api
CHAT_MODEL_DIR=$(CHAT_DIR)/model

gen-auth-service:
	goctl api go -api $(API_DIR)/auth.api -dir $(AUTH_API_DIR)

gen-chat-service:
	goctl api go -api $(API_DIR)/chat.api -dir $(CHAT_API_DIR)

gen-service: gen-auth-service gen-chat-service

gen-auth-model: 
	goctl model mysql ddl -src="${SCHEMA_DIR}/auth.sql" -dir="${AUTH_MODEL_DIR}" --ignore-columns=""

gen-chat-model:
	goctl model mysql ddl -src="${SCHEMA_DIR}/chat.sql" -dir="${CHAT_MODEL_DIR}" --ignore-columns=""

gen-model: gen-auth-model gen-chat-model

dep-init:
	go mod init $(MODULE_NAME)
	go mod tidy
dep:
	go mod tidy

runs:
	go run main.go -f etc/server.yaml