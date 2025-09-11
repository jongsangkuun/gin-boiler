include .env
export

.PHONY: dev build clean test migrate-up migrate-down migrate-create atlas-generate docker-up docker-down

# 개발 서버 실행
dev:
	go run cmd/main.go

# 빌드
build:
	go build -o bin/server cmd/main.go

# 테스트 실행
test:
	go test -v ./...

# Atlas 마이그레이션 생성 (localhost)
migrate-create:
	atlas migrate diff --env dev-local

# Atlas 마이그레이션 적용 (localhost)
migrate-up:
	atlas migrate apply --env dev-local

# Atlas 스키마 검증 (localhost)
atlas-validate:
	atlas schema validate --env dev-local

# Atlas 스키마 포맷
atlas-format:
	atlas schema fmt atlas/schema.hcl

# Atlas 스키마 검사 (localhost)
atlas-inspect:
	atlas schema inspect --env dev-local

# Docker 환경에서 Atlas 스키마 검사
atlas-inspect-docker:
	atlas schema inspect --env dev

# GORM으로부터 스키마 생성
gorm-inspect:
	atlas schema inspect --env gorm

# Docker 컨테이너 시작 (docker compose 사용)
docker-up:
	docker compose up -d

# Docker 컨테이너 중지 (docker compose 사용)
docker-down:
	docker compose down

# 클린업
clean:
	rm -rf bin/
	go clean

# migrations 폴더 생성
init-migrations:
	mkdir -p migrations atlas

# PostgreSQL 상태 확인
db-status:
	docker exec boiler-postgres-db pg_isready -h localhost -p 5432 -U boilerplate_user

# PostgreSQL 로그 확인
db-logs:
	docker logs boiler-postgres-db

# 전체 워크플로우 (Docker 시작 + 마이그레이션)
setup: docker-up
	@echo "Docker 컨테이너 시작 중..."
	@sleep 10
	@echo "스키마 검사 중..."
	make atlas-inspect
	@echo "마이그레이션 생성 중..."
	make migrate-create || echo "마이그레이션이 필요하지 않습니다."
	@echo "마이그레이션 적용 중..."
	make migrate-up || echo "적용할 마이그레이션이 없습니다."

# Docker 시스템 체크
docker-check:
	@echo "Docker 버전 확인:"
	@docker --version
	@echo ""
	@echo "Docker Compose 버전 확인:"
	@docker compose version || docker-compose --version || echo "Docker Compose를 찾을 수 없습니다."