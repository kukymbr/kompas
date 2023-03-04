package domain

// SpcStruct is a parsed struct of the spcStruct part of the Kompas' xml
type SpcStruct []*SpcStructSection

// SpcStructSection is a one section of the SpcStruct
type SpcStructSection struct {
	Name    string `example:"Сборочные единицы"`
	Objects []*SpcObject
}
