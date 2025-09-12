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

	// env 파일 존재 여부 확인
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		t.Fatalf(".env file does not exist: %s", envFile)
	}

	t.Logf(".env file exists: %s", envFile)
}

func TestEnvFileReadable(t *testing.T) {
	envFile := "../.env"

	// .env 파일 읽기 가능한지 확인
	file, err := os.Open(envFile)
	if err != nil {
		t.Fatalf("Cannot open .env file: %v", err)
	}
	defer file.Close()

	t.Log(".env file is readable")
}

func TestEnvFileContent(t *testing.T) {
	envFile := "../.env"

	// ../.env 파일 읽기
	file, err := os.Open(envFile)
	if err != nil {
		t.Fatalf("Cannot open ../.env file: %v", err)
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
			t.Logf("Warning: Invalid format at line %d: %s", lineNumber, line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		foundVars[key] = value
		t.Logf("Found: %s = %s", key, value)
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("Error reading ../.env file: %v", err)
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
		t.Errorf("Missing required environment variables: %v", missingVars)
	}

	if len(emptyVars) > 0 {
		t.Errorf("Empty environment variables: %v", emptyVars)
	}

	t.Logf("Found %d environment variables in ../.env file", len(foundVars))

	if len(missingVars) == 0 && len(emptyVars) == 0 {
		t.Log("All required environment variables are present and not empty")
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
		{"Environment State", "ENV_STATE", true},
		{"Postgres Host", "POSTGRES_HOST", true},
		{"Postgres Port", "POSTGRES_PORT", true},
		{"Dev DB User", "DEV_POSTGRES_USER", true},
		{"Dev DB Password", "DEV_POSTGRES_PASSWORD", true},
		{"Dev DB Name", "DEV_POSTGRES_DB", true},
		{"Dev SSL Mode", "DEV_SSL_MODE", false},
		{"Timezone", "TIMEZONE", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, exists := envVars[tt.key]

			if tt.required && !exists {
				t.Errorf("Required environment variable '%s' not found", tt.key)
				return
			}

			if tt.required && value == "" {
				t.Errorf("Required environment variable '%s' is empty", tt.key)
				return
			}

			if exists {
				t.Logf("✓ %s = %s", tt.key, value)
			} else {
				t.Logf("○ %s (optional, not found)", tt.key)
			}
		})
	}
}

func TestEnvFileFormat(t *testing.T) {
	envFile := "../.env"

	file, err := os.Open(envFile)
	if err != nil {
		t.Fatalf("Cannot open .env file: %v", err)
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
				fmt.Sprintf("Line %d: '%s' (missing '=')", lineNumber, line))
		} else {
			parts := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(parts[0])

			// 키 이름 검증 (영문자, 숫자, 언더스코어만 허용)
			if !isValidEnvKey(key) {
				invalidLines = append(invalidLines,
					fmt.Sprintf("Line %d: Invalid key format '%s'", lineNumber, key))
			}
		}
	}

	if len(invalidLines) > 0 {
		t.Errorf("Invalid .env file format:\n%s", strings.Join(invalidLines, "\n"))
	} else {
		t.Log(".env file format is valid")
	}
}

// 헬퍼 함수들

func loadEnvFile(t *testing.T, filename string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Cannot open file %s: %v", filename, err)
	}
	defer file.Close()

	envVars := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

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
