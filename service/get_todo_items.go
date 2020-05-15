package service

import (
	"swagger_test/models"
)

func (s *service) FetchTodoItems()[]*models.Item{
	items := make([]*models.Item,0)


	for _,v :=range List{

		val := models.Item(v)//create a copy as we can't get the address of a map element
		items = append(items, &val)

	}

 return items
}
