/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Tue May 09 2023 4:41:32 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package usecase

import (
	"github.com/eulbyvan/go-enigma-laundry/entity"
	"github.com/eulbyvan/go-enigma-laundry/repository"
)

type CreateBill interface {
	Do(billRequest *entity.BillRequest) error
}

type createBill struct {
	repo repository.BillRepository
}

func (uc *createBill) Do(billRequest *entity.BillRequest) error {
	return nil
}
func NewCreateBill(repository repository.BillRepository) CreateBill {
	return &createBill{
		repo: repository,
	}
}
