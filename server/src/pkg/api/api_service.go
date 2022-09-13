package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"src/cmd/main/src/pkg/usb"
)

const HOST = ":8081"

type ApiService struct {
	usbService *usb.UsbService
}

func NewApiService() *ApiService {
	return &ApiService{}
}

func (r *ApiService) Active(usbService *usb.UsbService) {
	log.Println("api service active")
	log.Println("listen: " + HOST)
	r.usbService = usbService

	http.HandleFunc("/usb-list", r.UsbList)
	go func() {
		err := http.ListenAndServe(HOST, nil)
		if err != nil {
			panic(err)
		}
	}()
}

func (r *ApiService) UsbList(response http.ResponseWriter, request *http.Request) {
	marshal, _ := json.Marshal(r.usbService.UsbStruct)
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	response.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	response.Header().Set("Access-Control-Expose-Headers", "Authorization")
	fmt.Fprintf(response, "%s", marshal)
}
