﻿![WIP](https://www.repostatus.org/badges/latest/wip.svg) ![build status](https://api.travis-ci.com/novaladip/qb.svg?token=pi2LtwHe97ad5UFUJcjx&branch=master)
### qb (Query Builder)
*when I learn to make a rest-api with Go I face a case where I need a query builder to avoid DRY but because I'm absolutely new with Go and SQL, so I think it will be a good practice for me to create query builder by my self.*

###  currently only support for postgres

*current features*
 - [x] Select
      - [x] Where
      - [ ] WhereIn
      - [ ] WhereBetween
      - [ ] WhereNotBetween
      - [x] AndWhere
      - [ ] AndWhereIn
      - [x] AndBetween
      - [ ] AndNotBetween
      - [x] OrderBy
      - [x] Limit
      - [x] Offset
      - [ ] Join
      - [ ] Left Join
      - [ ] Right Join
      - [ ] Inner Join
 - [x] Update
 - [x] Delete
 - [x] Insert

```go
import "github.com/novaladip/qb"

gender := "M"
age := 20

// Don't forget to call the Build() method to generate Stmt & Args
// THE Build() METHOD SHOULD ONLY CALLED ONCE
sd := qb.Select("*").From("person")
sd.Where("age >=", age).AndWhere("gender =", gender).Build()

fmt.Println(sd.Stmt) // SELECT * FROM person WHERE age >= $1 AND gender = $2
fmt.Println(sd.Args) // [20, M]

rows, err := db.Query(sd.Stmt, sd.Args...)
```