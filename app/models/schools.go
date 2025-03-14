package models

type School struct {
	Name     string `db:"name"`
	Npsn     string `db:"npsn"`
	Address  string `db:"address"`
	Type     string `db:"type"`
	IsActive bool   `db:"is_active"`
}
type SchoolSetting struct {
	SchoolID uint   `db:"name"`
	Logo     string `db:"logo"`
	Favicon  string `db:"favicon"`
}
