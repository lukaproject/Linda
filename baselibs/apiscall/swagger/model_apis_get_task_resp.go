/*
 * AgentCentral API
 *
 * This is agent central swagger API
 *
 * API version: 0.dev
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ApisGetTaskResp struct {
	BagName string `json:"bagName,omitempty"`
	CreateTimeMs int32 `json:"createTimeMs,omitempty"`
	FinishTimeMs int32 `json:"finishTimeMs,omitempty"`
	NodeId string `json:"nodeId,omitempty"`
	Priority int32 `json:"priority,omitempty"`
	ScheduledTimeMs int32 `json:"scheduledTimeMs,omitempty"`
	ScriptPath string `json:"scriptPath,omitempty"`
	TaskDisplayName string `json:"taskDisplayName,omitempty"`
	TaskName string `json:"taskName,omitempty"`
	WorkingDir string `json:"workingDir,omitempty"`
}
