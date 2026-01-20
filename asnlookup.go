package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const ripeBase = "https://stat.ripe.net/data"

var httpClient = &http.Client{
	Timeout: 20 * time.Second,
}

type asnNeighboursResp struct {
	Data struct {
		Neighbours []struct {
			ASN     int    `json:"asn"`
			Type    string `json:"type"`
			Power   int    `json:"power"`
			V4Peers int    `json:"v4_peers"`
			V6Peers int    `json:"v6_peers"`
		} `json:"neighbours"`
	} `json:"data"`
}

type asnNeighboursHistoryResp struct {
	Data struct {
		Neighbours []struct {
			ASN       int    `json:"asn"`
			Type      string `json:"type"`
			Power     int    `json:"power"`
			V4Peers   int    `json:"v4_peers"`
			V6Peers   int    `json:"v6_peers"`
			Timelines []struct {
				Starttime string `json:"starttime"`
				Endtime   string `json:"endtime"`
			} `json:"timelines"`
		} `json:"neighbours"`
	} `json:"data"`
}

type abuseContactResp struct {
	Data struct {
		AbuseContacts []string `json:"abuse_contacts"`
	} `json:"data"`
}

type historicalWhoisResp struct {
	Data struct {
		NumVersions int `json:"num_versions"`
	} `json:"data"`
}

type networkInfoResp struct {
	Data struct {
		Prefix string `json:"prefix"`
	} `json:"data"`
}

