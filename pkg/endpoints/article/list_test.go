package article

import (
	"testing"

	"github.com/Yuelioi/bilibili/tests"

	"github.com/stretchr/testify/assert"
)

func TestArticleList(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata()
	service := New(tc.Client)

	resp, err := service.ArticleList(upid, 1, 30, "publish_time")

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}

func TestReadList(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata()
	service := New(tc.Client)

	resp, err := service.ReadList(upid, 0)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}
