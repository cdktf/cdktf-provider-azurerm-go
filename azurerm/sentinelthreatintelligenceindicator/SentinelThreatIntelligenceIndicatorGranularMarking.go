package sentinelthreatintelligenceindicator


type SentinelThreatIntelligenceIndicatorGranularMarking struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_threat_intelligence_indicator#language SentinelThreatIntelligenceIndicator#language}.
	Language *string `field:"optional" json:"language" yaml:"language"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_threat_intelligence_indicator#marking_ref SentinelThreatIntelligenceIndicator#marking_ref}.
	MarkingRef *string `field:"optional" json:"markingRef" yaml:"markingRef"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_threat_intelligence_indicator#selectors SentinelThreatIntelligenceIndicator#selectors}.
	Selectors *[]*string `field:"optional" json:"selectors" yaml:"selectors"`
}

