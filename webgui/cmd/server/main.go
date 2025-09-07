package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shinichi.sunayama/md2html/internal/convert"
)

// Markdown -> HTML変換（CLIと共通）
func mdToHTML(md []byte) ([]byte, error) {
	htmlStr, err := convert.MarkdownToHTML(md, false, "", "")
	if err != nil {
		return nil, err
	}
	return []byte(htmlStr), nil
}

func main() {
	r := gin.Default()

	// HTMLテンプレート読み込み（2つ）
	r.LoadHTMLGlob("web/*.tmpl")
	r.Static("/static", "./web/static")

	// /healthz エンドポイント
	r.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// GET / → メイン画面（フォーム＋空のプレビュー）
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"HTML": template.HTML(""),
		})
	})

	// POST /preview → Markdownファイル → HTMLプレビュー（部分描画）
	r.POST("/preview", func(c *gin.Context) {
		file, _, err := c.Request.FormFile("mdfile")
		if err != nil {
			c.String(http.StatusBadRequest, "ファイル取得エラー: %v", err)
			return
		}
		defer file.Close()

		buf := new(bytes.Buffer)
		if _, err := io.Copy(buf, file); err != nil {
			c.String(http.StatusInternalServerError, "読み込みエラー: %v", err)
			return
		}

		htmlBytes, err := mdToHTML(buf.Bytes())
		if err != nil {
			c.String(http.StatusInternalServerError, "変換エラー: %v", err)
			return
		}

		c.HTML(http.StatusOK, "preview.tmpl", gin.H{
			"HTML": template.HTML(htmlBytes),
		})
	})

	// POST /export → テキスト入力からHTMLファイルをダウンロード
	r.POST("/export", func(c *gin.Context) {
		mdText := c.PostForm("mdtext")

		htmlBytes, err := mdToHTML([]byte(mdText))
		if err != nil {
			c.String(http.StatusInternalServerError, "変換エラー: %v", err)
			return
		}

		c.Header("Content-Disposition", "attachment; filename=output.html")
		c.Data(http.StatusOK, "text/html; charset=utf-8", htmlBytes)
	})

	// サーバ起動
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
