package main

import (
	"src/cmd/main/src/pkg"
	"src/cmd/main/src/pkg/api"
	"src/cmd/main/src/pkg/socket"
	"src/cmd/main/src/pkg/usb"
)

func main() {
	server := pkg.Server{
		UsbService:    *usb.NewUsbService(),
		ApiService:    *api.NewApiService(),
		SocketService: *socket.NewSocketService(),
	}
	server.Run()
}
