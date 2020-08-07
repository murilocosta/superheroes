package superhero

import (
	"github.com/jinzhu/gorm"
)

type SuperRepository interface {
	Save(s *Super) error
	SaveMany(s []*Super) error
	Delete(id int64) error
	FindByUuid(uuid int64) (*Super, error)
	FindByName(name string) ([]*Super, error)
	List(superType SuperType) ([]*Super, error)
}

type superRepositoryImpl struct {
	db *gorm.DB
}

func NewSuperRepository(db *gorm.DB) SuperRepository {
	return &superRepositoryImpl{db}
}

func (repo *superRepositoryImpl) Save(s *Super) error {
	return repo.db.Create(s).Error
}

func (repo *superRepositoryImpl) SaveMany(ss []*Super) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		for _, s := range ss {
			if err := tx.Create(s).Error; err != nil {
				// Will rollback
				return err
			}
		}
		// Will commit
		return nil
	})
}

func (repo *superRepositoryImpl) Delete(id int64) error {
	return nil
}

func (repo *superRepositoryImpl) FindByUuid(uuid int64) (*Super, error) {
	return nil, nil
}

func (repo *superRepositoryImpl) FindByName(name string) ([]*Super, error) {
	return nil, nil
}

func (repo *superRepositoryImpl) List(superType SuperType) ([]*Super, error) {
	var result []*Super
	if superType == "" {
		repo.db.Find(&result)
	} else {
		repo.db.Where("type LIKE ?", superType).Find(&result)
	}
	return result, nil
}
