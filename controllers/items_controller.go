package controllers

import (
	"GinAPI/constants"
	"GinAPI/middlewares"
	"GinAPI/models"
	service "GinAPI/services"
	"GinAPI/util"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type itemsController struct {
	service service.ItemsService
	router  *gin.Engine
	//tokenGen
}

func (itemsContr *itemsController) SetupRouter() (*gin.Engine, error) {
	//if another entity exists, change this router instance
	r := gin.Default()
	itemsContr.router = r
	itemsContr.router.Use(middlewares.LoggerToFile())

	api := itemsContr.router.Group("/api")
	api.GET("/items", itemsContr.listAllItems)          //tested
	api.GET("/items/:id", itemsContr.findItemById)      //tested
	api.POST("/items", itemsContr.insertItem)           //tested
	api.DELETE("/items/:id", itemsContr.deleteItemById) //tested
	api.PATCH("/items/:id", itemsContr.patchItem)       //tested
	return r, nil
}

func InitItemsController(itemsService service.ItemsService) (*itemsController, error) {
	controller := &itemsController{
		service: itemsService,
	}
	//controller.SetupRouter()
	return controller, nil
}

func (itemsContr *itemsController) listAllItems(c *gin.Context) {
	var req models.ListItemsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		util.HandleFailure(c, http.StatusBadRequest, constants.SubmitDataErr)
		return
	}

	arg := models.ListItemsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	listItems, err := itemsContr.service.FindAllItem(arg)
	if err != nil {
		util.HandleFailure(c, http.StatusInternalServerError, constants.ServerError)
		return
	}

	util.HandleSuccessWithData(c, listItems)
}

//tested, works fine
//this will be used on FindById, PatchItem, and DeleteItem
func (itemsContr *itemsController) findItem(c *gin.Context) (*models.Items, error) {
	var item *models.Items
	id := util.GetInt64IdFromContext(c)
	item, err := itemsContr.service.FindItemByID(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

//same as findItem, works fine
func (itemsContr *itemsController) findItemById(c *gin.Context) {
	finditem, err := itemsContr.findItem(c)
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

func (itemsContr *itemsController) insertItem(c *gin.Context) {
	var insert *models.CreateItemInput

	//body, _ := ioutil.ReadAll(c.Request.Body)
	//_ = json.Unmarshal(body, &insert)
	if err := c.ShouldBindJSON(&insert); err != nil {
		util.HandleFailure(c, http.StatusBadRequest, constants.RequiredFields)
		return
	}
	insertSuccess := itemsContr.service.InsertItem(insert)
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
func (itemsContr *itemsController) patchItem(c *gin.Context) {
	//find the item first
	var item *models.Items
	item, err := itemsContr.findItem(c)
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
	if input.Name == "" {
		log.Info("entered input name checking")
		input.Name = item.Name
	}
	if input.Price == 0 {
		log.Info("entered input price checking")
		input.Price = item.Price
	}
	//this will probably never return "missing fields" error
	//because of the field checking above
	result := itemsContr.service.UpdateItem(item.ID, input)
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

func (itemsContr *itemsController) deleteItemById(c *gin.Context) {
	//find the to-be-deleted data's id
	id := util.GetInt64IdFromContext(c)
	result := itemsContr.service.DeleteItem(id)
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
