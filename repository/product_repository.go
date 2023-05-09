/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Tue May 09 2023 4:41:22 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package repository

import (
	"database/sql"

	"github.com/eulbyvan/go-enigma-laundry/entity"
	"github.com/eulbyvan/go-enigma-laundry/util"
)

type ProductRepository interface {
	Create(newProduct *entity.Product) error
	Delete(id string) error
	FindOne(id string) (entity.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func (u *productRepository) Create(newProduct *entity.Product) error {
	stmt := `INSERT INTO product (id,name,price,uom_id) VALUES ($1,$2,$3,$4)`
	prodId := util.GenerateUuid()
	_, err := u.db.Exec(stmt, newProduct.Id, newProduct.Name, newProduct.Price, newProduct.Uom.Id)
	newProduct.Id = prodId
	return err
}

func (u *productRepository) Delete(id string) error {
	stmt := `UPDATE product SET is_delete = true where id=$1`
	_, err := u.db.Exec(stmt, id)
	return err
}

func (u *productRepository) FindOne(id string) (entity.Product, error) {
	stmt := `SELECT p.id,p.name,p.price,p.uom_id,u.name 
	FROM product p JOIN uom u ON p.uom_id = u.id AND p.id=$1`
	row := u.db.QueryRow(stmt, id)
	var product entity.Product
	switch err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Uom.Id, &product.Uom.Name); err {
	case sql.ErrNoRows:
		return entity.Product{}, err
	case nil:
		return product, nil
	default:
		panic(err)
	}
}

func NewProductRepository(db *sql.DB) ProductRepository {
	prodRepo := productRepository{
		db,
	}
	return &prodRepo
}
