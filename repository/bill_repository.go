/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Tue May 09 2023 4:41:10 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package repository

import (
	"database/sql"
	"time"

	"github.com/eulbyvan/go-enigma-laundry/entity"
	"github.com/eulbyvan/go-enigma-laundry/util"
)

type BillRepository interface {
	Create(newBill *entity.BillRequest) error
}

type billRepository struct {
	db *sql.DB
}

func (b *billRepository) Create(newBill *entity.BillRequest) error {
	tx, err := b.db.Begin()
	util.CheckErr(err)
	billStmt := `INSERT into bill (id,bill_date,finish_date, employee_id,customer_id) VALUES($1,$2,$3,$4,$5)`
	billId := util.GenerateUuid()
	billDate := time.Now()
	finishDate := billDate.AddDate(0, 0, 7)

	_, err = tx.Exec(billStmt, billId, billDate, finishDate, newBill.EmployeeId, newBill.CustomerId)
	util.ValidateTransaction(err, tx)

	billDetailStmt := `INSERT into bill_detail (id,bill_id,product_id, product_price,qty) VALUES($1,$2,$3,$4,$5)`
	for _, prodReq := range newBill.Items {
		billDetailId := util.GenerateUuid()
		_, err = tx.Exec(billDetailStmt, billDetailId, billId, prodReq.ProductId, prodReq.Price, prodReq.Qty)
		util.ValidateTransaction(err, tx)
	}
	tx.Commit()
	return err
}

func NewBillRepository(db *sql.DB) BillRepository {
	billRepo := billRepository{
		db,
	}
	return &billRepo
}
