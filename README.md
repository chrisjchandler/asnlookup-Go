# ASN Lookup Tool & RIPE Interrogator

An interactive command-line client for [RIPEstat](https://stat.ripe.net/). It queries routing, registration, address-space, DNS, and RPKI data for ASNs, IP addresses, and prefixes.

Run the program and enter the number for the lookup you need. Results are printed as formatted JSON and, after every request, also saved in the current directory as a timestamped JSON file (for example, `fetchASOverview_20260730_113000.json`). Saved files wrap the RIPEstat response with the lookup name and run time.

## Input conventions

- **ASNs:** Enter `12345` or `AS12345`; numeric ASNs are normalized to the `AS12345` form.
- **IP addresses:** IPv4 and IPv6 addresses are accepted where an IP is requested.
- **Prefixes/CIDRs:** Use a network such as `192.0.2.0/24` or `2001:db8::/32`.
- **IP-or-prefix prompts:** For the options noted below, an IP address is first resolved to its covering BGP prefix. Supplying a prefix skips that resolution.
- **Times:** BGP updates and allocation-history lookups accept `YYYY-MM-DD` or RFC 3339 timestamps such as `2026-07-30T12:00:00Z`.

## Menu reference

| # | Lookup | Input | What it returns |
| --- | --- | --- | --- |
| 1 | Neighbor ASNs | ASN | ASNs observed as neighbors of the selected ASN, including relationship type, peer counts, and power. |
| 2 | Historical neighbor ASNs | ASN | Neighbor relationships over time, including their observed date ranges. |
| 3 | Abuse contacts | IP, prefix, or ASN | Registered abuse-contact email addresses for the resource. |
| 4 | Historical WHOIS change count | IP, prefix, or ASN | The number of available historical WHOIS record versions. |
| 5 | Routing history | ASN, prefix, or IP | Historical routing and visibility information for the resource. |
| 6 | Announced prefixes | ASN | Prefixes announced by the ASN. |
| 7 | BGP updates | ASN; optional start/end time | BGP announcements and withdrawals for the ASN in the requested time window. A blank start is one hour ago; a blank end is now. |
| 8 | Geolocation | IP | Geolocation data associated with the IP address. |
| 9 | Reverse DNS | IP | PTR/reverse-DNS names for the IP address. |
| 10 | Network information | IP | The IP's covering prefix and related network information. |
| 11 | DNS blocklists | IP | Whether the IP appears in supported DNS-based blocklists. |
| 12 | Address-space hierarchy | Prefix or IP | The allocation hierarchy containing the prefix. An IP is resolved to its covering prefix first. |
| 13 | AS path length | ASN | AS-path-length distribution and related path information for the ASN. |
| 14 | Address-space usage | Prefix or IP | Address-usage information for the prefix. An IP is resolved to its covering prefix first. |
| 15 | Routing status by IP | IP | Current route/origin information for the IP address. |
| 16 | AS routing consistency | ASN | Whether routes and prefixes associated with the ASN are internally consistent. |
| 17 | Routing status by ASN | ASN | Current routing status and origin information for the ASN. |
| 18 | Prefix routing consistency by IP | IP | Routing consistency for the IP's covering prefix. |
| 19 | Routing status by prefix | Prefix | Current route/origin information for the exact prefix. |
| 20 | Prefix routing consistency by prefix | Prefix | Routing consistency checks for the exact prefix. |
| 21 | Routing status by CIDR | CIDR | Current routing status for an IP block; functionally the CIDR-specific form of option 19. |
| 22 | Prefix routing consistency by CIDR | CIDR | Routing consistency for an IP block; functionally the CIDR-specific form of option 20. |
| 23 | IP/prefix allocation history | IP or prefix; optional start/end time | Allocation and registration history. A blank start is one year ago; a blank end is now. |
| 24 | ASN allocation history | ASN; optional start/end time | Allocation and registration history for the ASN. A blank start is one year ago; a blank end is now. |
| 25 | Derive origin ASN(s) | IP or prefix | The originating ASN or ASNs inferred from current routing status. An IP is resolved to its covering prefix first. |
| 26 | AS overview | ASN | General AS metadata and overview information. |
| 27 | Prefix overview | Prefix or IP | General prefix metadata and overview information. An IP is resolved to its covering prefix first. |
| 28 | WHOIS | IP, prefix, or ASN | Current WHOIS/registry records for the resource. |
| 29 | RPKI validation | ASN and prefix or IP | RPKI validation for the ASN-prefix pair. An IP is resolved to its covering prefix first. |
| 30 | Visibility | IP, prefix, or ASN | Visibility measurements and related routing data for the resource. |
| 31 | Reverse-DNS consistency | ASN or prefix | Whether reverse-DNS data is consistent with the ASN or prefix. |
| 32 | DNS chain | Hostname or IP | DNS resolution chain and DNSSEC-related information for a hostname or IP address. |
| 33 | Reverse-DNS delegations | Prefix | Reverse-DNS delegations associated with the prefix. |
| 34 | Exit | — | Closes the program. |

## Notes

- The tool is an interactive client: it has no command-line flags or batch mode.
- Data availability and field meanings are determined by RIPEstat. An empty or partial response can mean that RIPEstat has no data for the requested resource.
- Requests time out after 20 seconds. HTTP errors are printed to the terminal; the tool then returns to the menu.
