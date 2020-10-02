package vojo

type DownloadFileReq struct {
	FileKeyCode string `form:"fileKeyCode" json:"fileKeyCode" `
	FileName    string `form:"fileName" json:"fileName" `
}
