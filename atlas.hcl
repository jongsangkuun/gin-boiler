# GORM 모델로부터 스키마 생성
data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./internal/models",
    "--dialect", "postgres",
  ]
}

# 개발 환경 설정 (GORM 기반)
env "dev-local" {
  src = data.external_schema.gorm.url
  url = "postgres://boilerplate_user:boilerplate@localhost:5432/boilerplate-dev-db?sslmode=disable"
  dev = "docker://postgres/15/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

# 개발 환경 설정 (Docker 내부)
env "dev" {
  src = data.external_schema.gorm.url
  url = "postgres://boilerplate_user:boilerplate@boiler-postgres-db:5432/boilerplate-dev-db?sslmode=disable"
  dev = "docker://postgres/15/dev"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

# 프로덕션 환경 설정
env "prod" {
  src = data.external_schema.gorm.url
  url = "postgres://boilerplate_user:boilerplate@localhost:5432/boilerplate-prod-db?sslmode=disable"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}