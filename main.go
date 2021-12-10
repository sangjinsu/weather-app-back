package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sangjinsu/weather-app-back/tools"
	"log"
)

type Request struct {
	SearchText string `json:"searchText"`
}

type Response struct {
	PlaceName   string `json:"placeName"`
	Temperature int    `json:"temperature"`
	Humidity    int    `json:"humidity"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	IsDay       string `json:"isDay"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Post("/weather", func(ctx *fiber.Ctx) error {
		request := new(Request)
		if bodyParserErr := ctx.BodyParser(request); err != nil {
			return bodyParserErr
		}

		ctx.Status(fiber.StatusOK)
		latitude, longitude, placeName, locationErr := tools.FetchLocation(request.SearchText)
		if locationErr != nil {
			return locationErr
		}

		temperature, humidity, description, icon, isDay, weatherErr := tools.FetchWeather(latitude, longitude)
		if weatherErr != nil {
			return weatherErr
		}

		response := Response{
			PlaceName:   placeName,
			Temperature: temperature,
			Humidity:    humidity,
			Description: description,
			Icon:        icon,
			IsDay:       isDay,
		}

		return ctx.JSON(response)
	})

	serverErr := app.Listen(":3000")
	if serverErr != nil {
		log.Fatalf("Server Listening Error %v", serverErr)
		return
	}
}
