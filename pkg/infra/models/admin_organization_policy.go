package models

import "time"

type OrganizationPolicyPageScheme struct {
	Data  []*OrganizationPolicyData `json:"data"`
	Links *LinkPageModelScheme      `json:"links"`
	Meta  struct {
		Next     string `json:"next"`
		PageSize int    `json:"page_size"`
	} `json:"meta"`
}

type OrganizationPolicyScheme struct {
	Data OrganizationPolicyData `json:"data"`
}
type OrganizationPolicyResource struct {
	ID                string `json:"id,omitempty"`
	ApplicationStatus string `json:"applicationStatus,omitempty"`
}
type OrganizationPolicyAttributes struct {
	Type      string                        `json:"type,omitempty"`
	Name      string                        `json:"name,omitempty"`
	Status    string                        `json:"status,omitempty"`
	Resources []*OrganizationPolicyResource `json:"resources,omitempty"`
	CreatedAt time.Time                     `json:"createdAt,omitempty"`
	UpdatedAt time.Time                     `json:"updatedAt,omitempty"`
}
type OrganizationPolicyData struct {
	ID         string                        `json:"id,omitempty"`
	Type       string                        `json:"type,omitempty"`
	Attributes *OrganizationPolicyAttributes `json:"attributes,omitempty"`
}
