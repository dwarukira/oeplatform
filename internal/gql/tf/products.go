package tf

import (
	"fmt"
	gql "oe/internal/gql/models"
	dbm "oe/internal/models"
)

func GQLInputProductToDBProduct(i *gql.ProductCreateInput, seller *dbm.Seller, update bool, ids ...string) (o *dbm.Product, err error) {
	variants := []dbm.ProductVariant{}
	for _, vrec := range i.Variants {
		if rec, err := GQLInputProductVariantToDBProductVariant(vrec, false); err != nil {
			fmt.Println(err)
		} else {
			variants = append(variants, *rec)
		}
	}

	o = &dbm.Product{
		Name:        *i.Name,
		Seller:      *seller,
		Description: *i.Description,
		IsPublished: *i.IsPublished,
		Variants:    variants,
	}
	// TODO update stuff

	return o, err
}

func GQLInputProductVariantToDBProductVariant(i *gql.ProductVariantInput, update bool, ids ...string) (o *dbm.ProductVariant, err error) {

	o = &dbm.ProductVariant{
		Name:              i.Name,
		Sku:               i.Sku,
		Taxable:           *i.Taxable,
		InventoryQuantity: *i.InventoryQuantity,
		TrackInventory:    *i.TrackInventory,
		Price:             *i.Price,
		CompareAtPrice:    *i.CompareAtPrice,
		Weight:            *i.Weight,
		WeightUnit:        *i.WeightUnit,
		RequiresShipping:  *i.RequiresShipping,
	}
	// TODO update stuff

	return o, err
}

func DBProductToGQLProduct(i *dbm.Product) (o *gql.Product, err error) {
	// user, err := DBUserToGQLUser(&i.User)
	// bank, _ := DBBankToGQLBank(i.Bank)
	seller, _ := DBSellerToGQLSeller(&i.Seller)
	variants := []*gql.ProductVariant{}
	for _, vrec := range i.Variants {
		if rec, err := DBProductVariantToGQLProductVariant(&vrec); err != nil {
			fmt.Println(err)
		} else {
			variants = append(variants, rec)
		}
	}

	// s, err := DBSellerToGQLSeller(&i.Seller)
	// if err != nil {
	// 	fmt.Println(s, err, "We have")
	// }
	o = &gql.Product{
		ID:          i.ID,
		Name:        i.Name,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   *i.UpdatedAt,
		Variants:    variants,
		Description: i.Description,
		Seller:      seller,
	}

	return o, err
}

func DBProductVariantToGQLProductVariant(i *dbm.ProductVariant) (o *gql.ProductVariant, err error) {
	// user, err := DBUserToGQLUser(&i.User)
	// bank, _ := DBBankToGQLBank(i.Bank)
	o = &gql.ProductVariant{
		ID:                i.ID,
		Name:              i.Name,
		Sku:               i.Sku,
		Taxable:           &i.Taxable,
		Barcode:           &i.Barcode,
		Price:             &i.Price,
		Weight:            &i.Weight,
		TrackInventory:    &i.TrackInventory,
		Grams:             &i.Grams,
		CompareAtPrice:    &i.CompareAtPrice,
		CreatedAt:         i.CreatedAt,
		UpdatedAt:         *i.UpdatedAt,
		InventoryQuantity: &i.InventoryQuantity,
	}

	return o, err
}
