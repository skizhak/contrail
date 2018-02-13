package models

// RouteTable

import "encoding/json"

// RouteTable
//proteus:generate
type RouteTable struct {
	UUID        string          `json:"uuid,omitempty"`
	ParentUUID  string          `json:"parent_uuid,omitempty"`
	ParentType  string          `json:"parent_type,omitempty"`
	FQName      []string        `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType    `json:"id_perms,omitempty"`
	DisplayName string          `json:"display_name,omitempty"`
	Annotations *KeyValuePairs  `json:"annotations,omitempty"`
	Perms2      *PermType2      `json:"perms2,omitempty"`
	Routes      *RouteTableType `json:"routes,omitempty"`
}

// String returns json representation of the object
func (model *RouteTable) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRouteTable makes RouteTable
func MakeRouteTable() *RouteTable {
	return &RouteTable{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		Routes:      MakeRouteTableType(),
	}
}

// MakeRouteTableSlice() makes a slice of RouteTable
func MakeRouteTableSlice() []*RouteTable {
	return []*RouteTable{}
}
