#!/bin/bash

echo "=== Networking Practice ==="
echo ""

echo "1. Test connectivity:"
ping -c 3 8.8.8.8
echo ""

echo "2. Fetch URL:"
curl -s https://httpbin.org/ip | head -5
echo ""

echo "3. Download file:"
curl -s -o /dev/null -w "HTTP Status: %{http_code}\n" https://example.com
echo ""

echo "4. Check hostname:"
hostname
echo ""

echo "Done!"
