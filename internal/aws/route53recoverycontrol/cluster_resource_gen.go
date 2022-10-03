// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53recoverycontrol

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
	registry.AddResourceFactory("awscc_route53recoverycontrol_cluster", clusterResource)
}

// clusterResource returns the Terraform awscc_route53recoverycontrol_cluster resource.
// This Terraform resource corresponds to the CloudFormation AWS::Route53RecoveryControl::Cluster resource.
func clusterResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"cluster_arn": {
			// Property: ClusterArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the cluster.",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the cluster.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"cluster_endpoints": {
			// Property: ClusterEndpoints
			// CloudFormation resource type schema:
			// {
			//   "description": "Endpoints for the cluster.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Endpoint": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Region": {
			//         "maxLength": 32,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "Endpoints for the cluster.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"endpoint": {
						// Property: Endpoint
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"region": {
						// Property: Region
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 32),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of a Cluster. You can use any non-white space character in the name",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of a Cluster. You can use any non-white space character in the name",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 64),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "Deployment status of a resource. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
			//   "enum": [
			//     "PENDING",
			//     "DEPLOYED",
			//     "PENDING_DELETION"
			//   ],
			//   "type": "string"
			// }
			Description: "Deployment status of a resource. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A collection of tags associated with a resource",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A collection of tags associated with a resource",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
			// Tags is a write-only property.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "AWS Route53 Recovery Control Cluster resource schema",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryControl::Cluster").WithTerraformTypeName("awscc_route53recoverycontrol_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cluster_arn":       "ClusterArn",
		"cluster_endpoints": "ClusterEndpoints",
		"endpoint":          "Endpoint",
		"key":               "Key",
		"name":              "Name",
		"region":            "Region",
		"status":            "Status",
		"tags":              "Tags",
		"value":             "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
