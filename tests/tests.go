package tests

import (
	"bilibili/pkg/client"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

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
	sessdata := os.Getenv("SESSDATA")
	t.Client.SESSDATA = sessdata
	return t
}
func (t *TestClient) WithDedeUserID() *TestClient {
	userID := os.Getenv("DedeUserID")
	uid, err := strconv.Atoi(userID)
	if err != nil {
		t.Client.DedeUserID = 0
	} else {
		t.Client.DedeUserID = uid
	}
	return t
}

func (t *TestClient) WithCRSF() *TestClient {
	csrf := os.Getenv("CSRF")
	t.Client.CSRF = csrf
	return t
}

func TestMain(m *testing.M) {
	LoadEnv()

	// 运行测试
	code := m.Run()
	os.Exit(code)
}
