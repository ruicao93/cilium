// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package types

import "github.com/cilium/cilium/pkg/ipam/types"

// Spec is the ENI specification of a node. This specification is considered
// by the cilium-operator to act as an IPAM operator and makes ENI IPs available
// via the IPAMSpec section.
//
// The ENI specification can either be provided explicitly by the user or the
// cilium agent running on the node can be instructed to create the CiliumNode
// custom resource along with an ENI specification when the node registers
// itself to the Kubernetes cluster.
type Spec struct {
	// InstanceType is the Volcengine ECS instance type, e.g. "ecs.g1.large"
	//
	// +kubebuilder:validation:Optional
	InstanceType string `json:"instance-type,omitempty"`

	// AvailabilityZone is the availability zone to use when allocating
	// ENIs.
	//
	// +kubebuilder:validation:Optional
	AvailabilityZone string `json:"availability-zone,omitempty"`

	// VPCID is the VPC ID to use when allocating ENIs.
	//
	// +kubebuilder:validation:Optional
	VPCID string `json:"vpc-id,omitempty"`

	// CIDRBlock is VPC IPv4 CIDR
	//
	// +kubebuilder:validation:Optional
	CIDRBlock string `json:"cidr-block,omitempty"`

	// SubnetIDs is the list of subnet ids to use when evaluating what Volcengine
	// subnets to use for ENI and IP allocation.
	//
	// +kubebuilder:validation:Optional
	SubnetIDs []string `json:"subnet-ids,omitempty"`

	// SubnetTags is the list of tags to use when evaluating what Volcengine
	// subnets to use for ENI and IP allocation.
	//
	// +kubebuilder:validation:Optional
	SubnetTags map[string]string `json:"subnet-tags,omitempty"`
}

const (
	// ENITypePrimary is the type for primary ENI
	ENITypePrimary = "primary"
	// ENITypeSecondary is the type for secondary ENI
	ENITypeSecondary = "secondary"
)

// ENI represents an Volcengine Network Interface
type ENI struct {
	// NetworkInterfaceID is the ENI ID
	NetworkInterfaceID string `json:"network-interface-id,omitempty"`

	// Type is the type of the ENI: Primary or Secondary
	Type string `json:"type,omitempty"`

	// MACAddress is the MAC address of the ENI
	MACAddress string `json:"mac-address,omitempty"`

	// PrimaryIPAddress is the primary private IPv4 address of the ENI
	PrimaryIPAddress string `json:"primary-ip-address,omitempty"`

	// PrivateIPSets is the list of all private IP addresses of the ENI
	PrivateIPSets []PrivateIPSet `json:"private-ipsets,omitempty"`

	// VPC is the VPC to which the ENI is attached
	VPC VPC `json:"vpc,omitempty"`

	// ZoneID is the zone to which the ENI is attached
	ZoneID string `json:"zone-id,omitempty"`

	// Subnet is the subnet to which the ENI is attached
	Subnet Subnet `json:"subnet,omitempty"`

	// InstanceID is the ID of the instance to which the ENI is attached
	InstanceID string `json:"instance-id,omitempty"`

	// SecurityGroupIDs is the list of security group IDs associated with the ENI
	SecurityGroupIds []string `json:"security-group-ids,omitempty"`

	// Tags is the set of tags associated with the ENI
	//
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

// InterfaceID returns the identifier of the interface
func (e *ENI) InterfaceID() string {
	return e.NetworkInterfaceID
}

// ForeachAddress iterates over all addresses and calls fn
func (e *ENI) ForeachAddress(id string, fn types.AddressIterator) error {
	for _, address := range e.PrivateIPSets {
		if address.Primary {
			continue
		}
		if err := fn(id, e.NetworkInterfaceID, address.PrivateIpAddress, "", address); err != nil {
			return err
		}
	}

	return nil
}

type VPC struct {
	// VPCID is the ID of the VPC
	VPCID string `json:"vpc-id,omitempty"`

	// CIDRBlock is the IPv4 CIDR of the VPC
	//
	// +optional
	CIDRBlock string `json:"cidr-block,omitempty"`

	// IPv6CIDRBlock is the IPv6 CIDR of the VPC
	//
	// +optional
	IPv6CIDRBlock string `json:"ipv6-cidr-block,omitempty"`

	// SecondaryCIDRBlocks is the list of secondary CIDR blocks of the VPC
	//
	// +optional
	SecondaryCIDRBlocks []string `json:"secondary-cidr-blocks,omitempty"`
}

// Subnet stores the information about a Volcengine subnet
type Subnet struct {
	// SubnetID is the ID of the subnet
	SubnetID string `json:"subnet-id,omitempty"`

	// CIDRBlock is the IPv4 CIDR of the subnet
	//
	// +optional
	CIDRBlock string `json:"cidr-block,omitempty"`

	// IPv6CIDRBlock is the IPv6 CIDR of the subnet
	//
	// +optional
	IPv6CIDRBlock string `json:"ipv6-cidr-block,omitempty"`
}

// PrivateIPSet stores the information about a private IP address
type PrivateIPSet struct {
	PrivateIpAddress string `json:"private-ip-address,omitempty"`
	Primary          bool   `json:"primary,omitempty" `
}

type ENIStatus struct {
	// ENIs is the list of ENIs on the node
	//
	// +optional
	ENIS []ENI `json:"enis,omitempty"`
}
