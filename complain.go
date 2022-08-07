package main

type Complain struct {
	// complain representation structure
	// with user and moderator information
	// inside
	id          int
	objectType  int
	objectId    int
	user        User
	userIp      string
	reason      int
	status      int
	moderator   User
	moderatorIp string
	meta        Meta
}

type ComplainInterface interface {
	AvailableBy(id int)
	ApplyTo()
	Object(objectType int, id int)
	From()
	User(user User)
	WithUserIp(userIp string)
	WithReason(reason int)
	ModeratedBy()
	Moderator(user User)
	WithModeratorIp(moderatorIp string)
	WithMeta(meta Meta)
}

type ComplainBuilder struct {
	// Complain object builder
	complain *Complain
}

func (b *ComplainBuilder) AvailableBy(id int) *ComplainIdBuilder {
	// create id for complain availability
	b.complain.id = id
	return &ComplainIdBuilder{*b}
}

type ComplainIdBuilder struct {
	// builder for sport credentials
	ComplainBuilder
}

func (b *ComplainBuilder) ApplyTo() *ComplainAppliedBuilder {
	// create id, slug and site for sport availability
	return &ComplainAppliedBuilder{*b}
}

type ComplainAppliedBuilder struct {
	// builder for sport credentials
	ComplainBuilder
}

func (b *ComplainBuilder) Object(objectType int, objectId int) *ComplainObjectBuilder {
	// create object type and id
	b.complain.objectType = objectType
	b.complain.objectId = objectId
	return &ComplainObjectBuilder{*b}
}

type ComplainObjectBuilder struct {
	// builder for complain object
	ComplainBuilder
}

func (b *ComplainBuilder) From() *ComplainFromBuilder {
	// create From section
	return &ComplainFromBuilder{*b}
}

type ComplainFromBuilder struct {
	// builder for From
	ComplainBuilder
}

func (b *ComplainBuilder) User(user User) *ComplainUserBuilder {
	// User that send request
	b.complain.user = user
	return &ComplainUserBuilder{*b}
}

type ComplainUserBuilder struct {
	// builder for complain
	ComplainBuilder
}

func (b *ComplainBuilder) WithUserIp(userIp string) *ComplainUserIPBuilder {
	// create user ip
	b.complain.userIp = userIp
	return &ComplainUserIPBuilder{*b}
}

type ComplainUserIPBuilder struct {
	// builder user ip field
	ComplainBuilder
}

func (b *ComplainBuilder) WithReason(reason int) *ComplainReasonBuilder {
	// create reason
	b.complain.reason = reason
	return &ComplainReasonBuilder{*b}
}

type ComplainReasonBuilder struct {
	// builder for complain reason
	ComplainBuilder
}

func (b *ComplainBuilder) ModeratedBy() *ComplainModerationBlockBuilder {
	// moderated by section
	return &ComplainModerationBlockBuilder{*b}
}

type ComplainModerationBlockBuilder struct {
	// builder for Moderation()
	ComplainBuilder
}

func (b *ComplainBuilder) Moderator(moderator User) *ComplainModeratorBuilder {
	// create moderator in ModeratedBy() section
	b.complain.moderator = moderator
	return &ComplainModeratorBuilder{*b}
}

type ComplainModeratorBuilder struct {
	// builder for complain moderator
	ComplainBuilder
}

func (b *ComplainBuilder) WithModeratorIp(moderatorIp string) *ComplainModeratorIPBuilder {
	// create moderator ip
	b.complain.moderatorIp = moderatorIp
	return &ComplainModeratorIPBuilder{*b}
}

type ComplainModeratorIPBuilder struct {
	// moderator ip in complain
	ComplainBuilder
}

func (b *ComplainBuilder) WithMeta(meta Meta) *ComplainMetaBuilder {
	// create meta fields
	b.complain.meta = meta
	return &ComplainMetaBuilder{*b}
}

type ComplainMetaBuilder struct {
	// builder for complain meta
	ComplainBuilder
}

func (b *ComplainBuilder) Build() *Complain {
	// Complain object builder
	return b.complain
}
