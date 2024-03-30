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
