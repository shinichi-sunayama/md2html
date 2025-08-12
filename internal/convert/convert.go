package convert

import (
	"bytes"
	"fmt"
	"html"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	htmlrender "github.com/yuin/goldmark/renderer/html"
)

// MarkdownToHTML converts Markdown bytes to HTML string.
// If standalone is true, wrap with full HTML document and embed CSS.
func MarkdownToHTML(md []byte, standalone bool, title, customCSS string) (string, error) {
	mdParser := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,      // テーブル/チェックボックスなど
			extension.Footnote, // 脚注
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			htmlrender.WithHardWraps(),
			htmlrender.WithXHTML(),
			htmlrender.WithUnsafe(), // 生HTML許可（自己責任）
		),
	)

	var buf bytes.Buffer
	if err := mdParser.Convert(md, &buf); err != nil {
		return "", fmt.Errorf("convert failed: %w", err)
	}
	body := buf.String()
	if !standalone {
		return body, nil
	}

	if title == "" {
		title = "Document"
	}

	defaultCSS := `
body{font-family:system-ui,-apple-system,Segoe UI,Roboto,Helvetica,Arial,'Noto Sans JP',sans-serif;max-width:860px;margin:40px auto;padding:0 16px;line-height:1.7;}
pre,code{font-family:ui-monospace,SFMono-Regular,Menlo,Consolas,Monaco,monospace;}
pre{background:#f6f8fa;padding:12px;border-radius:8px;overflow:auto;}
code{background:#f6f8fa;padding:2px 6px;border-radius:6px;}
table{border-collapse:collapse;}
td,th{border:1px solid #e3e3e3;padding:6px 10px;}
blockquote{border-left:4px solid #e3e3e3;margin:0;padding:0 12px;color:#555;}
h1,h2,h3{line-height:1.4}
`

	finalCSS := defaultCSS
	if customCSS != "" {
		finalCSS += "\n/* --- custom.css --- */\n" + customCSS + "\n"
	}

	template := `<!doctype html>
<html lang="ja">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<title>%s</title>
<style>
%s
</style>
</head>
<body>
%s
</body>
</html>`
	return fmt.Sprintf(template, html.EscapeString(title), finalCSS, body), nil
}
