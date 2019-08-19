package cmd

import (
	"fmt"
	"github.com/nutsdo/go-next/database"
	"github.com/nutsdo/go-next/models"
)

func Migrate()  {
	fmt.Println("migrating...")
	database.DB.AutoMigrate(&models.User{},&models.OauthWeixin{},&models.UserBaseinfo{},
		&models.UserRealnameAuth{},&models.ShopAuthentications{},&models.ShopAuthVerify{},
		&models.UserAccounts{},&models.UserBankcards{},&models.AccountDetail{},
		&models.UserAddress{},&models.ShopInfo{},&models.Orders{},
		&models.OrderLogs{},&models.ShopStaff{},&models.ProductCategory{},
		&models.CategoryAttribute{},&models.CategoryAttributeOptions{},&models.ProductLibraries{},
		&models.Product{},&models.ProductImages{},&models.ProductReviews{},
		&models.ShopAuctionSetup{},&models.UserFollower{},&models.UserFollowed{},
		&models.Video{},&models.UserFavorite{})

	defer database.DB.Close()
	fmt.Println("migrate complete...")
}
