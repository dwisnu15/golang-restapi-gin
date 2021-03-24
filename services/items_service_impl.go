package services
/*
Service for entity items
 */
import (
	"GinAPI/models"
	"GinAPI/repositories"
)

type ItemsServiceImpl struct {
	itemsRepo repositories.ItemsRepo
}

func CreateItemsService(itemsRepo repositories.ItemsRepo) ItemsService {
	return &ItemsServiceImpl{itemsRepo}
}

func (i ItemsServiceImpl) FindAllItem() (*[]models.Items, error) {
	return i.itemsRepo.FindAllItem()
}

func (i ItemsServiceImpl) FindItemByID(itemID int64) (*models.Items, error) {
	return i.itemsRepo.FindItemByID(itemID)
}

func (i ItemsServiceImpl) InsertItem(newItem *models.CreateItemInput) bool {
	return i.itemsRepo.InsertItem(newItem)
}

func (i ItemsServiceImpl) UpdateItem(itemID int64, update *models.UpdateItemInput) bool {
	return i.itemsRepo.UpdateItem(itemID, update)
}

func (i ItemsServiceImpl) DeleteItem(itemID int64) bool {
	return i.itemsRepo.DeleteItem(itemID)
}



