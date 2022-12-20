package models_test

import (
	"testing"

	"github.com/pocketbase/pocketbase/models"
)

func TestProjectTableName(t *testing.T) {
	m := models.Project{}

	if m.TableName() != "_projects" {
		t.Fatalf("Unexected table name, got %q", m.TableName())
	}
}