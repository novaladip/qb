package qb

import (
	"testing"
)

type table struct {
	table        string
	where        expWithArg
	andWhere     []expWithArg
	expectedStmt string
	expectedArgs []interface{}
}

func Test_deleteData_Build(t *testing.T) {
	tables := []table{
		{
			table: "user",
			where: expWithArg{"id =", 1},
			andWhere: []expWithArg{
				{"email =", "john@email.com"},
			},
			expectedStmt: "DELETE FROM user WHERE id = $1 AND email = $2 ",
			expectedArgs: []interface{}{1, "john@email.com"},
		},
		{
			table: "user",
			where: expWithArg{"age <=", 10},
			andWhere: []expWithArg{
				{"gender =", "M"},
			},
			expectedStmt: "DELETE FROM user WHERE age <= $1 AND gender = $2 ",
			expectedArgs: []interface{}{10, "M"},
		},
	}

	for _, v := range tables {
		dd := Delete(v.table).Where(v.where.exp, v.where.arg)

		for _, aw := range v.andWhere {
			dd.AndWhere(aw.exp, aw.arg)
		}

		dd.Build()

		if dd.Stmt != v.expectedStmt {
			t.Errorf("Stmt was incorecct \ngot: %v\nwant: %v", dd.Stmt, v.expectedStmt)
		}

		for index, arg := range dd.Args {
			if arg != v.expectedArgs[index] {
				t.Errorf("Args was incorrect \ngot: %v\n want: %v", arg, v.expectedArgs[index])
			}
		}

	}
}
