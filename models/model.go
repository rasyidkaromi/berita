package models

import "time"

type MysqlDB struct {
	ID      int       `json:"id" `
	Author  string    `json:"author" `
	Body    string    `json:"body" "`
	Created time.Time `json:"created"`
}

type BeritaStruct struct {
	ID      int       `json:"id"`
	Author  string    `json:"author"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

func (m *BeritaStruct) TName() string {
	return "berita"
}

type ElasticNews struct {
	ID      int       `json:"id"`
	Created time.Time `json:"created"`
}

func (m *ElasticNews) TName() string {
	return "berita"
}

type GetSetting struct {
	Filter map[string]interface{} `json:"filter"`
	Offset int                    `json:"offset"`
	Limit  int                    `json:"limit"`
	Order  map[string]bool        `json:"order"`
}

func (m *GetSetting) String() string {
	return "payload"
}
