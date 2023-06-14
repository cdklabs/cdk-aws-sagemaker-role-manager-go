package cdklabscdkawssagemakerrolemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Experimental.
type KMSOptions struct {
	// Experimental.
	DataKeys *[]awskms.IKey `field:"optional" json:"dataKeys" yaml:"dataKeys"`
	// Experimental.
	VolumeKeys *[]awskms.IKey `field:"optional" json:"volumeKeys" yaml:"volumeKeys"`
}

