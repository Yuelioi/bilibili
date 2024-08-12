package audio

import (
	"bilibili/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sid = 13598

func TestCollect(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithDedeUserID()
	service := New(tc.Client)

	resp, err := service.Collect(sid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}

func TestCoin(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithDedeUserID()
	service := New(tc.Client)

	resp, err := service.Coin(sid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}

func TestAddCoin(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata().WithDedeUserID().WithCRSF()
	service := New(tc.Client)

	resp, err := service.AddCoin(sid, 1)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
}
