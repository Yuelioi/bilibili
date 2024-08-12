package article

import (
	"bilibili/pkg/client"
	"bilibili/tests"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestArticle_Like(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")
	csrf := os.Getenv("CSRF")

	assert.NotEmpty(t, sessdata, "SESSDATA environment variable is empty")
	assert.NotEmpty(t, csrf, "CSRF environment variable is empty")

	restyClient := resty.New()

	client := &client.Client{
		HTTPClient: restyClient,
		SESSDATA:   sessdata,
		CSRF:       csrf,
	}

	service := New(client)

	resp, err := service.Like(cvid, 1)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}

func TestArticle_Coin(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")
	csrf := os.Getenv("CSRF")

	assert.NotEmpty(t, sessdata, "SESSDATA environment variable is empty")
	assert.NotEmpty(t, csrf, "CSRF environment variable is empty")

	client := client.New()

	client.SESSDATA = sessdata
	client.CSRF = csrf
	service := New(client)

	resp, err := service.Coin(cvid, 42793701, 1)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}

func TestArticle_Favorite(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")
	csrf := os.Getenv("CSRF")

	assert.NotEmpty(t, sessdata, "SESSDATA environment variable is empty")
	assert.NotEmpty(t, csrf, "CSRF environment variable is empty")

	client := client.New()

	client.SESSDATA = sessdata
	client.CSRF = csrf
	service := New(client)

	resp, err := service.Favorite(cvid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}
func TestArticle_UnFavorite(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")
	csrf := os.Getenv("CSRF")

	assert.NotEmpty(t, csrf, "CSRF environment variable is empty")

	client := client.New()

	client.SESSDATA = sessdata
	client.CSRF = csrf
	service := New(client)

	resp, err := service.UnFavorite(cvid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}
