// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_transit_gateway_vpc_attachment", transitGatewayVpcAttachmentResource)
}

// transitGatewayVpcAttachmentResource returns the Terraform awscc_ec2_transit_gateway_vpc_attachment resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::TransitGatewayVpcAttachment resource.
func transitGatewayVpcAttachmentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"add_subnet_ids": {
			// Property: AddSubnetIds
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"options": {
			// Property: Options
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The options for the transit gateway vpc attachment.",
			//   "properties": {
			//     "ApplianceModeSupport": {
			//       "description": "Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable",
			//       "type": "string"
			//     },
			//     "DnsSupport": {
			//       "description": "Indicates whether to enable DNS Support for Vpc Attachment. Valid Values: enable | disable",
			//       "type": "string"
			//     },
			//     "Ipv6Support": {
			//       "description": "Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The options for the transit gateway vpc attachment.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"appliance_mode_support": {
						// Property: ApplianceModeSupport
						Description: "Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"dns_support": {
						// Property: DnsSupport
						Description: "Indicates whether to enable DNS Support for Vpc Attachment. Valid Values: enable | disable",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"ipv_6_support": {
						// Property: Ipv6Support
						Description: "Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"remove_subnet_ids": {
			// Property: RemoveSubnetIds
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"subnet_ids": {
			// Property: SubnetIds
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
				resource.UseStateForUnknown(),
			},
		},
		"transit_gateway_id": {
			// Property: TransitGatewayId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::EC2::TransitGatewayVpcAttachment",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayVpcAttachment").WithTerraformTypeName("awscc_ec2_transit_gateway_vpc_attachment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"add_subnet_ids":         "AddSubnetIds",
		"appliance_mode_support": "ApplianceModeSupport",
		"dns_support":            "DnsSupport",
		"id":                     "Id",
		"ipv_6_support":          "Ipv6Support",
		"key":                    "Key",
		"options":                "Options",
		"remove_subnet_ids":      "RemoveSubnetIds",
		"subnet_ids":             "SubnetIds",
		"tags":                   "Tags",
		"transit_gateway_id":     "TransitGatewayId",
		"value":                  "Value",
		"vpc_id":                 "VpcId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
