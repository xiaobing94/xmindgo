package models

type MetaCreator struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type MetaData struct {
	Creator       *MetaCreator `json:"creator,omitempty"`
	ActiveSheetId string       `json:"active_sheet_id,omitempty"`
}

func NewMetaData() *MetaData {
	return &MetaData{
		Creator:       &MetaCreator{},
		ActiveSheetId: "",
	}
}