package sql

type index_type int

const (
	PRIMARY index_type = iota
	UNIQUE
	INDEX
	FULLTEXT
)

type index_ope int

const (
	DROP index_type = iota
	ADD
	MODIFY
)
