package main

type Penalty struct {
	// represent penalty structure
	// with content
	id            int
	content       []Content
	clause        string
	disciplines   []Discipline
	costInSeconds int
	amount        int
	active        bool
	meta          Meta
}

type PenaltyInterface interface {
	AvailableBy(id int)
	Clause(clause string)
	Content(content []Content)
	In(disciplines []Discipline)
	Cost(costInSeconds int)
	Amount(amount int)
	Active(active bool)
	WithMeta(meta Meta)
}

type PenaltyBuilder struct {
	penalty *Penalty
}

func (b *PenaltyBuilder) AvailableBy(id int) *PenaltyIdBuilder {
	// add Meta object to SlotResult
	b.penalty.id = id
	return &PenaltyIdBuilder{*b}
}

type PenaltyIdBuilder struct {
	// builder for meta
	PenaltyBuilder
}

func (b *PenaltyBuilder) Clause(clause string) *PenaltyClauseBuilder {
	// add Meta object to SlotResult
	b.penalty.clause = clause
	return &PenaltyClauseBuilder{*b}
}

type PenaltyClauseBuilder struct {
	// builder for meta
	PenaltyBuilder
}

func (b *PenaltyBuilder) Content(content []Content) *PenaltyContentBuilder {
	// add Meta object to SlotResult
	for i := range content {
		b.penalty.content = append(b.penalty.content, content[i])
	}
	return &PenaltyContentBuilder{*b}
}

type PenaltyContentBuilder struct {
	// builder for meta
	PenaltyBuilder
}

func (b *PenaltyBuilder) In(disciplines []Discipline) *PenaltyDisciplineBuilder {
	// add Meta object to SlotResult
	for i := range disciplines {
		b.penalty.disciplines = append(b.penalty.disciplines, disciplines[i])
	}
	return &PenaltyDisciplineBuilder{*b}
}

type PenaltyDisciplineBuilder struct {
	// builder for meta
	PenaltyBuilder
}

func (b *PenaltyBuilder) Cost(costInSeconds int) *PenaltyCostBuilder {
	// add Meta object to SlotResult
	b.penalty.costInSeconds = costInSeconds
	return &PenaltyCostBuilder{*b}
}

type PenaltyCostBuilder struct {
	// builder for meta
	PenaltyBuilder
}

func (b *PenaltyBuilder) Amount(amount int) *PenaltyAmountBuilder {
	// add Meta object to SlotResult
	b.penalty.amount = amount
	return &PenaltyAmountBuilder{*b}
}

type PenaltyAmountBuilder struct {
	// builder for meta
	PenaltyBuilder
}

func (b *PenaltyBuilder) Active(active bool) *PenaltyActiveBuilder {
	// add Meta object to SlotResult
	b.penalty.active = active
	return &PenaltyActiveBuilder{*b}
}

type PenaltyActiveBuilder struct {
	// builder for meta
	PenaltyBuilder
}

func (b *PenaltyBuilder) WithMeta(meta Meta) *PenaltyMetaBuilder {
	// add Meta object to SlotResult
	b.penalty.meta = meta
	return &PenaltyMetaBuilder{*b}
}

type PenaltyMetaBuilder struct {
	// builder for meta
	PenaltyBuilder
}

func (b *PenaltyBuilder) Build() *Penalty {
	// Slot result by specific Course
	return b.penalty
}
