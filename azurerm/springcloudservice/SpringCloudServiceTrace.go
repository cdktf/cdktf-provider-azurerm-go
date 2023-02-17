package springcloudservice


type SpringCloudServiceTrace struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#connection_string SpringCloudService#connection_string}.
	ConnectionString *string `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_service#sample_rate SpringCloudService#sample_rate}.
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

