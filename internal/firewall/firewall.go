package firewall

import (
	"context"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func SetupFirewall(ctx context.Context, client *hcloud.Client, server *hcloud.Server) error {
	// Get Cloudflare IPs
	cloudflareIPs, err := getCloudflareIPs()
	if err != nil {
		return err
	}

	// Create firewall rules
	rules := []hcloud.FirewallRule{
		newFirewallRule("80", cloudflareIPs),
		newFirewallRule("443", cloudflareIPs),
		newFirewallRule("22", []string{"0.0.0.0/0", "::/0"}),
	}

	firewallOpts := hcloud.FirewallCreateOpts{
		Name:  "cloudflare-firewall",
		Rules: rules,
		ApplyTo: []hcloud.FirewallResource{
			{
				Type: hcloud.FirewallResourceTypeServer,
				Server: &hcloud.FirewallResourceServer{
					ID: server.ID,
				},
			},
		},
	}

	_, _, err = client.Firewall.Create(ctx, firewallOpts)
	return err
}

func getCloudflareIPs() ([]string, error) {
	var ips []string

	// Get IPv4 ranges
	resp, err := http.Get("https://www.cloudflare.com/ips-v4")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ipv4, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Get IPv6 ranges
	resp, err = http.Get("https://www.cloudflare.com/ips-v6")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ipv6, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Combine and split IPs
	ips = append(
		strings.Split(strings.TrimSpace(string(ipv4)), "\n"),
		strings.Split(strings.TrimSpace(string(ipv6)), "\n")...,
	)

	return ips, nil
}

func newFirewallRule(port string, sourceIPs []string) hcloud.FirewallRule {
	return hcloud.FirewallRule{
		Direction: hcloud.FirewallRuleDirectionIn,
		Protocol:  hcloud.FirewallRuleProtocolTCP,
		Port:      &port,
		SourceIPs: convertToIPNets(sourceIPs),
	}
}

func convertToIPNets(ips []string) []net.IPNet {
	var ipNets []net.IPNet
	for _, ip := range ips {
		_, ipNet, err := net.ParseCIDR(ip)
		if err != nil {
			ipNet = &net.IPNet{
				IP:   net.ParseIP(ip),
				Mask: net.CIDRMask(32, 32),
			}
		}
		if ipNet != nil {
			ipNets = append(ipNets, *ipNet)
		}
	}
	return ipNets
}
