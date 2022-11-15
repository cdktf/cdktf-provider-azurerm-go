package vpnserverconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/vpnserverconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration azurerm_vpn_server_configuration}.
type VpnServerConfiguration interface {
	cdktf.TerraformResource
	AzureActiveDirectoryAuthentication() VpnServerConfigurationAzureActiveDirectoryAuthenticationList
	AzureActiveDirectoryAuthenticationInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientRevokedCertificate() VpnServerConfigurationClientRevokedCertificateList
	ClientRevokedCertificateInput() interface{}
	ClientRootCertificate() VpnServerConfigurationClientRootCertificateList
	ClientRootCertificateInput() interface{}
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
	IpsecPolicy() VpnServerConfigurationIpsecPolicyOutputReference
	IpsecPolicyInput() *VpnServerConfigurationIpsecPolicy
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
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Radius() VpnServerConfigurationRadiusOutputReference
	RadiusInput() *VpnServerConfigurationRadius
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VpnServerConfigurationTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpnAuthenticationTypes() *[]*string
	SetVpnAuthenticationTypes(val *[]*string)
	VpnAuthenticationTypesInput() *[]*string
	VpnProtocols() *[]*string
	SetVpnProtocols(val *[]*string)
	VpnProtocolsInput() *[]*string
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
	PutAzureActiveDirectoryAuthentication(value interface{})
	PutClientRevokedCertificate(value interface{})
	PutClientRootCertificate(value interface{})
	PutIpsecPolicy(value *VpnServerConfigurationIpsecPolicy)
	PutRadius(value *VpnServerConfigurationRadius)
	PutTimeouts(value *VpnServerConfigurationTimeouts)
	ResetAzureActiveDirectoryAuthentication()
	ResetClientRevokedCertificate()
	ResetClientRootCertificate()
	ResetId()
	ResetIpsecPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRadius()
	ResetTags()
	ResetTimeouts()
	ResetVpnProtocols()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VpnServerConfiguration
type jsiiProxy_VpnServerConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_VpnServerConfiguration) AzureActiveDirectoryAuthentication() VpnServerConfigurationAzureActiveDirectoryAuthenticationList {
	var returns VpnServerConfigurationAzureActiveDirectoryAuthenticationList
	_jsii_.Get(
		j,
		"azureActiveDirectoryAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) AzureActiveDirectoryAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureActiveDirectoryAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) ClientRevokedCertificate() VpnServerConfigurationClientRevokedCertificateList {
	var returns VpnServerConfigurationClientRevokedCertificateList
	_jsii_.Get(
		j,
		"clientRevokedCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) ClientRevokedCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientRevokedCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) ClientRootCertificate() VpnServerConfigurationClientRootCertificateList {
	var returns VpnServerConfigurationClientRootCertificateList
	_jsii_.Get(
		j,
		"clientRootCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) ClientRootCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientRootCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) IpsecPolicy() VpnServerConfigurationIpsecPolicyOutputReference {
	var returns VpnServerConfigurationIpsecPolicyOutputReference
	_jsii_.Get(
		j,
		"ipsecPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) IpsecPolicyInput() *VpnServerConfigurationIpsecPolicy {
	var returns *VpnServerConfigurationIpsecPolicy
	_jsii_.Get(
		j,
		"ipsecPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) Radius() VpnServerConfigurationRadiusOutputReference {
	var returns VpnServerConfigurationRadiusOutputReference
	_jsii_.Get(
		j,
		"radius",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) RadiusInput() *VpnServerConfigurationRadius {
	var returns *VpnServerConfigurationRadius
	_jsii_.Get(
		j,
		"radiusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) Timeouts() VpnServerConfigurationTimeoutsOutputReference {
	var returns VpnServerConfigurationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) VpnAuthenticationTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpnAuthenticationTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) VpnAuthenticationTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpnAuthenticationTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) VpnProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpnProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfiguration) VpnProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpnProtocolsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration azurerm_vpn_server_configuration} Resource.
func NewVpnServerConfiguration(scope constructs.Construct, id *string, config *VpnServerConfigurationConfig) VpnServerConfiguration {
	_init_.Initialize()

	if err := validateNewVpnServerConfigurationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnServerConfiguration{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnServerConfiguration.VpnServerConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration azurerm_vpn_server_configuration} Resource.
func NewVpnServerConfiguration_Override(v VpnServerConfiguration, scope constructs.Construct, id *string, config *VpnServerConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnServerConfiguration.VpnServerConfiguration",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetVpnAuthenticationTypes(val *[]*string) {
	if err := j.validateSetVpnAuthenticationTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnAuthenticationTypes",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfiguration)SetVpnProtocols(val *[]*string) {
	if err := j.validateSetVpnProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnProtocols",
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
func VpnServerConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpnServerConfiguration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.vpnServerConfiguration.VpnServerConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpnServerConfiguration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpnServerConfiguration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.vpnServerConfiguration.VpnServerConfiguration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpnServerConfiguration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpnServerConfiguration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.vpnServerConfiguration.VpnServerConfiguration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpnServerConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.vpnServerConfiguration.VpnServerConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpnServerConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpnServerConfiguration) PutAzureActiveDirectoryAuthentication(value interface{}) {
	if err := v.validatePutAzureActiveDirectoryAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAzureActiveDirectoryAuthentication",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnServerConfiguration) PutClientRevokedCertificate(value interface{}) {
	if err := v.validatePutClientRevokedCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putClientRevokedCertificate",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnServerConfiguration) PutClientRootCertificate(value interface{}) {
	if err := v.validatePutClientRootCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putClientRootCertificate",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnServerConfiguration) PutIpsecPolicy(value *VpnServerConfigurationIpsecPolicy) {
	if err := v.validatePutIpsecPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putIpsecPolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnServerConfiguration) PutRadius(value *VpnServerConfigurationRadius) {
	if err := v.validatePutRadiusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRadius",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnServerConfiguration) PutTimeouts(value *VpnServerConfigurationTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnServerConfiguration) ResetAzureActiveDirectoryAuthentication() {
	_jsii_.InvokeVoid(
		v,
		"resetAzureActiveDirectoryAuthentication",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnServerConfiguration) ResetClientRevokedCertificate() {
	_jsii_.InvokeVoid(
		v,
		"resetClientRevokedCertificate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnServerConfiguration) ResetClientRootCertificate() {
	_jsii_.InvokeVoid(
		v,
		"resetClientRootCertificate",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnServerConfiguration) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnServerConfiguration) ResetIpsecPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetIpsecPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnServerConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnServerConfiguration) ResetRadius() {
	_jsii_.InvokeVoid(
		v,
		"resetRadius",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnServerConfiguration) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnServerConfiguration) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnServerConfiguration) ResetVpnProtocols() {
	_jsii_.InvokeVoid(
		v,
		"resetVpnProtocols",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnServerConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

