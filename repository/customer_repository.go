/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Tue May 09 2023 4:41:14 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package repository

import (
	"database/sql"

	"github.com/eulbyvan/go-enigma-laundry/entity"
	"github.com/eulbyvan/go-enigma-laundry/util"
)

type CustomerRepository interface {
	Create(newCustomer *entity.Customer) error
	Delete(id string) error
	FindOne(id string) (entity.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func (u *customerRepository) Create(newCustomer *entity.Customer) error {
	stmt := `INSERT INTO customer (id,name,phone_number) VALUES ($1,$2)`
	custId := util.GenerateUuid()
	_, err := u.db.Exec(stmt, custId, newCustomer.Name)
	newCustomer.Id = custId
	return err
}

func (u *customerRepository) Delete(id string) error {
	stmt := `UPDATE Customer SET is_delete = true where id=$1`
	_, err := u.db.Exec(stmt, id)
	return err
}

func (u *customerRepository) FindOne(id string) (entity.Customer, error) {
	stmt := `SELECT id,name,phone_number FROM customer where id=$1`
	row := u.db.QueryRow(stmt, id)
	var Customer entity.Customer
	switch err := row.Scan(&Customer.Id, &Customer.Name); err {
	case sql.ErrNoRows:
		return entity.Customer{}, err
	case nil:
		return Customer, nil
	default:
		panic(err)
	}
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	custRepo := customerRepository{
		db,
	}
	return &custRepo
}
