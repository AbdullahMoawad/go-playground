package uploader

import (
	"encoding/json"
	"fmt"
	"github.com/h2non/filetype"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"real-estate/models"
)

type image struct {
	name 	string
	size 	int64
	// kind is the pic or file part on sys ( profile ,post , else)
	kind 	string
	headers textproto.MIMEHeader
}
func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("profile")
	if err != nil {
		json.NewEncoder(w).Encode(models.Logger(404, "Error Retrieving the File",err))
		return
	}
	defer file.Close()
	
	fmt.Println(image{name:handler.Filename},"/n")
	fmt.Println(image{size:handler.Size},"/n")
	fmt.Println(image{headers:handler.Header},"/n")

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp/profile", "upload-*.png")
	if err != nil {
		json.NewEncoder(w).Encode(models.Logger(404, " ",err))
		return
	}

	defer tempFile.Close()

	// read all of the contents of the uploaded file

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		json.NewEncoder(w).Encode(models.Logger(404, "can't read file ",err))
		return
	}
	// check if the file you gonna upload is image or no't
	if !filetype.IsImage(fileBytes) {
		json.NewEncoder(w).Encode(models.Logger(406, "this file isn't an image (no't accepted)",nil))
		return
	}

	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "")
	json.NewEncoder(w).Encode("Successfully Uploaded File\n")

}