package html

import (
	"html/template"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHTML(t *testing.T) {

	templatesPath := filepath.Join("./", "..", "..", "..", "templates")

	expected := []string{
		path.Join(templatesPath, "partial", "disqus.tpl"),
		path.Join(templatesPath, "partial", "footer.tpl"),
		path.Join(templatesPath, "partial", "heading.tpl"),
		path.Join(templatesPath, "partial", "info.tpl"),
		path.Join(templatesPath, "partial", "posts.tpl"),
	}

	instance := &HTML{
		config:    HTMLConfig{TemplatesPath: "../../../templates", DateFormat: "2 Jan 2006"},
		Templates: make(map[string]*template.Template)}
	partials, err := instance.GetPartials()
	require.NoError(t, err)
	assert.Equal(t, expected, partials)

	err = instance.InitTemplates()
	require.NoError(t, err)

}
