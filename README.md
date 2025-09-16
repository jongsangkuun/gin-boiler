# 프로젝트 이름

간단한 설명: 이 프로젝트는 Go (1.25)로 작성된 RESTful API 서버 템플릿입니다. 인증, 사용자 관리, 데이터베이스 연결, 미들웨어(로깅, CORS, 인증) 등을 포함합니다.

## 주요 기능
- 사용자 등록/로그인 (JWT 기반 인증)
- 사용자 CRUD 서비스
- 데이터베이스 연결 (마이그레이션 폴더 포함)
- 로깅 및 요청 미들웨어
- Docker / docker-compose 개발 환경 지원

## 요구사항
- Go 1.25
- Make (선택 사항)
- Docker & Docker Compose (컨테이너 실행 시)
- PostgreSQL (또는 프로젝트에 맞는 DB)

## 빠른 시작 (로컬)
1. 리포지토리 클론
   git clone <repo-url>
   cd <repo-dir>

2. 환경변수 설정
   프로젝트 루트 또는 `docker-compose`에서 사용하는 환경변수 파일(.env)을 준비합니다. 주요 변수 예시:
    - APP_PORT=8080
    - DB_HOST=localhost
    - DB_PORT=5432
    - DB_USER=postgres
    - DB_PASSWORD=yourpassword
    - DB_NAME=yourdb
    - JWT_SECRET=your_jwt_secret

3. 의존성 설치
   go mod download

4. 빌드 및 실행
   go run ./cmd/server

또는 Makefile에 정의된 명령 사용:
make run

## Docker 사용
개발용:
docker-compose up --build

프로덕션:
docker-compose -f docker-compose.prod.yml up --build

## 데이터베이스 마이그레이션
프로젝트에 `migrations` 디렉터리가 포함되어 있습니다. 사용하시는 마이그레이션 도구(atlas 또는 goose 등)에 맞춰 적용하세요.
예:
atlas migrate apply --dir file://migrations

(또는 프로젝트에 맞는 도구/명령을 사용)

## 프로젝트 디렉터리 트리
프로젝트 루트의 주요 파일/디렉터리 구조는 다음과 같습니다:
```
├── cmd
│   └── server
│       └── main.go
├── internal
│   ├── config
│   │   └── config.go
│   ├── database
│   │   └── connection.go
│   ├── dto
│   │   ├── auth_dto.go
│   │   └── user_dto.go
│   ├── logger
│   │   └── logger.go
│   ├── middleware
│   │   ├── auth.go
│   │   ├── cors.go
│   │   └── logger.go
│   ├── models
│   │   ├── base.go
│   │   └── user.go
│   ├── repository
│   │   ├── interface.go
│   │   └── user_repo.go
│   ├── router
│   │   ├── router.go
│   │   ├── auth_router.go
│   │   └── user_router.go
│   ├── service
│   │   ├── auth_service.go
│   │   └── user_service.go
│   └── utils
│       ├── crypto.go
│       ├── jwt.go
│       └── response.go
├── migrations
├── docs
├── tests
│   └── config_test.go
├── .gitignore
├── .git
├── .idea
├── Makefile
├── Dockerfile
├── docker-compose.yml
├── docker-compose.prod.yml
├── docker-compose.override.yml
├── atlas.hcl
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## 환경 변수
.env 파일 생성
```aiignore
# 공통 설정
ENV_STATE=env

# 공통 Postgresql 설정
POSTGRES_HOST=
POSTGRES_PORT=
TIMEZONE=

PORT=
JWT_TOKEN=

# dev 설정
DEV_DEBUG=
DEV_POSTGRES_USER=
DEV_POSTGRES_PASSWORD=
DEV_POSTGRES_DB=
DEV_POSTGRES_SSL_MODE=
DEV_POSTGRES_MAX_OPEN_CONNS=
DEV_POSTGRES_MAX_IDLE_CONNS=
DEV_POSTGRES_CONN_MAX_LIFETIME=
DEV_POSTGRES_CONN_MAX_IDLE_TIME=

# Atlas용 DATABASE_URL (Docker 내부 네트워크)
DEV_DATABASE_URL=

# Atlas용 DATABASE_URL (localhost 접근용)
DEV_DATABASE_URL_LOCAL=

# prod 설정
PROD_DEBUG=
PROD_POSTGRES_USER=
PROD_POSTGRES_PASSWORD=
PROD_POSTGRES_DB=
PROD_POSTGRES_SSL_MODE=
PROD_POSTGRES_MAX_OPEN_CONNS=
PROD_POSTGRES_MAX_IDLE_CONNS=
PROD_POSTGRES_CONN_MAX_LIFETIME=
PROD_POSTGRES_CONN_MAX_IDLE_TIME=

# Atlas용 DATABASE_URL (Docker 내부 네트워크)
PROD_DATABASE_URL=

# Atlas용 DATABASE_URL (localhost 접근용)
PROD_DATABASE_URL_LOCAL=
```