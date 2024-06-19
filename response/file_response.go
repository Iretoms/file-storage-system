package response

type FileResponse struct {
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	FileReference interface{} `json:"file_reference"`
}
