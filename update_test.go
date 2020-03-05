package qb

import "testing"

type updateDataTable struct {
	expectedStmt string
	expectedArgs []interface{}
	table        string
	set          []expWithArg
	where        expWithArg
	andWhere     []expWithArg
}

func Test_updateData_build(t *testing.T) {
	tables := []updateDataTable{
		{
			table: "user",
			set: []expWithArg{
				{"first_name", "jane"},
				{"last_name", "doe"},
			},
			where: expWithArg{"id =", 1},
			andWhere: []expWithArg{
				{"email =", "john@email.com"},
			},
			expectedStmt: "UPDATE user SET first_name = $1, last_name = $2 WHERE id = $3 AND email = $4 ",
			expectedArgs: []interface{}{"jane", "doe", 1, "john@email.com"},
		},
	}

	for _, v := range tables {
		ud := Update(v.table).Where(v.where.exp, v.where.arg)

		for _, s := range v.set {
			ud.Set(s.exp, s.arg)
		}

		for _, aw := range v.andWhere {
			ud.AndWhere(aw.exp, aw.arg)
		}

		ud.Build()

		if ud.Stmt != v.expectedStmt {
			t.Errorf("Stmt was incorecct \ngot: %v\nwant: %v", ud.Stmt, v.expectedStmt)
		}

		for index, arg := range ud.Args {
			if arg != v.expectedArgs[index] {
				t.Errorf("Args was incorrect \ngot: %v\n want: %v", arg, v.expectedArgs[index])
			}
		}
	}
}
