package handlers

import "time"

type CreateStackRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SegmentId   int    `json:"segmentId"`
	PurposeId   int    `json:"purposeId"`
}

type CreateStackResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SegmentName string `json:"segmentName"`
	PurposeName string `json:"purposeName"`
}
type SearchStackResponse struct {
	Id              int             `json:"id"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	SegmentResponse SegmentResponse `json:"segment"`
	PurposeResponse PurposeResponse `json:"purpose"`
	CreatedBy       string          `json:"createdBy"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedBy       string          `json:"updatedBy"`
	UpdatedAt       time.Time       `json:"updatedAt"`
}

type SegmentResponse struct {
	Name         string `json:"name"`
	IsProductive int64  `json:"isProductive"`
}

type PurposeResponse struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsProductive int64  `json:"isProductive"`
}
