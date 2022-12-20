package models

var (
	_ Model = (*Project)(nil)
	_ FilesManager = (*Project)(nil)
)

type Project struct {
	BaseModel

	Name string `db:"name" json:"name"`
}

// TableName retuns the Collection model SQL table name.
func (m *Project) TableName() string {
	return "_projects"
}

func (m *Project) BaseFilesPath() string {
	return m.Id
}