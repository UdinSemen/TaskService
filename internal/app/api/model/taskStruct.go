package model

type Task struct {
	Uri                 string      `json:"uri"`
	Id                  string      `json:"id"`
	Type                string      `json:"type"`
	CollectionFrequency int         `json:"collection_frequency"`
	Description         string      `json:"description"`
	Ti                  interface{} `json:"ti"`
}
