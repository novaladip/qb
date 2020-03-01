package main

import (
	"testing"
)

func Test_selectData_Build(t *testing.T) {
	tables := []struct {
		selec        string
		from         string
		where        expWithArg
		andWhere     []expWithArg
		andBetween   []betweenExpWithArgs
		orderBy      string
		limit        int
		offset       int
		expectedStmt string
		expectedArgs []interface{}
	}{
		{
			selec: "username, email, age",
			from:  "user",
			where: expWithArg{"username =", "john"},
			andWhere: []expWithArg{
				{"email =", "john@email.com"},
			},
			andBetween: []betweenExpWithArgs{
				{"age", []interface{}{20, 25}},
			},
			orderBy:      "age DESC",
			limit:        10,
			offset:       1,
			expectedStmt: "SELECT username, email, age FROM user WHERE username = $1 AND email = $2 AND age BETWEEN $3 AND $4 ORDER BY age DESC LIMIT 10 OFFSET 1 ",
			expectedArgs: []interface{}{"john", "john@email.com", 20, 25},
		},
		{
			selec: "*",
			from:  "user",
			where: expWithArg{"gender =", "M"},
			andWhere: []expWithArg{
				{"age >=", 22},
			},
			andBetween: []betweenExpWithArgs{
				{"height", []interface{}{160, 170}},
			},
			orderBy:      "gender ASC",
			limit:        100,
			offset:       10,
			expectedStmt: "SELECT * FROM user WHERE gender = $1 AND age >= $2 AND height BETWEEN $3 AND $4 ORDER BY gender ASC LIMIT 100 OFFSET 10 ",
			expectedArgs: []interface{}{"M", 22, 160, 170},
		},
	}

	for _, v := range tables {
		sd := Select(v.selec).From(v.from).Where(v.where.exp, v.where.arg).OrderBy(v.orderBy).Limit(v.limit).Offset(v.offset)

		for _, aw := range v.andWhere {
			sd.AndWhere(aw.exp, aw.arg)
		}

		for _, ab := range v.andBetween {
			sd.AndBetween(ab.exp, ab.args...)
		}
		sd.Build()

		if sd.Stmt != v.expectedStmt {
			t.Errorf("Stmt was incorrect \ngot: %v\nwant: %v", sd.Stmt, v.expectedStmt)
		}

		for index, arg := range sd.Args {
			if arg != v.expectedArgs[index] {
				t.Errorf("Args was incorrect \ngot: %v\n want: %v", arg, v.expectedArgs[index])
			}
		}
	}
}
