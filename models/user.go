package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/core/errors"
	"github.com/nutsdo/go-next/database"
)

var (
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("User not found")
	ErrUserExist = errors.New("User has exist.")
)

//用户
type User struct {
	gorm.Model
	Username   string `gorm:"size:60" json:"username"`
	Phone      string `gorm:"size:20" json:"phone" validate:"required"`
	Password   string `gorm:"size:60" json:"-"`
	Nickname   string `gorm:"size:60" json:"nickname"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	IsDisabled bool   `json:"is_disabled"`
}

/**
 * 通过 id 获取 user 记录
 * @method GetUserById
 * @param  {[type]}       user  *User [description]
 */
func GetUserById(id uint) (*User, error) {
	user := new(User)
	user.ID = id

	if err := database.DB.First(user).Error; err != nil {
		fmt.Printf("GetUserByIdErr:%s", err)
		return nil, err
	}

	return user, nil
}


/**
 * 通过 username 获取 user 记录
 * @method GetUserByUserName
 * @param  {[type]}       user  *User [description]
 */
func FindUserByUserName(username string) (*User, error) {
	user := &User{Username: username}

	notFound := database.DB.Where(user).First(user).RecordNotFound()

	if notFound {
		return nil, ErrUserNotFound
	}

	return user, nil
}


/**
 * 通过 Phone 获取 user 记录
 * @method FindUserByPhone
 * @param  {[type]}       user  *User [description]
 */
func FindUserByPhone(phone string) (*User, error) {

	user := new(User)
	notFound := database.DB.Where("phone = ?", phone).
		First(user).RecordNotFound()

	if notFound {
		return nil, ErrUserNotFound
	}

	return user, nil
}


/**
 * 创建
 * @method CreateUser
 * @param  {[type]} kw string [description]
 * @param  {[type]} cp int    [description]
 * @param  {[type]} mp int    [description]
 */
func CreateUser(u *User) (*User, error) {

	if _,err := FindUserByPhone(u.Phone); err==nil {
		return nil,ErrUserExist
	}

	//if _,err := FindUserByUserName(u.Username); err==nil {
	//	return nil,ErrUserExist
	//}

	if err := database.DB.Create(u).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
		return nil, err
	}

	return u, nil
}
//用户微信信息
type OauthWeixin struct {
	gorm.Model
	UserId     uint64 `json:"user_id"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	Nickname   string `json:"nickname"`
	Sex        int    `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Headimgurl string `json:"headimgurl"`
}

//用户基础信息
type UserBaseinfo struct {
	gorm.Model
	UserId    uint64 `json:"user_id"`
	Birthday  string `json:"birthday"`
	Signature string `json:"signature"`
	Sex       int    `json:"sex"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Country   string `json:"country"`
}

//用户实名认证
type UserRealnameAuth struct {
	gorm.Model
	UserId   uint64 `json:"user_id"`
	Realname string `json:"realname"`
	IdType   string `json:"id_type"`
	Idno     string `json:"idno"`
	IdHands  string `json:"id_hands"`
}

//店铺认证信息
type ShopAuthentications struct {
	gorm.Model
	UserId          uint64 `json:"user_id"`
	AuthType        int    `json:"auth_type"`
	Truename        string `json:"truename"`
	Idno            string `json:"idno"`
	IdFront         string `json:"id_front"`
	IdBack          string `json:"id_back"`
	IdHands         string `json:"id_hands"`
	EnterpriseName  string `json:"enterprise_name"`
	BusinessLicense string `json:"business_license"`
	AuthStatus      int    `json:"auth_status"`
}

type ShopAuthVerify struct {
	gorm.Model
	UserId       uint64 `json:"user_id"`
	ShopAuthId   uint64 `json:"shop_auth_id"`
	AuthStatus   int    `json:"auth_status"`
	RefuseReason string `json:"refuse_reason"`
}

//用户账户
type UserAccounts struct {
	gorm.Model
	UserId           uint64  `json:"user_id"`
	pay_password     string  `json:"pay_password" `
	balance          float64 `json:"balance"`
	frozen_capital   float64 `json:"frozen_capital"`
	security_deposit float64 `json:"security_deposit"`
}

//用户银行卡
type UserBankcards struct {
	gorm.Model
	UserId     uint64 `json:"user_id"`
	Realname   string `json:"realname"`
	BankcardNo string `json:"bankcard_no"`
	BankId     string `json:"bank_id"`
	Phone      string `json:"phone"`
}

//用户账户交易明细
type AccountDetail struct {
	gorm.Model
	UserId    uint64  `json:"user_id"`
	PayType   int     `json:"pay_type"`
	TradeType int     `json:"trade_type"`
	TradeNo   int     `json:"trade_no"`
	Amount    float64 `json:"amount"`
}

