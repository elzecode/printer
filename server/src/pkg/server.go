package pkg

import (
	"bufio"
	"os"
	"src/cmd/main/src/pkg/api"
	"src/cmd/main/src/pkg/socket"
	"src/cmd/main/src/pkg/usb"
)

type Server struct {
	UsbService    usb.UsbService
	ApiService    api.ApiService
	SocketService socket.SocketService
}

func (s *Server) Run() {
	s.SocketService.Active()
	s.ApiService.Active(&s.UsbService)
	s.UsbService.Active(&s.SocketService)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Нас ебут, а мы крепчаем
	}
}
