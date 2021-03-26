package repository

//e-wallet-simple-api

import (
	"GinAPI/models"
	"github.com/stretchr/testify/mock"
)
// mock itemsrepository, for testing
type ItemsRepoMock struct {
	mock.Mock
}
//i have not yet understand how this works
func (m *ItemsRepoMock) FindAllItem() (*[]models.Items, error) {

	ret := m.Called()

	var mItems []models.Items
	if rf, ok := ret.Get(0).(func() []models.Items); ok {
		mItems = rf()
	} else {
		if ret.Get(0) != nil {
			mItems = ret.Get(0).([]models.Items)
		}
	}

	var r2 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(1)
	}

	return &mItems, r2
}

func (m *ItemsRepoMock) FindItemByID(itemID int64) (*models.Items, error) {
	ret := m.Called(itemID)
	
	var mItem *models.Items

	if ret.Get(0) != nil {
		mItem = ret.Get(0).(*models.Items) //assert type
	}
	
	return mItem, ret.Error(1)
}

func (m *ItemsRepoMock) InsertItem(newItem *models.CreateItemInput)  bool {
	ret := m.Called(newItem)
	
	err := ret.Error(0)
	if err != nil {
		return false
	}
	return true
}

func (m *ItemsRepoMock) UpdateItem(id int64, updateItem *models.UpdateItemInput) bool {
	ret := m.Called(id, updateItem)
	err := ret.Error(0)
	if err != nil {
		return false
	}
	return true
}

func (m *ItemsRepoMock) DeleteItem(id int64) bool {
	ret := m.Called(id)
	err := ret.Error(0)
	if err != nil {
		return false
	}
	return true
}



	

