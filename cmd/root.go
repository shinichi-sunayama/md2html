package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/shinichi.sunayama/md2html/internal/convert"
	"github.com/spf13/cobra"
)

var (
	outPath    string
	title      string
	standalone bool
	cssPath    string
	fromStdin  bool
)

var rootCmd = &cobra.Command{
	Use:   "md2html <input.md> [-o output.html] [--title TITLE] [--standalone] [--css file.css] [--stdin]",
	Short: "Convert Markdown to HTML",
	Args:  cobra.ArbitraryArgs, // --stdin のときは引数なし可
	RunE: func(cmd *cobra.Command, args []string) error {
		var mdBytes []byte
		var inName string

		// 入力: stdin or ファイル
		if fromStdin {
			b, err := io.ReadAll(os.Stdin)
			if err != nil {
				return fmt.Errorf("failed to read stdin: %w", err)
			}
			mdBytes = b
			inName = "stdin"
			if outPath == "" {
				outPath = "out.html"
			}
		} else {
			if len(args) != 1 {
				return fmt.Errorf("input file is required (or use --stdin)")
			}
			in := args[0]
			if _, err := os.Stat(in); err != nil {
				return fmt.Errorf("input not found: %s", in)
			}
			b, err := os.ReadFile(in)
			if err != nil {
				return fmt.Errorf("read failed: %w", err)
			}
			mdBytes = b
			inName = filepath.Base(in)
			if outPath == "" {
				base := inName[:len(inName)-len(filepath.Ext(inName))]
				outPath = base + ".html"
			}
		}

		// CSS読み込み（任意）
		var customCSS string
		if cssPath != "" {
			b, err := os.ReadFile(cssPath)
			if err != nil {
				return fmt.Errorf("failed to read css file: %w", err)
			}
			customCSS = string(b)
		}

		// 変換
		html, err := convert.MarkdownToHTML(mdBytes, standalone, title, customCSS)
		if err != nil {
			return err
		}

		// 出力
		if err := os.WriteFile(outPath, []byte(html), 0644); err != nil {
			return fmt.Errorf("write failed: %w", err)
		}
		fmt.Printf("✓ wrote %s (from %s)\n", outPath, inName)
		return nil
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&outPath, "output", "o", "", "output HTML path")
	rootCmd.Flags().StringVar(&title, "title", "", "HTML title")
	rootCmd.Flags().BoolVar(&standalone, "standalone", false, "emit full HTML document")
	rootCmd.Flags().StringVar(&cssPath, "css", "", "path to custom CSS (inlined)")
	rootCmd.Flags().BoolVar(&fromStdin, "stdin", false, "read Markdown from STDIN")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1) // エラー時は非0終了
	}
}
