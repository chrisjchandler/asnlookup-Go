# ASN Lookup Tool & RIPE Interrogator

## Overview

This Go-based tool provides a command-line interface for querying ASN information using the RIPE Stat API. It includes functionalities for looking up ASN and organization names by IP address, fetching current and historical neighbor ASNs, and much more.

## Prerequisites

- Go (1.13 or later recommended)

## Installation

1. **Install Go**: If you haven't already installed Go on your system, download and install it from the official Go website: [https://go.dev/](https://go.dev/)

2. **Download the Script**: Clone this repository or download the `asnlookup.go` file to your local machine.

## Usage

After compiling the Go program, you can run it to access its functionalities. Here are the steps to use the tool:

1. **Run the Program**: Open a terminal and navigate to the directory containing the compiled program. Run it by typing `./asnlookup` if on Unix/Linux/MacOS or `asnlookup.exe` if on Windows.

2. **Choose an Option**: The program will prompt you with multiple options:

    1. Look up ASN and organization name by IP address
    2. Fetch neighbor ASNs for a given ASN
    3. Fetch historical neighbor ASNs for a given ASN
    4. Get Abuse contact information for an IP
    5. Get Historical whois change count for an IP Address
    6. Fetch routing history for a given ASN
    7. Fetch prefix information for a given ASN
    8. Fetch BGP updates for a given ASN
    9. Fetch geolocation information for a given IP address
    10. Fetch reverse DNS information for a given IP address
    11. Fetch network information for a given IP address
    12. Fetch blacklist information for a given IP address
    13. Fetch IP address space information for a given ASN
    14. Fetch AS path information for a given ASN
    15. Fetch IP address block information for a given IP address
    16. Fetch routing status for a given IP address
    17. Fetch routing consistency for a given ASN
    18. Fetch routing status for a given ASN
    19. Fetch routing consistency for a given IP address
    20. Fetch routing status for a given prefix
    21. Fetch routing consistency for a given prefix
    22. Fetch routing status for a given IP address block
    23. Fetch routing consistency for a given IP address block
    24. Fetch IP address history for a given IP address
    25. Fetch ASN history for a given ASN
    26. Exit

3. **Enter the Required Information**:

    - For option 1, you will be prompted to enter an IP address.
    - For options 2, 3, 6, 7, 8, 13, 14, 17, and 18, you will be prompted to enter an ASN.
    - For options 4, 9, 10, 11, 12, 15, 16, 19, and 20, you will be prompted to enter an IP address.
    - For options 21 and 22, you will be prompted to enter a prefix.
    - For options 23, 24, and 25, you will be prompted to enter an IP address block.

4. **View the Results**: The program will display the queried information directly in the terminal and give you an option to save the output to a JSON file.

```sh
$ ./asnlookup
Choose an option:
    1. Look up ASN and organization name by IP address
    2. Fetch neighbor ASNs for a given ASN
    3. Fetch historical neighbor ASNs for a given ASN
    4. Get Abuse contact information for an IP
    5. Get Historical whois change count for an IP Address
    6. Fetch routing history for a given ASN
    7. Fetch prefix information for a given ASN
    8. Fetch BGP updates for a given ASN
    9. Fetch geolocation information for a given IP address
    10. Fetch reverse DNS information for a given IP address
    11. Fetch network information for a given IP address
    12. Fetch blacklist information for a given IP address
    13. Fetch IP address space information for a given ASN
    14. Fetch AS path information for a given ASN
    15. Fetch IP address block information for a given IP address
    16. Fetch routing status for a given IP address
    17. Fetch routing consistency for a given ASN
    18. Fetch routing status for a given ASN
    19. Fetch routing consistency for a given IP address
    20. Fetch routing status for a given prefix
    21. Fetch routing consistency for a given prefix
    22. Fetch routing status for a given IP address block
    23. Fetch routing consistency for a given IP address block
    24. Fetch IP address history for a given IP address
    25. Fetch ASN history for a given ASN
    26. Exit

Enter an IP address to look up: 8.8.8.8
