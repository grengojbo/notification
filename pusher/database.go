package pusher

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/notification"
)

// SendDB - save message to database
// duplicate="resolved" check resolved_at, message_type, link
// duplicate="link" check resolved_at, message_type, link
func SendDB(message *notification.Message, duplicate string, db * gorm.DB) error {
	notice := notification.QorNotification{
		Title: message.Title,
		Body:  message.Body,
		Val: message.Val,
		MessageType: message.MessageType,
		// ResolvedAt:  message.ResolvedAt,
	}

	if strFrom, ok := message.From.(string); ok {
		// if message.From {
		notice.From = strFrom
	} else {
		notice.From = "system"
	}
	if strTo, ok := message.To.(string); ok {
		// if message.To && len(message.To.(string)) > 0 {
		notice.To = strTo
	} else {
		notice.To = "all"
	}
	if len(message.Link) > 0 {
		notice.Link = message.Link
	}
	if message.ItemID > 0 {
		notice.ItemID = message.ItemID
	}

	old := notification.QorNotification{}
	if duplicate == "link" {
		if !db.Where("message_type=? AND link=?", notice.MessageType, notice.Link).First(&old).RecordNotFound() {
			return nil
		}
	} else if duplicate == "resolved" {
		if !db.Where("resolved_at IS NULL AND message_type=? AND link=?", notice.MessageType, notice.Link).First(&old).RecordNotFound() {
			return nil
		}
	}
	return db.Save(&notice).Error
}
