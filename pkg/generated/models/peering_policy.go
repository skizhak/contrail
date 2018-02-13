package models

// PeeringPolicy

import "encoding/json"

// PeeringPolicy
//proteus:generate
type PeeringPolicy struct {
	UUID           string             `json:"uuid,omitempty"`
	ParentUUID     string             `json:"parent_uuid,omitempty"`
	ParentType     string             `json:"parent_type,omitempty"`
	FQName         []string           `json:"fq_name,omitempty"`
	IDPerms        *IdPermsType       `json:"id_perms,omitempty"`
	DisplayName    string             `json:"display_name,omitempty"`
	Annotations    *KeyValuePairs     `json:"annotations,omitempty"`
	Perms2         *PermType2         `json:"perms2,omitempty"`
	PeeringService PeeringServiceType `json:"peering_service,omitempty"`
}

// String returns json representation of the object
func (model *PeeringPolicy) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePeeringPolicy makes PeeringPolicy
func MakePeeringPolicy() *PeeringPolicy {
	return &PeeringPolicy{
		//TODO(nati): Apply default
		UUID:           "",
		ParentUUID:     "",
		ParentType:     "",
		FQName:         []string{},
		IDPerms:        MakeIdPermsType(),
		DisplayName:    "",
		Annotations:    MakeKeyValuePairs(),
		Perms2:         MakePermType2(),
		PeeringService: MakePeeringServiceType(),
	}
}

// MakePeeringPolicySlice() makes a slice of PeeringPolicy
func MakePeeringPolicySlice() []*PeeringPolicy {
	return []*PeeringPolicy{}
}
