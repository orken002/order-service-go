package services

import (
	"OrderProject/entities"
	"OrderProject/repository"
)

type CustomerSegmentsService interface {
	AddSegment(segments entities.CustomerSegments)
	UpdateSegment(segments entities.CustomerSegments)
	DeleteSegment(id int)
	GetAllSegments() []entities.CustomerSegments
	GetByIdSegment(id int) entities.CustomerSegments
}

type customerSegmentsService struct {
	customerSegmentsRepository repository.CustomerSegmentsRepository
}

func NewCustomerSegmentsServices (customerSegmentsRepository repository.CustomerSegmentsRepository) CustomerSegmentsService {
	return &customerSegmentsService{customerSegmentsRepository: customerSegmentsRepository}
}

func (customerSegmentsService *customerSegmentsService) AddSegment(segments entities.CustomerSegments) {
	customerSegmentsService.customerSegmentsRepository.AddSegment(segments)
}

func (customerSegmentsService *customerSegmentsService) UpdateSegment(segments entities.CustomerSegments) {
	customerSegmentsService.customerSegmentsRepository.UpdateSegment(segments)
}

func (customerSegmentsService *customerSegmentsService) DeleteSegment(id int) {
	customerSegmentsService.customerSegmentsRepository.DeleteSegment(id)
}

func (customerSegmentsService *customerSegmentsService) GetAllSegments() []entities.CustomerSegments {
	return customerSegmentsService.customerSegmentsRepository.GetAllSegments()
}

func (customerSegmentsService *customerSegmentsService) GetByIdSegment(id int) entities.CustomerSegments {
	return customerSegmentsService.customerSegmentsRepository.GetByIdSegment(id)
}