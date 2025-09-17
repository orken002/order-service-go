package repository

import (
	"OrderProject/entities"
	"gorm.io/gorm"
)

type CustomerSegmentsRepository interface {
	AddSegment(segments entities.CustomerSegments)
	UpdateSegment(segments entities.CustomerSegments)
	DeleteSegment(id int)
	GetAllSegments() []entities.CustomerSegments
	GetByIdSegment(id int) entities.CustomerSegments
}

type customerSegmentsRepository struct {
	gormDB *gorm.DB
}

func NewCustomerSegmentsRepository(gormDB *gorm.DB) CustomerSegmentsRepository {
	return &customerSegmentsRepository{gormDB: gormDB}
}

func (segmentRepository *customerSegmentsRepository) AddSegment(segments entities.CustomerSegments) {
	segmentRepository.gormDB.Create(&segments)
}

func (segmentRepository *customerSegmentsRepository) UpdateSegment(segments entities.CustomerSegments) {
	segmentRepository.gormDB.Save(&segments)
}

func (segmentRepository *customerSegmentsRepository) DeleteSegment(id int) {
	segmentRepository.gormDB.Delete(&entities.CustomerSegments{}, id)
}

func (segmentRepository *customerSegmentsRepository) GetAllSegments() []entities.CustomerSegments {
	var segments []entities.CustomerSegments
	segmentRepository.gormDB.Find(&segments)
	return segments
}

func (segmentRepository *customerSegmentsRepository) GetByIdSegment(id int) entities.CustomerSegments {
	var segment entities.CustomerSegments
	segmentRepository.gormDB.First(&segment, id)
	return segment
}
