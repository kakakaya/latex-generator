package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	TMP_DIR      = "tmp"
	LIB_DIR      = "lib"
	CONFIG_DIR   = "config"
	TEX_DIR      = "latex"
	APPENDIX_DIR = "appendix"
	PDF_DIR      = "pdf"
)

func main() {
	p := flag.Bool("pdf", false, "指定するとPDFまで作成します。(Windows未対応)")
	i := flag.Bool("init", false, "指定すると初期化をします。")
	flag.Parse()

	fmt.Println(*p)
	fmt.Println(*i)

	v := []bool{*p, *i}

	fmt.Println(v)

	initialize()
}

func initialize() {
	init_dirs := []string{
		TMP_DIR,
		CONFIG_DIR,
		TEX_DIR,
		APPENDIX_DIR,
		PDF_DIR,
	}

	for _, d := range init_dirs {
		if err := os.Mkdir(d, 0766); err != nil {
			log.Fatal(err)
		}
	}
}
