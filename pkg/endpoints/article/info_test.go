package article

import (
	"bilibili/pkg/client"
	"bilibili/tests"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArticle(t *testing.T) {
	tests.LoadEnv()
	sessdata := os.Getenv("SESSDATA")

	client := client.New()
	client.SESSDATA = sessdata
	service := New(client)

	resp, err := service.Article(cvid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}
