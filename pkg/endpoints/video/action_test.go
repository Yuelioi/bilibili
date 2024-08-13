package video

import (
	"bilibili/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLike(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithCRSF().WithBuvid3()
	service := New(tc.Client)

	resp, err := service.Like(aid, bvid, 1)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}

func TestLikeApp(t *testing.T) {
	tc := tests.NewTestClient().WithAccessKey()
	service := New(tc.Client)

	resp, err := service.LikeApp(aid, 0)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	// assert.GreaterOrEqual(t, resp.Code, 0)
}

func TestDisLikeApp(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithAccessKey()
	service := New(tc.Client)

	resp, err := service.DislikeApp(aid, 1)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	// assert.GreaterOrEqual(t, resp.Code, 0)
}

func TestCoin(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithBuvid3().WithCRSF()
	service := New(tc.Client)

	resp, err := service.Coin(aid, bvid, 1, 0)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}
func TestCoinApp(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithBuvid3().WithCRSF()
	service := New(tc.Client)

	resp, err := service.CoinApp(aid, 1, 0)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	// assert.GreaterOrEqual(t, resp.Code, 0)
}

func TestCoinsStatus(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithAccessKey()
	service := New(tc.Client)

	resp, err := service.CoinsStatus(aid, bvid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}

// TestCollect tests the Collect function.
func TestCollect(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithAccessKey().WithCRSF()
	service := New(tc.Client)

	resp, err := service.Collect(aid, mlid, "")

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}

// TestWebCollect tests the WebCollect function.
func TestCollectWeb(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithCRSF()
	service := New(tc.Client)

	resp, err := service.CollectWeb(aid, mlid, "")

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}

// TestIsFavoured tests the IsFavoured function.
func TestIsFavoured(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithAccessKey()
	service := New(tc.Client)

	resp, err := service.IsFavoured(aid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}

// TestTripleLike tests the TripleLike function.
func TestTripleLike(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithCRSF()
	service := New(tc.Client)

	resp, err := service.TripleLike(aid, bvid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}

// TestAppTripleLike tests the AppTripleLike function.
func TestAppTripleLike(t *testing.T) {
	tc := tests.NewTestClient().WithAccessKey()
	service := New(tc.Client)

	resp, err := service.TripleLikeApp(aid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	// assert.GreaterOrEqual(t, resp.Code, 0)
}

func TestShare(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithCRSF()
	service := New(tc.Client)

	resp, err := service.Share(aid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	// assert.GreaterOrEqual(t, resp.Code, 0)
}
