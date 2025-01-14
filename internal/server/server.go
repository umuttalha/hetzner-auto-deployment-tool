package server

import (
	"context"
	"fmt"

	"github.com/umuttalha/go-cli-tool/internal/config"
	"github.com/umuttalha/go-cli-tool/internal/dns"
	"github.com/umuttalha/go-cli-tool/internal/firewall"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	"github.com/umuttalha/go-cli-tool/internal/scripts"
)

func SetupInfrastructure(config config.Config) error {
	ctx := context.Background()

	// Initialize Hetzner client
	hetznerClient := hcloud.NewClient(hcloud.WithToken(config.HetznerToken))

	// Initialize Cloudflare client
	cloudflareClient, err := cloudflare.NewWithAPIToken(config.CloudflareToken)
	if err != nil {
		return fmt.Errorf("failed to create cloudflare client: %v", err)
	}

	// Create server
	server, err := createServer(ctx, hetznerClient, config)
	if err != nil {
		return fmt.Errorf("failed to create server: %v", err)
	}

	// Setup DNS records
	if err := dns.SetupDNSRecords(ctx, cloudflareClient, config.CloudflareZoneID, config.DomainName, server.PublicNet.IPv4.IP.String()); err != nil {
		return fmt.Errorf("failed to setup DNS records: %v", err)
	}

	// Setup firewall
	if err := firewall.SetupFirewall(ctx, hetznerClient, server); err != nil {
		return fmt.Errorf("failed to setup firewall: %v", err)
	}

	return nil
}

func createServer(ctx context.Context, client *hcloud.Client, config config.Config) (*hcloud.Server, error) {
	// Get SSH key
	sshKey, _, err := client.SSHKey.Get(ctx, config.SSHKeyName)
	if err != nil {
		return nil, err
	}

	// Get setup script
	userData := scripts.GetServerSetupScript(config.BackendRepoURL)

	serverCreateOpts := hcloud.ServerCreateOpts{
		Name:       "web-server",
		ServerType: &hcloud.ServerType{Name: config.ServerType},
		Image:      &hcloud.Image{Name: config.ServerImage},
		Location:   &hcloud.Location{Name: config.ServerLocation},
		SSHKeys:    []*hcloud.SSHKey{sshKey},
		UserData:   userData,
	}

	result, _, err := client.Server.Create(ctx, serverCreateOpts)
	if err != nil {
		return nil, err
	}

	return result.Server, nil
}
