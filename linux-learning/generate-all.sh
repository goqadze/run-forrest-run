#!/bin/bash
set -e
BASE="/Users/avtandilgokadze/Learning/linux-learning"

echo "🐧 Generating Linux Learning Project..."

# Main README
cat > "$BASE/README.md" << 'MAINREADME'
# Linux Commands Learning Project

Master Linux command line through hands-on practice and comprehensive examples.

## Prerequisites

- Linux, macOS, or WSL (Windows Subsystem for Linux)
- Terminal/Shell access
- Basic computer literacy

## Project Structure

```
linux-learning/
├── README.md                    # You are here!
├── LINUX_COMMANDS.md            # Complete command reference
├── 01-basics/                   # Navigation and basic commands
├── 02-file-operations/          # File management and permissions
├── 03-text-processing/          # grep, sed, awk, text tools
├── 04-system/                   # Processes, users, system info
├── 05-networking/               # Network commands and tools
├── 06-scripting/                # Bash scripting basics
└── exercises/                   # Practice challenges
```

## Getting Started

### Verify Your Environment

```bash
# Check your shell
echo $SHELL

# Check available commands
which bash ls grep

# Print working directory
pwd
```

### Start Learning

```bash
cd 01-basics
cat README.md
bash practice.sh
```

## Topics Covered

### 1. Basics (`01-basics/`)
**What you'll learn:**
- Navigation (cd, pwd, ls)
- Directory operations (mkdir, rmdir)
- Path concepts (absolute vs relative)
- Command structure and options
- Getting help (man, --help)

**Key concepts:**
- Everything is a file in Linux
- Case sensitivity matters
- Wildcards and patterns
- Command options (-a, --all)

### 2. File Operations (`02-file-operations/`)
**What you'll learn:**
- Creating and viewing files (touch, cat, less, head, tail)
- Copying and moving (cp, mv)
- Deleting files (rm)
- File permissions (chmod, chown)
- Links (ln, symbolic vs hard)
- Finding files (find, locate)

**Key concepts:**
- Permission bits (rwx)
- User, group, others
- Recursive operations
- File ownership

### 3. Text Processing (`03-text-processing/`)
**What you'll learn:**
- Searching text (grep)
- Stream editing (sed)
- Text processing (awk)
- Sorting and filtering (sort, uniq, cut)
- Word count (wc)
- Comparing files (diff)

**Key concepts:**
- Regular expressions
- Pipes and redirection
- Text streams
- Pattern matching

### 4. System Management (`04-system/`)
**What you'll learn:**
- Process management (ps, top, htop, kill)
- User management (useradd, passwd, sudo)
- System information (uname, df, du, free)
- Package management (apt, yum, brew)
- System logs (journalctl, dmesg)

**Key concepts:**
- Process IDs (PID)
- User permissions
- System resources
- Service management

### 5. Networking (`05-networking/`)
**What you'll learn:**
- Network info (ip, ifconfig, netstat)
- Connectivity (ping, traceroute)
- File transfer (scp, rsync, wget, curl)
- Remote access (ssh)
- DNS tools (dig, nslookup)

**Key concepts:**
- IP addresses and ports
- TCP/IP basics
- SSH keys
- HTTP requests

### 6. Bash Scripting (`06-scripting/`)
**What you'll learn:**
- Script structure
- Variables and arguments
- Control flow (if, for, while)
- Functions
- Exit codes
- Best practices

