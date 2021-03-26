package controllers

import (
	"GinAPI/services"
	"reflect"
	"testing"
)

func TestInitItemsController(t *testing.T) {
	type args struct {
		r            *gin.Engine
		itemsService services.ItemsService
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestItemsController_deleteItemById(t *testing.T) {
	type fields struct {
		itemsService services.ItemsService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ItemsController{
				itemsService: tt.fields.itemsService,
			}
		})
	}
}

func TestItemsController_findItem(t *testing.T) {
	type fields struct {
		itemsService services.ItemsService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Items
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ItemsController{
				itemsService: tt.fields.itemsService,
			}
			got, err := e.findItem(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("findItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemsController_findItemById(t *testing.T) {
	type fields struct {
		itemsService services.ItemsService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ItemsController{
				itemsService: tt.fields.itemsService,
			}
		})
	}
}

func TestItemsController_insertItem(t *testing.T) {
	type fields struct {
		itemsService services.ItemsService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ItemsController{
				itemsService: tt.fields.itemsService,
			}
		})
	}
}

func TestItemsController_listAllItems(t *testing.T) {
	type fields struct {
		itemsService services.ItemsService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ItemsController{
				itemsService: tt.fields.itemsService,
			}
		})
	}
}

func TestItemsController_patchItem(t *testing.T) {
	type fields struct {
		itemsService services.ItemsService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ItemsController{
				itemsService: tt.fields.itemsService,
			}
		})
	}
}