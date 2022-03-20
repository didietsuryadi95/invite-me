package sites

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type SiteSerializer struct {
	C         *gin.Context
	SiteModel SiteModel
}

type SiteResponse struct {
	Template  string      `json:"template"`
	Theme     string      `json:"theme"`
	Settings  interface{} `json:"settings"`
	UrlType   string      `json:"urlType"`
	EndDate   string      `json:"endDate"`
	Icon      string      `json:"icon"`
	Language  string      `json:"language"`
	Title     string      `json:"title"`
	Slug      string      `json:"slug"`
	SubSlug   string      `json:"sub_slug"`
	StartDate string      `json:"startDate"`
	Modules   interface{} `json:"modules"`
}

func (s *SiteSerializer) Response() SiteResponse {
	return SiteResponse{
		Template: s.SiteModel.Template,
		Theme: s.SiteModel.Theme,
		Settings: s.SiteModel.Settings,
		UrlType: s.SiteModel.UrlType,
		Slug: s.SiteModel.Slug,
		StartDate: s.SiteModel.StartDate,
		EndDate: s.SiteModel.EndDate,
		SubSlug: s.SiteModel.SubSlug,
		Icon: s.SiteModel.Icon,
		Language: s.SiteModel.Language,
		Modules: s.SiteModel.Modules,
	}
}

type SitesSerializer struct {
	C          *gin.Context
	SitesModel []SiteModel
}

type SitesResponse struct {
	ID       uint   `json:"ID"`
	Template string `json:"template"`
	Status   int    `json:"status"`
	Title    string `json:"title"`
}

func (s *SitesSerializer) Response() []SitesResponse {
	sites := make([]SitesResponse, 0)

	for _, model := range s.SitesModel {
		sites = append(sites, SitesResponse{
			ID:       model.ID,
			Template: model.Template,
			Status:   1,
			Title:    model.Title,
		})
	}

	return sites
}

type EndUserSiteSerializer struct {
	C          *gin.Context
	SiteModel SiteModel
	Language   string
}

type EndUserSiteResponse struct {
	Template  string      `json:"template"`
	Theme     string      `json:"theme"`
	Settings  interface{} `json:"settings"`
	Icon      string      `json:"icon"`
	Title     string      `json:"title"`
	Modules   interface{} `json:"modules"`
}

func (s *EndUserSiteSerializer) Response() EndUserSiteResponse {
	modules := Modules{}
	modulesModel := reflect.ValueOf(s.SiteModel.Modules).Interface().([]Modules)
	for _, m := range modulesModel {
		if strings.ToUpper(s.SiteModel.Language) == strings.ToUpper(m.Language) {
			modules = m
		}
	}
	return EndUserSiteResponse{
		Template: s.SiteModel.Template,
		Theme: s.SiteModel.Theme,
		Settings: s.SiteModel.Settings,
		Icon: s.SiteModel.Icon,
		Modules: modules.Contents,
	}
}
