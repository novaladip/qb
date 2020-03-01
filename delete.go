package qb

import "fmt"

type deleteData struct {
	Stmt     string
	Args     []interface{}
	table    string
	where    expWithArg
	andWhere []expWithArg
}

func Delete(table string) *deleteData {
	return &deleteData{table: table}
}

func (dd *deleteData) buildDelete() *deleteData {
	dd.Stmt = fmt.Sprintf("DELETE FROM %v ", dd.table)
	return dd
}

func (dd *deleteData) Where(exp string, arg interface{}) *deleteData {
	dd.where = expWithArg{exp, arg}
	return dd
}

func (dd *deleteData) buildWhere() *deleteData {
	dd.Args = append(dd.Args, dd.where.arg)
	dd.Stmt += fmt.Sprintf("WHERE %v $%v ", dd.where.exp, len(dd.Args))
	return dd
}

func (dd *deleteData) AndWhere(exp string, arg interface{}) *deleteData {
	dd.andWhere = append(dd.andWhere, expWithArg{exp, arg})
	return dd
}

func (dd *deleteData) buildAndWhere() *deleteData {
	for _, v := range dd.andWhere {
		dd.Args = append(dd.Args, v.arg)
		dd.Stmt += fmt.Sprintf("AND %v $%v ", v.exp, len(dd.Args))
	}

	return dd
}

func (dd *deleteData) Build() *deleteData {
	dd.buildDelete()

	if len(dd.where.exp) > 0 {
		dd.buildWhere()
	}

	if len(dd.andWhere) > 0 {
		dd.buildAndWhere()
	}

	return dd
}
