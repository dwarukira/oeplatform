package tf

import (
	"fmt"
	gql "oe/internal/gql/models"
	dbm "oe/internal/models"
)

// GQLInputSellerToDBSeller transforms [Seller] gql input to db model
func GQLInputSellerToDBSeller(i *gql.SellerInput, user *dbm.User, update bool, ids ...string) (o *dbm.Seller, err error) {
	bank, err := GQLInputBankToDBBank(i.Bank, false)
	o = &dbm.Seller{
		Name:        i.Name,
		User:        *user,
		DisplayName: i.DisplayName,
		Bank:        bank,
		Phone:       i.Phone,
	}
	// TODO update stuff

	return o, err
}

// GQLInputBankToDBBank transforms [Bank] gql input to db model
func GQLInputBankToDBBank(i *gql.BankInput, update bool, ids ...string) (o *dbm.Bank, err error) {
	o = &dbm.Bank{
		Name:              &i.Name,
		Location:          &i.Location,
		AccountHolderName: &i.HolderName,
		BankAccount:       &i.AccountNumber,
	}
	// TODO update stuff

	return o, err
}

func DBBankToGQLBank(i *dbm.Bank) (o *gql.Bank, err error) {
	o = &gql.Bank{
		Name:          *i.Name,
		ID:            i.ID,
		AccountNumber: i.AccountHolderName,
		CreatedAt:     i.CreatedAt,
		UpdatedAt:     i.UpdatedAt,
	}

	return o, err
}

func DBSellerToGQLSeller(i *dbm.Seller) (o *gql.Seller, err error) {
	user, err := DBUserToGQLUser(&i.User)
	fmt.Println(i.Bank)
	bank, _ := DBBankToGQLBank(i.Bank)
	o = &gql.Seller{
		ID:          i.ID,
		Name:        i.Name,
		Phone:       &i.Phone,
		DisplayName: &i.DisplayName,
		User:        user,
		Bank:        bank,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
	}

	return o, err
}
