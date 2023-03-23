/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PrivateSubnetsObservation struct {

	// VPC CIDR. Required to be empty for GCP provider, and non-empty for other providers. Example: "10.11.0.0/24".
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Name of the VPC to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of this subnet.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type PrivateSubnetsParameters struct {
}

type PublicSubnetsObservation struct {

	// VPC CIDR. Required to be empty for GCP provider, and non-empty for other providers. Example: "10.11.0.0/24".
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Name of the VPC to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of this subnet.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type PublicSubnetsParameters struct {
}

type SubnetsObservation struct {

	// ID of this subnet.
	// Subnet ID.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type SubnetsParameters struct {

	// VPC CIDR. Required to be empty for GCP provider, and non-empty for other providers. Example: "10.11.0.0/24".
	// Subnet cidr.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Name of the VPC to be created.
	// Subnet name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region of cloud provider. Required to be empty for GCP provider, and non-empty for other providers. Example: AWS: "us-east-1", Azure: "East US 2", OCI: "us-ashburn-1", AzureGov: "USGov Arizona", AWSGov: "us-gov-east-1", AWSChina: "cn-north-1", AzureChina: "China North".
	// Subnet region.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type VPCObservation struct {

	// List of OCI availability domains.
	// List of OCI availability domains.
	AvailabilityDomains []*string `json:"availabilityDomains,omitempty" tf:"availability_domains,omitempty"`

	// Azure VNet resource ID.
	// Azure vnet resource ID.
	AzureVnetResourceID *string `json:"azureVnetResourceId,omitempty" tf:"azure_vnet_resource_id,omitempty"`

	// List of OCI fault domains.
	// List of OCI fault domains.
	FaultDomains []*string `json:"faultDomains,omitempty" tf:"fault_domains,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of private subnet of the VPC(AWS, Azure) to be created.
	// List of private subnet of the VPC to be created.
	PrivateSubnets []PrivateSubnetsObservation `json:"privateSubnets,omitempty" tf:"private_subnets,omitempty"`

	// List of public subnet of the VPC(AWS, Azure) to be created.
	// List of public subnet of the VPC to be created.
	PublicSubnets []PublicSubnetsObservation `json:"publicSubnets,omitempty" tf:"public_subnets,omitempty"`

	// List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure VPC.
	// List of route table ids associated with this VPC.
	RouteTables []*string `json:"routeTables,omitempty" tf:"route_tables,omitempty"`

	// List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
	// List of subnet of the VPC to be created. Required to be non-empty for GCP provider, and empty for other providers.
	// +kubebuilder:validation:Optional
	Subnets []SubnetsObservation `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// ID of the VPC to be created.
	// ID of the VPC created.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCParameters struct {

	// This parameter represents the name of a Cloud-Account in Aviatrix controller.
	// Account name. This account will be used to create an Aviatrix VPC.
	// +kubebuilder:validation:Required
	AccountName *string `json:"accountName" tf:"account_name,omitempty"`

	// Specify whether it is an Aviatrix FireNet VPC to be used for Aviatrix FireNet and Transit FireNet solutions. Only AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina are supported. Required to be false for other providers. Valid values: true, false. Default: false.
	// Specify the VPC as Aviatrix FireNet VPC or not. Required to be false for GCP provider.
	// +kubebuilder:validation:Optional
	AviatrixFirenetVPC *bool `json:"aviatrixFirenetVpc,omitempty" tf:"aviatrix_firenet_vpc,omitempty"`

	// Specify whether it is an Aviatrix Transit VPC to be used for Transit Network or TGW solutions. Only AWS, AWSGov, AWSChina, and Alibaba Cloud are supported. Required to be false for other providers. Valid values: true, false. Default: false.
	// Specify the VPC as Aviatrix Transit VPC or not. Required to be false for GCP provider.
	// +kubebuilder:validation:Optional
	AviatrixTransitVPC *bool `json:"aviatrixTransitVpc,omitempty" tf:"aviatrix_transit_vpc,omitempty"`

	// VPC CIDR. Required to be empty for GCP provider, and non-empty for other providers. Example: "10.11.0.0/24".
	// Subnet of the VPC to be created. Required to be empty for GCP provider, and non-empty for other providers.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.
	// Type of cloud service provider.
	// +kubebuilder:validation:Required
	CloudType *float64 `json:"cloudType" tf:"cloud_type,omitempty"`

	// Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloud_type = 1 (AWS). This option is only applicable to TGW-integrated FireNet. Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
	// Enable Native AWS GWLB for FireNet Function. Only valid with cloud_type = 1 (AWS). Valid values: true or false. Default value: false. Available as of provider version R2.18+.
	// +kubebuilder:validation:Optional
	EnableNativeGwlb *bool `json:"enableNativeGwlb,omitempty" tf:"enable_native_gwlb,omitempty"`

	// Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.
	// Switch to enable private oob subnet. Only supported for AWS/AWSGov provider. Valid values: true, false. Default value: false.
	// +kubebuilder:validation:Optional
	EnablePrivateOobSubnet *bool `json:"enablePrivateOobSubnet,omitempty" tf:"enable_private_oob_subnet,omitempty"`

	// Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.
	// Number of public subnet and private subnet pair to be created.
	// +kubebuilder:validation:Optional
	NumOfSubnetPairs *float64 `json:"numOfSubnetPairs,omitempty" tf:"num_of_subnet_pairs,omitempty"`

	// Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+.
	// Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller.
	// +kubebuilder:validation:Optional
	PrivateModeSubnets *bool `json:"privateModeSubnets,omitempty" tf:"private_mode_subnets,omitempty"`

	// Region of cloud provider. Required to be empty for GCP provider, and non-empty for other providers. Example: AWS: "us-east-1", Azure: "East US 2", OCI: "us-ashburn-1", AzureGov: "USGov Arizona", AWSGov: "us-gov-east-1", AWSChina: "cn-north-1", AzureChina: "China North".
	// Region of cloud provider. Required to be empty for GCP provider, and non-empty for other providers.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The name of an existing resource group or a new resource group to be created for the Azure VNet.  A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+.
	// Resource group of the Azure VPC created.
	// +kubebuilder:validation:Optional
	ResourceGroup *string `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`

	// Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.
	// Subnet size.
	// +kubebuilder:validation:Optional
	SubnetSize *float64 `json:"subnetSize,omitempty" tf:"subnet_size,omitempty"`

	// List of subnets to be specify for GCP provider. Required to be non-empty for GCP provider, and empty for other providers.
	// List of subnet of the VPC to be created. Required to be non-empty for GCP provider, and empty for other providers.
	// +kubebuilder:validation:Optional
	Subnets []SubnetsParameters `json:"subnets,omitempty" tf:"subnets,omitempty"`
}

// VPCSpec defines the desired state of VPC
type VPCSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCParameters `json:"forProvider"`
}

// VPCStatus defines the observed state of VPC.
type VPCStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPC is the Schema for the VPCs API. Creates and manages Aviatrix-created VPCs
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aviatrix}
type VPC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCSpec   `json:"spec"`
	Status            VPCStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCList contains a list of VPCs
type VPCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPC `json:"items"`
}

// Repository type metadata.
var (
	VPC_Kind             = "VPC"
	VPC_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPC_Kind}.String()
	VPC_KindAPIVersion   = VPC_Kind + "." + CRDGroupVersion.String()
	VPC_GroupVersionKind = CRDGroupVersion.WithKind(VPC_Kind)
)

func init() {
	SchemeBuilder.Register(&VPC{}, &VPCList{})
}
