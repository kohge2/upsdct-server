package main

import (
	"context"
	"log"

	"github.com/kohge2/upsdct-server/adapter"
	"github.com/kohge2/upsdct-server/config"
	"github.com/kohge2/upsdct-server/database"
	"github.com/kohge2/upsdct-server/domain/models"
	"github.com/kohge2/upsdct-server/utils"
)

func main() {
	db, err := database.GetDB(config.Env.DbDsn)
	if err != nil {
		log.Fatal("DB ERROR", err)
	}
	transaction, err := database.NewTransaction(config.Env.DbDsn)
	if err != nil {
		log.Fatal("Txn ERROR", err)
	}

	partnerCompanyAdapter := adapter.NewPartnerCompanyRepository(db)

	ctx := context.Background()
	if err := transaction.RunTxn(ctx, func(ctx context.Context) error {
		partnerCompnay := &models.PartnerCompany{
			ID:   "pc001",
			Name: "partnerCompany",
		}

		if err := partnerCompanyAdapter.CreatePartnerCompany(ctx, partnerCompnay); err != nil {
			log.Fatal("CreatePartnerCompany ERROR", err)
		}
		accountNumber, err := utils.Encrypt("accountNumber", config.Env.EncryptKey)
		if err != nil {
			log.Fatal("Encrypt ERROR", err)
		}
		accountHolderName, err := utils.Encrypt("accountHolderName", config.Env.EncryptKey)
		if err != nil {
			log.Fatal("Encrypt ERROR", err)
		}

		if err := partnerCompanyAdapter.CreatePartnerCompanyBankAccount(ctx, &models.PartnerCompanyBankAccount{
			ID:                "pcba001",
			PartnerCompanyID:  "pc001",
			BankName:          "bankName",
			BranchName:        "branchName",
			AccountType:       "accountType",
			AccountNumber:     accountNumber,
			AccountHolderName: accountHolderName,
		}); err != nil {
			log.Fatal("CreatePartnerCompanyBankAccount ERROR", err)
		}

		return nil
	}); err != nil {
		log.Fatal("RunTxn ERROR", err)
		return
	}
	log.Println("Success")
}
