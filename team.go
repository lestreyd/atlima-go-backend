package main

type Team struct {
	// represents Team structure
	// with title and disciplines info
	id         int
	title      string
	discipline Discipline
	meta       Meta
}

type TeamInterface interface {
	AvailableBy(id int)
	Title(title string)
	Discipline(discipline Discipline)
	WithMeta(meta Meta)
}

type TeamBuilder struct {
	team *Team
}

func (b *TeamBuilder) AvailableBy(id int) *TeamIdBuilder {
	// add meta to squad
	b.team.id = id
	return &TeamIdBuilder{*b}
}

type TeamIdBuilder struct {
	// builder for meta
	TeamBuilder
}

func (b *TeamBuilder) Title(title string) *TeamTitleBuilder {
	// add meta to squad
	b.team.title = title
	return &TeamTitleBuilder{*b}
}

type TeamTitleBuilder struct {
	// builder for meta
	TeamBuilder
}

func (b *TeamBuilder) Discipline(discipline Discipline) *TeamDisciplineBuilder {
	// add discipline to team
	b.team.discipline = discipline
	return &TeamDisciplineBuilder{*b}
}

type TeamDisciplineBuilder struct {
	// builder for discipline
	TeamBuilder
}

func (b *TeamBuilder) WithMeta(meta Meta) *TeamMetaBuilder {
	// add meta to Team
	b.team.meta = meta
	return &TeamMetaBuilder{*b}
}

type TeamMetaBuilder struct {
	// builder for meta
	TeamBuilder
}

func (b *TeamBuilder) Build() *Team {
	// Squad object builder
	return b.team
}
