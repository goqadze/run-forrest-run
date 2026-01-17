#!/bin/bash

echo "=== File Operations Practice ==="
echo ""

# Create a practice directory to work in
DIR="practice-files"
# mkdir -p - Creates directory; -p flag prevents error if it already exists
mkdir -p $DIR
# cd - Change into the practice directory
cd $DIR

echo "1. Create files:"
# echo "text" > file - Creates a new file and writes text to it (overwrites if exists)
echo "Hello World" > file1.txt
# echo "text" > file - Creates file2.txt with "Line 1"
echo "Line 1" > file2.txt
# echo "text" >> file - Appends text to existing file (>> means append, > means overwrite)
echo "Line 2" >> file2.txt
# cat - Concatenate and display file contents to the terminal
cat file1.txt
echo ""

echo "2. Copy and move:"
# cp source dest - Copies a file from source to destination
cp file1.txt file1-copy.txt
# mv source dest - Moves/renames a file from source to destination
mv file1-copy.txt backup.txt
# ls -la - Lists all files in long format to verify the copy and move operations
ls -la
echo ""

echo "3. File permissions:"
# chmod 644 - Sets file permissions: owner can read/write (6), group and others can only read (4)
# Permission format: owner(rwx) group(rwx) others(rwx) where r=4, w=2, x=1
chmod 644 file1.txt
# chmod 755 - Sets permissions: owner has full access (7), group and others can read/execute (5)
chmod 755 backup.txt
# ls -l - Shows file permissions in the first column (e.g., -rw-r--r--)
ls -l
echo ""

echo "4. View file:"
# cat - Displays the entire contents of a file
cat file2.txt
echo ""

echo "5. Count lines:"
# wc -l - Word count with -l flag counts only the number of lines in the file
wc -l file2.txt
echo ""

# cd .. - Move up one directory level (back to parent directory)
cd ..
# rm -rf - Remove directory and all contents recursively (-r) and forcefully (-f)
rm -rf $DIR
echo "Cleanup done!"
