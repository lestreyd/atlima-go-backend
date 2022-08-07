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

type TransactionInterface interface {
	Operation(operation string)
	TerminalKey(terminalKey string)
	Amount(amount string)
	Order(orderId int)
	CustomerKey(customerKey string)
	Language(language string)
	Recurrent(recurrent string)
	Success(success bool)
	Status(status string)
	Payment(paymentId string)
	ErrorCode(errorCode string)
	CardData(cardData string)
	CardNumber(cardNumber string)
	Expired(exp string)
	CardHolder(cardHolder string)
	CVV(cvv string)
	CheckType(checkType string)
	RequestKey(requestKey string)
	CardId(cardId string)
	Currency(currency Currency)
	Token(token string)
	IP(ip string)
	Description(description string)
	Data(data string)
	Phone(phone string)
	SendMail(sendMail bool)
	InfoEmail(infoEmail string)
	Rebill(rebillId string)
	PayForm(payform string)
	Receipt(receipt string)
	RedirectDueDate(redirectDueDate string)
	PaymentURL(paymentUrl string)
	Message(message string)
	Details(details string)
	Md(md string)
	Response(response string)
	WithMeta(meta Meta)
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

type OrderItemInterface interface {
	AvailableBy(id int)
	WithType(objectType string)
	WithId(objectId int)
	Amount(amount int)
	WithPromoCode(promoCode PromoCode)
	WithMeta(meta Meta)
}

type Order struct {
	// contains information about transaction
	// user and status of order
	id          int
	status      bool
	user        User
	transaction Transaction
	items       []OrderItem
	meta        Meta
}

type OrderInterface interface {
	AvailableBy(id int)
	For(user User)
	Status(status bool)
	WithTransaction(transaction Transaction)
	WithItems(items []OrderItem)
	WithMeta(meta Meta)
}
