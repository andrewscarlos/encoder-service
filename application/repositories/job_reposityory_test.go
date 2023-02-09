package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/infrastructure/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "/test"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{
		Db: db,
	}

	repo.Insert(video)

	job, err := domain.NewJob("outputpach", "Pending", video)
	require.Nil(t, err)

	repojob := repositories.JobRepository{
		Db: db,
	}

	repojob.Insert(job)

	j, err := repojob.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)

}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Complete", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepository{Db: db}
	repoJob.Insert(job)

	job.Status = "Complete"

	repoJob.Update(job)

	j, err := repoJob.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)
}
