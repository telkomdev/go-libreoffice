package main

import (
	"fmt"
	"os"

	"github.com/telkomdev/go-libreoffice"
)

// docker run --rm -v /Users/wuriyanto/Documents/go-projects/go-libreoffice/_examples/docx_to_pdf/data/:/usr/app/data/ -e INPUT_FILE_NAME=hello.txt golo-ex-docxtopdf
func main() {
	println("golo (Go LibreOffice)")

	args := os.Args
	if len(args) < 2 {
		fmt.Println("required at least one argument")
		os.Exit(1)
	}

	docx, err := os.Open(args[1])
	if err != nil {
		fmt.Println("error open file", err.Error())
		os.Exit(1)
	}

	// pdfOut, err := os.Create("data/out_pdf.pdf")
	// if err != nil {
	// 	fmt.Println("error create file", err.Error())
	// 	os.Exit(1)
	// }

	htmlOut, err := os.Create("data/out_html.html")
	if err != nil {
		fmt.Println("error create file", err.Error())
		os.Exit(1)
	}

	// textOut, err := os.Create("data/out_text.txt")
	// if err != nil {
	// 	fmt.Println("error create file", err.Error())
	// 	os.Exit(1)
	// }

	defer func() {
		docx.Close()
	}()

	// defer func() {
	// 	pdfOut.Close()
	// }()

	defer func() {
		htmlOut.Close()
	}()

	// defer func() {
	// 	textOut.Close()
	// }()

	// err = libreoffice.ToPdf(docx, pdfOut)
	// if err != nil {
	// 	fmt.Println("error convert file to pdf", err.Error())
	// 	os.Exit(1)
	// }

	err = libreoffice.ToHTML(docx, htmlOut)
	if err != nil {
		fmt.Println("error convert file html", err.Error())
		os.Exit(1)
	}

	// err = libreoffice.ToTEXT(docx, textOut)
	// if err != nil {
	// 	fmt.Println("error convert file text", err.Error())
	// 	os.Exit(1)
	// }
}
