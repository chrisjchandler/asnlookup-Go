package main

import (
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
"time"
)

// ASNResponse represents the common structure of responses from RIPE Stat API
type ASNResponse struct {
Data struct {
Asns []struct {
Asn int `json:"asn"`
Holder string `json:"holder"`
} `json:"asns"`
Neighbours []struct {
Asn int `json:"asn"`
Type string `json:"type"`
Power int `json:"power"`
V4Peers int `json:"v4_peers"`
V6Peers int `json:"v6_peers"`
Timelines []struct {
Starttime string `json:"starttime"`
Endtime string `json:"endtime"`
} `json:"timelines,omitempty"`
} `json:"neighbours"`
} `json:"data"`
}

func main() {
for {
fmt.Println("Choose an option:")
fmt.Println("1. Fetch neighbor ASNs for a given ASN")
fmt.Println("2. Fetch historical neighbor ASNs for a given ASN")
fmt.Println("3. Get Abuse contact information for an IP")
fmt.Println("4. Get Historical whois change count for an IP Address")
fmt.Println("5. Fetch routing history for a given ASN")
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
fmt.Println("21. Fetch routing status for a given IP address block")
fmt.Println("22. Fetch routing consistency for a given IP address block")
fmt.Println("23. Fetch IP address history for a given IP address")
fmt.Println("24. Fetch ASN history for a given ASN")
fmt.Println("25. Exit")
var choice int
fmt.Scanln(&choice)

switch choice {
case 1:
fetchNeighborASNs()
case 2:
fetchHistoricalNeighborASNs()
case 3:
getAbuseContactInfo()
case 4:
getHistoricalWhoisChangeCount()
case 5:
fetchRoutingHistory()
case 6:
fetchPrefixInfo()
case 7:
fetchBGPUpdates()
case 8:
fetchGeolocationInfo()
case 9:
fetchReverseDNSInfo()
case 10:
fetchNetworkInfo()
case 11:
fetchBlacklistInfo()
case 12:
fetchIPAddressSpaceInfo()
case 13:
fetchASPathInfo()
case 14:
fetchIPAddressBlockInfo()
case 15:
fetchRoutingStatusByIP()
case 16:
fetchRoutingConsistencyByASN()
case 17:
fetchRoutingStatusByASN()
case 18:
fetchRoutingConsistencyByIP()
case 19:
fetchRoutingStatusByPrefix()
case 20:
fetchRoutingConsistencyByPrefix()
case 21:
fetchRoutingStatusByIPAddressBlock()
case 22:
fetchRoutingConsistencyByIPAddressBlock()
case 23:
fetchIPAddressHistory()
case 24:
fetchASNHistory()
case 25:
return
default:
fmt.Println("Invalid choice. Please try again.")
}
}
}

func fetchNeighborASNs() {
// Implementation for fetching neighbor ASNs
}

func fetchHistoricalNeighborASNs() {
// Implementation for fetching historical neighbor ASNs
}

func getAbuseContactInfo() {
// Implementation for getting abuse contact information
}

func getHistoricalWhoisChangeCount() {
// Implementation for getting historical whois change count
}

func fetchRoutingHistory() {
asn := getInput("Enter ASN: ")
url := fmt.Sprintf("https://stat.ripe.net/data/routing-history/data.json?resource=%s", asn)
response := fetchData(url)
fmt.Println("Routing History:", response)
writeToFile("fetchRoutingHistory", response)
}

func fetchPrefixInfo() {
asn := getInput("Enter ASN: ")
url := fmt.Sprintf("https://stat.ripe.net/data/announced-prefixes/data.json?resource=%s", asn)
response := fetchData(url)
fmt.Println("Prefix Information:", response)
writeToFile("fetchPrefixInfo", response)
}

func fetchBGPUpdates() {
asn := getInput("Enter ASN: ")
url := fmt.Sprintf("https://stat.ripe.net/data/bgp-updates/data.json?resource=%s", asn)
response := fetchData(url)
fmt.Println("BGP Updates:", response)
writeToFile("fetchBGPUpdates", response)
}

func fetchGeolocationInfo() {
ip := getInput("Enter IP address: ")
url := fmt.Sprintf("https://stat.ripe.net/data/geoloc/data.json?resource=%s", ip)
response := fetchData(url)
fmt.Println("Geolocation Information:", response)
writeToFile("fetchGeolocationInfo", response)
}

func fetchReverseDNSInfo() {
ip := getInput("Enter IP address: ")
url := fmt.Sprintf("https://stat.ripe.net/data/reverse-dns-ip/data.json?resource=%s", ip)
response := fetchData(url)
fmt.Println("Reverse DNS Information:", response)
writeToFile("fetchReverseDNSInfo", response)
}

func fetchNetworkInfo() {
ip := getInput("Enter IP address: ")
url := fmt.Sprintf("https://stat.ripe.net/data/network-info/data.json?resource=%s", ip)
response := fetchData(url)
fmt.Println("Network Information:", response)
writeToFile("fetchNetworkInfo", response)
}

func fetchBlacklistInfo() {
ip := getInput("Enter IP address: ")
url := fmt.Sprintf("https://stat.ripe.net/data/blacklist/data.json?resource=%s", ip)
response := fetchData(url)
fmt.Println("Blacklist Information:", response)
writeToFile("fetchBlacklistInfo", response)
}

