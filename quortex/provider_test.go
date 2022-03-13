package quortex

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"quortex": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func TestAccPreCheck(t *testing.T) {
	if err := os.Getenv("QUORTEX_USERNAME"); err == "" {
		t.Fatal("QUORTEX_USERNAME must be set for acceptance tests")
	}
	if err := os.Getenv("QUORTEX_PASSWORD"); err == "" {
		t.Fatal("QUORTEX_PASSWORD must be set for acceptance tests")
	}
}
