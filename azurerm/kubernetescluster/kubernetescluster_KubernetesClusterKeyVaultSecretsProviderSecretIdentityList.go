package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterKeyVaultSecretsProviderSecretIdentityList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterKeyVaultSecretsProviderSecretIdentityList
type jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewKubernetesClusterKeyVaultSecretsProviderSecretIdentityList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) KubernetesClusterKeyVaultSecretsProviderSecretIdentityList {
	_init_.Initialize()

	if err := validateNewKubernetesClusterKeyVaultSecretsProviderSecretIdentityListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKeyVaultSecretsProviderSecretIdentityList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewKubernetesClusterKeyVaultSecretsProviderSecretIdentityList_Override(k KubernetesClusterKeyVaultSecretsProviderSecretIdentityList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterKeyVaultSecretsProviderSecretIdentityList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) Get(index *float64) KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference {
	if err := k.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns KubernetesClusterKeyVaultSecretsProviderSecretIdentityOutputReference

	_jsii_.Invoke(
		k,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterKeyVaultSecretsProviderSecretIdentityList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

