package core

import (
	"drako/internal/server"
	"drako/pkg/globals"
	"drako/pkg/logger"
)

type Core struct {
	m_pLogger *logger.Logger
	m_pServer *server.Server
}

func NewCore() *Core {
	return &Core{
		m_pLogger: logger.NewLogger(),
		m_pServer: server.NewServer(),
	}
}

func (c *Core) Start() {
	globals.Logger = c.m_pLogger

	c.m_pLogger.Start()
	c.m_pServer.Start()
}
