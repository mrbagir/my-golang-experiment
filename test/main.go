package main

import (
	"fmt"
	"time"
)

func main() {
	timestamp := time.Now().UnixNano()
	fmt.Println(timestamp)
	fmt.Println(timestamp % 1000000)
}

// type User struct {
// 	ID   uint
// 	Name string
// }

// db, err := gorm.Open(postgres.Open("host=172.18.132.22 user=addons password=@DDONS!2022 dbname=addons_cellular_postpaid port=5000 sslmode=disable TimeZone=Asia/Jakarta"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	dataORM := DataTransactionORM{
// 		Amount:               1000.50,
// 		AmountBefore:         "500.00",
// 		AmountCurrency:       "USD",
// 		BatchId:              "batch12",
// 		BatchIndex:           1,
// 		BillData:             "{}",
// 		BillingId:            "billing123",
// 		CompanyId:            "1234",
// 		CreatedById:          1,
// 		CreatedByName:        "John Doe",
// 		CustomerName:         "Jane Smith",
// 		DebitAccountCurrency: "USD",
// 		DebitAccountName:     "John Doe",
// 		DebitAccountNumber:   "123456789",
// 		Emails:               "jane.smith@example.com",
// 		ErrorMessage:         "",
// 		Fee:                  5.00,
// 		FeeCurrency:          "USD",
// 		Id:                   1,
// 		Request:              "{}",
// 		Response:             "{}",
// 		StatusId:             1,
// 		TaskData:             "{}",
// 		TaskId:               1,
// 		TransactionId:        "trx126",
// 		UpdatedById:          1,
// 		UpdatedByName:        "John Doe",
// 		WorkflowDoc:          "{}",
// 	}

// 	db.Debug().Clauses(clause.OnConflict{
// 		UpdateAll: true,
// 	}).Create(&dataORM)
