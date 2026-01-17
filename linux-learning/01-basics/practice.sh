#!/bin/bash

echo "=== Linux Basics Practice ==="
echo ""

echo "1. Where are we?"
# pwd (Print Working Directory) - Displays the full path of the current directory you're in
pwd
echo ""

echo "2. What's here?"
# ls (List) - Lists all files and directories in the current directory
ls
echo ""

echo "3. Detailed listing:"
# ls -la - Lists all files including hidden ones (-a) in long format (-l) showing permissions, owner, size, and date
ls -la
echo ""

echo "4. Create test directory:"
# mkdir -p - Creates a directory and any necessary parent directories (-p prevents errors if directory exists)
mkdir -p test-dir/subdir
echo "Created: test-dir/subdir"
echo ""

echo "5. Navigate:"
# cd (Change Directory) - Changes the current working directory to the specified path
# pwd - Prints the new current directory to confirm the change
cd test-dir && pwd
# cd - - Returns to the previous directory; > /dev/null suppresses output
cd - > /dev/null
echo ""

echo "6. Create test files:"
# touch - Creates empty files or updates the timestamp of existing files
touch test-dir/file1.txt test-dir/file2.log test-dir/file3.txt
# ls - Lists contents of test-dir to verify files were created
ls test-dir
echo ""

echo "7. List only .txt files:"
# ls with wildcard (*.txt) - Lists only files matching the pattern (files ending in .txt)
ls test-dir/*.txt
echo ""

echo "8. Cleanup:"
# rm -rf - Removes files/directories recursively (-r) and forcefully (-f) without prompting
# WARNING: Use with caution as this permanently deletes files
rm -rf test-dir
echo "Done!"
