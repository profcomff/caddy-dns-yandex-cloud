package caddy_dns_yandex_cloud

import (
    "github.com/caddyserver/caddy/v2"
    "github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
    libdns "github.com/profcomff/libdns-yandex-cloud"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *libdns.Provider }

func init() {
    caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
    return caddy.ModuleInfo{
        ID:  "dns.providers.yandex_cloud",
        New: func() caddy.Module { return &Provider{new(libdns.Provider)} },
    }
}

// Before using the provider config, resolve placeholders in the API token.
// Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
    repl := caddy.NewReplacer()
    p.Provider.ServiceAccountConfigPath = repl.ReplaceAll(p.Provider.ServiceAccountConfigPath, "")
    return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// libdns [<service_account_config_path>] {
//     service_account_config_path <json config path> ~/.yc/config.json for example
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
    for d.Next() {
        if d.NextArg() {
            p.Provider.ServiceAccountConfigPath = d.Val()
        }
        if d.NextArg() {
            return d.ArgErr()
        }
        for nesting := d.Nesting(); d.NextBlock(nesting); {
            switch d.Val() {
            case "service_account_config_path":
                err := p.Provider.SetServiceConfig(d.Val())
                if err != nil{
                    return err
                }
                if p.Provider.ServiceAccountConfigPath != "" {
                    return d.Err("config_path already set")
                }
                p.Provider.ServiceAccountConfigPath = d.Val()
                if d.NextArg() {
                    return d.ArgErr()
                }
            default:
                return d.Errf("unrecognized subdirective '%s'", d.Val())
            }
        }
    }
    if p.Provider.ServiceAccountConfigPath == "" {
        return d.Err("missing service account config path")
    }
    return nil
}

// Interface guards
var (
    _ caddyfile.Unmarshaler = (*Provider)(nil)
    _ caddy.Provisioner     = (*Provider)(nil)
)
