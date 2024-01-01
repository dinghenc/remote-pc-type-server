package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/dinghenc/remote-pc-type-server/model"
	"github.com/dinghenc/remote-pc-type-server/robot"
	"github.com/gin-gonic/gin"
)

type Server struct {
	op robot.Operator
	mu sync.Mutex
}

func (s *Server) OnChangeText(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	req := model.Request{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &model.Response{
			RetCode: -1,
			ErrInfo: "bind request failed",
		})
		log.Println("bind request failed")
		return
	}
	log.Printf("request: %+v", req)

	if err := s.HandleOnChangeText(req); err != nil {
		c.JSON(http.StatusInternalServerError, &model.Response{
			RetCode: -2,
			ErrInfo: fmt.Sprintf("handle on change text failed: %v", err),
		})
		log.Printf("handle on change text failed: %v", err)
		return
	}
	c.JSON(http.StatusOK, &model.Response{})
	log.Printf("on change text suc: %+v", req)
}

func (s *Server) OnSearch(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	req := model.Request{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &model.Response{
			RetCode: -1,
			ErrInfo: "bind request failed",
		})
		log.Println("bind request failed")
		return
	}
	log.Printf("request: %+v", req)

	if err := s.HandleOnSearch(req); err != nil {
		c.JSON(http.StatusInternalServerError, &model.Response{
			RetCode: -3,
			ErrInfo: fmt.Sprintf("handle on search failed: %v", err),
		})
		log.Printf("handle on search failed: %v", err)
		return
	}
	c.JSON(http.StatusOK, &model.Response{ErrInfo: req.Text})
	log.Printf("on search suc: %+v", req)
}

func (s *Server) HandleOnChangeText(request model.Request) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.op.PasteString(request.Text); err != nil {
		return fmt.Errorf("paste string failed: %w", err)
	}

	return nil
}

func (s *Server) HandleOnSearch(request model.Request) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.op.PasteString(request.Text); err != nil {
		return fmt.Errorf("paste string failed: %w", err)
	}
	if err := s.op.Enter(); err != nil {
		return fmt.Errorf("enter failed: %w", err)
	}

	return nil
}
