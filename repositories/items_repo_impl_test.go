package repositories_test

import (
	"GinAPI/models"
	itemsRepoTest "GinAPI/repositories"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)
//create mock item object
var i = models.Items{
	ID: 32,
	Name: "Corn",
	Price: 3000,
}

/*
	i should create a Before Each equivalent in go
	so prepareMock() doesn't have to be called explicitly
	in every func
 */

func prepareMock() (*sql.DB, sqlmock.Sqlmock) {
	//create mock database
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("error when creating mock db :%s", err)
	}

	return db, mock
}

//test findbyid on items
func TestItemsRepoImpl_FindItemByID(t *testing.T) {
	db, mock := prepareMock()
	rows := sqlmock.NewRows([]string {"id", "name", "price"}).
		AddRow(i.ID, i.Name, i.Price)

	queryFindId := "SELECT * FROM items WHERE id=$1"

	//mock the query call with ID from 'i'
	mock.ExpectQuery(queryFindId).WithArgs(i.ID).WillReturnRows(rows)
	defer db.Close()

	repo := itemsRepoTest.InitItemRepo(db)
	item, err := repo.FindItemByID(i.ID)

	assert.NotNil(t, item)
	assert.NoError(t, err)
}

func TestItemsRepoImpl_FindAllItem(t *testing.T) {
	db, mock := prepareMock()
	rows := sqlmock.NewRows([]string {"id", "name", "price"}).
		AddRow(i.ID, i.Name, i.Price)

	queryFindId := "SELECT * FROM items"

	//mock the query call with ID from 'i'
	mock.ExpectQuery(queryFindId).WillReturnRows(rows)
	defer db.Close()

	repo := itemsRepoTest.InitItemRepo(db)
	items, err := repo.FindAllItem()

	assert.NotEmpty(t, items)
	assert.NoError(t, err)
}

func TestItemsRepoImpl_UpdateItem(t *testing.T) {
	db, mock := prepareMock()

	//update data for var i
	var updateData = models.UpdateItemInput{
		Name: "Flour",
		Price: 2000,
	}
	query := "UPDATE items SET name=$1, price=$3 WHERE id=$1"
	mock.ExpectQuery(query).WithArgs(i.ID, updateData.Name, updateData.Price)
	defer db.Close()

	repo := itemsRepoTest.InitItemRepo(db)

	updateSuccess:= repo.UpdateItem(i.ID, &updateData)
	assert.Equal(t, true, updateSuccess)
	assert.Equal(t, "Flour", i.Name)
	assert.Equal(t, 2000, i.Price)
}







