package main

type Transaction struct {
	// store information about merchant
	// transaction (operation type and
	// all available information from bank
	operation       string
	terminalKey     string
	amount          string
	orderId         int
	customerKey     string
	language        string
	recurrent       string
	success         bool
	status          string
	paymentId       string
	errorCode       string
	cardData        string
	cardNumber      string
	exp             string
	cardHolder      string
	cvv             string
	checkType       string
	requestKey      string
	cardId          string
	currency        string
	token           string
	ip              string
	description     string
	data            string
	phone           string
	sendMail        bool
	infoEmail       string
	rebillId        string
	payForm         string
	receipt         string
	redirectDueDate string
	paymentUrl      string
	message         string
	details         string
	md              string
	response        string
	meta            Meta
}

type OrderItem struct {
	// contains information about ordered item
	id         int
	objectType string
	objectId   int
	amount     int
	promoCode  PromoCode
	meta       Meta
}

type Order struct {
	// contains information about transaction
	// user and status of order
	status      bool
	user        User
	transaction Transaction
	items       []OrderItem
	meta        Meta
}
