package models

// Dashboard

import "encoding/json"

// Dashboard
type Dashboard struct {
	UUID            string         `json:"uuid,omitempty"`
	IDPerms         *IdPermsType   `json:"id_perms,omitempty"`
	Annotations     *KeyValuePairs `json:"annotations,omitempty"`
	Perms2          *PermType2     `json:"perms2,omitempty"`
	ContainerConfig string         `json:"container_config,omitempty"`
	ParentUUID      string         `json:"parent_uuid,omitempty"`
	ParentType      string         `json:"parent_type,omitempty"`
	FQName          []string       `json:"fq_name,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`
}

// String returns json representation of the object
func (model *Dashboard) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeDashboard makes Dashboard
func MakeDashboard() *Dashboard {
	return &Dashboard{
		//TODO(nati): Apply default
		UUID:            "",
		IDPerms:         MakeIdPermsType(),
		Annotations:     MakeKeyValuePairs(),
		Perms2:          MakePermType2(),
		DisplayName:     "",
		ContainerConfig: "",
		ParentUUID:      "",
		ParentType:      "",
		FQName:          []string{},
	}
}

// MakeDashboardSlice() makes a slice of Dashboard
func MakeDashboardSlice() []*Dashboard {
	return []*Dashboard{}
}
