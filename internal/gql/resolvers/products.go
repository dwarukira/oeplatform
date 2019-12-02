package resolvers

import (
	"context"
	"fmt"
	"oe/internal/gql/models"
	"oe/internal/gql/tf"
	"oe/internal/handlers/middleware"

	dbm "oe/internal/models"

	"github.com/pkg/errors"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input models.ProductCreateInput) (*models.Product, error) {
	user := middleware.ForContext(ctx)

	if user == nil {
		return &models.Product{}, fmt.Errorf("Access denied")
	}

	return productCreateUpdate(r, input, user, false)
}

func (r *mutationResolver) CreateProductVariant(ctx context.Context, input models.CreateProductVariantInput) (*models.ProductVariant, error) {
	// TODO create a product

	return productVariantCreateUpadate(input)
}

func (r *queryResolver) Products(ctx context.Context, id *string) (*models.Products, error) {
	return productsList(r, id)
}

func (r *queryResolver) SellerProducts(ctx context.Context, id *string) (*models.Products, error) {
	user := middleware.ForContext(ctx)

	if user == nil {
		return &models.Products{}, fmt.Errorf("Access denied")
	}

	seller, err := r.ORM.FindSellerUser(user)
	if err != nil {
		return &models.Products{}, fmt.Errorf("Access denied")
	}

	return sellerProductList(r, seller, id)
}

func sellerProductList(r *queryResolver, seller *dbm.Seller, id *string) (*models.Products, error) {
	whereID := "id = ?"
	whereSeller := "seller_id = ?"

	record := &models.Products{}
	dbRecords := []*dbm.Product{}

	db := r.ORM.DB.New()

	db = db.Preload("Seller.User")
	db = db.Preload("Seller.Bank")
	db = db.Preload("Variants")

	if id != nil {
		db = db.Where(whereID, *id)
	}

	db = db.Where(whereSeller, seller.ID)

	db = db.Find(&dbRecords).Count(&record.Count)

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBProductToGQLProduct(dbRec); err != nil {

		} else {
			record.List = append(record.List, rec)
		}
	}

	return record, db.Error

}

func productsList(r *queryResolver, id *string) (*models.Products, error) {

	whereID := "id = ?"
	record := &models.Products{}
	dbRecords := []*dbm.Product{}

	db := r.ORM.DB.New()

	db = db.Preload("Seller.User")
	db = db.Preload("Seller.Bank")

	if id != nil {
		db = db.Where(whereID, *id)
	}

	db = db.Find(&dbRecords).Count(&record.Count)

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBProductToGQLProduct(dbRec); err != nil {

		} else {
			record.List = append(record.List, rec)
		}
	}

	return record, db.Error
}

func productVariantCreateUpadate(r *mutationResolver, input models.CreateProductVariantInput, update bool, ids ...string) (*models.ProductVariant, error) {

	return nil, nil
}

func productCreateUpdate(r *mutationResolver, input models.ProductCreateInput, user *dbm.User, update bool, ids ...string) (*models.Product, error) {

	seller, _ := r.ORM.FindSellerUser(user)

	if len(seller.ID) < 0 {
		return nil, errors.New("Access denied. Not a seller")
	}

	dbo, err := tf.GQLInputProductToDBProduct(&input, seller, false)

	if err != nil {
		return nil, err
	}

	db := r.ORM.DB.New().Begin()

	db = db.Preload("Seller.User")
	db = db.Preload("Seller.Bank")

	db = db.Create(dbo).First(dbo)

	gql, err := tf.DBProductToGQLProduct(dbo)

	if err != nil {
		db.RollbackUnlessCommitted()
		return nil, err
	}

	db = db.Commit()

	return gql, nil
}
