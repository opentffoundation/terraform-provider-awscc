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
	registry.AddResourceFactory("awscc_ec2_local_gateway_route", localGatewayRouteResource)
}

// localGatewayRouteResource returns the Terraform awscc_ec2_local_gateway_route resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::LocalGatewayRoute resource.
func localGatewayRouteResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"destination_cidr_block": {
			// Property: DestinationCidrBlock
			// CloudFormation resource type schema:
			// {
			//   "description": "The CIDR block used for destination matches.",
			//   "type": "string"
			// }
			Description: "The CIDR block used for destination matches.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"local_gateway_route_table_id": {
			// Property: LocalGatewayRouteTableId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the local gateway route table.",
			//   "type": "string"
			// }
			Description: "The ID of the local gateway route table.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"local_gateway_virtual_interface_group_id": {
			// Property: LocalGatewayVirtualInterfaceGroupId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the virtual interface group.",
			//   "type": "string"
			// }
			Description: "The ID of the virtual interface group.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "The state of the route.",
			//   "type": "string"
			// }
			Description: "The state of the route.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The route type.",
			//   "type": "string"
			// }
			Description: "The route type.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
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
		Description: "Describes a route for a local gateway route table.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::LocalGatewayRoute").WithTerraformTypeName("awscc_ec2_local_gateway_route")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"destination_cidr_block":                   "DestinationCidrBlock",
		"local_gateway_route_table_id":             "LocalGatewayRouteTableId",
		"local_gateway_virtual_interface_group_id": "LocalGatewayVirtualInterfaceGroupId",
		"state": "State",
		"type":  "Type",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
