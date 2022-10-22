package schema

import "time"

type ReturnReason struct {
	Value                string         `json:"value"`
	Label                string         `json:"Label"`
	Id                   string         `json:"id"`
	Description          string         `json:"description"`
	ParentReturnReason   *ReturnReason  `json:"parent_return_reason"`
	ParentReturnReasonId string         `json:"parent_return_reason_id"`
	ReturnReasonChildren *ReturnReason  `json:"return_reason_children"`
	CreatedAt            *time.Time     `json:"created_at"`
	UpdatedAt            *time.Time     `json:"updated_at"`
	DeletedAt            *time.Time     `json:"deleted_at"`
	Metadata             map[string]any `json:"metadata"`
}
