package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/Yuelioi/bilibili/pkg/client"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	_, filename, _, _ := runtime.Caller(0)
	env := filepath.Join(filepath.Dir(filename), "..", ".env")

	// Attempt to load the .env file
	err := godotenv.Load(env)
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}
}

type TestClient struct {
	Client *client.Client
}

func NewTestClient() *TestClient {
	LoadEnv()
	return &TestClient{
		Client: client.New(),
	}
}

func (t *TestClient) WithSessdata() *TestClient {
	value := os.Getenv("SESSDATA")
	t.Client.SESSDATA = value
	return t
}
func (t *TestClient) WithDedeUserID() *TestClient {
	value := os.Getenv("DedeUserID")
	uid, err := strconv.Atoi(value)
	if err != nil {
		t.Client.DedeUserID = 0
	} else {
		t.Client.DedeUserID = uid
	}
	return t
}

func (t *TestClient) WithCRSF() *TestClient {
	value := os.Getenv("CSRF")
	t.Client.CSRF = value
	return t
}

func (t *TestClient) WithBuvid3() *TestClient {
	value := os.Getenv("Buvid3")
	t.Client.Buvid3 = value
	return t
}
func (t *TestClient) WithAccessKey() *TestClient {
	value := os.Getenv("AccessKey")
	t.Client.AccessKey = value
	return t
}

func TestMain(m *testing.M) {
	LoadEnv()

	// 运行测试
	code := m.Run()
	os.Exit(code)
}
