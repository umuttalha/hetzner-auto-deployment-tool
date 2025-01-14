package dns

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
)

func SetupDNSRecords(ctx context.Context, client *cloudflare.API, zoneID, domainName, serverIP string) error {
	// Create resource container
	rc := &cloudflare.ResourceContainer{
		Identifier: zoneID,
		Level:      cloudflare.ZoneRouteLevel,
	}

	// Create A record
	proxied := true
	aRecordParams := cloudflare.CreateDNSRecordParams{
		Type:    "A",
		Name:    domainName,
		Content: serverIP,
		Proxied: &proxied,
	}

	_, err := client.CreateDNSRecord(ctx, rc, aRecordParams)
	if err != nil {
		return err
	}

	// Create CNAME record
	cnameRecordParams := cloudflare.CreateDNSRecordParams{
		Type:    "CNAME",
		Name:    "www",
		Content: domainName,
		Proxied: &proxied,
	}

	_, err = client.CreateDNSRecord(ctx, rc, cnameRecordParams)
	return err
}
