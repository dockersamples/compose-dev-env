package usecase

import (
	"aquafarm-management/app/model"
	"aquafarm-management/app/repository"
)

// ItemUsecase menangani proses bisnis item
type ItemUsecase struct {
	itemRepository *repository.ItemRepository
}

// NewItemUsecase membuat instance baru ItemUsecase
func NewItemUsecase(r *repository.ItemRepository) *ItemUsecase {
	return &ItemUsecase{
		itemRepository: r,
	}
}

// Fetch mengambil semua item dari repository
func (u *ItemUsecase) Fetch() ([]model.Item, error) {
	return u.itemRepository.Fetch()
}

// Get mengambil item dari repository berdasarkan ID
func (u *ItemUsecase) Get(id string) (*model.Item, error) {
	return u.itemRepository.Get(id)
}

// Store menyimpan item baru ke repository
func (u *ItemUsecase) Store(item *model.Item) (*model.Item, error) {
	return u.itemRepository.Store(item)
}

// Update memperbarui item di repository
func (u *ItemUsecase) Update(id string, item *model.Item) error {
	i, err := u.itemRepository.Get(id)
	if err != nil {
		return err
	}
	item.ID = i.ID
	return u.itemRepository.Update(item)
}

// Delete menghapus item dari repository
func (u *ItemUsecase) Delete(id string) error {
	return u.itemRepository.Delete(id)
}
