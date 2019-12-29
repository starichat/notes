package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const maxUploadSize = 500 * 1024 * 1024
const uploadPath = ""

func main() {
	// 设置上传路由
	http.HandleFunc("/upload", uploadFileHandler())

	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("files", fs))

	http.Handle("/hello", sayHello())
	log.Println("satrt http in hello")
	log.Println("Server startd on 8080, use /upload for upload")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sayHello() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		w.Header().Set("content-type", "application/json")
		w.Write([]byte("hello world"))

	})
}

// 对应的上传功能
func uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		fileType := r.PostFormValue("type")
		log.Println("filetypr:", fileType)
		file, _, err := r.FormFile("uploadFile")
		log.Println(file)
		if err != nil {
			//log.Println(err)
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println("file is invild")
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		filetype := http.DetectContentType(fileBytes)
		if filetype != "image/jpeg" && filetype != "image/jpg" &&
			filetype != "image/gif" && filetype != "image/png" &&
			filetype != "application/pdf" {
			log.Println("file type is invild")
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}
		fileName := randToken(12)
		//fileEndings, err := mime.ExtensionsByType(fileType)
		if err != nil {
			renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
			return
		}
		newPath := filepath.Join(uploadPath, fileName+".pdf")
		fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)
		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			log.Println(err)
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("SUCCESS"))
	})
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
