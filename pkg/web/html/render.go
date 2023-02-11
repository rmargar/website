package html

import (
	"bytes"
	"fmt"

	"github.com/rmargar/website/pkg/domain"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

type RenderData struct {
	Posts      []domain.Post
	CurrentURL string
}

type HTMLPost struct {
	Post        domain.Post
	CurrentURL  string
	ContentHTML string
}

var md goldmark.Markdown = goldmark.New(
	goldmark.WithExtensions(extension.GFM, extension.Footnote, extension.Typographer),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(),
)

func RenderPost(post domain.Post, url string) HTMLPost {
	var buffer bytes.Buffer
	var contentHTML string

	if err := md.Convert([]byte(post.Content), &buffer); err != nil {
		contentHTML = fmt.Sprintf("<p>Error rendering Markdown: <code>%s</code></p>", err.Error())
	} else {
		contentHTML = buffer.String()
	}

	return HTMLPost{
		Post:        post,
		CurrentURL:  url,
		ContentHTML: contentHTML,
	}
}
