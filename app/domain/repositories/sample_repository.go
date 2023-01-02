package repositories

import (
	"github.com/althenlimzixuan/gorm_mysql_di/app/domain/entities"
	"github.com/althenlimzixuan/gorm_mysql_di/service"
	"github.com/sirupsen/logrus"
)

type SampleDBRepository struct {
	DBService service.GormMySqlServiceInt
}

type SampleDBRepositoryInt interface {
	Create(*entities.SampleEntities) (*entities.SampleEntities, error)
	ReadById(uint) (*entities.SampleEntities, error)
	Update(*entities.SampleEntities) (*entities.SampleEntities, error)
	Delete(*entities.SampleEntities) (uint, error)
}

func ProvideSampleDBRepository(db_service service.GormMySqlServiceInt) SampleDBRepositoryInt {
	return &SampleDBRepository{DBService: db_service}
}

func (svc *SampleDBRepository) Create(new_entity *entities.SampleEntities) (*entities.SampleEntities, error) {
	return new_entity, nil
}

func (svc *SampleDBRepository) ReadById(entity_id uint) (*entities.SampleEntities, error) {
	result := &entities.SampleEntities{}
	err := svc.DBService.First(&result).Where("id=?", entity_id).Error

	if err != nil {
		logrus.Errorf("Error @ ReadByID: %v", err)
		return nil, err
	}

	return result, nil
}

func (svc *SampleDBRepository) Update(entity *entities.SampleEntities) (*entities.SampleEntities, error) {
	return entity, nil
}

func (svc *SampleDBRepository) Delete(entity *entities.SampleEntities) (uint, error) {
	return 1, nil
}
