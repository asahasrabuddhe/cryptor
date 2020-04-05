package module

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func GetFileContent(path string) (content []byte, err error) {
	var (
		fp    *os.File
		fstat os.FileInfo
	)

	fp, err = os.Open(path)
	if err != nil {
		return
	}
	defer fp.Close()

	fstat, err = fp.Stat()
	if err != nil {
		return
	}

	if fstat.Size() == 0 {
		return content, errors.New("file is empty")
	}

	content, err = ioutil.ReadAll(fp)
	if err != nil {
		return content, errors.New("error is reading file :" + err.Error())
	}

	return
}

func WriteTo(w io.Writer, msg string) {
	fmt.Fprint(w, msg)
}

func WriteToExit(w io.Writer, msg string) {
	WriteTo(w, msg)
	os.Exit(1)
}
