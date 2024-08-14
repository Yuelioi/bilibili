package article

import (
	"testing"

	"github.com/Yuelioi/bilibili/tests"

	"github.com/stretchr/testify/assert"
)

func TestArticle_Like(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithCRSF()
	service := New(tc.Client)

	resp, err := service.Like(cvid, 1)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}

func TestArticle_Coin(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithCRSF()
	service := New(tc.Client)

	resp, err := service.Coin(cvid, 42793701, 1)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}

func TestArticle_Favorite(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithCRSF()
	service := New(tc.Client)

	resp, err := service.Favorite(cvid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}
func TestArticle_UnFavorite(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithCRSF()
	service := New(tc.Client)

	resp, err := service.UnFavorite(cvid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}
