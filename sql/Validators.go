package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type RowValidator struct {
	Name    string
	Website string
	Address string
}

func InsertValidators(name string, website string, address string) {
	DBMutex.Lock()
	stmt, err := DBConnection.Prepare(
		`INSERT INTO Validators
			(
				name,
				website,
				address
			) VALUES(?,?,?);`,
	)
	DBMutex.Unlock()
	if err != nil {
		log.Printf("SQL Statement Prepare Error: %s\n", err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		name,
		website,
		address,
	)
	if err != nil {
		return
	}
}

func SelectValidatorByAddress(address string) RowValidator {
	DBMutex.Lock()
	ret, err := DBConnection.Query(
		`SELECT 
			*
		FROM Validators
		WHERE address = ?;`,
		address,
	)
	DBMutex.Unlock()
	if err != nil {
		log.Printf("SQL Statement Prepare Error: %s\n", err.Error())
		return RowValidator{}
	}

	row_validator := RowValidator{}
	for ret.Next() {
		ret.Scan(
			&(row_validator.Name),
			&(row_validator.Website),
			&(row_validator.Address),
		)
	}

	return row_validator
}
