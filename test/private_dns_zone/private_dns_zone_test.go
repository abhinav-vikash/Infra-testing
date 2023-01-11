
package private_dns_zone

import (
	"fmt"
	"testing"
	"strings"

	test "github.com/abhinav-vikash/Infra-testing/test"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestPrivateDNSzoneFunction(t *testing.T){
	t.Parallel()

	subscriptionId := "3344a922-f246-4f27-a6f1-3c85586f7b99"
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../private_dns_zone",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			// "postfix": strings.ToLower(uniquePostfix),
			"dns_resource_group_name": "privatednsrg",
			"dns_location": "eastus",
			"domain_name": "adaprivatezone.com",
			"record_name": "ada-backstage",
		},
	}
	
	terraform.InitAndApply(t, terraformOptions)

	output_domain_name := terraform.Output(t, terraformOptions, "domain_name")
	output_record_name := terraform.Output(t, terraformOptions, "record_name")
	a_record_ip := terraform.Output(t, terraformOptions, "a_record_ip")
	fully_qualified_domain_name := terraform.Output(t, terraformOptions, "fully_qualified_domain_name")
	
	dnsZone, err := test.getPrivateDNSZoneMetadata(subscriptionId, "privatednsrg", "adaprivatezone.com")
	recordSet, err := test.getRecordSetZoneMetadata(subscriptionId, "privatednsrg", "adaprivatezone.com", "ada-backstage")
	recordSetProperties := *recordSet.Properties
	record_ips := []string{}
	for _, b := range recordSetProperties.ARecords {
		record_ips = append(record_ips, *b.IPv4Address)
	}

	fmt.Println("checking domain name")
	assert.Equal(t, *dnsZone.PrivateZone.Name, output_domain_name, "domain name mismatch")

	fmt.Println("checking record name")
	assert.Equal(t, *recordSet.Name, output_record_name, "record name mismatch")

	fmt.Println("checking record set presence and ips")
	assert.NotEmpty(t, *recordSet.ID, "record not created mismatch")
	assert.Equal(t, "["+strings.Join(record_ips, " ")+"]", a_record_ip, "record ip mismatch")

	fmt.Println("checking fully qualified domain name")
	assert.Equal(t, *recordSetProperties.Fqdn, fully_qualified_domain_name, "fully_qualified_domain_name mismatch")

	fmt.Println("testing completed")
	defer terraform.Destroy(t, terraformOptions)

}