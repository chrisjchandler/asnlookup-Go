# ASN Lookup Tool & RIPE Interrogator

## Overview

This Go-based tool provides a command-line interface for querying ASN, prefix, and IP information using the RIPEstat API. It includes current and historical routing views, registry lookups, abuse and whois queries, and routing security checks.

## Prerequisites

- Go (1.13 or later recommended)

## Installation

1. **Install Go**: If you haven't already installed Go on your system, download and install it from the official Go website: [https://go.dev/](https://go.dev/)

2. **Download the Script**: Clone this repository or download the `asnlookup.go` file to your local machine.

## Usage

After compiling the Go program, you can run it to access its functionalities. Here are the steps to use the tool:

1. **Run the Program**: Open a terminal and navigate to the directory containing the compiled program. Run it by typing `./asnlookup` if on Unix/Linux/MacOS or `asnlookup.exe` if on Windows.

2. **Choose an Option**: The program will prompt you with the following options:

    1. Fetch neighbor ASNs for a given ASN
    2. Fetch historical neighbor ASNs for a given ASN
    3. Get Abuse contact information for an IP/Prefix/ASN
    4. Get Historical whois change count for an IP/Prefix/ASN
    5. Fetch routing history for a given ASN/Prefix/IP
    6. Fetch prefix information for a given ASN
    7. Fetch BGP updates for a given ASN
    8. Fetch geolocation information for a given IP address
    9. Fetch reverse DNS information for a given IP address
    10. Fetch network information for a given IP address
    11. Fetch blacklist information for a given IP address
    12. Fetch IP address space hierarchy for a given prefix or IP
    13. Fetch AS path information for a given ASN
    14. Fetch address space usage for a given prefix or IP
    15. Fetch routing status for a given IP address
    16. Fetch routing consistency for a given ASN
    17. Fetch routing status for a given ASN
    18. Fetch routing consistency for a given IP address
    19. Fetch routing status for a given prefix
    20. Fetch routing consistency for a given prefix
    21. Fetch routing status for a given IP address block (CIDR)
    22. Fetch routing consistency for a given IP address block (CIDR)
    23. Fetch allocation history for a given IP address or prefix
    24. Fetch allocation history for a given ASN
    25. Derive originating ASN(s) from an IP or prefix (CIDR)
    26. Fetch AS overview for a given ASN
    27. Fetch prefix overview for a given prefix or IP
    28. Fetch whois information for an IP/Prefix/ASN
    29. Fetch RPKI validation status for an ASN and prefix
    30. Fetch visibility information for an IP/Prefix/ASN
    31. Exit

3. **Enter the Required Information**:

    - ASN-focused options prompt for an ASN and accept either `AS12345` or `12345`.
    - Prefix-focused options accept either a CIDR prefix directly or an IP address, which is resolved to its covering prefix when required.
    - Mixed lookup options accept an IP, prefix, or ASN depending on the RIPEstat endpoint being queried.

4. **View the Results**: The program will display the queried information directly in the terminal and give you an option to save the output to a JSON file.

```sh
$ ./asnlookup
Choose an option:
    1. Fetch neighbor ASNs for a given ASN
    ...
    29. Fetch RPKI validation status for an ASN and prefix
    30. Fetch visibility information for an IP/Prefix/ASN
    31. Exit
