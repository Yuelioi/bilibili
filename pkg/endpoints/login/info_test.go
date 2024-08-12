package login

import (
	"bilibili/pkg/client"
	"bilibili/tests"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserKeys(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")

	client := client.New()
	client.SESSDATA = sessdata
	service := New(client)

	resp, err := service.UserKeys()

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}

func TestSignUrl(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")

	urlStr := "https://api.bilibili.com/x/space/wbi/acc/info?mid=1850091"

	client := client.New()
	client.SESSDATA = sessdata
	service := New(client)

	var data interface{}

	newUrl, err := service.SignAndGenerateURL(urlStr)

	t.Logf("newUrl: %v\n", newUrl)

	assert.NoError(t, err)

	resp, err := service.client.HTTPClient.R().SetCookie(&http.Cookie{
		Name:  "SESSDATA",
		Value: service.client.SESSDATA,
	}).SetResult(&data).Get(newUrl)
	assert.NoError(t, err)

	t.Logf("Response Status Code: %d\n", resp.StatusCode())
	t.Logf("Response Data: %+v\n", data)
	assert.NoError(t, err)
}
func TestNavUserInfo(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")

	client := client.New()
	client.SESSDATA = sessdata
	service := New(client)

	resp, err := service.NavUserInfo()

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}

func TestUserState(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")

	client := client.New()
	client.SESSDATA = sessdata
	service := New(client)

	resp, err := service.UserState()

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}
