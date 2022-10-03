package hdinsighthadoopcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/hdinsighthadoopcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster azurerm_hdinsight_hadoop_cluster}.
type HdinsightHadoopCluster interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterVersion() *string
	SetClusterVersion(val *string)
	ClusterVersionInput() *string
	ComponentVersion() HdinsightHadoopClusterComponentVersionOutputReference
	ComponentVersionInput() *HdinsightHadoopClusterComponentVersion
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
	DiskEncryption() HdinsightHadoopClusterDiskEncryptionList
	DiskEncryptionInput() interface{}
	Extension() HdinsightHadoopClusterExtensionOutputReference
	ExtensionInput() *HdinsightHadoopClusterExtension
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Gateway() HdinsightHadoopClusterGatewayOutputReference
	GatewayInput() *HdinsightHadoopClusterGateway
	HttpsEndpoint() *string
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
	Metastores() HdinsightHadoopClusterMetastoresOutputReference
	MetastoresInput() *HdinsightHadoopClusterMetastores
	Monitor() HdinsightHadoopClusterMonitorOutputReference
	MonitorInput() *HdinsightHadoopClusterMonitor
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() HdinsightHadoopClusterNetworkOutputReference
	NetworkInput() *HdinsightHadoopClusterNetwork
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
	Roles() HdinsightHadoopClusterRolesOutputReference
	RolesInput() *HdinsightHadoopClusterRoles
	SecurityProfile() HdinsightHadoopClusterSecurityProfileOutputReference
	SecurityProfileInput() *HdinsightHadoopClusterSecurityProfile
	SshEndpoint() *string
	StorageAccount() HdinsightHadoopClusterStorageAccountList
	StorageAccountGen2() HdinsightHadoopClusterStorageAccountGen2OutputReference
	StorageAccountGen2Input() *HdinsightHadoopClusterStorageAccountGen2
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
	Timeouts() HdinsightHadoopClusterTimeoutsOutputReference
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
	PutComponentVersion(value *HdinsightHadoopClusterComponentVersion)
	PutDiskEncryption(value interface{})
	PutExtension(value *HdinsightHadoopClusterExtension)
	PutGateway(value *HdinsightHadoopClusterGateway)
	PutMetastores(value *HdinsightHadoopClusterMetastores)
	PutMonitor(value *HdinsightHadoopClusterMonitor)
	PutNetwork(value *HdinsightHadoopClusterNetwork)
	PutRoles(value *HdinsightHadoopClusterRoles)
	PutSecurityProfile(value *HdinsightHadoopClusterSecurityProfile)
	PutStorageAccount(value interface{})
	PutStorageAccountGen2(value *HdinsightHadoopClusterStorageAccountGen2)
	PutTimeouts(value *HdinsightHadoopClusterTimeouts)
	ResetDiskEncryption()
	ResetExtension()
	ResetId()
	ResetMetastores()
	ResetMonitor()
	ResetNetwork()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for HdinsightHadoopCluster
