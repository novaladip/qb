package qb

import "fmt"

type insertData struct {
	Stmt    string
	table   string
	columns []string
	Args    []interface{}
}

func Insert(table string) *insertData {
	return &insertData{
		table: table,
	}
}

func (i *insertData) buildInsert() *insertData {
	i.Stmt = fmt.Sprintf("INSERT INTO %v ", i.table)
	return i
}

func (i *insertData) Columns(columns ...string) *insertData {
	i.columns = append(i.columns, columns...)
	return i
}

func (i *insertData) buildColumns() *insertData {
	i.Stmt += "("

	for index, v := range i.columns {
		if index == len(i.columns)-1 {
			i.Stmt += fmt.Sprintf("%v) ", v)
		} else {
			i.Stmt += fmt.Sprintf("%v, ", v)
		}
	}
	return i
}

func (i *insertData) Values(values ...interface{}) *insertData {
	i.Args = append(i.Args, values...)
	return i
}

func (i *insertData) buildValues() *insertData {
	i.Stmt += "VALUES("

	for index := range i.Args {
		if index == len(i.Args)-1 {
			i.Stmt += fmt.Sprintf("$%v)", index+1)
		} else {
			i.Stmt += fmt.Sprintf("$%v, ", index+1)
		}
	}

	return i
}

func (i *insertData) Build() *insertData {
	i.buildInsert()
	i.buildColumns()
	i.buildValues()

	return i
}
