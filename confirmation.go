package main

type Confirmation struct {
	// information about User confirmation
	// activity (phone/email) where action
	// contains What (phone or email confirmation)
	// and data contains code
	id     int
	action string
	data   string
	target User
	status uint
	meta   Meta
}

type ConfirmationInterface interface {
	AvailableBy(id int)
	ProvideTo(user User)
	Action(action string)
	For(data string)
	InStatus(status uint)
	WithMeta(meta Meta)
}

type ConfirmationBuilder struct {
	confirmation *Confirmation
}

func (b *ConfirmationBuilder) AvailableBy(id int) *ConfirmationIdBuilder {
	// add Meta object to SlotResult
	b.confirmation.id = id
	return &ConfirmationIdBuilder{*b}
}

type ConfirmationIdBuilder struct {
	// builder for meta
	ConfirmationBuilder
}

func (b *ConfirmationBuilder) ProvideTo(target User) *ConfirmationProvideBuilder {
	// add Meta object to SlotResult
	b.confirmation.target = target
	return &ConfirmationProvideBuilder{*b}
}

type ConfirmationProvideBuilder struct {
	// builder for meta
	ConfirmationBuilder
}

func (b *ConfirmationBuilder) Action(action string) *ConfirmationActionBuilder {
	// add Meta object to SlotResult
	b.confirmation.action = action
	return &ConfirmationActionBuilder{*b}
}

type ConfirmationActionBuilder struct {
	// builder for meta
	ConfirmationBuilder
}

func (b *ConfirmationBuilder) For(data string) *ConfirmationDataBuilder {
	// add Meta object to SlotResult
	b.confirmation.data = data
	return &ConfirmationDataBuilder{*b}
}

type ConfirmationDataBuilder struct {
	// builder for meta
	ConfirmationBuilder
}

func (b *ConfirmationBuilder) InStatus(status uint) *ConfirmationStatusBuilder {
	// add Meta object to SlotResult
	b.confirmation.status = status
	return &ConfirmationStatusBuilder{*b}
}

type ConfirmationStatusBuilder struct {
	// builder for meta
	ConfirmationBuilder
}

func (b *ConfirmationBuilder) WithMeta(meta Meta) *ConfirmationMetaBuilder {
	// add Meta object to SlotResult
	b.confirmation.meta = meta
	return &ConfirmationMetaBuilder{*b}
}

type ConfirmationMetaBuilder struct {
	// builder for meta
	ConfirmationBuilder
}

func (b *ConfirmationBuilder) Build() *Confirmation {
	// Slot result by specific Course
	return b.confirmation
}
