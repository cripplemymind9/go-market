package serviceerrs

import "fmt"

var (
	ErrPasswordHashingFailed = fmt.Errorf("password hashing failed")
	ErrInvalidPassword = fmt.Errorf("invalid password")

	ErrCannotSignToken = fmt.Errorf("cannot sign token")
	ErrCannotParseToken = fmt.Errorf("cannot parse token")

	ErrCannotCreateUser = fmt.Errorf("cannot create user")
	ErrUserAlreadyExists = fmt.Errorf("user already exists")
	ErrUserNotFound = fmt.Errorf("user not found")
	ErrCannotGetUser = fmt.Errorf("cannot get user")

	ErrProductAlreadyExists = fmt.Errorf("product already exists")
	ErrCannotCreateProduct = fmt.Errorf("cannot create product")
	ErrCannotGetProducts = fmt.Errorf("cannot get products")
	ErrNoProductsAvailable = fmt.Errorf("no products available")

	ErrCannotCreatePurchase = fmt.Errorf("cannot create purchase")
	ErrNoUserPurchasesFound = fmt.Errorf("user purchases not found")
	ErrCannotGetUserPurchases = fmt.Errorf("cannot get user purchases")
	ErrNoProductPurchasesFound = fmt.Errorf("product purchases not found")
	ErrCannotGetProductPurchases = fmt.Errorf("cannot get product purchases")
)