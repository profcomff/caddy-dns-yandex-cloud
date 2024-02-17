yandex_cloud module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with yandex_cloud.

## Caddy module name

```
dns.providers.provider_name
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "provider_name",
				"service_account_config_path": "~/.yc/keys.json"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns provider_name ...
}
```

```
# one site
tls {
	dns provider_name ...
}
```
