package test

import (
	"log"
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)


func getPrivateDNSZoneMetadata(subscriptionId string, resourceGroupName string, zoneName string)(armprivatedns.PrivateZonesClientGetResponse, error){
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	privateZoneClient, err := armprivatedns.NewPrivateZonesClient(subscriptionId, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	dnsZone, err := privateZoneClient.Get(ctx,
		resourceGroupName,
		zoneName,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	return dnsZone, err
}

func getRecordSetZoneMetadata(subscriptionId string, resourceGroupName string, zoneName string, recordName string)(armprivatedns.RecordSetsClientGetResponse, error){
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	recordSetClient, err := armprivatedns.NewRecordSetsClient(subscriptionId, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	
	recordSet, err := recordSetClient.Get(ctx,
		resourceGroupName,
		zoneName,
		armprivatedns.RecordTypeA,
		recordName,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	return recordSet, err
}