package sites

import (
	"github.com/gin-gonic/gin"
)

type SiteSerializer struct {
	C *gin.Context
	SiteModel SiteModel
}

type SiteResponse struct {

}

func (self *SiteSerializer) Response() SiteResponse {

}

type SitesSerializer struct {
	C *gin.Context
	SitesModel []SiteModel
}

type SitesResponse struct {

}

func (self *SitesSerializer) Response() []SitesResponse {

}
