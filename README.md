
```
gin-boiler/
├── cmd/
│   └── server/
│       └── main.go              # 애플리케이션 진입점
├── internal/
│   ├── config/
│   │   └── config.go            # 설정 관리
│   ├── handlers/
│   │   ├── auth.go             # 인증 관련 핸들러
│   │   ├── user.go             # 사용자 관련 핸들러
│   │   └── health.go           # 헬스체크 핸들러
│   ├── middleware/
│   │   ├── auth.go             # 인증 미들웨어
│   │   ├── cors.go             # CORS 미들웨어
│   │   └── logger.go           # 로깅 미들웨어
│   ├── models/
│   │   ├── user.go             # GORM 사용자 모델
│   │   ├── post.go             # GORM 포스트 모델
│   │   └── base.go             # 공통 모델 구조체
│   ├── repository/
│   │   ├── user_repo.go        # 사용자 데이터 접근 계층
│   │   ├── post_repo.go        # 포스트 데이터 접근 계층
│   │   └── interface.go        # 리포지토리 인터페이스
│   ├── service/
│   │   ├── user_service.go     # 사용자 비즈니스 로직
│   │   └── auth_service.go     # 인증 비즈니스 로직
│   ├── router/
│   │   └── router.go           # 라우터 설정
│   └── utils/
│       ├── jwt.go              # JWT 유틸리티
│       └── response.go         # 응답 유틸리티
├── pkg/
│   ├── database/
│   │   ├── connection.go       # GORM 데이터베이스 연결
│   │   └── migrate.go          # GORM 마이그레이션 헬퍼
│   └── logger/
│       └── logger.go           # 로거 설정
├── migrations/                 # Atlas 마이그레이션
│   ├── 20240101000001_initial.sql
│   └── atlas.sum
├── docs/                       # API 문서
├── tests/
│   ├── integration/
│   └── unit/
├── .env.example
├── .gitignore
├── atlas.hcl                   # Atlas 설정 파일
├── go.mod
├── go.sum
├── Makefile
└── README.md
```