// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_vpc_endpoint_service", vPCEndpointServiceDataSource)
}

// vPCEndpointServiceDataSource returns the Terraform awscc_ec2_vpc_endpoint_service data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::VPCEndpointService resource.
func vPCEndpointServiceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AcceptanceRequired
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"acceptance_required": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ContributorInsightsEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"contributor_insights_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: GatewayLoadBalancerArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"gateway_load_balancer_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkLoadBalancerArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "relationshipRef": {
		//	      "propertyPath": "/properties/LoadBalancerArn",
		//	      "typeName": "AWS::ElasticLoadBalancingV2::LoadBalancer"
		//	    },
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"network_load_balancer_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PayerResponsibility
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"payer_responsibility": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"service_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::VPCEndpointService",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VPCEndpointService").WithTerraformTypeName("awscc_ec2_vpc_endpoint_service")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"acceptance_required":          "AcceptanceRequired",
		"contributor_insights_enabled": "ContributorInsightsEnabled",
		"gateway_load_balancer_arns":   "GatewayLoadBalancerArns",
		"network_load_balancer_arns":   "NetworkLoadBalancerArns",
		"payer_responsibility":         "PayerResponsibility",
		"service_id":                   "ServiceId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
