// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package opensearchserverless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceFactory("awscc_opensearchserverless_vpc_endpoint", vpcEndpointDataSource)
}

// vpcEndpointDataSource returns the Terraform awscc_opensearchserverless_vpc_endpoint data source.
// This Terraform data source corresponds to the CloudFormation AWS::OpenSearchServerless::VpcEndpoint resource.
func vpcEndpointDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the VPC Endpoint",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^vpce-[0-9a-z]*$",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the VPC Endpoint",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the VPC Endpoint",
		//	  "maxLength": 32,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z][a-z0-9-]{2,31}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the VPC Endpoint",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of one or more security groups to associate with the endpoint network interface",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 128,
		//	    "minLength": 1,
		//	    "pattern": "^[\\w+\\-]+$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 5,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "The ID of one or more security groups to associate with the endpoint network interface",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of one or more subnets in which to create an endpoint network interface",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 32,
		//	    "minLength": 1,
		//	    "pattern": "^subnet-([0-9a-f]{8}|[0-9a-f]{17})$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 6,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType:  cctypes.NewMultisetTypeOf[types.String](ctx),
			Description: "The ID of one or more subnets in which to create an endpoint network interface",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the VPC in which the endpoint will be used.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "^vpc-[0-9a-z]*$",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the VPC in which the endpoint will be used.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::OpenSearchServerless::VpcEndpoint",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::OpenSearchServerless::VpcEndpoint").WithTerraformTypeName("awscc_opensearchserverless_vpc_endpoint")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":                 "Id",
		"name":               "Name",
		"security_group_ids": "SecurityGroupIds",
		"subnet_ids":         "SubnetIds",
		"vpc_id":             "VpcId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
