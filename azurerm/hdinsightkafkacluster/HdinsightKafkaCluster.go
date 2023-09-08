// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightkafkacluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/hdinsightkafkacluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/hdinsight_kafka_cluster azurerm_hdinsight_kafka_cluster}.
type HdinsightKafkaCluster interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterVersion() *string
	SetClusterVersion(val *string)
	ClusterVersionInput() *string
	ComponentVersion() HdinsightKafkaClusterComponentVersionOutputReference
	ComponentVersionInput() *HdinsightKafkaClusterComponentVersion
	ComputeIsolation() HdinsightKafkaClusterComputeIsolationOutputReference
	ComputeIsolationInput() *HdinsightKafkaClusterComputeIsolation
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
	DiskEncryption() HdinsightKafkaClusterDiskEncryptionList
	DiskEncryptionInput() interface{}
	EncryptionInTransitEnabled() interface{}
	SetEncryptionInTransitEnabled(val interface{})
	EncryptionInTransitEnabledInput() interface{}
	Extension() HdinsightKafkaClusterExtensionOutputReference
	ExtensionInput() *HdinsightKafkaClusterExtension
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Gateway() HdinsightKafkaClusterGatewayOutputReference
	GatewayInput() *HdinsightKafkaClusterGateway
	HttpsEndpoint() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KafkaRestProxyEndpoint() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Metastores() HdinsightKafkaClusterMetastoresOutputReference
	MetastoresInput() *HdinsightKafkaClusterMetastores
	Monitor() HdinsightKafkaClusterMonitorOutputReference
	MonitorInput() *HdinsightKafkaClusterMonitor
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() HdinsightKafkaClusterNetworkOutputReference
	NetworkInput() *HdinsightKafkaClusterNetwork
	// The tree node.
	Node() constructs.Node
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RestProxy() HdinsightKafkaClusterRestProxyOutputReference
	RestProxyInput() *HdinsightKafkaClusterRestProxy
	Roles() HdinsightKafkaClusterRolesOutputReference
	RolesInput() *HdinsightKafkaClusterRoles
	SecurityProfile() HdinsightKafkaClusterSecurityProfileOutputReference
	SecurityProfileInput() *HdinsightKafkaClusterSecurityProfile
	SshEndpoint() *string
	StorageAccount() HdinsightKafkaClusterStorageAccountList
	StorageAccountGen2() HdinsightKafkaClusterStorageAccountGen2OutputReference
	StorageAccountGen2Input() *HdinsightKafkaClusterStorageAccountGen2
	StorageAccountInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	Timeouts() HdinsightKafkaClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	TlsMinVersion() *string
	SetTlsMinVersion(val *string)
	TlsMinVersionInput() *string
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
	PutComponentVersion(value *HdinsightKafkaClusterComponentVersion)
	PutComputeIsolation(value *HdinsightKafkaClusterComputeIsolation)
	PutDiskEncryption(value interface{})
	PutExtension(value *HdinsightKafkaClusterExtension)
	PutGateway(value *HdinsightKafkaClusterGateway)
	PutMetastores(value *HdinsightKafkaClusterMetastores)
	PutMonitor(value *HdinsightKafkaClusterMonitor)
	PutNetwork(value *HdinsightKafkaClusterNetwork)
	PutRestProxy(value *HdinsightKafkaClusterRestProxy)
	PutRoles(value *HdinsightKafkaClusterRoles)
	PutSecurityProfile(value *HdinsightKafkaClusterSecurityProfile)
	PutStorageAccount(value interface{})
	PutStorageAccountGen2(value *HdinsightKafkaClusterStorageAccountGen2)
	PutTimeouts(value *HdinsightKafkaClusterTimeouts)
	ResetComputeIsolation()
	ResetDiskEncryption()
	ResetEncryptionInTransitEnabled()
	ResetExtension()
	ResetId()
	ResetMetastores()
	ResetMonitor()
	ResetNetwork()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRestProxy()
	ResetSecurityProfile()
	ResetStorageAccount()
	ResetStorageAccountGen2()
	ResetTags()
	ResetTimeouts()
	ResetTlsMinVersion()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for HdinsightKafkaCluster
