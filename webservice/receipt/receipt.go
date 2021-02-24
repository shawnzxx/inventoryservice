package receipt

import (
	"io/ioutil"
	"path/filepath"
	"time"
)

//ReceiptDirectory path
var ReceiptDirectory string = filepath.Join("uploads")

//Receipt structure
type Receipt struct {
	ReceiptName string    `json:"name"`
	UploadDate  time.Time `json:"uploadDate"`
}

//GetReceipts will return list of receipt which structure have properties ReceiptName, UploadDate
func GetReceipts() ([]Receipt, error) {
	reciepts := make([]Receipt, 0)
	files, err := ioutil.ReadDir(ReceiptDirectory)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		reciepts = append(reciepts, Receipt{ReceiptName: f.Name(), UploadDate: f.ModTime()})
	}
	return reciepts, nil
}
