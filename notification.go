package main

type Notification struct {
	id                   int
	isRead               bool
	objectType           int
	objectId             int
	targetUser           User
	systemEvent          string
	notificationTemplate NotificationTemplate
	meta                 Meta
}

type NotificationInterface interface {
	AvailableBy(id int)
	Read(isRead bool)
	Type(objectType int)
	Id(objectId int)
	For(user User)
	From(systemEvent string)
	Template(notificationTemplate NotificationTemplate)
	WithMeta(meta Meta)
}
