package main

type Comment struct {
	// comment structure for Post
	commentableId int
	commentId     int
	user          User
	message       string
	replyId       int
	deleted       bool
	meta          Meta
}

type Attachment struct {
	// contains information about
	// attached to post entity
	id             int
	attachmentType int
	attachmentData byte
	meta           Meta
}

type Post struct {
	// contains information about
	// post, likes, views and etc
	content     Content
	text        string
	creator     User
	organizer   Organizer
	reposted    int
	comments    []Comment
	postViewers []User
	postLikes   []User
	attachments []Attachment
	meta        Meta
}
