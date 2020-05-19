package main

/**
This struct's handle the specific data that I needed for my project, fi you
need more info you have the complete struct's just below.
*/
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
	FirstName   *string `json:"firstName,omitempty"`
	LastName    *string `json:"lastName,omitempty"`
	CountryCode *string `json:"countryCode"`
	Phone       *string `json:"phone"`
}

/**
This struct's handle are the data that SquareSpace provides.
- If you see a []interface{} type is because of i don't know what if
  actually SquareSpace returns so feel free to change it!!!
*/

/*
func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Result struct {
	Result     []Order   `json:"result,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

 type Order struct {
	ID                     *string          `json:"id,omitempty"`
	OrderNumber            *string          `json:"orderNumber,omitempty"`
	CreatedOn              *string          `json:"createdOn,omitempty"`
	ModifiedOn             *string          `json:"modifiedOn,omitempty"`
	Channel                *string          `json:"channel,omitempty"`
	Testmode               *bool            `json:"testmode,omitempty"`
	CustomerEmail          *string          `json:"customerEmail,omitempty"`
	BillingAddress         *IngAddress      `json:"billingAddress,omitempty"`
	ShippingAddress        *IngAddress      `json:"shippingAddress,omitempty"`
	FulfillmentStatus      *string          `json:"fulfillmentStatus,omitempty"`
	LineItems              []LineItem       `json:"lineItems"`
	InternalNotes          []interface{}    `json:"internalNotes"`
	ShippingLines          []interface{}    `json:"shippingLines"`
	DiscountLines          []interface{}    `json:"discountLines"`
	FormSubmission         []FormSubmission `json:"formSubmission"`
	Fulfillments           []interface{}    `json:"fulfillments"`
	Subtotal               *DiscountTotal   `json:"subtotal,omitempty"`
	ShippingTotal          *DiscountTotal   `json:"shippingTotal,omitempty"`
	DiscountTotal          *DiscountTotal   `json:"discountTotal,omitempty"`
	TaxTotal               *DiscountTotal   `json:"taxTotal,omitempty"`
	RefundedTotal          *DiscountTotal   `json:"refundedTotal,omitempty"`
	GrandTotal             *DiscountTotal   `json:"grandTotal,omitempty"`
	ChannelName            *string          `json:"channelName,omitempty"`
	ExternalOrderReference interface{}      `json:"externalOrderReference"`
	FulfilledOn            *string          `json:"fulfilledOn,omitempty"`
	PriceTaxInterpretation *string          `json:"priceTaxInterpretation,omitempty"`
}

type IngAddress struct {
	FirstName   *string     `json:"firstName,omitempty"`
	LastName    *string     `json:"lastName,omitempty"`
	Address1    *string     `json:"address1"`
	Address2    interface{} `json:"address2"`
	City        *string     `json:"city"`
	State       *string     `json:"state"`
	CountryCode *string     `json:"countryCode"`
	PostalCode  *string     `json:"postalCode"`
	Phone       *string     `json:"phone"`
}

type DiscountTotal struct {
	Value    *string `json:"value,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

type FormSubmission struct {
	Label *string `json:"label,omitempty"`
	Value *string `json:"value,omitempty"`
}

type LineItem struct {
	ID             *string        `json:"id,omitempty"`
	VariantID      interface{}    `json:"variantId"`
	Sku            interface{}    `json:"sku"`
	ProductID      *string        `json:"productId,omitempty"`
	ProductName    *string        `json:"productName,omitempty"`
	Quantity       *int64         `json:"quantity,omitempty"`
	UnitPricePaid  *DiscountTotal `json:"unitPricePaid,omitempty"`
	VariantOptions interface{}    `json:"variantOptions"`
	Customizations interface{}    `json:"customizations"`
	ImageURL       *string        `json:"imageUrl,omitempty"`
	LineItemType   *string        `json:"lineItemType,omitempty"`
} */
