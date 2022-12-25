package forms_test

import (
	"testing"

	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tests"
)

func TestNewProjectUpsert(t *testing.T) {
	app, _ := tests.NewTestApp()
	defer app.Cleanup()

	project := &models.Project{}
	project.Name = "test_name"

	form := forms.NewProjectUpsert(app, project)

	if form.Name != project.Name {
		t.Errorf("Expected Name %q, got %q", project.Name, form.Name)
	}
}