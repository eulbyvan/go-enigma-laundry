/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Tue May 09 2023 4:40:34 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package delivery

import (
	"fmt"
	"os"
	"strings"

	"github.com/eulbyvan/go-enigma-laundry/config"
	"github.com/eulbyvan/go-enigma-laundry/db"
	"github.com/eulbyvan/go-enigma-laundry/entity"
	"github.com/eulbyvan/go-enigma-laundry/repository"
	"github.com/eulbyvan/go-enigma-laundry/usecase"
	"github.com/eulbyvan/go-enigma-laundry/util"
)

type Console struct {
	createBill      usecase.CreateBill
	uomRegistration usecase.UomRegistration
}

func (c *Console) mainMenuForm() {
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("Enigma Laundry")
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("1. Master UOM")
	fmt.Println("2. Master Produk")
	fmt.Println("3. Master Staf")
	fmt.Println("4. Master Pelanggan")
	fmt.Println("5. Transaksi Baru")
	fmt.Println("0. Keluar")
	fmt.Println("Pilih Menu (0-5): ")
}

func (c *Console) uomForm() entity.Uom {
	var uomId string
	var uomName string
	var saveUOMConfirmation string
	fmt.Print("UOM Id: ")
	fmt.Scanln(&uomId)
	fmt.Print("UOM Name: ")
	fmt.Scanln(&uomName)
	fmt.Printf("UOM Id: %s, %s akan disimpan (y/t) ?", uomId, uomName)
	fmt.Scanln(&saveUOMConfirmation)

	if saveUOMConfirmation == "y" {
		var uom entity.Uom
		uom.Id = uomId
		uom.Name = uomName
		return uom
	}
	return entity.Uom{}
}

func (c *Console) billItemForm(billDetail *[]entity.BillItemRequest) {
	for {
		var productId string
		var qty int
		var saveBillDetConfirmation string
		fmt.Println("Produk Id:")
		fmt.Scanln(&productId)
		fmt.Println("Jumlah:")
		fmt.Scanln(&qty)
		fmt.Print("simpan produk (y/t) ?")
		fmt.Scanln(&saveBillDetConfirmation)
		if saveBillDetConfirmation == "y" {
			*billDetail = append(*billDetail, entity.BillItemRequest{
				ProductId: productId,
				Qty:       qty,
			})
		}
		var finishBillDetConfirmation string
		fmt.Print("selesai (y/t) ?")
		fmt.Scanln(&finishBillDetConfirmation)
		if finishBillDetConfirmation == "y" {
			fmt.Println(billDetail)
			break
		}
	}
}
func (c *Console) billForm() entity.BillRequest {
	var employeeId string
	var customerId string
	var billDetail []entity.BillItemRequest
	var saveBillConfirmation string
	fmt.Print("Id Staff: ")
	fmt.Scanln(&employeeId)
	fmt.Print("Id Pelanggan: ")
	fmt.Scanln(&customerId)
	c.billItemForm(&billDetail)
	fmt.Print("Buat Struk (y/t) ?")
	fmt.Scanln(&saveBillConfirmation)

	if saveBillConfirmation == "y" {
		var billReq entity.BillRequest
		billReq.EmployeeId = employeeId
		billReq.CustomerId = customerId
		billReq.Items = billDetail
		return billReq
	}
	return entity.BillRequest{}
}

func (c *Console) Run() {
	c.mainMenuForm()
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			uom := c.uomForm()
			err := c.uomRegistration.Do(&uom)
			util.CheckErr(err)
			c.mainMenuForm()
			break
		case "2":
			//delivery.ListProductForm(repo)
			break
		case "3":
			//delivery.SearchProductForm(repo)
			break
		case "4":
			//delivery.SearchProductForm(repo)
			break
		case "5":
			billRequest := c.billForm()
			fmt.Println(billRequest)
			err := c.createBill.Do(&billRequest)
			util.CheckErr(err)
			c.mainMenuForm()
			break
		case "0":
			os.Exit(0)
		}
	}
}

func NewConsole() *Console {
	cfg := config.NewConfig()
	dbManager := db.NewDbManager(cfg)
	createBillRepo := repository.NewBillRepository(dbManager.SqlDb())
	uomRepo := repository.NewUomRepository(dbManager.SqlDb())
	return &Console{
		createBill:      usecase.NewCreateBill(createBillRepo),
		uomRegistration: usecase.NewUomRegistration(uomRepo),
	}
}
