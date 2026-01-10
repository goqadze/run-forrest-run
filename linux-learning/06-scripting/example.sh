#!/bin/bash
# Example bash script

# Variables
NAME="Linux"
VERSION="1.0"

# Function
greet() {
    echo "Hello from $1!"
}

# Main
echo "Script: $VERSION"
greet "$NAME"

# Control flow
for i in {1..3}; do
    echo "Count: $i"
done

# Conditional
if [ -f "/etc/passwd" ]; then
    echo "System file exists"
fi

exit 0
