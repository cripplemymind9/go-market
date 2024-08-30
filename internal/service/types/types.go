package types

type AuthRegisterUserInput struct {
	Username 	string
	Password 	string
	Email 		string
}

type AuthGenerateTokenInput struct {
	Username 	string
	Password 	string
}

type ProductAddProductInput struct {
	Name 		string
	Description string
	Price 		float64
	Quantity 	int
}

type ProductUpdateProductInput struct {
	Name 		string
	Description string
	Price 		float64
	Quantity 	int
}

type PurchaseMakePurchaseInput struct {
	UserID		int
	ProductID 	int
	Quantity 	int
}