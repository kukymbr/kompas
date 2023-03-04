package domain

type SpcObject struct {
	ID                string `example:"374729415333.000000"`
	Text              string `example:"УПИГ-01.000-32  Рама"`
	Columns           []*SpcObjectColumn
	AdditionalColumns []*SpcObjectColumn
}

type SpcObjectColumn struct {
	Name     string `example:"Наименование"`
	TypeName string `example:"name"`
	Value    string `example:"Рама"`
}
