package metainfo

// xmlDoc is a mapping for the kompas MetaInfo file
type xmlDoc struct {
	Descriptions struct {
		PropertyDescriptions struct {
			PropertyDescription []struct {
				ID           string `xml:"id,attr"`
				Name         string `xml:"name,attr"`
				TypeValue    string `xml:"typeValue,attr"`
				NatureId     string `xml:"natureId,attr"`
				UnitId       string `xml:"unitId,attr"`
				Comment      string `xml:"comment,attr"`
				ValueVariant []struct {
					ID    string `xml:"id,attr"`
					Value string `xml:"value,attr"`
				} `xml:"valueVariant"`
			} `xml:"propertyDescription"`
		} `xml:"propertyDescriptions"`
	} `xml:"descriptions"`
	Properties struct {
		Property []struct {
			ID             string `xml:"id,attr"`
			Value          string `xml:"value,attr"`
			Modified       string `xml:"modified,attr"`
			ValueVariantId string `xml:"valueVariantId,attr"`
		} `xml:"property"`
	} `xml:"properties"`
	SpcDescriptions struct {
		SpcDescription []struct {
			Style struct {
				FileName      string `xml:"fileName,attr"`
				ID            string `xml:"id,attr"`
				MassUnit      string `xml:"massUnit,attr"`
				MassCommaSize string `xml:"massCommaSize,attr"`
				Modified      string `xml:"modified,attr"`
				Sections      struct {
					Section []struct {
						Name          string `xml:"name,attr"`
						Number        string `xml:"number,attr"`
						NestingBlocks struct {
							Block []struct {
								Name     string `xml:"name,attr"`
								Number   string `xml:"number,attr"`
								Included string `xml:"included,attr"`
								Sections struct {
									Section []struct {
										Number   string `xml:"number,attr"`
										Included string `xml:"included,attr"`
									} `xml:"section"`
								} `xml:"sections"`
							} `xml:"block"`
						} `xml:"nestingBlocks"`
					} `xml:"section"`
				} `xml:"sections"`
				AdditionalBlocks struct {
					Block []struct {
						Name     string `xml:"name,attr"`
						Number   string `xml:"number,attr"`
						Included string `xml:"included,attr"`
						Sections struct {
							Section []struct {
								Number   string `xml:"number,attr"`
								Included string `xml:"included,attr"`
							} `xml:"section"`
						} `xml:"sections"`
					} `xml:"block"`
				} `xml:"additionalBlocks"`
				Columns struct {
					Column []struct {
						Name     string `xml:"name,attr"`
						TypeName string `xml:"typeName,attr"`
						Type     string `xml:"type,attr"`
						Number   string `xml:"number,attr"`
					} `xml:"column"`
				} `xml:"columns"`
				AdditionalColumns struct {
					Column []struct {
						Name     string `xml:"name,attr"`
						TypeName string `xml:"typeName,attr"`
						Type     string `xml:"type,attr"`
						Number   string `xml:"number,attr"`
					} `xml:"column"`
				} `xml:"additionalColumns"`
			} `xml:"style"`
			Documents struct {
				Document struct {
					FileName string `xml:"fileName,attr"`
				} `xml:"document"`
			} `xml:"documents"`
			SpcObjects struct {
				Object []*xmlDocSpcObject `xml:"object"`
			} `xml:"spcObjects"`
			SpcStruct struct {
				Section []struct {
					AttrText string                   `xml:"text,attr"`
					Object   []*xmlDocSpcStructObject `xml:"object"`
				} `xml:"section"`
			} `xml:"spcStruct"`
		} `xml:"spcDescription"`
	} `xml:"spcDescriptions"`
}

func (d *xmlDoc) isValid() bool {
	return len(d.SpcDescriptions.SpcDescription) > 0
}

func (d *xmlDoc) findSpcObjectByID(id string) *xmlDocSpcObject {
	if !d.isValid() {
		return nil
	}

	for _, obj := range d.SpcDescriptions.SpcDescription[0].SpcObjects.Object {
		if obj.ID == id {
			return obj
		}
	}

	return nil
}

type xmlDocSpcObject struct {
	ID       string `xml:"id,attr"`
	Modified string `xml:"modified,attr"`
	Section  struct {
		Number                string `xml:"number,attr"`
		SubSecNumber          string `xml:"subSecNumber,attr"`
		AdditionalBlockNumber string `xml:"additionalBlockNumber,attr"`
		AdditionalSecNumber   string `xml:"additionalSecNumber,attr"`
		NestingBlockNumber    string `xml:"nestingBlockNumber,attr"`
		NestingSecNumber      string `xml:"nestingSecNumber,attr"`
	} `xml:"section"`
	Columns struct {
		Column []struct {
			Name        string `xml:"name,attr"`
			TypeName    string `xml:"typeName,attr"`
			Type        string `xml:"type,attr"`
			Number      string `xml:"number,attr"`
			BlockNumber string `xml:"blockNumber,attr"`
			Value       string `xml:"value,attr"`
			Modified    string `xml:"modified,attr"`
		} `xml:"column"`
	} `xml:"columns"`
	AdditionalColumns struct {
		Column []struct {
			Name        string `xml:"name,attr"`
			TypeName    string `xml:"typeName,attr"`
			Type        string `xml:"type,attr"`
			Number      string `xml:"number,attr"`
			BlockNumber string `xml:"blockNumber,attr"`
			Value       string `xml:"value,attr"`
			Modified    string `xml:"modified,attr"`
		} `xml:"column"`
	} `xml:"additionalColumns"`
	Documents struct {
		Document struct {
			FileName string `xml:"fileName,attr"`
		} `xml:"document"`
	} `xml:"documents"`
}

type xmlDocSpcStructObject struct {
	ID       string `xml:"id,attr"`
	AttrText string `xml:"text,attr"`
}
