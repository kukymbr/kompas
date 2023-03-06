package domain

// SpcStructSections is a parsed struct of the spcStruct part of the Kompas' xml
type SpcStructSections []*SpcStructSection

// SpcStructSection is a one section of the SpcStructSections
type SpcStructSection struct {
	Name    string `example:"Сборочные единицы"`
	Objects []*SpcObject
}