type UserAddress struct {
	gorm.Model
	UserId    uint64 `json:"user_id"`
	IsDefault bool   `json:"is_default"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Province  string `json:"province"`
	City      string `json:"city"`
	District  string `json:"district"`
	Street    string `json:"street"`
	Address   string `json:"address"`
}

type ShopInfo struct {
	gorm.Model
	UserId       uint64 `json:"user_id"`
	ShopName     string `json:"shop_name"`
	Logo         string `json:"logo"`
	Introduction string `json:"introduction"`
	Telephone    string `json:"telephone"`
}

type Orders struct {
	gorm.Model
	UserId      uint64  `json:"user_id"`
	SellerId    uint64  `json:"seller_id"`
	ProductId   uint64  `json:"product_id"`
	TradeNo     string  `json:"trade_no"`
	Price       string  `json:"price"`
	PayPrice    float64 `json:"pay_price"`
	OrderStatus int     `json:"order_status"`
}

type OrderLogs struct {
	gorm.Model
	TradeNo      string `json:"trade_no"`
	Detail       string `json:"detail"`
	ChangeStatus int    `json:"change_status"`
}

type ShopStaff struct {
	gorm.Model
	UserId     uint64     `json:"user_id"`
	ShopId     uint64     `json:"shop_id"`
	StaffId    uint64     `json:"staff_id"`
	Permission Permission `json:"permission"`
}

type Permission struct {
}

type ProductCategory struct {
	gorm.Model
	Name       string `json:"name"`
	Alias      string `json:"alias"`
	Subtitle   string `json:"subtitle"`
	Icon       string `json:"icon"`
	Thumb      string `json:"thumb"`
	IsDisabled bool   `json:"is_disabled"`
	lft        int    `json:"lft"`
	rgt        int    `json:"rgt"`
	ParentId   uint64 `json:"parent_id"`
}

type CategoryAttribute struct {
	gorm.Model
	CategoryId    uint64 `json:"category_id"`
	Attribute     string `json:"attribute"`
	FieldName     string `json:"field_name"`
	AttributeType string `json:"attribute_type"`
	IsOptions     bool   `json:"is_options"`
}

type CategoryAttributeOptions struct {
	gorm.Model
	AttributeId uint64 `json:"attribute_id"`
	Option      string `json:"option"`
	OptionValue string `json:"option_value"`
}

type ProductLibraries struct {
	gorm.Model
	UserId    uint64 `json:"user_id"`
	Name      string `json:"name"`
	Remark    string `json:"remark"`
	Icon      string `json:"icon"`
	Thumb     string `json:"thumb"`
	IsDefault bool   `json:"is_default"`
}

type Product struct {
	gorm.Model
	UserId        uint64 `json:"user_id"`
	LibId         uint64 `json:"lib_id"`
	CategoryId    uint64 `json:"category_id"`
	Name          string `json:"name"`
	Attributes    string `json:"attributes"`
	Specification string `json:"specification"`
	SerialNum     string `json:"serial_num"`
	Stock         uint64 `json:"stock"`
	SaleType      int    `json:"sale_type"`
	IsSale        bool   `json:"is_sale"`
	FreeShipment  bool   `json:"free_shipment"`
	Detail        string `json:"detail"`
}

type ProductImages struct {
	gorm.Model
	ProductId uint64 `json:"product_id"`
	ImageUrl  string `json:"image_url"`
	Position  int    `json:"position"`
}

type ProductReviews struct {
	gorm.Model
	UserId    uint64 `json:"user_id"`
	ShopId    uint64 `json:"shop_id"`
	ProductId uint64 `json:"product_id"`
	Stars     int    `json:"stars"`
	Content   string `json:"content"`
}

type ShopAuctionSetup struct {
	gorm.Model
	UserId    uint64 `json:"user_id"`
	IsDeduct  string `json:"is_deduct"`
	IsAuth    bool   `json:"is_auth"`
	IsGradual bool   `json:"is_gradual"`
	IsBreak   bool   `json:"is_break"`
	IsReturn  bool   `json:"is_return"`
	Is_Level  bool   `json:"is_level"`
}

type UserFollower struct {
	gorm.Model
	UserId       uint64 `json:"user_id"`
	FollowerId   uint64 `json:"follower_id"`
	FollowedType int    `json:"followed_type"`
}

type UserFollowed struct {
	gorm.Model
	UserId     uint64 `json:"user_id"`
	FollowedId uint64 `json:"followed_id"`
}

type Video struct {
	gorm.Model
	UserId   uint64 `json:"user_id"`
	Title    string `json:"title"`
	VideoUrl string `json:"video_url"`
	Likes    uint64 `json:"likes"`
	Shares   uint64 `json:"shares"`
	Comments uint64 `json:"comments"`
}

type UserFavorite struct {
	gorm.Model
	UserId       uint64 `json:"user_id"`
	FavoriteId   uint64 `json:"favorite_id"`
	FavoriteType int    `json:"favorite_type"`
}
