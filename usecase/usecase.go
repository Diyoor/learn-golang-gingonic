package usecase

import (
	"errors"

	"github.com/maxdev/go-gingonic/entity"
	"github.com/maxdev/go-gingonic/repository"
)

type TodoUsecase struct {
	repo repository.RepositoryI
}

func CreateTodoUsecase(repo repository.RepositoryI) UsecaseI {
	return &TodoUsecase{repo: repo}
}

func (uc *TodoUsecase) AddTodo(todo *entity.Todo) (int64, error) {

	id, err := uc.repo.AddTodo(todo)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *TodoUsecase) GetTodos() ([]entity.Todo, error) {
	data, err := uc.repo.GetTodos()
	return data, err
}
func (uc *TodoUsecase) UpdateTodo(id int64, todo map[string]interface{}) (entity.Todo, error) {
	_, err := uc.repo.GetByID(id)

	if err != nil {
		return entity.Todo{}, errors.New("CAN'T FIND ID")
	} else {
		data, err := uc.repo.UpdateTodo(id, todo)
		return data, err
	}

}

func (uc *TodoUsecase) DeleteTodo(id int64) (string, error) {

	_, err := uc.repo.GetByID(id)

	if err != nil {
		return "", errors.New("CAN'T FIND ID")
	} else {
		data, err := uc.repo.DeleteTodo(id)
		return data, err
	}

}
