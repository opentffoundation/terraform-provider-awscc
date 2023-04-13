// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotwireless_wireless_device_import_task", wirelessDeviceImportTaskDataSource)
}

// wirelessDeviceImportTaskDataSource returns the Terraform awscc_iotwireless_wireless_device_import_task data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTWireless::WirelessDeviceImportTask resource.
func wirelessDeviceImportTaskDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn for Wireless Device Import Task, Returned upon successful start.",
		//	  "maxLength": 128,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn for Wireless Device Import Task, Returned upon successful start.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreationDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "CreationDate for import task",
		//	  "type": "string"
		//	}
		"creation_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "CreationDate for import task",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DestinationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Destination Name for import task",
		//	  "maxLength": 128,
		//	  "pattern": "[a-zA-Z0-9-_]+",
		//	  "type": "string"
		//	}
		"destination_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Destination Name for import task",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FailedImportedDevicesCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Failed Imported Devices Count",
		//	  "type": "integer"
		//	}
		"failed_imported_devices_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Failed Imported Devices Count",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id for Wireless Device Import Task, Returned upon successful start.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id for Wireless Device Import Task, Returned upon successful start.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InitializedImportedDevicesCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Initialized Imported Devices Count",
		//	  "type": "integer"
		//	}
		"initialized_imported_devices_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Initialized Imported Devices Count",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OnboardedImportedDevicesCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Onboarded Imported Devices Count",
		//	  "type": "integer"
		//	}
		"onboarded_imported_devices_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Onboarded Imported Devices Count",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PendingImportedDevicesCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Pending Imported Devices Count",
		//	  "type": "integer"
		//	}
		"pending_imported_devices_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Pending Imported Devices Count",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Sidewalk
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "sidewalk contain file for created device and role",
		//	  "oneOf": [
		//	    {
		//	      "allOf": [
		//	        {
		//	          "required": [
		//	            "DeviceCreationFile"
		//	          ]
		//	        },
		//	        {
		//	          "required": [
		//	            "Role"
		//	          ]
		//	        }
		//	      ]
		//	    },
		//	    {
		//	      "required": [
		//	        "SidewalkManufacturingSn"
		//	      ]
		//	    }
		//	  ],
		//	  "properties": {
		//	    "DeviceCreationFile": {
		//	      "maxLength": 1024,
		//	      "type": "string"
		//	    },
		//	    "DeviceCreationFileList": {
		//	      "description": "sidewalk create device's file path",
		//	      "items": {
		//	        "maxLength": 1024,
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "Role": {
		//	      "description": "sidewalk role",
		//	      "maxLength": 2048,
		//	      "type": "string"
		//	    },
		//	    "SidewalkManufacturingSn": {
		//	      "maxLength": 64,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"sidewalk": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DeviceCreationFile
				"device_creation_file": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DeviceCreationFileList
				"device_creation_file_list": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "sidewalk create device's file path",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Role
				"role": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "sidewalk role",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SidewalkManufacturingSn
				"sidewalk_manufacturing_sn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "sidewalk contain file for created device and role",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Status for import task",
		//	  "enum": [
		//	    "INITIALIZING",
		//	    "INITIALIZED",
		//	    "PENDING",
		//	    "COMPLETE",
		//	    "FAILED",
		//	    "DELETING"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Status for import task",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StatusReason
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "StatusReason for import task",
		//	  "type": "string"
		//	}
		"status_reason": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "StatusReason for import task",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTWireless::WirelessDeviceImportTask",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::WirelessDeviceImportTask").WithTerraformTypeName("awscc_iotwireless_wireless_device_import_task")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                "Arn",
		"creation_date":                      "CreationDate",
		"destination_name":                   "DestinationName",
		"device_creation_file":               "DeviceCreationFile",
		"device_creation_file_list":          "DeviceCreationFileList",
		"failed_imported_devices_count":      "FailedImportedDevicesCount",
		"id":                                 "Id",
		"initialized_imported_devices_count": "InitializedImportedDevicesCount",
		"key":                                "Key",
		"onboarded_imported_devices_count":   "OnboardedImportedDevicesCount",
		"pending_imported_devices_count":     "PendingImportedDevicesCount",
		"role":                               "Role",
		"sidewalk":                           "Sidewalk",
		"sidewalk_manufacturing_sn":          "SidewalkManufacturingSn",
		"status":                             "Status",
		"status_reason":                      "StatusReason",
		"tags":                               "Tags",
		"value":                              "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
