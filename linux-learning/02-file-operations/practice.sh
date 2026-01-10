#!/bin/bash

echo "=== File Operations Practice ==="
echo ""

DIR="practice-files"
mkdir -p $DIR
cd $DIR

echo "1. Create files:"
echo "Hello World" > file1.txt
echo "Line 1" > file2.txt
echo "Line 2" >> file2.txt
cat file1.txt
echo ""

echo "2. Copy and move:"
cp file1.txt file1-copy.txt
mv file1-copy.txt backup.txt
ls -la
echo ""

echo "3. File permissions:"
chmod 644 file1.txt
chmod 755 backup.txt
ls -l
echo ""

echo "4. View file:"
cat file2.txt
echo ""

echo "5. Count lines:"
wc -l file2.txt
echo ""

cd ..
rm -rf $DIR
echo "Cleanup done!"
