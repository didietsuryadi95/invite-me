package sites

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/didietsuryadi95/invite-me/users"
)

func Site(router *gin.RouterGroup) {
	router.GET("/:slug", getEndUserSite)
	router.GET("/:slug/trackers", trackSite)
}

func AdminSite(router *gin.RouterGroup) {
	router.GET("", getSiteList)
	router.GET("/:id", getSiteById)
	router.PUT("/:id", updateSiteById)
	router.POST("/", createSiteById)
	router.PUT("/:id/orders", getOrderSiteById)
	router.GET("/:id/tracker", getSiteTrackerById)
	router.POST("/:id/order", createOrderSite)
}

func SiteSettings(router *gin.RouterGroup) {
	router.GET("/templates/", getTemplateProperty)
	router.GET("/verify-slug/", getTemplateProperty)
	router.GET("/templates/:code/modules", getModulesProperty)
	router.GET("/themes", getTemplateProperty)
	router.GET("/features", getTemplateProperty)
}

var userModel = users.UserModel{
	ID:       0,
	Username: "didietsuryadi",
	Email:    "didietsuryadi@gmail.com",
	Bio:      "",
	Image:    nil,
}

var siteModels = map[string]SiteModel{
	"id1": {
		UserModel: userModel,
		HashID:    "id1",
		Template:  "WEDDING",
		Theme:     "FLOWER_1",
		Settings: map[string]interface{}{
			"bride": map[string]string{
				"fullname":  "Larasati",
				"shortname": "Laras",
			},
			"groom": map[string]string{
				"fullname":  "Ikhsan Mahendri ST",
				"shortname": "Ikhsan",
			},
		},
		UrlType:   "SUBDOMAIN",
		EndDate:   "1741852352",
		Icon:      "",
		Language:  "ID",
		Title:     "Wedding Ikhsan & Laras",
		Slug:      "ikhsan-laras",
		SubSlug:   "",
		StartDate: "1647183152",
		Modules: []Modules{
			{
				Language: "ID",
				Contents: []Module{
					{
						ModuleCode: "HOME",
						Content: map[string]string{
							"date":      "1231231242343",
							"placeName": "Gedung Serba Guna Sapta Taruna 3",
						},
					},
					{
						ModuleCode: "EVENT",
						Content: []map[string]string{
							{
								"eventType": "Ijab Qobul",
								"startDate": "1231231231321",
								"endDate":   "1231231231321",
								"placeName": "Gedung Serba Guna Sapta Taruna 3",
								"address":   "Jl. Koperpu II No.32, RT.001/RW.034, Bojong Rawalumbu, Kec. Rawalumbu, Kota Bks, Jawa Barat 17116",
								"useMaps":   "true",
								"maps":      "https://www.google.com/maps/place/Komplek+Sapta+Taruna+III/@-6.2833277,106.987009,17z/data=!3m1!4b1!4m5!3m4!1s0x2e698dedf4033c4f:0xbb28452d1532f7ea!8m2!3d-6.2833277!4d106.9891977?shorturl=1",
							},
							{
								"eventType": "Wedding Reception",
								"startDate": "1231231231321",
								"endDate":   "1231231231321",
								"placeName": "Gedung Serba Guna Sapta Taruna 3",
								"address":   "Jl. Koperpu II No.32, RT.001/RW.034, Bojong Rawalumbu, Kec. Rawalumbu, Kota Bks, Jawa Barat 17116",
								"useMaps":   "false",
								"maps":      "https://www.google.com/maps/place/Komplek+Sapta+Taruna+III/@-6.2833277,106.987009,17z/data=!3m1!4b1!4m5!3m4!1s0x2e698dedf4033c4f:0xbb28452d1532f7ea!8m2!3d-6.2833277!4d106.9891977?shorturl=1",
							},
						},
					},
					{
						ModuleCode: "BRIDE_PROFILE",
						Content: []map[string]string{
							{
								"type": "TITLE",
								"text": "Bismillahirrohmanirrahiim \n In the Name of Allah Subhanahu Wa Ta`ala the Most Beneficent and the Most Merciful. It is our moment to celebrate the love that unites our children",
							},
							{
								"type":          "BRIDE",
								"childrenOrder": "1",
								"fatherName":    "Drs. Iwa Karsiwa, M.Pd",
								"motherName":    "Dini Sri Handayani",
							},
							{
								"type":          "GROOM",
								"childrenOrder": "1",
								"fatherName":    "Nurhendri",
								"motherName":    "Magdalena",
							},
						},
					},
					{
						ModuleCode: "GALLERY",
						Content: []map[string]string{
							{
								"url":       "https://blablabalbla",
								"type":      "PHOTO",
								"dimention": "Potrait",
							},
							{
								"url":       "https://blablabalbla",
								"type":      "PHOTO",
								"dimention": "Potrait",
							},
							{
								"url":       "https://blablabalbla",
								"type":      "PHOTO",
								"dimention": "Potrait",
							},
							{
								"url":       "https://blablabalbla",
								"type":      "PHOTO",
								"dimention": "Potrait",
							},
							{
								"url":       "https://blablabalbla",
								"type":      "PHOTO",
								"dimention": "Potrait",
							},
							{
								"url":       "https://blablabalbla",
								"type":      "PHOTO",
								"dimention": "Potrait",
							},
						},
					},
					{
						ModuleCode: "CHAT_BOX",
						Content: map[string]string{
							"title":       "Whishes",
							"description": "Send a Wishes",
							"referenceId": "1231233423123123",
						},
					},
				},
			},
			{
				Language: "EN",
				Contents: []Module{
					{
						ModuleCode: "HOME",
						Content: map[string]string{
							"date":      "1231231242343",
							"placeName": "Gedung Serba Guna Sapta Taruna 3",
						},
					},
					{
						ModuleCode: "EVENT",
						Content: []map[string]string{
							{
								"eventType": "Ijab Qobul",
								"startDate": "1231231231321",
								"endDate":   "1231231231321",
								"placeName": "Gedung Serba Guna Sapta Taruna 3",
								"address":   "Jl. Koperpu II No.32, RT.001/RW.034, Bojong Rawalumbu, Kec. Rawalumbu, Kota Bks, Jawa Barat 17116",
								"useMaps":   "true",
								"maps":      "https://www.google.com/maps/place/Komplek+Sapta+Taruna+III/@-6.2833277,106.987009,17z/data=!3m1!4b1!4m5!3m4!1s0x2e698dedf4033c4f:0xbb28452d1532f7ea!8m2!3d-6.2833277!4d106.9891977?shorturl=1",
							},
							{
								"eventType": "Wedding Reception",
								"startDate": "1231231231321",
								"endDate":   "1231231231321",
								"placeName": "Gedung Serba Guna Sapta Taruna 3",
								"address":   "Jl. Koperpu II No.32, RT.001/RW.034, Bojong Rawalumbu, Kec. Rawalumbu, Kota Bks, Jawa Barat 17116",
								"useMaps":   "false",
								"maps":      "https://www.google.com/maps/place/Komplek+Sapta+Taruna+III/@-6.2833277,106.987009,17z/data=!3m1!4b1!4m5!3m4!1s0x2e698dedf4033c4f:0xbb28452d1532f7ea!8m2!3d-6.2833277!4d106.9891977?shorturl=1",
							},
						},
					},
					{
						ModuleCode: "BRIDE_PROFILE",
						Content: []map[string]string{
							{
								"type": "TITLE",
								"text": "Bismillahirrohmanirrahiim \n In the Name of Allah Subhanahu Wa Ta`ala the Most Beneficent and the Most Merciful. It is our moment to celebrate the love that unites our children",
							},
							{
								"type":          "BRIDE",
								"childrenOrder": "1",
								"fatherName":    "Drs. Iwa Karsiwa, M.Pd",
								"motherName":    "Dini Sri Handayani",
							},
							{
								"type":          "GROOM",
								"childrenOrder": "1",
								"fatherName":    "Nurhendri",
								"motherName":    "Magdalena",
							},
						},
					},
					{
						ModuleCode: "GALLERY",
						Content: []map[string]string{
							{
								"url":       "https://blablabalbla",
								"type":      "PHOTO",
								"dimention": "Potrait",
							},
							{
								"url":       "https://blablabalbla",
								"type":      "PHOTO",
								"dimention": "Potrait",
							},
							{
								"url":       "https://blablabalbla",
								"type":      "PHOTO",
								"dimention": "Potrait",
							},
							{
								"url":       "https://blablabalbla",
								"type":      "PHOTO",
								"dimention": "Potrait",
							},
							{
								"url":       "https://blablabalbla",
								"type":      "PHOTO",
								"dimention": "Potrait",
							},
							{
								"url":       "https://blablabalbla",
								"type":      "PHOTO",
								"dimention": "Potrait",
							},
						},
					},
					{
						ModuleCode: "CHAT_BOX",
						Content: map[string]string{
							"title":       "Whishes",
							"description": "Send a Wishes",
							"referenceId": "1231233423123123",
						},
					},
				},
			},
		},
	},
	"id2": {
		UserModel: userModel,
		HashID:    "id2",
		Template:  "MICRO_SITE",
		Theme:     "BIO_1",
		Settings: map[string]string{
			"avatar": "https://image_url",
			"title":  "Ikhsan Mahendri",
			"bio":    "Welcome to My Content",
		},
		UrlType:   "SLUG",
		EndDate:   "1741852352",
		Icon:      "",
		Language:  "ID",
		Title:     "My Bio",
		Slug:      "ikhsan",
		SubSlug:   "",
		StartDate: "1647183152",
		Modules: []Modules{
			{
				Language: "ID",
				Contents: []Module{
					{
						ModuleCode: "LINK",
						Content: []map[string]string{
							{
								"url":   "https://blablabalbla",
								"color": "#fffff",
								"text":  "Ini Link",
							},
							{
								"url":   "https://blablabalbla",
								"color": "#fffff",
								"text":  "Ini Link",
							},
							{
								"url":   "https://blablabalbla",
								"color": "#fffff",
								"text":  "Ini Link",
							},
							{
								"url":   "https://blablabalbla",
								"color": "#fffff",
								"text":  "Ini Link",
							},
							{
								"url":   "https://blablabalbla",
								"color": "#fffff",
								"text":  "Ini Link",
							},
							{
								"url":   "https://blablabalbla",
								"color": "#fffff",
								"text":  "Ini Link",
							},
						},
					},
				},
			},
		},
	},
}

