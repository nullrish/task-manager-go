package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	IP   string
	Port string
	DB   *sql.DB
}

func InitializeServer(ip, port string, db *sql.DB) *Server {
	return &Server{IP: ip, Port: port, DB: db}
}

func (s *Server) StartServer() {
	addr := fmt.Sprintf("%s:%s", s.IP, s.Port)

	listenConfig := fiber.ListenConfig{
		DisableStartupMessage: true,
		EnablePrefork:         false,
	}
	app := fiber.New(fiber.Config{
		AppName: "Task Manager Go",
	})
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Khello World!")
	})

	log.Println("🚀 Starting Server....")
	log.Printf("Mode: %s\n", app.Config().AppName)
	log.Printf("Address: http://%s\n", addr)
	log.Printf("Prefork: %v\n", listenConfig.EnablePrefork)
	log.Printf("PID: %d\n", os.Getpid())

	log.Fatal(app.Listen(addr, listenConfig))
}
