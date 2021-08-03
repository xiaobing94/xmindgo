package models

type Manifest struct {
	FileEntries map[string]interface{} `json:"file-entries"`
}

func NewManifest() *Manifest {
	manifest := &Manifest{
		FileEntries: make(map[string]interface{}, 0),
	}
	return manifest
}

func (m *Manifest) RemoveFile(filename string) {
	delete(m.FileEntries, filename)
}

func (m *Manifest) AddFile(filename string) {
	var empty struct{}
	if m.FileEntries == nil {
		m.FileEntries = make(map[string]interface{})
	}
	m.FileEntries[filename] = &empty
}
