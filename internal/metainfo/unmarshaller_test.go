package metainfo_test

import (
	"strings"
	"testing"

	"github.com/kukymbr/kompasreader/internal/metainfo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshaller_Unmarshall_WhenValidXML_ExpectNoError(t *testing.T) {
	reader := strings.NewReader(givenValidXML())

	unm := metainfo.NewUnmarshaller(reader)
	spc, err := unm.Unmarshall()

	require.NoError(t, err)
	require.NotNil(t, spc)
	require.Len(t, spc, 2)

	assert.Equal(t, "Сборочные единицы", spc[0].Name)
	assert.Equal(t, "Материалы", spc[1].Name)

	require.Len(t, spc[0].Objects, 3)
	assert.Equal(t, "374729415333.000000", spc[0].Objects[0].ID)
	assert.Equal(t, "УПИГ-01.000-32  Рама", spc[0].Objects[0].Text)

	require.Len(t, spc[0].Objects[0].Columns, 4)
	assert.Equal(t, "Позиция", spc[0].Objects[0].Columns[0].Name)
	assert.Equal(t, "pos", spc[0].Objects[0].Columns[0].TypeName)
	assert.Equal(t, "1", spc[0].Objects[0].Columns[0].Value)

	require.Len(t, spc[0].Objects[1].AdditionalColumns, 1)
	assert.Equal(t, "Масса", spc[0].Objects[1].AdditionalColumns[0].Name)
	assert.Equal(t, "massa", spc[0].Objects[1].AdditionalColumns[0].TypeName)
	assert.Equal(t, "540,5361", spc[0].Objects[1].AdditionalColumns[0].Value)
}

func givenValidXML() string {
	return `<?xml version="1.0" encoding="utf-16"?>
<document modified="0">
	<spcDescriptions>
		<spcDescription>
			<spcObjects>
				<object id="374729415333.000000" modified="0">
					<section number="15" subSecNumber="0" additionalBlockNumber="0" additionalSecNumber="0" nestingBlockNumber="0" nestingSecNumber="0"/>
					<columns>
						<column name="Позиция" typeName="pos" type="3" number="1" blockNumber="0" value="1" modified="0"/>
						<column name="Обозначение" typeName="mark" type="4" number="1" blockNumber="0" value="УПИГ-01.000-32" modified="0"/>
						<column name="Наименование" typeName="name" type="5" number="1" blockNumber="0" value="Рама" modified="0"/>
						<column name="Количество" typeName="count" type="6" number="1" blockNumber="0" value="1" modified="0"/>
					</columns>
					<additionalColumns>
						<column name="Масса" typeName="massa" type="8" number="1" blockNumber="0" value="116,524" modified="0"/>
					</additionalColumns>
				</object>
				<object id="156194343357.000000" modified="0">
					<section number="15" subSecNumber="0" additionalBlockNumber="0" additionalSecNumber="0" nestingBlockNumber="0" nestingSecNumber="0"/>
					<columns>
						<column name="Позиция" typeName="pos" type="3" number="1" blockNumber="0" value="2" modified="0"/>
						<column name="Обозначение" typeName="mark" type="4" number="1" blockNumber="0" value="УПИГ-07.000-32" modified="0"/>
						<column name="Наименование" typeName="name" type="5" number="1" blockNumber="0" value="Газопровод" modified="0"/>
						<column name="Количество" typeName="count" type="6" number="1" blockNumber="0" value="1" modified="0"/>
					</columns>
					<additionalColumns>
						<column name="Масса" typeName="massa" type="8" number="1" blockNumber="0" value="540,5361" modified="0"/>
					</additionalColumns>
				</object>
				<object id="374729415533.000000" modified="0">
					<section number="15" subSecNumber="0" additionalBlockNumber="0" additionalSecNumber="0" nestingBlockNumber="0" nestingSecNumber="0"/>
					<columns>
						<column name="Позиция" typeName="pos" type="3" number="1" blockNumber="0" value="3" modified="0"/>
						<column name="Обозначение" typeName="mark" type="4" number="1" blockNumber="0" value="УПИГ-09.000-32" modified="0"/>
						<column name="Наименование" typeName="name" type="5" number="1" blockNumber="0" value="Электрооборудование" modified="0"/>
						<column name="Количество" typeName="count" type="6" number="1" blockNumber="0" value="1" modified="0"/>
					</columns>
					<additionalColumns>
						<column name="Масса" typeName="massa" type="8" number="1" blockNumber="0" value="780,226" modified="0"/>
					</additionalColumns>
				</object>
				<object id="180313760979.000000" modified="0">
					<section number="35" subSecNumber="0" additionalBlockNumber="0" additionalSecNumber="0" nestingBlockNumber="0" nestingSecNumber="0"/>
					<columns>
						<column name="Позиция" typeName="pos" type="3" number="1" blockNumber="0" value="33" modified="0"/>
						<column name="Наименование" typeName="name" type="5" number="1" blockNumber="0" value="Проволока нержавеющая @/2 мм 08Х18Н10" modified="0"/>
						<column name="Примечание" typeName="note" type="7" number="1" blockNumber="0" value="25 м" modified="0"/>
					</columns>
					<additionalColumns>
						<column name="Обозначение материала" typeName="user" type="10" number="3" blockNumber="0" value="Проволока нержавеющая @/2 мм 08Х18Н10" modified="0"/>
					</additionalColumns>
				</object>
			</spcObjects>
			<spcStruct>
				<section text="Сборочные единицы">
					<object id="374729415333.000000" text="УПИГ-01.000-32  Рама"/>
					<object id="156194343357.000000" text="УПИГ-07.000-32  Газопровод"/>
					<object id="374729415533.000000" text="УПИГ-09.000-32  Электрооборудование"/>
				</section>
				<section text="Материалы">
					<object id="180313760979.000000" text="Проволока нержавеющая  2 мм 08Х18Н10"/>
				</section>
			</spcStruct>
		</spcDescription>
	</spcDescriptions>
</document>`
}
