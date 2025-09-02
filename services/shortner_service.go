// Package services provides - wrapper around the shortner util
package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iamveekthorr/models"
	"github.com/iamveekthorr/utils"
)

type createURLShornerBody struct {
	URL string `json:"url" binding:"required"`
}

type updateURL struct {
	ShortCode string `json:"shortCode" binding:"required"`
	URL       string `json:"url" binding:"required"`
}

type response struct {
	URL        string    `json:"url"`
	ShortCode  string    `json:"shortCode"`
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	ClickCount int       `json:"clickCount"`
}

type uriParam struct {
	ShortCode string `uri:"shortcode" binding:"required"`
}

func CreateURLShorner(c *gin.Context) {
	var body createURLShornerBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse json",
			"error":   err.Error(),
		})
		return
	}

	var shortcode string

	for {
		shortcode = utils.MakeShortCode(7)

		// Each request is already in a goroutine
		rows, err := models.ConnPool.Query(
			c.Request.Context(),
			"SELECT * FROM urls WHERE short_code = $1",
			shortcode,
		)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{
				"message": "Unable to fetch rows",
				"error":   err.Error(),
			})
		}

		if !rows.Next() {
			break
		}
	}

	shortcode = utils.MakeShortCode(7)

	var data response
	// insert when truly unique
	row := models.ConnPool.QueryRow(
		c.Request.Context(),
		`INSERT INTO urls (url, short_code) VALUES ($1, $2) RETURNING id, url, short_code, created_at, updated_at`,
		body.URL, shortcode)

	err := row.Scan(&data.ID, &data.URL, &data.ShortCode, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error inserting into table!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id":        data.ID,
			"shortcode": data.ShortCode,
			"url":       data.URL,
			"createdAt": data.CreatedAt,
			"updatedAt": data.UpdatedAt,
		},
		"message": "Short code generated!",
	})
}

func UpdateURL(c *gin.Context) {
	var body updateURL

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse json",
			"error":   err.Error(),
		})
		return
	}

	query := `UPDATE urls SET url=$2 WHERE short_code=$1 RETURNING id, url, short_code, created_at, updated_at`

	// find url in db using the id
	var data response
	row := models.ConnPool.QueryRow(c.Request.Context(), query, body.ShortCode, body.URL)

	err := row.Scan(&data.ID, &data.URL, &data.ShortCode, &data.CreatedAt, &data.UpdatedAt)
	// throws error if not found!
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found!",
			"error":   err.Error(),
		})
		return
	}

	// return updated body
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":        data.ID,
			"shortcode": data.ShortCode,
			"url":       data.URL,
			"createdAt": data.CreatedAt,
			"updatedAt": data.UpdatedAt,
		},
		"message": "Record updated successfully!",
	})
}

func GetURLShortCode(c *gin.Context) {
	var uri uriParam

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse uri",
			"error":   err.Error(),
		})
		return
	}

	query := `SELECT * FROM urls WHERE short_code=$1;`

	// find url in db using the id
	var data response
	err := models.ConnPool.QueryRow(c.Request.Context(), query, uri.ShortCode).Scan(&data.ID, &data.URL, &data.ShortCode, &data.CreatedAt, &data.UpdatedAt, &data.ClickCount)
	// throws error if not found!
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found!",
			"error":   err.Error(),
		})
		return
	}

	// return updated body
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":          data.ID,
			"shortcode":   data.ShortCode,
			"url":         data.URL,
			"accessCount": data.ClickCount,
			"createdAt":   data.CreatedAt,
			"updatedAt":   data.UpdatedAt,
		},
		"message": "Record retrieved successfully!",
	})
}

func DeleteShortCode(c *gin.Context) {
	var uri uriParam

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse uri",
			"error":   err.Error(),
		})
		return
	}

	query := `DELETE FROM urls WHERE short_code=$1`

	// find url in db using the id
	result, err := models.ConnPool.Exec(c.Request.Context(), query, uri.ShortCode)
	// throws error if not found!
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting record",
			"error":   err.Error(),
		})
		return
	}

	if result.RowsAffected() == 0 {
		// return updated body
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Record not found!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Record deleted successfully!",
	})
}

func HandleRedirect(c *gin.Context) {
	var uri uriParam

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse uri",
			"error":   err.Error(),
		})
		return
	}

	var originalURL string
	// find url in db using the shorcode
	err := models.ConnPool.QueryRow(c.Request.Context(), `SELECT url FROM urls WHERE short_code=$1`, uri.ShortCode).Scan(&originalURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Short code not found"})
		return
	}

	// Update click count (Exec)
	query := `UPDATE urls SET click_count = click_count + 1 WHERE short_code=$1`

	_, err = models.ConnPool.Exec(c.Request.Context(), query, uri.ShortCode)
	// throws error if not found!
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error occurred while updating record!",
			"error":   err.Error(),
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}
