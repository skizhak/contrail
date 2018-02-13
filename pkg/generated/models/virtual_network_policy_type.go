package models

// VirtualNetworkPolicyType

import "encoding/json"

// VirtualNetworkPolicyType
//proteus:generate
type VirtualNetworkPolicyType struct {
	Timer    *TimerType    `json:"timer,omitempty"`
	Sequence *SequenceType `json:"sequence,omitempty"`
}

// String returns json representation of the object
func (model *VirtualNetworkPolicyType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeVirtualNetworkPolicyType makes VirtualNetworkPolicyType
func MakeVirtualNetworkPolicyType() *VirtualNetworkPolicyType {
	return &VirtualNetworkPolicyType{
		//TODO(nati): Apply default
		Timer:    MakeTimerType(),
		Sequence: MakeSequenceType(),
	}
}

// MakeVirtualNetworkPolicyTypeSlice() makes a slice of VirtualNetworkPolicyType
func MakeVirtualNetworkPolicyTypeSlice() []*VirtualNetworkPolicyType {
	return []*VirtualNetworkPolicyType{}
}
