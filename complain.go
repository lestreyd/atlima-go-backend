package main

type Complain struct {
	// complain representation structure
	// with user and moderator information
	// inside
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
