package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("file", func() {
	Title("file export file")
	Description("ファイル返す")
	Scheme("http")
	Host("localhost:8080")
})

var _ = Resource("file", func() {
	BasePath("downloads")

	Action("file", func() {
		Description("Download file")
		Routing(GET("/"))
		Response(OK, "application/vnd.ms-excel")
		Response(NotFound)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger-ui/*filepath", "swagger-ui/dist")
	Files("/mock/*filepath", "assets/mock", func() {
		Description("employee.json(従業員詳細データ) / submission.json(書類一覧)")
	})
})
