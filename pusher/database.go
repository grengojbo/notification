package pusher

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/notification"
)

// SendDB - save message to database
func SendDB(message *notification.Message, db * gorm.DB) error {
	notice := notification.QorNotification{
		Title: message.Title,
		Body:  message.Body,
		// Val: message.Val,
		MessageType: message.MessageType,
		// ResolvedAt:  message.ResolvedAt,
	}

	if len(message.From.(string)) > 0 {
		notice.From = message.From.(string)
	} else {
		notice.From = "system"
	}
	if len(message.To.(string)) > 0 {
		notice.From = message.To.(string)
	} else {
		notice.From = "all"
	}
	if len(message.Link) > 0 {
		notice.Link = message.Link
	}

	return db.Save(&notice).Error
}