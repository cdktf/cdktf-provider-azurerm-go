package appservice


type AppServiceSourceControl struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#branch AppService#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#manual_integration AppService#manual_integration}.
	ManualIntegration interface{} `field:"optional" json:"manualIntegration" yaml:"manualIntegration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#repo_url AppService#repo_url}.
	RepoUrl *string `field:"optional" json:"repoUrl" yaml:"repoUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#rollback_enabled AppService#rollback_enabled}.
	RollbackEnabled interface{} `field:"optional" json:"rollbackEnabled" yaml:"rollbackEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#use_mercurial AppService#use_mercurial}.
	UseMercurial interface{} `field:"optional" json:"useMercurial" yaml:"useMercurial"`
}

