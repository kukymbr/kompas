package metainfo

import (
	"encoding/xml"
	"errors"
	"github.com/kukymbr/kompas"
	"io"
	"strings"
)

// NewUnmarshaler creates new Unmarshaler instance for the reader
func NewUnmarshaler(reader io.Reader) *Unmarshaler {
	return &Unmarshaler{reader: reader}
}

type Unmarshaler struct {
	reader io.Reader
	doc    *xmlDoc
}

func (u *Unmarshaler) Unmarshal() (spc kompas.SpcStructSections, err error) {
	u.doc = &xmlDoc{}

	data, err := u.prepareXML()
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(data, &u.doc)
	if err != nil {
		return nil, err
	}

	if !u.doc.isValid() {
		return nil, errors.New("mata info doc is not valid")
	}

	spc = make(kompas.SpcStructSections, len(u.doc.SpcDescriptions.SpcDescription[0].SpcStruct.Section))

	for sectIndex, xmlSect := range u.doc.SpcDescriptions.SpcDescription[0].SpcStruct.Section {
		sect := &kompas.SpcStructSection{
			Name:    xmlSect.AttrText,
			Objects: make([]*kompas.SpcObject, 0, len(xmlSect.Object)),
		}

		for _, xmlSectObj := range xmlSect.Object {
			obj := u.buildSpcObject(xmlSectObj)
			if obj == nil {
				continue
			}

			sect.Objects = append(sect.Objects, obj)
		}

		spc[sectIndex] = sect
	}

	return spc, nil
}

func (u *Unmarshaler) prepareXML() (data []byte, err error) {
	data, err = io.ReadAll(u.reader)
	if err != nil {
		return nil, err
	}

	str := string(data)
	str = strings.TrimSpace(str)

	header := `<?xml version="1.0" encoding="utf-16"?>`
	str = strings.TrimPrefix(str, header)

	return []byte(str), nil
}

func (u *Unmarshaler) buildSpcObject(xmlSectObj *xmlDocSpcStructObject) *kompas.SpcObject {
	xmlObj := u.doc.findSpcObjectByID(xmlSectObj.ID)
	if xmlObj == nil {
		return nil
	}

	obj := &kompas.SpcObject{
		ID:                xmlObj.ID,
		Text:              xmlSectObj.AttrText,
		Columns:           make([]*kompas.SpcObjectColumn, len(xmlObj.Columns.Column)),
		AdditionalColumns: make([]*kompas.SpcObjectColumn, len(xmlObj.AdditionalColumns.Column)),
	}

	for i, xmlCol := range xmlObj.Columns.Column {
		obj.Columns[i] = &kompas.SpcObjectColumn{
			Name:     xmlCol.Name,
			TypeName: xmlCol.TypeName,
			Value:    xmlCol.Value,
		}
	}

	for i, xmlCol := range xmlObj.AdditionalColumns.Column {
		obj.AdditionalColumns[i] = &kompas.SpcObjectColumn{
			Name:     xmlCol.Name,
			TypeName: xmlCol.TypeName,
			Value:    xmlCol.Value,
		}
	}

	return obj
}
