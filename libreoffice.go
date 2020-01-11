package libreoffice

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	LibreOfficeBaseCommand = "libreoffice"
	HeadlessOption         = "--headless"
	ConvertToOption        = "--convert-to"
	OutDirOption           = "--outdir"
	PDFType                = "pdf"
)

// DocxToPdf func will convert Doc or Docx to PDF
func DocxToPdf(in io.Reader, out io.Writer) error {

	tempDir := os.TempDir()
	inBuffer := make([]byte, 16384)

	tmpFileIn, err := ioutil.TempFile("", "golo")
	if err != nil {
		return err
	}

	// cleanup
	defer func() {
		os.Remove(tmpFileIn.Name())
		os.Remove(fmt.Sprintf("%s.pdf", tmpFileIn.Name()))
	}()

	for {
		line, err := in.Read(inBuffer)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		_, err = tmpFileIn.Write(inBuffer[:line])
		if err != nil {
			return err
		}
	}

	cmd := exec.Command(LibreOfficeBaseCommand, HeadlessOption, ConvertToOption, PDFType, tmpFileIn.Name(), OutDirOption, tempDir)
	err = cmd.Run()
	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut
	fmt.Println(cmdOut.String())
	if err != nil {
		return err
	}

	inData, err := ioutil.ReadFile(fmt.Sprintf("%s.pdf", tmpFileIn.Name()))
	if err != nil {
		return err
	}

	_, err = out.Write(inData)
	if err != nil {
		return err
	}

	err = tmpFileIn.Close()
	if err != nil {
		return err
	}

	return nil
}
