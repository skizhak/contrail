package models

// APIAccessList

import "encoding/json"

// APIAccessList
type APIAccessList struct {
	UUID                 string               `json:"uuid,omitempty"`
	FQName               []string             `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType         `json:"id_perms,omitempty"`
	DisplayName          string               `json:"display_name,omitempty"`
	Annotations          *KeyValuePairs       `json:"annotations,omitempty"`
	APIAccessListEntries *RbacRuleEntriesType `json:"api_access_list_entries,omitempty"`
	ParentUUID           string               `json:"parent_uuid,omitempty"`
	ParentType           string               `json:"parent_type,omitempty"`
	Perms2               *PermType2           `json:"perms2,omitempty"`
}

// String returns json representation of the object
func (model *APIAccessList) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeAPIAccessList makes APIAccessList
func MakeAPIAccessList() *APIAccessList {
	return &APIAccessList{
		//TODO(nati): Apply default
		Perms2:               MakePermType2(),
		APIAccessListEntries: MakeRbacRuleEntriesType(),
		ParentUUID:           "",
		ParentType:           "",
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		UUID:                 "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
	}
}

// MakeAPIAccessListSlice() makes a slice of APIAccessList
func MakeAPIAccessListSlice() []*APIAccessList {
	return []*APIAccessList{}
}
