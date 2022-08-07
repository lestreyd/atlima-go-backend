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

type CommentInterface interface {
	For(commentableId int)
	WithId(commentId int)
	From(user User)
	Send(message string)
	To(replyId int)
	SetDeleted(deleted bool)
	WithMeta(meta Meta)
}

type Attachment struct {
	// contains information about
	// attached to post entity
	id             int
	attachmentType int
	attachmentData []byte
	meta           Meta
}

type AttachmentInterface interface {
	AvailableBy(id int)
	WithType(attachmentType int)
	WithData(attachmentData []byte)
	WithMeta(meta Meta)
}

type Post struct {
	// contains information about
	// post, likes, views and etc
	id          int
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

type PostInterface interface {
	AvailableBy(id int)
	From(creator User)
	PresentedAs(organizer Organizer)
	WithText(text string)
	WithComments(comments []Comment)
	WithAttachments(attachments []Attachment)
	WithMeta(meta Meta)
}
