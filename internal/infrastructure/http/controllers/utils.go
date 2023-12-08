package controllers

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetCustomerReviewIDFromCookies(c *fiber.Ctx) uint {
	customerReviewID := c.Cookies("customer-review-id")
	parsedID, _ := strconv.ParseUint(customerReviewID, 10, 32)
	uintID := uint(parsedID)
	return uintID
}

func GetIdFromCookies(c *fiber.Ctx) uint {
	id := c.Cookies("id")
	parsedID, _ := strconv.ParseUint(id, 10, 32)
	uintID := uint(parsedID)
	return uintID
}
