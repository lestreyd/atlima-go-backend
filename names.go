package main

type Names struct {
	// represents a Names information
	// about User
	id               int
	firstName        string
	lastName         string
	middleName       string
	nativeFirstName  string
	nativeLastName   string
	nativeMiddleName string
	meta             Meta
}

type NamesInterface interface {
	AvailableBy(id int)
	Iam()
	First(firstName string)
	Middle(middleName string)
	Last(lastName string)
	FirstNative(firstNative string)
	LastNative(lastNative string)
	MiddleNative(middleNative string)
	WithMeta(meta Meta)
}

type NameBuilder struct {
	names *Names
}

func (b *NameBuilder) AvailableBy(id int) *NamesIdBuilder {
	// add id to names array
	b.names.id = id
	return &NamesIdBuilder{*b}
}

type NamesIdBuilder struct {
	// id builder
	NameBuilder
}

func (b *NameBuilder) Iam() *IamBuilder {
	// Iam builder will create new names
	return &IamBuilder{*b}
}

type IamBuilder struct {
	// builder for names
	NameBuilder
}

func (b *NameBuilder) First(firstName string) *IamBuilder {
	// add firstName to Names
	b.names.firstName = firstName
	return &IamBuilder{*b}
}

func (b *NameBuilder) Last(lastName string) *IamBuilder {
	// add lastName to Names
	b.names.lastName = lastName
	return &IamBuilder{*b}
}

func (b *NameBuilder) Middle(middleName string) *IamBuilder {
	// add Middle name to Names
	b.names.middleName = middleName
	return &IamBuilder{*b}
}

func (b *NameBuilder) FirstNative(nativeFirstName string) *IamBuilder {
	// add native first name to array
	b.names.nativeFirstName = nativeFirstName
	return &IamBuilder{*b}
}

func (b *NameBuilder) LastNative(nativeLastName string) *IamBuilder {
	// add native last name to array
	b.names.nativeLastName = nativeLastName
	return &IamBuilder{*b}
}

func (b *NameBuilder) MiddleNative(nativeMiddleName string) *IamBuilder {
	// add native middle name to array
	b.names.nativeMiddleName = nativeMiddleName
	return &IamBuilder{*b}
}

func (b *NameBuilder) WithMeta(meta Meta) *NamesMetaBuilder {
	// add meta to array
	b.names.meta = meta
	return &NamesMetaBuilder{*b}
}

type NamesMetaBuilder struct {
	// builder for meta
	NameBuilder
}

func (b *NameBuilder) Build() *Names {
	// create names array
	return b.names
}
