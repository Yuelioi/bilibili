package video

import (
	"bilibili/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeasonArchives(t *testing.T) {
	tc := tests.NewTestClient()
	service := New(tc.Client)

	resp, err := service.SeasonsArchives(uid, sid, false, 0, 0, "", "", "", 0)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}
func TestSeasonSeries(t *testing.T) {
	tc := tests.NewTestClient()
	service := New(tc.Client)

	resp, err := service.SeasonsSeries(uid, 0, 0, "", "", 0)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}

func TestSeasonSeriesList(t *testing.T) {
	tc := tests.NewTestClient()
	service := New(tc.Client)

	resp, err := service.SeasonsSeriesList(uid, 0, 0, "", 0, "")

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}
func TestSeries(t *testing.T) {
	tc := tests.NewTestClient()
	service := New(tc.Client)

	resp, err := service.Series(listid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}
func TestArchives(t *testing.T) {
	tc := tests.NewTestClient()
	service := New(tc.Client)

	resp, err := service.Archives(uid, listid, "desc", 1, 20, uid)

	t.Logf("Response: %+v", resp)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}
