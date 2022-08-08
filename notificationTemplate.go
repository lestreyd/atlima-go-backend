package main

type NotificationTemplate struct {
	// NotificationTemplate attaches to specific
	// SystemEvent and can generate Notification
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

type NotificationTemplateBuilder struct {
	notificationTemplate *NotificationTemplate
}

func (b *NotificationTemplateBuilder) AvailableBy(id int) *NotificationTemplateIdBuilder {
	// add Meta object to SlotResult
	b.notificationTemplate.id = id
	return &NotificationTemplateIdBuilder{*b}
}

type NotificationTemplateIdBuilder struct {
	// builder for meta
	NotificationTemplateBuilder
}

func (b *NotificationTemplateBuilder) WithContent(content []Content) *NotificationTemplateContentBuilder {
	// add Meta object to SlotResult
	for i := range content {
		b.notificationTemplate.content = append(b.notificationTemplate.content, content[i])
	}
	return &NotificationTemplateContentBuilder{*b}
}

type NotificationTemplateContentBuilder struct {
	// builder for meta
	NotificationTemplateBuilder
}

func (b *NotificationTemplateBuilder) EventType(systemEventType string) *NotificationTemplateEventBuilder {
	// add Meta object to SlotResult
	b.notificationTemplate.systemEventType = systemEventType
	return &NotificationTemplateEventBuilder{*b}
}

type NotificationTemplateEventBuilder struct {
	// builder for meta
	NotificationTemplateBuilder
}

func (b *NotificationTemplateBuilder) Description(description string) *NotificationTemplateDescriptionBuilder {
	// add Meta object to SlotResult
	b.notificationTemplate.description = description
	return &NotificationTemplateDescriptionBuilder{*b}
}

type NotificationTemplateDescriptionBuilder struct {
	// builder for meta
	NotificationTemplateBuilder
}

func (b *NotificationTemplateBuilder) WithMeta(meta Meta) *NotificationTemplateMetaBuilder {
	// add Meta object to SlotResult
	b.notificationTemplate.meta = meta
	return &NotificationTemplateMetaBuilder{*b}
}

type NotificationTemplateMetaBuilder struct {
	// builder for meta
	NotificationTemplateBuilder
}

func (b *NotificationTemplateBuilder) Build() *NotificationTemplate {
	// Privacy object
	return b.notificationTemplate
}
