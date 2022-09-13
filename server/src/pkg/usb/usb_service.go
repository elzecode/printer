package usb

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"src/cmd/main/src/pkg/helper"
	"src/cmd/main/src/pkg/socket"
	"strconv"
	"strings"
	"time"
)

const USB_PATH = "/var/server/tmp/usb_mount"

type UsbService struct {
	UsbInsert     bool
	UsbStruct     []File
	SocketService *socket.SocketService
	AllowedFiles  []string
	PdfDir        string
}

type File struct {
	Path     string `json:"path"`
	Filename string `json:"filename"`
	Ext      string `json:"ext"`
	Size     int64  `json:"size"`
	PdfPath  string `json:"pdf_path"`
}

func NewUsbService() *UsbService {
	thisDir, _ := os.Getwd()
	pdfDir := thisDir + "/pdf"
	return &UsbService{
		UsbInsert:    false,
		AllowedFiles: []string{".txt", ".doc", ".docx", ".xlsx"},
		PdfDir:       pdfDir,
	}
}

func (u *UsbService) Active(socketService *socket.SocketService) {
	u.SocketService = socketService
	log.Println("usb service active")
	go u.findUsb()
}

func (u *UsbService) findUsb() {
	for {
		dir, _ := os.ReadDir(USB_PATH)
		if len(dir) > 1 {
			u.Inserting(dir[0])
		} else {
			u.Conclusion()
		}
		time.Sleep(time.Second * 1)
	}
}

func (u *UsbService) Inserting(entry os.DirEntry) {
	if u.UsbInsert == false {
		u.UsbInsert = true
		log.Println("usb insert")
		u.SocketService.Say("usb_insert")
		files, err := u.Processing(USB_PATH + "/" + entry.Name())
		if err != nil {
			log.Println(err)
		} else {
			u.UsbStruct = *files
			u.SocketService.Say("usb_ready")
		}
	}
}

func (u *UsbService) Conclusion() {
	if u.UsbInsert == true {
		u.UsbInsert = false
		u.UsbStruct = make([]File, 0)
		u.SocketService.Say("usb_conclusion")
		log.Println("usb conclusion")
		err := os.RemoveAll(u.PdfDir)
		if err != nil {
			log.Println(err)
		}
	}
}

func (u *UsbService) Processing(usbPath string) (*[]File, error) {
	u.SocketService.Say("usb_processing")
	var files = make([]File, 0)

	var allowedPaths []string
	err := filepath.Walk(usbPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if (usbPath != path || !info.IsDir()) && helper.StringInSlice(filepath.Ext(path), u.AllowedFiles) {
			allowedPaths = append(allowedPaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	for index, path := range allowedPaths {
		stat, err := os.Stat(path)
		if err != nil {
			return nil, err
		}
		cmd := exec.Command("soffice",
			"--headless",
			"--convert-to",
			"pdf:writer_pdf_Export",
			"--outdir",
			u.PdfDir,
			path,
		)
		err = cmd.Run()
		if err != nil {
			return nil, err
		}
		files = append(files, File{
			Path:     strings.ReplaceAll(path, usbPath, ""),
			Filename: stat.Name(),
			Ext:      filepath.Ext(path),
			Size:     stat.Size(),
			PdfPath:  u.PdfDir + "/" + (strings.ReplaceAll(stat.Name(), filepath.Ext(path), ".pdf")),
		})
		u.SocketService.Say("processing:" + strconv.FormatInt(helper.PercentageChange(len(allowedPaths), index+1), 10))
	}

	return &files, nil
}
