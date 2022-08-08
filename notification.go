package main

type Notification struct {
	// Notification generates from NotificationTemplate
	// by rules in NotificationTemplate
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

type NotificationBuilder struct {
	notification *Notification
}

func (b *NotificationBuilder) AvailableBy(id int) *NotificationIdBuilder {
	// add Meta object to SlotResult
	b.notification.id = id
	return &NotificationIdBuilder{*b}
}

type NotificationIdBuilder struct {
	// builder for meta
	NotificationBuilder
}

func (b *NotificationBuilder) Read(isRead bool) *NotificationReadBuilder {
	// add Meta object to SlotResult
	b.notification.isRead = isRead
	return &NotificationReadBuilder{*b}
}

type NotificationReadBuilder struct {
	// builder for meta
	NotificationBuilder
}

func (b *NotificationBuilder) Type(objectType int) *NotificationTypeBuilder {
	// add Meta object to SlotResult
	b.notification.objectType = objectType
	return &NotificationTypeBuilder{*b}
}

type NotificationTypeBuilder struct {
	// builder for meta
	NotificationBuilder
}

func (b *NotificationBuilder) Id(objectId int) *NotificationObjIdBuilder {
	// add Meta object to SlotResult
	b.notification.objectId = objectId
	return &NotificationObjIdBuilder{*b}
}

type NotificationObjIdBuilder struct {
	// builder for meta
	NotificationBuilder
}

func (b *NotificationBuilder) For(user User) *NotificationTargetBuilder {
	// add Meta object to SlotResult
	b.notification.targetUser = user
	return &NotificationTargetBuilder{*b}
}

type NotificationTargetBuilder struct {
	// builder for meta
	NotificationBuilder
}

func (b *NotificationBuilder) From(systemEvent string) *NotificationEventBuilder {
	// add Meta object to SlotResult
	b.notification.systemEvent = systemEvent
	return &NotificationEventBuilder{*b}
}

type NotificationEventBuilder struct {
	// builder for meta
	NotificationBuilder
}

func (b *NotificationBuilder) Template(notificationTemplate NotificationTemplate) *NotificationSourceTemplateBuilder {
	// add Meta object to SlotResult
	b.notification.notificationTemplate = notificationTemplate
	return &NotificationSourceTemplateBuilder{*b}
}

type NotificationSourceTemplateBuilder struct {
	// builder for meta
	NotificationBuilder
}

func (b *NotificationBuilder) WithMeta(meta Meta) *NotificationMetaBuilder {
	// add Meta object to SlotResult
	b.notification.meta = meta
	return &NotificationMetaBuilder{*b}
}

type NotificationMetaBuilder struct {
	// builder for meta
	NotificationBuilder
}

func (b *NotificationBuilder) Build() *Notification {
	// Privacy object
	return b.notification
}
