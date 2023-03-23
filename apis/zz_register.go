/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/aviatrix/provider-aviatrix/apis/aviatrix/v1alpha1"
	v1alpha1gateway "github.com/aviatrix/provider-aviatrix/apis/gateway/v1alpha1"
	v1alpha1segmentation "github.com/aviatrix/provider-aviatrix/apis/segmentation/v1alpha1"
	v1alpha1spoke "github.com/aviatrix/provider-aviatrix/apis/spoke/v1alpha1"
	v1alpha1transit "github.com/aviatrix/provider-aviatrix/apis/transit/v1alpha1"
	v1alpha1apis "github.com/aviatrix/provider-aviatrix/apis/v1alpha1"
	v1beta1 "github.com/aviatrix/provider-aviatrix/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1gateway.SchemeBuilder.AddToScheme,
		v1alpha1segmentation.SchemeBuilder.AddToScheme,
		v1alpha1spoke.SchemeBuilder.AddToScheme,
		v1alpha1transit.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
