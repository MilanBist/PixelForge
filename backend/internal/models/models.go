package models

import "time"


type ContextMetaData struct{
	UserId 		int
	Email 		string
}

type UploadedFilesMetaData struct{
	Id 			int64		`json:"uploadedId"`
	UserId		int64		`json:"userId,omitempty"`
	Filename	string		`json:"filename"`
	StorageKey	string		`json:"storageKey,omitempty"`
	FileType	string		`json:"fileType"`
	Mimetype	string		`json:"mimeType"`
	FileSize 	any		    `json:"filesize"`
	Height		any			`json:"height"`
	CreatedAt 	time.Time 	`json:"createdAt"`
}

type GeneratedImageMetaData struct{
	Id				any     `json:"id,omitempty"`
	Userid 		 	int64 	`json:"userId"`
	SourceFieldId	int64	`json:"sourceFileId"`
	StorageKey		string	`json:"storageKey"`
	Filename		string	`json:"filename"`
	Mimetype		string	`json:"mimetype"`
	Width			int		`json:"width"`
	Height			int		`json:"height"`
	FileSize		int64	`json:"filesize"`
}
type BaseFileData struct{
	Id 			int64		`json:"uploadedId"`
	Filename	string		`json:"fileName"`
	FileType	string		`json:"fileType"`
	CreatedAt 	time.Time 	`json:"createdAt"`
}

type BaseImageMetaData struct{
	Id				any     `json:"imageId"`
	ImageName		string	`json:"imageName"`
	Mimetype		string	`json:"mimetype"`
	Width			int		`json:"width"`
	Height			int		`json:"height"`
	FileSize		int64	`json:"filesize"`
}
type FileBasedImageGenerationReturn struct{
	ActualFile		BaseFileData				`json:"inputFile"`
	GeneratedImage	[]BaseImageMetaData		 	`json:"images"`
}


type History struct {
	ID            int64          `json:"id"`
	UserID        int64          `json:"userId"`
	SourceFileID  int64          `json:"sourceFileId"`
	OutputImageID int64          `json:"outputImageId"`
	OperationType string         `json:"operationType"`
	Prompt        *string        `json:"prompt"`
	Parameters    map[string]any `json:"parameters"`
	Status        string         `json:"status"`
	CreatedAt     time.Time      `json:"createdAt"`
	CompletedAt   time.Time      `json:"completedAt"`
}

type HistoricalData struct{
	UploadedFileId		int			`json:"uploadedId"`
	UploadedFileName 	string		`json:"uploadedFileName"`
	CreatedAt 			time.Time   `json:"createdAt"`
	ImageId 			int 		`json:"imageid"`
	ImageName 			string 		`json:"imageName"`
	MimeType 			string 		`json:"mimeType"`
	Width			int		`json:"width"`
	Height			int		`json:"height"`
	FileSize		int64	`json:"filesize"`
}