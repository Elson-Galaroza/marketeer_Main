package itemmodels

type SellItemInput struct {
	//SellerUserID    int    `json:"User ID"`
	//ItemID    int    `json:"Item ID"`
	ItemImageLinkURL string `json:"ItemImageURL"`
	ItemName         string `json:"ItemName"`
	ItemPrice        int    `json:"ItemPrice"`
	ItemQty          int    `json:"ItemQuantity"`
	ItemDesc         string `json:"ItemDescription"`
}
