package controllers

import (
	"GinAPI/models"
	"GinAPI/util"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//insert, update should return the resulting row, so queries should add RETURNING [value] suffix


func ListAllItems(c *gin.Context) {
	//limit item list to 20 items to prepare pagination
	stmt, err := models.DB.Prepare("SELECT * FROM items limit 20")
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}
	//as long as there is an open result set (rows)
	//the underlying connection will be busy so defer close
	defer rows.Close()

	//var wg sync.WaitGroup
	listitems := make([]*models.Items, 0)
	for rows.Next() {
		item := new(models.Items)
			err := rows.Scan(&item.ID, &item.Name, &item.Price)
			if err != nil {
				c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
				return
			}
			listitems = append(listitems, item)
		//wg.Add(1)
		//go func() {
		//	defer wg.Done()
		//	item := new(models.Items)
		//	err := rows.Scan(&item.ID, &item.Name, &item.Price)
		//	if err != nil {
		//		c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		//		return
		//	}
		//	listitems = append(listitems, item)
		//}()
	}
	//wg.Wait()
	c.JSON(http.StatusOK, gin.H{"data": listitems})
}

//tested, works fine
func findItem(c *gin.Context) (models.Items, error) {
	var item models.Items
	id := util.GetInt64IdFromContext(c)
	stmt, err := models.DB.Prepare("SELECT * FROM items WHERE id = $1")
	if err != nil {
		//i shouldnt use this
		log.Fatal(err)
	}
	err = stmt.QueryRow(id).Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		return item, err
	}
	return item, nil
}

//same as findItem, works fine
func FindItemById(c *gin.Context) {
	var finditem models.Items
	finditem, err := findItem(c); if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, errorResponse("Data not found"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": finditem})
}

func InsertItem(c *gin.Context) {
	var insert models.CreateItemInput
	if err := c.ShouldBindJSON(&insert); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		log.Fatal(err)
	}
	insertStmt,err := models.DB.Prepare("insert into items(name, price) values($1, $2) RETURNING id, name, price")
	if err != nil {
		log.Fatal(err)
	}
	defer insertStmt.Close()
	//this will result in "Locker{}" if successful
	var result models.Items
	err = insertStmt.QueryRow(insert.Name, insert.Price).Scan(&result.ID, &result.Name, &result.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func EditItem(c *gin.Context) {

}

func PatchItem(c *gin.Context) {
	//find the item first
	var item models.Items
	item, err := findItem(c)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, errorResponse("Data not found"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
			return
		}
	}

	var input *models.UpdateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}
	//check whether one of request body field is nil
	if &input.Name == nil {
		input.Name = item.Name
	}
	if &input.Price == nil {
		input.Price = item.Price
	}
	updatestmt, err := models.DB.Prepare("UPDATE items SET name = $2, price = $3 WHERE id = $1 RETURNING *")
	if err != nil {
		log.Fatal(err)
	}
	defer updatestmt.Close()
	err = updatestmt.QueryRow(item.ID, &input.Name, &input.Price).Scan(&item.ID, &item.Name, &item.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": item})
	//how to patch on sql query?
	//i may be able to get the to-be-updated item from db first,
	//then i checked whether the request body contains one or all item's field
	//if all, then good, proceed normally
	//if there is only one field, then i can compensate the missing field
	//by using the to-be-updated field, like
	//JSON -> {"name": "Bean"}, then the price shall be item.price
}

//untested
func DeleteItem(c *gin.Context) {
	//find the to-be-deleted data's id
	id := util.GetInt64IdFromContext(c)
	stmt, err := models.DB.Prepare("DELETE FROM items WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	//we dont want to get any result, so use exec
	_, err = stmt.Exec(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, errorResponse("Data not found"))
			return
		} else {
			c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}

//standard response given to user whenever there is an error
//i should create consts for error
func errorResponse(err string) gin.H {
	return gin.H{"error": err}
}