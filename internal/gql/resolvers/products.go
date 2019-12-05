package resolvers

import (
	"context"
	"fmt"
	"io/ioutil"
	"oe/internal/gql/models"
	"oe/internal/gql/tf"
	"oe/internal/handlers/middleware"
	"oe/pkg/bucket"

	dbm "oe/internal/models"

	"github.com/99designs/gqlgen/graphql"
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

	// return productVariantCreateUpadate(input)

	return nil, nil
}

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*models.Image, error) {
	_, err := ioutil.ReadAll(file.File)

	if err != nil {
		return nil, err
	}

	// object, err = bucket.UploadFile(content)

	return nil, err
}

func (r *mutationResolver) AddProductPhoto(ctx context.Context, files []*graphql.Upload, productVariantID string) ([]*models.Image, error) {
	productVariant, _ := r.ORM.GetProductVariant(productVariantID)

	db := r.ORM.DB.New().Begin()
	var images []*dbm.Image

	for i := range files {
		content, err := ioutil.ReadAll(files[i].File)
		if err != nil {
			return []*models.Image{}, err
		}
		_, url, _ := bucket.UploadFile(content)
		images = append(images, &dbm.Image{
			Name:           "Image nname",
			Source:         url,
			ProductVariant: *productVariant,
		})
	}

	db.Model(&productVariant).Association("Images").Append(images)
	db.Commit()
	imagesresponse := []*models.Image{}

	for _, rec := range images {

		if rec, err := tf.DBImageToGQLImage(rec); err != nil {

		} else {
			imagesresponse = append(imagesresponse, rec)
		}
	}
	return imagesresponse, nil
}

func (r *mutationResolver) MultipleUploadWithPayload(ctx context.Context, req []*models.UploadFile) ([]*models.Image, error) {
	if len(req) == 0 {
		return nil, errors.New("empty list")
	}

	productVariant, _ := r.ORM.GetProductVariant(req[0].ProductVariantID)
	db := r.ORM.DB.New().Begin()

	var images []*dbm.Image
	for i := range req {
		content, err := ioutil.ReadAll(req[i].File.File)
		if err != nil {
			return []*models.Image{}, err
		}
		_, url, _ := bucket.UploadFile(content)
		images = append(images, &dbm.Image{
			Name:           "Image nname",
			Source:         url,
			ProductVariant: *productVariant,
		})
	}
	db.Model(&productVariant).Association("Images").Append(images)
	db.Commit()

	imagesresponse := []*models.Image{}

	for _, rec := range images {

		if rec, err := tf.DBImageToGQLImage(rec); err != nil {

		} else {
			imagesresponse = append(imagesresponse, rec)
		}
	}
	return imagesresponse, nil
}

func (r *queryResolver) Products(ctx context.Context, id *string, filter *models.FilterProduct) (*models.Products, error) {
	return productsList(r, id, filter)
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
	db = db.Preload("Variants.Images")
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

func productsList(r *queryResolver, id *string, filter *models.FilterProduct) (*models.Products, error) {
	whereID := "id = ?"
	record := &models.Products{}
	dbRecords := []*dbm.Product{}
	db := r.ORM.DB.New()
	db = db.Preload("Seller.User")
	db = db.Preload("Seller.Bank")
	db = db.Preload("Variants")
	db = db.Preload("Variants.Images")

	if filter != nil {
		db = db.Scopes(dbm.SellerScope(filter.Seller))
	}
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
