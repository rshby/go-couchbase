package entity

import (
	"context"
	"time"
)

type (
	EmailLog struct {
		ID        uint64     `json:"id"`
		Sources   string     `json:"sources"`
		Recipient string     `json:"recipient"`
		Status    string     `json:"status"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	EmailLogRepository interface {
		CreateEmailLog(ctx context.Context, input *EmailLog) (*EmailLog, error)
		GetEmailLogByID(ctx context.Context, id uint64) (*EmailLog, error)
		GetEmailLogByRecipient(ctx context.Context, recipient string) ([]EmailLog, error)
	}
)
