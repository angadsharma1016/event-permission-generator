package model

import "github.com/neo4j/neo4j-go-driver/neo4j"

var (
	session neo4j.Session
	result  *neo4j.Result
)

func SetDB(s neo4j.Session, r *neo4j.Result) {
	session = s
	result = r
}