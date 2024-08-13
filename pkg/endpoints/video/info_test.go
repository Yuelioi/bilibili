package video

import (
	"bilibili/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVideoInfo(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata()
	service := New(tc.Client)

	resp, err := service.Info(seasonAid, "")
	t.Logf("Response: %+v", resp)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}
func TestVideoDetail(t *testing.T) {
	tc := tests.NewTestClient().WithSessdata()
	service := New(tc.Client)

	resp, err := service.Detail(seasonAid, "")
	t.Logf("Response: %+v", resp)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}

func TestVideoDesc(t *testing.T) {
	tc := tests.NewTestClient()
	service := New(tc.Client)

	resp, err := service.Description(seasonAid, "")
	t.Logf("Response: %+v", resp)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}
func TestPageList(t *testing.T) {
	tc := tests.NewTestClient()
	service := New(tc.Client)

	resp, err := service.PageList(pagesAid, "")
	t.Logf("Response: %+v", resp)

	assert.NoError(t, err)
	assert.GreaterOrEqual(t, resp.Code, 0)
}
