package controllers

import (
	"GinAPI/constants"
	"GinAPI/models"
	service "GinAPI/services"
	"GinAPI/util"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//insert, update should return the resulting row,
//so queries should add RETURNING [value] suffix
//..should it? currently use boolean to determine
//whether statement is successful or failed

type ItemsController struct {
	itemsService service.ItemsService
}

func InitItemsController(r *gin.Engine, itemsService service.ItemsService) {
	itemsController := ItemsController{itemsService}
	api := r.Group("/api")
	api.GET("/items", itemsController.listAllItems)//tested
	api.GET("/items/:id", itemsController.findItemById)//tested
	api.POST("/items", itemsController.insertItem)//tested
	api.DELETE("/items/:id", itemsController.deleteItemById)//tested
	api.PATCH("/items/:id", itemsController.patchItem)//tested
}

func (e *ItemsController) listAllItems(c *gin.Context) {
	var req models.ListItemsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		util.HandleFailure(c, http.StatusBadRequest, constants.SubmitDataErr)
		return
	}

	arg := models.ListItemsParams{
		Limit: req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	listItems, err := e.itemsService.FindAllItem(arg)
	if err != nil {
		util.HandleFailure(c, http.StatusInternalServerError, constants.ServerError)
		return
	}

	util.HandleSuccessWithData(c, listItems)

}

//tested, works fine
//this will be used on FindById, PatchItem, and DeleteItem
func (e *ItemsController) findItem(c *gin.Context) (*models.Items, error) {
	var item *models.Items
	id := util.GetInt64IdFromContext(c)
	item, err := e.itemsService.FindItemByID(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

//same as findItem, works fine
func (e *ItemsController) findItemById(c *gin.Context) {
	finditem, err := e.findItem(c)
	if err != nil {
		if err == sql.ErrNoRows {
			//hit api success BUT the entity is nonexistent, so we should return status 200
			util.HandleFailure(c, http.StatusOK, constants.DataNotFound)
			return
		} else {
			util.HandleFailure(c, http.StatusInternalServerError, constants.ServerError)
			return
		}
	}
	util.HandleSuccessWithData(c, finditem)
}

func (e *ItemsController) insertItem(c *gin.Context) {
	var insert *models.CreateItemInput

	//body, _ := ioutil.ReadAll(c.Request.Body)
	//_ = json.Unmarshal(body, &insert)
	if err := c.ShouldBindJSON(&insert); err != nil {
		util.HandleFailure(c, http.StatusBadRequest, constants.RequiredFields)
		return
	}
	insertSuccess := e.itemsService.InsertItem(insert)
	if !insertSuccess {
		util.HandleFailure(c, http.StatusInternalServerError, constants.ServerError)
		return
	}
	util.HandleSuccess(c, constants.SuccessInsert)
}

//func EditItem(c *gin.Context) {
//
//}

//really bad code (but works)
func (e *ItemsController) patchItem(c *gin.Context) {
	//find the item first
	var item *models.Items
	item, err := e.findItem(c)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, constants.DataNotFound)
			return
		} else {
			c.JSON(http.StatusInternalServerError, constants.ServerError)
			return
		}
	}

	//get update data from request body
	var input *models.UpdateItemInput
	if err := c.Bind(&input); err != nil {
		util.HandleFailure(c, http.StatusBadRequest, constants.SubmitDataErr)
		return
	}
	//check whether one of request body field is nil
	if input.Name == ""  {
		log.Info("entered input name checking")
		input.Name = item.Name
	}
	if input.Price == 0 {
		log.Info("entered input price checking")
		input.Price = item.Price
	}
	//this will probably never return "missing fields" error
	//because of the field checking above
	result := e.itemsService.UpdateItem(item.ID, input)
	if !result {
		util.HandleFailure(c, http.StatusBadRequest, constants.UpdateError)
		return
	}
	util.HandleSuccess(c, constants.SuccessUpdate)
	//how to patch on sql query?
	//i may be able to get the to-be-updated item from db first,
	//then i checked whether the request body contains one or all item's field
	//if all, then good, proceed normally
	//if there is only one field, then i can compensate the missing field
	//by using the to-be-updated field, like
	//JSON -> {"name": "Bean"}, then the price shall be item.price
}

func (e *ItemsController) deleteItemById(c *gin.Context) {
	//find the to-be-deleted data's id
	id := util.GetInt64IdFromContext(c)
	result:= e.itemsService.DeleteItem(id)
	if !result {
		util.HandleFailure(c, http.StatusInternalServerError, constants.ServerError)
		return
	}

	util.HandleSuccess(c, constants.SuccessDelete)
}

//standard response given to user whenever there is an error
//i should create consts for error (done)
//func errorResponse(errmsg string) gin.H {
//	return gin.H{"error": errmsg}
//}