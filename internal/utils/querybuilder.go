package utils

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	clauses  []string
	args     []any
	position int
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		clauses:  []string{},
		args:     []any{},
		position: 1,
	}
}

func (qb *QueryBuilder) AddField(column string, value any) {
	qb.clauses = append(qb.clauses, fmt.Sprintf("%s = $%d", column, qb.position))
	qb.args = append(qb.args, value)
	qb.position++
}

func (qb *QueryBuilder) HasUpdates() bool {
	return len(qb.clauses) > 0
}

func (qb *QueryBuilder) BuildSetClause() string {
	return strings.Join(qb.clauses, ", ")
}

func (qb *QueryBuilder) GetArgs() []any {
	return qb.args
}

func (qb *QueryBuilder) GetNextPosition() int {
	return qb.position
}
