package hdinsightinteractivequerycluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/hdinsightinteractivequerycluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster azurerm_hdinsight_interactive_query_cluster}.
type HdinsightInteractiveQueryCluster interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterVersion() *string
	SetClusterVersion(val *string)
	ClusterVersionInput() *string
	ComponentVersion() HdinsightInteractiveQueryClusterComponentVersionOutputReference
	ComponentVersionInput() *HdinsightInteractiveQueryClusterComponentVersion
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
	DiskEncryption() HdinsightInteractiveQueryClusterDiskEncryptionList
	DiskEncryptionInput() interface{}
	EncryptionInTransitEnabled() interface{}
	SetEncryptionInTransitEnabled(val interface{})
	EncryptionInTransitEnabledInput() interface{}
	Extension() HdinsightInteractiveQueryClusterExtensionOutputReference
	ExtensionInput() *HdinsightInteractiveQueryClusterExtension
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Gateway() HdinsightInteractiveQueryClusterGatewayOutputReference
	GatewayInput() *HdinsightInteractiveQueryClusterGateway
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
	Metastores() HdinsightInteractiveQueryClusterMetastoresOutputReference
	MetastoresInput() *HdinsightInteractiveQueryClusterMetastores
	Monitor() HdinsightInteractiveQueryClusterMonitorOutputReference
	MonitorInput() *HdinsightInteractiveQueryClusterMonitor
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() HdinsightInteractiveQueryClusterNetworkOutputReference
	NetworkInput() *HdinsightInteractiveQueryClusterNetwork
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
	Roles() HdinsightInteractiveQueryClusterRolesOutputReference
	RolesInput() *HdinsightInteractiveQueryClusterRoles
	SecurityProfile() HdinsightInteractiveQueryClusterSecurityProfileOutputReference
	SecurityProfileInput() *HdinsightInteractiveQueryClusterSecurityProfile
	SshEndpoint() *string
	StorageAccount() HdinsightInteractiveQueryClusterStorageAccountList
	StorageAccountGen2() HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference
	StorageAccountGen2Input() *HdinsightInteractiveQueryClusterStorageAccountGen2
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
	Timeouts() HdinsightInteractiveQueryClusterTimeoutsOutputReference
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
	PutComponentVersion(value *HdinsightInteractiveQueryClusterComponentVersion)
	PutDiskEncryption(value interface{})
	PutExtension(value *HdinsightInteractiveQueryClusterExtension)
	PutGateway(value *HdinsightInteractiveQueryClusterGateway)
	PutMetastores(value *HdinsightInteractiveQueryClusterMetastores)
	PutMonitor(value *HdinsightInteractiveQueryClusterMonitor)
	PutNetwork(value *HdinsightInteractiveQueryClusterNetwork)
	PutRoles(value *HdinsightInteractiveQueryClusterRoles)
	PutSecurityProfile(value *HdinsightInteractiveQueryClusterSecurityProfile)
	PutStorageAccount(value interface{})
	PutStorageAccountGen2(value *HdinsightInteractiveQueryClusterStorageAccountGen2)
	PutTimeouts(value *HdinsightInteractiveQueryClusterTimeouts)
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

