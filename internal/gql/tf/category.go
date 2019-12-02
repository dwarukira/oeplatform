package tf

import (
	gql "oe/internal/gql/models"
	"oe/internal/models"
	"strconv"
)

func GQLInputCategoryToDBCategory(i *gql.CategoryInput) (o *models.Category, err error) {
	o = &models.Category{
		Name: i.Name,
	}

	return o, err
}

func DBCategoryTOGQLCategory(i *models.Category) (o *gql.Category, err error) {
	o = &gql.Category{
		ID:   strconv.Itoa(i.ID),
		Name: &i.Name,
	}

	return o, err
}
