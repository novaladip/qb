package qb

import "testing"

type insertDataTable struct {
	expectedStmt string
	expectedArgs []interface{}
	table        string
	columns      []string
	values       []interface{}
}

func Test_insertData_build(t *testing.T) {
	tables := []insertDataTable{
		{
			table:        "user",
			columns:      []string{"first_name", "last_name"},
			values:       []interface{}{"john", "doe"},
			expectedStmt: "INSERT INTO user (first_name, last_name) VALUES($1, $2)",
			expectedArgs: []interface{}{"john", "doe"},
		},
	}

	for _, v := range tables {
		i := Insert(v.table).Columns(v.columns...).Values(v.values...).Build()

		if i.Stmt != v.expectedStmt {
			t.Errorf("Stmt was incorecct \ngot: %v\nwant: %v", i.Stmt, v.expectedStmt)
		}

		for index, arg := range i.Args {
			if arg != v.expectedArgs[index] {
				t.Errorf("Args was incorrect \ngot: %v\n want: %v", arg, v.expectedArgs[index])
			}
		}

	}
}
