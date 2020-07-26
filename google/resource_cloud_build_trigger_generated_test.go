// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCloudBuildTrigger_cloudbuildTriggerFilenameExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudBuildTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudBuildTrigger_cloudbuildTriggerFilenameExample(context),
			},
			{
				ResourceName:      "google_cloudbuild_trigger.filename-trigger",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudBuildTrigger_cloudbuildTriggerFilenameExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudbuild_trigger" "filename-trigger" {
  trigger_template {
    branch_name = "master"
    repo_name   = "my-repo"
  }

  substitutions = {
    _FOO = "bar"
    _BAZ = "qux"
  }

  filename = "cloudbuild.yaml"
}
`, context)
}

func TestAccCloudBuildTrigger_cloudbuildTriggerBuildExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCloudBuildTriggerDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccCloudBuildTrigger_cloudbuildTriggerBuildExample(context),
			},
			{
				ResourceName:      "google_cloudbuild_trigger.build-trigger",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudBuildTrigger_cloudbuildTriggerBuildExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudbuild_trigger" "build-trigger" {
  trigger_template {
    branch_name = "master"
    repo_name   = "my-repo"
  }
  
  build {
    step {
      name = "gcr.io/cloud-builders/gsutil"
      args = ["cp", "gs://mybucket/remotefile.zip", "localfile.zip"]
      timeout = "120s"
    }

    source {
      storage_source {
        bucket = "mybucket"
        object = "source_code.tar.gz"
      }
    }
    tags = ["build", "newFeature"]
    substitutions = {
      _FOO = "bar"
      _BAZ = "qux"
    }
    queue_ttl = "20s"
    logs_bucket = "gs://mybucket/logs"
    secret {
      kms_key_name = "projects/myProject/locations/global/keyRings/keyring-name/cryptoKeys/key-name"
      secret_env = {
        PASSWORD = "ZW5jcnlwdGVkLXBhc3N3b3JkCg=="
      }
    }
  }  
}
`, context)
}

func testAccCheckCloudBuildTriggerDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_cloudbuild_trigger" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{CloudBuildBasePath}}projects/{{project}}/triggers/{{trigger_id}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("CloudBuildTrigger still exists at %s", url)
			}
		}

		return nil
	}
}
