package main

import (
	"fmt"
	"strings"
)

type selectData struct {
	Stmt       string
	Args       []interface{}
	selec      string
	from       string
	where      expWithArg
	andWhere   []expWithArg
	andBetween []betweenExpWithArgs
	orderBy    string
	limit      int
	offset     int
}

func Select(exp string) *selectData {
	return &selectData{
		selec: exp,
	}
}

func (s *selectData) buildSelect() *selectData {
	s.Stmt = fmt.Sprintf("SELECT %v ", s.selec)
	return s
}

func (s *selectData) From(exp string) *selectData {
	s.from = exp
	return s
}

func (s *selectData) buildFrom() *selectData {
	s.Stmt += fmt.Sprintf("FROM %v ", s.from)
	return s
}

func (s *selectData) Where(exp string, arg interface{}) *selectData {
	s.where = expWithArg{exp, arg}
	return s
}

func (s *selectData) buildWhere() *selectData {
	s.Args = append(s.Args, s.where.arg)
	s.Stmt += fmt.Sprintf("WHERE %v $%v ", s.where.exp, len(s.Args))
	return s
}

func (s *selectData) AndWhere(exp string, arg interface{}) *selectData {
	s.andWhere = append(s.andWhere, expWithArg{exp, arg})
	return s
}

func (s *selectData) buildAndWhere() *selectData {
	for _, v := range s.andWhere {
		s.Args = append(s.Args, v.arg)
		s.Stmt += fmt.Sprintf("AND %v $%v ", v.exp, len(s.Args))
	}

	return s
}

func (s *selectData) AndBetween(exp string, args ...interface{}) *selectData {
	s.andBetween = append(s.andBetween, betweenExpWithArgs{exp, args})
	return s
}

func (s *selectData) buildAndBetween() *selectData {
	for _, v := range s.andBetween {
		s.Args = append(s.Args, v.args...)
		count := len(s.Args)
		s.Stmt += fmt.Sprintf("AND %v BETWEEN $%v AND $%v ", v.exp, count-1, count)
	}

	return s
}

func (s *selectData) OrderBy(exp string) *selectData {
	s.orderBy = exp
	return s
}

func (s *selectData) buildOrderBy() *selectData {
	s.Stmt += fmt.Sprintf("ORDER BY %v ", s.orderBy)
	return s
}

func (s *selectData) Limit(value int) *selectData {
	s.limit = value
	return s
}

func (s *selectData) buildLimit() *selectData {
	s.Stmt += fmt.Sprintf("LIMIT %v ", s.limit)
	return s
}

func (s *selectData) Offset(value int) *selectData {
	s.offset = value
	return s
}

func (s *selectData) buildOffset() *selectData {
	s.Stmt += fmt.Sprintf("OFFSET %v ", s.offset)
	return s
}

func (s *selectData) Build() *selectData {
	s.buildSelect()
	s.buildFrom()

	if len(s.where.exp) > 0 {
		s.buildWhere()
	}

	if len(s.andWhere) > 0 {
		s.buildAndWhere()
	}

	if len(s.andBetween) > 0 {
		s.buildAndBetween()
	}

	if strings.TrimSpace(s.orderBy) != "" {
		s.buildOrderBy()
	}

	if s.limit > 0 {
		s.buildLimit()
	}

	if s.offset > 0 {
		s.buildOffset()
	}

	return s
}
