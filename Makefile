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

docker-build:
	docker compose build

# Docker 컨테이너 중지 (docker compose 사용)
docker-down:
	docker compose down