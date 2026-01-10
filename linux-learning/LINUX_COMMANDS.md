# Linux Commands Reference

A comprehensive reference for essential Linux commands. Master the command line with practical examples.

## Table of Contents
- [File System Navigation](#file-system-navigation)
- [File Operations](#file-operations)
- [Text Processing](#text-processing)
- [File Permissions](#file-permissions)
- [Process Management](#process-management)
- [System Information](#system-information)
- [User Management](#user-management)
- [Networking](#networking)
- [Package Management](#package-management)
- [Archiving & Compression](#archiving--compression)
- [Disk Management](#disk-management)
- [Quick Tips](#quick-tips)

---

## File System Navigation

### `pwd`
Print working directory - shows your current location.

**Examples:**
```bash
pwd
# Output: /home/user/documents
```

---

### `ls`
List directory contents.

**Syntax:**
```bash
ls [OPTIONS] [PATH]
```

**Examples:**
```bash
# Basic list
ls

# Long format with details
ls -l

# Include hidden files
ls -a

# Long format + hidden files
ls -la

# Human-readable sizes
ls -lh

# Sort by modification time
ls -lt

# Reverse sort
ls -lr

# Recursive listing
ls -R

# List only directories
ls -d */

# With file type indicators
ls -F
```

**Common Flags:**
- `-l` - Long format (permissions, owner, size, date)
- `-a` - Show all files (including hidden .-files)
- `-h` - Human-readable sizes (K, M, G)
- `-t` - Sort by modification time
- `-r` - Reverse order
- `-R` - Recursive
- `-S` - Sort by file size

---

### `cd`
Change directory.

**Examples:**
```bash
# Go to specific directory
cd /usr/local/bin

# Go to home directory
cd
cd ~

# Go to previous directory
cd -

# Go up one level
cd ..

# Go up two levels
cd ../..

# Relative path
cd documents/projects
```

---

### `tree`
Display directory structure as a tree.

**Examples:**
```bash
# Show tree
tree

# Limit depth
tree -L 2

# Show hidden files
tree -a

# Directories only
tree -d

# With file sizes
tree -h
```

---

## File Operations

### `touch`
Create empty file or update timestamp.

**Examples:**
```bash
# Create new file
touch file.txt

# Create multiple files
touch file1.txt file2.txt file3.txt

# Update timestamp
touch existing-file.txt
```

---

### `cat`
Concatenate and display file contents.

**Examples:**
```bash
# View file
cat file.txt

# View multiple files
cat file1.txt file2.txt

# Number lines
cat -n file.txt

# Create file with content
cat > newfile.txt
# Type content, then Ctrl+D to save

# Append to file
cat >> file.txt
```

---

### `less` / `more`
View file contents page by page.

**Examples:**
```bash
# View file (scrollable)
less file.txt

# Search: /pattern
# Next: n, Previous: N
# Quit: q

# More (simpler pager)
more file.txt
```

**less Navigation:**
- `Space` - Next page
- `b` - Previous page
- `/pattern` - Search forward
- `?pattern` - Search backward
- `q` - Quit

---

### `head` / `tail`
View beginning or end of files.

**Examples:**
```bash
# First 10 lines (default)
head file.txt

# First 20 lines
head -n 20 file.txt
head -20 file.txt

# Last 10 lines
tail file.txt

# Last 20 lines
tail -n 20 file.txt

# Follow file (live updates)
tail -f logfile.txt

# Last 100 lines and follow
tail -n 100 -f logfile.txt
```

---

### `cp`
Copy files and directories.

**Examples:**
```bash
# Copy file
cp source.txt destination.txt

# Copy to directory
cp file.txt /path/to/directory/

# Copy multiple files
cp file1.txt file2.txt directory/

# Copy directory recursively
cp -r source-dir/ destination-dir/

# Interactive (confirm overwrite)
cp -i file.txt destination.txt

# Preserve attributes
cp -p file.txt destination.txt

# Verbose output
cp -v file.txt destination.txt
```

**Common Flags:**
- `-r` - Recursive (for directories)
- `-i` - Interactive (confirm)
- `-p` - Preserve attributes
- `-v` - Verbose
- `-u` - Update (copy only if newer)

---

### `mv`
Move or rename files.

**Examples:**
```bash
# Rename file
mv oldname.txt newname.txt

# Move file
mv file.txt /path/to/directory/

# Move multiple files
mv file1.txt file2.txt directory/

# Move directory
mv old-dir/ new-location/

# Interactive
mv -i file.txt destination.txt

# Don't overwrite
mv -n file.txt destination.txt
```

---

### `rm`
Remove files and directories.

**Examples:**
```bash
# Remove file
rm file.txt

# Remove multiple files
rm file1.txt file2.txt

# Interactive confirmation
rm -i file.txt

# Force removal
rm -f file.txt

# Remove directory and contents
rm -r directory/

# Force remove directory
rm -rf directory/

# Verbose output
rm -v file.txt

# Remove with pattern
rm *.tmp
```

**⚠️ Warning:** `rm` is permanent - there's no recycle bin!

**Common Flags:**
- `-r` - Recursive (for directories)
- `-f` - Force (no confirmation)
- `-i` - Interactive
- `-v` - Verbose

---

### `mkdir`
Make directories.

**Examples:**
```bash
# Create directory
mkdir mydir

# Create multiple directories
mkdir dir1 dir2 dir3

# Create parent directories
mkdir -p path/to/deep/directory

# Set permissions
mkdir -m 755 mydir

# Verbose
mkdir -v mydir
```

---

### `rmdir`
Remove empty directories.

**Examples:**
```bash
# Remove empty directory
rmdir mydir

# Remove parent directories
rmdir -p path/to/empty/dirs
```

---

### `find`
Search for files and directories.

**Examples:**
```bash
# Find by name
find . -name "*.txt"

# Case-insensitive name
find . -iname "*.TXT"

# Find directories
find . -type d -name "test"

# Find files
find . -type f -name "*.log"

# Find by size
find . -size +10M          # Larger than 10MB
find . -size -1M           # Smaller than 1MB

# Find by modification time
find . -mtime -7           # Modified in last 7 days
find . -mtime +30          # Modified more than 30 days ago

# Find and execute
find . -name "*.tmp" -delete
find . -name "*.txt" -exec cat {} \;

# Find empty files
find . -type f -empty

# Find by permissions
find . -perm 644
```

---

### `locate`
Find files by name (uses database).

**Examples:**
```bash
# Find file
locate filename

# Case-insensitive
locate -i filename

# Update database
sudo updatedb

# Limit results
locate -n 10 filename
```

---

## Text Processing

### `grep`
Search for patterns in text.

**Examples:**
```bash
# Search in file
grep "pattern" file.txt

# Case-insensitive
grep -i "pattern" file.txt

# Recursive search
grep -r "pattern" directory/

# Line numbers
grep -n "pattern" file.txt

# Invert match (lines NOT matching)
grep -v "pattern" file.txt

# Count matches
grep -c "pattern" file.txt

# Multiple patterns
grep -e "pattern1" -e "pattern2" file.txt

# Extended regex
grep -E "pattern1|pattern2" file.txt

# Show context
grep -C 3 "pattern" file.txt  # 3 lines before and after
grep -A 2 "pattern" file.txt  # 2 lines after
grep -B 2 "pattern" file.txt  # 2 lines before

# Word match
grep -w "word" file.txt

# Files with matches
grep -l "pattern" *.txt

# Files without matches
grep -L "pattern" *.txt
```

**Common Flags:**
- `-i` - Case-insensitive
- `-r` - Recursive
- `-n` - Line numbers
- `-v` - Invert (show non-matches)
- `-c` - Count matches
- `-w` - Whole word
- `-l` - Files with matches

---

### `sed`
Stream editor for text transformation.

**Examples:**
```bash
# Replace first occurrence
sed 's/old/new/' file.txt

# Replace all occurrences
sed 's/old/new/g' file.txt

# Replace in-place
sed -i 's/old/new/g' file.txt

# Delete lines
sed '/pattern/d' file.txt

# Print specific lines
sed -n '10,20p' file.txt

# Multiple operations
sed -e 's/old1/new1/g' -e 's/old2/new2/g' file.txt

# Insert line
sed '5i\New line' file.txt

# Append line
sed '5a\New line' file.txt
```

---

### `awk`
Text processing and data extraction.

**Examples:**
```bash
# Print column
awk '{print $1}' file.txt

# Print multiple columns
awk '{print $1, $3}' file.txt

# Pattern matching
awk '/pattern/ {print $0}' file.txt

# Sum column
awk '{sum+=$2} END {print sum}' file.txt

# Custom field separator
awk -F: '{print $1}' /etc/passwd

# Conditional
awk '$3 > 100 {print $1}' file.txt

# Built-in variables
awk '{print NR, $0}' file.txt  # Line number
awk '{print NF}' file.txt      # Number of fields
```

---

### `cut`
Extract columns from text.

**Examples:**
```bash
# Cut by character position
cut -c 1-5 file.txt

# Cut by field (default delimiter: tab)
cut -f 1,3 file.txt

# Custom delimiter
cut -d: -f1 /etc/passwd

# Range of fields
cut -d, -f1-3 file.csv
```

---

### `sort`
Sort lines of text.

**Examples:**
```bash
# Sort alphabetically
sort file.txt

# Reverse sort
sort -r file.txt

# Numeric sort
sort -n numbers.txt

# Sort by column
sort -k 2 file.txt

# Unique values only
sort -u file.txt

# Case-insensitive
sort -f file.txt

# Sort by size (human-readable)
sort -h sizes.txt
```

---

### `uniq`
Remove or report duplicate lines.

**Examples:**
```bash
# Remove duplicates (must be sorted first)
sort file.txt | uniq

# Count occurrences
uniq -c file.txt

# Show only duplicates
uniq -d file.txt

# Show only unique lines
uniq -u file.txt
```

---

### `wc`
Word, line, character, and byte count.

**Examples:**
```bash
# Count everything
wc file.txt

# Lines only
wc -l file.txt

# Words only
wc -w file.txt

# Characters only
wc -m file.txt

# Bytes only
wc -c file.txt

# Multiple files
wc -l *.txt
```

---

### `diff`
Compare files line by line.

**Examples:**
```bash
# Show differences
diff file1.txt file2.txt

# Side-by-side
diff -y file1.txt file2.txt

# Unified format (like git diff)
diff -u file1.txt file2.txt

# Ignore whitespace
diff -w file1.txt file2.txt

# Brief output
diff -q file1.txt file2.txt
```

---

## File Permissions

### `chmod`
Change file permissions.

**Examples:**
```bash
# Numeric mode
chmod 644 file.txt   # rw-r--r--
chmod 755 script.sh  # rwxr-xr-x
chmod 600 secret.txt # rw-------

# Symbolic mode
chmod u+x script.sh  # Add execute for user
chmod g-w file.txt   # Remove write for group
chmod o+r file.txt   # Add read for others
chmod a+x script.sh  # Add execute for all

# Recursive
chmod -R 755 directory/

# Copy permissions
chmod --reference=file1 file2
```

**Permission Numbers:**
- `4` - Read (r)
- `2` - Write (w)
- `1` - Execute (x)
- `7` - rwx (4+2+1)
- `6` - rw- (4+2)
- `5` - r-x (4+1)
- `0` - ---

---

### `chown`
Change file owner and group.

**Examples:**
```bash
# Change owner
sudo chown user file.txt

# Change owner and group
sudo chown user:group file.txt

# Recursive
sudo chown -R user:group directory/

# Change group only
sudo chown :group file.txt
```

---

### `chgrp`
Change group ownership.

**Examples:**
```bash
# Change group
chgrp group file.txt

# Recursive
chgrp -R group directory/
```

---

## Process Management

### `ps`
Display process status.

**Examples:**
```bash
# Show user processes
ps

# All processes
ps aux

# Process tree
ps auxf

# Specific user
ps -u username

# By process name
ps aux | grep process-name

# Custom format
ps -eo pid,user,command
```

**Common Options:**
- `a` - All users
- `u` - User-oriented format
- `x` - Include processes without TTY
- `f` - Full format

---

### `top`
Display and update sorted process information.

**Examples:**
```bash
# Interactive process viewer
top

# Specific user
top -u username

# Single iteration
top -n 1
```

**top Commands:**
- `q` - Quit
- `k` - Kill process
- `M` - Sort by memory
- `P` - Sort by CPU
- `h` - Help

---

### `htop`
Interactive process viewer (if installed).

**Examples:**
```bash
# Launch htop
htop

# Filter by user
htop -u username
```

---

### `kill`
Terminate processes.

**Examples:**
```bash
# Kill by PID
kill 1234

# Force kill
kill -9 1234
kill -SIGKILL 1234

# Graceful termination
kill -15 1234
kill -SIGTERM 1234

# Kill by name
killall process-name

# Kill all user processes
killall -u username
```

**Common Signals:**
- `1` (HUP) - Hangup
- `2` (INT) - Interrupt
- `9` (KILL) - Force kill
- `15` (TERM) - Terminate gracefully

---

### `jobs`
List background jobs.

**Examples:**
```bash
# List jobs
jobs

# Detailed list
jobs -l

# Bring job to foreground
fg %1

# Send job to background
bg %1

# Start process in background
command &
```

---

### `nohup`
Run command immune to hangups.

**Examples:**
```bash
# Run and ignore hangup
nohup command &

# Redirect output
nohup command > output.log 2>&1 &
```

---

## System Information

### `uname`
Print system information.

**Examples:**
```bash
# System name
uname

# All information
uname -a

# Kernel name
uname -s

# Kernel version
uname -r

# Machine hardware
uname -m

# Operating system
uname -o
```

---

### `df`
Disk space usage.

**Examples:**
```bash
# Show disk space
df

# Human-readable
df -h

# Specific filesystem
df -h /home

# Inode information
df -i

# Exclude type
df -x tmpfs
```

---

### `du`
Disk usage of files and directories.

**Examples:**
```bash
# Current directory
du

# Human-readable
du -h

# Summary of directory
du -sh directory/

# All files
du -ah

# Sort by size
du -sh * | sort -h

# Limit depth
du -h --max-depth=1
```

---

### `free`
Display memory usage.

**Examples:**
```bash
# Show memory
free

# Human-readable
free -h

# In MB
free -m

# In GB
free -g

# Continuous display
free -h -s 2  # Update every 2 seconds
```

---

### `uptime`
Show system uptime and load.

**Examples:**
```bash
uptime

# Pretty format
uptime -p

# Since when
uptime -s
```

---

### `hostname`
Show or set system hostname.

**Examples:**
```bash
# Show hostname
hostname

# Show FQDN
hostname -f

# Show IP address
hostname -I
```

---

### `date`
Display or set date and time.

**Examples:**
```bash
# Current date and time
date

# Custom format
date "+%Y-%m-%d %H:%M:%S"

# ISO 8601
date -I

# Unix timestamp
date +%s

# Convert timestamp
date -d @1609459200
```

---

## User Management

### `whoami`
Print current user.

**Examples:**
```bash
whoami
```

---

### `who`
Show logged in users.

**Examples:**
```bash
# All users
who

# Only me
who am i

# Detailed
who -a
```

---

### `id`
Print user and group IDs.

**Examples:**
```bash
# Current user
id

# Specific user
id username

# Just UID
id -u

# Just GID
id -g

# All groups
id -G
```

---

### `sudo`
Execute command as superuser.

**Examples:**
```bash
# Run as root
sudo command

# Run as specific user
sudo -u username command

# Become root
sudo -i
sudo su -

# Keep environment
sudo -E command

# List privileges
sudo -l
```

---

### `useradd` / `adduser`
Create new user.

**Examples:**
```bash
# Create user
sudo useradd username

# With home directory
sudo useradd -m username

# With specific shell
sudo useradd -s /bin/bash username

# With groups
sudo useradd -G sudo,docker username
```

---

### `passwd`
Change user password.

**Examples:**
```bash
# Change own password
passwd

# Change user password
sudo passwd username

# Lock account
sudo passwd -l username

# Unlock account
sudo passwd -u username
```

---

## Networking

### `ping`
Test network connectivity.

**Examples:**
```bash
# Ping host
ping google.com

# Limit count
ping -c 5 google.com

# Specific interval
ping -i 2 google.com

# IPv6
ping6 google.com
```

---

### `curl`
Transfer data from URLs.

**Examples:**
```bash
# GET request
curl https://example.com

# Save to file
curl -o file.html https://example.com

# Follow redirects
curl -L https://example.com

# POST request
curl -X POST -d "key=value" https://api.example.com

# With headers
curl -H "Content-Type: application/json" https://api.example.com

# Show headers
curl -I https://example.com

# Authentication
curl -u user:pass https://api.example.com

# Upload file
curl -F "file=@path/to/file" https://upload.example.com
```

---

### `wget`
Download files from the web.

**Examples:**
```bash
# Download file
wget https://example.com/file.zip

# Save with different name
wget -O newname.zip https://example.com/file.zip

# Resume download
wget -c https://example.com/largefile.iso

# Download in background
wget -b https://example.com/file.zip

# Recursive download
wget -r https://example.com/

# Limit rate
wget --limit-rate=1m https://example.com/file.zip
```

---

### `ssh`
Secure shell remote login.

**Examples:**
```bash
# Connect to host
ssh user@hostname

# Specific port
ssh -p 2222 user@hostname

# With key
ssh -i ~/.ssh/id_rsa user@hostname

# Execute command
ssh user@hostname 'ls -la'

# Port forwarding
ssh -L 8080:localhost:80 user@hostname
```

---

### `scp`
Secure copy over SSH.

**Examples:**
```bash
# Copy to remote
scp file.txt user@host:/path/to/destination

# Copy from remote
scp user@host:/path/to/file.txt .

# Copy directory
scp -r directory/ user@host:/path/

# Specific port
scp -P 2222 file.txt user@host:/path/
```

---

### `rsync`
Efficient file transfer and synchronization.

**Examples:**
```bash
# Sync directories
rsync -av source/ destination/

# Over SSH
rsync -av source/ user@host:/path/

# Show progress
rsync -av --progress source/ destination/

# Delete files in destination
rsync -av --delete source/ destination/

# Dry run
rsync -av --dry-run source/ destination/

# Exclude files
rsync -av --exclude='*.log' source/ destination/
```

---

### `netstat`
Network statistics.

**Examples:**
```bash
# All connections
netstat -a

# Listening ports
netstat -l

# TCP connections
netstat -t

# UDP connections
netstat -u

# With process info
netstat -p

# Routing table
netstat -r

# Continuous display
netstat -c
```

---

### `ss`
Socket statistics (modern netstat).

**Examples:**
```bash
# All sockets
ss -a

# Listening sockets
ss -l

# TCP sockets
ss -t

# With process
ss -p

# Summary
ss -s
```

---

## Package Management

### APT (Debian/Ubuntu)

```bash
# Update package list
sudo apt update

# Upgrade packages
sudo apt upgrade

# Install package
sudo apt install package-name

# Remove package
sudo apt remove package-name

# Search packages
apt search keyword

# Show package info
apt show package-name

# Clean cache
sudo apt clean
```

### YUM/DNF (RedHat/CentOS/Fedora)

```bash
# Install package
sudo yum install package-name
sudo dnf install package-name

# Update packages
sudo yum update
sudo dnf update

# Remove package
sudo yum remove package-name

# Search packages
yum search keyword
```

### Homebrew (macOS)

```bash
# Install package
brew install package-name

# Update Homebrew
brew update

# Upgrade packages
brew upgrade

# Search packages
brew search keyword
```

---

## Archiving & Compression

### `tar`
Archive files.

**Examples:**
```bash
# Create archive
tar -czf archive.tar.gz directory/

# Extract archive
tar -xzf archive.tar.gz

# List contents
tar -tzf archive.tar.gz

# Extract to specific directory
tar -xzf archive.tar.gz -C /path/to/destination

# Create bzip2 archive
tar -cjf archive.tar.bz2 directory/

# Verbose
tar -czvf archive.tar.gz directory/
```

**Common Flags:**
- `c` - Create
- `x` - Extract
- `z` - gzip compression
- `j` - bzip2 compression
- `f` - File
- `v` - Verbose
- `t` - List

---

### `gzip` / `gunzip`
Compress files.

**Examples:**
```bash
# Compress file
gzip file.txt

# Decompress
gunzip file.txt.gz

# Keep original
gzip -k file.txt

# Recursive
gzip -r directory/

# Best compression
gzip -9 file.txt
```

---

### `zip` / `unzip`
ZIP archives.

**Examples:**
```bash
# Create zip
zip archive.zip file1.txt file2.txt

# Create from directory
zip -r archive.zip directory/

# Extract
unzip archive.zip

# List contents
unzip -l archive.zip

# Extract to directory
unzip archive.zip -d /path/to/destination
```

---

## Disk Management

### `mount`
Mount filesystems.

**Examples:**
```bash
# Show mounts
mount

# Mount device
sudo mount /dev/sdb1 /mnt

# Unmount
sudo umount /mnt

# Show disk labels
sudo blkid
```

---

### `fdisk`
Disk partitioning.

**Examples:**
```bash
# List disks
sudo fdisk -l

# Partition disk
sudo fdisk /dev/sdb
```

---

## Quick Tips

### Aliases
```bash
# Add to ~/.bashrc or ~/.zshrc
alias ll='ls -la'
alias ..='cd ..'
alias ...='cd ../..'
alias grep='grep --color=auto'
```

### History
```bash
# Show history
history

# Execute previous command
!!

# Execute command from history
!123

# Search history
Ctrl+R

# Clear history
history -c
```

### Redirects and Pipes
```bash
# Redirect output
command > file.txt

# Append output
command >> file.txt

# Redirect errors
command 2> errors.txt

# Redirect both
command > output.txt 2>&1

# Pipe to another command
command1 | command2

# Pipe and save
command | tee file.txt
```

### Command Substitution
```bash
# Using $()
result=$(command)
echo "Result: $result"

# In commands
cd $(dirname $0)

# Backticks (old style)
result=`command`
```

---

## See Also

- [README.md](README.md) - Project learning guide
- Man pages: `man <command>`
- TLDR pages: https://tldr.sh/
- Linux Journey: https://linuxjourney.com/

Happy learning!
