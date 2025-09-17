package mappers

import (
	"OrderProject/dtos"
	"OrderProject/entities"
)

func MapToDTO(customer entities.Customer) dtos.CustomerDTO {
	var customerDTO dtos.CustomerDTO
	customerDTO.ID = customer.ID
	customerDTO.Name = customer.Name
	customerDTO.Email = customer.Email
	customerDTO.Email = customer.Email
	customerDTO.OrderID = customer.OrderID
	customerDTO.Order = customer.Order
	customerDTO.CustomerSegments = customer.CustomerSegments
	return customerDTO
}

func MapToDTOList(customers []entities.Customer) []dtos.CustomerDTO {
	var customerDTOList []dtos.CustomerDTO
	for _, customer := range customers {
		customerDTOList = append(customerDTOList, MapToDTO(customer))
	}
	return customerDTOList
}
