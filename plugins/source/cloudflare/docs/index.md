## Cloudflare Provider

The CloudQuery Cloudflare provider pulls configuration out of Cloudflare resources, normalizes them and stores them in PostgreSQL database.

### Install

```shell
cloudquery init cloudflare
```

### Configuration

The following configuration section can be automatically generated by `cloudquery init cloudflare`:

```yaml
// Use can use either the API token or the API key
// API token is preferred

// API token to access Cloudflare resources, also can be set with the CLOUDFLARE_API_TOKEN environment variable
api_token: "<YOUR_API_TOKEN_HERE>"
// API key to access Cloudflare resources, also can be set with the CLOUDFLARE_API_KEY environment variable
api_key: "<YOUR_API_KEY_HERE>"
// API email to access Cloudflare resources, also can be set with the CLOUDFLARE_API_EMAIL environment variable
api_email: "<YOUR_API_EMAIL_HERE>"

// List of accounts to target, if empty, all accounts will be targeted
//accounts:
// - "<YOUR_ACCOUNT_ID>"

// List of accounts to target, if empty, all available zones will be targeted
//zones:
// - "<YOUR_ZONE_ID>"
```

### Environment variables

The following environment variables can be used instead of passing in them in the configuration:

1. `CLOUDFLARE_API_KEY` & `CLOUDFLARE_EMAIL` 
2. `CLOUDFLARE_API_TOKEN`

## Query Examples

### Find all zones with dev mode enabled

```sql
SELECT id, account_id, host_name, name, original_ns FROM cloudflare_zones WHERE dev_mode = true;
```

### Find all dns records

```sql
SELECT id, account_id, zone_id, name, type FROM cloudflare_dns_records;
```