**Key concepts:**
- Shebang (#!/bin/bash)
- Script execution
- Error handling
- Command substitution

## Learning Path

**Recommended progression:**

1. **Basics** - Master navigation and file basics
2. **File Operations** - Manage files and permissions
3. **Text Processing** - Search and manipulate text
4. **System** - Understand processes and system
5. **Networking** - Network operations and SSH
6. **Scripting** - Automate tasks with scripts
7. **Practice** - Complete exercises

## Essential Commands Quick Reference

### Navigation
```bash
pwd                    # Print working directory
ls                     # List files
ls -la                 # List all files with details
cd /path              # Change directory
cd ~                  # Go to home directory
cd ..                 # Go up one directory
```

### File Operations
```bash
touch file.txt        # Create empty file
cat file.txt          # View file contents
less file.txt         # View file (paginated)
cp file1 file2        # Copy file
mv file1 file2        # Move/rename file
rm file.txt           # Delete file
mkdir dir             # Create directory
rm -rf dir            # Delete directory recursively
```

### Text Processing
```bash
grep "pattern" file   # Search for pattern
grep -r "pattern" .   # Recursive search
sed 's/old/new/g' file  # Replace text
awk '{print $1}' file   # Print first column
sort file             # Sort lines
uniq file             # Remove duplicates
wc -l file            # Count lines
```

### System
```bash
ps aux                # List all processes
top                   # Monitor processes
kill PID              # Kill process
df -h                 # Disk space
du -sh *              # Directory sizes
free -h               # Memory usage
sudo command          # Run as superuser
```

### Networking
```bash
ping google.com       # Test connectivity
curl example.com      # Fetch URL
wget file.zip         # Download file
ssh user@host         # Remote login
scp file user@host:~  # Copy file to remote
```

## Common Patterns

### Pipes and Redirection
```bash
# Redirect output to file
ls > files.txt

# Append to file
echo "text" >> file.txt

# Redirect errors
command 2> errors.txt

# Pipe output to another command
ls -la | grep ".txt"
cat file.txt | sort | uniq

# Chain multiple commands
command1 && command2    # Run command2 if command1 succeeds
command1 || command2    # Run command2 if command1 fails
command1 ; command2     # Run both sequentially
```

### Wildcards
```bash
ls *.txt              # All .txt files
ls file?.txt          # file1.txt, fileA.txt, etc.
ls file[0-9].txt      # file0.txt through file9.txt
```

### Finding Files
```bash
find . -name "*.txt"  # Find files by name
find . -type d        # Find directories
find . -mtime -7      # Modified in last 7 days
find . -size +10M     # Larger than 10MB
```

## Resources

- **[Linux Commands Reference](LINUX_COMMANDS.md)** - Complete command guide
- Linux man pages: `man <command>`
- TLDR pages: https://tldr.sh/
- Linux Journey: https://linuxjourney.com/
- The Linux Command Line book: https://linuxcommand.org/tlcl.php
- Bash Guide: https://mywiki.wooledge.org/BashGuide

## Tips for Learning

- **Use man pages** - `man ls` for detailed help
- **Practice daily** - Use the terminal for everyday tasks
- **Read error messages** - They tell you what's wrong
- **Tab completion** - Press Tab to autocomplete
- **Command history** - Use Up arrow or `history`
- **Aliases** - Create shortcuts for common commands
- **Don't copy-paste blindly** - Understand what commands do
- **Experiment safely** - Use a test directory
- **Learn keyboard shortcuts** - Ctrl+C, Ctrl+Z, Ctrl+R

## Common Keyboard Shortcuts

```bash
Ctrl + C    # Cancel/interrupt command
Ctrl + D    # Exit/logout (EOF)
Ctrl + Z    # Suspend process
Ctrl + R    # Search command history
Ctrl + A    # Go to line start
Ctrl + E    # Go to line end
Ctrl + U    # Clear line before cursor
Ctrl + K    # Clear line after cursor
Ctrl + L    # Clear screen
Tab         # Autocomplete
```

## Safety Tips

1. **Be careful with rm** - There's no recycle bin
2. **Use -i for confirmation** - `rm -i file.txt`
3. **Test destructive commands** - Use `echo` first
4. **Backup important data** - Before major changes
5. **Understand before executing** - Especially with sudo
6. **Check current directory** - `pwd` before deleting
7. **Use version control** - Git for important files

## Dangerous Commands (Be Careful!)

```bash
rm -rf /              # Deletes everything (DON'T RUN!)
chmod -R 777 /        # Makes everything writable (BAD!)
dd if=/dev/zero of=/dev/sda  # Wipes disk (DESTRUCTIVE!)

# Always double-check destructive commands
# Use sudo responsibly
```

## Exercises

See [exercises/README.md](exercises/README.md) for practice challenges.

## Next Steps

After completing this project:

- **Advanced shell scripting** - Complex automation
- **System administration** - Systemd, cron, logs
- **Security** - Firewalls, SELinux, AppArmor
- **Performance tuning** - Optimization techniques
- **Container tools** - Docker, podman
- **DevOps tools** - Ansible, Terraform
- **Linux certification** - LPIC, RHCSA

Happy learning!
MAINREADME

echo "✅ Created README.md"

# Topic directories and files...
cat > "$BASE/01-basics/README.md" << 'TOPIC1'
# Linux Basics

Master fundamental Linux commands for navigation and file management.

## Topics

- Filesystem navigation
- Listing files
- Creating directories
- Understanding paths
- Getting help

## Practice

```bash
bash practice.sh
```

## Key Commands

- `pwd` - Print working directory
- `ls` - List directory contents
- `cd` - Change directory
- `mkdir` - Make directory
- `man` - Manual pages
TOPIC1

cat > "$BASE/01-basics/practice.sh" << 'PRAC1'
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
PRAC1

chmod +x "$BASE/01-basics/practice.sh"

cat > "$BASE/02-file-operations/README.md" << 'TOPIC2'
# File Operations

Learn to manage files, permissions, and ownership.

## Topics

- Creating and viewing files
- Copying and moving
- Permissions (chmod)
- Ownership (chown)
- Finding files

## Practice

```bash
bash practice.sh
```
TOPIC2

cat > "$BASE/02-file-operations/practice.sh" << 'PRAC2'
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
PRAC2

chmod +x "$BASE/02-file-operations/practice.sh"

cat > "$BASE/03-text-processing/README.md" << 'TOPIC3'
# Text Processing

Master text manipulation and search tools.

## Topics

- grep - Search patterns
- sed - Stream editor
- awk - Text processing
- sort, uniq - Sorting and filtering
- cut - Extract columns

## Practice

```bash
bash practice.sh
```
TOPIC3

cat > "$BASE/03-text-processing/practice.sh" << 'PRAC3'
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
PRAC3

chmod +x "$BASE/03-text-processing/practice.sh"

cat > "$BASE/04-system/README.md" << 'TOPIC4'
# System Management

Learn process and system management.

## Topics

- Process management (ps, top, kill)
- System info (uname, df, free)
- User management
- System monitoring

## Practice

```bash
bash practice.sh
```
TOPIC4

cat > "$BASE/04-system/practice.sh" << 'PRAC4'
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
PRAC4

chmod +x "$BASE/04-system/practice.sh"

cat > "$BASE/05-networking/README.md" << 'TOPIC5'
# Networking

Learn network commands and tools.

## Topics

- Connectivity testing (ping)
- Network info (ip, ifconfig)
- File transfer (curl, wget, scp)
- Remote access (ssh)

## Practice

```bash
bash practice.sh
```
TOPIC5

cat > "$BASE/05-networking/practice.sh" << 'PRAC5'
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
PRAC5

chmod +x "$BASE/05-networking/practice.sh"

cat > "$BASE/06-scripting/README.md" << 'TOPIC6'
# Bash Scripting

Learn to write shell scripts.

## Topics

- Script structure
- Variables
- Control flow
- Functions
- Best practices

## Examples

Check the example scripts in this directory.
TOPIC6

cat > "$BASE/06-scripting/example.sh" << 'SCRIPT1'
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
SCRIPT1

chmod +x "$BASE/06-scripting/example.sh"

cat > "$BASE/exercises/README.md" << 'EXER'
# Linux Exercises

Practice your Linux skills!

## Beginner

1. Navigate to /tmp and create a directory called "practice"
2. Create 5 files named file1.txt through file5.txt
3. List all .txt files
4. Copy all files to a backup directory
5. Delete the practice directory

## Intermediate

6. Find all .txt files in your home directory
7. Count how many lines are in /etc/passwd
8. List all running processes and save to a file
9. Search for lines containing "root" in /etc/passwd
10. Create a script that prints all arguments passed to it

## Advanced

11. Write a script that backs up a directory with timestamp
12. Find all files larger than 100MB in /var
13. Monitor system resources and alert if disk usage > 80%
14. Create a script that automates user creation
15. Parse a log file and extract specific information

## Solutions

Try solving these yourself first! Use:
- `man <command>` for help
- Google/Stack Overflow when stuck
- Practice in a safe environment

Good luck!
EXER

echo ""
echo "✅ All files created!"
echo ""
echo "Project structure:"
ls -la "$BASE"
