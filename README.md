# md2html

`md2html` は、MarkdownファイルをHTMLに変換するGo製CLIツールです。

## 特徴
- Markdown → HTML変換
- `--standalone` フルHTML生成
- `--title` でHTMLタイトル指定
- `--css` 独自CSS適用
- `--stdin` 標準入力対応
- エラー時は終了コード `1` を返す

---

## インストール
```bash
go install github.com/yourname/md2html@latest
