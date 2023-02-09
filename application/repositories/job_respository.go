package repositories

import (
	"encoder/domain"
	"fmt"
	"github.com/jinzhu/gorm"
	
)

type IJobRepository interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
}

type JobRepository struct {
	Db *gorm.DB
}

func (repo JobRepository) Insert(job *domain.Job) (*domain.Job, error) {

	err := repo.Db.Create(job).Error
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (repo JobRepository) Find(id string) (*domain.Job, error) {
	var job domain.Job
	repo.Db.First(&job, "id = ?", id)

	if job.ID == "" {
		return nil, fmt.Errorf("job not found")
	}

	return &job, nil
}

func (repo JobRepository) Update(job *domain.Job) (*domain.Job, error) {
	
	err := repo.Db.Update(&job).Error
	if err != nil {
		return nil, err
	}

	return job, nil
}
