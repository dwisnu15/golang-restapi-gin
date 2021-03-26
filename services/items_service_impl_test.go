package services

import (
	mockRepo "GinAPI/mocks/repository"
	"GinAPI/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

//for finditem with id
var mockItem = models.Items{
	ID: 51,
	Name: "Sago",
	Price: 1000,
}
//for insertitem
var mockItemInsert = models.CreateItemInput{
	Name: "Noodle",
	Price: 3500,
}
//for update item
var mockItemUpdate = models.UpdateItemInput{
	Name: "Cooking Oil",
	Price: 5500,
}

//how to create beforeEach equivalent in Go?
//so i dont have to create mock of items repository
//func TestCreateItemsService(t *testing.T) {
//
//}

func TestItemsServiceImpl_DeleteItem(t *testing.T) {
	t.Run("Update item test, should return true if successful", func(t *testing.T) {
		mockItemRepo := new(mockRepo.ItemsRepoMock)
		mockItemRepo.On("UpdateItem", mockItem.ID, &mockItemUpdate).Return(nil)

		mockService := CreateItemsService(mockItemRepo)
		successUpdate := mockService.UpdateItem(mockItem.ID, &mockItemUpdate)

		assert.Equal(t, true, successUpdate)
	})
}

func TestItemsServiceImpl_FindAllItem(t *testing.T) {
	t.Run("Find all item with limit 20, should return an array of Items if successful", func(t *testing.T) {
		mockItemRepo := new(mockRepo.ItemsRepoMock)
		var mockAllItem []models.Items
		// going to insert mockItem from above
		//so at least findall will return one Item
		mockAllItem = append(mockAllItem, mockItem)
		mockItemRepo.On("FindAllItem").Return(mockAllItem, nil)
		mockService := CreateItemsService(mockItemRepo)
		result, err := mockService.FindAllItem()

		assert.NotNil(t, result)
		assert.Equal(t, len(mockAllItem), len(*result))
		assert.Nil(t, err)
	})
}

func TestItemsServiceImpl_FindItemByID(t *testing.T) {
		t.Run("Should return item with id: 51", func(t *testing.T) {

			mockItemsRepo := new(mockRepo.ItemsRepoMock)

			mockItemsRepo.On("FindItemByID", mockItem.ID).Return(&mockItem, nil)
			itemsService := CreateItemsService(mockItemsRepo)
			result, err := itemsService.FindItemByID(mockItem.ID)

			assert.NotNil(t, result)
			assert.Equal(t, mockItem.ID, result.ID)
			assert.Nil(t, err)
		})
}

func TestItemsServiceImpl_InsertItem(t *testing.T) {
	t.Run("Insert item test, should return true if successful", func(t *testing.T) {
		//db, mockSql, err := sqlmock.New()
		//assert.Nil(t, err)
		//defer db.Close()
		//mockSql.ExpectBegin()

		mockItemRepo := new(mockRepo.ItemsRepoMock)
		mockItemRepo.On("InsertItem", &mockItemInsert).Return(nil)

		mockService := CreateItemsService(mockItemRepo)
		successInsert := mockService.InsertItem(&mockItemInsert)

		assert.Equal(t, true, successInsert)
	})
}

func TestItemsServiceImpl_UpdateItem(t *testing.T) {
	t.Run("Update item test, should return true if successful", func(t *testing.T) {
		mockItemRepo := new(mockRepo.ItemsRepoMock)
		mockItemRepo.On("UpdateItem", mockItem.ID, &mockItemUpdate).Return(nil)

		mockService := CreateItemsService(mockItemRepo)
		successUpdate := mockService.UpdateItem(mockItem.ID, &mockItemUpdate)

		assert.Equal(t, true, successUpdate)
	})
}