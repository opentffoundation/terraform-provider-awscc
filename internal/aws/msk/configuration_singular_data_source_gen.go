// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package msk

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
	registry.AddDataSourceFactory("awscc_msk_configuration", configurationDataSource)
}

// configurationDataSource returns the Terraform awscc_msk_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::MSK::Configuration resource.
func configurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: KafkaVersionsList
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"kafka_versions_list": schema.ListAttribute{ /*START ATTRIBUTE*/
			CustomType: cctypes.NewMultisetTypeOf[types.String](ctx),
			Computed:   true,
		}, /*END ATTRIBUTE*/
		// Property: LatestRevision
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "CreationTime": {
		//	      "type": "string"
		//	    },
		//	    "Description": {
		//	      "type": "string"
		//	    },
		//	    "Revision": {
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"latest_revision": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CreationTime
				"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Description
				"description": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Revision
				"revision": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ServerProperties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"server_properties": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MSK::Configuration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MSK::Configuration").WithTerraformTypeName("awscc_msk_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"creation_time":       "CreationTime",
		"description":         "Description",
		"kafka_versions_list": "KafkaVersionsList",
		"latest_revision":     "LatestRevision",
		"name":                "Name",
		"revision":            "Revision",
		"server_properties":   "ServerProperties",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
