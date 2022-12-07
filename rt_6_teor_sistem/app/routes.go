package app

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	app.Get("/", index) // Главная страница

	app.Get("/stool/create", createstoolGet) // Страница создания нового животного
	app.Post("/stool/create", createstool)   // Добавление нового животного

	app.Get("/stools/:id?", readstool) // Страница вывода данных о множестве животных, либо одном с сохранением шаблона

	app.Get("/stool/update/:id", updatestoolGet) // Страница вывода данных для обновления данных о животном
	app.Post("/stool/update/:id", updatestool)   // Обновление данных о животном

	app.Get("/stool/delete/:id", deletestool) // Страница, которая вызывается для удаления данных о животном
}
