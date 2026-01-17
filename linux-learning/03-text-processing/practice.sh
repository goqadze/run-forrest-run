#!/bin/bash

echo "=== Text Processing Practice ==="
echo ""

# Create sample file using heredoc
# cat > file << 'DELIMITER' - Creates a file with multi-line content (heredoc syntax)
# The 'SAMPLE' delimiter marks the start and end of the content block
cat > sample.txt << 'SAMPLE'
apple 10
banana 20
apple 15
cherry 30
banana 25
SAMPLE

echo "Sample file created:"
# cat - Display the contents of the sample file
cat sample.txt
echo ""

echo "1. Search with grep:"
# grep "pattern" file - Searches for lines containing the pattern and prints matching lines
# Here it finds all lines containing "apple"
grep "apple" sample.txt
echo ""

echo "2. Sort:"
# sort - Sorts lines of text alphabetically (default) or numerically with -n flag
# Outputs the sorted content without modifying the original file
sort sample.txt
echo ""

echo "3. Unique values:"
# cut -d' ' -f1 - Cuts/extracts the first field (-f1) using space as delimiter (-d' ')
# sort - Sorts the extracted values alphabetically
# uniq - Removes adjacent duplicate lines (requires sorted input to work properly)
# The pipe (|) passes output from one command as input to the next
cut -d' ' -f1 sample.txt | sort | uniq
echo ""

echo "4. Sum with awk:"
# awk - Powerful text processing tool that processes text line by line
# {sum+=$2} - For each line, add the second field ($2) to the sum variable
# END {print} - After processing all lines, print the total
awk '{sum+=$2} END {print "Total:", sum}' sample.txt
echo ""

echo "5. Replace with sed:"
# sed 's/old/new/g' - Stream editor that performs text substitution
# s = substitute, g = global (replace all occurrences, not just the first)
# This replaces all "apple" with "APPLE" and prints to stdout (doesn't modify file)
sed 's/apple/APPLE/g' sample.txt
echo ""

# rm - Remove/delete the sample file
rm sample.txt
echo "Done!"
