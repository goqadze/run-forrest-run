#!/bin/bash

echo "=== Text Processing Practice ==="
echo ""

# Create sample file
cat > sample.txt << 'SAMPLE'
apple 10
banana 20
apple 15
cherry 30
banana 25
SAMPLE

echo "Sample file created:"
cat sample.txt
echo ""

echo "1. Search with grep:"
grep "apple" sample.txt
echo ""

echo "2. Sort:"
sort sample.txt
echo ""

echo "3. Unique values:"
cut -d' ' -f1 sample.txt | sort | uniq
echo ""

echo "4. Sum with awk:"
awk '{sum+=$2} END {print "Total:", sum}' sample.txt
echo ""

echo "5. Replace with sed:"
sed 's/apple/APPLE/g' sample.txt
echo ""

rm sample.txt
echo "Done!"
