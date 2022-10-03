package frontdoorcustomhttpsconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/frontdoorcustomhttpsconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration azurerm_frontdoor_custom_https_configuration}.
type FrontdoorCustomHttpsConfiguration interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CustomHttpsConfiguration() FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference
	CustomHttpsConfigurationInput() *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration
	CustomHttpsProvisioningEnabled() interface{}
	SetCustomHttpsProvisioningEnabled(val interface{})
	CustomHttpsProvisioningEnabledInput() interface{}
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
	FrontendEndpointId() *string
	SetFrontendEndpointId(val *string)
	FrontendEndpointIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() FrontdoorCustomHttpsConfigurationTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutCustomHttpsConfiguration(value *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration)
	PutTimeouts(value *FrontdoorCustomHttpsConfigurationTimeouts)
	ResetCustomHttpsConfiguration()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FrontdoorCustomHttpsConfiguration
type jsiiProxy_FrontdoorCustomHttpsConfiguration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) CustomHttpsConfiguration() FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference {
	var returns FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference
	_jsii_.Get(
		j,
		"customHttpsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) CustomHttpsConfigurationInput() *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration {
	var returns *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration
	_jsii_.Get(
		j,
		"customHttpsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) CustomHttpsProvisioningEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customHttpsProvisioningEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) CustomHttpsProvisioningEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customHttpsProvisioningEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) FrontendEndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendEndpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) FrontendEndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"frontendEndpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) Timeouts() FrontdoorCustomHttpsConfigurationTimeoutsOutputReference {
	var returns FrontdoorCustomHttpsConfigurationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration azurerm_frontdoor_custom_https_configuration} Resource.
func NewFrontdoorCustomHttpsConfiguration(scope constructs.Construct, id *string, config *FrontdoorCustomHttpsConfigurationConfig) FrontdoorCustomHttpsConfiguration {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorCustomHttpsConfiguration{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoorCustomHttpsConfiguration.FrontdoorCustomHttpsConfiguration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration azurerm_frontdoor_custom_https_configuration} Resource.
func NewFrontdoorCustomHttpsConfiguration_Override(f FrontdoorCustomHttpsConfiguration, scope constructs.Construct, id *string, config *FrontdoorCustomHttpsConfigurationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoorCustomHttpsConfiguration.FrontdoorCustomHttpsConfiguration",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) SetCustomHttpsProvisioningEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"customHttpsProvisioningEnabled",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) SetFrontendEndpointId(val *string) {
	_jsii_.Set(
		j,
		"frontendEndpointId",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfiguration) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
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
func FrontdoorCustomHttpsConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.frontdoorCustomHttpsConfiguration.FrontdoorCustomHttpsConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FrontdoorCustomHttpsConfiguration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.frontdoorCustomHttpsConfiguration.FrontdoorCustomHttpsConfiguration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) PutCustomHttpsConfiguration(value *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration) {
	_jsii_.InvokeVoid(
		f,
		"putCustomHttpsConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) PutTimeouts(value *FrontdoorCustomHttpsConfigurationTimeouts) {
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) ResetCustomHttpsConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomHttpsConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfiguration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorCustomHttpsConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration#custom_https_provisioning_enabled FrontdoorCustomHttpsConfiguration#custom_https_provisioning_enabled}.
	CustomHttpsProvisioningEnabled interface{} `field:"required" json:"customHttpsProvisioningEnabled" yaml:"customHttpsProvisioningEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration#frontend_endpoint_id FrontdoorCustomHttpsConfiguration#frontend_endpoint_id}.
	FrontendEndpointId *string `field:"required" json:"frontendEndpointId" yaml:"frontendEndpointId"`
	// custom_https_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration#custom_https_configuration FrontdoorCustomHttpsConfiguration#custom_https_configuration}
	CustomHttpsConfiguration *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration `field:"optional" json:"customHttpsConfiguration" yaml:"customHttpsConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration#id FrontdoorCustomHttpsConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration#timeouts FrontdoorCustomHttpsConfiguration#timeouts}
	Timeouts *FrontdoorCustomHttpsConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration#azure_key_vault_certificate_secret_name FrontdoorCustomHttpsConfiguration#azure_key_vault_certificate_secret_name}.
	AzureKeyVaultCertificateSecretName *string `field:"optional" json:"azureKeyVaultCertificateSecretName" yaml:"azureKeyVaultCertificateSecretName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration#azure_key_vault_certificate_secret_version FrontdoorCustomHttpsConfiguration#azure_key_vault_certificate_secret_version}.
	AzureKeyVaultCertificateSecretVersion *string `field:"optional" json:"azureKeyVaultCertificateSecretVersion" yaml:"azureKeyVaultCertificateSecretVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration#azure_key_vault_certificate_vault_id FrontdoorCustomHttpsConfiguration#azure_key_vault_certificate_vault_id}.
	AzureKeyVaultCertificateVaultId *string `field:"optional" json:"azureKeyVaultCertificateVaultId" yaml:"azureKeyVaultCertificateVaultId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration#certificate_source FrontdoorCustomHttpsConfiguration#certificate_source}.
	CertificateSource *string `field:"optional" json:"certificateSource" yaml:"certificateSource"`
}

type FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference interface {
	cdktf.ComplexObject
	AzureKeyVaultCertificateSecretName() *string
	SetAzureKeyVaultCertificateSecretName(val *string)
	AzureKeyVaultCertificateSecretNameInput() *string
	AzureKeyVaultCertificateSecretVersion() *string
	SetAzureKeyVaultCertificateSecretVersion(val *string)
	AzureKeyVaultCertificateSecretVersionInput() *string
	AzureKeyVaultCertificateVaultId() *string
	SetAzureKeyVaultCertificateVaultId(val *string)
	AzureKeyVaultCertificateVaultIdInput() *string
	CertificateSource() *string
	SetCertificateSource(val *string)
	CertificateSourceInput() *string
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
	InternalValue() *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration
	SetInternalValue(val *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration)
	MinimumTlsVersion() *string
	ProvisioningState() *string
	ProvisioningSubstate() *string
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
	ResetAzureKeyVaultCertificateSecretName()
	ResetAzureKeyVaultCertificateSecretVersion()
	ResetAzureKeyVaultCertificateVaultId()
	ResetCertificateSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference
type jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) AzureKeyVaultCertificateSecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultCertificateSecretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) AzureKeyVaultCertificateSecretNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultCertificateSecretNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) AzureKeyVaultCertificateSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultCertificateSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) AzureKeyVaultCertificateSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultCertificateSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) AzureKeyVaultCertificateVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultCertificateVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) AzureKeyVaultCertificateVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azureKeyVaultCertificateVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) CertificateSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) CertificateSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) InternalValue() *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration {
	var returns *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) MinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ProvisioningState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ProvisioningSubstate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provisioningSubstate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoorCustomHttpsConfiguration.FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference_Override(f FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoorCustomHttpsConfiguration.FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) SetAzureKeyVaultCertificateSecretName(val *string) {
	_jsii_.Set(
		j,
		"azureKeyVaultCertificateSecretName",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) SetAzureKeyVaultCertificateSecretVersion(val *string) {
	_jsii_.Set(
		j,
		"azureKeyVaultCertificateSecretVersion",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) SetAzureKeyVaultCertificateVaultId(val *string) {
	_jsii_.Set(
		j,
		"azureKeyVaultCertificateVaultId",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) SetCertificateSource(val *string) {
	_jsii_.Set(
		j,
		"certificateSource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) SetInternalValue(val *FrontdoorCustomHttpsConfigurationCustomHttpsConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ResetAzureKeyVaultCertificateSecretName() {
	_jsii_.InvokeVoid(
		f,
		"resetAzureKeyVaultCertificateSecretName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ResetAzureKeyVaultCertificateSecretVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetAzureKeyVaultCertificateSecretVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ResetAzureKeyVaultCertificateVaultId() {
	_jsii_.InvokeVoid(
		f,
		"resetAzureKeyVaultCertificateVaultId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ResetCertificateSource() {
	_jsii_.InvokeVoid(
		f,
		"resetCertificateSource",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationCustomHttpsConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type FrontdoorCustomHttpsConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration#create FrontdoorCustomHttpsConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration#delete FrontdoorCustomHttpsConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration#read FrontdoorCustomHttpsConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor_custom_https_configuration#update FrontdoorCustomHttpsConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type FrontdoorCustomHttpsConfigurationTimeoutsOutputReference interface {
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

// The jsii proxy struct for FrontdoorCustomHttpsConfigurationTimeoutsOutputReference
type jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewFrontdoorCustomHttpsConfigurationTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FrontdoorCustomHttpsConfigurationTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoorCustomHttpsConfiguration.FrontdoorCustomHttpsConfigurationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFrontdoorCustomHttpsConfigurationTimeoutsOutputReference_Override(f FrontdoorCustomHttpsConfigurationTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoorCustomHttpsConfiguration.FrontdoorCustomHttpsConfigurationTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		f,
		"resetCreate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		f,
		"resetDelete",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		f,
		"resetRead",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		f,
		"resetUpdate",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorCustomHttpsConfigurationTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

