// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/stackhcideploymentsetting/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StackHciDeploymentSettingScaleUnitOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectoryOrganizationalUnitPath() *string
	SetActiveDirectoryOrganizationalUnitPath(val *string)
	ActiveDirectoryOrganizationalUnitPathInput() *string
	BitlockerBootVolumeEnabled() interface{}
	SetBitlockerBootVolumeEnabled(val interface{})
	BitlockerBootVolumeEnabledInput() interface{}
	BitlockerDataVolumeEnabled() interface{}
	SetBitlockerDataVolumeEnabled(val interface{})
	BitlockerDataVolumeEnabledInput() interface{}
	Cluster() StackHciDeploymentSettingScaleUnitClusterOutputReference
	ClusterInput() *StackHciDeploymentSettingScaleUnitCluster
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CredentialGuardEnabled() interface{}
	SetCredentialGuardEnabled(val interface{})
	CredentialGuardEnabledInput() interface{}
	DomainFqdn() *string
	SetDomainFqdn(val *string)
	DomainFqdnInput() *string
	DriftControlEnabled() interface{}
	SetDriftControlEnabled(val interface{})
	DriftControlEnabledInput() interface{}
	DrtmProtectionEnabled() interface{}
	SetDrtmProtectionEnabled(val interface{})
	DrtmProtectionEnabledInput() interface{}
	EpisodicDataUploadEnabled() interface{}
	SetEpisodicDataUploadEnabled(val interface{})
	EpisodicDataUploadEnabledInput() interface{}
	EuLocationEnabled() interface{}
	SetEuLocationEnabled(val interface{})
	EuLocationEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	HostNetwork() StackHciDeploymentSettingScaleUnitHostNetworkOutputReference
	HostNetworkInput() *StackHciDeploymentSettingScaleUnitHostNetwork
	HvciProtectionEnabled() interface{}
	SetHvciProtectionEnabled(val interface{})
	HvciProtectionEnabledInput() interface{}
	InfrastructureNetwork() StackHciDeploymentSettingScaleUnitInfrastructureNetworkList
	InfrastructureNetworkInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	OptionalService() StackHciDeploymentSettingScaleUnitOptionalServiceOutputReference
	OptionalServiceInput() *StackHciDeploymentSettingScaleUnitOptionalService
	PhysicalNode() StackHciDeploymentSettingScaleUnitPhysicalNodeList
	PhysicalNodeInput() interface{}
	SecretsLocation() *string
	SetSecretsLocation(val *string)
	SecretsLocationInput() *string
	SideChannelMitigationEnabled() interface{}
	SetSideChannelMitigationEnabled(val interface{})
	SideChannelMitigationEnabledInput() interface{}
	SmbClusterEncryptionEnabled() interface{}
	SetSmbClusterEncryptionEnabled(val interface{})
	SmbClusterEncryptionEnabledInput() interface{}
	SmbSigningEnabled() interface{}
	SetSmbSigningEnabled(val interface{})
	SmbSigningEnabledInput() interface{}
	Storage() StackHciDeploymentSettingScaleUnitStorageOutputReference
	StorageInput() *StackHciDeploymentSettingScaleUnitStorage
	StreamingDataClientEnabled() interface{}
	SetStreamingDataClientEnabled(val interface{})
	StreamingDataClientEnabledInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WdacEnabled() interface{}
	SetWdacEnabled(val interface{})
	WdacEnabledInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutCluster(value *StackHciDeploymentSettingScaleUnitCluster)
	PutHostNetwork(value *StackHciDeploymentSettingScaleUnitHostNetwork)
	PutInfrastructureNetwork(value interface{})
	PutOptionalService(value *StackHciDeploymentSettingScaleUnitOptionalService)
	PutPhysicalNode(value interface{})
	PutStorage(value *StackHciDeploymentSettingScaleUnitStorage)
	ResetBitlockerBootVolumeEnabled()
	ResetBitlockerDataVolumeEnabled()
	ResetCredentialGuardEnabled()
	ResetDriftControlEnabled()
	ResetDrtmProtectionEnabled()
	ResetEpisodicDataUploadEnabled()
	ResetEuLocationEnabled()
	ResetHvciProtectionEnabled()
	ResetSideChannelMitigationEnabled()
	ResetSmbClusterEncryptionEnabled()
	ResetSmbSigningEnabled()
	ResetStreamingDataClientEnabled()
	ResetWdacEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StackHciDeploymentSettingScaleUnitOutputReference
type jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ActiveDirectoryOrganizationalUnitPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryOrganizationalUnitPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ActiveDirectoryOrganizationalUnitPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectoryOrganizationalUnitPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) BitlockerBootVolumeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bitlockerBootVolumeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) BitlockerBootVolumeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bitlockerBootVolumeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) BitlockerDataVolumeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bitlockerDataVolumeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) BitlockerDataVolumeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bitlockerDataVolumeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) Cluster() StackHciDeploymentSettingScaleUnitClusterOutputReference {
	var returns StackHciDeploymentSettingScaleUnitClusterOutputReference
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ClusterInput() *StackHciDeploymentSettingScaleUnitCluster {
	var returns *StackHciDeploymentSettingScaleUnitCluster
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) CredentialGuardEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"credentialGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) CredentialGuardEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"credentialGuardEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) DomainFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) DomainFqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) DriftControlEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"driftControlEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) DriftControlEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"driftControlEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) DrtmProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drtmProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) DrtmProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"drtmProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) EpisodicDataUploadEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"episodicDataUploadEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) EpisodicDataUploadEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"episodicDataUploadEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) EuLocationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"euLocationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) EuLocationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"euLocationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) HostNetwork() StackHciDeploymentSettingScaleUnitHostNetworkOutputReference {
	var returns StackHciDeploymentSettingScaleUnitHostNetworkOutputReference
	_jsii_.Get(
		j,
		"hostNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) HostNetworkInput() *StackHciDeploymentSettingScaleUnitHostNetwork {
	var returns *StackHciDeploymentSettingScaleUnitHostNetwork
	_jsii_.Get(
		j,
		"hostNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) HvciProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hvciProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) HvciProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hvciProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) InfrastructureNetwork() StackHciDeploymentSettingScaleUnitInfrastructureNetworkList {
	var returns StackHciDeploymentSettingScaleUnitInfrastructureNetworkList
	_jsii_.Get(
		j,
		"infrastructureNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) InfrastructureNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infrastructureNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) OptionalService() StackHciDeploymentSettingScaleUnitOptionalServiceOutputReference {
	var returns StackHciDeploymentSettingScaleUnitOptionalServiceOutputReference
	_jsii_.Get(
		j,
		"optionalService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) OptionalServiceInput() *StackHciDeploymentSettingScaleUnitOptionalService {
	var returns *StackHciDeploymentSettingScaleUnitOptionalService
	_jsii_.Get(
		j,
		"optionalServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) PhysicalNode() StackHciDeploymentSettingScaleUnitPhysicalNodeList {
	var returns StackHciDeploymentSettingScaleUnitPhysicalNodeList
	_jsii_.Get(
		j,
		"physicalNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) PhysicalNodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"physicalNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) SecretsLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) SecretsLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) SideChannelMitigationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sideChannelMitigationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) SideChannelMitigationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sideChannelMitigationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) SmbClusterEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbClusterEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) SmbClusterEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbClusterEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) SmbSigningEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbSigningEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) SmbSigningEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbSigningEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) Storage() StackHciDeploymentSettingScaleUnitStorageOutputReference {
	var returns StackHciDeploymentSettingScaleUnitStorageOutputReference
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) StorageInput() *StackHciDeploymentSettingScaleUnitStorage {
	var returns *StackHciDeploymentSettingScaleUnitStorage
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) StreamingDataClientEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingDataClientEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) StreamingDataClientEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingDataClientEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) WdacEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wdacEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) WdacEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wdacEnabledInput",
		&returns,
	)
	return returns
}


