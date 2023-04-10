package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

type Employee struct {
	EmployeeID    int32  `json:"employeeID"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	UserName      string `json:"userName"`
	Password      string `json:"password"`
	ContactNumber string `json:"contactNumber"`
	Description   string `json:"description"`
}

func sendHTTPRequest(url string, c *fiber.Ctx) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError).Send("Error making request")
		return nil
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError).Send("Error occurred parsing this response")
		return nil
	}
	return response
}

func getEmployees(url string, c *fiber.Ctx) {
	employees := sendHTTPRequest(url, c)
	if employees == nil {
		return
	}

	c.Send(employees)
}

func getAllSupervisors(url string, c *fiber.Ctx) error {
	response := sendHTTPRequest(url, c)
	if response == nil {
		return fmt.Errorf("Error sending HTTP request")
	}

	var employees []Employee
	err := json.Unmarshal(response, &employees)
	if err != nil {
		return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	supervisors := make([]Employee, 0)

	for _, k := range employees {
		if k.Description == "SUPERVISOR" {
			supervisors = append(supervisors, k)
		}
	}

	jsonBytes, err := json.Marshal(supervisors)
	if err != nil {
		return fmt.Errorf("Error marshalling JSON: %v", err)
	}

	c.Send(jsonBytes)
	return nil
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("error accessing url")
	}

	employee_url := os.Getenv("employee_url")

	app := fiber.New()

	app.Get("/allEmployees", func(c *fiber.Ctx) {
		getEmployees(employee_url, c)
	})

	app.Get("/allSupervisors", func(c *fiber.Ctx) {
		getAllSupervisors(employee_url, c)
	})

	app.Listen(3000)
}

func getEmployeeByID(url string, id int32, c *fiber.Ctx) error {
	response := sendHTTPRequest(url, c)
	if response == nil {
		return fmt.Errorf("Error sending HTTP request")
	}

	var employees []Employee
	err := json.Unmarshal(response, &employees)
	if err != nil {
		return fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	for _, k := range employees {
		if k.EmployeeID == id {
			c.Send(json.Marshal(k))
			return nil
		}
	}
	return emplo
}
