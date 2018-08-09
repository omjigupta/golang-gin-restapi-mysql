package controllers

import (
	"net/http"
  "fmt"
  "bytes"
  "golang-gin-restapi-mysql/db"
	"golang-gin-restapi-mysql/models"
	"github.com/gin-gonic/gin"
)

// Create new flower details
func Create(c *gin.Context) {
		var buffer bytes.Buffer
		name := c.PostForm("name")
		category := c.PostForm("category")
    price := c.PostForm("price")
    photo := c.PostForm("photo")
    descriptions := c.PostForm("descriptions")

		stmt, err := db.Init().Prepare("insert into flower (name, category, price, photo, descriptions) values(?,?,?,?,?);")

		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(name, category, price, photo, descriptions)

		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(name)
		buffer.WriteString(" ")
		defer stmt.Close()
		flowername := buffer.String()

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf(" %s successfully created", flowername),
		})
}

// GET all flowers
func GetAllFlower(c *gin.Context) {
		var (
			flower models.Flower
			flowers []models.Flower
		)
		rows, err := db.Init().Query("select id, name, category, price, photo, descriptions from flower;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&flower.Id, &flower.Name, &flower.Category, &flower.Price, &flower.Photo, &flower.Descriptions)
			flowers = append(flowers, flower)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": flowers,
			"count":  len(flowers),
		})
}
