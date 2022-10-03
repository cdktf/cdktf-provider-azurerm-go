package servicefabriccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/servicefabriccluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster azurerm_service_fabric_cluster}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
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

func (j *jsiiProxy_ServiceFabricCluster) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster azurerm_service_fabric_cluster} Resource.
func NewServiceFabricCluster(scope constructs.Construct, id *string, config *ServiceFabricClusterConfig) ServiceFabricCluster {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster azurerm_service_fabric_cluster} Resource.
func NewServiceFabricCluster_Override(s ServiceFabricCluster, scope constructs.Construct, id *string, config *ServiceFabricClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricCluster",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetAddOnFeatures(val *[]*string) {
	_jsii_.Set(
		j,
		"addOnFeatures",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetClusterCodeVersion(val *string) {
	_jsii_.Set(
		j,
		"clusterCodeVersion",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetManagementEndpoint(val *string) {
	_jsii_.Set(
		j,
		"managementEndpoint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetReliabilityLevel(val *string) {
	_jsii_.Set(
		j,
		"reliabilityLevel",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetServiceFabricZonalUpgradeMode(val *string) {
	_jsii_.Set(
		j,
		"serviceFabricZonalUpgradeMode",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetUpgradeMode(val *string) {
	_jsii_.Set(
		j,
		"upgradeMode",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetVmImage(val *string) {
	_jsii_.Set(
		j,
		"vmImage",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricCluster) SetVmssZonalUpgradeMode(val *string) {
	_jsii_.Set(
		j,
		"vmssZonalUpgradeMode",
		val,
	)
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricCluster",
		"isConstruct",
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

func (s *jsiiProxy_ServiceFabricCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutAzureActiveDirectory(value *ServiceFabricClusterAzureActiveDirectory) {
	_jsii_.InvokeVoid(
		s,
		"putAzureActiveDirectory",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutCertificate(value *ServiceFabricClusterCertificate) {
	_jsii_.InvokeVoid(
		s,
		"putCertificate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutCertificateCommonNames(value *ServiceFabricClusterCertificateCommonNames) {
	_jsii_.InvokeVoid(
		s,
		"putCertificateCommonNames",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutClientCertificateCommonName(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putClientCertificateCommonName",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutClientCertificateThumbprint(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putClientCertificateThumbprint",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutDiagnosticsConfig(value *ServiceFabricClusterDiagnosticsConfig) {
	_jsii_.InvokeVoid(
		s,
		"putDiagnosticsConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutFabricSettings(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putFabricSettings",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutNodeType(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putNodeType",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutReverseProxyCertificate(value *ServiceFabricClusterReverseProxyCertificate) {
	_jsii_.InvokeVoid(
		s,
		"putReverseProxyCertificate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutReverseProxyCertificateCommonNames(value *ServiceFabricClusterReverseProxyCertificateCommonNames) {
	_jsii_.InvokeVoid(
		s,
		"putReverseProxyCertificateCommonNames",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutTimeouts(value *ServiceFabricClusterTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricCluster) PutUpgradePolicy(value *ServiceFabricClusterUpgradePolicy) {
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

type ServiceFabricClusterAzureActiveDirectory struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#client_application_id ServiceFabricCluster#client_application_id}.
	ClientApplicationId *string `field:"required" json:"clientApplicationId" yaml:"clientApplicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#cluster_application_id ServiceFabricCluster#cluster_application_id}.
	ClusterApplicationId *string `field:"required" json:"clusterApplicationId" yaml:"clusterApplicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#tenant_id ServiceFabricCluster#tenant_id}.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

type ServiceFabricClusterAzureActiveDirectoryOutputReference interface {
	cdktf.ComplexObject
	ClientApplicationId() *string
	SetClientApplicationId(val *string)
	ClientApplicationIdInput() *string
	ClusterApplicationId() *string
	SetClusterApplicationId(val *string)
	ClusterApplicationIdInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *ServiceFabricClusterAzureActiveDirectory
	SetInternalValue(val *ServiceFabricClusterAzureActiveDirectory)
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterAzureActiveDirectoryOutputReference
type jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) ClientApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) ClientApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientApplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) ClusterApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) ClusterApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterApplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) InternalValue() *ServiceFabricClusterAzureActiveDirectory {
	var returns *ServiceFabricClusterAzureActiveDirectory
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterAzureActiveDirectoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterAzureActiveDirectoryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterAzureActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterAzureActiveDirectoryOutputReference_Override(s ServiceFabricClusterAzureActiveDirectoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterAzureActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) SetClientApplicationId(val *string) {
	_jsii_.Set(
		j,
		"clientApplicationId",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) SetClusterApplicationId(val *string) {
	_jsii_.Set(
		j,
		"clusterApplicationId",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) SetInternalValue(val *ServiceFabricClusterAzureActiveDirectory) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterAzureActiveDirectoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#thumbprint ServiceFabricCluster#thumbprint}.
	Thumbprint *string `field:"required" json:"thumbprint" yaml:"thumbprint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#x509_store_name ServiceFabricCluster#x509_store_name}.
	X509StoreName *string `field:"required" json:"x509StoreName" yaml:"x509StoreName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#thumbprint_secondary ServiceFabricCluster#thumbprint_secondary}.
	ThumbprintSecondary *string `field:"optional" json:"thumbprintSecondary" yaml:"thumbprintSecondary"`
}

type ServiceFabricClusterCertificateCommonNames struct {
	// common_names block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#common_names ServiceFabricCluster#common_names}
	CommonNames interface{} `field:"required" json:"commonNames" yaml:"commonNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#x509_store_name ServiceFabricCluster#x509_store_name}.
	X509StoreName *string `field:"required" json:"x509StoreName" yaml:"x509StoreName"`
}

type ServiceFabricClusterCertificateCommonNamesCommonNames struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#certificate_common_name ServiceFabricCluster#certificate_common_name}.
	CertificateCommonName *string `field:"required" json:"certificateCommonName" yaml:"certificateCommonName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#certificate_issuer_thumbprint ServiceFabricCluster#certificate_issuer_thumbprint}.
	CertificateIssuerThumbprint *string `field:"optional" json:"certificateIssuerThumbprint" yaml:"certificateIssuerThumbprint"`
}

type ServiceFabricClusterCertificateCommonNamesCommonNamesList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterCertificateCommonNamesCommonNamesList
type jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterCertificateCommonNamesCommonNamesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ServiceFabricClusterCertificateCommonNamesCommonNamesList {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterCertificateCommonNamesCommonNamesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewServiceFabricClusterCertificateCommonNamesCommonNamesList_Override(s ServiceFabricClusterCertificateCommonNamesCommonNamesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterCertificateCommonNamesCommonNamesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) Get(index *float64) ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference {
	var returns ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference interface {
	cdktf.ComplexObject
	CertificateCommonName() *string
	SetCertificateCommonName(val *string)
	CertificateCommonNameInput() *string
	CertificateIssuerThumbprint() *string
	SetCertificateIssuerThumbprint(val *string)
	CertificateIssuerThumbprintInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetCertificateIssuerThumbprint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference
type jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) CertificateCommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateCommonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) CertificateCommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateCommonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) CertificateIssuerThumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateIssuerThumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) CertificateIssuerThumbprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateIssuerThumbprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference_Override(s ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) SetCertificateCommonName(val *string) {
	_jsii_.Set(
		j,
		"certificateCommonName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) SetCertificateIssuerThumbprint(val *string) {
	_jsii_.Set(
		j,
		"certificateIssuerThumbprint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) ResetCertificateIssuerThumbprint() {
	_jsii_.InvokeVoid(
		s,
		"resetCertificateIssuerThumbprint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesCommonNamesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterCertificateCommonNamesOutputReference interface {
	cdktf.ComplexObject
	CommonNames() ServiceFabricClusterCertificateCommonNamesCommonNamesList
	CommonNamesInput() interface{}
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
	// Experimental.
	Fqn() *string
	InternalValue() *ServiceFabricClusterCertificateCommonNames
	SetInternalValue(val *ServiceFabricClusterCertificateCommonNames)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	X509StoreName() *string
	SetX509StoreName(val *string)
	X509StoreNameInput() *string
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
	PutCommonNames(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterCertificateCommonNamesOutputReference
type jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) CommonNames() ServiceFabricClusterCertificateCommonNamesCommonNamesList {
	var returns ServiceFabricClusterCertificateCommonNamesCommonNamesList
	_jsii_.Get(
		j,
		"commonNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) CommonNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commonNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) InternalValue() *ServiceFabricClusterCertificateCommonNames {
	var returns *ServiceFabricClusterCertificateCommonNames
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) X509StoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"x509StoreName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) X509StoreNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"x509StoreNameInput",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterCertificateCommonNamesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterCertificateCommonNamesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterCertificateCommonNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterCertificateCommonNamesOutputReference_Override(s ServiceFabricClusterCertificateCommonNamesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterCertificateCommonNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) SetInternalValue(val *ServiceFabricClusterCertificateCommonNames) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) SetX509StoreName(val *string) {
	_jsii_.Set(
		j,
		"x509StoreName",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) PutCommonNames(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putCommonNames",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateCommonNamesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterCertificateOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *ServiceFabricClusterCertificate
	SetInternalValue(val *ServiceFabricClusterCertificate)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
	SetThumbprint(val *string)
	ThumbprintInput() *string
	ThumbprintSecondary() *string
	SetThumbprintSecondary(val *string)
	ThumbprintSecondaryInput() *string
	X509StoreName() *string
	SetX509StoreName(val *string)
	X509StoreNameInput() *string
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
	ResetThumbprintSecondary()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterCertificateOutputReference
type jsiiProxy_ServiceFabricClusterCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) InternalValue() *ServiceFabricClusterCertificate {
	var returns *ServiceFabricClusterCertificate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) ThumbprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) ThumbprintSecondary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprintSecondary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) ThumbprintSecondaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprintSecondaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) X509StoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"x509StoreName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) X509StoreNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"x509StoreNameInput",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterCertificateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterCertificateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterCertificateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterCertificateOutputReference_Override(s ServiceFabricClusterCertificateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) SetInternalValue(val *ServiceFabricClusterCertificate) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) SetThumbprint(val *string) {
	_jsii_.Set(
		j,
		"thumbprint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) SetThumbprintSecondary(val *string) {
	_jsii_.Set(
		j,
		"thumbprintSecondary",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterCertificateOutputReference) SetX509StoreName(val *string) {
	_jsii_.Set(
		j,
		"x509StoreName",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) ResetThumbprintSecondary() {
	_jsii_.InvokeVoid(
		s,
		"resetThumbprintSecondary",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterCertificateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterClientCertificateCommonName struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#common_name ServiceFabricCluster#common_name}.
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#is_admin ServiceFabricCluster#is_admin}.
	IsAdmin interface{} `field:"required" json:"isAdmin" yaml:"isAdmin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#issuer_thumbprint ServiceFabricCluster#issuer_thumbprint}.
	IssuerThumbprint *string `field:"optional" json:"issuerThumbprint" yaml:"issuerThumbprint"`
}

type ServiceFabricClusterClientCertificateCommonNameList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ServiceFabricClusterClientCertificateCommonNameOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterClientCertificateCommonNameList
type jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterClientCertificateCommonNameList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ServiceFabricClusterClientCertificateCommonNameList {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterClientCertificateCommonNameList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewServiceFabricClusterClientCertificateCommonNameList_Override(s ServiceFabricClusterClientCertificateCommonNameList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterClientCertificateCommonNameList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) Get(index *float64) ServiceFabricClusterClientCertificateCommonNameOutputReference {
	var returns ServiceFabricClusterClientCertificateCommonNameOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterClientCertificateCommonNameOutputReference interface {
	cdktf.ComplexObject
	CommonName() *string
	SetCommonName(val *string)
	CommonNameInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsAdmin() interface{}
	SetIsAdmin(val interface{})
	IsAdminInput() interface{}
	IssuerThumbprint() *string
	SetIssuerThumbprint(val *string)
	IssuerThumbprintInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetIssuerThumbprint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterClientCertificateCommonNameOutputReference
type jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) IsAdmin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) IsAdminInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAdminInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) IssuerThumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerThumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) IssuerThumbprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerThumbprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterClientCertificateCommonNameOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricClusterClientCertificateCommonNameOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterClientCertificateCommonNameOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricClusterClientCertificateCommonNameOutputReference_Override(s ServiceFabricClusterClientCertificateCommonNameOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterClientCertificateCommonNameOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) SetCommonName(val *string) {
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) SetIsAdmin(val interface{}) {
	_jsii_.Set(
		j,
		"isAdmin",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) SetIssuerThumbprint(val *string) {
	_jsii_.Set(
		j,
		"issuerThumbprint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) ResetIssuerThumbprint() {
	_jsii_.InvokeVoid(
		s,
		"resetIssuerThumbprint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateCommonNameOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterClientCertificateThumbprint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#is_admin ServiceFabricCluster#is_admin}.
	IsAdmin interface{} `field:"required" json:"isAdmin" yaml:"isAdmin"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#thumbprint ServiceFabricCluster#thumbprint}.
	Thumbprint *string `field:"required" json:"thumbprint" yaml:"thumbprint"`
}

type ServiceFabricClusterClientCertificateThumbprintList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ServiceFabricClusterClientCertificateThumbprintOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterClientCertificateThumbprintList
type jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterClientCertificateThumbprintList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ServiceFabricClusterClientCertificateThumbprintList {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterClientCertificateThumbprintList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewServiceFabricClusterClientCertificateThumbprintList_Override(s ServiceFabricClusterClientCertificateThumbprintList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterClientCertificateThumbprintList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) Get(index *float64) ServiceFabricClusterClientCertificateThumbprintOutputReference {
	var returns ServiceFabricClusterClientCertificateThumbprintOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterClientCertificateThumbprintOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsAdmin() interface{}
	SetIsAdmin(val interface{})
	IsAdminInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
	SetThumbprint(val *string)
	ThumbprintInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterClientCertificateThumbprintOutputReference
type jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) IsAdmin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) IsAdminInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAdminInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) ThumbprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprintInput",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterClientCertificateThumbprintOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricClusterClientCertificateThumbprintOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterClientCertificateThumbprintOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricClusterClientCertificateThumbprintOutputReference_Override(s ServiceFabricClusterClientCertificateThumbprintOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterClientCertificateThumbprintOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) SetIsAdmin(val interface{}) {
	_jsii_.Set(
		j,
		"isAdmin",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) SetThumbprint(val *string) {
	_jsii_.Set(
		j,
		"thumbprint",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterClientCertificateThumbprintOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#location ServiceFabricCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#management_endpoint ServiceFabricCluster#management_endpoint}.
	ManagementEndpoint *string `field:"required" json:"managementEndpoint" yaml:"managementEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#name ServiceFabricCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// node_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#node_type ServiceFabricCluster#node_type}
	NodeType interface{} `field:"required" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#reliability_level ServiceFabricCluster#reliability_level}.
	ReliabilityLevel *string `field:"required" json:"reliabilityLevel" yaml:"reliabilityLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#resource_group_name ServiceFabricCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#upgrade_mode ServiceFabricCluster#upgrade_mode}.
	UpgradeMode *string `field:"required" json:"upgradeMode" yaml:"upgradeMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#vm_image ServiceFabricCluster#vm_image}.
	VmImage *string `field:"required" json:"vmImage" yaml:"vmImage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#add_on_features ServiceFabricCluster#add_on_features}.
	AddOnFeatures *[]*string `field:"optional" json:"addOnFeatures" yaml:"addOnFeatures"`
	// azure_active_directory block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#azure_active_directory ServiceFabricCluster#azure_active_directory}
	AzureActiveDirectory *ServiceFabricClusterAzureActiveDirectory `field:"optional" json:"azureActiveDirectory" yaml:"azureActiveDirectory"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#certificate ServiceFabricCluster#certificate}
	Certificate *ServiceFabricClusterCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// certificate_common_names block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#certificate_common_names ServiceFabricCluster#certificate_common_names}
	CertificateCommonNames *ServiceFabricClusterCertificateCommonNames `field:"optional" json:"certificateCommonNames" yaml:"certificateCommonNames"`
	// client_certificate_common_name block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#client_certificate_common_name ServiceFabricCluster#client_certificate_common_name}
	ClientCertificateCommonName interface{} `field:"optional" json:"clientCertificateCommonName" yaml:"clientCertificateCommonName"`
	// client_certificate_thumbprint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#client_certificate_thumbprint ServiceFabricCluster#client_certificate_thumbprint}
	ClientCertificateThumbprint interface{} `field:"optional" json:"clientCertificateThumbprint" yaml:"clientCertificateThumbprint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#cluster_code_version ServiceFabricCluster#cluster_code_version}.
	ClusterCodeVersion *string `field:"optional" json:"clusterCodeVersion" yaml:"clusterCodeVersion"`
	// diagnostics_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#diagnostics_config ServiceFabricCluster#diagnostics_config}
	DiagnosticsConfig *ServiceFabricClusterDiagnosticsConfig `field:"optional" json:"diagnosticsConfig" yaml:"diagnosticsConfig"`
	// fabric_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#fabric_settings ServiceFabricCluster#fabric_settings}
	FabricSettings interface{} `field:"optional" json:"fabricSettings" yaml:"fabricSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#id ServiceFabricCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// reverse_proxy_certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#reverse_proxy_certificate ServiceFabricCluster#reverse_proxy_certificate}
	ReverseProxyCertificate *ServiceFabricClusterReverseProxyCertificate `field:"optional" json:"reverseProxyCertificate" yaml:"reverseProxyCertificate"`
	// reverse_proxy_certificate_common_names block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#reverse_proxy_certificate_common_names ServiceFabricCluster#reverse_proxy_certificate_common_names}
	ReverseProxyCertificateCommonNames *ServiceFabricClusterReverseProxyCertificateCommonNames `field:"optional" json:"reverseProxyCertificateCommonNames" yaml:"reverseProxyCertificateCommonNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#service_fabric_zonal_upgrade_mode ServiceFabricCluster#service_fabric_zonal_upgrade_mode}.
	ServiceFabricZonalUpgradeMode *string `field:"optional" json:"serviceFabricZonalUpgradeMode" yaml:"serviceFabricZonalUpgradeMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#tags ServiceFabricCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#timeouts ServiceFabricCluster#timeouts}
	Timeouts *ServiceFabricClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// upgrade_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#upgrade_policy ServiceFabricCluster#upgrade_policy}
	UpgradePolicy *ServiceFabricClusterUpgradePolicy `field:"optional" json:"upgradePolicy" yaml:"upgradePolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#vmss_zonal_upgrade_mode ServiceFabricCluster#vmss_zonal_upgrade_mode}.
	VmssZonalUpgradeMode *string `field:"optional" json:"vmssZonalUpgradeMode" yaml:"vmssZonalUpgradeMode"`
}

type ServiceFabricClusterDiagnosticsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#blob_endpoint ServiceFabricCluster#blob_endpoint}.
	BlobEndpoint *string `field:"required" json:"blobEndpoint" yaml:"blobEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#protected_account_key_name ServiceFabricCluster#protected_account_key_name}.
	ProtectedAccountKeyName *string `field:"required" json:"protectedAccountKeyName" yaml:"protectedAccountKeyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#queue_endpoint ServiceFabricCluster#queue_endpoint}.
	QueueEndpoint *string `field:"required" json:"queueEndpoint" yaml:"queueEndpoint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#storage_account_name ServiceFabricCluster#storage_account_name}.
	StorageAccountName *string `field:"required" json:"storageAccountName" yaml:"storageAccountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#table_endpoint ServiceFabricCluster#table_endpoint}.
	TableEndpoint *string `field:"required" json:"tableEndpoint" yaml:"tableEndpoint"`
}

type ServiceFabricClusterDiagnosticsConfigOutputReference interface {
	cdktf.ComplexObject
	BlobEndpoint() *string
	SetBlobEndpoint(val *string)
	BlobEndpointInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *ServiceFabricClusterDiagnosticsConfig
	SetInternalValue(val *ServiceFabricClusterDiagnosticsConfig)
	ProtectedAccountKeyName() *string
	SetProtectedAccountKeyName(val *string)
	ProtectedAccountKeyNameInput() *string
	QueueEndpoint() *string
	SetQueueEndpoint(val *string)
	QueueEndpointInput() *string
	StorageAccountName() *string
	SetStorageAccountName(val *string)
	StorageAccountNameInput() *string
	TableEndpoint() *string
	SetTableEndpoint(val *string)
	TableEndpointInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterDiagnosticsConfigOutputReference
type jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) BlobEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) BlobEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) InternalValue() *ServiceFabricClusterDiagnosticsConfig {
	var returns *ServiceFabricClusterDiagnosticsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) ProtectedAccountKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectedAccountKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) ProtectedAccountKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protectedAccountKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) QueueEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) QueueEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) StorageAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) StorageAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) TableEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) TableEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterDiagnosticsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterDiagnosticsConfigOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterDiagnosticsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterDiagnosticsConfigOutputReference_Override(s ServiceFabricClusterDiagnosticsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterDiagnosticsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) SetBlobEndpoint(val *string) {
	_jsii_.Set(
		j,
		"blobEndpoint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) SetInternalValue(val *ServiceFabricClusterDiagnosticsConfig) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) SetProtectedAccountKeyName(val *string) {
	_jsii_.Set(
		j,
		"protectedAccountKeyName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) SetQueueEndpoint(val *string) {
	_jsii_.Set(
		j,
		"queueEndpoint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) SetStorageAccountName(val *string) {
	_jsii_.Set(
		j,
		"storageAccountName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) SetTableEndpoint(val *string) {
	_jsii_.Set(
		j,
		"tableEndpoint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterDiagnosticsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterFabricSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#name ServiceFabricCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#parameters ServiceFabricCluster#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

type ServiceFabricClusterFabricSettingsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ServiceFabricClusterFabricSettingsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterFabricSettingsList
type jsiiProxy_ServiceFabricClusterFabricSettingsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterFabricSettingsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ServiceFabricClusterFabricSettingsList {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterFabricSettingsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterFabricSettingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewServiceFabricClusterFabricSettingsList_Override(s ServiceFabricClusterFabricSettingsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterFabricSettingsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsList) Get(index *float64) ServiceFabricClusterFabricSettingsOutputReference {
	var returns ServiceFabricClusterFabricSettingsOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterFabricSettingsOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterFabricSettingsOutputReference
type jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterFabricSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricClusterFabricSettingsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterFabricSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricClusterFabricSettingsOutputReference_Override(s ServiceFabricClusterFabricSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterFabricSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) SetParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		s,
		"resetParameters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterFabricSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterNodeType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#client_endpoint_port ServiceFabricCluster#client_endpoint_port}.
	ClientEndpointPort *float64 `field:"required" json:"clientEndpointPort" yaml:"clientEndpointPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#http_endpoint_port ServiceFabricCluster#http_endpoint_port}.
	HttpEndpointPort *float64 `field:"required" json:"httpEndpointPort" yaml:"httpEndpointPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#instance_count ServiceFabricCluster#instance_count}.
	InstanceCount *float64 `field:"required" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#is_primary ServiceFabricCluster#is_primary}.
	IsPrimary interface{} `field:"required" json:"isPrimary" yaml:"isPrimary"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#name ServiceFabricCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// application_ports block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#application_ports ServiceFabricCluster#application_ports}
	ApplicationPorts *ServiceFabricClusterNodeTypeApplicationPorts `field:"optional" json:"applicationPorts" yaml:"applicationPorts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#capacities ServiceFabricCluster#capacities}.
	Capacities *map[string]*string `field:"optional" json:"capacities" yaml:"capacities"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#durability_level ServiceFabricCluster#durability_level}.
	DurabilityLevel *string `field:"optional" json:"durabilityLevel" yaml:"durabilityLevel"`
	// ephemeral_ports block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#ephemeral_ports ServiceFabricCluster#ephemeral_ports}
	EphemeralPorts *ServiceFabricClusterNodeTypeEphemeralPorts `field:"optional" json:"ephemeralPorts" yaml:"ephemeralPorts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#is_stateless ServiceFabricCluster#is_stateless}.
	IsStateless interface{} `field:"optional" json:"isStateless" yaml:"isStateless"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#multiple_availability_zones ServiceFabricCluster#multiple_availability_zones}.
	MultipleAvailabilityZones interface{} `field:"optional" json:"multipleAvailabilityZones" yaml:"multipleAvailabilityZones"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#placement_properties ServiceFabricCluster#placement_properties}.
	PlacementProperties *map[string]*string `field:"optional" json:"placementProperties" yaml:"placementProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#reverse_proxy_endpoint_port ServiceFabricCluster#reverse_proxy_endpoint_port}.
	ReverseProxyEndpointPort *float64 `field:"optional" json:"reverseProxyEndpointPort" yaml:"reverseProxyEndpointPort"`
}

type ServiceFabricClusterNodeTypeApplicationPorts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#end_port ServiceFabricCluster#end_port}.
	EndPort *float64 `field:"required" json:"endPort" yaml:"endPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#start_port ServiceFabricCluster#start_port}.
	StartPort *float64 `field:"required" json:"startPort" yaml:"startPort"`
}

type ServiceFabricClusterNodeTypeApplicationPortsOutputReference interface {
	cdktf.ComplexObject
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
	EndPort() *float64
	SetEndPort(val *float64)
	EndPortInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *ServiceFabricClusterNodeTypeApplicationPorts
	SetInternalValue(val *ServiceFabricClusterNodeTypeApplicationPorts)
	StartPort() *float64
	SetStartPort(val *float64)
	StartPortInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterNodeTypeApplicationPortsOutputReference
type jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) EndPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) EndPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) InternalValue() *ServiceFabricClusterNodeTypeApplicationPorts {
	var returns *ServiceFabricClusterNodeTypeApplicationPorts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) StartPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) StartPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterNodeTypeApplicationPortsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterNodeTypeApplicationPortsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterNodeTypeApplicationPortsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterNodeTypeApplicationPortsOutputReference_Override(s ServiceFabricClusterNodeTypeApplicationPortsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterNodeTypeApplicationPortsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) SetEndPort(val *float64) {
	_jsii_.Set(
		j,
		"endPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) SetInternalValue(val *ServiceFabricClusterNodeTypeApplicationPorts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) SetStartPort(val *float64) {
	_jsii_.Set(
		j,
		"startPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeApplicationPortsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterNodeTypeEphemeralPorts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#end_port ServiceFabricCluster#end_port}.
	EndPort *float64 `field:"required" json:"endPort" yaml:"endPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#start_port ServiceFabricCluster#start_port}.
	StartPort *float64 `field:"required" json:"startPort" yaml:"startPort"`
}

type ServiceFabricClusterNodeTypeEphemeralPortsOutputReference interface {
	cdktf.ComplexObject
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
	EndPort() *float64
	SetEndPort(val *float64)
	EndPortInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *ServiceFabricClusterNodeTypeEphemeralPorts
	SetInternalValue(val *ServiceFabricClusterNodeTypeEphemeralPorts)
	StartPort() *float64
	SetStartPort(val *float64)
	StartPortInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterNodeTypeEphemeralPortsOutputReference
type jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) EndPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) EndPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) InternalValue() *ServiceFabricClusterNodeTypeEphemeralPorts {
	var returns *ServiceFabricClusterNodeTypeEphemeralPorts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) StartPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) StartPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterNodeTypeEphemeralPortsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterNodeTypeEphemeralPortsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterNodeTypeEphemeralPortsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterNodeTypeEphemeralPortsOutputReference_Override(s ServiceFabricClusterNodeTypeEphemeralPortsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterNodeTypeEphemeralPortsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) SetEndPort(val *float64) {
	_jsii_.Set(
		j,
		"endPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) SetInternalValue(val *ServiceFabricClusterNodeTypeEphemeralPorts) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) SetStartPort(val *float64) {
	_jsii_.Set(
		j,
		"startPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeEphemeralPortsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterNodeTypeList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ServiceFabricClusterNodeTypeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterNodeTypeList
type jsiiProxy_ServiceFabricClusterNodeTypeList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterNodeTypeList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ServiceFabricClusterNodeTypeList {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterNodeTypeList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterNodeTypeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewServiceFabricClusterNodeTypeList_Override(s ServiceFabricClusterNodeTypeList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterNodeTypeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeList) Get(index *float64) ServiceFabricClusterNodeTypeOutputReference {
	var returns ServiceFabricClusterNodeTypeOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterNodeTypeOutputReference interface {
	cdktf.ComplexObject
	ApplicationPorts() ServiceFabricClusterNodeTypeApplicationPortsOutputReference
	ApplicationPortsInput() *ServiceFabricClusterNodeTypeApplicationPorts
	Capacities() *map[string]*string
	SetCapacities(val *map[string]*string)
	CapacitiesInput() *map[string]*string
	ClientEndpointPort() *float64
	SetClientEndpointPort(val *float64)
	ClientEndpointPortInput() *float64
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
	DurabilityLevel() *string
	SetDurabilityLevel(val *string)
	DurabilityLevelInput() *string
	EphemeralPorts() ServiceFabricClusterNodeTypeEphemeralPortsOutputReference
	EphemeralPortsInput() *ServiceFabricClusterNodeTypeEphemeralPorts
	// Experimental.
	Fqn() *string
	HttpEndpointPort() *float64
	SetHttpEndpointPort(val *float64)
	HttpEndpointPortInput() *float64
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceCountInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsPrimary() interface{}
	SetIsPrimary(val interface{})
	IsPrimaryInput() interface{}
	IsStateless() interface{}
	SetIsStateless(val interface{})
	IsStatelessInput() interface{}
	MultipleAvailabilityZones() interface{}
	SetMultipleAvailabilityZones(val interface{})
	MultipleAvailabilityZonesInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	PlacementProperties() *map[string]*string
	SetPlacementProperties(val *map[string]*string)
	PlacementPropertiesInput() *map[string]*string
	ReverseProxyEndpointPort() *float64
	SetReverseProxyEndpointPort(val *float64)
	ReverseProxyEndpointPortInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutApplicationPorts(value *ServiceFabricClusterNodeTypeApplicationPorts)
	PutEphemeralPorts(value *ServiceFabricClusterNodeTypeEphemeralPorts)
	ResetApplicationPorts()
	ResetCapacities()
	ResetDurabilityLevel()
	ResetEphemeralPorts()
	ResetIsStateless()
	ResetMultipleAvailabilityZones()
	ResetPlacementProperties()
	ResetReverseProxyEndpointPort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterNodeTypeOutputReference
type jsiiProxy_ServiceFabricClusterNodeTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ApplicationPorts() ServiceFabricClusterNodeTypeApplicationPortsOutputReference {
	var returns ServiceFabricClusterNodeTypeApplicationPortsOutputReference
	_jsii_.Get(
		j,
		"applicationPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ApplicationPortsInput() *ServiceFabricClusterNodeTypeApplicationPorts {
	var returns *ServiceFabricClusterNodeTypeApplicationPorts
	_jsii_.Get(
		j,
		"applicationPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) Capacities() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"capacities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) CapacitiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"capacitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ClientEndpointPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ClientEndpointPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientEndpointPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) DurabilityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durabilityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) DurabilityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durabilityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) EphemeralPorts() ServiceFabricClusterNodeTypeEphemeralPortsOutputReference {
	var returns ServiceFabricClusterNodeTypeEphemeralPortsOutputReference
	_jsii_.Get(
		j,
		"ephemeralPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) EphemeralPortsInput() *ServiceFabricClusterNodeTypeEphemeralPorts {
	var returns *ServiceFabricClusterNodeTypeEphemeralPorts
	_jsii_.Get(
		j,
		"ephemeralPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) HttpEndpointPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) HttpEndpointPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpEndpointPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) InstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) IsPrimary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) IsPrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPrimaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) IsStateless() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStateless",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) IsStatelessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStatelessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) MultipleAvailabilityZones() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multipleAvailabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) MultipleAvailabilityZonesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multipleAvailabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) PlacementProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"placementProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) PlacementPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"placementPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ReverseProxyEndpointPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reverseProxyEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ReverseProxyEndpointPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reverseProxyEndpointPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterNodeTypeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricClusterNodeTypeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterNodeTypeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterNodeTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricClusterNodeTypeOutputReference_Override(s ServiceFabricClusterNodeTypeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterNodeTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetCapacities(val *map[string]*string) {
	_jsii_.Set(
		j,
		"capacities",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetClientEndpointPort(val *float64) {
	_jsii_.Set(
		j,
		"clientEndpointPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetDurabilityLevel(val *string) {
	_jsii_.Set(
		j,
		"durabilityLevel",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetHttpEndpointPort(val *float64) {
	_jsii_.Set(
		j,
		"httpEndpointPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetIsPrimary(val interface{}) {
	_jsii_.Set(
		j,
		"isPrimary",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetIsStateless(val interface{}) {
	_jsii_.Set(
		j,
		"isStateless",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetMultipleAvailabilityZones(val interface{}) {
	_jsii_.Set(
		j,
		"multipleAvailabilityZones",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetPlacementProperties(val *map[string]*string) {
	_jsii_.Set(
		j,
		"placementProperties",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetReverseProxyEndpointPort(val *float64) {
	_jsii_.Set(
		j,
		"reverseProxyEndpointPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) PutApplicationPorts(value *ServiceFabricClusterNodeTypeApplicationPorts) {
	_jsii_.InvokeVoid(
		s,
		"putApplicationPorts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) PutEphemeralPorts(value *ServiceFabricClusterNodeTypeEphemeralPorts) {
	_jsii_.InvokeVoid(
		s,
		"putEphemeralPorts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetApplicationPorts() {
	_jsii_.InvokeVoid(
		s,
		"resetApplicationPorts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetCapacities() {
	_jsii_.InvokeVoid(
		s,
		"resetCapacities",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetDurabilityLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetDurabilityLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetEphemeralPorts() {
	_jsii_.InvokeVoid(
		s,
		"resetEphemeralPorts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetIsStateless() {
	_jsii_.InvokeVoid(
		s,
		"resetIsStateless",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetMultipleAvailabilityZones() {
	_jsii_.InvokeVoid(
		s,
		"resetMultipleAvailabilityZones",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetPlacementProperties() {
	_jsii_.InvokeVoid(
		s,
		"resetPlacementProperties",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ResetReverseProxyEndpointPort() {
	_jsii_.InvokeVoid(
		s,
		"resetReverseProxyEndpointPort",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterNodeTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterReverseProxyCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#thumbprint ServiceFabricCluster#thumbprint}.
	Thumbprint *string `field:"required" json:"thumbprint" yaml:"thumbprint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#x509_store_name ServiceFabricCluster#x509_store_name}.
	X509StoreName *string `field:"required" json:"x509StoreName" yaml:"x509StoreName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#thumbprint_secondary ServiceFabricCluster#thumbprint_secondary}.
	ThumbprintSecondary *string `field:"optional" json:"thumbprintSecondary" yaml:"thumbprintSecondary"`
}

type ServiceFabricClusterReverseProxyCertificateCommonNames struct {
	// common_names block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#common_names ServiceFabricCluster#common_names}
	CommonNames interface{} `field:"required" json:"commonNames" yaml:"commonNames"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#x509_store_name ServiceFabricCluster#x509_store_name}.
	X509StoreName *string `field:"required" json:"x509StoreName" yaml:"x509StoreName"`
}

type ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNames struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#certificate_common_name ServiceFabricCluster#certificate_common_name}.
	CertificateCommonName *string `field:"required" json:"certificateCommonName" yaml:"certificateCommonName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#certificate_issuer_thumbprint ServiceFabricCluster#certificate_issuer_thumbprint}.
	CertificateIssuerThumbprint *string `field:"optional" json:"certificateIssuerThumbprint" yaml:"certificateIssuerThumbprint"`
}

type ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList
type jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList_Override(s ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) Get(index *float64) ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference {
	var returns ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference interface {
	cdktf.ComplexObject
	CertificateCommonName() *string
	SetCertificateCommonName(val *string)
	CertificateCommonNameInput() *string
	CertificateIssuerThumbprint() *string
	SetCertificateIssuerThumbprint(val *string)
	CertificateIssuerThumbprintInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetCertificateIssuerThumbprint()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference
type jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) CertificateCommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateCommonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) CertificateCommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateCommonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) CertificateIssuerThumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateIssuerThumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) CertificateIssuerThumbprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateIssuerThumbprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference_Override(s ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) SetCertificateCommonName(val *string) {
	_jsii_.Set(
		j,
		"certificateCommonName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) SetCertificateIssuerThumbprint(val *string) {
	_jsii_.Set(
		j,
		"certificateIssuerThumbprint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) ResetCertificateIssuerThumbprint() {
	_jsii_.InvokeVoid(
		s,
		"resetCertificateIssuerThumbprint",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference interface {
	cdktf.ComplexObject
	CommonNames() ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList
	CommonNamesInput() interface{}
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
	// Experimental.
	Fqn() *string
	InternalValue() *ServiceFabricClusterReverseProxyCertificateCommonNames
	SetInternalValue(val *ServiceFabricClusterReverseProxyCertificateCommonNames)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	X509StoreName() *string
	SetX509StoreName(val *string)
	X509StoreNameInput() *string
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
	PutCommonNames(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference
type jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) CommonNames() ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList {
	var returns ServiceFabricClusterReverseProxyCertificateCommonNamesCommonNamesList
	_jsii_.Get(
		j,
		"commonNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) CommonNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commonNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) InternalValue() *ServiceFabricClusterReverseProxyCertificateCommonNames {
	var returns *ServiceFabricClusterReverseProxyCertificateCommonNames
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) X509StoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"x509StoreName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) X509StoreNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"x509StoreNameInput",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference_Override(s ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) SetInternalValue(val *ServiceFabricClusterReverseProxyCertificateCommonNames) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) SetX509StoreName(val *string) {
	_jsii_.Set(
		j,
		"x509StoreName",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) PutCommonNames(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putCommonNames",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateCommonNamesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterReverseProxyCertificateOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *ServiceFabricClusterReverseProxyCertificate
	SetInternalValue(val *ServiceFabricClusterReverseProxyCertificate)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
	SetThumbprint(val *string)
	ThumbprintInput() *string
	ThumbprintSecondary() *string
	SetThumbprintSecondary(val *string)
	ThumbprintSecondaryInput() *string
	X509StoreName() *string
	SetX509StoreName(val *string)
	X509StoreNameInput() *string
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
	ResetThumbprintSecondary()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterReverseProxyCertificateOutputReference
type jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) InternalValue() *ServiceFabricClusterReverseProxyCertificate {
	var returns *ServiceFabricClusterReverseProxyCertificate
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) ThumbprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) ThumbprintSecondary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprintSecondary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) ThumbprintSecondaryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprintSecondaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) X509StoreName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"x509StoreName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) X509StoreNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"x509StoreNameInput",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterReverseProxyCertificateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterReverseProxyCertificateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterReverseProxyCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterReverseProxyCertificateOutputReference_Override(s ServiceFabricClusterReverseProxyCertificateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterReverseProxyCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) SetInternalValue(val *ServiceFabricClusterReverseProxyCertificate) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) SetThumbprint(val *string) {
	_jsii_.Set(
		j,
		"thumbprint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) SetThumbprintSecondary(val *string) {
	_jsii_.Set(
		j,
		"thumbprintSecondary",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) SetX509StoreName(val *string) {
	_jsii_.Set(
		j,
		"x509StoreName",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) ResetThumbprintSecondary() {
	_jsii_.InvokeVoid(
		s,
		"resetThumbprintSecondary",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterReverseProxyCertificateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#create ServiceFabricCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#delete ServiceFabricCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#read ServiceFabricCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#update ServiceFabricCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type ServiceFabricClusterTimeoutsOutputReference interface {
	cdktf.ComplexObject
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
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
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
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterTimeoutsOutputReference
type jsiiProxy_ServiceFabricClusterTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterTimeoutsOutputReference_Override(s ServiceFabricClusterTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterUpgradePolicy struct {
	// delta_health_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#delta_health_policy ServiceFabricCluster#delta_health_policy}
	DeltaHealthPolicy *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy `field:"optional" json:"deltaHealthPolicy" yaml:"deltaHealthPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#force_restart_enabled ServiceFabricCluster#force_restart_enabled}.
	ForceRestartEnabled interface{} `field:"optional" json:"forceRestartEnabled" yaml:"forceRestartEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#health_check_retry_timeout ServiceFabricCluster#health_check_retry_timeout}.
	HealthCheckRetryTimeout *string `field:"optional" json:"healthCheckRetryTimeout" yaml:"healthCheckRetryTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#health_check_stable_duration ServiceFabricCluster#health_check_stable_duration}.
	HealthCheckStableDuration *string `field:"optional" json:"healthCheckStableDuration" yaml:"healthCheckStableDuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#health_check_wait_duration ServiceFabricCluster#health_check_wait_duration}.
	HealthCheckWaitDuration *string `field:"optional" json:"healthCheckWaitDuration" yaml:"healthCheckWaitDuration"`
	// health_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#health_policy ServiceFabricCluster#health_policy}
	HealthPolicy *ServiceFabricClusterUpgradePolicyHealthPolicy `field:"optional" json:"healthPolicy" yaml:"healthPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#upgrade_domain_timeout ServiceFabricCluster#upgrade_domain_timeout}.
	UpgradeDomainTimeout *string `field:"optional" json:"upgradeDomainTimeout" yaml:"upgradeDomainTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#upgrade_replica_set_check_timeout ServiceFabricCluster#upgrade_replica_set_check_timeout}.
	UpgradeReplicaSetCheckTimeout *string `field:"optional" json:"upgradeReplicaSetCheckTimeout" yaml:"upgradeReplicaSetCheckTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#upgrade_timeout ServiceFabricCluster#upgrade_timeout}.
	UpgradeTimeout *string `field:"optional" json:"upgradeTimeout" yaml:"upgradeTimeout"`
}

type ServiceFabricClusterUpgradePolicyDeltaHealthPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#max_delta_unhealthy_applications_percent ServiceFabricCluster#max_delta_unhealthy_applications_percent}.
	MaxDeltaUnhealthyApplicationsPercent *float64 `field:"optional" json:"maxDeltaUnhealthyApplicationsPercent" yaml:"maxDeltaUnhealthyApplicationsPercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#max_delta_unhealthy_nodes_percent ServiceFabricCluster#max_delta_unhealthy_nodes_percent}.
	MaxDeltaUnhealthyNodesPercent *float64 `field:"optional" json:"maxDeltaUnhealthyNodesPercent" yaml:"maxDeltaUnhealthyNodesPercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#max_upgrade_domain_delta_unhealthy_nodes_percent ServiceFabricCluster#max_upgrade_domain_delta_unhealthy_nodes_percent}.
	MaxUpgradeDomainDeltaUnhealthyNodesPercent *float64 `field:"optional" json:"maxUpgradeDomainDeltaUnhealthyNodesPercent" yaml:"maxUpgradeDomainDeltaUnhealthyNodesPercent"`
}

type ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy
	SetInternalValue(val *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy)
	MaxDeltaUnhealthyApplicationsPercent() *float64
	SetMaxDeltaUnhealthyApplicationsPercent(val *float64)
	MaxDeltaUnhealthyApplicationsPercentInput() *float64
	MaxDeltaUnhealthyNodesPercent() *float64
	SetMaxDeltaUnhealthyNodesPercent(val *float64)
	MaxDeltaUnhealthyNodesPercentInput() *float64
	MaxUpgradeDomainDeltaUnhealthyNodesPercent() *float64
	SetMaxUpgradeDomainDeltaUnhealthyNodesPercent(val *float64)
	MaxUpgradeDomainDeltaUnhealthyNodesPercentInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetMaxDeltaUnhealthyApplicationsPercent()
	ResetMaxDeltaUnhealthyNodesPercent()
	ResetMaxUpgradeDomainDeltaUnhealthyNodesPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference
type jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) InternalValue() *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy {
	var returns *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) MaxDeltaUnhealthyApplicationsPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeltaUnhealthyApplicationsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) MaxDeltaUnhealthyApplicationsPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeltaUnhealthyApplicationsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) MaxDeltaUnhealthyNodesPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeltaUnhealthyNodesPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) MaxDeltaUnhealthyNodesPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeltaUnhealthyNodesPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) MaxUpgradeDomainDeltaUnhealthyNodesPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUpgradeDomainDeltaUnhealthyNodesPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) MaxUpgradeDomainDeltaUnhealthyNodesPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUpgradeDomainDeltaUnhealthyNodesPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference_Override(s ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) SetInternalValue(val *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) SetMaxDeltaUnhealthyApplicationsPercent(val *float64) {
	_jsii_.Set(
		j,
		"maxDeltaUnhealthyApplicationsPercent",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) SetMaxDeltaUnhealthyNodesPercent(val *float64) {
	_jsii_.Set(
		j,
		"maxDeltaUnhealthyNodesPercent",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) SetMaxUpgradeDomainDeltaUnhealthyNodesPercent(val *float64) {
	_jsii_.Set(
		j,
		"maxUpgradeDomainDeltaUnhealthyNodesPercent",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ResetMaxDeltaUnhealthyApplicationsPercent() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxDeltaUnhealthyApplicationsPercent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ResetMaxDeltaUnhealthyNodesPercent() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxDeltaUnhealthyNodesPercent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ResetMaxUpgradeDomainDeltaUnhealthyNodesPercent() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxUpgradeDomainDeltaUnhealthyNodesPercent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterUpgradePolicyHealthPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#max_unhealthy_applications_percent ServiceFabricCluster#max_unhealthy_applications_percent}.
	MaxUnhealthyApplicationsPercent *float64 `field:"optional" json:"maxUnhealthyApplicationsPercent" yaml:"maxUnhealthyApplicationsPercent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#max_unhealthy_nodes_percent ServiceFabricCluster#max_unhealthy_nodes_percent}.
	MaxUnhealthyNodesPercent *float64 `field:"optional" json:"maxUnhealthyNodesPercent" yaml:"maxUnhealthyNodesPercent"`
}

type ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference interface {
	cdktf.ComplexObject
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
	// Experimental.
	Fqn() *string
	InternalValue() *ServiceFabricClusterUpgradePolicyHealthPolicy
	SetInternalValue(val *ServiceFabricClusterUpgradePolicyHealthPolicy)
	MaxUnhealthyApplicationsPercent() *float64
	SetMaxUnhealthyApplicationsPercent(val *float64)
	MaxUnhealthyApplicationsPercentInput() *float64
	MaxUnhealthyNodesPercent() *float64
	SetMaxUnhealthyNodesPercent(val *float64)
	MaxUnhealthyNodesPercentInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetMaxUnhealthyApplicationsPercent()
	ResetMaxUnhealthyNodesPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference
type jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) InternalValue() *ServiceFabricClusterUpgradePolicyHealthPolicy {
	var returns *ServiceFabricClusterUpgradePolicyHealthPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) MaxUnhealthyApplicationsPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyApplicationsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) MaxUnhealthyApplicationsPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyApplicationsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) MaxUnhealthyNodesPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodesPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) MaxUnhealthyNodesPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodesPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterUpgradePolicyHealthPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterUpgradePolicyHealthPolicyOutputReference_Override(s ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) SetInternalValue(val *ServiceFabricClusterUpgradePolicyHealthPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) SetMaxUnhealthyApplicationsPercent(val *float64) {
	_jsii_.Set(
		j,
		"maxUnhealthyApplicationsPercent",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) SetMaxUnhealthyNodesPercent(val *float64) {
	_jsii_.Set(
		j,
		"maxUnhealthyNodesPercent",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) ResetMaxUnhealthyApplicationsPercent() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxUnhealthyApplicationsPercent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) ResetMaxUnhealthyNodesPercent() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxUnhealthyNodesPercent",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricClusterUpgradePolicyOutputReference interface {
	cdktf.ComplexObject
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
	DeltaHealthPolicy() ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference
	DeltaHealthPolicyInput() *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy
	ForceRestartEnabled() interface{}
	SetForceRestartEnabled(val interface{})
	ForceRestartEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	HealthCheckRetryTimeout() *string
	SetHealthCheckRetryTimeout(val *string)
	HealthCheckRetryTimeoutInput() *string
	HealthCheckStableDuration() *string
	SetHealthCheckStableDuration(val *string)
	HealthCheckStableDurationInput() *string
	HealthCheckWaitDuration() *string
	SetHealthCheckWaitDuration(val *string)
	HealthCheckWaitDurationInput() *string
	HealthPolicy() ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference
	HealthPolicyInput() *ServiceFabricClusterUpgradePolicyHealthPolicy
	InternalValue() *ServiceFabricClusterUpgradePolicy
	SetInternalValue(val *ServiceFabricClusterUpgradePolicy)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpgradeDomainTimeout() *string
	SetUpgradeDomainTimeout(val *string)
	UpgradeDomainTimeoutInput() *string
	UpgradeReplicaSetCheckTimeout() *string
	SetUpgradeReplicaSetCheckTimeout(val *string)
	UpgradeReplicaSetCheckTimeoutInput() *string
	UpgradeTimeout() *string
	SetUpgradeTimeout(val *string)
	UpgradeTimeoutInput() *string
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
	PutDeltaHealthPolicy(value *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy)
	PutHealthPolicy(value *ServiceFabricClusterUpgradePolicyHealthPolicy)
	ResetDeltaHealthPolicy()
	ResetForceRestartEnabled()
	ResetHealthCheckRetryTimeout()
	ResetHealthCheckStableDuration()
	ResetHealthCheckWaitDuration()
	ResetHealthPolicy()
	ResetUpgradeDomainTimeout()
	ResetUpgradeReplicaSetCheckTimeout()
	ResetUpgradeTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricClusterUpgradePolicyOutputReference
type jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) DeltaHealthPolicy() ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference {
	var returns ServiceFabricClusterUpgradePolicyDeltaHealthPolicyOutputReference
	_jsii_.Get(
		j,
		"deltaHealthPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) DeltaHealthPolicyInput() *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy {
	var returns *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy
	_jsii_.Get(
		j,
		"deltaHealthPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ForceRestartEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRestartEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ForceRestartEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceRestartEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthCheckRetryTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckRetryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthCheckRetryTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckRetryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthCheckStableDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckStableDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthCheckStableDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckStableDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthCheckWaitDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckWaitDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthCheckWaitDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckWaitDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthPolicy() ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference {
	var returns ServiceFabricClusterUpgradePolicyHealthPolicyOutputReference
	_jsii_.Get(
		j,
		"healthPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) HealthPolicyInput() *ServiceFabricClusterUpgradePolicyHealthPolicy {
	var returns *ServiceFabricClusterUpgradePolicyHealthPolicy
	_jsii_.Get(
		j,
		"healthPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) InternalValue() *ServiceFabricClusterUpgradePolicy {
	var returns *ServiceFabricClusterUpgradePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) UpgradeDomainTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeDomainTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) UpgradeDomainTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeDomainTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) UpgradeReplicaSetCheckTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeReplicaSetCheckTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) UpgradeReplicaSetCheckTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeReplicaSetCheckTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) UpgradeTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) UpgradeTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeTimeoutInput",
		&returns,
	)
	return returns
}


func NewServiceFabricClusterUpgradePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricClusterUpgradePolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterUpgradePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricClusterUpgradePolicyOutputReference_Override(s ServiceFabricClusterUpgradePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricCluster.ServiceFabricClusterUpgradePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) SetForceRestartEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"forceRestartEnabled",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) SetHealthCheckRetryTimeout(val *string) {
	_jsii_.Set(
		j,
		"healthCheckRetryTimeout",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) SetHealthCheckStableDuration(val *string) {
	_jsii_.Set(
		j,
		"healthCheckStableDuration",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) SetHealthCheckWaitDuration(val *string) {
	_jsii_.Set(
		j,
		"healthCheckWaitDuration",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) SetInternalValue(val *ServiceFabricClusterUpgradePolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) SetUpgradeDomainTimeout(val *string) {
	_jsii_.Set(
		j,
		"upgradeDomainTimeout",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) SetUpgradeReplicaSetCheckTimeout(val *string) {
	_jsii_.Set(
		j,
		"upgradeReplicaSetCheckTimeout",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) SetUpgradeTimeout(val *string) {
	_jsii_.Set(
		j,
		"upgradeTimeout",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) PutDeltaHealthPolicy(value *ServiceFabricClusterUpgradePolicyDeltaHealthPolicy) {
	_jsii_.InvokeVoid(
		s,
		"putDeltaHealthPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) PutHealthPolicy(value *ServiceFabricClusterUpgradePolicyHealthPolicy) {
	_jsii_.InvokeVoid(
		s,
		"putHealthPolicy",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetDeltaHealthPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetDeltaHealthPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetForceRestartEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetForceRestartEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetHealthCheckRetryTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetHealthCheckRetryTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetHealthCheckStableDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetHealthCheckStableDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetHealthCheckWaitDuration() {
	_jsii_.InvokeVoid(
		s,
		"resetHealthCheckWaitDuration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetHealthPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetHealthPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetUpgradeDomainTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetUpgradeDomainTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetUpgradeReplicaSetCheckTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetUpgradeReplicaSetCheckTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ResetUpgradeTimeout() {
	_jsii_.InvokeVoid(
		s,
		"resetUpgradeTimeout",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricClusterUpgradePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

