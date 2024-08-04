# Yandex Cloud module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Yandex Cloud accounts.

## Caddy module name

```
dns.providers.yandex_cloud
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "yandex_cloud",
        "api_token": "YCLOUD_KEYS_FILE"
      }
    }
  }
}
```

or with the Caddyfile:

```
your.domain.com {
  respond "Hello World"	# replace with whatever config you need...
  tls {
    dns yandex_cloud {env.YCLOUD_KEYS_FILE}
  }
}
```

You can replace `{env.YCLOUD_KEYS_FILE}` with the actual authorized keys file if you prefer to put it directly in your config instead of an environment variable.

## Authenticating

See [the associated README in the libdns package](https://github.com/github.com/profcomff/libdns-yandex-cloud) for important information about credentials.
