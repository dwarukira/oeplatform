package resolvers

import (
	"context"
	"oe/internal/gql/models"
	"oe/internal/gql/tf"
	dbm "oe/internal/models"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, input models.CategoryInput) (*models.Category, error) {

	return categoryCreateUpdate(r, input, false)
}

func (r *queryResolver) Categories(ctx context.Context, id *string) (*models.Categories, error) {
	return categoriesList(r, id)
}

func (r *queryResolver) SubCategories(ctx context.Context, categoryId *string) (*models.Categories, error) {
	return categoriesSubList(r, categoryId)
}

func categoriesList(r *queryResolver, id *string) (*models.Categories, error) {
	whereID := "id = ?"

	record := &models.Categories{}
	dbRecords := []*dbm.Category{}

	db := r.ORM.DB.New()

	db = db.Set("gorm:auto_preload", true)
	if id != nil {
		db = db.Where(whereID, *id)
	}

	db = db.Find(&dbRecords).Count(&record.Count)

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBCategoryTOGQLCategory(dbRec); err != nil {

		} else {
			record.List = append(record.List, rec)
		}
	}

	return record, db.Error
}

func categoriesSubList(r *queryResolver, id *string) (*models.Categories, error) {
	record := &models.Categories{}
	var err error
	var dbRecords []*dbm.Category
	if id == nil || *id == "" {
		dbRecords, err = r.ORM.FindRootCategories()
	} else {
		dbRecords, err = r.ORM.FindSubCategories(*id)
	}

	if err != nil {
		return record, err
	}

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBCategoryTOGQLCategory(dbRec); err != nil {

		} else {
			record.List = append(record.List, rec)
		}
	}

	return record, nil
}

func categoryCreateUpdate(r *mutationResolver, input models.CategoryInput, update bool) (*models.Category, error) {

	db := r.ORM.DB.New().Begin()
	parent, err := r.ORM.GetCategoryParent(*input.Parent)

	if err != nil {
		// TODO do something about the error
	}

	dbo, err := tf.GQLInputCategoryToDBCategory(&input)

	if err != nil {
		return nil, err
	}

	dbo.Parent = parent
	db = db.Create(dbo).First(dbo)

	db = db.Commit()

	gql, err := tf.DBCategoryTOGQLCategory(dbo)

	return gql, nil
}
