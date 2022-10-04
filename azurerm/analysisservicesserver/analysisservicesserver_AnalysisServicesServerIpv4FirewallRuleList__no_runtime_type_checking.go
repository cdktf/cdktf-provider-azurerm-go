//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package analysisservicesserver

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AnalysisServicesServerIpv4FirewallRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AnalysisServicesServerIpv4FirewallRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AnalysisServicesServerIpv4FirewallRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AnalysisServicesServerIpv4FirewallRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AnalysisServicesServerIpv4FirewallRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AnalysisServicesServerIpv4FirewallRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAnalysisServicesServerIpv4FirewallRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

