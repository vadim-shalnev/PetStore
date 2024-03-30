package storeService

import "github.com/vadim-shalnev/PetStore/models"

func NewStoreservice(repo *StoreRepository) *Storeservice {
	return &Storeservice{
		repo: repo,
	}
}

func (s *Storeservice) NewOrder(order models.Order) (models.Order, error) {
	err := s.repo.NewOrder(&order)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (s *Storeservice) GetOrder(id int) (models.Order, error) {
	order, err := s.repo.GetOrder(id)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (s *Storeservice) DeleteOrder(id int) error {
	err := s.repo.DeleteOrder(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storeservice) Getinventory(inventory models.Inventory) (models.Inventory, error) {
	err := s.repo.Getinventory(&inventory)
	if err != nil {
		return models.Inventory{}, err
	}
	return inventory, nil
}
