package repositories

import (
	"github.com/2gbeh/drygin-api/configs/db_config"
	"github.com/2gbeh/drygin-api/models"
	"gorm.io/gorm"
)

type Resource models.Tag
type Collection []models.Tag

var collection Collection
var resource Resource

type ITagRepository interface {
	SelectAll() (res Collection, err error)
	// SelectOne(id int) (res Resource, err error)
	// Select(field string, value any) (res Collection, err error)
	// Insert(params Resource) (res int, err error)
	InsertMany(params Collection) (res int64, err error)
	// Update(params Resource, id int) (res Resource, err error)
	// DeleteAll() (res int, err error)
	// DeleteOne(id int) (res Resource, err error)
	// Delete(field string, value any) (res int, err error)
	Migrate() (res Resource, err error)
	DropTable() (res bool, err error)
}

type TagRepository struct {
	db *gorm.DB
	tb *db_config.Table
}

func UseTagRpository(db *gorm.DB, tb *db_config.Table) ITagRepository {
	return &TagRepository{db: db, tb: tb}
}

func (repository *TagRepository) SelectAll() (res Collection, err error) {
	result := repository.db.Find(&collection)
	return collection, result.Error
}

func (repository *TagRepository) InsertMany(params Collection) (res int64, err error) {
	result := repository.db.Create(params)
	return result.RowsAffected, result.Error
}

func (repository *TagRepository) Migrate() (res Resource, err error) {
	result := repository.db.Table(repository.tb.Tags).AutoMigrate(&Resource{})
	return Resource{}, result
}

func (repository *TagRepository) DropTable() (res bool, err error) {
	result := repository.db.Exec("DROP TABLE " + repository.tb.Tags)
	return true, result.Error
}