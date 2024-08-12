package audio

import (
	"bilibili/pkg/client"
	"bilibili/tests"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sid = 13598

func TestCollect(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")
	userID := os.Getenv("DedeUserID")

	client := client.New()
	client.SESSDATA = sessdata
	uid, err := strconv.Atoi(userID)
	assert.NoError(t, err)

	client.DedeUserID = uid
	service := New(client)

	resp, err := service.Collect(sid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}

func TestCoin(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")
	userID := os.Getenv("DedeUserID")

	client := client.New()
	client.SESSDATA = sessdata
	uid, err := strconv.Atoi(userID)
	assert.NoError(t, err)

	client.DedeUserID = uid
	service := New(client)

	resp, err := service.Coin(sid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}

func TestAddCoin(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")
	userID := os.Getenv("DedeUserID")
	csrf := os.Getenv("CSRF")

	client := client.New()

	client.SESSDATA = sessdata
	uid, err := strconv.Atoi(userID)
	assert.NoError(t, err)
	client.DedeUserID = uid
	client.CSRF = csrf
	service := New(client)

	resp, err := service.AddCoin(sid, 1)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}
