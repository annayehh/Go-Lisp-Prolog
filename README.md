# Go-Lisp-Prolog
please install Go, lisp and prolog before running code
## Windows
### Go for windows
1. Go to https://golang.org/dl/.
2. Scroll down and expand “Archived versions”.
3. Expand “go1.17.4”.
4. Download “go1.17.4.windows-amd64.msi” (for x86 architecture, try “go1.17.4.windows-386.msi”)
5. Install following instructions.
6. Open cmd, then type go version, and you should see something like “go version go1.17.4 windows/amd64”

### LISP for Windows
1. Go to https://clisp.sourceforge.io/.
2. Click on “Cygwin” under “Get CLISP”
3. “Install Cygwin by running setup-x86 64.exe”
4. During installing, select “Install from Internet”
5. Choose any mirror.
6. In “Cygwin Setup - Select Packages”, search for “clisp” and then expand “All” and “Devel”, then change column “New” from “Skip” to “2.49-...”.
7. After install, double click “Cygwin64 Terminal” on desktop.
8. Type clisp --version, then you should see something like “GNU CLISP 2.49+ (2010-07-17)”
9. Type cd /cygdrive/c/, then you are in drive C.

### SWI-Prolog for Windows
1. Go to https://www.swi-prolog.org/download/stable?show=all
2. Search for “7.6.4”, then you can find SWI-Prolog 7.6.4 for Microsoft Windows (64 bit) and SWI-Prolog 7.6.4 for Microsoft Windows (32 bit). Use the one corresponding to your architecture
3. Download executable and install.
4. After install, there should be “SWI-Prolog” in Start menu.
5. You should be able to open a “.pl” file using SWI-Prolog. The title bar should show “SWI-Prolog (..., version 7.6.4)”.
6. Open cmd, then type "C:\Program Files\swipl\bin\swipl.exe" --version, and you should see something like “SWI-Prolog version 7.6.4 for x64-win64”

## macOS
### Go for macOS
1. Go to https://golang.org/dl/.
2. Scroll down and expand “Archived versions”.
3. Expand “go1.17.4”.
4. Download “go1.17.4.darwin-amd64.pkg” and Install.
5. Open terminal, type go version, and you should see something like “go version go1.17.4 darwin/amd64”

### LISP for macOS
1. Go to https://clisp.sourceforge.io/
2. Click on “MacPorts” under “Get CLISP”
3. Go to “installation section”, follow instructions.
4. Open a new terminal window, type port. You should see something like “MacPorts 2.6.2” (and enter shell mode).
• If the port cannot be found, quit terminal and open a new terminal, or try to use /opt/local/bin/port instead.
5. Quit the port shell mode (control + D), then type sudo port install clisp
6. Open a terminal and type clisp --version, and you should see something like “GNU CLISP 2.49 (2010-07-07) ...”

### SWI-Prolog for macOS
1. Go to https://www.swi-prolog.org/download/stable?show=all
2. Search for “7.6.4”, then you can find SWI-Prolog 7.6.4 for MacOSX 10.6 (Snow Leopard) and later on intel
3. Download executable and install.
4. Add the following line to your ∼/.zshrc file: export PATH=$PATH:/Applications/SWI-Prolog.app/Contents/MacOS
5. Open a terminal and type swipl --version, and you should see something like “SWI-Prolog version 7.6.4 ...”

## Linux
### Go for Linux
1. You need around 250 MiB free disk space to install
2. Go to https://golang.org/dl/
3. Scroll down and expand “Archived versions”
4. Expand “go1.17.4”
5. Download go1.17.4.linux-amd64.tar.gz
6. Extract the archive to /usr/local
```
$ ls go1.17.4.linux-amd64.tar.gz
go1.17.4.linux-amd64.tar.gz
$ pwd=$PWD
$ cd /usr/local
$ sudo tar xvf "$pwd/go1.17.4.linux-amd64.tar.gz"
...
$ /usr/local/go/bin/go version
go version go1.17.4 linux/amd64
```
7. You should add /usr/local/go/bin to your PATH.
```
$ echo 'PATH=/usr/local/go/bin:$PATH' >> ∼/.bashrc
$ echo 'PATH=/usr/local/go/bin:$PATH' >> ∼/.bash_profile
$ exec bash
$ go version
go version go1.17.4 linux/amd64
```
### LISP for Linux
1. In https://clisp.sourceforge.io/, see “Linux packages” in “Get CLISP”.
2. For RPM-based Linux, the following command installs the correct version of CLISP: sudo yum install "clisp-2.49.*"
3. For Debian-based Linux, the following command should be able to install the correct version of CLISP (though not tested):
```
sudo apt-get update
sudo apt-get install clisp
```
4. For Debian-based Linux, also try: https://unix.stackexchange.com/a/487556

### SWI-Prolog for Linux
1. On Ubuntu (≥ 18.04), you can install SWI-Prolog 7.6.4 using following command:
```
sudo apt-get update
sudo apt-get install swi-prolog
```
2. For other distributions it is possible to build from source: https://www.swi-prolog.org/build/unixautotools.txt
3. Reference: http://www.codecompiling.net/node/137
4. You need around 500 MiB free disk space to install
5. For RPM-based Linux, install prerequisites following https://www.swi-prolog.org/
```
build/Redhat.html
sudo dnf install \
cmake \
ninja-build \
libunwind \
gperftools-devel \
freetype-devel \
gmp-devel \
java-1.8.0-openjdk-devel \
jpackage-utils \
libICE-devel \
libjpeg-turbo-devel \
libSM-devel \
libX11-devel \
libXaw-devel \
libXext-devel \
libXft-devel \
libXinerama-devel \
libXmu-devel \
libXpm-devel \
libXrender-devel \
libXt-devel \
ncurses-devel \
openssl-devel \
pkgconfig \
readline-devel \
libedit-devel \
unixODBC-devel \
zlib-devel \
uuid-devel \
libarchive-devel \
libyaml-devel
```
6. For Debian-based Linux, install prerequisites following https://www.swi-prolog.org/build/Debian.html
```
sudo apt-get update
sudo apt-get install \
build-essential cmake ninja-build pkg-config \
ncurses-dev libreadline-dev libedit-dev \
libgoogle-perftools-dev \
libunwind-dev \
libgmp-dev \
libssl-dev \
unixodbc-dev \
zlib1g-dev libarchive-dev \
libossp-uuid-dev \
libxext-dev libice-dev libjpeg-dev libxinerama-dev libxft-dev \
libxpm-dev libxt-dev \
libdb-dev \
libpcre3-dev \
libyaml-dev \
default-jdk junit
```
7. It seems that the latest version of SWI-Prolog is using cmake (https://www.swi-prolog.org/build/unix.html), but not for SWI-Prolog 7.6.4)
8. Go to https://www.swi-prolog.org/download/stable?show=all
9. Search for “7.6.4”, then you can find and download SWI-Prolog source for 7.6.4
10. Execute the following
```
$ tar xvf swipl-7.6.4.tar.gz
...
$ cd swipl-7.6.4/
$ ./configure
$ make
...
$ sudo make install
...
$ # Now we need to install plunit, or we cannot run test cases
$ cd packages
$ ./configure
...
$ make
...
$ sudo make install
...
$ swipl --version
SWI-Prolog version 7.6.4 for x86_64-linux
```
