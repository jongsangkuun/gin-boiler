package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"gin-boiler/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(env config.Env) (*gorm.DB, error) {
	dsn := buildPostgresDSN(&env)
	log.Printf("PostgreSQL DSN: %s", dsn)
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}

	// GORM으로 PostgreSQL 연결
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("PostgreSQL GORM 연결 실패: %v", err)
	}

	// 기본 SQL DB 인스턴스 가져오기 (연결 풀 설정용)
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("SQL DB 인스턴스 가져오기 실패: %v", err)
	}

	// 연결 풀 설정
	if err := setupConnectionPool(sqlDB, env.DbConnectionPool); err != nil {
		return nil, fmt.Errorf("연결 풀 설정 실패: %v", err)
	}

	// 연결 테스트
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("PostgreSQL 연결 테스트 실패: %v", err)
	}

	DB = db
	log.Printf("PostgreSQL GORM 데이터베이스 연결 성공")
	return db, nil
}

// buildPostgresDSN PostgreSQL 연결 문자열을 생성합니다
func buildPostgresDSN(env *config.Env) string {
	return fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s",
		env.DbConfig.Host,
		env.DbConfig.User,
		env.DbConfig.Password,
		env.DbConfig.DbName,
		env.DbConfig.Port,
	)
}

// setupConnectionPool 데이터베이스 연결 풀을 설정합니다
func setupConnectionPool(sqlDB *sql.DB, poolConfig config.DbConnectionPool) error {
	// 최대 유휴 연결 수 설정
	if poolConfig.MaxIdleConns != "" {
		maxIdleConns, err := strconv.Atoi(poolConfig.MaxIdleConns)
		if err != nil {
			return fmt.Errorf("MaxIdleConns 변환 실패: %v", err)
		}
		sqlDB.SetMaxIdleConns(maxIdleConns)
	}

	// 최대 연결 수 설정
	if poolConfig.MaxOpenConns != "" {
		maxOpenConns, err := strconv.Atoi(poolConfig.MaxOpenConns)
		if err != nil {
			return fmt.Errorf("MaxOpenConns 변환 실패: %v", err)
		}
		sqlDB.SetMaxOpenConns(maxOpenConns)
	}

	// 연결 최대 생존 시간 설정
	if poolConfig.ConnMaxLifetime != "" {
		connMaxLifetime, err := time.ParseDuration(poolConfig.ConnMaxLifetime)
		if err != nil {
			return fmt.Errorf("ConnMaxLifetime 변환 실패: %v", err)
		}
		sqlDB.SetConnMaxLifetime(connMaxLifetime)
	}

	// 연결 최대 유휴 시간 설정
	if poolConfig.ConnMaxIdleTime != "" {
		connMaxIdleTime, err := time.ParseDuration(poolConfig.ConnMaxIdleTime)
		if err != nil {
			return fmt.Errorf("ConnMaxIdleTime 변환 실패: %v", err)
		}
		sqlDB.SetConnMaxIdleTime(connMaxIdleTime)
	}

	log.Printf("데이터베이스 연결 풀 설정 완료")
	return nil
}

// Close 데이터베이스 연결을 종료합니다
func Close() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return fmt.Errorf("SQL DB 인스턴스 가져오기 실패: %v", err)
		}
		return sqlDB.Close()
	}
	return nil
}
