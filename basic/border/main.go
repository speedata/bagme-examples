package main

import (
	"log"
	"os"

	"github.com/speedata/bagme/document"
)

func read(filename string) string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func dothings() error {
	d, err := document.New("out.pdf")
	if err != nil {
		return err
	}

	if err = d.ReadCSSFile("styles.css"); err != nil {
		return err
	}
	pageSize, err := d.PageSize()
	if err != nil {
		return err
	}
	wd := pageSize.Width - pageSize.MarginLeft - pageSize.MarginRight
	x := pageSize.MarginLeft
	y := pageSize.Height - pageSize.MarginTop

	if err = d.OutputAt(read("chunk.html"), wd, x, y); err != nil {
		return err
	}

	return d.Finish()
}

func main() {
	if err := dothings(); err != nil {
		log.Fatal(err)
	}
}
