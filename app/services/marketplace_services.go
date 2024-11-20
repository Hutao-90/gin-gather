package services

import (
	messagePackages "debox/message"
	"errors"
)

type GameProductConfig struct {
	ID    int    `json:"id"`    //ID
	Name  string `json:"name"`  // 名称
	Value int    `json:"value"` // 值
	Type  int    `json:"type"`  // 1:钻石（1U可购买钻石数量） 2:魔法药水（隐身概率） 3:邀请好友可获得钻石数 4:魔法药水价格(U或者钻石 可通过1转换
}

// GetProductConfigALl 商品列表
func GetProductConfigALl(columns string) ([]GameProductConfig, error) {
	db := GetDbConnection()
	var productConfig []GameProductConfig
	err := db.Table("game_product_config").
		Select(columns).
		Order("created_at desc").
		Scan(&productConfig).Error

	return productConfig, err
}

// BuyProduct 购买产品
func BuyProduct(productId, num int) (bool, err error) {
	db := GetDbConnection()
	var productConfig GameProductConfig
	err = db.Where("id =?", productId).First(&productConfig).Error
	if err != nil {
		return bool, errors.New(messagePackages.ProductNotFound)
	}

	//var member Member
	// TODO 区块链钱包支付
	//if productConfig.Type == models.GameProductConfigTypeOne { // 钻石
	//	member.DiamondNum += productConfig.Value + num
	//	return false
	//} else if productConfig.Type == models.GameProductConfigTypeTwo { //魔法药水
	//	member.MagicPotionNum += productConfig.Value + num
	//	return false
	//
	//} else if productConfig.Type == models.GameProductConfigTypeThree { //魔法药水
	//
	//}
	//err = db.Save(&member).Error
	//if err != nil {
	//	return errors.New(messagePackages.BuyFailure)
	//}
	return

}
