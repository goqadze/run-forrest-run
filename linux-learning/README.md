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
