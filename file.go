package xmindgo

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/xiaobing94/xmindgo/models"
)

const (
	// 默认的meta信息文件名
	defaultMetaFilename = "metadata.json"

	// 默认的xml内容文件名
	defaultXmlContentFilename = "content.xml"

	// 默认的manifest文件名
	defaultManifestFilename = "manifest.json"

	// 默认的json内容文件名
	defaultJsonContentFilename = "content.json"
)

type File struct {
	Path     string
	Workbook *models.Workbook
	Manifest *models.Manifest
	MetaData *models.MetaData
	files    map[string][]byte
}

func NewFile() *File {
	files := make(map[string][]byte)
	file := &File{
		Workbook: &models.Workbook{},
		Manifest: models.NewManifest(),
		MetaData: &models.MetaData{},
		files:    files,
	}
	return file
}

func (f *File) Save() error {
	if f.Path == "" {
		return fmt.Errorf("no path defined for file, consider File.WriteTo or File.Write")
	}
	return f.SaveAs(f.Path)
}

func (f *File) SaveAs(path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	return f.Write(file)
}

func (f *File) Write(w io.Writer) error {
	_, err := f.WriteTo(w)
	return err
}

func (f *File) WriteTo(w io.Writer) (int64, error) {
	buf, err := f.WriteToBuffer()
	if err != nil {
		return 0, err
	}
	return buf.WriteTo(w)
}

func (f *File) WriteToBuffer() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	if err := f.generateFiles(); err != nil {
		return nil, err
	}
	for path, content := range f.files {
		fi, err := zw.Create(path)
		if err != nil {
			return buf, err
		}
		_, err = fi.Write(content)
		if err != nil {
			return buf, err
		}
	}
	return buf, zw.Close()
}
