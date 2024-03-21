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
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_customer_gateway", customerGatewayDataSource)
}

// customerGatewayDataSource returns the Terraform awscc_ec2_customer_gateway data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::CustomerGateway resource.
func customerGatewayDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BgpAsn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 65000,
		//	  "description": "For devices that support BGP, the customer gateway's BGP ASN.",
		//	  "type": "integer"
		//	}
		"bgp_asn": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "For devices that support BGP, the customer gateway's BGP ASN.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CustomerGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "CustomerGateway ID generated after customer gateway is created. Each customer gateway has a unique ID.",
		//	  "type": "string"
		//	}
		"customer_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "CustomerGateway ID generated after customer gateway is created. Each customer gateway has a unique ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeviceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the customer gateway device.",
		//	  "type": "string"
		//	}
		"device_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the customer gateway device.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IpAddress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The internet-routable IP address for the customer gateway's outside interface. The address must be static.",
		//	  "type": "string"
		//	}
		"ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The internet-routable IP address for the customer gateway's outside interface. The address must be static.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more tags for the customer gateway.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.NewMultisetTypeOf[types.Object](ctx),
			Description: "One or more tags for the customer gateway.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of VPN connection that this customer gateway supports.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of VPN connection that this customer gateway supports.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::CustomerGateway",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::CustomerGateway").WithTerraformTypeName("awscc_ec2_customer_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bgp_asn":             "BgpAsn",
		"customer_gateway_id": "CustomerGatewayId",
		"device_name":         "DeviceName",
		"ip_address":          "IpAddress",
		"key":                 "Key",
		"tags":                "Tags",
		"type":                "Type",
		"value":               "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
