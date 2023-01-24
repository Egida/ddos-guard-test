package usecase

import "github.com/Shteyd/ddos-guard-test/internal/entity"

type UsersUseCase struct {
	repo UserRepo
}

func NewUsersUC(repo UserRepo) *UsersUseCase {
	return &UsersUseCase{repo: repo}
}

func (uc *UsersUseCase) Metric() (entity.Metric, error) {
	metric, err := uc.repo.GetMetric()
	if err != nil {
		return metric, err
	}

	return metric, nil
}

func (uc *UsersUseCase) GetUserID(username string) (int, error) {
	id, err := uc.repo.GetUserID(username)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *UsersUseCase) Store(username string) error {
	return uc.repo.Store(username)
}
