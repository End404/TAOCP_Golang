package main

import (
	cfg "filestore-server/config"
	"filestore-server/route"
)

func main() {
	/*
		//文件存取接口
		http.HandleFunc("/file/upload", handler.UploadHandler)
		http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
		http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
		http.HandleFunc("/file/download", handler.DownloadHandler)
		http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
		http.HandleFunc("/file/delete", handler.FileDeleteHandler)
		//秒传接口
		http.HandleFunc("/file/fastupload", handler.HTTPInterceptor(
			handler.TryFastUploadHandler))

		http.HandleFunc("/file/downloadurl", handler.HTTPInterceptor(
			handler.DownloadURLHandler))

		// 分块上传接口
		http.HandleFunc("file/mpupload/init", //初始化分块信息
			handler.HTTPInterceptor(handler.InitialMultipartUploadHandler))
		http.HandleFunc("/file/mpupload/uppart", //上传分块
			handler.HTTPInterceptor(handler.UploadPartHandler))
		http.HandleFunc("/file/mpupload/complete", //通知分块上传完成
			handler.HTTPInterceptor(handler.CompleteUploadHandler))

		//用户相关接口
		http.HandleFunc("/user/signup", handler.SignupHandler)
		http.HandleFunc("/user/signin", handler.SignupHandler)
		http.HandleFunc("/user/info", handler.HTTPInterceptor(handler.UserInfoHandler))

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Printf("Failed to start server, err:%s", err.Error())
		}
	*/

	//gin framework
	route := route.Router()
	route.Run(cfg.UploadSucHandler)

}
