# Cloud Server Deployment Tool

A CLI tool for automated server deployment on Hetzner Cloud with Cloudflare DNS integration. This tool simplifies the process of creating servers, configuring DNS records, and setting up basic security with firewall rules.

## Features

- 🚀 One-command server deployment on Hetzner Cloud
- 🔒 Automatic Cloudflare DNS configuration
- 🛡️ Firewall setup with Cloudflare IP allowlist
- 🔑 SSH key integration
- 🐳 Automatic Docker installation
- 📦 Caddy server setup
- ⚙️ Customizable server configurations

## Prerequisites

- Hetzner Cloud account and API token
- Cloudflare account and API token
- Domain name managed by Cloudflare
- SSH key uploaded to Hetzner Cloud
- Go 1.19 or later

## Installation

```bash
# Clone the repository
git clone https://github.com/umuttalha/go-cli-tool
cd go-cli-tool

# Build the tool
cd cmd
go build -o deploy-tool
```

## Configuration

Create a `.env` file in the project root with the following variables:

```env
HETZNER_API_KEY=your_hetzner_api_token
CLOUDFLARE_API_TOKEN=your_cloudflare_api_token
CLOUDFLARE_ZONE_ID=your_zone_id
DOMAIN_NAME=your_domain_name
SSH_KEY_NAME=your_ssh_key_name
```

## Usage

### Basic Usage

```bash
# Show help and available options
./deploy-tool -help

# Deploy with default settings
./deploy-tool -repo-url https://github.com/your/backend-repo.git

# Deploy with custom configuration
./deploy-tool \
  -server-type cx22 \
  -server-image ubuntu-24.04 \
  -location fsn1 \
  -repo-url https://github.com/your/backend-repo.git
```

### Available Options

| Flag | Description | Default |
|------|-------------|---------|
| `-server-type` | Hetzner server type | `cx22` |
| `-server-image` | Server operating system | `ubuntu-24.04` |
| `-location` | Server location | `nbg1` |
| `-repo-url` | Backend repository URL | - |
| `-help` | Show help message | `false` |

### Server Types


- `cx22`: 2 vCPU, 4GB RAM
- `cx32`: 4 vCPU, 8GB RAM
- `cx42`: 8 vCPU, 16GB RAM

### Locations

- `nbg1`: Nuremberg, Germany
- `fsn1`: Falkenstein, Germany
- `hel1`: Helsinki, Finland

### Available Images

- `ubuntu-24.04`
- `ubuntu-22.04`
- `ubuntu-20.04`
- `debian-11`
- `debian-12`

## What Gets Deployed

The tool sets up:

1. A new server on Hetzner Cloud
2. DNS records on Cloudflare (A and CNAME records)
3. Firewall rules allowing:
   - HTTP (80) and HTTPS (443) from Cloudflare IPs only
   - SSH (22) from any IP
4. Basic software:
   - Docker and Docker Compose
   - Caddy web server
   - Git
   - Your specified backend application

## Security Features

- 🔒 Automatic firewall configuration
- 🛡️ Cloudflare proxy protection
- 🔐 SSH key authentication
- 🌐 HTTPS/TLS support via Caddy

## Development

### Requirements

```bash
go mod download
```

### Building

```bash
go build -o deploy-tool
```

## Acknowledgments

- [Hetzner Cloud Go Client](https://github.com/hetznercloud/hcloud-go)
- [Cloudflare Go Client](https://github.com/cloudflare/cloudflare-go)
- [GoDotEnv](https://github.com/joho/godotenv)


