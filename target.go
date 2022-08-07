package main

type TargetType struct {
	// represent Target type
	// structure with allowedResults
	// as "A" or "ACD"
	id             int
	paper          bool
	image          []byte
	allowedResults []string
	content        Content
	meta           Meta
}

type TargetInterface interface {
	AvailableById(id int)
	SetContent(content Content)
	Paper(paper bool)
	Image(image []byte)
	Allowed(allowedResults []string)
	WithMeta(meta Meta)
}

type TargetTypeBuilder struct {
	targetType *TargetType
}

func (b *TargetTypeBuilder) AvailableById(id int) *TargetTypeAvailabilityBuilder {
	// set id for target types
	b.targetType.id = id
	return &TargetTypeAvailabilityBuilder{*b}
}

type TargetTypeAvailabilityBuilder struct {
	// id builder for target type
	TargetTypeBuilder
}

func (b *TargetTypeBuilder) SetContent(content Content) *TargetTypeContentBuilder {
	// set content for target type
	b.targetType.content = content
	return &TargetTypeContentBuilder{*b}
}

type TargetTypeContentBuilder struct {
	//content builder for target type
	TargetTypeBuilder
}

func (b *TargetTypeBuilder) Paper(paper bool) *TargetTypePaperBuilder {
	// paper or not paper
	b.targetType.paper = paper
	return &TargetTypePaperBuilder{*b}
}

type TargetTypePaperBuilder struct {
	// builder for paper targets
	TargetTypeBuilder
}

func (b *TargetTypeBuilder) Image(image []byte) *TargetTypeImageBuilder {
	// image for target
	b.targetType.image = image
	return &TargetTypeImageBuilder{*b}
}

type TargetTypeImageBuilder struct {
	TargetTypeBuilder
}

func (b *TargetTypeBuilder) Allowed(allowedResults []string) *TargetTypeARBuilder {
	// add allowed results in array
	// as NS, M, A, C, D
	b.targetType.allowedResults = allowedResults
	return &TargetTypeARBuilder{*b}
}

type TargetTypeARBuilder struct {
	TargetTypeBuilder
}

func (b *TargetTypeBuilder) WithMeta(meta Meta) *TargetTypeMetaBuilder {
	// add meta to target type
	b.targetType.meta = meta
	return &TargetTypeMetaBuilder{*b}
}

type TargetTypeMetaBuilder struct {
	// meta builder for target type
	TargetTypeBuilder
}

func (b *TargetTypeBuilder) Build() *TargetType {
	// target type builder object
	return b.targetType
}

type TargetSet struct {
	// represent set of targets
	// with specific amount and
	// alphaCost for target
	id         int
	targetType TargetType
	course     Course
	amount     int
	alphaCost  int
	meta       Meta
}

type TargetSetInterface interface {
	AvailableBy(id int)
	For(course Course)
	Type(targetType TargetType)
	Amount(amount int)
	AlphaCost(alphaCost int)
	WithMeta(meta Meta)
}

type TargetSetBuilder struct {
	// builder for target set
	targetSet *TargetSet
}

func (b *TargetSetBuilder) AvailableById(id int) *TargetSetAvailabilityBuilder {
	// set id for target set
	b.targetSet.id = id
	return &TargetSetAvailabilityBuilder{*b}
}

type TargetSetAvailabilityBuilder struct {
	// target set id builder
	TargetSetBuilder
}

func (b *TargetSetBuilder) For(course Course) *TargetSetForBuilder {
	// course builder for target set
	b.targetSet.course = course
	return &TargetSetForBuilder{*b}
}

type TargetSetForBuilder struct {
	TargetSetBuilder
}

func (b *TargetSetBuilder) Type(targetType TargetType) *TargetSetTypeBuilder {
	// set type for target set
	b.targetSet.targetType = targetType
	return &TargetSetTypeBuilder{*b}
}

type TargetSetTypeBuilder struct {
	// builder for target type
	TargetSetBuilder
}

func (b *TargetSetBuilder) Amount(targetType TargetType) *TargetSetAmountBuilder {
	// builder for targets amount in target set
	b.targetSet.targetType = targetType
	return &TargetSetAmountBuilder{*b}
}

type TargetSetAmountBuilder struct {
	// builder for amount of targets in target set
	TargetSetBuilder
}

func (b *TargetSetBuilder) AlphaCost(alphaCost int) *TargetSetAlphaCostBuilder {
	// set alpha cost in current target set
	b.targetSet.alphaCost = alphaCost
	return &TargetSetAlphaCostBuilder{*b}
}

type TargetSetAlphaCostBuilder struct {
	TargetSetBuilder
}

func (b *TargetSetBuilder) WithMeta(meta Meta) *TargetSetMetaBuilder {
	// add meta to target set
	b.targetSet.meta = meta
	return &TargetSetMetaBuilder{*b}
}

type TargetSetMetaBuilder struct {
	// builder for meta of target set
	TargetSetBuilder
}

func (b *TargetSetBuilder) Build() *TargetSet {
	// TargetSet Builder
	return b.targetSet
}
