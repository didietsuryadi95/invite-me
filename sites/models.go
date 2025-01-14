package sites

import (
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/didietsuryadi95/invite-me/common"
	"github.com/didietsuryadi95/invite-me/users"
)

type SiteModel struct {
	gorm.Model
	ID        uint   `json:"ID" gorm:"primary_key"`
	HashID    string `json:"HashID" gorm:"column:hash_id"`
	UserModel users.UserModel
	Template  string      `json:"template" gorm:"column:template"`
	Theme     string      `json:"theme" gorm:"column:theme"`
	Settings  interface{} `json:"settings" sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
	UrlType   string      `json:"urlType" gorm:"column:urlType"`
	EndDate   string      `json:"endDate" gorm:"column:endDate"`
	Icon      string      `json:"icon" gorm:"column:icon"`
	Language  string      `json:"language" gorm:"column:language"`
	Title     string      `json:"title" gorm:"column:title"`
	Slug      string      `json:"slug" gorm:"column:slug"`
	SubSlug   string      `json:"sub_slug" gorm:"column:sub_slug"`
	StartDate string      `json:"startDate" gorm:"column:startDate"`
	Modules   interface{} `json:"modules" sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

type WeddingSetting struct {
	Bride         BrideInfo   `json:"bride" gorm:"column:bride"`
	Groom         BrideInfo   `json:"groom" gorm:"column:groom"`
	WeddingInfo   PlaceInfo   `json:"wedding_info" gorm:"column:wedding_info"`
	ReceptionInfo []PlaceInfo `json:"reception_info" gorm:"column"`
	Flower        string      `json:"flower" gorm:"column:flower"`
}

type BrideInfo struct {
	FullName   string `json:"full_name" gorm:"column:full_name"`
	ShortName  string `json:"short_name" gorm:"column:short_name"`
	FatherName string `json:"father_name" gorm:"column:father_name"`
	MotherName string `json:"mother_name" gorm:"column:mother_name"`
}

type PlaceInfo struct {
	Address   string `json:"address" gorm:"column:address"`
	Maps      string `json:"maps" gorm:"column:maps"`
	StartTime string `json:"start_time" gorm:"column:start"`
	EndTime   string `json:"endTime" gorm:"column:endTime"`
}

type Modules struct {
	Language string   `json:"language" gorm:"column:language"`
	Contents []Module `json:"contents" gorm:"column:contents"`
}

type Module struct {
	ModuleCode string      `json:"moduleCode" gorm:"column:moduleCode"`
	Content    interface{} `json:"content" gorm:"column:content"`
}

type OrderModel struct {
	gorm.Model
	UserModel     users.UserModel
	SiteModel     SiteModel
	Price         float64 `json:"price" gorm:"column:price"`
	PriceWithTax  float64 `json:"price_with_tax" gorm:"column:price_with_tax"`
	UniqueCode    string  `json:"unique_code" gorm:"column:unique_code"`
	Discount      float64 `json:"discount" gorm:"discount"`
	Voucher       string  `json:"voucher" gorm:"column:voucher"`
	PriceTotals   float64 `json:"price_totals" gorm:"column:price_totals"`
	PaymentMethod string  `json:"payment_method" gorm:"column:payment_method"`
	OrderId       string  `json:"order_id" gorm:"column:order_id"`
	// Features        []string `json:"features" gorm:"column:features"`
	PaymentStatus  string `json:"payment_status" gorm:"column:payment_status"`
	VirtualAccount string `json:"virtual_account" gorm:"column:virtual_account"`
}

// Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&SiteModel{})
	db.AutoMigrate(&OrderModel{})
}

func FindOneSite(condition interface{}) (SiteModel, error) {
	db := common.GetDB()
	var model SiteModel
	tx := db.Begin()
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	return model, err
}

func FindSiteByUser(userModel users.UserModel, search, limit, offset string) ([]SiteModel, int, error) {
	db := common.GetDB()
	var models []SiteModel
	var count int

	if userModel.ID == 0 {
		return models, 0, nil
	}

	offset_int, err := strconv.Atoi(offset)
	if err != nil {
		offset_int = 0
	}

	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		limit_int = 20
	}

	tx := db.Begin()
	tx.Where(SiteModel{UserModel: userModel})
	if search != "" {
		tx.Where(SiteModel{UserModel: userModel, Slug: search})
	}

	db.Model(&models).Count(&count)
	db.Offset(offset_int).Limit(limit_int).Find(&models)

	err = tx.Commit().Error
	return models, count, err
}
