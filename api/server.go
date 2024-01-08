package api

import (
	"fmt"

	//"github.com/dortaedward/financeTracker/auth"
	"github.com/dortaedward/financeTracker/auth"
	"github.com/dortaedward/financeTracker/types"
	"github.com/gofiber/fiber/v2"
)

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
  err := c.JSON(data); if err != nil {
    return err
  }
	return nil
}

func CreateServer(addr string, server *fiber.App) *Server{
	return &Server{
		addr: addr,
		server: server,
	}
}

func HandleIndexRoute(c *fiber.Ctx) error {
  return nil
}

func HandleSignIn(c *fiber.Ctx) error{
  var signin_request types.SigninRequest
  if err := c.BodyParser(&signin_request); err != nil{
    c.Status(fiber.StatusBadRequest)
    return err
  }
  fmt.Println(signin_request.Email, signin_request.Password)
  auth.AuthenticateUser(signin_request)  
  return nil
}

func (s *Server) Run(){
	s.server.Get("/",HandleIndexRoute)
	routes := s.server.Group("/api")
	routes.Get("/",HandleHealthCheck)
  	routes.Post("/signin",HandleSignIn)
}
