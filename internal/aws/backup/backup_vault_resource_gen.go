// Code generated by generators/resource/main.go; DO NOT EDIT.

package backup

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_backup_backup_vault", backupVaultResourceType)
}

// backupVaultResourceType returns the Terraform awscc_backup_backup_vault resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Backup::BackupVault resource type.
func backupVaultResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"access_policy": {
			// Property: AccessPolicy
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"backup_vault_arn": {
			// Property: BackupVaultArn
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
		"backup_vault_name": {
			// Property: BackupVaultName
			// CloudFormation resource type schema:
			// {
			//   "pattern": "^[a-zA-Z0-9\\-\\_]{2,50}$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9\\-\\_]{2,50}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"backup_vault_tags": {
			// Property: BackupVaultTags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"encryption_key_arn": {
			// Property: EncryptionKeyArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"lock_configuration": {
			// Property: LockConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ChangeableForDays": {
			//       "type": "number"
			//     },
			//     "MaxRetentionDays": {
			//       "type": "number"
			//     },
			//     "MinRetentionDays": {
			//       "type": "number"
			//     }
			//   },
			//   "required": [
			//     "MinRetentionDays"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"changeable_for_days": {
						// Property: ChangeableForDays
						Type:     types.Float64Type,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"max_retention_days": {
						// Property: MaxRetentionDays
						Type:     types.Float64Type,
						Optional: true,
						Computed: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"min_retention_days": {
						// Property: MinRetentionDays
						Type:     types.Float64Type,
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"notifications": {
			// Property: Notifications
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "BackupVaultEvents": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "SNSTopicArn": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "SNSTopicArn",
			//     "BackupVaultEvents"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"backup_vault_events": {
						// Property: BackupVaultEvents
						Type:     types.ListType{ElemType: types.StringType},
						Required: true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							Multiset(),
						},
					},
					"sns_topic_arn": {
						// Property: SNSTopicArn
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
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
		Description: "Resource Type definition for AWS::Backup::BackupVault",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Backup::BackupVault").WithTerraformTypeName("awscc_backup_backup_vault")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_policy":       "AccessPolicy",
		"backup_vault_arn":    "BackupVaultArn",
		"backup_vault_events": "BackupVaultEvents",
		"backup_vault_name":   "BackupVaultName",
		"backup_vault_tags":   "BackupVaultTags",
		"changeable_for_days": "ChangeableForDays",
		"encryption_key_arn":  "EncryptionKeyArn",
		"lock_configuration":  "LockConfiguration",
		"max_retention_days":  "MaxRetentionDays",
		"min_retention_days":  "MinRetentionDays",
		"notifications":       "Notifications",
		"sns_topic_arn":       "SNSTopicArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
