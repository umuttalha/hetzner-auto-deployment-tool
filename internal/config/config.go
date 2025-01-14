package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	HetznerToken     string
	CloudflareToken  string
	CloudflareZoneID string
	DomainName       string
	SSHKeyName       string
	BackendRepoURL   string
	ServerType       string
	ServerImage      string
	ServerLocation   string
}

func ShowHelp() {
	fmt.Println("Server Deployment Tool")
	fmt.Println("\nAvailable Server Types:")
	fmt.Println("  - cx22  (2 vCPU, 4GB RAM)")
	fmt.Println("  - cx32  (4 vCPU, 8GB RAM)")
	fmt.Println("  - cx42  (8 vCPU, 16GB RAM)")
	fmt.Println("  - cx52  (16 vCPU, 32GB RAM)")

	fmt.Println("\nAvailable Locations:")
	fmt.Println("  - nbg1  (Nuremberg, Germany)")
	fmt.Println("  - fsn1  (Falkenstein, Germany)")
	fmt.Println("  - hel1  (Helsinki, Finland)")

	fmt.Println("\nAvailable Images:")
	fmt.Println("  - ubuntu-24.04")
	fmt.Println("  - ubuntu-22.04")
	fmt.Println("  - debian-11")
	fmt.Println("  - debian-12")

	fmt.Println("\nFlags:")
	flag.PrintDefaults()
}

func (c *Config) LoadFromEnv() {
	c.HetznerToken = os.Getenv("HETZNER_API_KEY")
	c.CloudflareToken = os.Getenv("CLOUDFLARE_API_TOKEN")
	c.CloudflareZoneID = os.Getenv("CLOUDFLARE_ZONE_ID")
	c.DomainName = os.Getenv("DOMAIN_NAME")
	c.SSHKeyName = os.Getenv("SSH_KEY_NAME")
}

func (c *Config) Validate() error {
	if c.HetznerToken == "" {
		return fmt.Errorf("HETZNER_API_KEY is required")
	}
	if c.CloudflareToken == "" {
		return fmt.Errorf("CLOUDFLARE_API_TOKEN is required")
	}
	if c.CloudflareZoneID == "" {
		return fmt.Errorf("CLOUDFLARE_ZONE_ID is required")
	}
	if c.DomainName == "" {
		return fmt.Errorf("DOMAIN_NAME is required")
	}
	if c.SSHKeyName == "" {
		return fmt.Errorf("SSH_KEY_NAME is required")
	}
	return nil
}
