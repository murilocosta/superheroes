package superhero

import (
	"github.com/jinzhu/gorm"
)

type SuperRepository interface {
	Save(s *Super) error
	SaveMany(s []*Super) error
	Delete(id int64) error
	FindByID(id int64) (*Super, error)
	FindByUUID(uuid int64) (*Super, error)
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
	return repo.db.Delete(&id).Error
}

func (repo *superRepositoryImpl) FindByID(id int64) (*Super, error) {
	var result Super
	err := repo.db.Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo *superRepositoryImpl) FindByUUID(uuid int64) (*Super, error) {
	var result Super
	err := repo.db.Where("uuid = ?", uuid).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repo *superRepositoryImpl) FindByName(name string) ([]*Super, error) {
	var result []*Super
	err := repo.db.Where("name LIKE ?", "%"+name+"%").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *superRepositoryImpl) List(superType SuperType) ([]*Super, error) {
	var result []*Super
	var err error
	if superType == "" {
		err = repo.db.Find(&result).Error
	} else {
		err = repo.db.Where("type LIKE ?", superType).Find(&result).Error
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}
