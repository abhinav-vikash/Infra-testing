package test

import (
	"fmt"
	// "strings"
	"encoding/json"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	// "github.com/davecgh/go-spew/spew"
	// "github.com/gruntwork-io/terratest/modules/random"
	// "github.com/gruntwork-io/terratest/modules/terraform"
	// "github.com/stretchr/testify/assert"
)

func TestStorageFunction(t *testing.T) {
	t.Parallel()
	
	resourceGroupName := "rg_ada_adls"
	// storageV2ContainerName := "stagingadls"
	storageAccountName := "adaadls"
	subscriptionID := "3344a922-f246-4f27-a6f1-3c85586f7b99"	
	account,er := azure.GetStorageAccountPropertyE(storageAccountName, resourceGroupName, subscriptionID)
	result,rerr := json.Marshal(account)
	fmt.Println(string(result))
	fmt.Println((account))
	fmt.Println(er)
	fmt.Println(rerr)
}