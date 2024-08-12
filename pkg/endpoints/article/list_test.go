package article

import (
	"bilibili/pkg/client"
	"bilibili/tests"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArticleList(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")

	client := client.New()
	client.SESSDATA = sessdata
	service := New(client)

	resp, err := service.ArticleList(upid, 1, 30, "publish_time")

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}

func TestReadList(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")

	client := client.New()
	client.SESSDATA = sessdata
	service := New(client)

	resp, err := service.ReadList(upid, 0)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}
