package tf

import (
	gql "oe/internal/gql/models"
	"oe/internal/models"
	"strconv"
)

func GQLInputCategoryToDBCategory(i *gql.CategoryInput) (o *models.Category, err error) {
	o = &models.Category{
		Name:        i.Name,
		Description: i.Description,
	}

	return o, err
}

func DBCategoryTOGQLCategory(i *models.Category) (o *gql.Category, err error) {
	if i.Parent == nil {
		o = &gql.Category{
			ID:          strconv.Itoa(i.ID),
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
		ID:          strconv.Itoa(i.ID),
		Name:        &i.Name,
		Description: i.Description,
		Parent:      category,
		CreatedAt:   i.CreatedAt,
	}

	return o, err
}