func getEndUserSite(c *gin.Context) {
	language := c.Query("language")
	if language == "" {
		language = "ID"
	}
	slugs := map[string]string{
		"wedding": "id1",
		"bio":     "id2",
	}

	slug := c.Param("slug")
	if slugs[slug] == "" {
		c.JSON(http.StatusNotFound, gin.H{"data": nil})
		return
	}
	siteSerializer := EndUserSiteSerializer{c, siteModels[slugs[slug]], language}
	c.JSON(http.StatusOK, gin.H{"data": siteSerializer.Response()})
}

func trackSite(c *gin.Context) {

}

func getSiteTrackerById(c *gin.Context) {

}

func getSiteList(c *gin.Context) {

	sitesSerializer := SitesSerializer{c, []SiteModel{siteModels["id1"], siteModels["id2"]}}
	c.JSON(http.StatusOK, gin.H{"data": sitesSerializer.Response()})

}

func getSiteById(c *gin.Context) {

	id := c.Param("id")
	siteSerializer := SiteSerializer{c, siteModels[id]}
	c.JSON(http.StatusOK, gin.H{"data": siteSerializer.Response()})

}

func updateSiteById(c *gin.Context) {

	siteSerializer := SiteSerializer{c, siteModels["id1"]}
	c.JSON(http.StatusOK, gin.H{"data": siteSerializer.Response()})

}

func getOrderSiteById(c *gin.Context) {

}

func createSiteById(c *gin.Context) {

	siteSerializer := SiteSerializer{c, siteModels["id1"]}
	c.JSON(http.StatusOK, gin.H{"data": siteSerializer.Response()})

}

func createOrderSite(c *gin.Context) {

}

func getModulesProperty(c *gin.Context) {

}

func getTemplateProperty(c *gin.Context) {

}
