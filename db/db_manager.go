/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Tue May 09 2023 4:40:23 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package db

import (
	"database/sql"
	"fmt"

	"github.com/eulbyvan/go-enigma-laundry/config"
	"github.com/eulbyvan/go-enigma-laundry/util"
)

type DbManager interface {
	SqlDb() *sql.DB
}
type dbManager struct {
	db *sql.DB
}

func (i *dbManager) SqlDb() *sql.DB {
	return i.db
}

func NewDbManager(config config.Config) DbManager {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Name)
	db, err := sql.Open("postgres", psqlInfo)
	err = db.Ping()
	util.CheckErr(err)
	fmt.Println("Successfully connected!")
	return &dbManager{
		db,
	}
}
