package tests

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestEnvFileExists(t *testing.T) {
	envFile := "../.env"

	// .env 파일 존재 여부 확인
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		t.Fatalf(".env 파일이 존재하지 않습니다: %s", envFile)
	}

	t.Logf(".env 파일이 존재합니다: %s", envFile)
}

func TestEnvFileReadable(t *testing.T) {
	envFile := "../.env"

	// .env 파일 읽기 가능한지 확인
	file, err := os.Open(envFile)
	if err != nil {
		t.Fatalf(".env 파일을 열 수 없습니다: %v", err)
	}
	defer file.Close()

	t.Log(".env 파일을 읽을 수 있습니다")
}

func TestEnvFileContent(t *testing.T) {
	envFile := "../.env"

	// .env 파일 읽기
	file, err := os.Open(envFile)
	if err != nil {
		t.Fatalf(".env 파일을 열 수 없습니다: %v", err)
	}
	defer file.Close()

	// 필수 환경변수 목록
	requiredVars := []string{
		"ENV_STATE",
		"POSTGRES_HOST",
		"POSTGRES_PORT",
		"DEV_POSTGRES_USER",
		"DEV_POSTGRES_PASSWORD",
		"DEV_POSTGRES_DB",
		"PROD_POSTGRES_USER",
		"PROD_POSTGRES_PASSWORD",
		"PROD_POSTGRES_DB",
	}

	// 찾은 환경변수들 저장
	foundVars := make(map[string]string)

	// 파일 내용 한 줄씩 읽기
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		// 빈 줄이나 주석은 건너뛰기
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// KEY=VALUE 형태 파싱
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			t.Logf("경고: %d번째 줄의 형식이 올바르지 않습니다: %s", lineNumber, line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		foundVars[key] = value
		t.Logf("발견됨: %s = %s", key, value)
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf(".env 파일 읽기 오류: %v", err)
	}

	// 필수 환경변수 존재 여부 확인
	missingVars := []string{}
	emptyVars := []string{}

	for _, reqVar := range requiredVars {
		if value, exists := foundVars[reqVar]; !exists {
			missingVars = append(missingVars, reqVar)
		} else if value == "" {
			emptyVars = append(emptyVars, reqVar)
		}
	}

	// 결과 출력
	if len(missingVars) > 0 {
		t.Errorf("누락된 필수 환경변수들: %v", missingVars)
	}

	if len(emptyVars) > 0 {
		t.Errorf("값이 비어있는 환경변수들: %v", emptyVars)
	}

	t.Logf(".env 파일에서 %d개의 환경변수를 찾았습니다", len(foundVars))

	if len(missingVars) == 0 && len(emptyVars) == 0 {
		t.Log("모든 필수 환경변수가 존재하고 값이 설정되어 있습니다")
	}
}

func TestSpecificEnvVariables(t *testing.T) {
	envFile := "../.env"

	// .env 파일에서 환경변수 로드
	envVars := loadEnvFile(t, envFile)

	// 특정 값들 확인
	tests := []struct {
		name     string
		key      string
		required bool
	}{
		{"환경 상태", "ENV_STATE", true},
		{"Postgres 호스트", "POSTGRES_HOST", true},
		{"Postgres 포트", "POSTGRES_PORT", true},
		{"개발용 DB 사용자", "DEV_POSTGRES_USER", true},
		{"개발용 DB 비밀번호", "DEV_POSTGRES_PASSWORD", true},
		{"개발용 DB 이름", "DEV_POSTGRES_DB", true},
		{"개발용 SSL 모드", "DEV_SSL_MODE", false},
		{"시간대", "TIMEZONE", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, exists := envVars[tt.key]

			if tt.required && !exists {
				t.Errorf("필수 환경변수 '%s'를 찾을 수 없습니다", tt.key)
				return
			}

			if tt.required && value == "" {
				t.Errorf("필수 환경변수 '%s'의 값이 비어있습니다", tt.key)
				return
			}

			if exists {
				t.Logf("✓ %s = %s", tt.key, value)
			} else {
				t.Logf("○ %s (선택사항, 찾을 수 없음)", tt.key)
			}
		})
	}
}

func TestEnvFileFormat(t *testing.T) {
	envFile := "../.env"

	file, err := os.Open(envFile)
	if err != nil {
		t.Fatalf(".env 파일을 열 수 없습니다: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	invalidLines := []string{}

	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		// 빈 줄이나 주석은 유효함
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// KEY=VALUE 형태 검증
		if !strings.Contains(line, "=") {
			invalidLines = append(invalidLines,
				fmt.Sprintf("%d번째 줄: '%s' ('=' 누락)", lineNumber, line))
		} else {
			parts := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(parts[0])

			// 키 이름 검증 (영문자, 숫자, 언더스코어만 허용)
			if !isValidEnvKey(key) {
				invalidLines = append(invalidLines,
					fmt.Sprintf("%d번째 줄: 유효하지 않은 키 형식 '%s'", lineNumber, key))
			}
		}
	}

	if len(invalidLines) > 0 {
		t.Errorf("유효하지 않은 .env 파일 형식:\n%s", strings.Join(invalidLines, "\n"))
	} else {
		t.Log(".env 파일 형식이 올바릅니다")
	}
}

// 헬퍼 함수들

func loadEnvFile(t *testing.T, filename string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("파일 %s를 열 수 없습니다: %v", filename, err)
	}
	defer file.Close()

	envVars := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// 빈 줄이나 주석은 건너뛰기
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// KEY=VALUE 형태로 파싱
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			envVars[key] = value
		}
	}

	return envVars
}

func isValidEnvKey(key string) bool {
	if key == "" {
		return false
	}

	// 영문자, 숫자, 언더스코어만 허용하는지 검사
	for _, char := range key {
		if !((char >= 'A' && char <= 'Z') ||
			(char >= 'a' && char <= 'z') ||
			(char >= '0' && char <= '9') ||
			char == '_') {
			return false
		}
	}

	return true
}
