package main

import (
	"time"
)

type DataTransactionORM struct {
	Amount               float64    `gorm:"column:amount;type:decimal;not null"`
	AmountBefore         string     `gorm:"column:amount_before;type:text"`
	AmountCurrency       string     `gorm:"column:amount_currency;type:text;not null"`
	BatchId              string     `gorm:"column:batch_id;type:text"`
	BatchIndex           uint32     `gorm:"column:batch_index;type:int4"`
	BillData             string     `gorm:"column:bill_data;type:jsonb;not null"`
	BillingId            string     `gorm:"column:billing_id;type:text;not null"`
	CompanyId            string     `gorm:"column:company_id;type:varchar(20);not null"`
	CreatedAt            *time.Time `gorm:"column:created_at"`
	CreatedById          uint64     `gorm:"column:created_by_id;type:int8;not null"`
	CreatedByName        string     `gorm:"column:created_by_name;type:text;not null"`
	CurrentSchedule      *time.Time `gorm:"column:current_schedule"`
	CustomerName         string     `gorm:"column:customer_name;type:text;not null"`
	DebitAccountCurrency string     `gorm:"column:debit_account_currency;type:text;not null"`
	DebitAccountName     string     `gorm:"column:debit_account_name;type:text;not null"`
	DebitAccountNumber   string     `gorm:"column:debit_account_number;type:text;not null"`
	Emails               string     `gorm:"column:emails;type:text;not null"`
	ErrorMessage         string     `gorm:"column:error_message;type:text"`
	Fee                  float64    `gorm:"column:fee;type:decimal;not null"`
	FeeCurrency          string     `gorm:"column:fee_currency;type:text;not null"`
	Id                   uint64     `gorm:"column:id;type:int8;primary_key;not null"`
	Request              string     `gorm:"column:request;type:jsonb;not null"`
	Response             string     `gorm:"column:response;type:jsonb;not null"`
	StatusId             uint64     `gorm:"column:status_id;type:int4;not null"`
	TaskData             string     `gorm:"column:task_data;type:jsonb;not null"`
	TaskId               uint64     `gorm:"column:task_id;type:int8;unique;not null"`
	TransactionId        string     `gorm:"column:transaction_id;type:text;not null"`
	TransactionSchedule  string     `gorm:"column:transaction_schedule;type:text;not null"`
	UpdatedAt            *time.Time `gorm:"column:updated_at"`
	UpdatedById          uint64     `gorm:"column:updated_by_id;type:int8"`
	UpdatedByName        string     `gorm:"column:updated_by_name;type:text"`
	WorkflowDoc          string     `gorm:"column:workflow_doc;type:jsonb;not null"`
}

func (DataTransactionORM) TableName() string {
	return "transactions"
}
