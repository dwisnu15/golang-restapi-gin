package models

//For list Items pagination
type ListItemsRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=20"`//maximum item per page = 20
}

//params sent by user (actually web but you know), will be used as settings for items per page
type ListItemsParams struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

