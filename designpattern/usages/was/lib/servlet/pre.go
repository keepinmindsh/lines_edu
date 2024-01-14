package servlet

import (
	"design_patterns/domain"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func (s servletContainer) Pre(conn net.Conn) (domain.ParseResult, error) {
	buf := make([]byte, 1024)

	// Read the request header
	n, err := conn.Read(buf)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return domain.ParseResult{}, err
		} else {
			fmt.Println(err.Error())
		}
	}

	request := string(buf[:n])

	// Parse the request method and path
	parseResult := parseRequestLine(request)

	// todo content 타입을 체크하는 로직
	if parseResult.Err != nil {
		conn.Close()
		return domain.ParseResult{}, errors.New("servlet container - pre handle error")
	}

	return parseResult, nil
}

func parseRequestLine(request string) domain.ParseResult {
	log.Print(request)

	parts := strings.Split(request, "\r\n")

	var split []string
	var contentType string
	for _, part := range parts {
		if strings.Contains(part, "HTTP") {
			split = strings.Split(parts[0], " ")
			break
		}

		if strings.Contains(part, "Content-Type") {
			contentType = strings.Split(part, " ")[1]
		}
	}

	if contentType == "" {
		return domain.ParseResult{
			Err:         errors.New("content-Type is empty, please check your header info"),
			Method:      "",
			Path:        "",
			ContentType: "",
		}
	}

	return domain.ParseResult{
		Err:         nil,
		Method:      split[0],
		Path:        split[1],
		ContentType: contentType,
	}
}
