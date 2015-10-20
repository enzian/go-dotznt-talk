type Vertex struct {
    X int 	  // public member
    xi int  	// private member
    Y int
}

func (vertex *Vertex) Scale(int factor) { … }

// Function Type definition
type HandlerFunc func(request HttpRequest, response HttpResponse)( status int, err error )

type Reader interface {
    Read(b []byte) (n int, err error)
}

type FileReader struct { … }
func (reader *FileReader) Read(b []byte) (n int, err error){ … }
