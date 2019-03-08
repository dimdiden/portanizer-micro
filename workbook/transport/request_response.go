package transport

import (
	"github.com/dimdiden/portanizer-micro/workbook"
)

// CreatePostRequest holds the request parameters for the CreatePost method.
type CreatePostRequest struct {
	Post workbook.Post
}

// CreatePostResponse holds the response values for the CreatePost method.
type CreatePostResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

// UpdatePostRequest holds the request parameters for the UpdatePost method.
type UpdatePostRequest struct {
	ID   string
	Post workbook.Post
}

// UpdatePostResponse holds the response values for the UpdatePost method.
type UpdatePostResponse struct {
	Err error `json:"error,omitempty"`
}

// GetByIDPostRequest holds the request parameters for the GetByIDPost method.
type GetByIDPostRequest struct {
	ID string
}

// GetByIDPostResponse holds the response values for the GetByIDPost method.
type GetByIDPostResponse struct {
	Post *workbook.Post `json:"post"`
	Err  error          `json:"error,omitempty"`
}

func (r GetByIDPostResponse) Error() error { return r.Err }

// GetAllPostRequest holds the request parameters for the GetAllPost method.
type GetAllPostRequest struct {
}

// GetAllPostResponse holds the response values for the GetAllPost method.
type GetAllPostResponse struct {
	Posts []*workbook.Post `json:"posts"`
	Err   error            `json:"error,omitempty"`
}

// DeletePostRequest holds the request parameters for the DeletePost method.
type DeletePostRequest struct {
	ID string
}

// DeletePostResponse holds the response values for the DeletePost method.
type DeletePostResponse struct {
	Err error `json:"error,omitempty"`
}

// ======================
// CreateTagRequest holds the request parameters for the CreateTag method.
type CreateTagRequest struct {
	PostID string
	Tag    workbook.Tag
}

// CreateTagResponse holds the response values for the CreateTag method.
type CreateTagResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

// UpdateTagRequest holds the request parameters for the UpdateTag method.
type UpdateTagRequest struct {
	ID  string
	Tag workbook.Tag
}

// UpdateTagResponse holds the response values for the UpdateTag method.
type UpdateTagResponse struct {
	Err error `json:"error,omitempty"`
}

// GetAllTagRequest holds the request parameters for the GetAllTag method.
type GetAllTagRequest struct {
}

// GetAllTagResponse holds the response values for the GetAllTag method.
type GetAllTagResponse struct {
	Tags []*workbook.Tag `json:"tags"`
	Err  error           `json:"error,omitempty"`
}

// DeleteTagRequest holds the request parameters for the DeleteTag method.
type DeleteTagRequest struct {
	ID string
}

// DeleteTagResponse holds the response values for the DeleteTag method.
type DeleteTagResponse struct {
	Err error `json:"error,omitempty"`
}
