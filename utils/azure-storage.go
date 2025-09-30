package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
)

// GenerateSASToken returns the SAS URL and filename
func GenerateSASToken(fileName string) (string, string, error) {
	accountName := os.Getenv("AZURE_STORAGE_ACCOUNT")
	accountKey := os.Getenv("AZURE_STORAGE_KEY")
	containerName := os.Getenv("AZURE_CONTAINER_NAME")

	if fileName == "" {
		return "", "", fmt.Errorf("filename is required")
	}

	cred, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to create credential: %v", err)
	}

	TimefileName := fmt.Sprintf("%s-%s", time.Now().Format("20060102150405"), fileName)

	permissions := sas.BlobPermissions{
		Read:   true,
		Write:  true,
		Create: true,
		Add:    true,
	}

	sasQueryParams, err := sas.BlobSignatureValues{
		ContainerName: containerName,
		// Protocol:      sas.ProtocolHTTPS,
		BlobName:    TimefileName,
		Permissions: permissions.String(),
		StartTime:   time.Now().UTC(),
		ExpiryTime:  time.Now().UTC().Add(1 * time.Hour),
	}.SignWithSharedKey(cred)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate SAS query parameters: %v", err)
	}

	blobURL := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s", accountName, containerName, TimefileName)
	sasURL := fmt.Sprintf("%s?%s", blobURL, sasQueryParams.Encode())

	return sasURL, fileName, nil
}

// UploadFileToBlob handles file upload to Azure Blob Storage using the SAS URL
func UploadFileToBlob(sasUrl string, fileContent io.Reader, contentType string) error {
	// Create a buffer to hold the file content and calculate its length
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, fileContent); err != nil {
		return fmt.Errorf("failed to copy file content: %v", err)
	}

	// Create the PUT request
	req, err := http.NewRequest("PUT", sasUrl, buf)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers required for blob storage
	req.Header.Set("x-ms-blob-type", "BlockBlob")
	req.Header.Set("Content-Type", contentType)     // Set Content-Type based on file type
	req.Header.Set("Content-Disposition", "inline") // Set Content-Disposition to inline
	req.ContentLength = int64(buf.Len())            // Set Content-Length based on buffer length

	// Create the HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to upload file: %v", err)
	}
	defer resp.Body.Close()

	// Log the full response for debugging
	body, _ := io.ReadAll(resp.Body)
	log.Printf("Response Status: %s", resp.Status)
	log.Printf("Response Body: %s", body)

	// Check for success (201 Created or 202 Accepted)
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("failed to upload file, status code: %d", resp.StatusCode)
	}

	return nil
}

func GetContentType(fileName string) string {
	switch {
	case strings.HasSuffix(fileName, ".pdf"):
		return "application/pdf"
	case strings.HasSuffix(fileName, ".jpg"), strings.HasSuffix(fileName, ".jpeg"):
		return "image/jpeg"
	case strings.HasSuffix(fileName, ".png"):
		return "image/png"
	case strings.HasSuffix(fileName, ".gif"):
		return "image/gif"
	case strings.HasSuffix(fileName, ".txt"):
		return "text/plain"
	// Add more cases as needed
	default:
		return "application/zip"
	}
}
