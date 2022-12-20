package apis

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/search"
)

// bindProjectApi registers the project api endpoint and the corresponding handlers
func bindProjectApi(app core.App, rg *echo.Group) {
	api := projectApi{app: app}

	subGroup := rg.Group("/projects", ActivityLogger(app), RequireAdminAuth())
	subGroup.GET("", api.list)
	subGroup.POST("", api.create)
}

type projectApi struct {
	app core.App
}

func (api *projectApi) list(c echo.Context) error {
	fieldResolver := search.NewSimpleFieldResolver(
		"id", "created", "updated", "name",
	)

	project := []*models.Project{}

	result, err := search.NewProvider(fieldResolver).
		Query(api.app.Dao().ProjectQuery()).
		ParseAndExec(c.QueryParams().Encode(), &project)

	if err != nil {
		return NewBadRequestError("", err)
	}

	event := &core.ProjectsListEvent{
		HttpContext: c,
		Project: project,
		Result: result,
	}

	return api.app.OnProjectsListRequest().Trigger(event, func(e *core.ProjectsListEvent) error {
		return e.HttpContext.JSON(http.StatusOK, e.Result)
	})
}

func (api *projectApi) create(c echo.Context) error {
	project := &models.Project{}

	form := forms.NewProjectUpsert(api.app, project)

	// load request
	if err := c.Bind(form); err != nil {
		return NewBadRequestError("Failed to load the submitted data due to invalid formatiing", err)
	}

	event := &core.ProjectCreateEvent{
		HttpContext: c,
		Project:  project,
	}

	// create the project
	submitErr := form.Submit(func(next forms.InterceptorNextFunc) forms.InterceptorNextFunc {
		return func() error {
			return api.app.OnProjectBeforeCreateRequest().Trigger(event, func(e *core.ProjectCreateEvent) error {
				if err := next(); err != nil {
					return NewBadRequestError("Failed to create the collection.", err)
				}

				return e.HttpContext.JSON(http.StatusOK, e.Project)
			})
		}
	})

	if submitErr == nil {
		api.app.OnProjectAfterCreateRequest().Trigger(event)
	}

	return submitErr

}