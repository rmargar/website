package html_test

import (
	"testing"
	"time"

	"github.com/rmargar/website/pkg/domain"
	"github.com/rmargar/website/pkg/web/html"
	"github.com/stretchr/testify/assert"
)

var currentTime time.Time = time.Now()

var mockPost domain.Post = domain.Post{
	ID:       1,
	Added:    currentTime,
	Modified: currentTime,
	Author:   "Test",
	Content:  "*This* is a test",
	Title:    "Test",
	Tags:     []string{},
	Summary:  "",
	URLPath:  "test-post",
}

func TestRenderPost(t *testing.T) {
	expected := "<p><em>This</em> is a test</p>\n"
	got := html.RenderPost(mockPost, "http://test.com")
	assert.Equal(t, expected, got.ContentHTML)
}
