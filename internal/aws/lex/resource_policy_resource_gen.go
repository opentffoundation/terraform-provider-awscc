// Code generated by generators/resource/main.go; DO NOT EDIT.

package lex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_lex_resource_policy", resourcePolicyResource)
}

// resourcePolicyResource returns the Terraform awscc_lex_resource_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::Lex::ResourcePolicy resource.
func resourcePolicyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The Physical ID of the resource policy.",
			//   "type": "string"
			// }
			Description: "The Physical ID of the resource policy.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"policy": {
			// Property: Policy
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource policy to add to the resource. The policy is a JSON structure following the IAM syntax that contains one or more statements that define the policy.",
			//   "type": "object"
			// }
			Description: "A resource policy to add to the resource. The policy is a JSON structure following the IAM syntax that contains one or more statements that define the policy.",
			Type:        types.MapType{ElemType: types.StringType},
			Required:    true,
		},
		"resource_arn": {
			// Property: ResourceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.",
			//   "maxLength": 1011,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 1011),
			},
		},
		"revision_id": {
			// Property: RevisionId
			// CloudFormation resource type schema:
			// {
			//   "description": "The current revision of the resource policy. Use the revision ID to make sure that you are updating the most current version of a resource policy when you add a policy statement to a resource, delete a resource, or update a resource.",
			//   "maxLength": 5,
			//   "minLength": 1,
			//   "pattern": "^[0-9]+$",
			//   "type": "string"
			// }
			Description: "The current revision of the resource policy. Use the revision ID to make sure that you are updating the most current version of a resource policy when you add a policy statement to a resource, delete a resource, or update a resource.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "A resource policy with specified policy statements that attaches to a Lex bot or bot alias.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lex::ResourcePolicy").WithTerraformTypeName("awscc_lex_resource_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":           "Id",
		"policy":       "Policy",
		"resource_arn": "ResourceArn",
		"revision_id":  "RevisionId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
