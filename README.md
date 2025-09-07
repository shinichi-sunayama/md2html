# md2html

📄 Markdown → HTML 変換ツール（CLI + Web GUI対応）

---

## 🧩 概要

`md2html` は、Markdown ファイルを HTML に変換する Go 製ツールです。以下の2通りの使い方が可能です：

- ✅ **CLIモード**：指定したフォルダ内の `.md` ファイルを一括で `.html` に変換
- ✅ **Web GUIモード**：ブラウザ上で `.md` ファイルをアップロードして変換＆プレビュー

---

## 🚀 特徴と機能

- ✅ 複数の `.md` ファイルを一括で HTML に変換（CLIモード）
- ✅ Webブラウザで Markdown をアップロード → 変換＆プレビュー（GUIモード）
- ✅ CSSスタイル（`styles/custom.css`）を自動適用
- ✅ Docker対応、簡単にGUIを立ち上げ可能
- ✅ Goのみで完結。外部ライブラリ最小限

---

## 🧱 使用技術・フレームワーク

| 種別             | 使用技術                     |
|------------------|------------------------------|
| 言語             | Go 1.20+                      |
| HTMLレンダリング | `html/template`（標準）       |
| Webフレームワーク| [Echo](https://echo.labstack.com/) v4 |
| Markdown変換     | `github.com/gomarkdown/markdown` |
| スタイル         | 独自CSS（`styles/custom.css`）|
| GUIテンプレート  | `text/template`＋`.tmpl`形式 |
| Docker対応       | コンテナでCLI or GUIを実行可 |

## 📁 ディレクトリ構成

```plaintext
md2html/
├─ cmd/                   # CLIエントリーポイント（root.go）
│   └─ root.go
│
├─ examples/             # サンプルMarkdown
│   └─ sample.md
│
├─ internal/
│   └─ convert/           # Markdown→HTML変換処理
│       └─ convert.go
│
├─ styles/
│   └─ custom.css         # HTML出力に使われるカスタムCSS
│
├─ webgui/
│   ├─ cmd/server/        # Web GUIのHTTPサーバ
│   │   └─ main.go
│   └─ web/               # GUI用HTMLテンプレート＆素材
│       ├─ index.tmpl
│       ├─ preview.tmpl
│       └─ screenshot.png
│
├─ .gitignore
├─ Dockerfile
├─ go.mod / go.sum
├─ main.go                # CLI起動用エントリーポイント
├─ md2html                # CLIバイナリ（例）
└─ README.md              # このファイル

🚀 機能一覧
✅ CLIモード
bash
go run main.go -in ./examples -out ./output
オプション	内容
-in	入力ディレクトリ（Markdown）
-out	出力ディレクトリ（HTML）
-style	CSSファイルの指定（任意）

変換された .html ファイルには <head> に custom.css が自動で読み込まれます。

✅ Web GUIモード
bash
コピーする
編集する
go run webgui/cmd/server/main.go
起動後、ブラウザでアクセス：
http://localhost:8080

機能	説明
アップロード	.md ファイルをドラッグ＆ドロップ、または選択
HTMLプレビュー	即座に変換結果を確認可能（preview.tmplを使用）
スクリーンショット	web/screenshot.png にGUIイメージ付き
カスタムCSS対応	styles/custom.css を読み込んだスタイルで表示

🐳 Dockerで動かす
docker build -t md2html .
docker run -p 8080:8080 md2html

🧪 テスト実行（ユニットテスト）
go test ./internal/...