// The jsii proxy struct for HdinsightInteractiveQueryCluster
type jsiiProxy_HdinsightInteractiveQueryCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) ClusterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) ClusterVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) ComponentVersion() HdinsightInteractiveQueryClusterComponentVersionOutputReference {
	var returns HdinsightInteractiveQueryClusterComponentVersionOutputReference
	_jsii_.Get(
		j,
		"componentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) ComponentVersionInput() *HdinsightInteractiveQueryClusterComponentVersion {
	var returns *HdinsightInteractiveQueryClusterComponentVersion
	_jsii_.Get(
		j,
		"componentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) DiskEncryption() HdinsightInteractiveQueryClusterDiskEncryptionList {
	var returns HdinsightInteractiveQueryClusterDiskEncryptionList
	_jsii_.Get(
		j,
		"diskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) DiskEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) EncryptionInTransitEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInTransitEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) EncryptionInTransitEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionInTransitEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Extension() HdinsightInteractiveQueryClusterExtensionOutputReference {
	var returns HdinsightInteractiveQueryClusterExtensionOutputReference
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) ExtensionInput() *HdinsightInteractiveQueryClusterExtension {
	var returns *HdinsightInteractiveQueryClusterExtension
	_jsii_.Get(
		j,
		"extensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Gateway() HdinsightInteractiveQueryClusterGatewayOutputReference {
	var returns HdinsightInteractiveQueryClusterGatewayOutputReference
	_jsii_.Get(
		j,
		"gateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) GatewayInput() *HdinsightInteractiveQueryClusterGateway {
	var returns *HdinsightInteractiveQueryClusterGateway
	_jsii_.Get(
		j,
		"gatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) HttpsEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Metastores() HdinsightInteractiveQueryClusterMetastoresOutputReference {
	var returns HdinsightInteractiveQueryClusterMetastoresOutputReference
	_jsii_.Get(
		j,
		"metastores",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) MetastoresInput() *HdinsightInteractiveQueryClusterMetastores {
	var returns *HdinsightInteractiveQueryClusterMetastores
	_jsii_.Get(
		j,
		"metastoresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Monitor() HdinsightInteractiveQueryClusterMonitorOutputReference {
	var returns HdinsightInteractiveQueryClusterMonitorOutputReference
	_jsii_.Get(
		j,
		"monitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) MonitorInput() *HdinsightInteractiveQueryClusterMonitor {
	var returns *HdinsightInteractiveQueryClusterMonitor
	_jsii_.Get(
		j,
		"monitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Network() HdinsightInteractiveQueryClusterNetworkOutputReference {
	var returns HdinsightInteractiveQueryClusterNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) NetworkInput() *HdinsightInteractiveQueryClusterNetwork {
	var returns *HdinsightInteractiveQueryClusterNetwork
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Roles() HdinsightInteractiveQueryClusterRolesOutputReference {
	var returns HdinsightInteractiveQueryClusterRolesOutputReference
	_jsii_.Get(
		j,
		"roles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) RolesInput() *HdinsightInteractiveQueryClusterRoles {
	var returns *HdinsightInteractiveQueryClusterRoles
	_jsii_.Get(
		j,
		"rolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SecurityProfile() HdinsightInteractiveQueryClusterSecurityProfileOutputReference {
	var returns HdinsightInteractiveQueryClusterSecurityProfileOutputReference
	_jsii_.Get(
		j,
		"securityProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SecurityProfileInput() *HdinsightInteractiveQueryClusterSecurityProfile {
	var returns *HdinsightInteractiveQueryClusterSecurityProfile
	_jsii_.Get(
		j,
		"securityProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SshEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) StorageAccount() HdinsightInteractiveQueryClusterStorageAccountList {
	var returns HdinsightInteractiveQueryClusterStorageAccountList
	_jsii_.Get(
		j,
		"storageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) StorageAccountGen2() HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference {
	var returns HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference
	_jsii_.Get(
		j,
		"storageAccountGen2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) StorageAccountGen2Input() *HdinsightInteractiveQueryClusterStorageAccountGen2 {
	var returns *HdinsightInteractiveQueryClusterStorageAccountGen2
	_jsii_.Get(
		j,
		"storageAccountGen2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) StorageAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) Timeouts() HdinsightInteractiveQueryClusterTimeoutsOutputReference {
	var returns HdinsightInteractiveQueryClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) TlsMinVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) TlsMinVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster azurerm_hdinsight_interactive_query_cluster} Resource.
func NewHdinsightInteractiveQueryCluster(scope constructs.Construct, id *string, config *HdinsightInteractiveQueryClusterConfig) HdinsightInteractiveQueryCluster {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster azurerm_hdinsight_interactive_query_cluster} Resource.
func NewHdinsightInteractiveQueryCluster_Override(h HdinsightInteractiveQueryCluster, scope constructs.Construct, id *string, config *HdinsightInteractiveQueryClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryCluster",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetClusterVersion(val *string) {
	_jsii_.Set(
		j,
		"clusterVersion",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetEncryptionInTransitEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionInTransitEnabled",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetTier(val *string) {
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryCluster) SetTlsMinVersion(val *string) {
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
func HdinsightInteractiveQueryCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HdinsightInteractiveQueryCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) PutComponentVersion(value *HdinsightInteractiveQueryClusterComponentVersion) {
	_jsii_.InvokeVoid(
		h,
		"putComponentVersion",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) PutDiskEncryption(value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"putDiskEncryption",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) PutExtension(value *HdinsightInteractiveQueryClusterExtension) {
	_jsii_.InvokeVoid(
		h,
		"putExtension",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) PutGateway(value *HdinsightInteractiveQueryClusterGateway) {
	_jsii_.InvokeVoid(
		h,
		"putGateway",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) PutMetastores(value *HdinsightInteractiveQueryClusterMetastores) {
	_jsii_.InvokeVoid(
		h,
		"putMetastores",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) PutMonitor(value *HdinsightInteractiveQueryClusterMonitor) {
	_jsii_.InvokeVoid(
		h,
		"putMonitor",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) PutNetwork(value *HdinsightInteractiveQueryClusterNetwork) {
	_jsii_.InvokeVoid(
		h,
		"putNetwork",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) PutRoles(value *HdinsightInteractiveQueryClusterRoles) {
	_jsii_.InvokeVoid(
		h,
		"putRoles",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) PutSecurityProfile(value *HdinsightInteractiveQueryClusterSecurityProfile) {
	_jsii_.InvokeVoid(
		h,
		"putSecurityProfile",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) PutStorageAccount(value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"putStorageAccount",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) PutStorageAccountGen2(value *HdinsightInteractiveQueryClusterStorageAccountGen2) {
	_jsii_.InvokeVoid(
		h,
		"putStorageAccountGen2",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) PutTimeouts(value *HdinsightInteractiveQueryClusterTimeouts) {
	_jsii_.InvokeVoid(
		h,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetDiskEncryption() {
	_jsii_.InvokeVoid(
		h,
		"resetDiskEncryption",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetEncryptionInTransitEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetEncryptionInTransitEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetExtension() {
	_jsii_.InvokeVoid(
		h,
		"resetExtension",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetMetastores() {
	_jsii_.InvokeVoid(
		h,
		"resetMetastores",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetMonitor() {
	_jsii_.InvokeVoid(
		h,
		"resetMonitor",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetNetwork() {
	_jsii_.InvokeVoid(
		h,
		"resetNetwork",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetSecurityProfile() {
	_jsii_.InvokeVoid(
		h,
		"resetSecurityProfile",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetStorageAccount() {
	_jsii_.InvokeVoid(
		h,
		"resetStorageAccount",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetStorageAccountGen2() {
	_jsii_.InvokeVoid(
		h,
		"resetStorageAccountGen2",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetTags() {
	_jsii_.InvokeVoid(
		h,
		"resetTags",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		h,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ResetTlsMinVersion() {
	_jsii_.InvokeVoid(
		h,
		"resetTlsMinVersion",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterComponentVersion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#interactive_hive HdinsightInteractiveQueryCluster#interactive_hive}.
	InteractiveHive *string `field:"required" json:"interactiveHive" yaml:"interactiveHive"`
}

type HdinsightInteractiveQueryClusterComponentVersionOutputReference interface {
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
	InteractiveHive() *string
	SetInteractiveHive(val *string)
	InteractiveHiveInput() *string
	InternalValue() *HdinsightInteractiveQueryClusterComponentVersion
	SetInternalValue(val *HdinsightInteractiveQueryClusterComponentVersion)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterComponentVersionOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) InteractiveHive() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interactiveHive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) InteractiveHiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interactiveHiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) InternalValue() *HdinsightInteractiveQueryClusterComponentVersion {
	var returns *HdinsightInteractiveQueryClusterComponentVersion
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterComponentVersionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterComponentVersionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterComponentVersionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterComponentVersionOutputReference_Override(h HdinsightInteractiveQueryClusterComponentVersionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterComponentVersionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) SetInteractiveHive(val *string) {
	_jsii_.Set(
		j,
		"interactiveHive",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterComponentVersion) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterComponentVersionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#cluster_version HdinsightInteractiveQueryCluster#cluster_version}.
	ClusterVersion *string `field:"required" json:"clusterVersion" yaml:"clusterVersion"`
	// component_version block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#component_version HdinsightInteractiveQueryCluster#component_version}
	ComponentVersion *HdinsightInteractiveQueryClusterComponentVersion `field:"required" json:"componentVersion" yaml:"componentVersion"`
	// gateway block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#gateway HdinsightInteractiveQueryCluster#gateway}
	Gateway *HdinsightInteractiveQueryClusterGateway `field:"required" json:"gateway" yaml:"gateway"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#location HdinsightInteractiveQueryCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#name HdinsightInteractiveQueryCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#resource_group_name HdinsightInteractiveQueryCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// roles block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#roles HdinsightInteractiveQueryCluster#roles}
	Roles *HdinsightInteractiveQueryClusterRoles `field:"required" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#tier HdinsightInteractiveQueryCluster#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// disk_encryption block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#disk_encryption HdinsightInteractiveQueryCluster#disk_encryption}
	DiskEncryption interface{} `field:"optional" json:"diskEncryption" yaml:"diskEncryption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#encryption_in_transit_enabled HdinsightInteractiveQueryCluster#encryption_in_transit_enabled}.
	EncryptionInTransitEnabled interface{} `field:"optional" json:"encryptionInTransitEnabled" yaml:"encryptionInTransitEnabled"`
	// extension block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#extension HdinsightInteractiveQueryCluster#extension}
	Extension *HdinsightInteractiveQueryClusterExtension `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#id HdinsightInteractiveQueryCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// metastores block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#metastores HdinsightInteractiveQueryCluster#metastores}
	Metastores *HdinsightInteractiveQueryClusterMetastores `field:"optional" json:"metastores" yaml:"metastores"`
	// monitor block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#monitor HdinsightInteractiveQueryCluster#monitor}
	Monitor *HdinsightInteractiveQueryClusterMonitor `field:"optional" json:"monitor" yaml:"monitor"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#network HdinsightInteractiveQueryCluster#network}
	Network *HdinsightInteractiveQueryClusterNetwork `field:"optional" json:"network" yaml:"network"`
	// security_profile block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#security_profile HdinsightInteractiveQueryCluster#security_profile}
	SecurityProfile *HdinsightInteractiveQueryClusterSecurityProfile `field:"optional" json:"securityProfile" yaml:"securityProfile"`
	// storage_account block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#storage_account HdinsightInteractiveQueryCluster#storage_account}
	StorageAccount interface{} `field:"optional" json:"storageAccount" yaml:"storageAccount"`
	// storage_account_gen2 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#storage_account_gen2 HdinsightInteractiveQueryCluster#storage_account_gen2}
	StorageAccountGen2 *HdinsightInteractiveQueryClusterStorageAccountGen2 `field:"optional" json:"storageAccountGen2" yaml:"storageAccountGen2"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#tags HdinsightInteractiveQueryCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#timeouts HdinsightInteractiveQueryCluster#timeouts}
	Timeouts *HdinsightInteractiveQueryClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#tls_min_version HdinsightInteractiveQueryCluster#tls_min_version}.
	TlsMinVersion *string `field:"optional" json:"tlsMinVersion" yaml:"tlsMinVersion"`
}

type HdinsightInteractiveQueryClusterDiskEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#encryption_algorithm HdinsightInteractiveQueryCluster#encryption_algorithm}.
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#encryption_at_host_enabled HdinsightInteractiveQueryCluster#encryption_at_host_enabled}.
	EncryptionAtHostEnabled interface{} `field:"optional" json:"encryptionAtHostEnabled" yaml:"encryptionAtHostEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#key_vault_key_id HdinsightInteractiveQueryCluster#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#key_vault_managed_identity_id HdinsightInteractiveQueryCluster#key_vault_managed_identity_id}.
	KeyVaultManagedIdentityId *string `field:"optional" json:"keyVaultManagedIdentityId" yaml:"keyVaultManagedIdentityId"`
}

type HdinsightInteractiveQueryClusterDiskEncryptionList interface {
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
	Get(index *float64) HdinsightInteractiveQueryClusterDiskEncryptionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightInteractiveQueryClusterDiskEncryptionList
type jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterDiskEncryptionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HdinsightInteractiveQueryClusterDiskEncryptionList {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterDiskEncryptionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterDiskEncryptionList_Override(h HdinsightInteractiveQueryClusterDiskEncryptionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterDiskEncryptionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) Get(index *float64) HdinsightInteractiveQueryClusterDiskEncryptionOutputReference {
	var returns HdinsightInteractiveQueryClusterDiskEncryptionOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterDiskEncryptionOutputReference interface {
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterDiskEncryptionOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) EncryptionAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) EncryptionAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) EncryptionAtHostEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) EncryptionAtHostEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) KeyVaultKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) KeyVaultKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) KeyVaultManagedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultManagedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) KeyVaultManagedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultManagedIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterDiskEncryptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HdinsightInteractiveQueryClusterDiskEncryptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterDiskEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterDiskEncryptionOutputReference_Override(h HdinsightInteractiveQueryClusterDiskEncryptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterDiskEncryptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) SetEncryptionAlgorithm(val *string) {
	_jsii_.Set(
		j,
		"encryptionAlgorithm",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) SetEncryptionAtHostEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"encryptionAtHostEnabled",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) SetKeyVaultKeyId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultKeyId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) SetKeyVaultManagedIdentityId(val *string) {
	_jsii_.Set(
		j,
		"keyVaultManagedIdentityId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) ResetEncryptionAlgorithm() {
	_jsii_.InvokeVoid(
		h,
		"resetEncryptionAlgorithm",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) ResetEncryptionAtHostEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetEncryptionAtHostEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) ResetKeyVaultKeyId() {
	_jsii_.InvokeVoid(
		h,
		"resetKeyVaultKeyId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) ResetKeyVaultManagedIdentityId() {
	_jsii_.InvokeVoid(
		h,
		"resetKeyVaultManagedIdentityId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterDiskEncryptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterExtension struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#log_analytics_workspace_id HdinsightInteractiveQueryCluster#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#primary_key HdinsightInteractiveQueryCluster#primary_key}.
	PrimaryKey *string `field:"required" json:"primaryKey" yaml:"primaryKey"`
}

type HdinsightInteractiveQueryClusterExtensionOutputReference interface {
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
	InternalValue() *HdinsightInteractiveQueryClusterExtension
	SetInternalValue(val *HdinsightInteractiveQueryClusterExtension)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterExtensionOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) InternalValue() *HdinsightInteractiveQueryClusterExtension {
	var returns *HdinsightInteractiveQueryClusterExtension
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) PrimaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) PrimaryKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterExtensionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterExtensionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterExtensionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterExtensionOutputReference_Override(h HdinsightInteractiveQueryClusterExtensionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterExtensionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterExtension) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) SetLogAnalyticsWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) SetPrimaryKey(val *string) {
	_jsii_.Set(
		j,
		"primaryKey",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterExtensionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterGateway struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#password HdinsightInteractiveQueryCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#username HdinsightInteractiveQueryCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

type HdinsightInteractiveQueryClusterGatewayOutputReference interface {
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
	InternalValue() *HdinsightInteractiveQueryClusterGateway
	SetInternalValue(val *HdinsightInteractiveQueryClusterGateway)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterGatewayOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) InternalValue() *HdinsightInteractiveQueryClusterGateway {
	var returns *HdinsightInteractiveQueryClusterGateway
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterGatewayOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterGatewayOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterGatewayOutputReference_Override(h HdinsightInteractiveQueryClusterGatewayOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterGateway) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterGatewayOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterMetastores struct {
	// ambari block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#ambari HdinsightInteractiveQueryCluster#ambari}
	Ambari *HdinsightInteractiveQueryClusterMetastoresAmbari `field:"optional" json:"ambari" yaml:"ambari"`
	// hive block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#hive HdinsightInteractiveQueryCluster#hive}
	Hive *HdinsightInteractiveQueryClusterMetastoresHive `field:"optional" json:"hive" yaml:"hive"`
	// oozie block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#oozie HdinsightInteractiveQueryCluster#oozie}
	Oozie *HdinsightInteractiveQueryClusterMetastoresOozie `field:"optional" json:"oozie" yaml:"oozie"`
}

type HdinsightInteractiveQueryClusterMetastoresAmbari struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#database_name HdinsightInteractiveQueryCluster#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#password HdinsightInteractiveQueryCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#server HdinsightInteractiveQueryCluster#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#username HdinsightInteractiveQueryCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

type HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference interface {
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
	InternalValue() *HdinsightInteractiveQueryClusterMetastoresAmbari
	SetInternalValue(val *HdinsightInteractiveQueryClusterMetastoresAmbari)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) InternalValue() *HdinsightInteractiveQueryClusterMetastoresAmbari {
	var returns *HdinsightInteractiveQueryClusterMetastoresAmbari
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) Server() *string {
	var returns *string
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) ServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterMetastoresAmbariOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterMetastoresAmbariOutputReference_Override(h HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterMetastoresAmbari) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) SetServer(val *string) {
	_jsii_.Set(
		j,
		"server",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterMetastoresHive struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#database_name HdinsightInteractiveQueryCluster#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#password HdinsightInteractiveQueryCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#server HdinsightInteractiveQueryCluster#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#username HdinsightInteractiveQueryCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

type HdinsightInteractiveQueryClusterMetastoresHiveOutputReference interface {
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
	InternalValue() *HdinsightInteractiveQueryClusterMetastoresHive
	SetInternalValue(val *HdinsightInteractiveQueryClusterMetastoresHive)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterMetastoresHiveOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) InternalValue() *HdinsightInteractiveQueryClusterMetastoresHive {
	var returns *HdinsightInteractiveQueryClusterMetastoresHive
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) Server() *string {
	var returns *string
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) ServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterMetastoresHiveOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterMetastoresHiveOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterMetastoresHiveOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterMetastoresHiveOutputReference_Override(h HdinsightInteractiveQueryClusterMetastoresHiveOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterMetastoresHiveOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterMetastoresHive) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) SetServer(val *string) {
	_jsii_.Set(
		j,
		"server",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresHiveOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterMetastoresOozie struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#database_name HdinsightInteractiveQueryCluster#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#password HdinsightInteractiveQueryCluster#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#server HdinsightInteractiveQueryCluster#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#username HdinsightInteractiveQueryCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

type HdinsightInteractiveQueryClusterMetastoresOozieOutputReference interface {
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
	InternalValue() *HdinsightInteractiveQueryClusterMetastoresOozie
	SetInternalValue(val *HdinsightInteractiveQueryClusterMetastoresOozie)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterMetastoresOozieOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) InternalValue() *HdinsightInteractiveQueryClusterMetastoresOozie {
	var returns *HdinsightInteractiveQueryClusterMetastoresOozie
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) Server() *string {
	var returns *string
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) ServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterMetastoresOozieOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterMetastoresOozieOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterMetastoresOozieOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterMetastoresOozieOutputReference_Override(h HdinsightInteractiveQueryClusterMetastoresOozieOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterMetastoresOozieOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterMetastoresOozie) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) SetServer(val *string) {
	_jsii_.Set(
		j,
		"server",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOozieOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterMetastoresOutputReference interface {
	cdktf.ComplexObject
	Ambari() HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference
	AmbariInput() *HdinsightInteractiveQueryClusterMetastoresAmbari
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
	Hive() HdinsightInteractiveQueryClusterMetastoresHiveOutputReference
	HiveInput() *HdinsightInteractiveQueryClusterMetastoresHive
	InternalValue() *HdinsightInteractiveQueryClusterMetastores
	SetInternalValue(val *HdinsightInteractiveQueryClusterMetastores)
	Oozie() HdinsightInteractiveQueryClusterMetastoresOozieOutputReference
	OozieInput() *HdinsightInteractiveQueryClusterMetastoresOozie
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
	PutAmbari(value *HdinsightInteractiveQueryClusterMetastoresAmbari)
	PutHive(value *HdinsightInteractiveQueryClusterMetastoresHive)
	PutOozie(value *HdinsightInteractiveQueryClusterMetastoresOozie)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterMetastoresOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) Ambari() HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference {
	var returns HdinsightInteractiveQueryClusterMetastoresAmbariOutputReference
	_jsii_.Get(
		j,
		"ambari",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) AmbariInput() *HdinsightInteractiveQueryClusterMetastoresAmbari {
	var returns *HdinsightInteractiveQueryClusterMetastoresAmbari
	_jsii_.Get(
		j,
		"ambariInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) Hive() HdinsightInteractiveQueryClusterMetastoresHiveOutputReference {
	var returns HdinsightInteractiveQueryClusterMetastoresHiveOutputReference
	_jsii_.Get(
		j,
		"hive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) HiveInput() *HdinsightInteractiveQueryClusterMetastoresHive {
	var returns *HdinsightInteractiveQueryClusterMetastoresHive
	_jsii_.Get(
		j,
		"hiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) InternalValue() *HdinsightInteractiveQueryClusterMetastores {
	var returns *HdinsightInteractiveQueryClusterMetastores
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) Oozie() HdinsightInteractiveQueryClusterMetastoresOozieOutputReference {
	var returns HdinsightInteractiveQueryClusterMetastoresOozieOutputReference
	_jsii_.Get(
		j,
		"oozie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) OozieInput() *HdinsightInteractiveQueryClusterMetastoresOozie {
	var returns *HdinsightInteractiveQueryClusterMetastoresOozie
	_jsii_.Get(
		j,
		"oozieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterMetastoresOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterMetastoresOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterMetastoresOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterMetastoresOutputReference_Override(h HdinsightInteractiveQueryClusterMetastoresOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterMetastoresOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterMetastores) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) PutAmbari(value *HdinsightInteractiveQueryClusterMetastoresAmbari) {
	_jsii_.InvokeVoid(
		h,
		"putAmbari",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) PutHive(value *HdinsightInteractiveQueryClusterMetastoresHive) {
	_jsii_.InvokeVoid(
		h,
		"putHive",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) PutOozie(value *HdinsightInteractiveQueryClusterMetastoresOozie) {
	_jsii_.InvokeVoid(
		h,
		"putOozie",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) ResetAmbari() {
	_jsii_.InvokeVoid(
		h,
		"resetAmbari",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) ResetHive() {
	_jsii_.InvokeVoid(
		h,
		"resetHive",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) ResetOozie() {
	_jsii_.InvokeVoid(
		h,
		"resetOozie",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMetastoresOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterMonitor struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#log_analytics_workspace_id HdinsightInteractiveQueryCluster#log_analytics_workspace_id}.
	LogAnalyticsWorkspaceId *string `field:"required" json:"logAnalyticsWorkspaceId" yaml:"logAnalyticsWorkspaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#primary_key HdinsightInteractiveQueryCluster#primary_key}.
	PrimaryKey *string `field:"required" json:"primaryKey" yaml:"primaryKey"`
}

type HdinsightInteractiveQueryClusterMonitorOutputReference interface {
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
	InternalValue() *HdinsightInteractiveQueryClusterMonitor
	SetInternalValue(val *HdinsightInteractiveQueryClusterMonitor)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterMonitorOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) InternalValue() *HdinsightInteractiveQueryClusterMonitor {
	var returns *HdinsightInteractiveQueryClusterMonitor
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) LogAnalyticsWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) LogAnalyticsWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logAnalyticsWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) PrimaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) PrimaryKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterMonitorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterMonitorOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterMonitorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterMonitorOutputReference_Override(h HdinsightInteractiveQueryClusterMonitorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterMonitorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterMonitor) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) SetLogAnalyticsWorkspaceId(val *string) {
	_jsii_.Set(
		j,
		"logAnalyticsWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) SetPrimaryKey(val *string) {
	_jsii_.Set(
		j,
		"primaryKey",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterMonitorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterNetwork struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#connection_direction HdinsightInteractiveQueryCluster#connection_direction}.
	ConnectionDirection *string `field:"optional" json:"connectionDirection" yaml:"connectionDirection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#private_link_enabled HdinsightInteractiveQueryCluster#private_link_enabled}.
	PrivateLinkEnabled interface{} `field:"optional" json:"privateLinkEnabled" yaml:"privateLinkEnabled"`
}

type HdinsightInteractiveQueryClusterNetworkOutputReference interface {
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
	InternalValue() *HdinsightInteractiveQueryClusterNetwork
	SetInternalValue(val *HdinsightInteractiveQueryClusterNetwork)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterNetworkOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) ConnectionDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) ConnectionDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) InternalValue() *HdinsightInteractiveQueryClusterNetwork {
	var returns *HdinsightInteractiveQueryClusterNetwork
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) PrivateLinkEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateLinkEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) PrivateLinkEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateLinkEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterNetworkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterNetworkOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterNetworkOutputReference_Override(h HdinsightInteractiveQueryClusterNetworkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) SetConnectionDirection(val *string) {
	_jsii_.Set(
		j,
		"connectionDirection",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterNetwork) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) SetPrivateLinkEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"privateLinkEnabled",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) ResetConnectionDirection() {
	_jsii_.InvokeVoid(
		h,
		"resetConnectionDirection",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) ResetPrivateLinkEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetPrivateLinkEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterRoles struct {
	// head_node block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#head_node HdinsightInteractiveQueryCluster#head_node}
	HeadNode *HdinsightInteractiveQueryClusterRolesHeadNode `field:"required" json:"headNode" yaml:"headNode"`
	// worker_node block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#worker_node HdinsightInteractiveQueryCluster#worker_node}
	WorkerNode *HdinsightInteractiveQueryClusterRolesWorkerNode `field:"required" json:"workerNode" yaml:"workerNode"`
	// zookeeper_node block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#zookeeper_node HdinsightInteractiveQueryCluster#zookeeper_node}
	ZookeeperNode *HdinsightInteractiveQueryClusterRolesZookeeperNode `field:"required" json:"zookeeperNode" yaml:"zookeeperNode"`
}

type HdinsightInteractiveQueryClusterRolesHeadNode struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#username HdinsightInteractiveQueryCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#vm_size HdinsightInteractiveQueryCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#password HdinsightInteractiveQueryCluster#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#ssh_keys HdinsightInteractiveQueryCluster#ssh_keys}.
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#subnet_id HdinsightInteractiveQueryCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#virtual_network_id HdinsightInteractiveQueryCluster#virtual_network_id}.
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

type HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference interface {
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
	InternalValue() *HdinsightInteractiveQueryClusterRolesHeadNode
	SetInternalValue(val *HdinsightInteractiveQueryClusterRolesHeadNode)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) InternalValue() *HdinsightInteractiveQueryClusterRolesHeadNode {
	var returns *HdinsightInteractiveQueryClusterRolesHeadNode
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SshKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterRolesHeadNodeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterRolesHeadNodeOutputReference_Override(h HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterRolesHeadNode) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SetSshKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"sshKeys",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SetVirtualNetworkId(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) SetVmSize(val *string) {
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		h,
		"resetPassword",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) ResetSshKeys() {
	_jsii_.InvokeVoid(
		h,
		"resetSshKeys",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		h,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		h,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterRolesOutputReference interface {
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
	HeadNode() HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference
	HeadNodeInput() *HdinsightInteractiveQueryClusterRolesHeadNode
	InternalValue() *HdinsightInteractiveQueryClusterRoles
	SetInternalValue(val *HdinsightInteractiveQueryClusterRoles)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkerNode() HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference
	WorkerNodeInput() *HdinsightInteractiveQueryClusterRolesWorkerNode
	ZookeeperNode() HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference
	ZookeeperNodeInput() *HdinsightInteractiveQueryClusterRolesZookeeperNode
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
	PutHeadNode(value *HdinsightInteractiveQueryClusterRolesHeadNode)
	PutWorkerNode(value *HdinsightInteractiveQueryClusterRolesWorkerNode)
	PutZookeeperNode(value *HdinsightInteractiveQueryClusterRolesZookeeperNode)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightInteractiveQueryClusterRolesOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) HeadNode() HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference {
	var returns HdinsightInteractiveQueryClusterRolesHeadNodeOutputReference
	_jsii_.Get(
		j,
		"headNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) HeadNodeInput() *HdinsightInteractiveQueryClusterRolesHeadNode {
	var returns *HdinsightInteractiveQueryClusterRolesHeadNode
	_jsii_.Get(
		j,
		"headNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) InternalValue() *HdinsightInteractiveQueryClusterRoles {
	var returns *HdinsightInteractiveQueryClusterRoles
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) WorkerNode() HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference {
	var returns HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference
	_jsii_.Get(
		j,
		"workerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) WorkerNodeInput() *HdinsightInteractiveQueryClusterRolesWorkerNode {
	var returns *HdinsightInteractiveQueryClusterRolesWorkerNode
	_jsii_.Get(
		j,
		"workerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) ZookeeperNode() HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference {
	var returns HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference
	_jsii_.Get(
		j,
		"zookeeperNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) ZookeeperNodeInput() *HdinsightInteractiveQueryClusterRolesZookeeperNode {
	var returns *HdinsightInteractiveQueryClusterRolesZookeeperNode
	_jsii_.Get(
		j,
		"zookeeperNodeInput",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterRolesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterRolesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterRolesOutputReference_Override(h HdinsightInteractiveQueryClusterRolesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterRoles) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) PutHeadNode(value *HdinsightInteractiveQueryClusterRolesHeadNode) {
	_jsii_.InvokeVoid(
		h,
		"putHeadNode",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) PutWorkerNode(value *HdinsightInteractiveQueryClusterRolesWorkerNode) {
	_jsii_.InvokeVoid(
		h,
		"putWorkerNode",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) PutZookeeperNode(value *HdinsightInteractiveQueryClusterRolesZookeeperNode) {
	_jsii_.InvokeVoid(
		h,
		"putZookeeperNode",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterRolesWorkerNode struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#target_instance_count HdinsightInteractiveQueryCluster#target_instance_count}.
	TargetInstanceCount *float64 `field:"required" json:"targetInstanceCount" yaml:"targetInstanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#username HdinsightInteractiveQueryCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#vm_size HdinsightInteractiveQueryCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// autoscale block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#autoscale HdinsightInteractiveQueryCluster#autoscale}
	Autoscale *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale `field:"optional" json:"autoscale" yaml:"autoscale"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#password HdinsightInteractiveQueryCluster#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#ssh_keys HdinsightInteractiveQueryCluster#ssh_keys}.
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#subnet_id HdinsightInteractiveQueryCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#virtual_network_id HdinsightInteractiveQueryCluster#virtual_network_id}.
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

type HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale struct {
	// capacity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#capacity HdinsightInteractiveQueryCluster#capacity}
	Capacity *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity `field:"optional" json:"capacity" yaml:"capacity"`
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#recurrence HdinsightInteractiveQueryCluster#recurrence}
	Recurrence *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence `field:"optional" json:"recurrence" yaml:"recurrence"`
}

type HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#max_instance_count HdinsightInteractiveQueryCluster#max_instance_count}.
	MaxInstanceCount *float64 `field:"required" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#min_instance_count HdinsightInteractiveQueryCluster#min_instance_count}.
	MinInstanceCount *float64 `field:"required" json:"minInstanceCount" yaml:"minInstanceCount"`
}

type HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference interface {
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
	InternalValue() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity
	SetInternalValue(val *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) InternalValue() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity {
	var returns *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) MaxInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) MaxInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) MinInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) MinInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference_Override(h HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetMaxInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"maxInstanceCount",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetMinInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"minInstanceCount",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference interface {
	cdktf.ComplexObject
	Capacity() HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference
	CapacityInput() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity
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
	InternalValue() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale
	SetInternalValue(val *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale)
	Recurrence() HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference
	RecurrenceInput() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence
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
	PutCapacity(value *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity)
	PutRecurrence(value *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) Capacity() HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference {
	var returns HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacityOutputReference
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) CapacityInput() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity {
	var returns *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) InternalValue() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale {
	var returns *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) Recurrence() HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference {
	var returns HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference
	_jsii_.Get(
		j,
		"recurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) RecurrenceInput() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence {
	var returns *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence
	_jsii_.Get(
		j,
		"recurrenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference_Override(h HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) PutCapacity(value *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleCapacity) {
	_jsii_.InvokeVoid(
		h,
		"putCapacity",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) PutRecurrence(value *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence) {
	_jsii_.InvokeVoid(
		h,
		"putRecurrence",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) ResetCapacity() {
	_jsii_.InvokeVoid(
		h,
		"resetCapacity",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) ResetRecurrence() {
	_jsii_.InvokeVoid(
		h,
		"resetRecurrence",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence struct {
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#schedule HdinsightInteractiveQueryCluster#schedule}
	Schedule interface{} `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#timezone HdinsightInteractiveQueryCluster#timezone}.
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

type HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference interface {
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
	InternalValue() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence
	SetInternalValue(val *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence)
	Schedule() HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) InternalValue() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence {
	var returns *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) Schedule() HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList {
	var returns HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) ScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference_Override(h HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrence) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) SetTimezone(val *string) {
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) PutSchedule(value interface{}) {
	_jsii_.InvokeVoid(
		h,
		"putSchedule",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#days HdinsightInteractiveQueryCluster#days}.
	Days *[]*string `field:"required" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#target_instance_count HdinsightInteractiveQueryCluster#target_instance_count}.
	TargetInstanceCount *float64 `field:"required" json:"targetInstanceCount" yaml:"targetInstanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#time HdinsightInteractiveQueryCluster#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
}

type HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList interface {
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
	Get(index *float64) HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList
type jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList_Override(h HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) Get(index *float64) HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference {
	var returns HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference interface {
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) Days() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"days",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) DaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) TargetInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) TargetInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) Time() *string {
	var returns *string
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) TimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference_Override(h HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetDays(val *[]*string) {
	_jsii_.Set(
		j,
		"days",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetTargetInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"targetInstanceCount",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) SetTime(val *string) {
	_jsii_.Set(
		j,
		"time",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleRecurrenceScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference interface {
	cdktf.ComplexObject
	Autoscale() HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference
	AutoscaleInput() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale
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
	InternalValue() *HdinsightInteractiveQueryClusterRolesWorkerNode
	SetInternalValue(val *HdinsightInteractiveQueryClusterRolesWorkerNode)
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
	PutAutoscale(value *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) Autoscale() HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference {
	var returns HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscaleOutputReference
	_jsii_.Get(
		j,
		"autoscale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) AutoscaleInput() *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale {
	var returns *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale
	_jsii_.Get(
		j,
		"autoscaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) InternalValue() *HdinsightInteractiveQueryClusterRolesWorkerNode {
	var returns *HdinsightInteractiveQueryClusterRolesWorkerNode
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SshKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) TargetInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) TargetInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference_Override(h HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterRolesWorkerNode) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SetSshKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"sshKeys",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SetTargetInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"targetInstanceCount",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SetVirtualNetworkId(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) SetVmSize(val *string) {
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) PutAutoscale(value *HdinsightInteractiveQueryClusterRolesWorkerNodeAutoscale) {
	_jsii_.InvokeVoid(
		h,
		"putAutoscale",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) ResetAutoscale() {
	_jsii_.InvokeVoid(
		h,
		"resetAutoscale",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		h,
		"resetPassword",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) ResetSshKeys() {
	_jsii_.InvokeVoid(
		h,
		"resetSshKeys",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		h,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		h,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesWorkerNodeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterRolesZookeeperNode struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#username HdinsightInteractiveQueryCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#vm_size HdinsightInteractiveQueryCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#password HdinsightInteractiveQueryCluster#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#ssh_keys HdinsightInteractiveQueryCluster#ssh_keys}.
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#subnet_id HdinsightInteractiveQueryCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#virtual_network_id HdinsightInteractiveQueryCluster#virtual_network_id}.
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}

type HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference interface {
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
	InternalValue() *HdinsightInteractiveQueryClusterRolesZookeeperNode
	SetInternalValue(val *HdinsightInteractiveQueryClusterRolesZookeeperNode)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) InternalValue() *HdinsightInteractiveQueryClusterRolesZookeeperNode {
	var returns *HdinsightInteractiveQueryClusterRolesZookeeperNode
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SshKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SshKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference_Override(h HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterRolesZookeeperNode) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SetSshKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"sshKeys",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SetSubnetId(val *string) {
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SetVirtualNetworkId(val *string) {
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) SetVmSize(val *string) {
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		h,
		"resetPassword",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) ResetSshKeys() {
	_jsii_.InvokeVoid(
		h,
		"resetSshKeys",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		h,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		h,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterRolesZookeeperNodeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterSecurityProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#aadds_resource_id HdinsightInteractiveQueryCluster#aadds_resource_id}.
	AaddsResourceId *string `field:"required" json:"aaddsResourceId" yaml:"aaddsResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#domain_name HdinsightInteractiveQueryCluster#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#domain_username HdinsightInteractiveQueryCluster#domain_username}.
	DomainUsername *string `field:"required" json:"domainUsername" yaml:"domainUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#domain_user_password HdinsightInteractiveQueryCluster#domain_user_password}.
	DomainUserPassword *string `field:"required" json:"domainUserPassword" yaml:"domainUserPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#ldaps_urls HdinsightInteractiveQueryCluster#ldaps_urls}.
	LdapsUrls *[]*string `field:"required" json:"ldapsUrls" yaml:"ldapsUrls"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#msi_resource_id HdinsightInteractiveQueryCluster#msi_resource_id}.
	MsiResourceId *string `field:"required" json:"msiResourceId" yaml:"msiResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#cluster_users_group_dns HdinsightInteractiveQueryCluster#cluster_users_group_dns}.
	ClusterUsersGroupDns *[]*string `field:"optional" json:"clusterUsersGroupDns" yaml:"clusterUsersGroupDns"`
}

type HdinsightInteractiveQueryClusterSecurityProfileOutputReference interface {
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
	InternalValue() *HdinsightInteractiveQueryClusterSecurityProfile
	SetInternalValue(val *HdinsightInteractiveQueryClusterSecurityProfile)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterSecurityProfileOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) AaddsResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aaddsResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) AaddsResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aaddsResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ClusterUsersGroupDns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterUsersGroupDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ClusterUsersGroupDnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterUsersGroupDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) DomainUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) DomainUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) DomainUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) DomainUserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) InternalValue() *HdinsightInteractiveQueryClusterSecurityProfile {
	var returns *HdinsightInteractiveQueryClusterSecurityProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) LdapsUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ldapsUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) LdapsUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ldapsUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) MsiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) MsiResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterSecurityProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterSecurityProfileOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterSecurityProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterSecurityProfileOutputReference_Override(h HdinsightInteractiveQueryClusterSecurityProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterSecurityProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) SetAaddsResourceId(val *string) {
	_jsii_.Set(
		j,
		"aaddsResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) SetClusterUsersGroupDns(val *[]*string) {
	_jsii_.Set(
		j,
		"clusterUsersGroupDns",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) SetDomainName(val *string) {
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) SetDomainUsername(val *string) {
	_jsii_.Set(
		j,
		"domainUsername",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) SetDomainUserPassword(val *string) {
	_jsii_.Set(
		j,
		"domainUserPassword",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterSecurityProfile) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) SetLdapsUrls(val *[]*string) {
	_jsii_.Set(
		j,
		"ldapsUrls",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) SetMsiResourceId(val *string) {
	_jsii_.Set(
		j,
		"msiResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ResetClusterUsersGroupDns() {
	_jsii_.InvokeVoid(
		h,
		"resetClusterUsersGroupDns",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterSecurityProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterStorageAccount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#is_default HdinsightInteractiveQueryCluster#is_default}.
	IsDefault interface{} `field:"required" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#storage_account_key HdinsightInteractiveQueryCluster#storage_account_key}.
	StorageAccountKey *string `field:"required" json:"storageAccountKey" yaml:"storageAccountKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#storage_container_id HdinsightInteractiveQueryCluster#storage_container_id}.
	StorageContainerId *string `field:"required" json:"storageContainerId" yaml:"storageContainerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#storage_resource_id HdinsightInteractiveQueryCluster#storage_resource_id}.
	StorageResourceId *string `field:"optional" json:"storageResourceId" yaml:"storageResourceId"`
}

type HdinsightInteractiveQueryClusterStorageAccountGen2 struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#filesystem_id HdinsightInteractiveQueryCluster#filesystem_id}.
	FilesystemId *string `field:"required" json:"filesystemId" yaml:"filesystemId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#is_default HdinsightInteractiveQueryCluster#is_default}.
	IsDefault interface{} `field:"required" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#managed_identity_resource_id HdinsightInteractiveQueryCluster#managed_identity_resource_id}.
	ManagedIdentityResourceId *string `field:"required" json:"managedIdentityResourceId" yaml:"managedIdentityResourceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#storage_resource_id HdinsightInteractiveQueryCluster#storage_resource_id}.
	StorageResourceId *string `field:"required" json:"storageResourceId" yaml:"storageResourceId"`
}

type HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference interface {
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
	InternalValue() *HdinsightInteractiveQueryClusterStorageAccountGen2
	SetInternalValue(val *HdinsightInteractiveQueryClusterStorageAccountGen2)
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) FilesystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filesystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) FilesystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filesystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) InternalValue() *HdinsightInteractiveQueryClusterStorageAccountGen2 {
	var returns *HdinsightInteractiveQueryClusterStorageAccountGen2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) IsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) IsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) ManagedIdentityResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) ManagedIdentityResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) StorageResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) StorageResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterStorageAccountGen2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterStorageAccountGen2OutputReference_Override(h HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) SetFilesystemId(val *string) {
	_jsii_.Set(
		j,
		"filesystemId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) SetInternalValue(val *HdinsightInteractiveQueryClusterStorageAccountGen2) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) SetIsDefault(val interface{}) {
	_jsii_.Set(
		j,
		"isDefault",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) SetManagedIdentityResourceId(val *string) {
	_jsii_.Set(
		j,
		"managedIdentityResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) SetStorageResourceId(val *string) {
	_jsii_.Set(
		j,
		"storageResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountGen2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterStorageAccountList interface {
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
	Get(index *float64) HdinsightInteractiveQueryClusterStorageAccountOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightInteractiveQueryClusterStorageAccountList
type jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterStorageAccountList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HdinsightInteractiveQueryClusterStorageAccountList {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterStorageAccountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterStorageAccountList_Override(h HdinsightInteractiveQueryClusterStorageAccountList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterStorageAccountList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) Get(index *float64) HdinsightInteractiveQueryClusterStorageAccountOutputReference {
	var returns HdinsightInteractiveQueryClusterStorageAccountOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterStorageAccountOutputReference interface {
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterStorageAccountOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) IsDefault() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) IsDefaultInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) StorageAccountKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) StorageAccountKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) StorageContainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) StorageContainerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) StorageResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) StorageResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterStorageAccountOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HdinsightInteractiveQueryClusterStorageAccountOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterStorageAccountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterStorageAccountOutputReference_Override(h HdinsightInteractiveQueryClusterStorageAccountOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterStorageAccountOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) SetIsDefault(val interface{}) {
	_jsii_.Set(
		j,
		"isDefault",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) SetStorageAccountKey(val *string) {
	_jsii_.Set(
		j,
		"storageAccountKey",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) SetStorageContainerId(val *string) {
	_jsii_.Set(
		j,
		"storageContainerId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) SetStorageResourceId(val *string) {
	_jsii_.Set(
		j,
		"storageResourceId",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) ResetStorageResourceId() {
	_jsii_.InvokeVoid(
		h,
		"resetStorageResourceId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterStorageAccountOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type HdinsightInteractiveQueryClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#create HdinsightInteractiveQueryCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#delete HdinsightInteractiveQueryCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#read HdinsightInteractiveQueryCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_interactive_query_cluster#update HdinsightInteractiveQueryCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type HdinsightInteractiveQueryClusterTimeoutsOutputReference interface {
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

// The jsii proxy struct for HdinsightInteractiveQueryClusterTimeoutsOutputReference
type jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewHdinsightInteractiveQueryClusterTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightInteractiveQueryClusterTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightInteractiveQueryClusterTimeoutsOutputReference_Override(h HdinsightInteractiveQueryClusterTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightInteractiveQueryCluster.HdinsightInteractiveQueryClusterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		h,
		"resetCreate",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		h,
		"resetDelete",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		h,
		"resetRead",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		h,
		"resetUpdate",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightInteractiveQueryClusterTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

