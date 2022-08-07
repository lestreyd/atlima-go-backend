package main

type NotificationTemplate struct {
	id              int
	content         []Content
	systemEventType string
	description     string
	meta            Meta
}

type NotificationTemplateInterface interface {
	AvailableBy(id int)
	WithContent(content []Content)
	EventType(systemEventType string)
	Description(description string)
	WithMeta(meta Meta)
}
