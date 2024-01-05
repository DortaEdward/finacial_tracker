package types

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type Server struct{
	addr string
	server *fiber.App
  db *sql.DB
}

func HandleHealthCheck(c *fiber.Ctx) error {
	data := map[string]interface{}{
		"message": "Health Check",
		"status": fiber.StatusOK,
	}
	c.Status(fiber.StatusOK)
  err := c.JSON(data); if err != nil {
    return err
  }
	return nil
}

func CreateServer(addr string, server *fiber.App, db *sql.DB) *Server{
	return &Server{
		addr: addr,
		server: server,
    db: db,
	}
}

func HandleIndexRoute(c *fiber.Ctx) error {
  return nil
}

func (s *Server) Run(){
	s.server.Get("/",HandleIndexRoute)
	routes := s.server.Group("/api")
	routes.Get("/",HandleHealthCheck)

}
