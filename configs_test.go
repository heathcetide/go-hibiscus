package hibiscus

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func writeTempDotEnv(content string) {
	_ = os.WriteFile(".env", []byte(content), 0644)
}

func cleanupDotEnv() {
	_ = os.Remove(".env")
}

func TestLookupEnv_FromDotEnv(t *testing.T) {
	defer cleanupDotEnv()

	writeTempDotEnv(`
# This is a comment
APP_NAME=MyApp
PORT=8080

EMPTY=
INVALIDLINE
	`)

	envCache = NewExpiredLRUCache[string, string](10, 1*time.Minute)

	val, found := LookupEnv("APP_NAME")
	assert.True(t, found)
	assert.Equal(t, "MyApp", val)

	val, found = LookupEnv("PORT")
	assert.True(t, found)
	assert.Equal(t, "8080", val)

	val, found = LookupEnv("EMPTY")
	assert.True(t, found)
	assert.Equal(t, "", val)

	val, found = LookupEnv("INVALID") // 不存在
	assert.False(t, found)
}

func TestLookupEnv_FromOS(t *testing.T) {
	defer cleanupDotEnv()

	os.Unsetenv("MY_ENV_VAR")
	t.Setenv("MY_ENV_VAR", "from-os")

	envCache = NewExpiredLRUCache[string, string](10, 1*time.Minute)

	val, found := LookupEnv("MY_ENV_VAR")
	assert.True(t, found)
	assert.Equal(t, "from-os", val)
}

func TestLookupEnv_CacheHit(t *testing.T) {
	defer cleanupDotEnv()

	writeTempDotEnv(`CACHED_KEY=123`)
	envCache = NewExpiredLRUCache[string, string](10, 1*time.Minute)

	// 第一次读取：会读文件并写入缓存
	val, found := LookupEnv("CACHED_KEY")
	assert.True(t, found)
	assert.Equal(t, "123", val)

	// 第二次读取：来自缓存，不读取文件（测试无法直接验证是否访问文件，但逻辑路径会走缓存）
	val, found = LookupEnv("CACHED_KEY")
	assert.True(t, found)
	assert.Equal(t, "123", val)
}

func TestLookupEnv_NotFound(t *testing.T) {
	defer cleanupDotEnv()

	writeTempDotEnv(`UNRELATED_KEY=xyz`)
	envCache = NewExpiredLRUCache[string, string](10, 1*time.Minute)

	val, found := LookupEnv("NON_EXISTENT")
	assert.False(t, found)
	assert.Equal(t, "", val)
}

// 定义一个测试用配置结构体
type MyConfig struct {
	AppName string `env:"APP_NAME"`
	Port    int    `env:"APP_PORT"`
	Debug   bool   `env:"DEBUG"`
	Ignored string `env:"-"` // 不应被设置
	Default string // 无 tag，自动匹配 "Default"
}

func TestLoadEnvs(t *testing.T) {
	// 设置模拟 .env 内容
	os.WriteFile(".env", []byte(`
APP_NAME=TestApp
APP_PORT=8888
DEBUG=true
DEFAULT=default-val
`), 0644)
	defer os.Remove(".env")

	envCache = NewExpiredLRUCache[string, string](10, 1*time.Minute)

	var cfg MyConfig
	LoadEnvs(&cfg)

	assert.Equal(t, "TestApp", cfg.AppName)
	assert.Equal(t, 8888, cfg.Port)
	assert.True(t, cfg.Debug)
	assert.Equal(t, "default-val", cfg.Default)
	assert.Equal(t, "", cfg.Ignored)
}

func TestLoadEnvs_NilPointer(t *testing.T) {
	// 验证空指针不会 panic
	LoadEnvs(nil)
}
