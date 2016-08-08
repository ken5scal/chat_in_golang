package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"path/filepath"
)

func uploaderHandler(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userid") // read from hidden field in html
	file, header, err := req.FormFile("avatarFile") // retrieve reader for reading uploaded bytes
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	filename := filepath.Join("avatars", userId + filepath.Ext(header.Filename))
	err = ioutil.WriteFile(filename, data, 0777)
	if err !=nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, "Success")
}
