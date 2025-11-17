package controller

import (
	"fmt"
	"log"
	"strconv"

	m "ex/model"
	valid "ex/validator"

	v "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Factorial(c *fiber.Ctx) error {
	n := c.Params("num")
	log.Printf("num: %s", n)

	num, err := strconv.Atoi(n)

	e := fiber.StatusBadRequest
	if num < 0 {
		return c.Status(e).JSON(fiber.Map{
			"Message": "Please input number > 0",
			"Data":    num,
		})
	}
	if err != nil {
		return c.Status(e).JSON(fiber.Map{
			"Message": "Invalid Input",
			"Data":    n,
		})
	}

	result := 1
	for i := 1; i < num+1; i++ {
		log.Println(result)
		result = result * i
	}
	log.Println(result)
	return c.JSON((fiber.Map{"Message": "success", "factorial": result}))
}

func ToAscii(c *fiber.Ctx) error {

	name := c.Query("tax_id")
	log.Println("name:" + name)

	a := []byte(name)
	result := ""
	for i := range a {
		result += fmt.Sprintf("%d", a[i])
		if i < len(a)-1 {
			result += " "
		}
	}

	return c.JSON(fiber.Map{"Message": "success", "AscII": result})
}

func CreateUser(c *fiber.Ctx) error {
	p := new(m.Users)

	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Invalid request", "error": err.Error()})
	}

	if err := valid.Validate.Struct(p); err != nil {
		r := make(map[string]string)
		for _, e := range err.(v.ValidationErrors) {
			r[e.Field()] = e.Error()
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Invalid request", "Errors": r})
	}

	log.Println(p)
	return c.JSON(fiber.Map{"Message": "success", "User": p})

}
