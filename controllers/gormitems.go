package controllers

//func FindAllItems(c *gin.Context) {
//	var allitems []models.Items
//	//limit 20 item
//	models.GormDB.Find(&allitems).Limit(10) //select * from items
//
//	c.JSON(http.StatusOK, gin.H{"data": allitems})
//}
//
//func findItem(c *gin.Context) (models.Items, error) {
//	var item models.Items
//	if err := models.GormDB.Where("id = ?", c.Param("id")).First(&item).Error;
//		err != nil {
//		return item, err
//	}
//	return item, nil
//}
//
//func FindItemById(c *gin.Context) {
//	var item models.Items
//	item, err := findItem(c); if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found"})
//		return
//	}
//
//	log.Info("Success finding item with id:", item.ID)
//	c.JSON(http.StatusOK, gin.H{"data": item})
//}
//
//func AddItemsInput(c *gin.Context) {
//	//validate input
//	var input models.CreateItemInput
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	//create item
//	item := models.Items{Name: input.Name, Price: input.Price}
//	models.GormDB.Create(&item)
//
//	c.JSON(http.StatusOK, gin.H{"data": item})
//}
//
//func UpdateItem(c *gin.Context) {
//	//first, get item if it exists
//	var item models.Items
//	item, err := findItem(c); if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found"})
//		return
//	}
//
//	//validate update input
//	var input models.UpdateItemInput
//	if err := c.ShouldBindJSON(&input);
//		err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	models.GormDB.Model(&item).Updates(input)
//	c.JSON(http.StatusOK, gin.H{"data": item})
//}
//
//func DeleteItem(c *gin.Context) {
//	//find the item first
//	var item models.Items
//	item, err := findItem(c); if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found"})
//		return
//	}
//
//	models.GormDB.Delete(&item)
//	log.Info("Deleting item with id:", item.ID)
//	c.JSON(http.StatusOK, gin.H{"data": true})
//}
