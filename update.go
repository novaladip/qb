package qb

import "fmt"

type updateData struct {
	Stmt     string
	Args     []interface{}
	table    string
	set      []expWithArg
	where    expWithArg
	andWhere []expWithArg
}

func Update(table string) *updateData {
	return &updateData{table: table}
}

func (ud *updateData) buildUpdate() *updateData {
	ud.Stmt += fmt.Sprintf("UPDATE %v ", ud.table)
	return ud
}

func (ud *updateData) Set(exp string, arg interface{}) *updateData {
	ud.set = append(ud.set, expWithArg{exp, arg})
	return ud
}

func (ud *updateData) buildSet() *updateData {
	ud.Stmt += "SET "
	for i, v := range ud.set {
		ud.Args = append(ud.Args, v.arg)
		if i == len(ud.set)-1 {
			ud.Stmt += fmt.Sprintf("%v = $%v ", v.exp, len(ud.Args))
		} else {
			ud.Stmt += fmt.Sprintf("%v = $%v, ", v.exp, len(ud.Args))
		}
	}

	return ud
}

func (ud *updateData) Where(exp string, arg interface{}) *updateData {
	ud.where = expWithArg{exp, arg}
	return ud
}

func (ud *updateData) buildWhere() *updateData {
	ud.Args = append(ud.Args, ud.where.arg)
	ud.Stmt += fmt.Sprintf("WHERE %v $%v ", ud.where.exp, len(ud.Args))

	return ud
}

func (ud *updateData) AndWhere(exp string, arg interface{}) *updateData {
	ud.andWhere = append(ud.andWhere, expWithArg{exp, arg})
	return ud
}

func (ud *updateData) buildAndWhere() *updateData {
	for _, v := range ud.andWhere {
		ud.Args = append(ud.Args, v.arg)
		ud.Stmt += fmt.Sprintf("AND %v $%v ", v.exp, len(ud.Args))
	}

	return ud
}

func (ud *updateData) Build() *updateData {
	ud.buildUpdate()
	ud.buildSet()

	if len(ud.where.exp) > 0 {
		ud.buildWhere()
	}

	if len(ud.andWhere) > 0 {
		ud.buildAndWhere()
	}

	return ud
}
