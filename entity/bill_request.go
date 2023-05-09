/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Tue May 09 2023 4:40:46 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package entity

type BillRequest struct {
	EmployeeId string
	CustomerId string
	Items      []BillItemRequest
}
