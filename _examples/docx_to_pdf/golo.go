package main

import (
	"fmt"
	"os"

	"github.com/telkomdev/go-libreoffice"
)

func main() {
	println("golo (Go LibreOffice)")
	// println(os.TempDir())

	// content := []byte("temporary file's content")
	// tmpfile, err := ioutil.TempFile("", "example")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer os.Remove(tmpfile.Name()) // clean up

	// if _, err := tmpfile.Write(content); err != nil {
	// 	log.Fatal(err)
	// }

	// println(fmt.Sprintf("%s.pdf", tmpfile.Name()))

	// if err := tmpfile.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	docx, err := os.Open("mydoc.docx")
	if err != nil {
		fmt.Println("error open file", err.Error())
		os.Exit(1)
	}

	pdfOut, err := os.Create("data/out.pdf")
	if err != nil {
		fmt.Println("error create file", err.Error())
		os.Exit(1)
	}

	defer func() {
		docx.Close()
	}()

	defer func() {
		pdfOut.Close()
	}()

	err = libreoffice.DocxToPdf(docx, pdfOut)
	if err != nil {
		fmt.Println("error convert file", err.Error())
		os.Exit(1)
	}
}
