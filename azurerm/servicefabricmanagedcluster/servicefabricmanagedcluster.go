package servicefabricmanagedcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/servicefabricmanagedcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster azurerm_service_fabric_managed_cluster}.
type ServiceFabricManagedCluster interface {
	cdktf.TerraformResource
	Authentication() ServiceFabricManagedClusterAuthenticationOutputReference
	AuthenticationInput() *ServiceFabricManagedClusterAuthentication
	BackupServiceEnabled() interface{}
	SetBackupServiceEnabled(val interface{})
	BackupServiceEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientConnectionPort() *float64
	SetClientConnectionPort(val *float64)
	ClientConnectionPortInput() *float64
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
	CustomFabricSetting() ServiceFabricManagedClusterCustomFabricSettingList
	CustomFabricSettingInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DnsName() *string
	SetDnsName(val *string)
	DnsNameInput() *string
	DnsServiceEnabled() interface{}
	SetDnsServiceEnabled(val interface{})
	DnsServiceEnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpGatewayPort() *float64
	SetHttpGatewayPort(val *float64)
	HttpGatewayPortInput() *float64
	Id() *string
	SetId(val *string)
	IdInput() *string
	LbRule() ServiceFabricManagedClusterLbRuleList
	LbRuleInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeType() ServiceFabricManagedClusterNodeTypeList
	NodeTypeInput() interface{}
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
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
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ServiceFabricManagedClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpgradeWave() *string
	SetUpgradeWave(val *string)
	UpgradeWaveInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	PutAuthentication(value *ServiceFabricManagedClusterAuthentication)
	PutCustomFabricSetting(value interface{})
	PutLbRule(value interface{})
	PutNodeType(value interface{})
	PutTimeouts(value *ServiceFabricManagedClusterTimeouts)
	ResetAuthentication()
	ResetBackupServiceEnabled()
	ResetCustomFabricSetting()
	ResetDnsName()
	ResetDnsServiceEnabled()
	ResetId()
	ResetNodeType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetSku()
	ResetTags()
	ResetTimeouts()
	ResetUpgradeWave()
	ResetUsername()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ServiceFabricManagedCluster
type jsiiProxy_ServiceFabricManagedCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Authentication() ServiceFabricManagedClusterAuthenticationOutputReference {
	var returns ServiceFabricManagedClusterAuthenticationOutputReference
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) AuthenticationInput() *ServiceFabricManagedClusterAuthentication {
	var returns *ServiceFabricManagedClusterAuthentication
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) BackupServiceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupServiceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) BackupServiceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backupServiceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) ClientConnectionPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientConnectionPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) ClientConnectionPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientConnectionPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) CustomFabricSetting() ServiceFabricManagedClusterCustomFabricSettingList {
	var returns ServiceFabricManagedClusterCustomFabricSettingList
	_jsii_.Get(
		j,
		"customFabricSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) CustomFabricSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFabricSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) DnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) DnsServiceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsServiceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) DnsServiceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsServiceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) HttpGatewayPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpGatewayPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) HttpGatewayPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpGatewayPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) LbRule() ServiceFabricManagedClusterLbRuleList {
	var returns ServiceFabricManagedClusterLbRuleList
	_jsii_.Get(
		j,
		"lbRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) LbRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lbRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) NodeType() ServiceFabricManagedClusterNodeTypeList {
	var returns ServiceFabricManagedClusterNodeTypeList
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) NodeTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Timeouts() ServiceFabricManagedClusterTimeoutsOutputReference {
	var returns ServiceFabricManagedClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) UpgradeWave() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeWave",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) UpgradeWaveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeWaveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedCluster) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster azurerm_service_fabric_managed_cluster} Resource.
func NewServiceFabricManagedCluster(scope constructs.Construct, id *string, config *ServiceFabricManagedClusterConfig) ServiceFabricManagedCluster {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedCluster{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster azurerm_service_fabric_managed_cluster} Resource.
func NewServiceFabricManagedCluster_Override(s ServiceFabricManagedCluster, scope constructs.Construct, id *string, config *ServiceFabricManagedClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedCluster",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetBackupServiceEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"backupServiceEnabled",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetClientConnectionPort(val *float64) {
	_jsii_.Set(
		j,
		"clientConnectionPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetDnsName(val *string) {
	_jsii_.Set(
		j,
		"dnsName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetDnsServiceEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"dnsServiceEnabled",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetHttpGatewayPort(val *float64) {
	_jsii_.Set(
		j,
		"httpGatewayPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetSku(val *string) {
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetUpgradeWave(val *string) {
	_jsii_.Set(
		j,
		"upgradeWave",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedCluster) SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
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
func ServiceFabricManagedCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ServiceFabricManagedCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) PutAuthentication(value *ServiceFabricManagedClusterAuthentication) {
	_jsii_.InvokeVoid(
		s,
		"putAuthentication",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) PutCustomFabricSetting(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putCustomFabricSetting",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) PutLbRule(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putLbRule",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) PutNodeType(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putNodeType",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) PutTimeouts(value *ServiceFabricManagedClusterTimeouts) {
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetAuthentication() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetBackupServiceEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetBackupServiceEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetCustomFabricSetting() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomFabricSetting",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetDnsName() {
	_jsii_.InvokeVoid(
		s,
		"resetDnsName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetDnsServiceEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetDnsServiceEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetNodeType() {
	_jsii_.InvokeVoid(
		s,
		"resetNodeType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetPassword() {
	_jsii_.InvokeVoid(
		s,
		"resetPassword",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetSku() {
	_jsii_.InvokeVoid(
		s,
		"resetSku",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetUpgradeWave() {
	_jsii_.InvokeVoid(
		s,
		"resetUpgradeWave",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ResetUsername() {
	_jsii_.InvokeVoid(
		s,
		"resetUsername",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterAuthentication struct {
	// active_directory block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#active_directory ServiceFabricManagedCluster#active_directory}
	ActiveDirectory *ServiceFabricManagedClusterAuthenticationActiveDirectory `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#certificate ServiceFabricManagedCluster#certificate}
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
}

type ServiceFabricManagedClusterAuthenticationActiveDirectory struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#client_application_id ServiceFabricManagedCluster#client_application_id}.
	ClientApplicationId *string `field:"required" json:"clientApplicationId" yaml:"clientApplicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#cluster_application_id ServiceFabricManagedCluster#cluster_application_id}.
	ClusterApplicationId *string `field:"required" json:"clusterApplicationId" yaml:"clusterApplicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#tenant_id ServiceFabricManagedCluster#tenant_id}.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

type ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference interface {
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
	InternalValue() *ServiceFabricManagedClusterAuthenticationActiveDirectory
	SetInternalValue(val *ServiceFabricManagedClusterAuthenticationActiveDirectory)
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

// The jsii proxy struct for ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference
type jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) ClientApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) ClientApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientApplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) ClusterApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) ClusterApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterApplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) InternalValue() *ServiceFabricManagedClusterAuthenticationActiveDirectory {
	var returns *ServiceFabricManagedClusterAuthenticationActiveDirectory
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference_Override(s ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) SetClientApplicationId(val *string) {
	_jsii_.Set(
		j,
		"clientApplicationId",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) SetClusterApplicationId(val *string) {
	_jsii_.Set(
		j,
		"clusterApplicationId",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) SetInternalValue(val *ServiceFabricManagedClusterAuthenticationActiveDirectory) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterAuthenticationCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#thumbprint ServiceFabricManagedCluster#thumbprint}.
	Thumbprint *string `field:"required" json:"thumbprint" yaml:"thumbprint"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#type ServiceFabricManagedCluster#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#common_name ServiceFabricManagedCluster#common_name}.
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
}

type ServiceFabricManagedClusterAuthenticationCertificateList interface {
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
	Get(index *float64) ServiceFabricManagedClusterAuthenticationCertificateOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricManagedClusterAuthenticationCertificateList
type jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterAuthenticationCertificateList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ServiceFabricManagedClusterAuthenticationCertificateList {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterAuthenticationCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterAuthenticationCertificateList_Override(s ServiceFabricManagedClusterAuthenticationCertificateList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterAuthenticationCertificateList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) Get(index *float64) ServiceFabricManagedClusterAuthenticationCertificateOutputReference {
	var returns ServiceFabricManagedClusterAuthenticationCertificateOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterAuthenticationCertificateOutputReference interface {
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
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetCommonName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricManagedClusterAuthenticationCertificateOutputReference
type jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) ThumbprintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterAuthenticationCertificateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricManagedClusterAuthenticationCertificateOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterAuthenticationCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterAuthenticationCertificateOutputReference_Override(s ServiceFabricManagedClusterAuthenticationCertificateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterAuthenticationCertificateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) SetCommonName(val *string) {
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) SetThumbprint(val *string) {
	_jsii_.Set(
		j,
		"thumbprint",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		s,
		"resetCommonName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationCertificateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterAuthenticationOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectory() ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference
	ActiveDirectoryInput() *ServiceFabricManagedClusterAuthenticationActiveDirectory
	Certificate() ServiceFabricManagedClusterAuthenticationCertificateList
	CertificateInput() interface{}
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
	InternalValue() *ServiceFabricManagedClusterAuthentication
	SetInternalValue(val *ServiceFabricManagedClusterAuthentication)
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
	PutActiveDirectory(value *ServiceFabricManagedClusterAuthenticationActiveDirectory)
	PutCertificate(value interface{})
	ResetActiveDirectory()
	ResetCertificate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricManagedClusterAuthenticationOutputReference
type jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) ActiveDirectory() ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference {
	var returns ServiceFabricManagedClusterAuthenticationActiveDirectoryOutputReference
	_jsii_.Get(
		j,
		"activeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) ActiveDirectoryInput() *ServiceFabricManagedClusterAuthenticationActiveDirectory {
	var returns *ServiceFabricManagedClusterAuthenticationActiveDirectory
	_jsii_.Get(
		j,
		"activeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) Certificate() ServiceFabricManagedClusterAuthenticationCertificateList {
	var returns ServiceFabricManagedClusterAuthenticationCertificateList
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) CertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) InternalValue() *ServiceFabricManagedClusterAuthentication {
	var returns *ServiceFabricManagedClusterAuthentication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterAuthenticationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricManagedClusterAuthenticationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterAuthenticationOutputReference_Override(s ServiceFabricManagedClusterAuthenticationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterAuthenticationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) SetInternalValue(val *ServiceFabricManagedClusterAuthentication) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) PutActiveDirectory(value *ServiceFabricManagedClusterAuthenticationActiveDirectory) {
	_jsii_.InvokeVoid(
		s,
		"putActiveDirectory",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) PutCertificate(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putCertificate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) ResetActiveDirectory() {
	_jsii_.InvokeVoid(
		s,
		"resetActiveDirectory",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		s,
		"resetCertificate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterAuthenticationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#client_connection_port ServiceFabricManagedCluster#client_connection_port}.
	ClientConnectionPort *float64 `field:"required" json:"clientConnectionPort" yaml:"clientConnectionPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#http_gateway_port ServiceFabricManagedCluster#http_gateway_port}.
	HttpGatewayPort *float64 `field:"required" json:"httpGatewayPort" yaml:"httpGatewayPort"`
	// lb_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#lb_rule ServiceFabricManagedCluster#lb_rule}
	LbRule interface{} `field:"required" json:"lbRule" yaml:"lbRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#location ServiceFabricManagedCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#name ServiceFabricManagedCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#resource_group_name ServiceFabricManagedCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#authentication ServiceFabricManagedCluster#authentication}
	Authentication *ServiceFabricManagedClusterAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#backup_service_enabled ServiceFabricManagedCluster#backup_service_enabled}.
	BackupServiceEnabled interface{} `field:"optional" json:"backupServiceEnabled" yaml:"backupServiceEnabled"`
	// custom_fabric_setting block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#custom_fabric_setting ServiceFabricManagedCluster#custom_fabric_setting}
	CustomFabricSetting interface{} `field:"optional" json:"customFabricSetting" yaml:"customFabricSetting"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#dns_name ServiceFabricManagedCluster#dns_name}.
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#dns_service_enabled ServiceFabricManagedCluster#dns_service_enabled}.
	DnsServiceEnabled interface{} `field:"optional" json:"dnsServiceEnabled" yaml:"dnsServiceEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#id ServiceFabricManagedCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// node_type block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#node_type ServiceFabricManagedCluster#node_type}
	NodeType interface{} `field:"optional" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#password ServiceFabricManagedCluster#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#sku ServiceFabricManagedCluster#sku}.
	Sku *string `field:"optional" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#tags ServiceFabricManagedCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#timeouts ServiceFabricManagedCluster#timeouts}
	Timeouts *ServiceFabricManagedClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#upgrade_wave ServiceFabricManagedCluster#upgrade_wave}.
	UpgradeWave *string `field:"optional" json:"upgradeWave" yaml:"upgradeWave"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#username ServiceFabricManagedCluster#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

type ServiceFabricManagedClusterCustomFabricSetting struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#parameter ServiceFabricManagedCluster#parameter}.
	Parameter *string `field:"required" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#section ServiceFabricManagedCluster#section}.
	Section *string `field:"required" json:"section" yaml:"section"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#value ServiceFabricManagedCluster#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

type ServiceFabricManagedClusterCustomFabricSettingList interface {
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
	Get(index *float64) ServiceFabricManagedClusterCustomFabricSettingOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricManagedClusterCustomFabricSettingList
type jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterCustomFabricSettingList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ServiceFabricManagedClusterCustomFabricSettingList {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterCustomFabricSettingList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterCustomFabricSettingList_Override(s ServiceFabricManagedClusterCustomFabricSettingList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterCustomFabricSettingList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) Get(index *float64) ServiceFabricManagedClusterCustomFabricSettingOutputReference {
	var returns ServiceFabricManagedClusterCustomFabricSettingOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterCustomFabricSettingOutputReference interface {
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
	Parameter() *string
	SetParameter(val *string)
	ParameterInput() *string
	Section() *string
	SetSection(val *string)
	SectionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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

// The jsii proxy struct for ServiceFabricManagedClusterCustomFabricSettingOutputReference
type jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) Parameter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) ParameterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) Section() *string {
	var returns *string
	_jsii_.Get(
		j,
		"section",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) SectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterCustomFabricSettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricManagedClusterCustomFabricSettingOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterCustomFabricSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterCustomFabricSettingOutputReference_Override(s ServiceFabricManagedClusterCustomFabricSettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterCustomFabricSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) SetParameter(val *string) {
	_jsii_.Set(
		j,
		"parameter",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) SetSection(val *string) {
	_jsii_.Set(
		j,
		"section",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterCustomFabricSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterLbRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#backend_port ServiceFabricManagedCluster#backend_port}.
	BackendPort *float64 `field:"required" json:"backendPort" yaml:"backendPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#frontend_port ServiceFabricManagedCluster#frontend_port}.
	FrontendPort *float64 `field:"required" json:"frontendPort" yaml:"frontendPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#probe_protocol ServiceFabricManagedCluster#probe_protocol}.
	ProbeProtocol *string `field:"required" json:"probeProtocol" yaml:"probeProtocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#protocol ServiceFabricManagedCluster#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#probe_request_path ServiceFabricManagedCluster#probe_request_path}.
	ProbeRequestPath *string `field:"optional" json:"probeRequestPath" yaml:"probeRequestPath"`
}

type ServiceFabricManagedClusterLbRuleList interface {
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
	Get(index *float64) ServiceFabricManagedClusterLbRuleOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricManagedClusterLbRuleList
type jsiiProxy_ServiceFabricManagedClusterLbRuleList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterLbRuleList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ServiceFabricManagedClusterLbRuleList {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterLbRuleList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterLbRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterLbRuleList_Override(s ServiceFabricManagedClusterLbRuleList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterLbRuleList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleList) Get(index *float64) ServiceFabricManagedClusterLbRuleOutputReference {
	var returns ServiceFabricManagedClusterLbRuleOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterLbRuleOutputReference interface {
	cdktf.ComplexObject
	BackendPort() *float64
	SetBackendPort(val *float64)
	BackendPortInput() *float64
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
	FrontendPort() *float64
	SetFrontendPort(val *float64)
	FrontendPortInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProbeProtocol() *string
	SetProbeProtocol(val *string)
	ProbeProtocolInput() *string
	ProbeRequestPath() *string
	SetProbeRequestPath(val *string)
	ProbeRequestPathInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	ResetProbeRequestPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricManagedClusterLbRuleOutputReference
type jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) BackendPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) BackendPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backendPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) FrontendPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frontendPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) FrontendPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"frontendPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) ProbeProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probeProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) ProbeProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probeProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) ProbeRequestPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probeRequestPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) ProbeRequestPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probeRequestPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterLbRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricManagedClusterLbRuleOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterLbRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterLbRuleOutputReference_Override(s ServiceFabricManagedClusterLbRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterLbRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) SetBackendPort(val *float64) {
	_jsii_.Set(
		j,
		"backendPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) SetFrontendPort(val *float64) {
	_jsii_.Set(
		j,
		"frontendPort",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) SetProbeProtocol(val *string) {
	_jsii_.Set(
		j,
		"probeProtocol",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) SetProbeRequestPath(val *string) {
	_jsii_.Set(
		j,
		"probeRequestPath",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) SetProtocol(val *string) {
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) ResetProbeRequestPath() {
	_jsii_.InvokeVoid(
		s,
		"resetProbeRequestPath",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterLbRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterNodeType struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#application_port_range ServiceFabricManagedCluster#application_port_range}.
	ApplicationPortRange *string `field:"required" json:"applicationPortRange" yaml:"applicationPortRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#data_disk_size_gb ServiceFabricManagedCluster#data_disk_size_gb}.
	DataDiskSizeGb *float64 `field:"required" json:"dataDiskSizeGb" yaml:"dataDiskSizeGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#ephemeral_port_range ServiceFabricManagedCluster#ephemeral_port_range}.
	EphemeralPortRange *string `field:"required" json:"ephemeralPortRange" yaml:"ephemeralPortRange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#name ServiceFabricManagedCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#vm_image_offer ServiceFabricManagedCluster#vm_image_offer}.
	VmImageOffer *string `field:"required" json:"vmImageOffer" yaml:"vmImageOffer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#vm_image_publisher ServiceFabricManagedCluster#vm_image_publisher}.
	VmImagePublisher *string `field:"required" json:"vmImagePublisher" yaml:"vmImagePublisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#vm_image_sku ServiceFabricManagedCluster#vm_image_sku}.
	VmImageSku *string `field:"required" json:"vmImageSku" yaml:"vmImageSku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#vm_image_version ServiceFabricManagedCluster#vm_image_version}.
	VmImageVersion *string `field:"required" json:"vmImageVersion" yaml:"vmImageVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#vm_instance_count ServiceFabricManagedCluster#vm_instance_count}.
	VmInstanceCount *float64 `field:"required" json:"vmInstanceCount" yaml:"vmInstanceCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#vm_size ServiceFabricManagedCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#capacities ServiceFabricManagedCluster#capacities}.
	Capacities *map[string]*string `field:"optional" json:"capacities" yaml:"capacities"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#data_disk_type ServiceFabricManagedCluster#data_disk_type}.
	DataDiskType *string `field:"optional" json:"dataDiskType" yaml:"dataDiskType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#multiple_placement_groups_enabled ServiceFabricManagedCluster#multiple_placement_groups_enabled}.
	MultiplePlacementGroupsEnabled interface{} `field:"optional" json:"multiplePlacementGroupsEnabled" yaml:"multiplePlacementGroupsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#placement_properties ServiceFabricManagedCluster#placement_properties}.
	PlacementProperties *map[string]*string `field:"optional" json:"placementProperties" yaml:"placementProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#primary ServiceFabricManagedCluster#primary}.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#stateless ServiceFabricManagedCluster#stateless}.
	Stateless interface{} `field:"optional" json:"stateless" yaml:"stateless"`
	// vm_secrets block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#vm_secrets ServiceFabricManagedCluster#vm_secrets}
	VmSecrets interface{} `field:"optional" json:"vmSecrets" yaml:"vmSecrets"`
}

type ServiceFabricManagedClusterNodeTypeList interface {
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
	Get(index *float64) ServiceFabricManagedClusterNodeTypeOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricManagedClusterNodeTypeList
type jsiiProxy_ServiceFabricManagedClusterNodeTypeList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterNodeTypeList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ServiceFabricManagedClusterNodeTypeList {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterNodeTypeList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterNodeTypeList_Override(s ServiceFabricManagedClusterNodeTypeList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) Get(index *float64) ServiceFabricManagedClusterNodeTypeOutputReference {
	var returns ServiceFabricManagedClusterNodeTypeOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterNodeTypeOutputReference interface {
	cdktf.ComplexObject
	ApplicationPortRange() *string
	SetApplicationPortRange(val *string)
	ApplicationPortRangeInput() *string
	Capacities() *map[string]*string
	SetCapacities(val *map[string]*string)
	CapacitiesInput() *map[string]*string
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
	DataDiskSizeGb() *float64
	SetDataDiskSizeGb(val *float64)
	DataDiskSizeGbInput() *float64
	DataDiskType() *string
	SetDataDiskType(val *string)
	DataDiskTypeInput() *string
	EphemeralPortRange() *string
	SetEphemeralPortRange(val *string)
	EphemeralPortRangeInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MultiplePlacementGroupsEnabled() interface{}
	SetMultiplePlacementGroupsEnabled(val interface{})
	MultiplePlacementGroupsEnabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	PlacementProperties() *map[string]*string
	SetPlacementProperties(val *map[string]*string)
	PlacementPropertiesInput() *map[string]*string
	Primary() interface{}
	SetPrimary(val interface{})
	PrimaryInput() interface{}
	Stateless() interface{}
	SetStateless(val interface{})
	StatelessInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VmImageOffer() *string
	SetVmImageOffer(val *string)
	VmImageOfferInput() *string
	VmImagePublisher() *string
	SetVmImagePublisher(val *string)
	VmImagePublisherInput() *string
	VmImageSku() *string
	SetVmImageSku(val *string)
	VmImageSkuInput() *string
	VmImageVersion() *string
	SetVmImageVersion(val *string)
	VmImageVersionInput() *string
	VmInstanceCount() *float64
	SetVmInstanceCount(val *float64)
	VmInstanceCountInput() *float64
	VmSecrets() ServiceFabricManagedClusterNodeTypeVmSecretsList
	VmSecretsInput() interface{}
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
	PutVmSecrets(value interface{})
	ResetCapacities()
	ResetDataDiskType()
	ResetMultiplePlacementGroupsEnabled()
	ResetPlacementProperties()
	ResetPrimary()
	ResetStateless()
	ResetVmSecrets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricManagedClusterNodeTypeOutputReference
type jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ApplicationPortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ApplicationPortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Capacities() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"capacities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) CapacitiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"capacitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) DataDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) DataDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) DataDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) DataDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) EphemeralPortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ephemeralPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) EphemeralPortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ephemeralPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) MultiplePlacementGroupsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiplePlacementGroupsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) MultiplePlacementGroupsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiplePlacementGroupsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) PlacementProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"placementProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) PlacementPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"placementPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Primary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) PrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Stateless() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stateless",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) StatelessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statelessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImageOffer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImageOffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImageOfferInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImageOfferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImagePublisher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImagePublisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImagePublisherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImagePublisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImageSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImageSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImageSkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImageSkuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImageVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImageVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImageVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmSecrets() ServiceFabricManagedClusterNodeTypeVmSecretsList {
	var returns ServiceFabricManagedClusterNodeTypeVmSecretsList
	_jsii_.Get(
		j,
		"vmSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmSecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vmSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterNodeTypeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricManagedClusterNodeTypeOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterNodeTypeOutputReference_Override(s ServiceFabricManagedClusterNodeTypeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetApplicationPortRange(val *string) {
	_jsii_.Set(
		j,
		"applicationPortRange",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetCapacities(val *map[string]*string) {
	_jsii_.Set(
		j,
		"capacities",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetDataDiskSizeGb(val *float64) {
	_jsii_.Set(
		j,
		"dataDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetDataDiskType(val *string) {
	_jsii_.Set(
		j,
		"dataDiskType",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetEphemeralPortRange(val *string) {
	_jsii_.Set(
		j,
		"ephemeralPortRange",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetMultiplePlacementGroupsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"multiplePlacementGroupsEnabled",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetPlacementProperties(val *map[string]*string) {
	_jsii_.Set(
		j,
		"placementProperties",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetPrimary(val interface{}) {
	_jsii_.Set(
		j,
		"primary",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetStateless(val interface{}) {
	_jsii_.Set(
		j,
		"stateless",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetVmImageOffer(val *string) {
	_jsii_.Set(
		j,
		"vmImageOffer",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetVmImagePublisher(val *string) {
	_jsii_.Set(
		j,
		"vmImagePublisher",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetVmImageSku(val *string) {
	_jsii_.Set(
		j,
		"vmImageSku",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetVmImageVersion(val *string) {
	_jsii_.Set(
		j,
		"vmImageVersion",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetVmInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"vmInstanceCount",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) SetVmSize(val *string) {
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) PutVmSecrets(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putVmSecrets",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetCapacities() {
	_jsii_.InvokeVoid(
		s,
		"resetCapacities",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetDataDiskType() {
	_jsii_.InvokeVoid(
		s,
		"resetDataDiskType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetMultiplePlacementGroupsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetMultiplePlacementGroupsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetPlacementProperties() {
	_jsii_.InvokeVoid(
		s,
		"resetPlacementProperties",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetPrimary() {
	_jsii_.InvokeVoid(
		s,
		"resetPrimary",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetStateless() {
	_jsii_.InvokeVoid(
		s,
		"resetStateless",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetVmSecrets() {
	_jsii_.InvokeVoid(
		s,
		"resetVmSecrets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterNodeTypeVmSecrets struct {
	// certificates block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#certificates ServiceFabricManagedCluster#certificates}
	Certificates interface{} `field:"required" json:"certificates" yaml:"certificates"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#vault_id ServiceFabricManagedCluster#vault_id}.
	VaultId *string `field:"required" json:"vaultId" yaml:"vaultId"`
}

type ServiceFabricManagedClusterNodeTypeVmSecretsCertificates struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#store ServiceFabricManagedCluster#store}.
	Store *string `field:"required" json:"store" yaml:"store"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#url ServiceFabricManagedCluster#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

type ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList interface {
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
	Get(index *float64) ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList
type jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList_Override(s ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) Get(index *float64) ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference {
	var returns ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference interface {
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
	Store() *string
	SetStore(val *string)
	StoreInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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

// The jsii proxy struct for ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference
type jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) Store() *string {
	var returns *string
	_jsii_.Get(
		j,
		"store",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) StoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference_Override(s ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) SetStore(val *string) {
	_jsii_.Set(
		j,
		"store",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterNodeTypeVmSecretsList interface {
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
	Get(index *float64) ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricManagedClusterNodeTypeVmSecretsList
type jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterNodeTypeVmSecretsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ServiceFabricManagedClusterNodeTypeVmSecretsList {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeVmSecretsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterNodeTypeVmSecretsList_Override(s ServiceFabricManagedClusterNodeTypeVmSecretsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeVmSecretsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) Get(index *float64) ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference {
	var returns ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference interface {
	cdktf.ComplexObject
	Certificates() ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList
	CertificatesInput() interface{}
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
	VaultId() *string
	SetVaultId(val *string)
	VaultIdInput() *string
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
	PutCertificates(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference
type jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) Certificates() ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList {
	var returns ServiceFabricManagedClusterNodeTypeVmSecretsCertificatesList
	_jsii_.Get(
		j,
		"certificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) CertificatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) VaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) VaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vaultIdInput",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterNodeTypeVmSecretsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterNodeTypeVmSecretsOutputReference_Override(s ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) SetVaultId(val *string) {
	_jsii_.Set(
		j,
		"vaultId",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) PutCertificates(value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"putCertificates",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeVmSecretsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ServiceFabricManagedClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#create ServiceFabricManagedCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#delete ServiceFabricManagedCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#read ServiceFabricManagedCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_managed_cluster#update ServiceFabricManagedCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type ServiceFabricManagedClusterTimeoutsOutputReference interface {
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

// The jsii proxy struct for ServiceFabricManagedClusterTimeoutsOutputReference
type jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceFabricManagedClusterTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterTimeoutsOutputReference_Override(s ServiceFabricManagedClusterTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

