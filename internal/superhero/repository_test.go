package superhero

import (
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Should get from migration scripts, maybe later.
const createSchema = `
	CREATE TABLE supers (
		id           serial,
		uuid         integer UNIQUE,
		"type"       varchar(8),
		"name"       varchar(100),
		full_name    varchar(250),
		intelligence integer,
		"power"      integer,
		occupation   varchar(100),
		"image"      varchar(100)
	);
`

func TestSuperRepositorySave(t *testing.T) {
	// Test setup
	db, err := gorm.Open("sqlite3", "./repository_test.db")
	if err != nil {
		t.Errorf("Could not create database instance:\n%s", err)
	}

	// Must create the schema first
	db.Exec(createSchema)

	expected := &Super{
		UUID: 111,
		Name: "test",
		Type: HeroType,
	}

	repo := NewSuperRepository(db)
	err = repo.Save(expected)

	if err != nil {
		t.Errorf("Could not save entity on database:\n%s", err)
	}

	var result []*Super
	db.Find(&result)
	if len(result) != 1 {
		t.Errorf("The entity was not saved on database")
	}

	if result[0].UUID != expected.UUID || result[0].Name != expected.Name || result[0].Type != expected.Type {
		t.Errorf("The entity saved is different from expected")
	}

	// Test teardown
	db.Close()
	err = os.Remove("./repository_test.db")
	if err != nil {
		log.Fatal(err)
	}
}
