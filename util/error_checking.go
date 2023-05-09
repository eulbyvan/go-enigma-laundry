/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Tue May 09 2023 4:41:40 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package util

import (
	"database/sql"
	"fmt"
)

func ValidateTransaction(err error, tx *sql.Tx) {
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
