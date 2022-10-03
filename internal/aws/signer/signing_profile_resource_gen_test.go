// Code generated by generators/resource/main.go; DO NOT EDIT.

package signer_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSSignerSigningProfile_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Signer::SigningProfile", "awscc_signer_signing_profile", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
