package tf

import (
	"fmt"
	gql "oe/internal/gql/models"
	"oe/internal/models"
)

func GQLInputCategoryToDBCategory(i *gql.CategoryInput) (o *models.Category, err error) {
	o = &models.Category{
		Name:        i.Name,
		Description: i.Description,
	}

	return o, err
}

func DBCategoryTOGQLCategory(i *models.Category) (o *gql.Category, err error) {
	fmt.Println(i.Sub, "----------------------> is this loaded")
	if i.Parent == nil {
		o = &gql.Category{
			ID:          i.ID,
			Name:        &i.Name,
			Description: i.Description,
			CreatedAt:   i.CreatedAt,
			Parent:      nil,
		}
		return o, err

	}

	category, err := DBCategoryTOGQLCategory(i.Parent)

	if err != nil {
		// TODO do something orr not I dont want ot think  abot it now
		return nil, err
	}

	o = &gql.Category{
		ID:          i.ID,
		Name:        &i.Name,
		Description: i.Description,
		Parent:      category,
		CreatedAt:   i.CreatedAt,
	}

	return o, err
}
