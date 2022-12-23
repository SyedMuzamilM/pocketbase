package forms

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

// var collectionNameRegex = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_]*$`)

// CollectionUpsert is a [models.Collection] upsert (create/update) form.
type ProjectUpsert struct {
	app        core.App
	dao        *daos.Dao
	project    *models.Project

	Id         string        `form:"id" json:"id"`
	Name       string        `form:"name" json:"name"`
}

// NewProjectUpsert creates a new [ProjectUpsert] form with initializer
// config created from the provided [core.App] and [models.Collection] instances
// (for create you could pass a pointer to an empty Collection - `&models.Collection{}`).
//
// If you want to submit the form as part of a transaction,
// you can change the default Dao via [SetDao()].
func NewProjectUpsert(app core.App, project *models.Project) *ProjectUpsert {
	form := &ProjectUpsert{
		app:        app,
		dao:        app.Dao(),
		project: 	project,
	}

	// load defaults
	form.Id = form.project.Id
	form.Name = form.project.Name

	print(form)

	return form
}


// Submit validates the form and upserts the form's Projct model.
//
// On success the related record table schema will be auto updated.
//
// You can optionally provide a list of InterceptorFunc to further
// modify the form behavior before persisting it.
func (form *ProjectUpsert) Submit(interceptors ...InterceptorFunc) error {
	if err := form.Validate(); err != nil {
		return err
	}

	if form.project.IsNew() {
		// custom insertion id can be set only on create
		if form.Id != "" {
			form.project.MarkAsNew()
			form.project.SetId(form.Id)
		}
	}

	form.project.Name = form.Name
	// form.collection.ListRule = form.ListRule
	// form.collection.ViewRule = form.ViewRule
	// form.collection.CreateRule = form.CreateRule
	// form.collection.UpdateRule = form.UpdateRule
	// form.collection.DeleteRule = form.DeleteRule
	// form.collection.SetOptions(form.Options)

	return runInterceptors(func() error {
		return form.dao.SaveProject(form.project)
	}, interceptors...)
}

func (form *ProjectUpsert) Validate() error {
	return validation.ValidateStruct(form,
		validation.Field(
			&form.Id,
			validation.When(
				form.project.IsNew(),
				validation.Length(models.DefaultIdLength, models.DefaultIdLength),
				validation.Match(idRegex),
			).Else(validation.In(form.project.Id)),
		),
		validation.Field(
			&form.Name,
			validation.Required,
			validation.Length(1, 255),
		),
	)
}

// func (form *ProjectUpsert) checkUniqueName(value any) error {
// 	v, _ := value.(string)

// 	// ensure unique collection name
// 	if !form.dao.IsCollectionNameUnique(v, form.project.Id) {
// 		return validation.NewError("validation_collection_name_exists", "Collection name must be unique (case insensitive).")
// 	}

// 	// ensure that the collection name doesn't collide with the id of any collection
// 	if form.dao.FindById(&models.Collection{}, v) == nil {
// 		return validation.NewError("validation_collection_name_id_duplicate", "The name must not match an existing collection id.")
// 	}

// 	// ensure that there is no existing table name with the same name
// 	if (form.project.IsNew() || !strings.EqualFold(v, form.project.Name)) && form.dao.HasTable(v) {
// 		return validation.NewError("validation_collection_name_table_exists", "The collection name must be also unique table name.")
// 	}

// 	return nil
// }