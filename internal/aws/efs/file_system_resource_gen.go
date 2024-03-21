// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package efs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_efs_file_system", fileSystemResource)
}

// fileSystemResource returns the Terraform awscc_efs_file_system resource.
// This Terraform resource corresponds to the CloudFormation AWS::EFS::FileSystem resource.
func fileSystemResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZoneName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "For One Zone file systems, specify the AWS Availability Zone in which to create the file system. Use the format ``us-east-1a`` to specify the Availability Zone. For more information about One Zone file systems, see [EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/availability-durability.html#file-system-type) in the *Amazon EFS User Guide*.\n  One Zone file systems are not available in all Availability Zones in AWS-Regions where Amazon EFS is available.",
		//	  "type": "string"
		//	}
		"availability_zone_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "For One Zone file systems, specify the AWS Availability Zone in which to create the file system. Use the format ``us-east-1a`` to specify the Availability Zone. For more information about One Zone file systems, see [EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/availability-durability.html#file-system-type) in the *Amazon EFS User Guide*.\n  One Zone file systems are not available in all Availability Zones in AWS-Regions where Amazon EFS is available.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BackupPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Use the ``BackupPolicy`` to turn automatic backups on or off for the file system.",
		//	  "properties": {
		//	    "Status": {
		//	      "description": "Set the backup policy status for the file system.\n  +   *ENABLED* - Turns automatic backups on for the file system. \n  +   *DISABLED* - Turns automatic backups off for the file system.",
		//	      "enum": [
		//	        "DISABLED",
		//	        "ENABLED"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Status"
		//	  ],
		//	  "type": "object"
		//	}
		"backup_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Set the backup policy status for the file system.\n  +   *ENABLED* - Turns automatic backups on for the file system. \n  +   *DISABLED* - Turns automatic backups off for the file system.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"DISABLED",
							"ENABLED",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Use the ``BackupPolicy`` to turn automatic backups on or off for the file system.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BypassPolicyLockoutSafetyCheck
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "(Optional) A boolean that specifies whether or not to bypass the ``FileSystemPolicy`` lockout safety check. The lockout safety check determines whether the policy in the request will lock out, or prevent, the IAM principal that is making the request from making future ``PutFileSystemPolicy`` requests on this file system. Set ``BypassPolicyLockoutSafetyCheck`` to ``True`` only when you intend to prevent the IAM principal that is making the request from making subsequent ``PutFileSystemPolicy`` requests on this file system. The default value is ``False``.",
		//	  "type": "boolean"
		//	}
		"bypass_policy_lockout_safety_check": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "(Optional) A boolean that specifies whether or not to bypass the ``FileSystemPolicy`` lockout safety check. The lockout safety check determines whether the policy in the request will lock out, or prevent, the IAM principal that is making the request from making future ``PutFileSystemPolicy`` requests on this file system. Set ``BypassPolicyLockoutSafetyCheck`` to ``True`` only when you intend to prevent the IAM principal that is making the request from making subsequent ``PutFileSystemPolicy`` requests on this file system. The default value is ``False``.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// BypassPolicyLockoutSafetyCheck is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Encrypted
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A Boolean value that, if true, creates an encrypted file system. When creating an encrypted file system, you have the option of specifying a KmsKeyId for an existing kms-key-long. If you don't specify a kms-key, then the default kms-key for EFS, ``/aws/elasticfilesystem``, is used to protect the encrypted file system.",
		//	  "type": "boolean"
		//	}
		"encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A Boolean value that, if true, creates an encrypted file system. When creating an encrypted file system, you have the option of specifying a KmsKeyId for an existing kms-key-long. If you don't specify a kms-key, then the default kms-key for EFS, ``/aws/elasticfilesystem``, is used to protect the encrypted file system.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FileSystemId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"file_system_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FileSystemPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ``FileSystemPolicy`` for the EFS file system. A file system policy is an IAM resource policy used to control NFS access to an EFS file system. For more information, see [Using to control NFS access to Amazon EFS](https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html) in the *Amazon EFS User Guide*.",
		//	  "type": "object"
		//	}
		"file_system_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "The ``FileSystemPolicy`` for the EFS file system. A file system policy is an IAM resource policy used to control NFS access to an EFS file system. For more information, see [Using to control NFS access to Amazon EFS](https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html) in the *Amazon EFS User Guide*.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FileSystemProtection
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Describes the protection on the file system.",
		//	  "properties": {
		//	    "ReplicationOverwriteProtection": {
		//	      "description": "The status of the file system's replication overwrite protection.\n  +   ``ENABLED`` ? The file system cannot be used as the destination file system in a replication configuration. The file system is writeable. Replication overwrite protection is ``ENABLED`` by default. \n  +   ``DISABLED`` ? The file system can be used as the destination file system in a replication configuration. The file system is read-only and can only be modified by EFS replication.\n  +   ``REPLICATING`` ? The file system is being used as the destination file system in a replication configuration. The file system is read-only and is only modified only by EFS replication.\n  \n If the replication configuration is deleted, the file system's replication overwrite protection is re-enabled, the file system becomes writeable.",
		//	      "enum": [
		//	        "DISABLED",
		//	        "ENABLED"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"file_system_protection": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ReplicationOverwriteProtection
				"replication_overwrite_protection": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The status of the file system's replication overwrite protection.\n  +   ``ENABLED`` ? The file system cannot be used as the destination file system in a replication configuration. The file system is writeable. Replication overwrite protection is ``ENABLED`` by default. \n  +   ``DISABLED`` ? The file system can be used as the destination file system in a replication configuration. The file system is read-only and can only be modified by EFS replication.\n  +   ``REPLICATING`` ? The file system is being used as the destination file system in a replication configuration. The file system is read-only and is only modified only by EFS replication.\n  \n If the replication configuration is deleted, the file system's replication overwrite protection is re-enabled, the file system becomes writeable.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"DISABLED",
							"ENABLED",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Describes the protection on the file system.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FileSystemTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Use to create one or more tags associated with the file system. Each tag is a user-defined key-value pair. Name your file system on creation by including a ``\"Key\":\"Name\",\"Value\":\"{value}\"`` key-value pair. Each key must be unique. For more information, see [Tagging resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *General Reference Guide*.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A tag is a key-value pair attached to a file system. Allowed characters in the ``Key`` and ``Value`` properties are letters, white space, and numbers that can be represented in UTF-8, and the following characters:``+ - = . _ : /``",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag key (String). The key can't start with ``aws:``.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value of the tag key.",
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
		//	  "uniqueItems": true
		//	}
		"file_system_tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag key (String). The key can't start with ``aws:``.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value of the tag key.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Use to create one or more tags associated with the file system. Each tag is a user-defined key-value pair. Name your file system on creation by including a ``\"Key\":\"Name\",\"Value\":\"{value}\"`` key-value pair. Each key must be unique. For more information, see [Tagging resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *General Reference Guide*.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the kms-key-long to be used to protect the encrypted file system. This parameter is only required if you want to use a nondefault kms-key. If this parameter is not specified, the default kms-key for EFS is used. This ID can be in one of the following formats:\n  +  Key ID - A unique identifier of the key, for example ``1234abcd-12ab-34cd-56ef-1234567890ab``.\n  +  ARN - An Amazon Resource Name (ARN) for the key, for example ``arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab``.\n  +  Key alias - A previously created display name for a key, for example ``alias/projectKey1``.\n  +  Key alias ARN - An ARN for a key alias, for example ``arn:aws:kms:us-west-2:444455556666:alias/projectKey1``.\n  \n If ``KmsKeyId`` is specified, the ``Encrypted`` parameter must be set to true.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the kms-key-long to be used to protect the encrypted file system. This parameter is only required if you want to use a nondefault kms-key. If this parameter is not specified, the default kms-key for EFS is used. This ID can be in one of the following formats:\n  +  Key ID - A unique identifier of the key, for example ``1234abcd-12ab-34cd-56ef-1234567890ab``.\n  +  ARN - An Amazon Resource Name (ARN) for the key, for example ``arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab``.\n  +  Key alias - A previously created display name for a key, for example ``alias/projectKey1``.\n  +  Key alias ARN - An ARN for a key alias, for example ``arn:aws:kms:us-west-2:444455556666:alias/projectKey1``.\n  \n If ``KmsKeyId`` is specified, the ``Encrypted`` parameter must be set to true.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LifecyclePolicies
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of ``LifecyclePolicy`` objects that define the file system's ``LifecycleConfiguration`` object. A ``LifecycleConfiguration`` object informs Lifecycle management of the following:\n  +  When to move files in the file system from primary storage to IA storage.\n  + When to move files in the file system from primary storage or IA storage to Archive storage.\n +  When to move files that are in IA or Archive storage to primary storage.\n  \n  EFS requires that each ``LifecyclePolicy`` object have only a single transition. This means that in a request body, ``LifecyclePolicies`` needs to be structured as an array of ``LifecyclePolicy`` objects, one object for each transition, ``TransitionToIA``, ``TransitionToArchive`` ``TransitionToPrimaryStorageClass``. See the example requests in the following section for more information.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Describes a policy used by Lifecycle management that specifies when to transition files into and out of the EFS storage classes. For more information, see [Managing file system storage](https://docs.aws.amazon.com/efs/latest/ug/lifecycle-management-efs.html).\n  + Each ``LifecyclePolicy`` object can have only a single transition. This means that in a request body, ``LifecyclePolicies`` must be structured as an array of ``LifecyclePolicy`` objects, one object for each transition, ``TransitionToIA``, ``TransitionToArchive``, ``TransitionToPrimaryStorageClass``.\n + See the AWS::EFS::FileSystem examples for the correct ``LifecyclePolicy`` structure. Do not use the syntax shown on this page.",
		//	    "properties": {
		//	      "TransitionToArchive": {
		//	        "description": "The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Archive storage. Metadata operations such as listing the contents of a directory don't count as file access events.",
		//	        "type": "string"
		//	      },
		//	      "TransitionToIA": {
		//	        "description": "The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Infrequent Access (IA) storage. Metadata operations such as listing the contents of a directory don't count as file access events.",
		//	        "type": "string"
		//	      },
		//	      "TransitionToPrimaryStorageClass": {
		//	        "description": "Whether to move files back to primary (Standard) storage after they are accessed in IA or Archive storage. Metadata operations such as listing the contents of a directory don't count as file access events.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"lifecycle_policies": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: TransitionToArchive
					"transition_to_archive": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Archive storage. Metadata operations such as listing the contents of a directory don't count as file access events.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: TransitionToIA
					"transition_to_ia": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Infrequent Access (IA) storage. Metadata operations such as listing the contents of a directory don't count as file access events.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: TransitionToPrimaryStorageClass
					"transition_to_primary_storage_class": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Whether to move files back to primary (Standard) storage after they are accessed in IA or Archive storage. Metadata operations such as listing the contents of a directory don't count as file access events.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of ``LifecyclePolicy`` objects that define the file system's ``LifecycleConfiguration`` object. A ``LifecycleConfiguration`` object informs Lifecycle management of the following:\n  +  When to move files in the file system from primary storage to IA storage.\n  + When to move files in the file system from primary storage or IA storage to Archive storage.\n +  When to move files that are in IA or Archive storage to primary storage.\n  \n  EFS requires that each ``LifecyclePolicy`` object have only a single transition. This means that in a request body, ``LifecyclePolicies`` needs to be structured as an array of ``LifecyclePolicy`` objects, one object for each transition, ``TransitionToIA``, ``TransitionToArchive`` ``TransitionToPrimaryStorageClass``. See the example requests in the following section for more information.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PerformanceMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Performance mode of the file system. We recommend ``generalPurpose`` performance mode for all file systems. File systems using the ``maxIO`` performance mode can scale to higher levels of aggregate throughput and operations per second with a tradeoff of slightly higher latencies for most file operations. The performance mode can't be changed after the file system has been created. The ``maxIO`` mode is not supported on One Zone file systems.\n  Due to the higher per-operation latencies with Max I/O, we recommend using General Purpose performance mode for all file systems.\n  Default is ``generalPurpose``.",
		//	  "type": "string"
		//	}
		"performance_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Performance mode of the file system. We recommend ``generalPurpose`` performance mode for all file systems. File systems using the ``maxIO`` performance mode can scale to higher levels of aggregate throughput and operations per second with a tradeoff of slightly higher latencies for most file operations. The performance mode can't be changed after the file system has been created. The ``maxIO`` mode is not supported on One Zone file systems.\n  Due to the higher per-operation latencies with Max I/O, we recommend using General Purpose performance mode for all file systems.\n  Default is ``generalPurpose``.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ProvisionedThroughputInMibps
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The throughput, measured in mebibytes per second (MiBps), that you want to provision for a file system that you're creating. Required if ``ThroughputMode`` is set to ``provisioned``. Valid values are 1-3414 MiBps, with the upper limit depending on Region. To increase this limit, contact SUP. For more information, see [Amazon EFS quotas that you can increase](https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the *Amazon EFS User Guide*.",
		//	  "type": "number"
		//	}
		"provisioned_throughput_in_mibps": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "The throughput, measured in mebibytes per second (MiBps), that you want to provision for a file system that you're creating. Required if ``ThroughputMode`` is set to ``provisioned``. Valid values are 1-3414 MiBps, with the upper limit depending on Region. To increase this limit, contact SUP. For more information, see [Amazon EFS quotas that you can increase](https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the *Amazon EFS User Guide*.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				float64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReplicationConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Describes the replication configuration for a specific file system.",
		//	  "properties": {
		//	    "Destinations": {
		//	      "description": "An array of destination objects. Only one destination object is supported.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Describes the destination file system in the replication configuration.",
		//	        "properties": {
		//	          "AvailabilityZoneName": {
		//	            "description": "The AWS For One Zone file systems, the replication configuration must specify the Availability Zone in which the destination file system is located. \n Use the format ``us-east-1a`` to specify the Availability Zone. For more information about One Zone file systems, see [EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the *Amazon EFS User Guide*.\n  One Zone file system type is not available in all Availability Zones in AWS-Regions where Amazon EFS is available.",
		//	            "type": "string"
		//	          },
		//	          "FileSystemId": {
		//	            "description": "The ID of the destination Amazon EFS file system.",
		//	            "type": "string"
		//	          },
		//	          "KmsKeyId": {
		//	            "description": "The ID of an kms-key-long used to protect the encrypted file system.",
		//	            "type": "string"
		//	          },
		//	          "Region": {
		//	            "description": "The AWS-Region in which the destination file system is located.\n  For One Zone file systems, the replication configuration must specify the AWS-Region in which the destination file system is located.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "maxItems": 1,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"replication_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Destinations
				"destinations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AvailabilityZoneName
							"availability_zone_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The AWS For One Zone file systems, the replication configuration must specify the Availability Zone in which the destination file system is located. \n Use the format ``us-east-1a`` to specify the Availability Zone. For more information about One Zone file systems, see [EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the *Amazon EFS User Guide*.\n  One Zone file system type is not available in all Availability Zones in AWS-Regions where Amazon EFS is available.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: FileSystemId
							"file_system_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The ID of the destination Amazon EFS file system.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: KmsKeyId
							"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The ID of an kms-key-long used to protect the encrypted file system.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Region
							"region": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The AWS-Region in which the destination file system is located.\n  For One Zone file systems, the replication configuration must specify the AWS-Region in which the destination file system is located.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "An array of destination objects. Only one destination object is supported.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(1, 1),
						listvalidator.UniqueValues(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Describes the replication configuration for a specific file system.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ThroughputMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the throughput mode for the file system. The mode can be ``bursting``, ``provisioned``, or ``elastic``. If you set ``ThroughputMode`` to ``provisioned``, you must also set a value for ``ProvisionedThroughputInMibps``. After you create the file system, you can decrease your file system's Provisioned throughput or change between the throughput modes, with certain time restrictions. For more information, see [Specifying throughput with provisioned mode](https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput) in the *Amazon EFS User Guide*. \n Default is ``bursting``.",
		//	  "type": "string"
		//	}
		"throughput_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the throughput mode for the file system. The mode can be ``bursting``, ``provisioned``, or ``elastic``. If you set ``ThroughputMode`` to ``provisioned``, you must also set a value for ``ProvisionedThroughputInMibps``. After you create the file system, you can decrease your file system's Provisioned throughput or change between the throughput modes, with certain time restrictions. For more information, see [Specifying throughput with provisioned mode](https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput) in the *Amazon EFS User Guide*. \n Default is ``bursting``.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The ``AWS::EFS::FileSystem`` resource creates a new, empty file system in EFSlong (EFS). You must create a mount target ([AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html)) to mount your EFS file system on an EC2 or other AWS cloud compute resource.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EFS::FileSystem").WithTerraformTypeName("awscc_efs_file_system")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                 "Arn",
		"availability_zone_name":              "AvailabilityZoneName",
		"backup_policy":                       "BackupPolicy",
		"bypass_policy_lockout_safety_check":  "BypassPolicyLockoutSafetyCheck",
		"destinations":                        "Destinations",
		"encrypted":                           "Encrypted",
		"file_system_id":                      "FileSystemId",
		"file_system_policy":                  "FileSystemPolicy",
		"file_system_protection":              "FileSystemProtection",
		"file_system_tags":                    "FileSystemTags",
		"key":                                 "Key",
		"kms_key_id":                          "KmsKeyId",
		"lifecycle_policies":                  "LifecyclePolicies",
		"performance_mode":                    "PerformanceMode",
		"provisioned_throughput_in_mibps":     "ProvisionedThroughputInMibps",
		"region":                              "Region",
		"replication_configuration":           "ReplicationConfiguration",
		"replication_overwrite_protection":    "ReplicationOverwriteProtection",
		"status":                              "Status",
		"throughput_mode":                     "ThroughputMode",
		"transition_to_archive":               "TransitionToArchive",
		"transition_to_ia":                    "TransitionToIA",
		"transition_to_primary_storage_class": "TransitionToPrimaryStorageClass",
		"value":                               "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/BypassPolicyLockoutSafetyCheck",
		"/properties/ReplicationConfiguration/Destinations/0/AvailabilityZoneName",
		"/properties/ReplicationConfiguration/Destinations/0/KmsKeyId",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
