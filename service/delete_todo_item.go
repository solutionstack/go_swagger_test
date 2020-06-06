package service

import (
	"errors"
)

//function bar
func (s *service) DeleteTodoItem(id int64) error {
	//check if item exist
	if _, ok := List[id]; !ok {
		return errors.New("404")
	}
	//delete
	delete(List, id)
	return nil
}
