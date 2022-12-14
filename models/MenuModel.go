package models

type MenuModel struct {
	Mid    int
	Parent int
	Name   string
	Seq    int
	Format string
}

type MenuTree struct {
	MenuModel
	Child []MenuModel
}

func (m *MenuModel) TableName() string {
	return "xcms_menu"
}
