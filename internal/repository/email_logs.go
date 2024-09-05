package repository

import (
	"context"
	"errors"
	"github.com/couchbase/gocb/v2"
	"github.com/sirupsen/logrus"
	"go-couchbase/internal/entity"
	"go-couchbase/utils"
	"time"
)

type emailLogRepository struct {
	couchbase           *gocb.Bucket
	emailLogCollection  *gocb.Collection
	templatesCollection *gocb.Collection
}

func NewEmailLogRepository(couchbase *gocb.Bucket) entity.EmailLogRepository {
	emailLogRepo := emailLogRepository{
		couchbase: couchbase,
	}

	// define collection
	emailLogRepo.emailLogCollection = emailLogRepo.couchbase.Scope("email").Collection("email")
	emailLogRepo.templatesCollection = emailLogRepo.couchbase.Scope("email").Collection("templates")

	return &emailLogRepo
}

// CreateEmailLog is method to create new data email
func (e *emailLogRepository) CreateEmailLog(ctx context.Context, input *entity.EmailLog) (*entity.EmailLog, error) {
	id := utils.GenerateID()

	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"id": id,
	})

	input.ID = id
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()
	input.DeletedAt = nil

	_, err := e.templatesCollection.Upsert(utils.ExpectString(id), input, &gocb.UpsertOptions{})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return input, nil
}

// GetEmailLogByID is method to get data by id
func (e *emailLogRepository) GetEmailLogByID(ctx context.Context, id uint64) (*entity.EmailLog, error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"id": id,
	})

	var err error
	result, err := e.templatesCollection.Get(utils.ExpectString(id), nil)
	switch err {
	case nil:
		var emailLog entity.EmailLog
		if err = result.Content(&emailLog); err != nil {
			logger.Error(err)
			return nil, err
		}

		cas := result.Cas()
		logrus.Infof("cas : %v", cas)
		return &emailLog, nil
	default:
		logger.Error(err)
		if errors.Is(err, gocb.ErrDocumentNotFound) {
			return nil, nil
		}

		return nil, err
	}
}

// GetEmailLogByRecipient is method to get email logs by recipient
func (e *emailLogRepository) GetEmailLogByRecipient(ctx context.Context, recipient string) ([]entity.EmailLog, error) {
	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"recipient": recipient,
	})

	query := `SELECT id, sources, recipient, status, created_at, updated_at, deleted_at FROM templates WHERE recipient = $1`
	rows, err := e.couchbase.Scope("email").Query(query, &gocb.QueryOptions{
		PositionalParameters: []interface{}{recipient},
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var emailLogs []entity.EmailLog
	for rows.Next() {
		var eml entity.EmailLog
		if err = rows.Row(&eml); err != nil {
			logrus.Error(err)
			return nil, err
		}

		emailLogs = append(emailLogs, eml)
	}

	if err = rows.Err(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	// if not found
	if len(emailLogs) == 0 {
		return nil, nil
	}

	// success get
	return emailLogs, nil
}
