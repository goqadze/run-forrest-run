#!/bin/bash

echo "=== Linux Basics Practice ==="
echo ""

echo "1. Where are we?"
pwd
echo ""

echo "2. What's here?"
ls
echo ""

echo "3. Detailed listing:"
ls -la
echo ""

echo "4. Create test directory:"
mkdir -p test-dir/subdir
echo "Created: test-dir/subdir"
echo ""

echo "5. Navigate:"
cd test-dir && pwd
cd - > /dev/null
echo ""

echo "6. Create test files:"
touch test-dir/file1.txt test-dir/file2.log test-dir/file3.txt
ls test-dir
echo ""

echo "7. List only .txt files:"
ls test-dir/*.txt
echo ""

echo "8. Cleanup:"
rm -rf test-dir
echo "Done!"
