# Squarespace Orders v0.1
This api checks every 24h your Squarespace orders api and retrieves the data of all
the orders and create docs for every email and stores the orders by email also saves
basic user info. You can customize the info that you want with the struct's that 
are commented in the types.go file.

You can use it if you want to get filtered data from a user (because squarespace
api don't provide any filter option).

## Firestore structure
> Users
> > [user email]
> > > User data
> > 
> > > Orders
> > > > Product name
> > > > > Product data 

## Google cloud platform
Create a project and create a Firestore instance.
Create credentials for your project and download the json file.

```zsh
export GOOGLE_APPLICATION_CREDENTIALS="[PATH]"
```

## Go config 
Get your Google cloud platform project ID and create your squarespace orders api
key.
```go
package main

const (
	projectID      = "YOUR GOOGLE PLATFORM PROJECT ID"
	ApiKey         = "YOUR SQUARESPACE API KEY"
	squareSpaceURL = "https://api.squarespace.com/1.0/commerce/orders/"
)
```


For testing the app in dev mode.
```go
go run *.go
```

For production or deploy.
```go
go build
```

Default data that I store is. The fields that have an interface type is because
I don't know what field type are so feel free to improve it for yourself.
```go
package main

type FilteredTicket struct {
	ID            string               `json:"id,omitempty"`
	OrderNumber   string               `json:"orderNumber,omitempty"`
	CreatedOn     string               `json:"createdOn,omitempty"`
	CustomerEmail string               `json:"customerEmail,omitempty"`
	Items         []SimplifiedLineItem `json:"lineItems,omitempty"`
	UserData      UserData             `json:"billingAddress,omitempty"`
}

type SimplifiedLineItem struct {
	ID          string `json:"id,omitempty"`
	ProductID   string `json:"productId,omitempty"`
	ProductName string `json:"productName,omitempty"`
	Quantity    int64  `json:"quantity,omitempty"`
}

type FilteredResult struct {
	Result     []FilteredTicket `json:"result,omitempty"`
	Pagination Pagination       `json:"pagination,omitempty"`
}

type Pagination struct {
	NextPageUrl    string `json:"nextPageUrl,omitempty"`
	NextPageCursor string `json:"nextPageCursor,omitempty"`
	HasNextPage    bool   `json:"hasNextPage,omitempty"`
}

type UserData struct {
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	CountryCode string `json:"countryCode"`
	Phone       string `json:"phone"`
}
```

You can customize the data that you want to save.
```go
package main

type Result struct {
	Result     []Order    `json:"result,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	NextPageUrl    string `json:"nextPageUrl,omitempty"`
	NextPageCursor string `json:"nextPageCursor,omitempty"`
	HasNextPage    bool   `json:"hasNextPage,omitempty"`
}

type Order struct {
	ID                     string           `json:"id,omitempty"`
	OrderNumber            string           `json:"orderNumber,omitempty"`
	CreatedOn              string           `json:"createdOn,omitempty"`
	ModifiedOn             string           `json:"modifiedOn,omitempty"`
	Channel                string           `json:"channel,omitempty"`
	Testmode               bool             `json:"testmode,omitempty"`
	CustomerEmail          string           `json:"customerEmail,omitempty"`
	BillingAddress         IngAddress       `json:"billingAddress,omitempty"`
	ShippingAddress        IngAddress       `json:"shippingAddress,omitempty"`
	FulfillmentStatus      string           `json:"fulfillmentStatus,omitempty"`
	LineItems              []LineItem       `json:"lineItems"`
	InternalNotes          []interface{}    `json:"internalNotes"`
	ShippingLines          []interface{}    `json:"shippingLines"`
	DiscountLines          []interface{}    `json:"discountLines"`
	FormSubmission         []FormSubmission `json:"formSubmission"`
	Fulfillments           []interface{}    `json:"fulfillments"`
	Subtotal               DiscountTotal    `json:"subtotal,omitempty"`
	ShippingTotal          DiscountTotal    `json:"shippingTotal,omitempty"`
	DiscountTotal          DiscountTotal    `json:"discountTotal,omitempty"`
	TaxTotal               DiscountTotal    `json:"taxTotal,omitempty"`
	RefundedTotal          DiscountTotal    `json:"refundedTotal,omitempty"`
	GrandTotal             DiscountTotal    `json:"grandTotal,omitempty"`
	ChannelName            string           `json:"channelName,omitempty"`
	ExternalOrderReference interface{}      `json:"externalOrderReference"`
	FulfilledOn            string           `json:"fulfilledOn,omitempty"`
	PriceTaxInterpretation string           `json:"priceTaxInterpretation,omitempty"`
}

type IngAddress struct {
	FirstName   string      `json:"firstName,omitempty"`
	LastName    string      `json:"lastName,omitempty"`
	Address1    string      `json:"address1"`
	Address2    string      `json:"address2"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	CountryCode string      `json:"countryCode"`
	PostalCode  string      `json:"postalCode"`
	Phone       string      `json:"phone"`
}

type DiscountTotal struct {
	Value    string `json:"value,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type FormSubmission struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

type LineItem struct {
	ID             string        `json:"id,omitempty"`
	VariantID      string        `json:"variantId"`
	Sku            string        `json:"sku"`
	ProductID      string        `json:"productId,omitempty"`
	ProductName    string        `json:"productName,omitempty"`
	Quantity       int64         `json:"quantity,omitempty"`
	UnitPricePaid  DiscountTotal `json:"unitPricePaid,omitempty"`
	VariantOptions interface{}   `json:"variantOptions"`
	Customizations interface{}   `json:"customizations"`
	ImageURL       string        `json:"imageUrl,omitempty"`
	LineItemType   string        `json:"lineItemType,omitempty"`
}
```