package mediacontentkeypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/mediacontentkeypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy azurerm_media_content_key_policy}.
type MediaContentKeyPolicy interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	MediaServicesAccountName() *string
	SetMediaServicesAccountName(val *string)
	MediaServicesAccountNameInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PolicyOption() MediaContentKeyPolicyPolicyOptionList
	PolicyOptionInput() interface{}
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MediaContentKeyPolicyTimeoutsOutputReference
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
	PutPolicyOption(value interface{})
	PutTimeouts(value *MediaContentKeyPolicyTimeouts)
	ResetDescription()
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

// The jsii proxy struct for MediaContentKeyPolicy
type jsiiProxy_MediaContentKeyPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MediaContentKeyPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) MediaServicesAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaServicesAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) MediaServicesAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaServicesAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) PolicyOption() MediaContentKeyPolicyPolicyOptionList {
	var returns MediaContentKeyPolicyPolicyOptionList
	_jsii_.Get(
		j,
		"policyOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) PolicyOptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) Timeouts() MediaContentKeyPolicyTimeoutsOutputReference {
	var returns MediaContentKeyPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy azurerm_media_content_key_policy} Resource.
func NewMediaContentKeyPolicy(scope constructs.Construct, id *string, config *MediaContentKeyPolicyConfig) MediaContentKeyPolicy {
	_init_.Initialize()

	j := jsiiProxy_MediaContentKeyPolicy{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy azurerm_media_content_key_policy} Resource.
func NewMediaContentKeyPolicy_Override(m MediaContentKeyPolicy, scope constructs.Construct, id *string, config *MediaContentKeyPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicy",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicy) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicy) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicy) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicy) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicy) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicy) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicy) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicy) SetMediaServicesAccountName(val *string) {
	_jsii_.Set(
		j,
		"mediaServicesAccountName",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicy) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicy) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicy) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicy) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
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
func MediaContentKeyPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MediaContentKeyPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicy) PutPolicyOption(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putPolicyOption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicy) PutTimeouts(value *MediaContentKeyPolicyTimeouts) {
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicy) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicy) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaContentKeyPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#media_services_account_name MediaContentKeyPolicy#media_services_account_name}.
	MediaServicesAccountName *string `field:"required" json:"mediaServicesAccountName" yaml:"mediaServicesAccountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#name MediaContentKeyPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// policy_option block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#policy_option MediaContentKeyPolicy#policy_option}
	PolicyOption interface{} `field:"required" json:"policyOption" yaml:"policyOption"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#resource_group_name MediaContentKeyPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#description MediaContentKeyPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#id MediaContentKeyPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#timeouts MediaContentKeyPolicy#timeouts}
	Timeouts *MediaContentKeyPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type MediaContentKeyPolicyPolicyOption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#name MediaContentKeyPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#clear_key_configuration_enabled MediaContentKeyPolicy#clear_key_configuration_enabled}.
	ClearKeyConfigurationEnabled interface{} `field:"optional" json:"clearKeyConfigurationEnabled" yaml:"clearKeyConfigurationEnabled"`
	// fairplay_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#fairplay_configuration MediaContentKeyPolicy#fairplay_configuration}
	FairplayConfiguration *MediaContentKeyPolicyPolicyOptionFairplayConfiguration `field:"optional" json:"fairplayConfiguration" yaml:"fairplayConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#open_restriction_enabled MediaContentKeyPolicy#open_restriction_enabled}.
	OpenRestrictionEnabled interface{} `field:"optional" json:"openRestrictionEnabled" yaml:"openRestrictionEnabled"`
	// playready_configuration_license block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#playready_configuration_license MediaContentKeyPolicy#playready_configuration_license}
	PlayreadyConfigurationLicense interface{} `field:"optional" json:"playreadyConfigurationLicense" yaml:"playreadyConfigurationLicense"`
	// token_restriction block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#token_restriction MediaContentKeyPolicy#token_restriction}
	TokenRestriction *MediaContentKeyPolicyPolicyOptionTokenRestriction `field:"optional" json:"tokenRestriction" yaml:"tokenRestriction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#widevine_configuration_template MediaContentKeyPolicy#widevine_configuration_template}.
	WidevineConfigurationTemplate *string `field:"optional" json:"widevineConfigurationTemplate" yaml:"widevineConfigurationTemplate"`
}

type MediaContentKeyPolicyPolicyOptionFairplayConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#ask MediaContentKeyPolicy#ask}.
	Ask *string `field:"optional" json:"ask" yaml:"ask"`
	// offline_rental_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#offline_rental_configuration MediaContentKeyPolicy#offline_rental_configuration}
	OfflineRentalConfiguration *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration `field:"optional" json:"offlineRentalConfiguration" yaml:"offlineRentalConfiguration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#pfx MediaContentKeyPolicy#pfx}.
	Pfx *string `field:"optional" json:"pfx" yaml:"pfx"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#pfx_password MediaContentKeyPolicy#pfx_password}.
	PfxPassword *string `field:"optional" json:"pfxPassword" yaml:"pfxPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#rental_and_lease_key_type MediaContentKeyPolicy#rental_and_lease_key_type}.
	RentalAndLeaseKeyType *string `field:"optional" json:"rentalAndLeaseKeyType" yaml:"rentalAndLeaseKeyType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#rental_duration_seconds MediaContentKeyPolicy#rental_duration_seconds}.
	RentalDurationSeconds *float64 `field:"optional" json:"rentalDurationSeconds" yaml:"rentalDurationSeconds"`
}

type MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#playback_duration_seconds MediaContentKeyPolicy#playback_duration_seconds}.
	PlaybackDurationSeconds *float64 `field:"optional" json:"playbackDurationSeconds" yaml:"playbackDurationSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#storage_duration_seconds MediaContentKeyPolicy#storage_duration_seconds}.
	StorageDurationSeconds *float64 `field:"optional" json:"storageDurationSeconds" yaml:"storageDurationSeconds"`
}

type MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference interface {
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
	InternalValue() *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration
	SetInternalValue(val *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration)
	PlaybackDurationSeconds() *float64
	SetPlaybackDurationSeconds(val *float64)
	PlaybackDurationSecondsInput() *float64
	StorageDurationSeconds() *float64
	SetStorageDurationSeconds(val *float64)
	StorageDurationSecondsInput() *float64
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
	ResetPlaybackDurationSeconds()
	ResetStorageDurationSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference
type jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) InternalValue() *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration {
	var returns *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) PlaybackDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"playbackDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) PlaybackDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"playbackDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) StorageDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) StorageDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference_Override(m MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) SetInternalValue(val *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) SetPlaybackDurationSeconds(val *float64) {
	_jsii_.Set(
		j,
		"playbackDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) SetStorageDurationSeconds(val *float64) {
	_jsii_.Set(
		j,
		"storageDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) ResetPlaybackDurationSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetPlaybackDurationSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) ResetStorageDurationSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageDurationSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference interface {
	cdktf.ComplexObject
	Ask() *string
	SetAsk(val *string)
	AskInput() *string
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
	InternalValue() *MediaContentKeyPolicyPolicyOptionFairplayConfiguration
	SetInternalValue(val *MediaContentKeyPolicyPolicyOptionFairplayConfiguration)
	OfflineRentalConfiguration() MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference
	OfflineRentalConfigurationInput() *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration
	Pfx() *string
	SetPfx(val *string)
	PfxInput() *string
	PfxPassword() *string
	SetPfxPassword(val *string)
	PfxPasswordInput() *string
	RentalAndLeaseKeyType() *string
	SetRentalAndLeaseKeyType(val *string)
	RentalAndLeaseKeyTypeInput() *string
	RentalDurationSeconds() *float64
	SetRentalDurationSeconds(val *float64)
	RentalDurationSecondsInput() *float64
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
	PutOfflineRentalConfiguration(value *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration)
	ResetAsk()
	ResetOfflineRentalConfiguration()
	ResetPfx()
	ResetPfxPassword()
	ResetRentalAndLeaseKeyType()
	ResetRentalDurationSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference
type jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) Ask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) AskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"askInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) InternalValue() *MediaContentKeyPolicyPolicyOptionFairplayConfiguration {
	var returns *MediaContentKeyPolicyPolicyOptionFairplayConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) OfflineRentalConfiguration() MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference {
	var returns MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfigurationOutputReference
	_jsii_.Get(
		j,
		"offlineRentalConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) OfflineRentalConfigurationInput() *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration {
	var returns *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration
	_jsii_.Get(
		j,
		"offlineRentalConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) Pfx() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) PfxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) PfxPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) PfxPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) RentalAndLeaseKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rentalAndLeaseKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) RentalAndLeaseKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rentalAndLeaseKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) RentalDurationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rentalDurationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) RentalDurationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rentalDurationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference_Override(m MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) SetAsk(val *string) {
	_jsii_.Set(
		j,
		"ask",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) SetInternalValue(val *MediaContentKeyPolicyPolicyOptionFairplayConfiguration) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) SetPfx(val *string) {
	_jsii_.Set(
		j,
		"pfx",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) SetPfxPassword(val *string) {
	_jsii_.Set(
		j,
		"pfxPassword",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) SetRentalAndLeaseKeyType(val *string) {
	_jsii_.Set(
		j,
		"rentalAndLeaseKeyType",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) SetRentalDurationSeconds(val *float64) {
	_jsii_.Set(
		j,
		"rentalDurationSeconds",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) PutOfflineRentalConfiguration(value *MediaContentKeyPolicyPolicyOptionFairplayConfigurationOfflineRentalConfiguration) {
	_jsii_.InvokeVoid(
		m,
		"putOfflineRentalConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ResetAsk() {
	_jsii_.InvokeVoid(
		m,
		"resetAsk",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ResetOfflineRentalConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetOfflineRentalConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ResetPfx() {
	_jsii_.InvokeVoid(
		m,
		"resetPfx",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ResetPfxPassword() {
	_jsii_.InvokeVoid(
		m,
		"resetPfxPassword",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ResetRentalAndLeaseKeyType() {
	_jsii_.InvokeVoid(
		m,
		"resetRentalAndLeaseKeyType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ResetRentalDurationSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetRentalDurationSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaContentKeyPolicyPolicyOptionList interface {
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
	Get(index *float64) MediaContentKeyPolicyPolicyOptionOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionList
type jsiiProxy_MediaContentKeyPolicyPolicyOptionList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MediaContentKeyPolicyPolicyOptionList {
	_init_.Initialize()

	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionList_Override(m MediaContentKeyPolicyPolicyOptionList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) Get(index *float64) MediaContentKeyPolicyPolicyOptionOutputReference {
	var returns MediaContentKeyPolicyPolicyOptionOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaContentKeyPolicyPolicyOptionOutputReference interface {
	cdktf.ComplexObject
	ClearKeyConfigurationEnabled() interface{}
	SetClearKeyConfigurationEnabled(val interface{})
	ClearKeyConfigurationEnabledInput() interface{}
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
	FairplayConfiguration() MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference
	FairplayConfigurationInput() *MediaContentKeyPolicyPolicyOptionFairplayConfiguration
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	OpenRestrictionEnabled() interface{}
	SetOpenRestrictionEnabled(val interface{})
	OpenRestrictionEnabledInput() interface{}
	PlayreadyConfigurationLicense() MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList
	PlayreadyConfigurationLicenseInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenRestriction() MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference
	TokenRestrictionInput() *MediaContentKeyPolicyPolicyOptionTokenRestriction
	WidevineConfigurationTemplate() *string
	SetWidevineConfigurationTemplate(val *string)
	WidevineConfigurationTemplateInput() *string
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
	PutFairplayConfiguration(value *MediaContentKeyPolicyPolicyOptionFairplayConfiguration)
	PutPlayreadyConfigurationLicense(value interface{})
	PutTokenRestriction(value *MediaContentKeyPolicyPolicyOptionTokenRestriction)
	ResetClearKeyConfigurationEnabled()
	ResetFairplayConfiguration()
	ResetOpenRestrictionEnabled()
	ResetPlayreadyConfigurationLicense()
	ResetTokenRestriction()
	ResetWidevineConfigurationTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionOutputReference
type jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ClearKeyConfigurationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clearKeyConfigurationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ClearKeyConfigurationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clearKeyConfigurationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) FairplayConfiguration() MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference {
	var returns MediaContentKeyPolicyPolicyOptionFairplayConfigurationOutputReference
	_jsii_.Get(
		j,
		"fairplayConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) FairplayConfigurationInput() *MediaContentKeyPolicyPolicyOptionFairplayConfiguration {
	var returns *MediaContentKeyPolicyPolicyOptionFairplayConfiguration
	_jsii_.Get(
		j,
		"fairplayConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) OpenRestrictionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openRestrictionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) OpenRestrictionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openRestrictionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) PlayreadyConfigurationLicense() MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList {
	var returns MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList
	_jsii_.Get(
		j,
		"playreadyConfigurationLicense",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) PlayreadyConfigurationLicenseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"playreadyConfigurationLicenseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) TokenRestriction() MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference {
	var returns MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference
	_jsii_.Get(
		j,
		"tokenRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) TokenRestrictionInput() *MediaContentKeyPolicyPolicyOptionTokenRestriction {
	var returns *MediaContentKeyPolicyPolicyOptionTokenRestriction
	_jsii_.Get(
		j,
		"tokenRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) WidevineConfigurationTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widevineConfigurationTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) WidevineConfigurationTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"widevineConfigurationTemplateInput",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediaContentKeyPolicyPolicyOptionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionOutputReference_Override(m MediaContentKeyPolicyPolicyOptionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) SetClearKeyConfigurationEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"clearKeyConfigurationEnabled",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) SetOpenRestrictionEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"openRestrictionEnabled",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) SetWidevineConfigurationTemplate(val *string) {
	_jsii_.Set(
		j,
		"widevineConfigurationTemplate",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) PutFairplayConfiguration(value *MediaContentKeyPolicyPolicyOptionFairplayConfiguration) {
	_jsii_.InvokeVoid(
		m,
		"putFairplayConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) PutPlayreadyConfigurationLicense(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putPlayreadyConfigurationLicense",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) PutTokenRestriction(value *MediaContentKeyPolicyPolicyOptionTokenRestriction) {
	_jsii_.InvokeVoid(
		m,
		"putTokenRestriction",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ResetClearKeyConfigurationEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetClearKeyConfigurationEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ResetFairplayConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetFairplayConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ResetOpenRestrictionEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetOpenRestrictionEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ResetPlayreadyConfigurationLicense() {
	_jsii_.InvokeVoid(
		m,
		"resetPlayreadyConfigurationLicense",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ResetTokenRestriction() {
	_jsii_.InvokeVoid(
		m,
		"resetTokenRestriction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ResetWidevineConfigurationTemplate() {
	_jsii_.InvokeVoid(
		m,
		"resetWidevineConfigurationTemplate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicense struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#allow_test_devices MediaContentKeyPolicy#allow_test_devices}.
	AllowTestDevices interface{} `field:"optional" json:"allowTestDevices" yaml:"allowTestDevices"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#begin_date MediaContentKeyPolicy#begin_date}.
	BeginDate *string `field:"optional" json:"beginDate" yaml:"beginDate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#content_key_location_from_header_enabled MediaContentKeyPolicy#content_key_location_from_header_enabled}.
	ContentKeyLocationFromHeaderEnabled interface{} `field:"optional" json:"contentKeyLocationFromHeaderEnabled" yaml:"contentKeyLocationFromHeaderEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#content_key_location_from_key_id MediaContentKeyPolicy#content_key_location_from_key_id}.
	ContentKeyLocationFromKeyId *string `field:"optional" json:"contentKeyLocationFromKeyId" yaml:"contentKeyLocationFromKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#content_type MediaContentKeyPolicy#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#expiration_date MediaContentKeyPolicy#expiration_date}.
	ExpirationDate *string `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#grace_period MediaContentKeyPolicy#grace_period}.
	GracePeriod *string `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#license_type MediaContentKeyPolicy#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// play_right block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#play_right MediaContentKeyPolicy#play_right}
	PlayRight *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight `field:"optional" json:"playRight" yaml:"playRight"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#relative_begin_date MediaContentKeyPolicy#relative_begin_date}.
	RelativeBeginDate *string `field:"optional" json:"relativeBeginDate" yaml:"relativeBeginDate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#relative_expiration_date MediaContentKeyPolicy#relative_expiration_date}.
	RelativeExpirationDate *string `field:"optional" json:"relativeExpirationDate" yaml:"relativeExpirationDate"`
}

type MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList interface {
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
	Get(index *float64) MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList
type jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList {
	_init_.Initialize()

	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList_Override(m MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) Get(index *float64) MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference {
	var returns MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference interface {
	cdktf.ComplexObject
	AllowTestDevices() interface{}
	SetAllowTestDevices(val interface{})
	AllowTestDevicesInput() interface{}
	BeginDate() *string
	SetBeginDate(val *string)
	BeginDateInput() *string
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
	ContentKeyLocationFromHeaderEnabled() interface{}
	SetContentKeyLocationFromHeaderEnabled(val interface{})
	ContentKeyLocationFromHeaderEnabledInput() interface{}
	ContentKeyLocationFromKeyId() *string
	SetContentKeyLocationFromKeyId(val *string)
	ContentKeyLocationFromKeyIdInput() *string
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExpirationDate() *string
	SetExpirationDate(val *string)
	ExpirationDateInput() *string
	// Experimental.
	Fqn() *string
	GracePeriod() *string
	SetGracePeriod(val *string)
	GracePeriodInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	PlayRight() MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference
	PlayRightInput() *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight
	RelativeBeginDate() *string
	SetRelativeBeginDate(val *string)
	RelativeBeginDateInput() *string
	RelativeExpirationDate() *string
	SetRelativeExpirationDate(val *string)
	RelativeExpirationDateInput() *string
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
	PutPlayRight(value *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight)
	ResetAllowTestDevices()
	ResetBeginDate()
	ResetContentKeyLocationFromHeaderEnabled()
	ResetContentKeyLocationFromKeyId()
	ResetContentType()
	ResetExpirationDate()
	ResetGracePeriod()
	ResetLicenseType()
	ResetPlayRight()
	ResetRelativeBeginDate()
	ResetRelativeExpirationDate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference
type jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) AllowTestDevices() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowTestDevices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) AllowTestDevicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowTestDevicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) BeginDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) BeginDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ContentKeyLocationFromHeaderEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentKeyLocationFromHeaderEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ContentKeyLocationFromHeaderEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentKeyLocationFromHeaderEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ContentKeyLocationFromKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentKeyLocationFromKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ContentKeyLocationFromKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentKeyLocationFromKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ExpirationDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GracePeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GracePeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) PlayRight() MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference {
	var returns MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference
	_jsii_.Get(
		j,
		"playRight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) PlayRightInput() *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight {
	var returns *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight
	_jsii_.Get(
		j,
		"playRightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) RelativeBeginDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeBeginDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) RelativeBeginDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeBeginDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) RelativeExpirationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeExpirationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) RelativeExpirationDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeExpirationDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference_Override(m MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetAllowTestDevices(val interface{}) {
	_jsii_.Set(
		j,
		"allowTestDevices",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetBeginDate(val *string) {
	_jsii_.Set(
		j,
		"beginDate",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetContentKeyLocationFromHeaderEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"contentKeyLocationFromHeaderEnabled",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetContentKeyLocationFromKeyId(val *string) {
	_jsii_.Set(
		j,
		"contentKeyLocationFromKeyId",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetContentType(val *string) {
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetExpirationDate(val *string) {
	_jsii_.Set(
		j,
		"expirationDate",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetGracePeriod(val *string) {
	_jsii_.Set(
		j,
		"gracePeriod",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetLicenseType(val *string) {
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetRelativeBeginDate(val *string) {
	_jsii_.Set(
		j,
		"relativeBeginDate",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetRelativeExpirationDate(val *string) {
	_jsii_.Set(
		j,
		"relativeExpirationDate",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) PutPlayRight(value *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight) {
	_jsii_.InvokeVoid(
		m,
		"putPlayRight",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetAllowTestDevices() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowTestDevices",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetBeginDate() {
	_jsii_.InvokeVoid(
		m,
		"resetBeginDate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetContentKeyLocationFromHeaderEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetContentKeyLocationFromHeaderEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetContentKeyLocationFromKeyId() {
	_jsii_.InvokeVoid(
		m,
		"resetContentKeyLocationFromKeyId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetContentType() {
	_jsii_.InvokeVoid(
		m,
		"resetContentType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetExpirationDate() {
	_jsii_.InvokeVoid(
		m,
		"resetExpirationDate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetGracePeriod() {
	_jsii_.InvokeVoid(
		m,
		"resetGracePeriod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetLicenseType() {
	_jsii_.InvokeVoid(
		m,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetPlayRight() {
	_jsii_.InvokeVoid(
		m,
		"resetPlayRight",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetRelativeBeginDate() {
	_jsii_.InvokeVoid(
		m,
		"resetRelativeBeginDate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ResetRelativeExpirationDate() {
	_jsii_.InvokeVoid(
		m,
		"resetRelativeExpirationDate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicenseOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#agc_and_color_stripe_restriction MediaContentKeyPolicy#agc_and_color_stripe_restriction}.
	AgcAndColorStripeRestriction *float64 `field:"optional" json:"agcAndColorStripeRestriction" yaml:"agcAndColorStripeRestriction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#allow_passing_video_content_to_unknown_output MediaContentKeyPolicy#allow_passing_video_content_to_unknown_output}.
	AllowPassingVideoContentToUnknownOutput *string `field:"optional" json:"allowPassingVideoContentToUnknownOutput" yaml:"allowPassingVideoContentToUnknownOutput"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#analog_video_opl MediaContentKeyPolicy#analog_video_opl}.
	AnalogVideoOpl *float64 `field:"optional" json:"analogVideoOpl" yaml:"analogVideoOpl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#compressed_digital_audio_opl MediaContentKeyPolicy#compressed_digital_audio_opl}.
	CompressedDigitalAudioOpl *float64 `field:"optional" json:"compressedDigitalAudioOpl" yaml:"compressedDigitalAudioOpl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#digital_video_only_content_restriction MediaContentKeyPolicy#digital_video_only_content_restriction}.
	DigitalVideoOnlyContentRestriction interface{} `field:"optional" json:"digitalVideoOnlyContentRestriction" yaml:"digitalVideoOnlyContentRestriction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#first_play_expiration MediaContentKeyPolicy#first_play_expiration}.
	FirstPlayExpiration *string `field:"optional" json:"firstPlayExpiration" yaml:"firstPlayExpiration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#image_constraint_for_analog_component_video_restriction MediaContentKeyPolicy#image_constraint_for_analog_component_video_restriction}.
	ImageConstraintForAnalogComponentVideoRestriction interface{} `field:"optional" json:"imageConstraintForAnalogComponentVideoRestriction" yaml:"imageConstraintForAnalogComponentVideoRestriction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#image_constraint_for_analog_computer_monitor_restriction MediaContentKeyPolicy#image_constraint_for_analog_computer_monitor_restriction}.
	ImageConstraintForAnalogComputerMonitorRestriction interface{} `field:"optional" json:"imageConstraintForAnalogComputerMonitorRestriction" yaml:"imageConstraintForAnalogComputerMonitorRestriction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#scms_restriction MediaContentKeyPolicy#scms_restriction}.
	ScmsRestriction *float64 `field:"optional" json:"scmsRestriction" yaml:"scmsRestriction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#uncompressed_digital_audio_opl MediaContentKeyPolicy#uncompressed_digital_audio_opl}.
	UncompressedDigitalAudioOpl *float64 `field:"optional" json:"uncompressedDigitalAudioOpl" yaml:"uncompressedDigitalAudioOpl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#uncompressed_digital_video_opl MediaContentKeyPolicy#uncompressed_digital_video_opl}.
	UncompressedDigitalVideoOpl *float64 `field:"optional" json:"uncompressedDigitalVideoOpl" yaml:"uncompressedDigitalVideoOpl"`
}

type MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference interface {
	cdktf.ComplexObject
	AgcAndColorStripeRestriction() *float64
	SetAgcAndColorStripeRestriction(val *float64)
	AgcAndColorStripeRestrictionInput() *float64
	AllowPassingVideoContentToUnknownOutput() *string
	SetAllowPassingVideoContentToUnknownOutput(val *string)
	AllowPassingVideoContentToUnknownOutputInput() *string
	AnalogVideoOpl() *float64
	SetAnalogVideoOpl(val *float64)
	AnalogVideoOplInput() *float64
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
	CompressedDigitalAudioOpl() *float64
	SetCompressedDigitalAudioOpl(val *float64)
	CompressedDigitalAudioOplInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DigitalVideoOnlyContentRestriction() interface{}
	SetDigitalVideoOnlyContentRestriction(val interface{})
	DigitalVideoOnlyContentRestrictionInput() interface{}
	FirstPlayExpiration() *string
	SetFirstPlayExpiration(val *string)
	FirstPlayExpirationInput() *string
	// Experimental.
	Fqn() *string
	ImageConstraintForAnalogComponentVideoRestriction() interface{}
	SetImageConstraintForAnalogComponentVideoRestriction(val interface{})
	ImageConstraintForAnalogComponentVideoRestrictionInput() interface{}
	ImageConstraintForAnalogComputerMonitorRestriction() interface{}
	SetImageConstraintForAnalogComputerMonitorRestriction(val interface{})
	ImageConstraintForAnalogComputerMonitorRestrictionInput() interface{}
	InternalValue() *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight
	SetInternalValue(val *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight)
	ScmsRestriction() *float64
	SetScmsRestriction(val *float64)
	ScmsRestrictionInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UncompressedDigitalAudioOpl() *float64
	SetUncompressedDigitalAudioOpl(val *float64)
	UncompressedDigitalAudioOplInput() *float64
	UncompressedDigitalVideoOpl() *float64
	SetUncompressedDigitalVideoOpl(val *float64)
	UncompressedDigitalVideoOplInput() *float64
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
	ResetAgcAndColorStripeRestriction()
	ResetAllowPassingVideoContentToUnknownOutput()
	ResetAnalogVideoOpl()
	ResetCompressedDigitalAudioOpl()
	ResetDigitalVideoOnlyContentRestriction()
	ResetFirstPlayExpiration()
	ResetImageConstraintForAnalogComponentVideoRestriction()
	ResetImageConstraintForAnalogComputerMonitorRestriction()
	ResetScmsRestriction()
	ResetUncompressedDigitalAudioOpl()
	ResetUncompressedDigitalVideoOpl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference
type jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) AgcAndColorStripeRestriction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"agcAndColorStripeRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) AgcAndColorStripeRestrictionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"agcAndColorStripeRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) AllowPassingVideoContentToUnknownOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowPassingVideoContentToUnknownOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) AllowPassingVideoContentToUnknownOutputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowPassingVideoContentToUnknownOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) AnalogVideoOpl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"analogVideoOpl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) AnalogVideoOplInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"analogVideoOplInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) CompressedDigitalAudioOpl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"compressedDigitalAudioOpl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) CompressedDigitalAudioOplInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"compressedDigitalAudioOplInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) DigitalVideoOnlyContentRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"digitalVideoOnlyContentRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) DigitalVideoOnlyContentRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"digitalVideoOnlyContentRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) FirstPlayExpiration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstPlayExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) FirstPlayExpirationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstPlayExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ImageConstraintForAnalogComponentVideoRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConstraintForAnalogComponentVideoRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ImageConstraintForAnalogComponentVideoRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConstraintForAnalogComponentVideoRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ImageConstraintForAnalogComputerMonitorRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConstraintForAnalogComputerMonitorRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ImageConstraintForAnalogComputerMonitorRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConstraintForAnalogComputerMonitorRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) InternalValue() *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight {
	var returns *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ScmsRestriction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scmsRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ScmsRestrictionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scmsRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) UncompressedDigitalAudioOpl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uncompressedDigitalAudioOpl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) UncompressedDigitalAudioOplInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uncompressedDigitalAudioOplInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) UncompressedDigitalVideoOpl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uncompressedDigitalVideoOpl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) UncompressedDigitalVideoOplInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"uncompressedDigitalVideoOplInput",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference_Override(m MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetAgcAndColorStripeRestriction(val *float64) {
	_jsii_.Set(
		j,
		"agcAndColorStripeRestriction",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetAllowPassingVideoContentToUnknownOutput(val *string) {
	_jsii_.Set(
		j,
		"allowPassingVideoContentToUnknownOutput",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetAnalogVideoOpl(val *float64) {
	_jsii_.Set(
		j,
		"analogVideoOpl",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetCompressedDigitalAudioOpl(val *float64) {
	_jsii_.Set(
		j,
		"compressedDigitalAudioOpl",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetDigitalVideoOnlyContentRestriction(val interface{}) {
	_jsii_.Set(
		j,
		"digitalVideoOnlyContentRestriction",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetFirstPlayExpiration(val *string) {
	_jsii_.Set(
		j,
		"firstPlayExpiration",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetImageConstraintForAnalogComponentVideoRestriction(val interface{}) {
	_jsii_.Set(
		j,
		"imageConstraintForAnalogComponentVideoRestriction",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetImageConstraintForAnalogComputerMonitorRestriction(val interface{}) {
	_jsii_.Set(
		j,
		"imageConstraintForAnalogComputerMonitorRestriction",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetInternalValue(val *MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRight) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetScmsRestriction(val *float64) {
	_jsii_.Set(
		j,
		"scmsRestriction",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetUncompressedDigitalAudioOpl(val *float64) {
	_jsii_.Set(
		j,
		"uncompressedDigitalAudioOpl",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) SetUncompressedDigitalVideoOpl(val *float64) {
	_jsii_.Set(
		j,
		"uncompressedDigitalVideoOpl",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetAgcAndColorStripeRestriction() {
	_jsii_.InvokeVoid(
		m,
		"resetAgcAndColorStripeRestriction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetAllowPassingVideoContentToUnknownOutput() {
	_jsii_.InvokeVoid(
		m,
		"resetAllowPassingVideoContentToUnknownOutput",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetAnalogVideoOpl() {
	_jsii_.InvokeVoid(
		m,
		"resetAnalogVideoOpl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetCompressedDigitalAudioOpl() {
	_jsii_.InvokeVoid(
		m,
		"resetCompressedDigitalAudioOpl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetDigitalVideoOnlyContentRestriction() {
	_jsii_.InvokeVoid(
		m,
		"resetDigitalVideoOnlyContentRestriction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetFirstPlayExpiration() {
	_jsii_.InvokeVoid(
		m,
		"resetFirstPlayExpiration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetImageConstraintForAnalogComponentVideoRestriction() {
	_jsii_.InvokeVoid(
		m,
		"resetImageConstraintForAnalogComponentVideoRestriction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetImageConstraintForAnalogComputerMonitorRestriction() {
	_jsii_.InvokeVoid(
		m,
		"resetImageConstraintForAnalogComputerMonitorRestriction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetScmsRestriction() {
	_jsii_.InvokeVoid(
		m,
		"resetScmsRestriction",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetUncompressedDigitalAudioOpl() {
	_jsii_.InvokeVoid(
		m,
		"resetUncompressedDigitalAudioOpl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ResetUncompressedDigitalVideoOpl() {
	_jsii_.InvokeVoid(
		m,
		"resetUncompressedDigitalVideoOpl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionPlayreadyConfigurationLicensePlayRightOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaContentKeyPolicyPolicyOptionTokenRestriction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#audience MediaContentKeyPolicy#audience}.
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#issuer MediaContentKeyPolicy#issuer}.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#open_id_connect_discovery_document MediaContentKeyPolicy#open_id_connect_discovery_document}.
	OpenIdConnectDiscoveryDocument *string `field:"optional" json:"openIdConnectDiscoveryDocument" yaml:"openIdConnectDiscoveryDocument"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#primary_rsa_token_key_exponent MediaContentKeyPolicy#primary_rsa_token_key_exponent}.
	PrimaryRsaTokenKeyExponent *string `field:"optional" json:"primaryRsaTokenKeyExponent" yaml:"primaryRsaTokenKeyExponent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#primary_rsa_token_key_modulus MediaContentKeyPolicy#primary_rsa_token_key_modulus}.
	PrimaryRsaTokenKeyModulus *string `field:"optional" json:"primaryRsaTokenKeyModulus" yaml:"primaryRsaTokenKeyModulus"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#primary_symmetric_token_key MediaContentKeyPolicy#primary_symmetric_token_key}.
	PrimarySymmetricTokenKey *string `field:"optional" json:"primarySymmetricTokenKey" yaml:"primarySymmetricTokenKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#primary_x509_token_key_raw MediaContentKeyPolicy#primary_x509_token_key_raw}.
	PrimaryX509TokenKeyRaw *string `field:"optional" json:"primaryX509TokenKeyRaw" yaml:"primaryX509TokenKeyRaw"`
	// required_claim block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#required_claim MediaContentKeyPolicy#required_claim}
	RequiredClaim interface{} `field:"optional" json:"requiredClaim" yaml:"requiredClaim"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#token_type MediaContentKeyPolicy#token_type}.
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
}

type MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference interface {
	cdktf.ComplexObject
	Audience() *string
	SetAudience(val *string)
	AudienceInput() *string
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
	InternalValue() *MediaContentKeyPolicyPolicyOptionTokenRestriction
	SetInternalValue(val *MediaContentKeyPolicyPolicyOptionTokenRestriction)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	OpenIdConnectDiscoveryDocument() *string
	SetOpenIdConnectDiscoveryDocument(val *string)
	OpenIdConnectDiscoveryDocumentInput() *string
	PrimaryRsaTokenKeyExponent() *string
	SetPrimaryRsaTokenKeyExponent(val *string)
	PrimaryRsaTokenKeyExponentInput() *string
	PrimaryRsaTokenKeyModulus() *string
	SetPrimaryRsaTokenKeyModulus(val *string)
	PrimaryRsaTokenKeyModulusInput() *string
	PrimarySymmetricTokenKey() *string
	SetPrimarySymmetricTokenKey(val *string)
	PrimarySymmetricTokenKeyInput() *string
	PrimaryX509TokenKeyRaw() *string
	SetPrimaryX509TokenKeyRaw(val *string)
	PrimaryX509TokenKeyRawInput() *string
	RequiredClaim() MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList
	RequiredClaimInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenType() *string
	SetTokenType(val *string)
	TokenTypeInput() *string
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
	PutRequiredClaim(value interface{})
	ResetAudience()
	ResetIssuer()
	ResetOpenIdConnectDiscoveryDocument()
	ResetPrimaryRsaTokenKeyExponent()
	ResetPrimaryRsaTokenKeyModulus()
	ResetPrimarySymmetricTokenKey()
	ResetPrimaryX509TokenKeyRaw()
	ResetRequiredClaim()
	ResetTokenType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference
type jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) Audience() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audience",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) AudienceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audienceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) InternalValue() *MediaContentKeyPolicyPolicyOptionTokenRestriction {
	var returns *MediaContentKeyPolicyPolicyOptionTokenRestriction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) OpenIdConnectDiscoveryDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectDiscoveryDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) OpenIdConnectDiscoveryDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openIdConnectDiscoveryDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimaryRsaTokenKeyExponent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryRsaTokenKeyExponent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimaryRsaTokenKeyExponentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryRsaTokenKeyExponentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimaryRsaTokenKeyModulus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryRsaTokenKeyModulus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimaryRsaTokenKeyModulusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryRsaTokenKeyModulusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimarySymmetricTokenKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primarySymmetricTokenKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimarySymmetricTokenKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primarySymmetricTokenKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimaryX509TokenKeyRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryX509TokenKeyRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PrimaryX509TokenKeyRawInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryX509TokenKeyRawInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) RequiredClaim() MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList {
	var returns MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList
	_jsii_.Get(
		j,
		"requiredClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) RequiredClaimInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference_Override(m MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) SetAudience(val *string) {
	_jsii_.Set(
		j,
		"audience",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) SetInternalValue(val *MediaContentKeyPolicyPolicyOptionTokenRestriction) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) SetIssuer(val *string) {
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) SetOpenIdConnectDiscoveryDocument(val *string) {
	_jsii_.Set(
		j,
		"openIdConnectDiscoveryDocument",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) SetPrimaryRsaTokenKeyExponent(val *string) {
	_jsii_.Set(
		j,
		"primaryRsaTokenKeyExponent",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) SetPrimaryRsaTokenKeyModulus(val *string) {
	_jsii_.Set(
		j,
		"primaryRsaTokenKeyModulus",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) SetPrimarySymmetricTokenKey(val *string) {
	_jsii_.Set(
		j,
		"primarySymmetricTokenKey",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) SetPrimaryX509TokenKeyRaw(val *string) {
	_jsii_.Set(
		j,
		"primaryX509TokenKeyRaw",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) SetTokenType(val *string) {
	_jsii_.Set(
		j,
		"tokenType",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) PutRequiredClaim(value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"putRequiredClaim",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetAudience() {
	_jsii_.InvokeVoid(
		m,
		"resetAudience",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		m,
		"resetIssuer",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetOpenIdConnectDiscoveryDocument() {
	_jsii_.InvokeVoid(
		m,
		"resetOpenIdConnectDiscoveryDocument",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetPrimaryRsaTokenKeyExponent() {
	_jsii_.InvokeVoid(
		m,
		"resetPrimaryRsaTokenKeyExponent",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetPrimaryRsaTokenKeyModulus() {
	_jsii_.InvokeVoid(
		m,
		"resetPrimaryRsaTokenKeyModulus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetPrimarySymmetricTokenKey() {
	_jsii_.InvokeVoid(
		m,
		"resetPrimarySymmetricTokenKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetPrimaryX509TokenKeyRaw() {
	_jsii_.InvokeVoid(
		m,
		"resetPrimaryX509TokenKeyRaw",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetRequiredClaim() {
	_jsii_.InvokeVoid(
		m,
		"resetRequiredClaim",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ResetTokenType() {
	_jsii_.InvokeVoid(
		m,
		"resetTokenType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaim struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#type MediaContentKeyPolicy#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#value MediaContentKeyPolicy#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

type MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList interface {
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
	Get(index *float64) MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList
type jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList {
	_init_.Initialize()

	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList_Override(m MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) Get(index *float64) MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference {
	var returns MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference

	_jsii_.Invoke(
		m,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference interface {
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
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetType()
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference
type jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference_Override(m MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) SetValue(val *string) {
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		m,
		"resetType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		m,
		"resetValue",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyPolicyOptionTokenRestrictionRequiredClaimOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type MediaContentKeyPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#create MediaContentKeyPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#delete MediaContentKeyPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#read MediaContentKeyPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/media_content_key_policy#update MediaContentKeyPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type MediaContentKeyPolicyTimeoutsOutputReference interface {
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

// The jsii proxy struct for MediaContentKeyPolicyTimeoutsOutputReference
type jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewMediaContentKeyPolicyTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MediaContentKeyPolicyTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaContentKeyPolicyTimeoutsOutputReference_Override(m MediaContentKeyPolicyTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mediaContentKeyPolicy.MediaContentKeyPolicyTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		m,
		"resetCreate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		m,
		"resetDelete",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		m,
		"resetRead",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		m,
		"resetUpdate",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaContentKeyPolicyTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