func main() {
	in := bufio.NewReader(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("Choose an option:")
		fmt.Println("1. Fetch neighbor ASNs for a given ASN")
		fmt.Println("2. Fetch historical neighbor ASNs for a given ASN")
		fmt.Println("3. Get Abuse contact information for an IP/Prefix/ASN")
		fmt.Println("4. Get Historical whois change count for an IP/Prefix/ASN")
		fmt.Println("5. Fetch routing history for a given ASN/Prefix/IP")
		fmt.Println("6. Fetch prefix information for a given ASN")
		fmt.Println("7. Fetch BGP updates for a given ASN")
		fmt.Println("8. Fetch geolocation information for a given IP address")
		fmt.Println("9. Fetch reverse DNS information for a given IP address")
		fmt.Println("10. Fetch network information for a given IP address")
		fmt.Println("11. Fetch blacklist information for a given IP address")
		fmt.Println("12. Fetch IP address space information for a given ASN")
		fmt.Println("13. Fetch AS path information for a given ASN")
		fmt.Println("14. Fetch IP address block information for a given IP address")
		fmt.Println("15. Fetch routing status for a given IP address")
		fmt.Println("16. Fetch routing consistency for a given ASN")
		fmt.Println("17. Fetch routing status for a given ASN")
		fmt.Println("18. Fetch routing consistency for a given IP address")
		fmt.Println("19. Fetch routing status for a given prefix")
		fmt.Println("20. Fetch routing consistency for a given prefix")
		fmt.Println("21. Fetch routing status for a given IP address block (CIDR)")
		fmt.Println("22. Fetch routing consistency for a given IP address block (CIDR)")
		fmt.Println("23. Fetch IP address history for a given IP address")
		fmt.Println("24. Fetch ASN history for a given ASN")
		fmt.Println("25. Exit")
		fmt.Println("26. Derive originating ASN(s) from an IP or prefix (CIDR)")

		choice, ok := readInt(in, "Enter choice (1-26): ")
		if !ok {
			return
		}

		switch choice {
		case 1:
			fetchNeighborASNs(in)
		case 2:
			fetchHistoricalNeighborASNs(in)
		case 3:
			getAbuseContactInfo(in)
		case 4:
			getHistoricalWhoisChangeCount(in)
		case 5:
			fetchRoutingHistory(in)
		case 6:
			fetchPrefixInfo(in)
		case 7:
			fetchBGPUpdates(in)
		case 8:
			fetchGeolocationInfo(in)
		case 9:
			fetchReverseDNSInfo(in)
		case 10:
			fetchNetworkInfo(in)
		case 11:
			fetchBlacklistInfo(in)
		case 12:
			fetchIPAddressSpaceInfo(in)
		case 13:
			fetchASPathInfo(in)
		case 14:
			fetchIPAddressBlockInfo(in)
		case 15:
			fetchRoutingStatusByIP(in)
		case 16:
			fetchRoutingConsistencyByASN(in)
		case 17:
			fetchRoutingStatusByASN(in)
		case 18:
			fetchRoutingConsistencyByIP(in)
		case 19:
			fetchRoutingStatusByPrefix(in)
		case 20:
			fetchRoutingConsistencyByPrefix(in)
		case 21:
			fetchRoutingStatusByIPAddressBlock(in)
		case 22:
			fetchRoutingConsistencyByIPAddressBlock(in)
		case 23:
			fetchIPAddressHistory(in)
		case 24:
			fetchASNHistory(in)
		case 25:
			return
		case 26:
			deriveASNFromIPOrPrefix(in)
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func fetchNeighborASNs(in *bufio.Reader) {
	asn := normalizeASN(readLine(in, "Enter ASN: "))
	if asn == "" {
		return
	}
	url := fmt.Sprintf("%s/asn-neighbours/data.json?resource=%s", ripeBase, asn)
	raw, err := fetchBytes(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchNeighborASNs", raw)
}

func fetchHistoricalNeighborASNs(in *bufio.Reader) {
	asn := normalizeASN(readLine(in, "Enter ASN: "))
	if asn == "" {
		return
	}
	url := fmt.Sprintf("%s/asn-neighbours-history/data.json?resource=%s", ripeBase, asn)
	raw, err := fetchBytes(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchHistoricalNeighborASNs", raw)
}

func getAbuseContactInfo(in *bufio.Reader) {
	res := normalizeASN(readLine(in, "Enter IP/Prefix/ASN: "))
	url := fmt.Sprintf("%s/abuse-contact-finder/data.json?resource=%s", ripeBase, res)
	raw, err := fetchBytes(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(prettyJSON(raw))
	writeToFile("getAbuseContactInfo", raw)
}

func getHistoricalWhoisChangeCount(in *bufio.Reader) {
	res := normalizeASN(readLine(in, "Enter IP/Prefix/ASN: "))
	url := fmt.Sprintf("%s/historical-whois/data.json?resource=%s", ripeBase, res)
	raw, err := fetchBytes(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(prettyJSON(raw))
	writeToFile("getHistoricalWhoisChangeCount", raw)
}

func deriveASNFromIPOrPrefix(in *bufio.Reader) {
	res := normalizeASN(readLine(in, "Enter IP or prefix: "))
	if res == "" {
		return
	}
	if !strings.Contains(res, "/") && !strings.HasPrefix(strings.ToUpper(res), "AS") {
		prefix, err := coveringPrefixForIP(res)
		if err != nil {
			fmt.Println(err)
			return
		}
		res = prefix
	}
	url := fmt.Sprintf("%s/routing-status/data.json?resource=%s", ripeBase, res)
	raw := mustFetch(url)
	fmt.Println(prettyJSON(raw))
	writeToFile("deriveASNFromIPOrPrefix", raw)
}

func fetchRoutingHistory(in *bufio.Reader) {
	res := normalizeASN(readLine(in, "Enter ASN/Prefix/IP: "))
	raw := mustFetch(fmt.Sprintf("%s/routing-history/data.json?resource=%s", ripeBase, res))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchRoutingHistory", raw)
}

func fetchPrefixInfo(in *bufio.Reader) {
	asn := normalizeASN(readLine(in, "Enter ASN: "))
	raw := mustFetch(fmt.Sprintf("%s/announced-prefixes/data.json?resource=%s", ripeBase, asn))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchPrefixInfo", raw)
}

func fetchBGPUpdates(in *bufio.Reader) {
	asn := normalizeASN(readLine(in, "Enter ASN: "))
	raw := mustFetch(fmt.Sprintf("%s/bgp-updates/data.json?resource=%s", ripeBase, asn))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchBGPUpdates", raw)
}

func fetchGeolocationInfo(in *bufio.Reader) {
	ip := readLine(in, "Enter IP: ")
	raw := mustFetch(fmt.Sprintf("%s/geoloc/data.json?resource=%s", ripeBase, ip))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchGeolocationInfo", raw)
}

func fetchReverseDNSInfo(in *bufio.Reader) {
	ip := readLine(in, "Enter IP: ")
	raw := mustFetch(fmt.Sprintf("%s/reverse-dns-ip/data.json?resource=%s", ripeBase, ip))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchReverseDNSInfo", raw)
}

func fetchNetworkInfo(in *bufio.Reader) {
	ip := readLine(in, "Enter IP: ")
	raw := mustFetch(fmt.Sprintf("%s/network-info/data.json?resource=%s", ripeBase, ip))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchNetworkInfo", raw)
}

func fetchBlacklistInfo(in *bufio.Reader) {
	ip := readLine(in, "Enter IP: ")
	raw := mustFetch(fmt.Sprintf("%s/blacklist/data.json?resource=%s", ripeBase, ip))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchBlacklistInfo", raw)
}

func fetchIPAddressSpaceInfo(in *bufio.Reader) {
	asn := normalizeASN(readLine(in, "Enter ASN: "))
	raw := mustFetch(fmt.Sprintf("%s/address-space-hierarchy/data.json?resource=%s", ripeBase, asn))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchIPAddressSpaceInfo", raw)
}

func fetchASPathInfo(in *bufio.Reader) {
	asn := normalizeASN(readLine(in, "Enter ASN: "))
	raw := mustFetch(fmt.Sprintf("%s/as-path-length/data.json?resource=%s", ripeBase, asn))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchASPathInfo", raw)
}

func fetchIPAddressBlockInfo(in *bufio.Reader) {
	ip := readLine(in, "Enter IP: ")
	raw := mustFetch(fmt.Sprintf("%s/address-blocks/data.json?resource=%s", ripeBase, ip))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchIPAddressBlockInfo", raw)
}

func fetchRoutingStatusByIP(in *bufio.Reader) {
	ip := readLine(in, "Enter IP: ")
	raw := mustFetch(fmt.Sprintf("%s/routing-status/data.json?resource=%s", ripeBase, ip))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchRoutingStatusByIP", raw)
}

func fetchRoutingConsistencyByASN(in *bufio.Reader) {
	asn := normalizeASN(readLine(in, "Enter ASN: "))
	raw := mustFetch(fmt.Sprintf("%s/as-routing-consistency/data.json?resource=%s", ripeBase, asn))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchRoutingConsistencyByASN", raw)
}

func fetchRoutingStatusByASN(in *bufio.Reader) {
	asn := normalizeASN(readLine(in, "Enter ASN: "))
	raw := mustFetch(fmt.Sprintf("%s/routing-status/data.json?resource=%s", ripeBase, asn))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchRoutingStatusByASN", raw)
}

func fetchRoutingConsistencyByIP(in *bufio.Reader) {
	ip := readLine(in, "Enter IP: ")
	prefix, err := coveringPrefixForIP(ip)
	if err != nil {
		fmt.Println(err)
		return
	}
	raw := mustFetch(fmt.Sprintf("%s/prefix-routing-consistency/data.json?resource=%s", ripeBase, prefix))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchRoutingConsistencyByIP", raw)
}

func fetchRoutingStatusByPrefix(in *bufio.Reader) {
	prefix := readLine(in, "Enter prefix: ")
	raw := mustFetch(fmt.Sprintf("%s/routing-status/data.json?resource=%s", ripeBase, prefix))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchRoutingStatusByPrefix", raw)
}

func fetchRoutingConsistencyByPrefix(in *bufio.Reader) {
	prefix := readLine(in, "Enter prefix: ")
	raw := mustFetch(fmt.Sprintf("%s/prefix-routing-consistency/data.json?resource=%s", ripeBase, prefix))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchRoutingConsistencyByPrefix", raw)
}

func fetchRoutingStatusByIPAddressBlock(in *bufio.Reader) {
	cidr := readLine(in, "Enter CIDR: ")
	raw := mustFetch(fmt.Sprintf("%s/routing-status/data.json?resource=%s", ripeBase, cidr))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchRoutingStatusByIPAddressBlock", raw)
}

func fetchRoutingConsistencyByIPAddressBlock(in *bufio.Reader) {
	cidr := readLine(in, "Enter CIDR: ")
	raw := mustFetch(fmt.Sprintf("%s/prefix-routing-consistency/data.json?resource=%s", ripeBase, cidr))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchRoutingConsistencyByIPAddressBlock", raw)
}

func fetchIPAddressHistory(in *bufio.Reader) {
	ip := readLine(in, "Enter IP: ")
	raw := mustFetch(fmt.Sprintf("%s/address-history/data.json?resource=%s", ripeBase, ip))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchIPAddressHistory", raw)
}

func fetchASNHistory(in *bufio.Reader) {
	asn := normalizeASN(readLine(in, "Enter ASN: "))
	raw := mustFetch(fmt.Sprintf("%s/asn-history/data.json?resource=%s", ripeBase, asn))
	fmt.Println(prettyJSON(raw))
	writeToFile("fetchASNHistory", raw)
}

// ---------- Helpers ----------

func readLine(in *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	s, _ := in.ReadString('\n')
	return strings.TrimSpace(s)
}

func readInt(in *bufio.Reader, prompt string) (int, bool) {
	for {
		s := readLine(in, prompt)
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Enter a number.")
			continue
		}
		return n, true
	}
}

func normalizeASN(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	u := strings.ToUpper(s)
	if strings.HasPrefix(u, "AS") {
		return u
	}
	if _, err := strconv.Atoi(s); err == nil {
		return "AS" + s
	}
	return s
}

func fetchBytes(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "ripestat-cli-go")
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, body)
	}
	return body, nil
}

func mustFetch(url string) []byte {
	raw, err := fetchBytes(url)
	if err != nil {
		fmt.Println(err)
		return []byte("{}")
	}
	return raw
}

func prettyJSON(raw []byte) string {
	var out bytes.Buffer
	if err := json.Indent(&out, raw, "", "  "); err != nil {
		return string(raw)
	}
	return out.String()
}

func coveringPrefixForIP(ip string) (string, error) {
	raw, err := fetchBytes(fmt.Sprintf("%s/network-info/data.json?resource=%s", ripeBase, ip))
	if err != nil {
		return "", err
	}
	var parsed networkInfoResp
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return "", err
	}
	if parsed.Data.Prefix == "" {
		return "", fmt.Errorf("no prefix found")
	}
	return parsed.Data.Prefix, nil
}

func writeToFile(testName string, raw []byte) {
	fileName := fmt.Sprintf("%s_%s.json", testName, time.Now().Format("20060102_150405"))
	var data any
	if json.Valid(raw) {
		data = json.RawMessage(raw)
	} else {
		data = string(raw)
	}
	wrapper := map[string]any{
		"test_name": testName,
		"date_run":  time.Now().Format(time.RFC3339),
		"data":      data,
	}
	content, _ := json.MarshalIndent(wrapper, "", "  ")
	_ = os.WriteFile(fileName, content, 0644)
}
