#!/bin/bash

echo "=== System Management Practice ==="
echo ""

echo "1. System information:"
uname -a
echo ""

echo "2. Disk space:"
df -h | head -5
echo ""

echo "3. Memory usage:"
free -h
echo ""

echo "4. Current processes:"
ps aux | head -10
echo ""

echo "5. Current user:"
whoami
echo ""

echo "6. Uptime:"
uptime
echo ""

echo "Done!"
