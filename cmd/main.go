package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mmartinjoo/explainer/internal/analyzer"
	"github.com/mmartinjoo/explainer/internal/parser"
	"github.com/mmartinjoo/explainer/internal/runner"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/analytics")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	queries, err := parser.Parse("/Users/joomartin/code/explainer/queries.log")
	if err != nil {
		panic(err)
	}

	explains, err := runner.Run(db, queries)
	if err != nil {
		panic(err)
	}

	results, err := analyzer.Analyze(explains)
	if err != nil {
		panic(err)
	}

	for _, res := range results {
		fmt.Printf("%s\n", res.String())
	}
}
