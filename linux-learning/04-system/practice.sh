#!/bin/bash

echo "=== System Management Practice ==="
echo ""

echo "1. System information:"
# uname -a - Displays all system information including:
# kernel name, hostname, kernel version, machine hardware, processor, OS
uname -a
echo ""

echo "2. Disk space:"
# df -h - Displays disk filesystem usage in human-readable format (-h)
# Shows total size, used space, available space, and mount points
# head -5 - Limits output to first 5 lines (header + 4 filesystems)
df -h | head -5
echo ""

echo "3. Memory usage:"
# free -h - Displays memory (RAM) and swap usage in human-readable format (-h)
# Shows total, used, free, shared, buff/cache, and available memory
free -h
echo ""

echo "4. Current processes:"
# ps aux - Lists all running processes with detailed information
# a = show processes for all users, u = user-oriented format, x = include processes without terminals
# Columns: USER, PID, %CPU, %MEM, VSZ, RSS, TTY, STAT, START, TIME, COMMAND
# head -10 - Limits output to first 10 processes
ps aux | head -10
echo ""

echo "5. Current user:"
# whoami - Prints the username of the current logged-in user
whoami
echo ""

echo "6. Uptime:"
# uptime - Shows how long the system has been running, number of users, and load averages
# Load averages show system load over the last 1, 5, and 15 minutes
uptime
echo ""

echo "Done!"
