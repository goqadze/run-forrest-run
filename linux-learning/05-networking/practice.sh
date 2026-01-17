#!/bin/bash

echo "=== Networking Practice ==="
echo ""

echo "1. Test connectivity:"
# ping - Sends ICMP echo requests to test network connectivity to a host
# -c 3 - Limits to 3 packets (without -c, ping runs indefinitely)
# 8.8.8.8 - Google's public DNS server, commonly used to test internet connectivity
ping -c 3 8.8.8.8
echo ""

echo "2. Fetch URL:"
# curl - Command-line tool for transferring data from/to servers using various protocols
# -s - Silent mode, suppresses progress meter and error messages
# https://httpbin.org/ip - A test endpoint that returns your public IP address in JSON format
# head -5 - Shows only the first 5 lines of the response
curl -s https://httpbin.org/ip | head -5
echo ""

echo "3. Download file:"
# curl with options to check HTTP response without saving the content
# -s - Silent mode (no progress bar)
# -o /dev/null - Discards the downloaded content (sends to null device)
# -w "format" - Writes out specified information after transfer; %{http_code} prints the HTTP status code
# Used to verify a URL is reachable without downloading the actual content
curl -s -o /dev/null -w "HTTP Status: %{http_code}\n" https://example.com
echo ""

echo "4. Check hostname:"
# hostname - Displays the system's network hostname (the name identifying this computer on the network)
hostname
echo ""

echo "Done!"
