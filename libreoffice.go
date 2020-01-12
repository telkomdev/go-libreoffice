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
	HTMLType               = "html"
)

func command(tmpFileIn, tempDir, outputType string) error {
	cmd := exec.Command(LibreOfficeBaseCommand, HeadlessOption, ConvertToOption, outputType, tmpFileIn, OutDirOption, tempDir)

	var (
		cmdOut, cmdErr bytes.Buffer
	)
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr

	fmt.Println(fmt.Sprintf("stdout : %s", cmdOut.String()))
	fmt.Println(fmt.Sprintf("stderr : %s", cmdErr.String()))

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func execute(in io.Reader, out io.Writer, outputType string) error {
	tempDir := os.TempDir()
	inBuffer := make([]byte, 16384)

	tmpFileIn, err := ioutil.TempFile("", "golo")
	if err != nil {
		return err
	}

	// cleanup
	defer func() {
		os.Remove(tmpFileIn.Name())
		os.Remove(fmt.Sprintf("%s.%s", tmpFileIn.Name(), outputType))
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

	err = command(tmpFileIn.Name(), tempDir, outputType)
	if err != nil {
		return err
	}

	inData, err := ioutil.ReadFile(fmt.Sprintf("%s.%s", tmpFileIn.Name(), outputType))
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

// ToPdf func will convert any type to PDF
func ToPdf(in io.Reader, out io.Writer) error {
	return execute(in, out, PDFType)
}

// ToHTML func will convert any type to HTML
func ToHTML(in io.Reader, out io.Writer) error {
	return execute(in, out, HTMLType)
}
