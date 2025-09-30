package utils

import (
	"context"

	"tripatra-dct-service-config/database/model"

	"gorm.io/gorm"
)

func SendAppNotification(ctx context.Context, tx *gorm.DB, notifications model.Notifications) error {
	// Specify the model or table you want to query
	err := tx.WithContext(ctx).Create(&notifications).Error
	if err != nil {
		return err
	}

	return err
}
