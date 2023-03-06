package metainfo

import (
	"encoding/xml"
	"errors"
	"io"
	"strings"

	"github.com/kukymbr/kompasreader/domain"
)

// NewUnmarshaller creates new Unmarshaller instance for the reader
func NewUnmarshaller(reader io.Reader) *Unmarshaller {
	return &Unmarshaller{reader: reader}
}

type Unmarshaller struct {
	reader io.Reader
	doc    *xmlDoc
}

func (u *Unmarshaller) Unmarshall() (spc domain.SpcStruct, err error) {
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

	spc = make(domain.SpcStruct, len(u.doc.SpcDescriptions.SpcDescription[0].SpcStruct.Section))

	for sectIndex, xmlSect := range u.doc.SpcDescriptions.SpcDescription[0].SpcStruct.Section {
		sect := &domain.SpcStructSection{
			Name:    xmlSect.AttrText,
			Objects: make([]*domain.SpcObject, 0, len(xmlSect.Object)),
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

func (u *Unmarshaller) prepareXML() (data []byte, err error) {
	data, err = io.ReadAll(u.reader)
	if err != nil {
		return nil, err
	}

	str := string(data)
	str = strings.TrimSpace(str)

	header := `<?xml version="1.0" encoding="utf-16"?>`
	if strings.HasPrefix(str, header) {
		str = strings.TrimPrefix(str, header)
	}

	return []byte(str), nil
}

func (u *Unmarshaller) buildSpcObject(xmlSectObj *xmlDocSpcStructObject) *domain.SpcObject {
	xmlObj := u.doc.findSpcObjectByID(xmlSectObj.ID)
	if xmlObj == nil {
		return nil
	}

	obj := &domain.SpcObject{
		ID:                xmlObj.ID,
		Text:              xmlSectObj.AttrText,
		Columns:           make([]*domain.SpcObjectColumn, len(xmlObj.Columns.Column)),
		AdditionalColumns: make([]*domain.SpcObjectColumn, len(xmlObj.AdditionalColumns.Column)),
	}

	for i, xmlCol := range xmlObj.Columns.Column {
		obj.Columns[i] = &domain.SpcObjectColumn{
			Name:     xmlCol.Name,
			TypeName: xmlCol.TypeName,
			Value:    xmlCol.Value,
		}
	}

	for i, xmlCol := range xmlObj.AdditionalColumns.Column {
		obj.AdditionalColumns[i] = &domain.SpcObjectColumn{
			Name:     xmlCol.Name,
			TypeName: xmlCol.TypeName,
			Value:    xmlCol.Value,
		}
	}

	return obj
}
