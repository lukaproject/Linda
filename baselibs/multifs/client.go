package multifs

import (
	"Linda/baselibs/abstractions/xio"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	EndPoint string

	client *http.Client
}

func (c *Client) Upload(
	targetFilePath string,
	fileName string,
	reader io.Reader,
) (err error) {
	targetFilePath = strings.TrimPrefix(targetFilePath, "/")
	url := c.EndPoint + "/upload/" + targetFilePath
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	fileWriter, _ := bodyWriter.CreateFormFile("file", fileName)
	io.Copy(fileWriter, reader)
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(url, contentType, bodyBuffer)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("upload file failed, status code is %d", resp.StatusCode)
	}
	return
}

func (c *Client) DownloadToWriter(
	filePath string,
	writer io.Writer,
) (err error) {
	filePath = strings.TrimPrefix(filePath, "/")
	url := c.EndPoint + "/files/" + filePath
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download file failed, status code is %d", resp.StatusCode)
	}
	return xio.Transport(resp.Body, writer)
}

func NewClient(endpoint string) *Client {
	return &Client{
		EndPoint: endpoint,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
