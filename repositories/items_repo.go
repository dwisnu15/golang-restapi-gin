package repositories

import "GinAPI/models"

type ItemsRepo interface {
	FindAllItem(params models.ListItemsParams) (*[]models.Items, error)
	FindItemByID(itemID int64) (*models.Items, error)
	InsertItem(newItem *models.CreateItemInput) bool
	UpdateItem(itemID int64, update *models.UpdateItemInput) bool
	DeleteItem(itemID int64) bool

}
