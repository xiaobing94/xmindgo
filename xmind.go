package xmindgo

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/xiaobing94/xmindgo/models"
)

const (
	resourcePrefix = "resources/"

	xapPrefix = "xap:"
)

func OpenFile(path string) (*File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	f, err := OpenReader(file)
	if err != nil {
		return nil, err
	}
	f.Path = path
	return f, nil
}

func OpenReader(r io.Reader) (*File, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	if err != nil {
		return nil, err
	}
	files, err := ReadZipReader(zr)
	if err != nil {
		return nil, err
	}
	content := files[defaultJsonContentFilename]
	workbook, err := ParseWorkbook(content)
	if err != nil {
		return nil, err
	}
	manifestContent := files[defaultManifestFilename]
	manifest, err := ParseManifest(manifestContent)
	if err != nil {
		return nil, err
	}
	metaDataContent := files[defaultMetaFilename]
	metaData, err := ParseMetaData(metaDataContent)
	if err != nil {
		return nil, err
	}
	file := &File{
		Workbook: workbook,
		Manifest: manifest,
		MetaData: metaData,
		files:    files,
	}
	return file, nil
}

func ParseWorkbook(content []byte) (*models.Workbook, error) {
	var workbook models.Workbook
	var sheets []*models.Sheet
	if err := json.Unmarshal(content, &sheets); err != nil {
		return nil, err
	}
	workbook.SetSheets(sheets)
	return &workbook, nil
}

func ParseManifest(content []byte) (*models.Manifest, error) {
	var manifest models.Manifest
	err := json.Unmarshal(content, &manifest)
	return &manifest, err
}

func ParseMetaData(content []byte) (*models.MetaData, error) {
	var metaData models.MetaData
	err := json.Unmarshal(content, &metaData)
	return &metaData, err
}

func ReadZipReader(r *zip.Reader) (map[string][]byte, error) {
	fileList := make(map[string][]byte)
	for _, v := range r.File {
		data, err := readFile(v)
		if err != nil {
			return nil, err
		}
		fileList[v.Name] = data
	}
	return fileList, nil
}

func (f *File) GetWorkbook() *models.Workbook {
	return f.Workbook
}

func (f *File) DeleteFile(filename string) {
	f.removeFile(filename)
}

func (f *File) addMetadata() error {
	data, err := json.Marshal(f.MetaData)
	if err != nil {
		return err
	}
	f.addFile(defaultMetaFilename, data)
	return nil
}

func (f *File) addDefaultXml() {
	f.addFile(defaultXmlContentFilename, []byte(xmlContent))
}

func (f *File) addManifest() error {
	data, err := json.Marshal(f.Manifest)
	if err != nil {
		return err
	}
	f.files[defaultManifestFilename] = data
	return nil
}

func (f *File) addWorkbook() error {
	sheets := f.Workbook.GetSheets()
	data, err := json.Marshal(sheets)
	if err != nil {
		return err
	}
	f.addFile(defaultJsonContentFilename, data)
	return nil
}

func (f *File) addFile(path string, content []byte) {
	f.files[path] = content
	f.Manifest.AddFile(path)
}

func (f *File) CreateTopicImage(content []byte, width int, height int) *models.TopicImage {
	image := &models.TopicImage{
		Image: models.Image{},
	}
	image.Width = width
	image.Height = height
	return image
}

func (f *File) AddResource(content []byte, filename string) (string, error) {
	checkSum, err := CheckSum(content)
	if err != nil {
		return "", err
	}
	filenames := strings.Split(filename, ".")
	// 无后缀文件直接替换为hash
	if len(filenames) < 2 {
		return resourcePrefix + checkSum, nil
	}
	suffix := filenames[len(filenames)-1]
	resourcePath := fmt.Sprintf("%s%s.%s", resourcePrefix, checkSum, suffix)
	f.addFile(resourcePath, content)
	resourcePath = xapPrefix + resourcePath
	return resourcePath, nil
}

func (f *File) removeFile(filename string) {
	delete(f.files, filename)
	f.Manifest.RemoveFile(filename)
}

func (f *File) GetResource(src string) ([]byte, error) {
	filename := strings.TrimPrefix(src, xapPrefix)
	content, ok := f.files[filename]
	if !ok {
		return nil, fmt.Errorf("%s not found", filename)
	}
	return content, nil
}

func (f *File) CreateImage(content []byte, filename string) (*models.TopicImage, error) {
	xapResource, err := f.AddResource(content, filename)
	if err != nil {
		return nil, err
	}
	topicImage := &models.TopicImage{
		Image: models.Image{
			Src: xapResource,
		},
	}
	return topicImage, nil
}

func (f *File) generateFiles() error {
	f.addDefaultXml()
	if err := f.addWorkbook(); err != nil {
		return err
	}
	if err := f.addMetadata(); err != nil {
		return err
	}
	return f.addManifest()
}
