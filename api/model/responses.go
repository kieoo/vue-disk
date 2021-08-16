package model

import "time"

type ResForm struct {
	Suc 		bool	`json:"success"`
	ErrorCode 	*int	`json:"errorCode"`
	ErrorText	string 	`json:"errorText"`
	Result 		*[]InfoList 	`json:"result"`
}

type InfoList struct {
	Key				string 		`json:"key"`
	Name			string		`json:"name"`
	DateModified	time.Time	`time_format:"2019-10-25T12:52:22.9529346Z" json:"dateModified"`
	IsDirectory		bool 		`json:"isDirectory"`
	Size			int64 		`json:"size"`
	HasSubD			bool 		`json:"hasSubDirectories"`
	Url				string 		`json:"url"`
}
