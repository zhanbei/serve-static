package main

import (
	"fmt"
	"net/http"

	"github.com/zhanbei/serve-static"
)

var mStaticServer *servestatic.FileServer

const mAddress = "0.0.0.0:1234"
const mRootDir = "./static"
const mVirtualHost = false

func init() {
	server, err := servestatic.NewFileServer(mRootDir, mVirtualHost)
	if err != nil {
		panic(err)
	}
	mStaticServer = server
}

func main() {
	fmt.Println("Server is listening:", mAddress, "and looking after:", mRootDir, "; Using virtual host:", mVirtualHost, ".")
	// @see https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server
	http.ListenAndServe(mAddress, mStaticServer)
}
