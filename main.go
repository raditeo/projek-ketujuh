package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

var (
	db  *sql.DB
	err error
)

func main() {
	db, err = sql.Open("mysql", "xx:xx@tcp(127.0.0.1:3306)/db-go-sql")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("successfully connected to database")

	// CreateEmployee()
	// GetEmployee()
	// UpdateEmployee()
	DeleteEmployee()
}

func CreateEmployee() {
	// var employee = Employee{}

	// sqlStatement := `
	// INSERT INTO employees (full_name, email, age, division)
	// VALUES ($1, $2, $3, $4);
	// SELECT LAST_INSERT_ID();
	// `

	// err = db.QueryRow(sqlStatement, "Raditeo", "raditeo@wlaisongo.ac.id", 23, "PTIPD").
	// 	Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("new employee data: %+v\n", employee)

	var employee = Employee{}

	query := "INSERT INTO employees (full_name, email, age, division) VALUES (?, ?, ?, ?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		fmt.Printf(err.Error())
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, "Raditeo", "raditeo@gmail.com", 23, "PTIPD")
	if err != nil {
		fmt.Printf(err.Error())
	}

	employeeID, err := res.LastInsertId()
	if err != nil {
		fmt.Printf(err.Error())
	}

	err = db.QueryRow("SELECT id, full_name, email, age, division FROM employees where id = ?", employeeID).Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	fmt.Printf("new employee data: %+v\n", employee)
}

func GetEmployee() {
	var results = []Employee{}

	sqlStatement := "SELECT id, full_name, email, age, division FROM employees"

	rows, err := db.Query(sqlStatement)

	if err != nil {
		fmt.Printf(err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}
		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			fmt.Printf(err.Error())
		}

		results = append(results, employee)
	}

	fmt.Println("employee data:", results)
}

func UpdateEmployee() {
	sqlStatement := "UPDATE employees SET full_name = ?, email = ?, division = ?, age = ? where id = ?"

	res, err := db.Exec(sqlStatement, "raditeo", "raditeo@walisongo.ac.id", "ptipd", 24, 1)

	if err != nil {
		fmt.Printf(err.Error())
	}

	count, err := res.RowsAffected()

	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println("updated rows:", count)
}

func DeleteEmployee() {
	sqlStatement := "DELETE FROM employees WHERE id = ?"

	res, err := db.Exec(sqlStatement, 1)
	if err != nil {
		fmt.Printf(err.Error())
	}

	count, err := res.RowsAffected()

	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println("deleted rows:", count)
}
