PKG_MGR ?= pnpm
API_DIR := apps/api
WEB_DIR := apps/web
INFRA_DIR := infra

.PHONY: install
install:
	$(PKG_MGR) install
	cd $(API_DIR) && go mod download

.PHONY: up
up:
	cd $(INFRA_DIR) && docker compose up -d

.PHONY: down
down:
	cd $(INFRA_DIR) && docker compose down

.PHONY: dev
dev:
	@echo "Starting API + Web (dev)..."
	cd $(API_DIR) && go run ./cmd/server &
	$(PKG_MGR) -C $(WEB_DIR) dev

.PHONY: test
test:
	cd $(API_DIR) && go test ./...
	$(PKG_MGR) -C $(WEB_DIR) test || true

.PHONY: lint
lint:
	cd $(API_DIR) && gofmt -w .
	$(PKG_MGR) -C $(WEB_DIR) lint || true

.PHONY: build
build:
	cd $(API_DIR) && go build ./cmd/server
	$(PKG_MGR) -C $(WEB_DIR) build

.PHONY: ci
ci: install test build