func fetchIPAddressSpaceInfo() {
asn := getInput("Enter ASN: ")
url := fmt.Sprintf("https://stat.ripe.net/data/address-space-hierarchy/data.json?resource=%s", asn)
response := fetchData(url)
fmt.Println("IP Address Space Information:", response)
writeToFile("fetchIPAddressSpaceInfo", response)
}

func fetchASPathInfo() {
asn := getInput("Enter ASN: ")
url := fmt.Sprintf("https://stat.ripe.net/data/as-path-length/data.json?resource=%s", asn)
response := fetchData(url)
fmt.Println("AS Path Information:", response)
writeToFile("fetchASPathInfo", response)
}

func fetchIPAddressBlockInfo() {
ip := getInput("Enter IP address: ")
url := fmt.Sprintf("https://stat.ripe.net/data/address-blocks/data.json?resource=%s", ip)
response := fetchData(url)
fmt.Println("IP Address Block Information:", response)
writeToFile("fetchIPAddressBlockInfo", response)
}

func fetchRoutingStatusByIP() {
ip := getInput("Enter IP address: ")
url := fmt.Sprintf("https://stat.ripe.net/data/routing-status/data.json?resource=%s", ip)
response := fetchData(url)
fmt.Println("Routing Status by IP:", response)
writeToFile("fetchRoutingStatusByIP", response)
}

func fetchRoutingConsistencyByASN() {
asn := getInput("Enter ASN: ")
url := fmt.Sprintf("https://stat.ripe.net/data/routing-consistency/data.json?resource=%s", asn)
response := fetchData(url)
fmt.Println("Routing Consistency by ASN:", response)
writeToFile("fetchRoutingConsistencyByASN", response)
}

func fetchRoutingStatusByASN() {
asn := getInput("Enter ASN: ")
url := fmt.Sprintf("https://stat.ripe.net/data/routing-status/data.json?resource=%s", asn)
response := fetchData(url)
fmt.Println("Routing Status by ASN:", response)
writeToFile("fetchRoutingStatusByASN", response)
}

func fetchRoutingConsistencyByIP() {
ip := getInput("Enter IP address: ")
url := fmt.Sprintf("https://stat.ripe.net/data/routing-consistency/data.json?resource=%s", ip)
response := fetchData(url)
fmt.Println("Routing Consistency by IP:", response)
writeToFile("fetchRoutingConsistencyByIP", response)
}

func fetchRoutingStatusByPrefix() {
prefix := getInput("Enter prefix: ")
url := fmt.Sprintf("https://stat.ripe.net/data/routing-status/data.json?resource=%s", prefix)
response := fetchData(url)
fmt.Println("Routing Status by Prefix:", response)
writeToFile("fetchRoutingStatusByPrefix", response)
}

func fetchRoutingConsistencyByPrefix() {
prefix := getInput("Enter prefix: ")
url := fmt.Sprintf("https://stat.ripe.net/data/routing-consistency/data.json?resource=%s", prefix)
response := fetchData(url)
fmt.Println("Routing Consistency by Prefix:", response)
writeToFile("fetchRoutingConsistencyByPrefix", response)
}

func fetchRoutingStatusByIPAddressBlock() {
ipBlock := getInput("Enter IP address block: ")
url := fmt.Sprintf("https://stat.ripe.net/data/routing-status/data.json?resource=%s", ipBlock)
response := fetchData(url)
fmt.Println("Routing Status by IP Address Block:", response)
writeToFile("fetchRoutingStatusByIPAddressBlock", response)
}

func fetchRoutingConsistencyByIPAddressBlock() {
ipBlock := getInput("Enter IP address block: ")
url := fmt.Sprintf("https://stat.ripe.net/data/routing-consistency/data.json?resource=%s", ipBlock)
response := fetchData(url)
fmt.Println("Routing Consistency by IP Address Block:", response)
writeToFile("fetchRoutingConsistencyByIPAddressBlock", response)
}

func fetchIPAddressHistory() {
ip := getInput("Enter IP address: ")
url := fmt.Sprintf("https://stat.ripe.net/data/address-history/data.json?resource=%s", ip)
response := fetchData(url)
fmt.Println("IP Address History:", response)
writeToFile("fetchIPAddressHistory", response)
}

func fetchASNHistory() {
asn := getInput("Enter ASN: ")
url := fmt.Sprintf("https://stat.ripe.net/data/asn-history/data.json?resource=%s", asn)
response := fetchData(url)
fmt.Println("ASN History:", response)
writeToFile("fetchASNHistory", response)
}

func getInput(prompt string) string {
var input string
fmt.Print(prompt)
fmt.Scanln(&input)
return input
}

func fetchData(url string) string {
resp, err := http.Get(url)
if err != nil {
fmt.Println("Error fetching data:", err)
return ""
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
fmt.Println("Error reading response body:", err)
return ""
}
return string(body)
}

func writeToFile(testName, data string) {
fileName := fmt.Sprintf("%s_%s.json", testName, time.Now().Format("20060102_150405"))
fileData := map[string]interface{}{
"test_name": testName,
"date_run": time.Now().Format(time.RFC3339),
"data": json.RawMessage(data),
}
fileContent, err := json.MarshalIndent(fileData, "", " ")
if err != nil {
fmt.Println("Error marshalling data to JSON:", err)
return
}
err = ioutil.WriteFile(fileName, fileContent, 0644)
if err != nil {
fmt.Println("Error writing to file:", err)
return
}
fmt.Printf("Data written to file: %s\n", fileName)
}
