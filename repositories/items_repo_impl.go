package repositories

import (
	"GinAPI/constants"
	"GinAPI/models"
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"

	"runtime"
	"strings"
)

type ItemsRepoImpl struct {
	db *sql.DB
}

func InitItemRepo(db *sql.DB) ItemsRepo {
	return &ItemsRepoImpl{db}
}

func (i ItemsRepoImpl) FindAllItem(arg models.ListItemsParams) (*[]models.Items, error) {
	//limit item list to 20 items to prepare pagination
	stmt, err := i.db.Prepare("SELECT * FROM items LIMIT $1 OFFSET $2")
	if err != nil {
		log.Error(errorResponse(err.Error()))
		return nil, fmt.Errorf(constants.ServerError)
	}

	rows, err := stmt.Query(arg.Limit, arg.Offset)
	if err != nil {
		log.Error(errorResponse(err.Error()))
		return nil, fmt.Errorf(constants.ServerError)
	}

	var listitems []models.Items
	for rows.Next() {
		var item models.Items
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Price)
		if err != nil {
			log.Error(errorResponse(err.Error()))
			return nil, fmt.Errorf(constants.ServerError)
		}
		listitems = append(listitems, item)
	}
	//as long as there is an open result set (rows)
	//the underlying connection will be busy so defer close
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &listitems, nil
}

func (i ItemsRepoImpl) FindItemByID(itemID int64) (*models.Items, error) {

	stmt, err := i.db.Prepare("SELECT * FROM items WHERE id = $1")
	if err != nil {
		log.Error(errorResponse(err.Error()))
		return nil, fmt.Errorf(constants.ServerError)
	}
	result, err := stmt.Query(itemID)
	if err != nil {
		log.Error(errorResponse(err.Error()))
		return nil, fmt.Errorf(constants.ServerError)
	}
	defer result.Close()
	var item models.Items
	//i could use queryrow.scan, but...
	for result.Next() {
		err := result.Scan(
			&item.ID,
			&item.Name,
			&item.Price,
		)
		if err != nil {
			log.Error(errorResponse(err.Error()))
			return nil, fmt.Errorf(constants.ServerError)
		}
	}

	return &item, nil
}

func (i ItemsRepoImpl) InsertItem(newItem *models.CreateItemInput) bool {
	insertStmt, err := i.db.Prepare("INSERT INTO items(name, price) VALUES ($1, $2)")
	if err != nil {
		log.Error(errorResponse(err.Error()))
		return false
	}

	_, err = insertStmt.Exec(newItem.Name, newItem.Price)
	if err != nil {
		log.Error(errorResponse(err.Error()))
		return false
	}
	return true
}

//i cannot yet impl patch request so
//put request
func (i ItemsRepoImpl) UpdateItem(id int64, update *models.UpdateItemInput) bool {
	//mutex?

	stmt, err := i.db.Prepare("UPDATE items SET name = $2, price = $3 WHERE id = $1")
	if err != nil {
		log.Error(errorResponse(err.Error()))
		return false
	}
	_, err = stmt.Exec(id, update.Name, update.Price)
	if err != nil {
		log.Error(errorResponse(err.Error()))
		return false
	}
	return true
}

func (i ItemsRepoImpl) DeleteItem(itemID int64) bool {
	//find the to-be-deleted data's id
	stmt, err := i.db.Prepare("DELETE FROM items WHERE id = $1")
	if err != nil {
		log.Error(errorResponse(err.Error()))
		return false
	}

	defer stmt.Close()
	//we dont want to get any result, so use exec
	_, err = stmt.Exec(itemID)
	if err != nil {
		log.Error(errorResponse(err.Error()))
		return false
	}
	return true
}

//traverse the stack trace
func errorResponse(message string) string {
	fpcs := make([]uintptr, 1)
	//skip 2 levels to get the caller
	n := runtime.Callers(2, fpcs)
	if n == 0 {
		return "No caller"
	}

	caller := runtime.FuncForPC(fpcs[0] - 1)
	if caller == nil {
		return "Caller was nil"
	}
	var response strings.Builder
	response.WriteString(caller.Name())
	response.WriteString(message)
	return response.String()
}
