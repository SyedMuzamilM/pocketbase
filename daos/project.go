package daos

import (
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/list"
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

// IsProjectNameUnique checks that there is no existing collection
// with the provided name (case insensitive!).
//
// Note: case insensitive check because the name is used also as a table name for the records.
func (dao *Dao) IsProjectNameUnique(name string, excludeIds ...string) bool {
	if name == "" {
		return false
	}

	query := dao.ProjectQuery().
		Select("count(*)").
		AndWhere(dbx.NewExp("LOWER([[name]])={:name}", dbx.Params{
			"name": strings.ToLower(name),
		})).
		Limit(1)

	if len(excludeIds) > 0 {
		uniqueExcludeIds := list.NonzeroUniques(excludeIds)
		query.AndWhere(dbx.NotIn("id", list.ToInterfaceSlice(uniqueExcludeIds)...))
	}

	var exists bool

	return query.Row(&exists) == nil && !exists
}

// FindProjectByName finds a single db project with the specified name and
// scans the result into m.
func (dao *Dao) FindProjectByName(m models.Model, name string) error {
	return dao.ModelQuery(m).Where(dbx.HashExp{"name": name}).Limit(1).One(m)
}


func (dao *Dao) DoesProjectExist(name string) bool {
	if name == "" {
		return false
	}

	query := dao.ProjectQuery().
		Select("COUNT(*)").
		AndWhere(dbx.NewExp("LOWER([name]])={:name}", dbx.Params{
			"name": strings.ToLower(name),
		})).
		Limit(1)
	
	var exists bool
	return query.Row(&exists) == nil && !exists
}


func (dao *Dao) SaveProject(project *models.Project) error {
	// Also create a new user collection (system) for that project
	return dao.RunInTransaction(func(txDao *Dao) error {
		if err := txDao.Save(project); err != nil {
			return err
		}
		return txDao.CreateUserCollectionForProject(project.Name)
	})
}