type jsiiProxy_HdinsightKafkaCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_HdinsightKafkaCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) ClusterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) ClusterVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) ComponentVersion() HdinsightKafkaClusterComponentVersionOutputReference {
	var returns HdinsightKafkaClusterComponentVersionOutputReference
	_jsii_.Get(
		j,
		"componentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) ComponentVersionInput() *HdinsightKafkaClusterComponentVersion {
	var returns *HdinsightKafkaClusterComponentVersion
	_jsii_.Get(
		j,
		"componentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) ComputeIsolation() HdinsightKafkaClusterComputeIsolationOutputReference {
	var returns HdinsightKafkaClusterComputeIsolationOutputReference
	_jsii_.Get(
		j,
		"computeIsolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) ComputeIsolationInput() *HdinsightKafkaClusterComputeIsolation {
	var returns *HdinsightKafkaClusterComputeIsolation
	_jsii_.Get(
		j,
		"computeIsolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) DiskEncryption() HdinsightKafkaClusterDiskEncryptionList {
	var returns HdinsightKafkaClusterDiskEncryptionList
	_jsii_.Get(
		j,
		"diskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) DiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) EncryptionInTransitEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInTransitEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) EncryptionInTransitEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInTransitEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Extension() HdinsightKafkaClusterExtensionOutputReference {
	var returns HdinsightKafkaClusterExtensionOutputReference
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) ExtensionInput() *HdinsightKafkaClusterExtension {
	var returns *HdinsightKafkaClusterExtension
	_jsii_.Get(
		j,
		"extensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Gateway() HdinsightKafkaClusterGatewayOutputReference {
	var returns HdinsightKafkaClusterGatewayOutputReference
	_jsii_.Get(
		j,
		"gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) GatewayInput() *HdinsightKafkaClusterGateway {
	var returns *HdinsightKafkaClusterGateway
	_jsii_.Get(
		j,
		"gatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) HttpsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) KafkaRestProxyEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaRestProxyEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Metastores() HdinsightKafkaClusterMetastoresOutputReference {
	var returns HdinsightKafkaClusterMetastoresOutputReference
	_jsii_.Get(
		j,
		"metastores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) MetastoresInput() *HdinsightKafkaClusterMetastores {
	var returns *HdinsightKafkaClusterMetastores
	_jsii_.Get(
		j,
		"metastoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Monitor() HdinsightKafkaClusterMonitorOutputReference {
	var returns HdinsightKafkaClusterMonitorOutputReference
	_jsii_.Get(
		j,
		"monitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) MonitorInput() *HdinsightKafkaClusterMonitor {
	var returns *HdinsightKafkaClusterMonitor
	_jsii_.Get(
		j,
		"monitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Network() HdinsightKafkaClusterNetworkOutputReference {
	var returns HdinsightKafkaClusterNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) NetworkInput() *HdinsightKafkaClusterNetwork {
	var returns *HdinsightKafkaClusterNetwork
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) RestProxy() HdinsightKafkaClusterRestProxyOutputReference {
	var returns HdinsightKafkaClusterRestProxyOutputReference
	_jsii_.Get(
		j,
		"restProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) RestProxyInput() *HdinsightKafkaClusterRestProxy {
	var returns *HdinsightKafkaClusterRestProxy
	_jsii_.Get(
		j,
		"restProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Roles() HdinsightKafkaClusterRolesOutputReference {
	var returns HdinsightKafkaClusterRolesOutputReference
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) RolesInput() *HdinsightKafkaClusterRoles {
	var returns *HdinsightKafkaClusterRoles
	_jsii_.Get(
		j,
		"rolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) SecurityProfile() HdinsightKafkaClusterSecurityProfileOutputReference {
	var returns HdinsightKafkaClusterSecurityProfileOutputReference
	_jsii_.Get(
		j,
		"securityProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) SecurityProfileInput() *HdinsightKafkaClusterSecurityProfile {
	var returns *HdinsightKafkaClusterSecurityProfile
	_jsii_.Get(
		j,
		"securityProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) SshEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) StorageAccount() HdinsightKafkaClusterStorageAccountList {
	var returns HdinsightKafkaClusterStorageAccountList
	_jsii_.Get(
		j,
		"storageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) StorageAccountGen2() HdinsightKafkaClusterStorageAccountGen2OutputReference {
	var returns HdinsightKafkaClusterStorageAccountGen2OutputReference
	_jsii_.Get(
		j,
		"storageAccountGen2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) StorageAccountGen2Input() *HdinsightKafkaClusterStorageAccountGen2 {
	var returns *HdinsightKafkaClusterStorageAccountGen2
	_jsii_.Get(
		j,
		"storageAccountGen2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) StorageAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) Timeouts() HdinsightKafkaClusterTimeoutsOutputReference {
	var returns HdinsightKafkaClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) TlsMinVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaCluster) TlsMinVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/hdinsight_kafka_cluster azurerm_hdinsight_kafka_cluster} Resource.
func NewHdinsightKafkaCluster(scope constructs.Construct, id *string, config *HdinsightKafkaClusterConfig) HdinsightKafkaCluster {
	_init_.Initialize()

	if err := validateNewHdinsightKafkaClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HdinsightKafkaCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightKafkaCluster.HdinsightKafkaCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/hdinsight_kafka_cluster azurerm_hdinsight_kafka_cluster} Resource.
func NewHdinsightKafkaCluster_Override(h HdinsightKafkaCluster, scope constructs.Construct, id *string, config *HdinsightKafkaClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightKafkaCluster.HdinsightKafkaCluster",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetClusterVersion(val *string) {
	if err := j.validateSetClusterVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterVersion",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetEncryptionInTransitEnabled(val interface{}) {
	if err := j.validateSetEncryptionInTransitEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionInTransitEnabled",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaCluster)SetTlsMinVersion(val *string) {
	if err := j.validateSetTlsMinVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsMinVersion",
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
func HdinsightKafkaCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHdinsightKafkaCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.hdinsightKafkaCluster.HdinsightKafkaCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HdinsightKafkaCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHdinsightKafkaCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.hdinsightKafkaCluster.HdinsightKafkaCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HdinsightKafkaCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHdinsightKafkaCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.hdinsightKafkaCluster.HdinsightKafkaCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HdinsightKafkaCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.hdinsightKafkaCluster.HdinsightKafkaCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutComponentVersion(value *HdinsightKafkaClusterComponentVersion) {
	if err := h.validatePutComponentVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putComponentVersion",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutComputeIsolation(value *HdinsightKafkaClusterComputeIsolation) {
	if err := h.validatePutComputeIsolationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putComputeIsolation",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutDiskEncryption(value interface{}) {
	if err := h.validatePutDiskEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putDiskEncryption",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutExtension(value *HdinsightKafkaClusterExtension) {
	if err := h.validatePutExtensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putExtension",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutGateway(value *HdinsightKafkaClusterGateway) {
	if err := h.validatePutGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putGateway",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutMetastores(value *HdinsightKafkaClusterMetastores) {
	if err := h.validatePutMetastoresParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putMetastores",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutMonitor(value *HdinsightKafkaClusterMonitor) {
	if err := h.validatePutMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putMonitor",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutNetwork(value *HdinsightKafkaClusterNetwork) {
	if err := h.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNetwork",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutRestProxy(value *HdinsightKafkaClusterRestProxy) {
	if err := h.validatePutRestProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putRestProxy",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutRoles(value *HdinsightKafkaClusterRoles) {
	if err := h.validatePutRolesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putRoles",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutSecurityProfile(value *HdinsightKafkaClusterSecurityProfile) {
	if err := h.validatePutSecurityProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putSecurityProfile",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutStorageAccount(value interface{}) {
	if err := h.validatePutStorageAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putStorageAccount",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutStorageAccountGen2(value *HdinsightKafkaClusterStorageAccountGen2) {
	if err := h.validatePutStorageAccountGen2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putStorageAccountGen2",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) PutTimeouts(value *HdinsightKafkaClusterTimeouts) {
	if err := h.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetComputeIsolation() {
	_jsii_.InvokeVoid(
		h,
		"resetComputeIsolation",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetDiskEncryption() {
	_jsii_.InvokeVoid(
		h,
		"resetDiskEncryption",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetEncryptionInTransitEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetEncryptionInTransitEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetExtension() {
	_jsii_.InvokeVoid(
		h,
		"resetExtension",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetMetastores() {
	_jsii_.InvokeVoid(
		h,
		"resetMetastores",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetMonitor() {
	_jsii_.InvokeVoid(
		h,
		"resetMonitor",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetNetwork() {
	_jsii_.InvokeVoid(
		h,
		"resetNetwork",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetRestProxy() {
	_jsii_.InvokeVoid(
		h,
		"resetRestProxy",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetSecurityProfile() {
	_jsii_.InvokeVoid(
		h,
		"resetSecurityProfile",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetStorageAccount() {
	_jsii_.InvokeVoid(
		h,
		"resetStorageAccount",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetStorageAccountGen2() {
	_jsii_.InvokeVoid(
		h,
		"resetStorageAccountGen2",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetTags() {
	_jsii_.InvokeVoid(
		h,
		"resetTags",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		h,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) ResetTlsMinVersion() {
	_jsii_.InvokeVoid(
		h,
		"resetTlsMinVersion",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

