package types

import "github.com/gofiber/fiber/v2"

type Server struct{
	addr string
	server *fiber.App
}

func HandleHealthCheck(c *fiber.Ctx) error {
	data := map[string]interface{}{
		"message": "Health Check",
		"status": fiber.StatusOK,
	}
	c.Status(fiber.StatusOK)
	c.JSON(data)
	return nil
}

func CreateServer(addr string, server *fiber.App) *Server{
	return &Server{
		addr: addr,
		server: server,
	}
}

func HandleIndexRoute(c *fiber.Ctx) error {
	return c.Render("index",fiber.Map{})
}

func (s *Server) Run(){
	s.server.Get("/",HandleIndexRoute)
	routes := s.server.Group("/api")
	routes.Get("/",HandleHealthCheck)

}
