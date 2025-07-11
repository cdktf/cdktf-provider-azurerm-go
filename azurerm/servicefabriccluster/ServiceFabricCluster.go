// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabriccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/servicefabriccluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/service_fabric_cluster azurerm_service_fabric_cluster}.
type ServiceFabricCluster interface {
	cdktf.TerraformResource
	AddOnFeatures() *[]*string
	SetAddOnFeatures(val *[]*string)
	AddOnFeaturesInput() *[]*string
	AzureActiveDirectory() ServiceFabricClusterAzureActiveDirectoryOutputReference
	AzureActiveDirectoryInput() *ServiceFabricClusterAzureActiveDirectory
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificate() ServiceFabricClusterCertificateOutputReference
	CertificateCommonNames() ServiceFabricClusterCertificateCommonNamesOutputReference
	CertificateCommonNamesInput() *ServiceFabricClusterCertificateCommonNames
	CertificateInput() *ServiceFabricClusterCertificate
	ClientCertificateCommonName() ServiceFabricClusterClientCertificateCommonNameList
	ClientCertificateCommonNameInput() interface{}
	ClientCertificateThumbprint() ServiceFabricClusterClientCertificateThumbprintList
	ClientCertificateThumbprintInput() interface{}
	ClusterCodeVersion() *string
	SetClusterCodeVersion(val *string)
	ClusterCodeVersionInput() *string
	ClusterEndpoint() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiagnosticsConfig() ServiceFabricClusterDiagnosticsConfigOutputReference
	DiagnosticsConfigInput() *ServiceFabricClusterDiagnosticsConfig
	FabricSettings() ServiceFabricClusterFabricSettingsList
	FabricSettingsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagementEndpoint() *string
	SetManagementEndpoint(val *string)
	ManagementEndpointInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeType() ServiceFabricClusterNodeTypeList
	NodeTypeInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReliabilityLevel() *string
	SetReliabilityLevel(val *string)
	ReliabilityLevelInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ReverseProxyCertificate() ServiceFabricClusterReverseProxyCertificateOutputReference
	ReverseProxyCertificateCommonNames() ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference
	ReverseProxyCertificateCommonNamesInput() *ServiceFabricClusterReverseProxyCertificateCommonNames
	ReverseProxyCertificateInput() *ServiceFabricClusterReverseProxyCertificate
	ServiceFabricZonalUpgradeMode() *string
	SetServiceFabricZonalUpgradeMode(val *string)
	ServiceFabricZonalUpgradeModeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ServiceFabricClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpgradeMode() *string
	SetUpgradeMode(val *string)
	UpgradeModeInput() *string
	UpgradePolicy() ServiceFabricClusterUpgradePolicyOutputReference
	UpgradePolicyInput() *ServiceFabricClusterUpgradePolicy
	VmImage() *string
	SetVmImage(val *string)
	VmImageInput() *string
	VmssZonalUpgradeMode() *string
	SetVmssZonalUpgradeMode(val *string)
	VmssZonalUpgradeModeInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAzureActiveDirectory(value *ServiceFabricClusterAzureActiveDirectory)
	PutCertificate(value *ServiceFabricClusterCertificate)
	PutCertificateCommonNames(value *ServiceFabricClusterCertificateCommonNames)
	PutClientCertificateCommonName(value interface{})
	PutClientCertificateThumbprint(value interface{})
	PutDiagnosticsConfig(value *ServiceFabricClusterDiagnosticsConfig)
	PutFabricSettings(value interface{})
	PutNodeType(value interface{})
	PutReverseProxyCertificate(value *ServiceFabricClusterReverseProxyCertificate)
	PutReverseProxyCertificateCommonNames(value *ServiceFabricClusterReverseProxyCertificateCommonNames)
	PutTimeouts(value *ServiceFabricClusterTimeouts)
	PutUpgradePolicy(value *ServiceFabricClusterUpgradePolicy)
	ResetAddOnFeatures()
	ResetAzureActiveDirectory()
	ResetCertificate()
	ResetCertificateCommonNames()
	ResetClientCertificateCommonName()
	ResetClientCertificateThumbprint()
	ResetClusterCodeVersion()
	ResetDiagnosticsConfig()
	ResetFabricSettings()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReverseProxyCertificate()
	ResetReverseProxyCertificateCommonNames()
	ResetServiceFabricZonalUpgradeMode()
	ResetTags()
	ResetTimeouts()
	ResetUpgradePolicy()
	ResetVmssZonalUpgradeMode()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ServiceFabricCluster
type jsiiProxy_ServiceFabricCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServiceFabricCluster) AddOnFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addOnFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) AddOnFeaturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addOnFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) AzureActiveDirectory() ServiceFabricClusterAzureActiveDirectoryOutputReference {
	var returns ServiceFabricClusterAzureActiveDirectoryOutputReference
	_jsii_.Get(
		j,
		"azureActiveDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) AzureActiveDirectoryInput() *ServiceFabricClusterAzureActiveDirectory {
	var returns *ServiceFabricClusterAzureActiveDirectory
	_jsii_.Get(
		j,
		"azureActiveDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) Certificate() ServiceFabricClusterCertificateOutputReference {
	var returns ServiceFabricClusterCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) CertificateCommonNames() ServiceFabricClusterCertificateCommonNamesOutputReference {
	var returns ServiceFabricClusterCertificateCommonNamesOutputReference
	_jsii_.Get(
		j,
		"certificateCommonNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) CertificateCommonNamesInput() *ServiceFabricClusterCertificateCommonNames {
	var returns *ServiceFabricClusterCertificateCommonNames
	_jsii_.Get(
		j,
		"certificateCommonNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) CertificateInput() *ServiceFabricClusterCertificate {
	var returns *ServiceFabricClusterCertificate
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ClientCertificateCommonName() ServiceFabricClusterClientCertificateCommonNameList {
	var returns ServiceFabricClusterClientCertificateCommonNameList
	_jsii_.Get(
		j,
		"clientCertificateCommonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ClientCertificateCommonNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateCommonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ClientCertificateThumbprint() ServiceFabricClusterClientCertificateThumbprintList {
	var returns ServiceFabricClusterClientCertificateThumbprintList
	_jsii_.Get(
		j,
		"clientCertificateThumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ClientCertificateThumbprintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateThumbprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ClusterCodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ClusterCodeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCodeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ClusterEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) DiagnosticsConfig() ServiceFabricClusterDiagnosticsConfigOutputReference {
	var returns ServiceFabricClusterDiagnosticsConfigOutputReference
	_jsii_.Get(
		j,
		"diagnosticsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) DiagnosticsConfigInput() *ServiceFabricClusterDiagnosticsConfig {
	var returns *ServiceFabricClusterDiagnosticsConfig
	_jsii_.Get(
		j,
		"diagnosticsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) FabricSettings() ServiceFabricClusterFabricSettingsList {
	var returns ServiceFabricClusterFabricSettingsList
	_jsii_.Get(
		j,
		"fabricSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) FabricSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fabricSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ManagementEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ManagementEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managementEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) NodeType() ServiceFabricClusterNodeTypeList {
	var returns ServiceFabricClusterNodeTypeList
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) NodeTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ReliabilityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reliabilityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ReliabilityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reliabilityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ReverseProxyCertificate() ServiceFabricClusterReverseProxyCertificateOutputReference {
	var returns ServiceFabricClusterReverseProxyCertificateOutputReference
	_jsii_.Get(
		j,
		"reverseProxyCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ReverseProxyCertificateCommonNames() ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference {
	var returns ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference
	_jsii_.Get(
		j,
		"reverseProxyCertificateCommonNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ReverseProxyCertificateCommonNamesInput() *ServiceFabricClusterReverseProxyCertificateCommonNames {
	var returns *ServiceFabricClusterReverseProxyCertificateCommonNames
	_jsii_.Get(
		j,
		"reverseProxyCertificateCommonNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ReverseProxyCertificateInput() *ServiceFabricClusterReverseProxyCertificate {
	var returns *ServiceFabricClusterReverseProxyCertificate
	_jsii_.Get(
		j,
		"reverseProxyCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ServiceFabricZonalUpgradeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceFabricZonalUpgradeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) ServiceFabricZonalUpgradeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceFabricZonalUpgradeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) Timeouts() ServiceFabricClusterTimeoutsOutputReference {
	var returns ServiceFabricClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) UpgradeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) UpgradeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) UpgradePolicy() ServiceFabricClusterUpgradePolicyOutputReference {
	var returns ServiceFabricClusterUpgradePolicyOutputReference
	_jsii_.Get(
		j,
		"upgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) UpgradePolicyInput() *ServiceFabricClusterUpgradePolicy {
	var returns *ServiceFabricClusterUpgradePolicy
	_jsii_.Get(
		j,
		"upgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) VmImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) VmImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) VmssZonalUpgradeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmssZonalUpgradeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricCluster) VmssZonalUpgradeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmssZonalUpgradeModeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/service_fabric_cluster azurerm_service_fabric_cluster} Resource.
func NewServiceFabricCluster(scope constructs.Construct, id *string, config *ServiceFabricClusterConfig) ServiceFabricCluster {
	_init_.Initialize()

	if err := validateNewServiceFabricClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceFabricCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/service_fabric_cluster azurerm_service_fabric_cluster} Resource.
func NewServiceFabricCluster_Override(s ServiceFabricCluster, scope constructs.Construct, id *string, config *ServiceFabricClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricCluster",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetAddOnFeatures(val *[]*string) {
	if err := j.validateSetAddOnFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addOnFeatures",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetClusterCodeVersion(val *string) {
	if err := j.validateSetClusterCodeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterCodeVersion",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetManagementEndpoint(val *string) {
	if err := j.validateSetManagementEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managementEndpoint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetReliabilityLevel(val *string) {
	if err := j.validateSetReliabilityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reliabilityLevel",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetServiceFabricZonalUpgradeMode(val *string) {
	if err := j.validateSetServiceFabricZonalUpgradeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceFabricZonalUpgradeMode",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetUpgradeMode(val *string) {
	if err := j.validateSetUpgradeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeMode",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetVmImage(val *string) {
	if err := j.validateSetVmImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmImage",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster)SetVmssZonalUpgradeMode(val *string) {
	if err := j.validateSetVmssZonalUpgradeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmssZonalUpgradeMode",
		val,
	)
}

// Generates CDKTF code for importing a ServiceFabricCluster resource upon running "cdktf plan <stack-name>".
func ServiceFabricCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateServiceFabricCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricCluster",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ServiceFabricCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceFabricCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceFabricCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceFabricCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ServiceFabricCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServiceFabricCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServiceFabricCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServiceFabricCluster) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_ServiceFabricCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_ServiceFabricCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_ServiceFabricCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_ServiceFabricCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_ServiceFabricCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_ServiceFabricCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_ServiceFabricCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_ServiceFabricCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_ServiceFabricCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricCluster) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutAzureActiveDirectory(value *ServiceFabricClusterAzureActiveDirectory) {
	if err := s.validatePutAzureActiveDirectoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAzureActiveDirectory",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutCertificate(value *ServiceFabricClusterCertificate) {
	if err := s.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCertificate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutCertificateCommonNames(value *ServiceFabricClusterCertificateCommonNames) {
	if err := s.validatePutCertificateCommonNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCertificateCommonNames",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutClientCertificateCommonName(value interface{}) {
	if err := s.validatePutClientCertificateCommonNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putClientCertificateCommonName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutClientCertificateThumbprint(value interface{}) {
	if err := s.validatePutClientCertificateThumbprintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putClientCertificateThumbprint",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutDiagnosticsConfig(value *ServiceFabricClusterDiagnosticsConfig) {
	if err := s.validatePutDiagnosticsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDiagnosticsConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutFabricSettings(value interface{}) {
	if err := s.validatePutFabricSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putFabricSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutNodeType(value interface{}) {
	if err := s.validatePutNodeTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putNodeType",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutReverseProxyCertificate(value *ServiceFabricClusterReverseProxyCertificate) {
	if err := s.validatePutReverseProxyCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReverseProxyCertificate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutReverseProxyCertificateCommonNames(value *ServiceFabricClusterReverseProxyCertificateCommonNames) {
	if err := s.validatePutReverseProxyCertificateCommonNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putReverseProxyCertificateCommonNames",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutTimeouts(value *ServiceFabricClusterTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutUpgradePolicy(value *ServiceFabricClusterUpgradePolicy) {
	if err := s.validatePutUpgradePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putUpgradePolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetAddOnFeatures() {
	_jsii_.InvokeVoid(
		s,
		"resetAddOnFeatures",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetAzureActiveDirectory() {
	_jsii_.InvokeVoid(
		s,
		"resetAzureActiveDirectory",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetCertificate() {
	_jsii_.InvokeVoid(
		s,
		"resetCertificate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetCertificateCommonNames() {
	_jsii_.InvokeVoid(
		s,
		"resetCertificateCommonNames",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetClientCertificateCommonName() {
	_jsii_.InvokeVoid(
		s,
		"resetClientCertificateCommonName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetClientCertificateThumbprint() {
	_jsii_.InvokeVoid(
		s,
		"resetClientCertificateThumbprint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetClusterCodeVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetClusterCodeVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetDiagnosticsConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetDiagnosticsConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetFabricSettings() {
	_jsii_.InvokeVoid(
		s,
		"resetFabricSettings",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetReverseProxyCertificate() {
	_jsii_.InvokeVoid(
		s,
		"resetReverseProxyCertificate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetReverseProxyCertificateCommonNames() {
	_jsii_.InvokeVoid(
		s,
		"resetReverseProxyCertificateCommonNames",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetServiceFabricZonalUpgradeMode() {
	_jsii_.InvokeVoid(
		s,
		"resetServiceFabricZonalUpgradeMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetUpgradePolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetUpgradePolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) ResetVmssZonalUpgradeMode() {
	_jsii_.InvokeVoid(
		s,
		"resetVmssZonalUpgradeMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

