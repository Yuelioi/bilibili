package article

import (
	"bilibili/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArticle(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata()
	service := New(tc.Client)

	resp, err := service.Article(cvid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}
