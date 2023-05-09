/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Tue May 09 2023 4:41:45 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package util

import "github.com/google/uuid"

func GenerateUuid() string {
	return uuid.New().String()
}
