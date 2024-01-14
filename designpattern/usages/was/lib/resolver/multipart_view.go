package resolver

import (
	"design_patterns/domain"
	"fmt"
	"io"
	"os"
	"strconv"
)

type multipartView struct{}

func (i multipartView) Resolve(param domain.ViewResolverConfig) {
	fileName, ok := param.Data.(string)
	if ok {
		// Open the file for reading
		file, err := os.Open("/Users/lines/sources/02_linesgits/lines_edu/designpattern/usages/was/lib/" + fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Determine the file size
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Println("Error getting file info:", err)
			return
		}
		fileSize := fileInfo.Size()

		// Construct the HTTP response header
		responseHeader := "HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: " + strconv.Itoa(int(fileSize)) + "\r\n\r\n"

		// Send the HTTP response header
		_, err = param.Conn.Write([]byte(responseHeader))

		// Send the file contents
		_, err = io.Copy(param.Conn, file)
		if err != nil {
			fmt.Println("Error sending file contents:", err)
			return
		}
	} else {
		content := "Path is empty"

		fmt.Fprintf(param.Conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(param.Conn, "Content-Length: %d\r\n", len(content))
		fmt.Fprintf(param.Conn, "\r\n")
		fmt.Fprintf(param.Conn, "%s\r\n", content)
	}
}

func NewMultiPartView() domain.ViewResolver {
	return &multipartView{}
}
