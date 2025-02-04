/*
 * AgentCentral API
 *
 * This is agent central swagger API
 *
 * API version: 0.dev
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"os"
)

type FilesUploadBody struct {
	// file name
	FileName string `json:"fileName"`
	// file block
	Block string `json:"block"`
	// this is a file
	File **os.File `json:"file"`
}