type jsiiProxy_HdinsightHadoopCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_HdinsightHadoopCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) ClusterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) ClusterVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) ComponentVersion() HdinsightHadoopClusterComponentVersionOutputReference {
	var returns HdinsightHadoopClusterComponentVersionOutputReference
	_jsii_.Get(
		j,
		"componentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) ComponentVersionInput() *HdinsightHadoopClusterComponentVersion {
	var returns *HdinsightHadoopClusterComponentVersion
	_jsii_.Get(
		j,
		"componentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) DiskEncryption() HdinsightHadoopClusterDiskEncryptionList {
	var returns HdinsightHadoopClusterDiskEncryptionList
	_jsii_.Get(
		j,
		"diskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) DiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Extension() HdinsightHadoopClusterExtensionOutputReference {
	var returns HdinsightHadoopClusterExtensionOutputReference
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) ExtensionInput() *HdinsightHadoopClusterExtension {
	var returns *HdinsightHadoopClusterExtension
	_jsii_.Get(
		j,
		"extensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Gateway() HdinsightHadoopClusterGatewayOutputReference {
	var returns HdinsightHadoopClusterGatewayOutputReference
	_jsii_.Get(
		j,
		"gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) GatewayInput() *HdinsightHadoopClusterGateway {
	var returns *HdinsightHadoopClusterGateway
	_jsii_.Get(
		j,
		"gatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) HttpsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Metastores() HdinsightHadoopClusterMetastoresOutputReference {
	var returns HdinsightHadoopClusterMetastoresOutputReference
	_jsii_.Get(
		j,
		"metastores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) MetastoresInput() *HdinsightHadoopClusterMetastores {
	var returns *HdinsightHadoopClusterMetastores
	_jsii_.Get(
		j,
		"metastoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Monitor() HdinsightHadoopClusterMonitorOutputReference {
	var returns HdinsightHadoopClusterMonitorOutputReference
	_jsii_.Get(
		j,
		"monitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) MonitorInput() *HdinsightHadoopClusterMonitor {
	var returns *HdinsightHadoopClusterMonitor
	_jsii_.Get(
		j,
		"monitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Network() HdinsightHadoopClusterNetworkOutputReference {
	var returns HdinsightHadoopClusterNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) NetworkInput() *HdinsightHadoopClusterNetwork {
	var returns *HdinsightHadoopClusterNetwork
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Roles() HdinsightHadoopClusterRolesOutputReference {
	var returns HdinsightHadoopClusterRolesOutputReference
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) RolesInput() *HdinsightHadoopClusterRoles {
	var returns *HdinsightHadoopClusterRoles
	_jsii_.Get(
		j,
		"rolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) SecurityProfile() HdinsightHadoopClusterSecurityProfileOutputReference {
	var returns HdinsightHadoopClusterSecurityProfileOutputReference
	_jsii_.Get(
		j,
		"securityProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) SecurityProfileInput() *HdinsightHadoopClusterSecurityProfile {
	var returns *HdinsightHadoopClusterSecurityProfile
	_jsii_.Get(
		j,
		"securityProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) SshEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) StorageAccount() HdinsightHadoopClusterStorageAccountList {
	var returns HdinsightHadoopClusterStorageAccountList
	_jsii_.Get(
		j,
		"storageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) StorageAccountGen2() HdinsightHadoopClusterStorageAccountGen2OutputReference {
	var returns HdinsightHadoopClusterStorageAccountGen2OutputReference
	_jsii_.Get(
		j,
		"storageAccountGen2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) StorageAccountGen2Input() *HdinsightHadoopClusterStorageAccountGen2 {
	var returns *HdinsightHadoopClusterStorageAccountGen2
	_jsii_.Get(
		j,
		"storageAccountGen2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) StorageAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) Timeouts() HdinsightHadoopClusterTimeoutsOutputReference {
	var returns HdinsightHadoopClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) TlsMinVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopCluster) TlsMinVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster azurerm_hdinsight_hadoop_cluster} Resource.
func NewHdinsightHadoopCluster(scope constructs.Construct, id *string, config *HdinsightHadoopClusterConfig) HdinsightHadoopCluster {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster azurerm_hdinsight_hadoop_cluster} Resource.
func NewHdinsightHadoopCluster_Override(h HdinsightHadoopCluster, scope constructs.Construct, id *string, config *HdinsightHadoopClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopCluster",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetClusterVersion(val *string) {
	_jsii_.Set(
		j,
		"clusterVersion",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetTier(val *string) {
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopCluster) SetTlsMinVersion(val *string) {
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
func HdinsightHadoopCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HdinsightHadoopCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) PutComponentVersion(value *HdinsightHadoopClusterComponentVersion) {
	_jsii_.InvokeVoid(
		h,
		"putComponentVersion",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) PutDiskEncryption(value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"putDiskEncryption",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) PutExtension(value *HdinsightHadoopClusterExtension) {
	_jsii_.InvokeVoid(
		h,
		"putExtension",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) PutGateway(value *HdinsightHadoopClusterGateway) {
	_jsii_.InvokeVoid(
		h,
		"putGateway",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) PutMetastores(value *HdinsightHadoopClusterMetastores) {
	_jsii_.InvokeVoid(
		h,
		"putMetastores",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) PutMonitor(value *HdinsightHadoopClusterMonitor) {
	_jsii_.InvokeVoid(
		h,
		"putMonitor",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) PutNetwork(value *HdinsightHadoopClusterNetwork) {
	_jsii_.InvokeVoid(
		h,
		"putNetwork",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) PutRoles(value *HdinsightHadoopClusterRoles) {
	_jsii_.InvokeVoid(
		h,
		"putRoles",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) PutSecurityProfile(value *HdinsightHadoopClusterSecurityProfile) {
	_jsii_.InvokeVoid(
		h,
		"putSecurityProfile",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) PutStorageAccount(value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"putStorageAccount",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) PutStorageAccountGen2(value *HdinsightHadoopClusterStorageAccountGen2) {
	_jsii_.InvokeVoid(
		h,
		"putStorageAccountGen2",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) PutTimeouts(value *HdinsightHadoopClusterTimeouts) {
	_jsii_.InvokeVoid(
		h,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) ResetDiskEncryption() {
	_jsii_.InvokeVoid(
		h,
		"resetDiskEncryption",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) ResetExtension() {
	_jsii_.InvokeVoid(
		h,
		"resetExtension",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) ResetMetastores() {
	_jsii_.InvokeVoid(
		h,
		"resetMetastores",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) ResetMonitor() {
	_jsii_.InvokeVoid(
		h,
		"resetMonitor",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) ResetNetwork() {
	_jsii_.InvokeVoid(
		h,
		"resetNetwork",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) ResetSecurityProfile() {
	_jsii_.InvokeVoid(
		h,
		"resetSecurityProfile",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) ResetStorageAccount() {
	_jsii_.InvokeVoid(
		h,
		"resetStorageAccount",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) ResetStorageAccountGen2() {
	_jsii_.InvokeVoid(
		h,
		"resetStorageAccountGen2",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) ResetTags() {
	_jsii_.InvokeVoid(
		h,
		"resetTags",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		h,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) ResetTlsMinVersion() {
	_jsii_.InvokeVoid(
		h,
		"resetTlsMinVersion",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterComponentVersion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#hadoop HdinsightHadoopCluster#hadoop}.
	Hadoop *string `field:"required" json:"hadoop" yaml:"hadoop"`
}

type HdinsightHadoopClusterComponentVersionOutputReference interface {
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
	Hadoop() *string
	SetHadoop(val *string)
	HadoopInput() *string
	InternalValue() *HdinsightHadoopClusterComponentVersion
	SetInternalValue(val *HdinsightHadoopClusterComponentVersion)
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

// The jsii proxy struct for HdinsightHadoopClusterComponentVersionOutputReference
type jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) Hadoop() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hadoop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) HadoopInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hadoopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) InternalValue() *HdinsightHadoopClusterComponentVersion {
	var returns *HdinsightHadoopClusterComponentVersion
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterComponentVersionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterComponentVersionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterComponentVersionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterComponentVersionOutputReference_Override(h HdinsightHadoopClusterComponentVersionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterComponentVersionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) SetHadoop(val *string) {
	_jsii_.Set(
		j,
		"hadoop",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) SetInternalValue(val *HdinsightHadoopClusterComponentVersion) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterComponentVersionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#cluster_version HdinsightHadoopCluster#cluster_version}.
	ClusterVersion *string `field:"required" json:"clusterVersion" yaml:"clusterVersion"`
	// component_version block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#component_version HdinsightHadoopCluster#component_version}
	ComponentVersion *HdinsightHadoopClusterComponentVersion `field:"required" json:"componentVersion" yaml:"componentVersion"`
	// gateway block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#gateway HdinsightHadoopCluster#gateway}
	Gateway *HdinsightHadoopClusterGateway `field:"required" json:"gateway" yaml:"gateway"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#location HdinsightHadoopCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#name HdinsightHadoopCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#resource_group_name HdinsightHadoopCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// roles block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#roles HdinsightHadoopCluster#roles}
	Roles *HdinsightHadoopClusterRoles `field:"required" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#tier HdinsightHadoopCluster#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// disk_encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#disk_encryption HdinsightHadoopCluster#disk_encryption}
	DiskEncryption interface{} `field:"optional" json:"diskEncryption" yaml:"diskEncryption"`
	// extension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#extension HdinsightHadoopCluster#extension}
	Extension *HdinsightHadoopClusterExtension `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#id HdinsightHadoopCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// metastores block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#metastores HdinsightHadoopCluster#metastores}
	Metastores *HdinsightHadoopClusterMetastores `field:"optional" json:"metastores" yaml:"metastores"`
	// monitor block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#monitor HdinsightHadoopCluster#monitor}
	Monitor *HdinsightHadoopClusterMonitor `field:"optional" json:"monitor" yaml:"monitor"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#network HdinsightHadoopCluster#network}
	Network *HdinsightHadoopClusterNetwork `field:"optional" json:"network" yaml:"network"`
	// security_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#security_profile HdinsightHadoopCluster#security_profile}
	SecurityProfile *HdinsightHadoopClusterSecurityProfile `field:"optional" json:"securityProfile" yaml:"securityProfile"`
	// storage_account block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#storage_account HdinsightHadoopCluster#storage_account}
	StorageAccount interface{} `field:"optional" json:"storageAccount" yaml:"storageAccount"`
	// storage_account_gen2 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#storage_account_gen2 HdinsightHadoopCluster#storage_account_gen2}
	StorageAccountGen2 *HdinsightHadoopClusterStorageAccountGen2 `field:"optional" json:"storageAccountGen2" yaml:"storageAccountGen2"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#tags HdinsightHadoopCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#timeouts HdinsightHadoopCluster#timeouts}
	Timeouts *HdinsightHadoopClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#tls_min_version HdinsightHadoopCluster#tls_min_version}.
	TlsMinVersion *string `field:"optional" json:"tlsMinVersion" yaml:"tlsMinVersion"`
}

type HdinsightHadoopClusterDiskEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#encryption_algorithm HdinsightHadoopCluster#encryption_algorithm}.
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#encryption_at_host_enabled HdinsightHadoopCluster#encryption_at_host_enabled}.
	EncryptionAtHostEnabled interface{} `field:"optional" json:"encryptionAtHostEnabled" yaml:"encryptionAtHostEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#key_vault_key_id HdinsightHadoopCluster#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#key_vault_managed_identity_id HdinsightHadoopCluster#key_vault_managed_identity_id}.
	KeyVaultManagedIdentityId *string `field:"optional" json:"keyVaultManagedIdentityId" yaml:"keyVaultManagedIdentityId"`
}

type HdinsightHadoopClusterDiskEncryptionList interface {
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
	Get(index *float64) HdinsightHadoopClusterDiskEncryptionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterDiskEncryptionList
type jsiiProxy_HdinsightHadoopClusterDiskEncryptionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterDiskEncryptionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HdinsightHadoopClusterDiskEncryptionList {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterDiskEncryptionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterDiskEncryptionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterDiskEncryptionList_Override(h HdinsightHadoopClusterDiskEncryptionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterDiskEncryptionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) Get(index *float64) HdinsightHadoopClusterDiskEncryptionOutputReference {
	var returns HdinsightHadoopClusterDiskEncryptionOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterDiskEncryptionOutputReference interface {
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
	EncryptionAlgorithm() *string
	SetEncryptionAlgorithm(val *string)
	EncryptionAlgorithmInput() *string
	EncryptionAtHostEnabled() interface{}
	SetEncryptionAtHostEnabled(val interface{})
	EncryptionAtHostEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyVaultKeyId() *string
	SetKeyVaultKeyId(val *string)
	KeyVaultKeyIdInput() *string
	KeyVaultManagedIdentityId() *string
	SetKeyVaultManagedIdentityId(val *string)
	KeyVaultManagedIdentityIdInput() *string
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
	ResetEncryptionAlgorithm()
	ResetEncryptionAtHostEnabled()
	ResetKeyVaultKeyId()
	ResetKeyVaultManagedIdentityId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterDiskEncryptionOutputReference
type jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) EncryptionAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) EncryptionAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) EncryptionAtHostEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) EncryptionAtHostEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) KeyVaultKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) KeyVaultKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) KeyVaultManagedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultManagedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) KeyVaultManagedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultManagedIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterDiskEncryptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HdinsightHadoopClusterDiskEncryptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterDiskEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterDiskEncryptionOutputReference_Override(h HdinsightHadoopClusterDiskEncryptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterDiskEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) SetEncryptionAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"encryptionAlgorithm",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) SetEncryptionAtHostEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionAtHostEnabled",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) SetKeyVaultKeyId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultKeyId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) SetKeyVaultManagedIdentityId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultManagedIdentityId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) ResetEncryptionAlgorithm() {
	_jsii_.InvokeVoid(
		h,
		"resetEncryptionAlgorithm",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) ResetEncryptionAtHostEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetEncryptionAtHostEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) ResetKeyVaultKeyId() {
	_jsii_.InvokeVoid(
		h,
		"resetKeyVaultKeyId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) ResetKeyVaultManagedIdentityId() {
	_jsii_.InvokeVoid(
		h,
		"resetKeyVaultManagedIdentityId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterDiskEncryptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterExtension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#log_analytics_workspace_id HdinsightHadoopCluster#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#primary_key HdinsightHadoopCluster#primary_key}.
	PrimaryKey *string `field:"required" json:"primaryKey" yaml:"primaryKey"`
}

type HdinsightHadoopClusterExtensionOutputReference interface {
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
	InternalValue() *HdinsightHadoopClusterExtension
	SetInternalValue(val *HdinsightHadoopClusterExtension)
	LogAnalyticsWorkspaceId() *string
	SetLogAnalyticsWorkspaceId(val *string)
	LogAnalyticsWorkspaceIdInput() *string
	PrimaryKey() *string
	SetPrimaryKey(val *string)
	PrimaryKeyInput() *string
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

// The jsii proxy struct for HdinsightHadoopClusterExtensionOutputReference
type jsiiProxy_HdinsightHadoopClusterExtensionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) InternalValue() *HdinsightHadoopClusterExtension {
	var returns *HdinsightHadoopClusterExtension
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) PrimaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) PrimaryKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterExtensionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterExtensionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterExtensionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterExtensionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterExtensionOutputReference_Override(h HdinsightHadoopClusterExtensionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterExtensionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) SetInternalValue(val *HdinsightHadoopClusterExtension) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) SetLogAnalyticsWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) SetPrimaryKey(val *string) {
	_jsii_.Set(
		j,
		"primaryKey",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterExtensionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterGateway struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#password HdinsightHadoopCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#username HdinsightHadoopCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

type HdinsightHadoopClusterGatewayOutputReference interface {
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
	InternalValue() *HdinsightHadoopClusterGateway
	SetInternalValue(val *HdinsightHadoopClusterGateway)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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

// The jsii proxy struct for HdinsightHadoopClusterGatewayOutputReference
type jsiiProxy_HdinsightHadoopClusterGatewayOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) InternalValue() *HdinsightHadoopClusterGateway {
	var returns *HdinsightHadoopClusterGateway
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterGatewayOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterGatewayOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterGatewayOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterGatewayOutputReference_Override(h HdinsightHadoopClusterGatewayOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) SetInternalValue(val *HdinsightHadoopClusterGateway) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterGatewayOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterMetastores struct {
	// ambari block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#ambari HdinsightHadoopCluster#ambari}
	Ambari *HdinsightHadoopClusterMetastoresAmbari `field:"optional" json:"ambari" yaml:"ambari"`
	// hive block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#hive HdinsightHadoopCluster#hive}
	Hive *HdinsightHadoopClusterMetastoresHive `field:"optional" json:"hive" yaml:"hive"`
	// oozie block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#oozie HdinsightHadoopCluster#oozie}
	Oozie *HdinsightHadoopClusterMetastoresOozie `field:"optional" json:"oozie" yaml:"oozie"`
}

type HdinsightHadoopClusterMetastoresAmbari struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#database_name HdinsightHadoopCluster#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#password HdinsightHadoopCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#server HdinsightHadoopCluster#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#username HdinsightHadoopCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

type HdinsightHadoopClusterMetastoresAmbariOutputReference interface {
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *HdinsightHadoopClusterMetastoresAmbari
	SetInternalValue(val *HdinsightHadoopClusterMetastoresAmbari)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Server() *string
	SetServer(val *string)
	ServerInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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

// The jsii proxy struct for HdinsightHadoopClusterMetastoresAmbariOutputReference
type jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) InternalValue() *HdinsightHadoopClusterMetastoresAmbari {
	var returns *HdinsightHadoopClusterMetastoresAmbari
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) Server() *string {
	var returns *string
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) ServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterMetastoresAmbariOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterMetastoresAmbariOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterMetastoresAmbariOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterMetastoresAmbariOutputReference_Override(h HdinsightHadoopClusterMetastoresAmbariOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterMetastoresAmbariOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) SetInternalValue(val *HdinsightHadoopClusterMetastoresAmbari) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) SetServer(val *string) {
	_jsii_.Set(
		j,
		"server",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresAmbariOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterMetastoresHive struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#database_name HdinsightHadoopCluster#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#password HdinsightHadoopCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#server HdinsightHadoopCluster#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#username HdinsightHadoopCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

type HdinsightHadoopClusterMetastoresHiveOutputReference interface {
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *HdinsightHadoopClusterMetastoresHive
	SetInternalValue(val *HdinsightHadoopClusterMetastoresHive)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Server() *string
	SetServer(val *string)
	ServerInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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

// The jsii proxy struct for HdinsightHadoopClusterMetastoresHiveOutputReference
type jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) InternalValue() *HdinsightHadoopClusterMetastoresHive {
	var returns *HdinsightHadoopClusterMetastoresHive
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) Server() *string {
	var returns *string
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) ServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterMetastoresHiveOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterMetastoresHiveOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterMetastoresHiveOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterMetastoresHiveOutputReference_Override(h HdinsightHadoopClusterMetastoresHiveOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterMetastoresHiveOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) SetInternalValue(val *HdinsightHadoopClusterMetastoresHive) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) SetServer(val *string) {
	_jsii_.Set(
		j,
		"server",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresHiveOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterMetastoresOozie struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#database_name HdinsightHadoopCluster#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#password HdinsightHadoopCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#server HdinsightHadoopCluster#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#username HdinsightHadoopCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

type HdinsightHadoopClusterMetastoresOozieOutputReference interface {
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *HdinsightHadoopClusterMetastoresOozie
	SetInternalValue(val *HdinsightHadoopClusterMetastoresOozie)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Server() *string
	SetServer(val *string)
	ServerInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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

// The jsii proxy struct for HdinsightHadoopClusterMetastoresOozieOutputReference
type jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) InternalValue() *HdinsightHadoopClusterMetastoresOozie {
	var returns *HdinsightHadoopClusterMetastoresOozie
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) Server() *string {
	var returns *string
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) ServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterMetastoresOozieOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterMetastoresOozieOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterMetastoresOozieOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterMetastoresOozieOutputReference_Override(h HdinsightHadoopClusterMetastoresOozieOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterMetastoresOozieOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) SetInternalValue(val *HdinsightHadoopClusterMetastoresOozie) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) SetServer(val *string) {
	_jsii_.Set(
		j,
		"server",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOozieOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterMetastoresOutputReference interface {
	cdktf.ComplexObject
	Ambari() HdinsightHadoopClusterMetastoresAmbariOutputReference
	AmbariInput() *HdinsightHadoopClusterMetastoresAmbari
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
	Hive() HdinsightHadoopClusterMetastoresHiveOutputReference
	HiveInput() *HdinsightHadoopClusterMetastoresHive
	InternalValue() *HdinsightHadoopClusterMetastores
	SetInternalValue(val *HdinsightHadoopClusterMetastores)
	Oozie() HdinsightHadoopClusterMetastoresOozieOutputReference
	OozieInput() *HdinsightHadoopClusterMetastoresOozie
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
	PutAmbari(value *HdinsightHadoopClusterMetastoresAmbari)
	PutHive(value *HdinsightHadoopClusterMetastoresHive)
	PutOozie(value *HdinsightHadoopClusterMetastoresOozie)
	ResetAmbari()
	ResetHive()
	ResetOozie()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterMetastoresOutputReference
type jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) Ambari() HdinsightHadoopClusterMetastoresAmbariOutputReference {
	var returns HdinsightHadoopClusterMetastoresAmbariOutputReference
	_jsii_.Get(
		j,
		"ambari",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) AmbariInput() *HdinsightHadoopClusterMetastoresAmbari {
	var returns *HdinsightHadoopClusterMetastoresAmbari
	_jsii_.Get(
		j,
		"ambariInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) Hive() HdinsightHadoopClusterMetastoresHiveOutputReference {
	var returns HdinsightHadoopClusterMetastoresHiveOutputReference
	_jsii_.Get(
		j,
		"hive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) HiveInput() *HdinsightHadoopClusterMetastoresHive {
	var returns *HdinsightHadoopClusterMetastoresHive
	_jsii_.Get(
		j,
		"hiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) InternalValue() *HdinsightHadoopClusterMetastores {
	var returns *HdinsightHadoopClusterMetastores
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) Oozie() HdinsightHadoopClusterMetastoresOozieOutputReference {
	var returns HdinsightHadoopClusterMetastoresOozieOutputReference
	_jsii_.Get(
		j,
		"oozie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) OozieInput() *HdinsightHadoopClusterMetastoresOozie {
	var returns *HdinsightHadoopClusterMetastoresOozie
	_jsii_.Get(
		j,
		"oozieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterMetastoresOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterMetastoresOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterMetastoresOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterMetastoresOutputReference_Override(h HdinsightHadoopClusterMetastoresOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterMetastoresOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) SetInternalValue(val *HdinsightHadoopClusterMetastores) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) PutAmbari(value *HdinsightHadoopClusterMetastoresAmbari) {
	_jsii_.InvokeVoid(
		h,
		"putAmbari",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) PutHive(value *HdinsightHadoopClusterMetastoresHive) {
	_jsii_.InvokeVoid(
		h,
		"putHive",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) PutOozie(value *HdinsightHadoopClusterMetastoresOozie) {
	_jsii_.InvokeVoid(
		h,
		"putOozie",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) ResetAmbari() {
	_jsii_.InvokeVoid(
		h,
		"resetAmbari",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) ResetHive() {
	_jsii_.InvokeVoid(
		h,
		"resetHive",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) ResetOozie() {
	_jsii_.InvokeVoid(
		h,
		"resetOozie",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMetastoresOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterMonitor struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#log_analytics_workspace_id HdinsightHadoopCluster#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#primary_key HdinsightHadoopCluster#primary_key}.
	PrimaryKey *string `field:"required" json:"primaryKey" yaml:"primaryKey"`
}

type HdinsightHadoopClusterMonitorOutputReference interface {
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
	InternalValue() *HdinsightHadoopClusterMonitor
	SetInternalValue(val *HdinsightHadoopClusterMonitor)
	LogAnalyticsWorkspaceId() *string
	SetLogAnalyticsWorkspaceId(val *string)
	LogAnalyticsWorkspaceIdInput() *string
	PrimaryKey() *string
	SetPrimaryKey(val *string)
	PrimaryKeyInput() *string
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

// The jsii proxy struct for HdinsightHadoopClusterMonitorOutputReference
type jsiiProxy_HdinsightHadoopClusterMonitorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) InternalValue() *HdinsightHadoopClusterMonitor {
	var returns *HdinsightHadoopClusterMonitor
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) PrimaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) PrimaryKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterMonitorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterMonitorOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterMonitorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterMonitorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterMonitorOutputReference_Override(h HdinsightHadoopClusterMonitorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterMonitorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) SetInternalValue(val *HdinsightHadoopClusterMonitor) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) SetLogAnalyticsWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) SetPrimaryKey(val *string) {
	_jsii_.Set(
		j,
		"primaryKey",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterMonitorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterNetwork struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#connection_direction HdinsightHadoopCluster#connection_direction}.
	ConnectionDirection *string `field:"optional" json:"connectionDirection" yaml:"connectionDirection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#private_link_enabled HdinsightHadoopCluster#private_link_enabled}.
	PrivateLinkEnabled interface{} `field:"optional" json:"privateLinkEnabled" yaml:"privateLinkEnabled"`
}

type HdinsightHadoopClusterNetworkOutputReference interface {
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
	ConnectionDirection() *string
	SetConnectionDirection(val *string)
	ConnectionDirectionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *HdinsightHadoopClusterNetwork
	SetInternalValue(val *HdinsightHadoopClusterNetwork)
	PrivateLinkEnabled() interface{}
	SetPrivateLinkEnabled(val interface{})
	PrivateLinkEnabledInput() interface{}
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
	ResetConnectionDirection()
	ResetPrivateLinkEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterNetworkOutputReference
type jsiiProxy_HdinsightHadoopClusterNetworkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) ConnectionDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) ConnectionDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) InternalValue() *HdinsightHadoopClusterNetwork {
	var returns *HdinsightHadoopClusterNetwork
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) PrivateLinkEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateLinkEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) PrivateLinkEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateLinkEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterNetworkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterNetworkOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterNetworkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterNetworkOutputReference_Override(h HdinsightHadoopClusterNetworkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) SetConnectionDirection(val *string) {
	_jsii_.Set(
		j,
		"connectionDirection",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) SetInternalValue(val *HdinsightHadoopClusterNetwork) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) SetPrivateLinkEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"privateLinkEnabled",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) ResetConnectionDirection() {
	_jsii_.InvokeVoid(
		h,
		"resetConnectionDirection",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) ResetPrivateLinkEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetPrivateLinkEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRoles struct {
	// head_node block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#head_node HdinsightHadoopCluster#head_node}
	HeadNode *HdinsightHadoopClusterRolesHeadNode `field:"required" json:"headNode" yaml:"headNode"`
	// worker_node block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#worker_node HdinsightHadoopCluster#worker_node}
	WorkerNode *HdinsightHadoopClusterRolesWorkerNode `field:"required" json:"workerNode" yaml:"workerNode"`
	// zookeeper_node block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#zookeeper_node HdinsightHadoopCluster#zookeeper_node}
	ZookeeperNode *HdinsightHadoopClusterRolesZookeeperNode `field:"required" json:"zookeeperNode" yaml:"zookeeperNode"`
	// edge_node block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#edge_node HdinsightHadoopCluster#edge_node}
	EdgeNode *HdinsightHadoopClusterRolesEdgeNode `field:"optional" json:"edgeNode" yaml:"edgeNode"`
}

type HdinsightHadoopClusterRolesEdgeNode struct {
	// install_script_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#install_script_action HdinsightHadoopCluster#install_script_action}
	InstallScriptAction interface{} `field:"required" json:"installScriptAction" yaml:"installScriptAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#target_instance_count HdinsightHadoopCluster#target_instance_count}.
	TargetInstanceCount *float64 `field:"required" json:"targetInstanceCount" yaml:"targetInstanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#vm_size HdinsightHadoopCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// https_endpoints block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#https_endpoints HdinsightHadoopCluster#https_endpoints}
	HttpsEndpoints interface{} `field:"optional" json:"httpsEndpoints" yaml:"httpsEndpoints"`
	// uninstall_script_actions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#uninstall_script_actions HdinsightHadoopCluster#uninstall_script_actions}
	UninstallScriptActions interface{} `field:"optional" json:"uninstallScriptActions" yaml:"uninstallScriptActions"`
}

type HdinsightHadoopClusterRolesEdgeNodeHttpsEndpoints struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#access_modes HdinsightHadoopCluster#access_modes}.
	AccessModes *[]*string `field:"optional" json:"accessModes" yaml:"accessModes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#destination_port HdinsightHadoopCluster#destination_port}.
	DestinationPort *float64 `field:"optional" json:"destinationPort" yaml:"destinationPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#disable_gateway_auth HdinsightHadoopCluster#disable_gateway_auth}.
	DisableGatewayAuth interface{} `field:"optional" json:"disableGatewayAuth" yaml:"disableGatewayAuth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#private_ip_address HdinsightHadoopCluster#private_ip_address}.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#sub_domain_suffix HdinsightHadoopCluster#sub_domain_suffix}.
	SubDomainSuffix *string `field:"optional" json:"subDomainSuffix" yaml:"subDomainSuffix"`
}

type HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList interface {
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
	Get(index *float64) HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList
type jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList_Override(h HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) Get(index *float64) HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference {
	var returns HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference interface {
	cdktf.ComplexObject
	AccessModes() *[]*string
	SetAccessModes(val *[]*string)
	AccessModesInput() *[]*string
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
	DestinationPort() *float64
	SetDestinationPort(val *float64)
	DestinationPortInput() *float64
	DisableGatewayAuth() interface{}
	SetDisableGatewayAuth(val interface{})
	DisableGatewayAuthInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrivateIpAddress() *string
	SetPrivateIpAddress(val *string)
	PrivateIpAddressInput() *string
	SubDomainSuffix() *string
	SetSubDomainSuffix(val *string)
	SubDomainSuffixInput() *string
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
	ResetAccessModes()
	ResetDestinationPort()
	ResetDisableGatewayAuth()
	ResetPrivateIpAddress()
	ResetSubDomainSuffix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference
type jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) AccessModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) AccessModesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessModesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) DestinationPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) DestinationPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"destinationPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) DisableGatewayAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableGatewayAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) DisableGatewayAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableGatewayAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) PrivateIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) SubDomainSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subDomainSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) SubDomainSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subDomainSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference_Override(h HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) SetAccessModes(val *[]*string) {
	_jsii_.Set(
		j,
		"accessModes",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) SetDestinationPort(val *float64) {
	_jsii_.Set(
		j,
		"destinationPort",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) SetDisableGatewayAuth(val interface{}) {
	_jsii_.Set(
		j,
		"disableGatewayAuth",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) SetPrivateIpAddress(val *string) {
	_jsii_.Set(
		j,
		"privateIpAddress",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) SetSubDomainSuffix(val *string) {
	_jsii_.Set(
		j,
		"subDomainSuffix",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) ResetAccessModes() {
	_jsii_.InvokeVoid(
		h,
		"resetAccessModes",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) ResetDestinationPort() {
	_jsii_.InvokeVoid(
		h,
		"resetDestinationPort",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) ResetDisableGatewayAuth() {
	_jsii_.InvokeVoid(
		h,
		"resetDisableGatewayAuth",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) ResetPrivateIpAddress() {
	_jsii_.InvokeVoid(
		h,
		"resetPrivateIpAddress",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) ResetSubDomainSuffix() {
	_jsii_.InvokeVoid(
		h,
		"resetSubDomainSuffix",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesEdgeNodeInstallScriptAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#name HdinsightHadoopCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#uri HdinsightHadoopCluster#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#parameters HdinsightHadoopCluster#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

type HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList interface {
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
	Get(index *float64) HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList
type jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList_Override(h HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) Get(index *float64) HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference {
	var returns HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference interface {
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
	Parameters() *string
	SetParameters(val *string)
	ParametersInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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

// The jsii proxy struct for HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference
type jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference_Override(h HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) SetParameters(val *string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) SetUri(val *string) {
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		h,
		"resetParameters",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesEdgeNodeOutputReference interface {
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
	HttpsEndpoints() HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList
	HttpsEndpointsInput() interface{}
	InstallScriptAction() HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList
	InstallScriptActionInput() interface{}
	InternalValue() *HdinsightHadoopClusterRolesEdgeNode
	SetInternalValue(val *HdinsightHadoopClusterRolesEdgeNode)
	TargetInstanceCount() *float64
	SetTargetInstanceCount(val *float64)
	TargetInstanceCountInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UninstallScriptActions() HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList
	UninstallScriptActionsInput() interface{}
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
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
	PutHttpsEndpoints(value interface{})
	PutInstallScriptAction(value interface{})
	PutUninstallScriptActions(value interface{})
	ResetHttpsEndpoints()
	ResetUninstallScriptActions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterRolesEdgeNodeOutputReference
type jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) HttpsEndpoints() HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList {
	var returns HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList
	_jsii_.Get(
		j,
		"httpsEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) HttpsEndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) InstallScriptAction() HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList {
	var returns HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList
	_jsii_.Get(
		j,
		"installScriptAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) InstallScriptActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installScriptActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) InternalValue() *HdinsightHadoopClusterRolesEdgeNode {
	var returns *HdinsightHadoopClusterRolesEdgeNode
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) TargetInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) TargetInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) UninstallScriptActions() HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList {
	var returns HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList
	_jsii_.Get(
		j,
		"uninstallScriptActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) UninstallScriptActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uninstallScriptActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesEdgeNodeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterRolesEdgeNodeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesEdgeNodeOutputReference_Override(h HdinsightHadoopClusterRolesEdgeNodeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) SetInternalValue(val *HdinsightHadoopClusterRolesEdgeNode) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) SetTargetInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"targetInstanceCount",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) SetVmSize(val *string) {
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) PutHttpsEndpoints(value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"putHttpsEndpoints",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) PutInstallScriptAction(value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"putInstallScriptAction",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) PutUninstallScriptActions(value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"putUninstallScriptActions",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) ResetHttpsEndpoints() {
	_jsii_.InvokeVoid(
		h,
		"resetHttpsEndpoints",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) ResetUninstallScriptActions() {
	_jsii_.InvokeVoid(
		h,
		"resetUninstallScriptActions",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#name HdinsightHadoopCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#uri HdinsightHadoopCluster#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#parameters HdinsightHadoopCluster#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

type HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList interface {
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
	Get(index *float64) HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList
type jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList_Override(h HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) Get(index *float64) HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference {
	var returns HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference interface {
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
	Parameters() *string
	SetParameters(val *string)
	ParametersInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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

// The jsii proxy struct for HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference
type jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) Parameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) ParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference_Override(h HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) SetParameters(val *string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) SetUri(val *string) {
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		h,
		"resetParameters",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesHeadNode struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#username HdinsightHadoopCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#vm_size HdinsightHadoopCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#password HdinsightHadoopCluster#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#ssh_keys HdinsightHadoopCluster#ssh_keys}.
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#subnet_id HdinsightHadoopCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#virtual_network_id HdinsightHadoopCluster#virtual_network_id}.
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

type HdinsightHadoopClusterRolesHeadNodeOutputReference interface {
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
	InternalValue() *HdinsightHadoopClusterRolesHeadNode
	SetInternalValue(val *HdinsightHadoopClusterRolesHeadNode)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	SshKeys() *[]*string
	SetSshKeys(val *[]*string)
	SshKeysInput() *[]*string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	VirtualNetworkId() *string
	SetVirtualNetworkId(val *string)
	VirtualNetworkIdInput() *string
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
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
	ResetPassword()
	ResetSshKeys()
	ResetSubnetId()
	ResetVirtualNetworkId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterRolesHeadNodeOutputReference
type jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) InternalValue() *HdinsightHadoopClusterRolesHeadNode {
	var returns *HdinsightHadoopClusterRolesHeadNode
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SshKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesHeadNodeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterRolesHeadNodeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesHeadNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesHeadNodeOutputReference_Override(h HdinsightHadoopClusterRolesHeadNodeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesHeadNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SetInternalValue(val *HdinsightHadoopClusterRolesHeadNode) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SetSshKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"sshKeys",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SetVirtualNetworkId(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) SetVmSize(val *string) {
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		h,
		"resetPassword",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) ResetSshKeys() {
	_jsii_.InvokeVoid(
		h,
		"resetSshKeys",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		h,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		h,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesHeadNodeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesOutputReference interface {
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
	EdgeNode() HdinsightHadoopClusterRolesEdgeNodeOutputReference
	EdgeNodeInput() *HdinsightHadoopClusterRolesEdgeNode
	// Experimental.
	Fqn() *string
	HeadNode() HdinsightHadoopClusterRolesHeadNodeOutputReference
	HeadNodeInput() *HdinsightHadoopClusterRolesHeadNode
	InternalValue() *HdinsightHadoopClusterRoles
	SetInternalValue(val *HdinsightHadoopClusterRoles)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkerNode() HdinsightHadoopClusterRolesWorkerNodeOutputReference
	WorkerNodeInput() *HdinsightHadoopClusterRolesWorkerNode
	ZookeeperNode() HdinsightHadoopClusterRolesZookeeperNodeOutputReference
	ZookeeperNodeInput() *HdinsightHadoopClusterRolesZookeeperNode
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
	PutEdgeNode(value *HdinsightHadoopClusterRolesEdgeNode)
	PutHeadNode(value *HdinsightHadoopClusterRolesHeadNode)
	PutWorkerNode(value *HdinsightHadoopClusterRolesWorkerNode)
	PutZookeeperNode(value *HdinsightHadoopClusterRolesZookeeperNode)
	ResetEdgeNode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterRolesOutputReference
type jsiiProxy_HdinsightHadoopClusterRolesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) EdgeNode() HdinsightHadoopClusterRolesEdgeNodeOutputReference {
	var returns HdinsightHadoopClusterRolesEdgeNodeOutputReference
	_jsii_.Get(
		j,
		"edgeNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) EdgeNodeInput() *HdinsightHadoopClusterRolesEdgeNode {
	var returns *HdinsightHadoopClusterRolesEdgeNode
	_jsii_.Get(
		j,
		"edgeNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) HeadNode() HdinsightHadoopClusterRolesHeadNodeOutputReference {
	var returns HdinsightHadoopClusterRolesHeadNodeOutputReference
	_jsii_.Get(
		j,
		"headNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) HeadNodeInput() *HdinsightHadoopClusterRolesHeadNode {
	var returns *HdinsightHadoopClusterRolesHeadNode
	_jsii_.Get(
		j,
		"headNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) InternalValue() *HdinsightHadoopClusterRoles {
	var returns *HdinsightHadoopClusterRoles
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) WorkerNode() HdinsightHadoopClusterRolesWorkerNodeOutputReference {
	var returns HdinsightHadoopClusterRolesWorkerNodeOutputReference
	_jsii_.Get(
		j,
		"workerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) WorkerNodeInput() *HdinsightHadoopClusterRolesWorkerNode {
	var returns *HdinsightHadoopClusterRolesWorkerNode
	_jsii_.Get(
		j,
		"workerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) ZookeeperNode() HdinsightHadoopClusterRolesZookeeperNodeOutputReference {
	var returns HdinsightHadoopClusterRolesZookeeperNodeOutputReference
	_jsii_.Get(
		j,
		"zookeeperNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) ZookeeperNodeInput() *HdinsightHadoopClusterRolesZookeeperNode {
	var returns *HdinsightHadoopClusterRolesZookeeperNode
	_jsii_.Get(
		j,
		"zookeeperNodeInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterRolesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesOutputReference_Override(h HdinsightHadoopClusterRolesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) SetInternalValue(val *HdinsightHadoopClusterRoles) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) PutEdgeNode(value *HdinsightHadoopClusterRolesEdgeNode) {
	_jsii_.InvokeVoid(
		h,
		"putEdgeNode",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) PutHeadNode(value *HdinsightHadoopClusterRolesHeadNode) {
	_jsii_.InvokeVoid(
		h,
		"putHeadNode",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) PutWorkerNode(value *HdinsightHadoopClusterRolesWorkerNode) {
	_jsii_.InvokeVoid(
		h,
		"putWorkerNode",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) PutZookeeperNode(value *HdinsightHadoopClusterRolesZookeeperNode) {
	_jsii_.InvokeVoid(
		h,
		"putZookeeperNode",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) ResetEdgeNode() {
	_jsii_.InvokeVoid(
		h,
		"resetEdgeNode",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesWorkerNode struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#target_instance_count HdinsightHadoopCluster#target_instance_count}.
	TargetInstanceCount *float64 `field:"required" json:"targetInstanceCount" yaml:"targetInstanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#username HdinsightHadoopCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#vm_size HdinsightHadoopCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// autoscale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#autoscale HdinsightHadoopCluster#autoscale}
	Autoscale *HdinsightHadoopClusterRolesWorkerNodeAutoscale `field:"optional" json:"autoscale" yaml:"autoscale"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#password HdinsightHadoopCluster#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#ssh_keys HdinsightHadoopCluster#ssh_keys}.
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#subnet_id HdinsightHadoopCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#virtual_network_id HdinsightHadoopCluster#virtual_network_id}.
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

type HdinsightHadoopClusterRolesWorkerNodeAutoscale struct {
	// capacity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#capacity HdinsightHadoopCluster#capacity}
	Capacity *HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity `field:"optional" json:"capacity" yaml:"capacity"`
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#recurrence HdinsightHadoopCluster#recurrence}
	Recurrence *HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence `field:"optional" json:"recurrence" yaml:"recurrence"`
}

type HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#max_instance_count HdinsightHadoopCluster#max_instance_count}.
	MaxInstanceCount *float64 `field:"required" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#min_instance_count HdinsightHadoopCluster#min_instance_count}.
	MinInstanceCount *float64 `field:"required" json:"minInstanceCount" yaml:"minInstanceCount"`
}

type HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference interface {
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
	InternalValue() *HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity
	SetInternalValue(val *HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity)
	MaxInstanceCount() *float64
	SetMaxInstanceCount(val *float64)
	MaxInstanceCountInput() *float64
	MinInstanceCount() *float64
	SetMinInstanceCount(val *float64)
	MinInstanceCountInput() *float64
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

// The jsii proxy struct for HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference
type jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) InternalValue() *HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity {
	var returns *HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) MaxInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) MaxInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) MinInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) MinInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference_Override(h HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetInternalValue(val *HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetMaxInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"maxInstanceCount",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetMinInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"minInstanceCount",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference interface {
	cdktf.ComplexObject
	Capacity() HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference
	CapacityInput() *HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity
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
	InternalValue() *HdinsightHadoopClusterRolesWorkerNodeAutoscale
	SetInternalValue(val *HdinsightHadoopClusterRolesWorkerNodeAutoscale)
	Recurrence() HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference
	RecurrenceInput() *HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence
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
	PutCapacity(value *HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity)
	PutRecurrence(value *HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence)
	ResetCapacity()
	ResetRecurrence()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference
type jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) Capacity() HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference {
	var returns HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacityOutputReference
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) CapacityInput() *HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity {
	var returns *HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) InternalValue() *HdinsightHadoopClusterRolesWorkerNodeAutoscale {
	var returns *HdinsightHadoopClusterRolesWorkerNodeAutoscale
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) Recurrence() HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference {
	var returns HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference
	_jsii_.Get(
		j,
		"recurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) RecurrenceInput() *HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence {
	var returns *HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence
	_jsii_.Get(
		j,
		"recurrenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference_Override(h HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) SetInternalValue(val *HdinsightHadoopClusterRolesWorkerNodeAutoscale) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) PutCapacity(value *HdinsightHadoopClusterRolesWorkerNodeAutoscaleCapacity) {
	_jsii_.InvokeVoid(
		h,
		"putCapacity",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) PutRecurrence(value *HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence) {
	_jsii_.InvokeVoid(
		h,
		"putRecurrence",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) ResetCapacity() {
	_jsii_.InvokeVoid(
		h,
		"resetCapacity",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) ResetRecurrence() {
	_jsii_.InvokeVoid(
		h,
		"resetRecurrence",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence struct {
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#schedule HdinsightHadoopCluster#schedule}
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#timezone HdinsightHadoopCluster#timezone}.
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

type HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference interface {
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
	InternalValue() *HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence
	SetInternalValue(val *HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence)
	Schedule() HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList
	ScheduleInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
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
	PutSchedule(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference
type jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) InternalValue() *HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence {
	var returns *HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) Schedule() HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList {
	var returns HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference_Override(h HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) SetInternalValue(val *HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrence) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) SetTimezone(val *string) {
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) PutSchedule(value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"putSchedule",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#days HdinsightHadoopCluster#days}.
	Days *[]*string `field:"required" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#target_instance_count HdinsightHadoopCluster#target_instance_count}.
	TargetInstanceCount *float64 `field:"required" json:"targetInstanceCount" yaml:"targetInstanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#time HdinsightHadoopCluster#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
}

type HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList interface {
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
	Get(index *float64) HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList
type jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList_Override(h HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) Get(index *float64) HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference {
	var returns HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference interface {
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
	Days() *[]*string
	SetDays(val *[]*string)
	DaysInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	TargetInstanceCount() *float64
	SetTargetInstanceCount(val *float64)
	TargetInstanceCountInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Time() *string
	SetTime(val *string)
	TimeInput() *string
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

// The jsii proxy struct for HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference
type jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) Days() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"days",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) DaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) TargetInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) TargetInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) Time() *string {
	var returns *string
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) TimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference_Override(h HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetDays(val *[]*string) {
	_jsii_.Set(
		j,
		"days",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetTargetInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"targetInstanceCount",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetTime(val *string) {
	_jsii_.Set(
		j,
		"time",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesWorkerNodeOutputReference interface {
	cdktf.ComplexObject
	Autoscale() HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference
	AutoscaleInput() *HdinsightHadoopClusterRolesWorkerNodeAutoscale
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
	InternalValue() *HdinsightHadoopClusterRolesWorkerNode
	SetInternalValue(val *HdinsightHadoopClusterRolesWorkerNode)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	SshKeys() *[]*string
	SetSshKeys(val *[]*string)
	SshKeysInput() *[]*string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	TargetInstanceCount() *float64
	SetTargetInstanceCount(val *float64)
	TargetInstanceCountInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	VirtualNetworkId() *string
	SetVirtualNetworkId(val *string)
	VirtualNetworkIdInput() *string
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
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
	PutAutoscale(value *HdinsightHadoopClusterRolesWorkerNodeAutoscale)
	ResetAutoscale()
	ResetPassword()
	ResetSshKeys()
	ResetSubnetId()
	ResetVirtualNetworkId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterRolesWorkerNodeOutputReference
type jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) Autoscale() HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference {
	var returns HdinsightHadoopClusterRolesWorkerNodeAutoscaleOutputReference
	_jsii_.Get(
		j,
		"autoscale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) AutoscaleInput() *HdinsightHadoopClusterRolesWorkerNodeAutoscale {
	var returns *HdinsightHadoopClusterRolesWorkerNodeAutoscale
	_jsii_.Get(
		j,
		"autoscaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) InternalValue() *HdinsightHadoopClusterRolesWorkerNode {
	var returns *HdinsightHadoopClusterRolesWorkerNode
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SshKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) TargetInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) TargetInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesWorkerNodeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterRolesWorkerNodeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesWorkerNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesWorkerNodeOutputReference_Override(h HdinsightHadoopClusterRolesWorkerNodeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesWorkerNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SetInternalValue(val *HdinsightHadoopClusterRolesWorkerNode) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SetSshKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"sshKeys",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SetTargetInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"targetInstanceCount",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SetVirtualNetworkId(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) SetVmSize(val *string) {
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) PutAutoscale(value *HdinsightHadoopClusterRolesWorkerNodeAutoscale) {
	_jsii_.InvokeVoid(
		h,
		"putAutoscale",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) ResetAutoscale() {
	_jsii_.InvokeVoid(
		h,
		"resetAutoscale",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		h,
		"resetPassword",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) ResetSshKeys() {
	_jsii_.InvokeVoid(
		h,
		"resetSshKeys",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		h,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		h,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesWorkerNodeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterRolesZookeeperNode struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#username HdinsightHadoopCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#vm_size HdinsightHadoopCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#password HdinsightHadoopCluster#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#ssh_keys HdinsightHadoopCluster#ssh_keys}.
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#subnet_id HdinsightHadoopCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#virtual_network_id HdinsightHadoopCluster#virtual_network_id}.
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

type HdinsightHadoopClusterRolesZookeeperNodeOutputReference interface {
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
	InternalValue() *HdinsightHadoopClusterRolesZookeeperNode
	SetInternalValue(val *HdinsightHadoopClusterRolesZookeeperNode)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	SshKeys() *[]*string
	SetSshKeys(val *[]*string)
	SshKeysInput() *[]*string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	VirtualNetworkId() *string
	SetVirtualNetworkId(val *string)
	VirtualNetworkIdInput() *string
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
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
	ResetPassword()
	ResetSshKeys()
	ResetSubnetId()
	ResetVirtualNetworkId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterRolesZookeeperNodeOutputReference
type jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) InternalValue() *HdinsightHadoopClusterRolesZookeeperNode {
	var returns *HdinsightHadoopClusterRolesZookeeperNode
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SshKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesZookeeperNodeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterRolesZookeeperNodeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesZookeeperNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesZookeeperNodeOutputReference_Override(h HdinsightHadoopClusterRolesZookeeperNodeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesZookeeperNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SetInternalValue(val *HdinsightHadoopClusterRolesZookeeperNode) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SetSshKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"sshKeys",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SetVirtualNetworkId(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) SetVmSize(val *string) {
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		h,
		"resetPassword",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) ResetSshKeys() {
	_jsii_.InvokeVoid(
		h,
		"resetSshKeys",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		h,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		h,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesZookeeperNodeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterSecurityProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#aadds_resource_id HdinsightHadoopCluster#aadds_resource_id}.
	AaddsResourceId *string `field:"required" json:"aaddsResourceId" yaml:"aaddsResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#domain_name HdinsightHadoopCluster#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#domain_username HdinsightHadoopCluster#domain_username}.
	DomainUsername *string `field:"required" json:"domainUsername" yaml:"domainUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#domain_user_password HdinsightHadoopCluster#domain_user_password}.
	DomainUserPassword *string `field:"required" json:"domainUserPassword" yaml:"domainUserPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#ldaps_urls HdinsightHadoopCluster#ldaps_urls}.
	LdapsUrls *[]*string `field:"required" json:"ldapsUrls" yaml:"ldapsUrls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#msi_resource_id HdinsightHadoopCluster#msi_resource_id}.
	MsiResourceId *string `field:"required" json:"msiResourceId" yaml:"msiResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#cluster_users_group_dns HdinsightHadoopCluster#cluster_users_group_dns}.
	ClusterUsersGroupDns *[]*string `field:"optional" json:"clusterUsersGroupDns" yaml:"clusterUsersGroupDns"`
}

type HdinsightHadoopClusterSecurityProfileOutputReference interface {
	cdktf.ComplexObject
	AaddsResourceId() *string
	SetAaddsResourceId(val *string)
	AaddsResourceIdInput() *string
	ClusterUsersGroupDns() *[]*string
	SetClusterUsersGroupDns(val *[]*string)
	ClusterUsersGroupDnsInput() *[]*string
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
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	DomainUsername() *string
	SetDomainUsername(val *string)
	DomainUsernameInput() *string
	DomainUserPassword() *string
	SetDomainUserPassword(val *string)
	DomainUserPasswordInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *HdinsightHadoopClusterSecurityProfile
	SetInternalValue(val *HdinsightHadoopClusterSecurityProfile)
	LdapsUrls() *[]*string
	SetLdapsUrls(val *[]*string)
	LdapsUrlsInput() *[]*string
	MsiResourceId() *string
	SetMsiResourceId(val *string)
	MsiResourceIdInput() *string
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
	ResetClusterUsersGroupDns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterSecurityProfileOutputReference
type jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) AaddsResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aaddsResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) AaddsResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aaddsResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) ClusterUsersGroupDns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterUsersGroupDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) ClusterUsersGroupDnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterUsersGroupDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) DomainUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) DomainUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) DomainUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) DomainUserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) InternalValue() *HdinsightHadoopClusterSecurityProfile {
	var returns *HdinsightHadoopClusterSecurityProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) LdapsUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ldapsUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) LdapsUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ldapsUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) MsiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) MsiResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterSecurityProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterSecurityProfileOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterSecurityProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterSecurityProfileOutputReference_Override(h HdinsightHadoopClusterSecurityProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterSecurityProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) SetAaddsResourceId(val *string) {
	_jsii_.Set(
		j,
		"aaddsResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) SetClusterUsersGroupDns(val *[]*string) {
	_jsii_.Set(
		j,
		"clusterUsersGroupDns",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) SetDomainUsername(val *string) {
	_jsii_.Set(
		j,
		"domainUsername",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) SetDomainUserPassword(val *string) {
	_jsii_.Set(
		j,
		"domainUserPassword",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) SetInternalValue(val *HdinsightHadoopClusterSecurityProfile) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) SetLdapsUrls(val *[]*string) {
	_jsii_.Set(
		j,
		"ldapsUrls",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) SetMsiResourceId(val *string) {
	_jsii_.Set(
		j,
		"msiResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) ResetClusterUsersGroupDns() {
	_jsii_.InvokeVoid(
		h,
		"resetClusterUsersGroupDns",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterSecurityProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterStorageAccount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#is_default HdinsightHadoopCluster#is_default}.
	IsDefault interface{} `field:"required" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#storage_account_key HdinsightHadoopCluster#storage_account_key}.
	StorageAccountKey *string `field:"required" json:"storageAccountKey" yaml:"storageAccountKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#storage_container_id HdinsightHadoopCluster#storage_container_id}.
	StorageContainerId *string `field:"required" json:"storageContainerId" yaml:"storageContainerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#storage_resource_id HdinsightHadoopCluster#storage_resource_id}.
	StorageResourceId *string `field:"optional" json:"storageResourceId" yaml:"storageResourceId"`
}

type HdinsightHadoopClusterStorageAccountGen2 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#filesystem_id HdinsightHadoopCluster#filesystem_id}.
	FilesystemId *string `field:"required" json:"filesystemId" yaml:"filesystemId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#is_default HdinsightHadoopCluster#is_default}.
	IsDefault interface{} `field:"required" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#managed_identity_resource_id HdinsightHadoopCluster#managed_identity_resource_id}.
	ManagedIdentityResourceId *string `field:"required" json:"managedIdentityResourceId" yaml:"managedIdentityResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#storage_resource_id HdinsightHadoopCluster#storage_resource_id}.
	StorageResourceId *string `field:"required" json:"storageResourceId" yaml:"storageResourceId"`
}

type HdinsightHadoopClusterStorageAccountGen2OutputReference interface {
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
	FilesystemId() *string
	SetFilesystemId(val *string)
	FilesystemIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *HdinsightHadoopClusterStorageAccountGen2
	SetInternalValue(val *HdinsightHadoopClusterStorageAccountGen2)
	IsDefault() interface{}
	SetIsDefault(val interface{})
	IsDefaultInput() interface{}
	ManagedIdentityResourceId() *string
	SetManagedIdentityResourceId(val *string)
	ManagedIdentityResourceIdInput() *string
	StorageResourceId() *string
	SetStorageResourceId(val *string)
	StorageResourceIdInput() *string
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

// The jsii proxy struct for HdinsightHadoopClusterStorageAccountGen2OutputReference
type jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) FilesystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filesystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) FilesystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filesystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) InternalValue() *HdinsightHadoopClusterStorageAccountGen2 {
	var returns *HdinsightHadoopClusterStorageAccountGen2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) IsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) IsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) ManagedIdentityResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) ManagedIdentityResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) StorageResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) StorageResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterStorageAccountGen2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterStorageAccountGen2OutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterStorageAccountGen2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterStorageAccountGen2OutputReference_Override(h HdinsightHadoopClusterStorageAccountGen2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterStorageAccountGen2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) SetFilesystemId(val *string) {
	_jsii_.Set(
		j,
		"filesystemId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) SetInternalValue(val *HdinsightHadoopClusterStorageAccountGen2) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) SetIsDefault(val interface{}) {
	_jsii_.Set(
		j,
		"isDefault",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) SetManagedIdentityResourceId(val *string) {
	_jsii_.Set(
		j,
		"managedIdentityResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) SetStorageResourceId(val *string) {
	_jsii_.Set(
		j,
		"storageResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountGen2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterStorageAccountList interface {
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
	Get(index *float64) HdinsightHadoopClusterStorageAccountOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterStorageAccountList
type jsiiProxy_HdinsightHadoopClusterStorageAccountList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterStorageAccountList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HdinsightHadoopClusterStorageAccountList {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterStorageAccountList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterStorageAccountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterStorageAccountList_Override(h HdinsightHadoopClusterStorageAccountList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterStorageAccountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountList) Get(index *float64) HdinsightHadoopClusterStorageAccountOutputReference {
	var returns HdinsightHadoopClusterStorageAccountOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterStorageAccountOutputReference interface {
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
	IsDefault() interface{}
	SetIsDefault(val interface{})
	IsDefaultInput() interface{}
	StorageAccountKey() *string
	SetStorageAccountKey(val *string)
	StorageAccountKeyInput() *string
	StorageContainerId() *string
	SetStorageContainerId(val *string)
	StorageContainerIdInput() *string
	StorageResourceId() *string
	SetStorageResourceId(val *string)
	StorageResourceIdInput() *string
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
	ResetStorageResourceId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterStorageAccountOutputReference
type jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) IsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) IsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) StorageAccountKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) StorageAccountKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) StorageContainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) StorageContainerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) StorageResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) StorageResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterStorageAccountOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HdinsightHadoopClusterStorageAccountOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterStorageAccountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterStorageAccountOutputReference_Override(h HdinsightHadoopClusterStorageAccountOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterStorageAccountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) SetIsDefault(val interface{}) {
	_jsii_.Set(
		j,
		"isDefault",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) SetStorageAccountKey(val *string) {
	_jsii_.Set(
		j,
		"storageAccountKey",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) SetStorageContainerId(val *string) {
	_jsii_.Set(
		j,
		"storageContainerId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) SetStorageResourceId(val *string) {
	_jsii_.Set(
		j,
		"storageResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) ResetStorageResourceId() {
	_jsii_.InvokeVoid(
		h,
		"resetStorageResourceId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterStorageAccountOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightHadoopClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#create HdinsightHadoopCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#delete HdinsightHadoopCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#read HdinsightHadoopCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hadoop_cluster#update HdinsightHadoopCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type HdinsightHadoopClusterTimeoutsOutputReference interface {
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

// The jsii proxy struct for HdinsightHadoopClusterTimeoutsOutputReference
type jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterTimeoutsOutputReference_Override(h HdinsightHadoopClusterTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		h,
		"resetCreate",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		h,
		"resetDelete",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		h,
		"resetRead",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		h,
		"resetUpdate",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

