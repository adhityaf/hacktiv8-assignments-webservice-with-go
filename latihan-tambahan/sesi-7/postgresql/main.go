package main

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

type Employee struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

func (emp *Employee) Print(){
	fmt.Println("ID :", emp.ID)
	fmt.Println("FullName :", emp.FullName)
	fmt.Println("Email :", emp.Email)
	fmt.Println("Age :", emp.Age)
	fmt.Println("Division :", emp.Division)
	fmt.Println()
}

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "pgadmin"
	DB_PASSWORD = "pgadmin"
	DB_NAME     = "db-go-psql"
)

func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	// manually create employee
	employees := []Employee{
		{Email: "joko@gmail.com", FullName: "joko", Age: 22, Division: "Tech"},
		{Email: "budi@gmail.com", FullName: "budi", Age: 23, Division: "Tech"},
		{Email: "santoso@gmail.com", FullName: "santoso", Age: 23, Division: "Tech"},
	}

	// loop through slice of Employee
	// and insert employee data
	for _, employee := range employees{
		err = createEmployee(db, &employee)
		if err != nil{
			fmt.Println("error :", err.Error())
			return
		}
	}

	emps, err := getAllEmployees(db)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	for _, employee := range *emps{
		employee.Print()
	}

	var updateRequest = Employee{
		FullName: "diki",
		Email: "diki@gmail.com",
		Age: 22,
		Division: "Tech",
	}

	// Update Employee By id 
	err = updateEmployeeById(db, 2, &updateRequest)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	// Get Employee By id 
	fmt.Println("==== Get Employee By id ====")
	employee, err := getEmployeeById(db, 2)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}
	employee.Print()

	// Delete Employee By id
	err = deleteEmployeeById(db, 5)
	if err != nil{
		fmt.Println("error :", err.Error())
		return
	}
}

func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dsn)
	if err != nil{
		panic(err)
	}

	err = db.Ping()
	if err != nil{
		return nil, err
	}

	// test connect to db purpose
	// defer db.Close()

	// connection pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(10 * time.Second)

	return db, nil
}


func getAllEmployees(db *sql.DB) (*[]Employee, error)  {
	query := `
		SELECT id, full_name, email, age, division
		FROM employees
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var employees []Employee

	rows, err := stmt.Query()
	if err != nil{
		return nil, err
	}

	for rows.Next(){
		var employee Employee
		err := rows.Scan(
			&employee.ID, 
			&employee.FullName,
			&employee.Email,
			&employee.Age,
			&employee.Division,
		)
		
		if err != nil{
			return nil, err
		}

		employees = append(employees, employee)
	}

	return &employees, nil
}

func createEmployee(db *sql.DB, request *Employee) error{
	query := `
		INSERT INTO employees(full_name, email, age, division)
		VALUES($1, $2, $3, $4)
	`

	// begin transaction
	tx, err := db.Begin()
	if err != nil{
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		request.FullName, 
		request.Email,
		request.Age,
		request.Division,
	)
	if err != nil {
		// rollback transaction
		err := tx.Rollback()
		if err!= nil{
			return err
		}
		return err
	}

	// commit transaction
	return tx.Commit()
}

func updateEmployeeById(db *sql.DB, id int, request *Employee) error{
	query := `
		UPDATE employees
		SET full_name=$1,
			email=$2,
			age=$3,
			division=$4
		WHERE id=$5
	`

	// begin transaction
	tx, err := db.Begin()
	if err != nil{
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		request.FullName,
		request.Email,
		request.Age,
		request.Division,
		id,
	)
	if err != nil{
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	// commit transaction
	return tx.Commit()
}

func deleteEmployeeById(db *sql.DB, id int) error{
	query := `
		DELETE FROM employees
		WHERE id=$1
	`

	// begin transaction
	tx, err := db.Begin()
	if err != nil{
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil{
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	// commit transaction
	return tx.Commit()
}


func getEmployeeById(db *sql.DB, id int) (*Employee, error)  {
	query := `
		SELECT id, full_name, email, age, division
		FROM employees
		WHERE id=$1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var employee Employee

	err = row.Scan(
		&employee.ID,
		&employee.FullName,
		&employee.Email,
		&employee.Age,
		&employee.Division,
	)

	if err != nil{
		return nil, err
	}

	return &employee, nil
}