package main

type Squad struct {
	// represent Squad structure
	// with number and comment info
	// about squad
	id       int
	number   int
	comment  string
	blocked  bool
	password string
	meta     Meta
}

type SquadInterface interface {
	AvailableBy(id int)
	SetNumber(number string)
	SetComment(comment string)
	Blocked(blocked bool)
	WithPassword(password string)
	WithMeta(meta Meta)
}

type SquadBuilder struct {
	squad *Squad
}

func (b *SquadBuilder) AvailableBy(id int) *SquadIdBuilder {
	// add content array
	b.squad.id = id
	return &SquadIdBuilder{*b}
}

type SquadIdBuilder struct {
	// builder for country of sport administrator
	SquadBuilder
}

func (b *SquadBuilder) SetComment(comment string) *SquadCommentBuilder {
	// add content array
	b.squad.comment = comment
	return &SquadCommentBuilder{*b}
}

type SquadCommentBuilder struct {
	// builder for country of sport administrator
	SquadBuilder
}

func (b *SquadBuilder) SetNumber(number int) *SquadNumberBuilder {
	// set number (usually number of Squad set automatically)
	b.squad.number = number
	return &SquadNumberBuilder{*b}
}

type SquadNumberBuilder struct {
	// builder for number
	SquadBuilder
}

func (b *SquadBuilder) Blocked(blocked bool) *SquadBlockedBuilder {
	// add blocked or not
	b.squad.blocked = blocked
	return &SquadBlockedBuilder{*b}
}

type SquadBlockedBuilder struct {
	// builder for country of sport administrator
	SquadBuilder
}

func (b *SquadBuilder) WithPassword(password string) *SquadPasswordBuilder {
	// add password, block squad
	b.squad.password = password
	b.squad.blocked = true
	return &SquadPasswordBuilder{*b}
}

type SquadPasswordBuilder struct {
	// builder for password for squad block
	SquadBuilder
}

func (b *SquadBuilder) WithMeta(meta Meta) *SquadMetaBuilder {
	// add meta to squad
	b.squad.meta = meta
	return &SquadMetaBuilder{*b}
}

type SquadMetaBuilder struct {
	// builder for meta
	SquadBuilder
}

func (b *SquadBuilder) Build() *Squad {
	// Squad object builder
	return b.squad
}
