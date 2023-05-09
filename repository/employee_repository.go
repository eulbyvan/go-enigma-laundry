/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Tue May 09 2023 4:41:18 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package repository

import (
	"database/sql"

	"github.com/eulbyvan/go-enigma-laundry/entity"
	"github.com/eulbyvan/go-enigma-laundry/util"
)

type EmployeeRepository interface {
	Create(newEmployee *entity.Employee) error
	Delete(id string) error
	FindOne(id string) (entity.Employee, error)
}

type employeeRepository struct {
	db *sql.DB
}

func (u *employeeRepository) Create(newEmployee *entity.Employee) error {
	stmt := `INSERT INTO employee (id,name) VALUES ($1,$2)`
	empId := util.GenerateUuid()
	_, err := u.db.Exec(stmt, empId, newEmployee.Name)
	newEmployee.Id = empId
	return err
}

func (u *employeeRepository) Delete(id string) error {
	stmt := `UPDATE employee SET is_delete = true where id=$1`
	_, err := u.db.Exec(stmt, id)
	return err
}

func (u *employeeRepository) FindOne(id string) (entity.Employee, error) {
	stmt := `SELECT id,name FROM employee where id=$1`
	row := u.db.QueryRow(stmt, id)
	var employee entity.Employee
	switch err := row.Scan(&employee.Id, &employee.Name); err {
	case sql.ErrNoRows:
		return entity.Employee{}, err
	case nil:
		return employee, nil
	default:
		panic(err)
	}
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	empRepo := employeeRepository{
		db,
	}
	return &empRepo
}
