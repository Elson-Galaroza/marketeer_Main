package itemmodels

type ItemRetrieveResponse struct {
	Code     int32   `json:"code,omitempty"`
	Status   string  `json:"status,omitempty"`
	Message  string `json:"message,omitempty"`
	ItemID   int    `json:"Item ID"`
	SellerID int    `json:"Seller ID"`

	// FirstName  string `json:"FirstName"`
	// LastName   string `json:"LastName"`
	// EMail      string `json:"E-Mail"`
	// ContactNum string `json:"ContactNumber"`
	// BirthDate  string `json:"BirthDate"`
	// Address    string `json:"Address"`
	// UserName   string `json:"Username"`
	ItemImageLinkURL string `json:"Item Image URL"`
	ItemName         string `json:"Item Name"`
	ItemPrice        int    `json:"Item Price"`
	ItemQty          int    `json:"Item Quantity"`
	ItemDesc         string `json:"Item Description"`
}
