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

type ComplaintInterface interface {
	AvailableBy(id int)
	ApplyTo()
	Object(objectType int)
	WithId(objectId int)
	From()
	User(user User)
	WithUserIp(userIp string)
	WithReason(reason int)
	ModeratedBy()
	Moderator(user User)
	WithModeratorIp(moderatorIp string)
	WithMeta(meta Meta)
}
