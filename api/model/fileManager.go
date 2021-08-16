package model

type FileManager struct {
	Com		string	`form:"command" `
	Arg 	ArgMap	`form:"arguments"`

}

type ArgMap struct {
	PathInfo				[]PathMap 	`form:"pathInfo"`
	PathInfoList			[][]PathMap	`form:"pathInfoList"`
	Name 					string 		`form:"name"`
	SourcePathInfo			[]PathMap	`form:"sourcePathInfo"`
	DestinationPathInfo 	[]PathMap	`form:"destinationPathInfo"`
	ChunkMetadata			string 	`form:"chunkMetadata"`
	UploadId				string 	`form:"uploadId"`
}

type PathMap struct {
	Key 	string	`form:"key"`
	Name 	string 	`form:"name"`
}

type ChunkMetadataMap struct {
	UploadId	string	`Json:""`
	FileName	string	`Json:""`
	Index		int 	`Json:""`
	TotalCount	int 	`Json:""`
	FileSize	int64 	`Json:""`
}