func NewStackHciDeploymentSettingScaleUnitOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) StackHciDeploymentSettingScaleUnitOutputReference {
	_init_.Initialize()

	if err := validateNewStackHciDeploymentSettingScaleUnitOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciDeploymentSetting.StackHciDeploymentSettingScaleUnitOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewStackHciDeploymentSettingScaleUnitOutputReference_Override(s StackHciDeploymentSettingScaleUnitOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.stackHciDeploymentSetting.StackHciDeploymentSettingScaleUnitOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetActiveDirectoryOrganizationalUnitPath(val *string) {
	if err := j.validateSetActiveDirectoryOrganizationalUnitPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeDirectoryOrganizationalUnitPath",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetBitlockerBootVolumeEnabled(val interface{}) {
	if err := j.validateSetBitlockerBootVolumeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitlockerBootVolumeEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetBitlockerDataVolumeEnabled(val interface{}) {
	if err := j.validateSetBitlockerDataVolumeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitlockerDataVolumeEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetCredentialGuardEnabled(val interface{}) {
	if err := j.validateSetCredentialGuardEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialGuardEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetDomainFqdn(val *string) {
	if err := j.validateSetDomainFqdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainFqdn",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetDriftControlEnabled(val interface{}) {
	if err := j.validateSetDriftControlEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driftControlEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetDrtmProtectionEnabled(val interface{}) {
	if err := j.validateSetDrtmProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"drtmProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetEpisodicDataUploadEnabled(val interface{}) {
	if err := j.validateSetEpisodicDataUploadEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"episodicDataUploadEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetEuLocationEnabled(val interface{}) {
	if err := j.validateSetEuLocationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"euLocationEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetHvciProtectionEnabled(val interface{}) {
	if err := j.validateSetHvciProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hvciProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetSecretsLocation(val *string) {
	if err := j.validateSetSecretsLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsLocation",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetSideChannelMitigationEnabled(val interface{}) {
	if err := j.validateSetSideChannelMitigationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sideChannelMitigationEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetSmbClusterEncryptionEnabled(val interface{}) {
	if err := j.validateSetSmbClusterEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbClusterEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetSmbSigningEnabled(val interface{}) {
	if err := j.validateSetSmbSigningEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbSigningEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetStreamingDataClientEnabled(val interface{}) {
	if err := j.validateSetStreamingDataClientEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamingDataClientEnabled",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference)SetWdacEnabled(val interface{}) {
	if err := j.validateSetWdacEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wdacEnabled",
		val,
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) PutCluster(value *StackHciDeploymentSettingScaleUnitCluster) {
	if err := s.validatePutClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCluster",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) PutHostNetwork(value *StackHciDeploymentSettingScaleUnitHostNetwork) {
	if err := s.validatePutHostNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putHostNetwork",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) PutInfrastructureNetwork(value interface{}) {
	if err := s.validatePutInfrastructureNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInfrastructureNetwork",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) PutOptionalService(value *StackHciDeploymentSettingScaleUnitOptionalService) {
	if err := s.validatePutOptionalServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putOptionalService",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) PutPhysicalNode(value interface{}) {
	if err := s.validatePutPhysicalNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPhysicalNode",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) PutStorage(value *StackHciDeploymentSettingScaleUnitStorage) {
	if err := s.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStorage",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ResetBitlockerBootVolumeEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetBitlockerBootVolumeEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ResetBitlockerDataVolumeEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetBitlockerDataVolumeEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ResetCredentialGuardEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetCredentialGuardEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ResetDriftControlEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetDriftControlEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ResetDrtmProtectionEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetDrtmProtectionEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ResetEpisodicDataUploadEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEpisodicDataUploadEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ResetEuLocationEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEuLocationEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ResetHvciProtectionEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetHvciProtectionEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ResetSideChannelMitigationEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSideChannelMitigationEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ResetSmbClusterEncryptionEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSmbClusterEncryptionEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ResetSmbSigningEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSmbSigningEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ResetStreamingDataClientEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetStreamingDataClientEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ResetWdacEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetWdacEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StackHciDeploymentSettingScaleUnitOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

