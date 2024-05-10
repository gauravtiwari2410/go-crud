package model

import "time"

type Product struct {
	Id               int       `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Sku              string    `json:"sku"`
	Name             string    `json:"name"`
	Stocks           int       `json:stocks`
	Price            int       `json:"price"`
	Image            string    `json:image`
	TotalFinalPrize  int       `json:"Total_final_prize"`
	TotalNormalPrize int       `json:"Total_normal_prize"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CategoryId       int       `json:"categoryId"`
	DiscountId       int       `json:"discountId"`
}
