/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Tue May 09 2023 4:41:26 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package repository

import (
	"database/sql"

	"github.com/eulbyvan/go-enigma-laundry/entity"
	"github.com/eulbyvan/go-enigma-laundry/util"
)

type UomRepository interface {
	Create(newUom *entity.Uom) error
	Delete(id string) error
	FindOne(id string) (entity.Uom, error)
}

type uomRepository struct {
	db *sql.DB
}

func (u *uomRepository) Create(newUom *entity.Uom) error {
	stmt := `INSERT INTO uom (id,name) VALUES ($1,$2)`
	uomId := util.GenerateUuid()
	_, err := u.db.Exec(stmt, uomId, newUom.Name)
	newUom.Id = uomId
	return err
}

func (u *uomRepository) Delete(id string) error {
	stmt := `UPDATE uom SET is_delete = true where id=$1`
	_, err := u.db.Exec(stmt, id)
	return err
}

func (u *uomRepository) FindOne(id string) (entity.Uom, error) {
	stmt := `SELECT id,name FROM uom where id=$1`
	row := u.db.QueryRow(stmt, id)
	var uom entity.Uom
	switch err := row.Scan(&uom.Id, &uom.Name); err {
	case sql.ErrNoRows:
		return entity.Uom{}, err
	case nil:
		return uom, nil
	default:
		panic(err)
	}
}

func NewUomRepository(db *sql.DB) UomRepository {
	uomRepo := uomRepository{
		db,
	}
	return &uomRepo
}
