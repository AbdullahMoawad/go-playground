package uploader

import (
	"encoding/json"
	"fmt"
	"github.com/h2non/filetype"
	"io/ioutil"
	"log"
	"net/http"
	"net/textproto"
	"os"
	"real-estate/App"
	"real-estate/common"
)

type image struct {
	name string
	size int64
	// kind is the pic or file part on sys ( profile ,post , else)
	kind     string
	userid   string
	estateID string
	headers  textproto.MIMEHeader
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	kind := r.Header.Get("kind")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile(kind)
	if err != nil {

		App.JsonLogger(w, 404, "Error Retrieving the File", err)
		App.Logger("Error Retrieving the File", "error")
		return
	}
	defer file.Close()

	//fmt.Println(image{name:handler.Filename},"/n")
	//fmt.Println(image{size:handler.Size},"/n")
	//fmt.Println(image{headers:handler.Header},"/n")

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	dir := ""

	if kind == "profile" {
		sessId := r.Header.Get("sessionId")
		err, userId := common.GetCurrentUserIdFromHeaders(sessId)
		if err != nil {
			fmt.Println(err)
			return
		}

		dir += "temp/profile/" + userId + "/"
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.Mkdir(dir, 0777); err != nil {
				log.Println("failed to create test sub-directory:", err)
			}

		}

	} else if kind == "estate" {
		dir += "temp/estate/" + ""
	}

	tempFile, err := ioutil.TempFile(dir, "upload-*.png")
	if err != nil {

		App.JsonLogger(w, 500, "", err)
		App.Logger("", "error")
		return
	}
	defer tempFile.Close()

	// read all of the contents of the uploaded file

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		App.JsonLogger(w, 500, "can't read file", err)
		App.Logger("can't read file", "error")
		return
	}
	// check if the file you gonna upload is image or no't
	if !filetype.IsImage(fileBytes) {
		App.JsonLogger(w, 406, "this file isn't an image (no't accepted ", nil)
		App.Logger("this file isn't an image (no't accepted", "error")
		return
	}

	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	json.NewEncoder(w).Encode("Successfully Uploaded File\n")

}
