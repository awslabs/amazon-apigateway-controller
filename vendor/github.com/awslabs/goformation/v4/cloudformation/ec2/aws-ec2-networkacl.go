package ec2

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/awslabs/goformation/v4/cloudformation/policies"
	"github.com/awslabs/goformation/v4/cloudformation/tags"
)

// NetworkAcl AWS CloudFormation Resource (AWS::EC2::NetworkAcl)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html
type NetworkAcl struct {

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html#cfn-ec2-networkacl-tags
	Tags []tags.Tag `json:"Tags,omitempty"`

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html#cfn-ec2-networkacl-vpcid
	VpcId string `json:"VpcId,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *NetworkAcl) AWSCloudFormationType() string {
	return "AWS::EC2::NetworkAcl"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r NetworkAcl) MarshalJSON() ([]byte, error) {
	type Properties NetworkAcl
	return json.Marshal(&struct {
		Type           string
		Properties     Properties
		DependsOn      []string                `json:"DependsOn,omitempty"`
		Metadata       map[string]interface{}  `json:"Metadata,omitempty"`
		DeletionPolicy policies.DeletionPolicy `json:"DeletionPolicy,omitempty"`
		Condition      string                  `json:"Condition,omitempty"`
	}{
		Type:           r.AWSCloudFormationType(),
		Properties:     (Properties)(r),
		DependsOn:      r.AWSCloudFormationDependsOn,
		Metadata:       r.AWSCloudFormationMetadata,
		DeletionPolicy: r.AWSCloudFormationDeletionPolicy,
		Condition:      r.AWSCloudFormationCondition,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *NetworkAcl) UnmarshalJSON(b []byte) error {
	type Properties NetworkAcl
	res := &struct {
		Type           string
		Properties     *Properties
		DependsOn      []string
		Metadata       map[string]interface{}
		DeletionPolicy string
		Condition      string
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	if err := dec.Decode(&res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = NetworkAcl(*res.Properties)
	}
	if res.DependsOn != nil {
		r.AWSCloudFormationDependsOn = res.DependsOn
	}
	if res.Metadata != nil {
		r.AWSCloudFormationMetadata = res.Metadata
	}
	if res.DeletionPolicy != "" {
		r.AWSCloudFormationDeletionPolicy = policies.DeletionPolicy(res.DeletionPolicy)
	}
	if res.Condition != "" {
		r.AWSCloudFormationCondition = res.Condition
	}
	return nil
}