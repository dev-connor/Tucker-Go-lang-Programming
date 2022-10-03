package main

import (
	"ch26/ex26.1/appB/exB5/marketplaceV2"
)

type ItemPurchaser interface { // 인터페이스 정의
	PurchaseItem()
}

func main() {
	var purchaser ItemPurchaser
	purchaser = marketplaceV2.NewMarketplace() // 1
	purchaser.PurchaseItem()
	// ...
}
