/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Tue May 09 2023 4:41:36 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package usecase

import (
	"github.com/eulbyvan/go-enigma-laundry/entity"
	"github.com/eulbyvan/go-enigma-laundry/repository"
)

type UomRegistration interface {
	Do(newUom *entity.Uom) error
}

type uomRegistration struct {
	repo repository.UomRepository
}

func (uc *uomRegistration) Do(newUom *entity.Uom) error {
	return nil
}
func NewUomRegistration(repository repository.UomRepository) UomRegistration {
	return &uomRegistration{
		repo: repository,
	}
}
