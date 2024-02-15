package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ASNResponse represents the common structure of responses from RIPE Stat API
type ASNResponse struct {
	Data struct {
		Asns []struct {
			Asn    int    `json:"asn"`   // Change this to int
			Holder string `json:"holder"`
		} `json:"asns"`
		Neighbours []struct {
			Asn      int    `json:"asn"` // Also change here to int
			Type     string `json:"type"`
			Power    int    `json:"power"`
			V4Peers  int    `json:"v4_peers"`
			V6Peers  int    `json:"v6_peers"`
			Timelines []struct {
				Starttime string `json:"starttime"`
				Endtime   string `json:"endtime"`
			} `json:"timelines,omitempty"`
		} `json:"neighbours"`
	} `json:"data"`
}

func main() {
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Look up ASN and organization name by IP address")
		fmt.Println("2. Fetch neighbor ASNs for a given ASN")
		fmt.Println("3. Fetch historical neighbor ASNs for a given ASN")
		fmt.Println("4. Get Abuse contact information for an ip")
		fmt.Println("5. Get Historical whois information for a domain")
		fmt.Println("6. Exit")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var ip string
			fmt.Println("Enter an IP address to look up: ")
			fmt.Scanln(&ip)
			lookupASNByIP(ip)
		case 2:
			var asn string
			fmt.Println("Enter an ASN to find its neighbors: ")
			fmt.Scanln(&asn)
			fetchASNNeighbors(asn, false)
		case 3:
			var asn string
			fmt.Println("Enter an ASN to find its historical neighbors: ")
			fmt.Scanln(&asn)
			fetchASNNeighbors(asn, true)
		case 4:
			var resource string
			fmt.Println("Enter an ASN or IP address to find its abuse contact:")
			fmt.Scanf("%s", &resource)
			fetchAbuseContact(resource)
		case 5:
			var domain string
			fmt.Println("Enter a domain to find its historical WHOIS information:")
			fmt.Scanf("%s", &domain)
			fetchHistoricalWhois(domain)
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 6.")
		}
	}
}

func lookupASNByIP(ip string) {
	fetchData("https://stat.ripe.net/data/prefix-overview/data.json?resource="+ip, false)
}

func fetchAbuseContact(resource string) {
	resp, err := http.Get(fmt.Sprintf("https://stat.ripe.net/data/abuse-contact-finder/data.json?resource=%s", resource))
	if err != nil {
		fmt.Println("Error fetching abuse contact:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var abuseContact struct {
		Data struct {
			AbuseContacts []string `json:"abuse_contacts"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &abuseContact); err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return
	}

	fmt.Println("Abuse Contacts:")
	for _, contact := range abuseContact.Data.AbuseContacts {
		fmt.Println(contact)
	}
}

func fetchHistoricalWhois(domain string) {
	resp, err := http.Get(fmt.Sprintf("https://stat.ripe.net/data/historical-whois/data.json?resource=%s", domain))
	if err != nil {
		fmt.Println("Error fetching historical WHOIS information:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var historicalWhois struct {
		Data struct {
			Versions []struct {
				Version int    `json:"version"`
				Time    string `json:"time"`
			} `json:"versions"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &historicalWhois); err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return
	}

	fmt.Println("Historical WHOIS Versions:")
	for _, version := range historicalWhois.Data.Versions {
		fmt.Printf("Version: %d, Time: %s\n", version.Version, version.Time)
	}
}

func fetchASNNeighbors(asn string, historical bool) {
	var url string
	if historical {
		url = "https://stat.ripe.net/data/asn-neighbours-history/data.json?resource=" + asn
	} else {
		url = "https://stat.ripe.net/data/asn-neighbours/data.json?resource=" + asn
	}
	fetchData(url, historical)
}

func fetchData(url string, historical bool) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var response ASNResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	if historical {
		for _, neighbour := range response.Data.Neighbours {
			fmt.Printf("ASN: %d, Timelines:\n", neighbour.Asn)
			for _, timeline := range neighbour.Timelines {
				fmt.Printf("    Start: %s, End: %s\n", timeline.Starttime, timeline.Endtime)
			}
		}
	} else {
		for _, asnInfo := range response.Data.Asns {
			fmt.Printf("ASN: %d, Holder: %s\n", asnInfo.Asn, asnInfo.Holder)
		}
		for _, neighbour := range response.Data.Neighbours {
			fmt.Printf("ASN: %d, Type: %s, Power: %d, IPv4 Peers: %d, IPv6 Peers: %d\n",
				neighbour.Asn, neighbour.Type, neighbour.Power, neighbour.V4Peers, neighbour.V6Peers)
		}
	}
}
