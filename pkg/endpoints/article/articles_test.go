package article

import (
	"bilibili/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArticles(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata()
	service := New(tc.Client)

	resp, err := service.Articles(rlid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}
