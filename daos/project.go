package daos

import (
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/models"
)

// ProjectQuery returns a new Project select query.
func (dao *Dao) ProjectQuery() *dbx.SelectQuery {
	return dao.ModelQuery(&models.Project{})
}

// FindProjectByNameOrId finds a single collection by its name (case insensitive) or id.
func (dao *Dao) FindProjectByNameOrId(nameOrId string) (*models.Project, error) {
	model := &models.Project{}

	err := dao.ProjectQuery().
		AndWhere(dbx.NewExp("[[id]] = {:id} OR LOWER([[name]])={:name}", dbx.Params{
			"id":   nameOrId,
			"name": strings.ToLower(nameOrId),
		})).
		Limit(1).
		One(model)

	if err != nil {
		return nil, err
	}

	return model, nil
}

func (dao *Dao) SaveProject(project *models.Project) error {
	return dao.Save(project)
}