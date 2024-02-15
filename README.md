ASN Lookup Tool
Overview
This Go-based tool provides a command-line interface for querying ASN information, including ASN and organization names by IP address, current neighbor ASNs, and historical neighbor ASNs using the RIPE Stat API.

Prerequisites
Go (1.13 or later recommended)
Installation
Install Go: If you haven't already installed Go on your system, download and install it from the official Go website. https://go.dev/

Download the Script: Clone this repository or download the asnlookup.go file to your local machine.

Usage
After compiling the Go program, you can run it to access its functionalities. Here are the steps to use the tool:

Run the Program:
Open a terminal and navigate to the directory containing the compiled program. Run it by typing ./asnlookup if on Unix/Linux/MacOS or asnlookup.exe if on Windows.

Choose an Option:
The program will prompt you with three options:

1 for looking up ASN and organization name by IP address.
2 for fetching neighbor ASNs for a given ASN.
3 for fetching historical neighbor ASNs for a given ASN.
4 shows abuse contacts for an asn or an ip
5 shows the historical nubmer of times whois records for an IP have changed
Enter the Required Information:

For option 1, you will be prompted to enter an IP address.
For options 2 and 3, you will be prompted to enter an ASN.
View the Results:
The program will display the queried information directly in the terminal.

$ ./asnlookup
Choose an option:
1. Look up ASN and organization name by IP address
2. Fetch neighbor ASNs for a given ASN
3. Fetch historical neighbor ASNs for a given ASN
4. Get Abuse contact information for an ip
5. Get Historical whois information for a domain
6. Exit

Enter an IP address to look up: 8.8.8.8


Note
The accuracy and availability of data depend on the RIPE Stat API's current status and data coverage.

An IP prefix may have multiple associated ASN's (for a variety of reasons) you can use my RPKI checker to get other ASN's https://github.com/chrisjchandler/RPKI_check_tool
