package htmx

import (
	"html/template"
	"strings"
	"template/src/models"

	"github.com/gin-gonic/gin"
	htmxmodel "github.com/yudhaananda/go-common/htmx_model"
)

const (
	Dashboard = "Dashboard"
)

func (h *htmx) GetDashboard(ctx *gin.Context) {
	name := make(map[string]string)
	name["Name"] = strings.ToLower(Dashboard)
	tmpl := template.Must(template.ParseFiles("./src/view/middleware.html"))
	tmpl.Execute(ctx.Writer, name)
}

func (h *htmx) DashboardContent(ctx *gin.Context) {
	htmxGet := htmxmodel.HTMXGet{
		SectionName: Dashboard,
	}
	for _, feature := range models.Features {
		temp := htmxmodel.SideBar{
			Name: template.HTML(feature.Name),
			Link: template.HTML(feature.Link),
		}
		if feature.Name == Dashboard {
			temp.Active = template.HTML("active")
		}
		htmxGet.SideBar = append(htmxGet.SideBar, temp)
	}
	tmpl := template.Must(template.ParseFiles("./src/view/index.html"))
	tmpl.Execute(ctx.Writer, htmxGet)
}
