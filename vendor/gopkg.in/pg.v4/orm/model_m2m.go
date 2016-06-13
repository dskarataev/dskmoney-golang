package orm

import (
	"fmt"
	"reflect"
)

type m2mModel struct {
	*sliceTableModel
	baseTable *Table
	rel       *Relation

	buf       []byte
	dstValues map[string][]reflect.Value
	columns   map[string]string
}

var _ tableModel = (*m2mModel)(nil)

func newM2MModel(join *join) *m2mModel {
	baseTable := join.BaseModel.Table()
	joinModel := join.JoinModel.(*sliceTableModel)
	dstValues := dstValues(joinModel.Root(), joinModel.Path(), baseTable.PKs)
	return &m2mModel{
		sliceTableModel: joinModel,
		baseTable:       baseTable,
		rel:             join.Rel,

		dstValues: dstValues,
		columns:   make(map[string]string),
	}
}

func (m *m2mModel) NewModel() ColumnScanner {
	m.strct = reflect.New(m.table.Type).Elem()
	m.structTableModel.NewModel()
	return m
}

func (m *m2mModel) AddModel(_ ColumnScanner) error {
	m.buf = modelIdMap(m.buf[:0], m.columns, m.baseTable.ModelName+"_", m.baseTable.PKs)
	dstValues, ok := m.dstValues[string(m.buf)]
	if !ok {
		return fmt.Errorf("pg: can't find dst value for model id=%q", m.buf)
	}
	for _, v := range dstValues {
		v.Set(reflect.Append(v, m.strct))
	}
	return nil
}

func (m *m2mModel) ScanColumn(colIdx int, colName string, b []byte) error {
	ok, err := m.sliceTableModel.scanColumn(colIdx, colName, b)
	if ok {
		return err
	}

	m.columns[colName] = string(b)
	return nil
}
