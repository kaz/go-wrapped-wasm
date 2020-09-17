package md

import (
	"bytes"
	"fmt"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

var (
	md = goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
			extension.GFM,
			extension.Footnote,
		),
		goldmark.WithParserOptions(
			parser.WithAttribute(),
		),
	)
)

func Render(src string) (string, map[string]interface{}, error) {
	buf := bytes.NewBuffer(nil)
	ctx := parser.NewContext()

	if err := md.Convert([]byte(src), buf, parser.WithContext(ctx)); err != nil {
		return "", nil, fmt.Errorf("md.Convert failed %w", err)
	}
	return buf.String(), meta.Get(ctx), nil
}
