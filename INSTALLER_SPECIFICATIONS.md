# Athera Installer Specifications

**Version:** 1.0  
**Format:** Platform-specific installation specifications  
**Purpose:** Technical details for building installers

---

## Overview

Athera installers should be:
- **Easy to use** - Download → Click → Done
- **Professional** - Signed, notarized, verified
- **Self-contained** - No external dependencies
- **Reversible** - Easy uninstall
- **Silent mode** - For CI/CD automation

---

## macOS Installer

### Format: PKG (macOS Package)

**File:** `athera-1.0.0.pkg` (~25-30 MB)

### Build Process

**Requirements:**
- Apple Developer ID certificate (for signing)
- macOS 10.13+
- codesign utility
- productbuild utility

**Package Structure:**
```
athera-1.0.0.pkg
├── Contents/
│   ├── Packages/
│   │   └── Athera.pkg (the actual payload)
│   └── Distribution (XML metadata)
```

**Payload Contents:**
```
Athera.pkg/
├── Scripts/       (pre/post install scripts)
│   ├── preinstall    (check system, create backup)
│   ├── postinstall   (set permissions, register)
│   └── postuninstall (cleanup)
├── Payload/       (files to install)
│   ├── usr/
│   │   └── local/
│   │       └── bin/
│   │           └── athera (binary, 24 MB)
│   └── Library/
│       └── LaunchDaemons/  (if needed for services)
└── PackageInfo    (metadata)
```

### Installation Details

**Installation Location:**
- Binary: `/usr/local/bin/athera`
- Stdlib: Embedded in binary
- Config: `~/.athera/config`
- Cache: `~/.athera/cache`
- Man page: `/usr/local/share/man/man1/athera.1`

**Signing & Notarization:**

```bash
# Sign the binary
codesign -s "Developer ID Application: Company Name" \
         --timestamp \
         --options=runtime \
         athera

# Verify signature
codesign -v athera

# Notarize (required for macOS 10.15+)
xcrun altool --notarize-app \
    -f athera-1.0.0.pkg \
    -t osx \
    --team-id "XXXXXXXXXX" \
    --username "apple@company.com" \
    --password "@keychain:AC_PASSWORD"

# Staple notarization
xcrun stapler staple athera-1.0.0.pkg

# Verify notarization
spctl -a -v -t install athera-1.0.0.pkg
```

**Installation Script:**

```bash
#!/bin/sh
# preinstall script

# Check if Athera is already installed
if [ -f /usr/local/bin/athera ]; then
    # Backup existing version
    cp /usr/local/bin/athera /usr/local/bin/athera.bak
fi

# Create necessary directories
mkdir -p ~/.athera/cache
mkdir -p ~/.athera/modules

exit 0
```

```bash
#!/bin/sh
# postinstall script

# Make binary executable
chmod +x /usr/local/bin/athera

# Install man page
mkdir -p /usr/local/share/man/man1
cp athera.1 /usr/local/share/man/man1/athera.1
chmod 644 /usr/local/share/man/man1/athera.1

# Create symlink for easy access
ln -sf /usr/local/bin/athera /usr/local/bin/ath 2>/dev/null || true

# Verify installation
/usr/local/bin/athera --version

exit 0
```

**Distribution XML:**

```xml
<?xml version="1.0" encoding="utf-8"?>
<installer-gui-script minSpecVersion="1">
    <title>Athera 1.0.0</title>
    <organization>Athera Project</organization>
    <options customize="never" require-scripts="false" 
             rootVolumeOnly="true" hostArchitectures="x86_64,arm64"/>
    
    <domains enable_localSystem="true"/>
    
    <installation-check script="check_version()"/>
    
    <script>
        function check_version() {
            return true;
        }
    </script>
    
    <pkg-ref id="athera.pkg" onConclusion="RequireRestart">Athera.pkg</pkg-ref>
    
    <choice id="athera.pkg" 
            title="Athera" 
            description="Athera Programming Language"
            start_selected="true">
        <pkg-ref id="athera.pkg"/>
    </choice>
    
    <choice-group id="installation">
        <choice id="default" title="Easy Install" description="Install for current user">
            <pkg-ref id="athera.pkg"/>
        </choice>
    </choice-group>
</installer-gui-script>
```

### Building macOS Installer

```bash
#!/bin/bash
set -e

VERSION="1.0.0"
BUILD_DIR="build/macos"
OUTPUT_DIR="dist"

mkdir -p "$BUILD_DIR/Athera.pkg/Scripts"
mkdir -p "$BUILD_DIR/Athera.pkg/Payload/usr/local/bin"
mkdir -p "$BUILD_DIR/Athera.pkg/Payload/Library/LaunchDaemons"

# Copy binary
cp build/athera "$BUILD_DIR/Athera.pkg/Payload/usr/local/bin/"

# Copy scripts
cp installer/macos/preinstall "$BUILD_DIR/Athera.pkg/Scripts/"
cp installer/macos/postinstall "$BUILD_DIR/Athera.pkg/Scripts/"
chmod +x "$BUILD_DIR/Athera.pkg/Scripts/"*

# Create PackageInfo
cat > "$BUILD_DIR/Athera.pkg/PackageInfo" <<EOF
<?xml version="1.0" encoding="utf-8"?>
<pkg-info format-version="2" identifier="com.athera.athera" version="$VERSION"
          install-location="/" auth="root">
    <payload installKBytes="26000"/>
    <scripts>
        <preinstall file="./preinstall"/>
        <postinstall file="./postinstall"/>
    </scripts>
</pkg-info>
EOF

# Build distribution package
productbuild \
    --identifier com.athera.athera \
    --distribution installer/macos/Distribution.xml \
    --resources installer/macos/resources \
    --package-path "$BUILD_DIR" \
    "$OUTPUT_DIR/athera-$VERSION.pkg"

# Sign
codesign -s "Developer ID Installer" "$OUTPUT_DIR/athera-$VERSION.pkg"

# Notarize (requires developer account)
echo "Notarizing package (requires Apple Developer account)..."
xcrun altool --notarize-app \
    -f "$OUTPUT_DIR/athera-$VERSION.pkg" \
    -t osx \
    --team-id "XXXXXXXXXX" \
    --username "$APPLE_USERNAME" \
    --password "$APPLE_PASSWORD"

echo "✓ macOS installer created: $OUTPUT_DIR/athera-$VERSION.pkg"
```

### User Experience (macOS)

```
1. Download athera-1.0.0.pkg
2. Double-click installer
3. Installation wizard appears
   - Welcome screen
   - License agreement
   - Installation location choice
   - Summary screen
4. Click "Install"
5. Enter admin password (if prompted)
6. "Installation successful" message
7. Close installer
8. Open Terminal, type: athera --version
```

---

## Windows Installer

### Format: MSI (Windows Installer Package)

**File:** `athera-1.0.0-setup.msi` (~25-30 MB)

**Also Provide:** `athera-1.0.0-installer.exe` (wrapper for user convenience)

### Build Process

**Requirements:**
- WiX Toolset 3.x
- Visual Studio or WiX compiler
- Signing certificate (optional but recommended)

**WiX Project Structure:**

```xml
<!-- Product.wxs -->
<?xml version="1.0" encoding="UTF-8"?>
<Wix xmlns="http://schemas.microsoft.com/wix/2006/wi">
    <Product Id="*" 
             Name="Athera" 
             Language="1033" 
             Version="1.0.0.0"
             Manufacturer="Athera Project" 
             UpgradeCode="12345678-1234-1234-1234-123456789012">
        
        <Package InstallerVersion="200" 
                 Compressed="yes" 
                 InstallScope="perMachine"
                 Platform="x64"/>
        
        <Media Id="1" Cabinet="Athera.cab" EmbedCab="yes"/>
        
        <!-- License and UI -->
        <UIRef Id="WixUI_InstallDir"/>
        <WixVariable Id="WixUILicenseRtf" 
                     Value="installer\windows\License.rtf"/>
        <WixVariable Id="WixUIBannerBmp" 
                     Value="installer\windows\banner.bmp"/>
        
        <!-- Installation directory -->
        <Directory Id="TARGETDIR" Name="SourceDir">
            <Directory Id="ProgramFilesFolder">
                <Directory Id="INSTALLFOLDER" Name="Athera"/>
            </Directory>
            <Directory Id="ProgramMenuFolder">
                <Directory Id="AtheraMenuFolder" Name="Athera"/>
            </Directory>
            <Directory Id="EnvironmentFolder" Name="."/>
        </Directory>
        
        <!-- Features -->
        <Feature Id="ProductFeature" Title="Athera Runtime" Level="1">
            <ComponentRef Id="AtheraExecutable"/>
            <ComponentRef Id="AtheraConfig"/>
            <ComponentRef Id="PathEnvironment"/>
        </Feature>
        
        <!-- Components -->
        <DirectoryRef Id="INSTALLFOLDER">
            <Component Id="AtheraExecutable" Guid="*" Win64="yes">
                <File Id="AtheraExe" Source="build\athera.exe" 
                      KeyPath="yes" Checksum="yes"/>
            </Component>
            
            <Component Id="AtheraConfig" Guid="*">
                <RegistryKey Root="HKCU" 
                            Key="Software\Athera\Athera"
                            Action="createAndRemoveOnUninstall">
                    <RegistryValue Type="string" 
                                  Name="InstallPath"
                                  Value="[INSTALLFOLDER]"/>
                </RegistryKey>
            </Component>
        </DirectoryRef>
        
        <!-- Add to PATH -->
        <DirectoryRef Id="INSTALLFOLDER">
            <Component Id="PathEnvironment" Guid="*">
                <Environment Id="PATH" 
                            Name="PATH" 
                            Value="[INSTALLFOLDER]" 
                            Permanent="no" 
                            Part="last" 
                            Action="set" 
                            System="yes"/>
            </Component>
        </DirectoryRef>
        
    </Product>
</Wix>
```

### Installation Details

**Installation Locations:**
- Binary: `C:\Program Files\Athera\athera.exe`
- Stdlib: Embedded in binary
- Config: `%APPDATA%\Athera\config`
- Cache: `%APPDATA%\Athera\cache`
- Start Menu: `C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Athera`

**Path Configuration:**
- Automatically adds `C:\Program Files\Athera\` to PATH
- Available in PowerShell, CMD, Git Bash, WSL

**Signing:**

```powershell
# Sign the MSI
signtool.exe sign /f "certificate.pfx" `
    /p "password" `
    /t http://timestamp.authority.com `
    /d "Athera Programming Language" `
    "athera-1.0.0-setup.msi"

# Verify signature
signtool.exe verify /pa "athera-1.0.0-setup.msi"
```

### Building Windows Installer

```batch
@echo off
REM build_windows_installer.bat

set VERSION=1.0.0
set BUILD_DIR=build\windows
set OUTPUT_DIR=dist

REM Compile WiX source
candle.exe -out %BUILD_DIR%\ ^
    -d ProductVersion=%VERSION% ^
    -d SourcePath=build\athera.exe ^
    installer\windows\Product.wxs

REM Link to MSI
light.exe -out %OUTPUT_DIR%\athera-%VERSION%-setup.msi ^
    %BUILD_DIR%\Product.wixobj

REM Sign MSI
signtool.exe sign /f "cert.pfx" /p "password" ^
    /t http://timestamp.authority.com ^
    /d "Athera Programming Language" ^
    %OUTPUT_DIR%\athera-%VERSION%-setup.msi

echo ✓ Windows installer created: %OUTPUT_DIR%\athera-%VERSION%-setup.msi
```

### User Experience (Windows)

```
1. Download athera-1.0.0-installer.exe
2. Double-click installer
3. Windows UAC dialog (if admin install)
4. Installation wizard appears
   - Welcome screen
   - License agreement
   - Installation path choice (default: C:\Program Files\Athera)
   - Features selection
   - Ready to install screen
5. Click "Install"
6. Installation progress bar
7. "Installation successful" message
8. Option to launch Athera
9. Close installer
10. Open PowerShell/CMD, type: athera --version
```

### Uninstall (Windows)

**Method 1: Add/Remove Programs**
```
Settings → Apps → Apps & Features
Find "Athera" → Click → Uninstall
```

**Method 2: Command Line**
```batch
msiexec /x athera-1.0.0-setup.msi /quiet
```

---

## Linux Installer

### Format 1: DEB (Ubuntu/Debian)

**File:** `athera_1.0.0_amd64.deb` (~25-30 MB)

**Build Process:**

```
athera_1.0.0_amd64/
├── DEBIAN/
│   ├── control          (package metadata)
│   ├── postinst         (post-install script)
│   ├── prerm            (pre-remove script)
│   ├── postrm           (post-remove script)
│   └── conffiles        (config files list)
├── usr/
│   ├── bin/
│   │   └── athera       (binary, 24 MB)
│   ├── share/
│   │   ├── man/man1/
│   │   │   └── athera.1 (man page)
│   │   ├── doc/athera/
│   │   │   ├── LICENSE
│   │   │   └── README.md
│   │   └── applications/
│   │       └── athera.desktop
│   └── lib/athera/      (optional libraries)
└── etc/
    └── athera/          (system config)
```

**Control File:**

```
Package: athera
Version: 1.0.0
Architecture: amd64
Maintainer: Athera Project <team@athera.dev>
Homepage: https://athera.dev
Description: Athera Programming Language
 A professional programming language designed for clarity and safety.
 Execute .ath files globally with the 'athera' command.
Depends: libc6 (>= 2.31)
Suggests: git
Conflicts: athera-dev
```

**Build Script:**

```bash
#!/bin/bash
set -e

VERSION="1.0.0"
BUILD_DIR="build/deb"
OUTPUT_DIR="dist"

# Create directory structure
mkdir -p "$BUILD_DIR/DEBIAN"
mkdir -p "$BUILD_DIR/usr/bin"
mkdir -p "$BUILD_DIR/usr/share/man/man1"
mkdir -p "$BUILD_DIR/usr/share/doc/athera"

# Copy binary
cp build/athera "$BUILD_DIR/usr/bin/"
chmod +x "$BUILD_DIR/usr/bin/athera"

# Copy man page
cp installer/linux/athera.1 "$BUILD_DIR/usr/share/man/man1/"
gzip "$BUILD_DIR/usr/share/man/man1/athera.1"

# Copy documentation
cp LICENSE "$BUILD_DIR/usr/share/doc/athera/"
cp README.md "$BUILD_DIR/usr/share/doc/athera/"
cp CHANGELOG.md "$BUILD_DIR/usr/share/doc/athera/"

# Create control file
cat > "$BUILD_DIR/DEBIAN/control" <<EOF
Package: athera
Version: $VERSION
Architecture: amd64
Maintainer: Athera Project <team@athera.dev>
Homepage: https://athera.dev
Description: Athera Programming Language
 A professional programming language designed for clarity and safety.
Depends: libc6 (>= 2.31)
EOF

# Post-install script
cat > "$BUILD_DIR/DEBIAN/postinst" <<'EOF'
#!/bin/bash
set -e

# Update man page database
if command -v mandb &> /dev/null; then
    mandb
fi

# Verify installation
/usr/bin/athera --version

echo "Athera installed successfully!"
exit 0
EOF

chmod +x "$BUILD_DIR/DEBIAN/postinst"

# Build DEB file
dpkg-deb --build "$BUILD_DIR" \
    "$OUTPUT_DIR/athera_${VERSION}_amd64.deb"

# Sign (optional)
if [ -f ~/.gnupg/secring.gpg ]; then
    dpkg-sig -k "KEY_ID" -s builder \
        "$OUTPUT_DIR/athera_${VERSION}_amd64.deb"
fi

echo "✓ DEB package created: $OUTPUT_DIR/athera_${VERSION}_amd64.deb"
```

### Installation (Ubuntu/Debian)

```bash
# Download
wget https://athera.dev/releases/athera_1.0.0_amd64.deb

# Install
sudo apt install ./athera_1.0.0_amd64.deb

# Verify
athera --version

# To uninstall
sudo apt remove athera
```

### Format 2: RPM (Fedora/RHEL)

**File:** `athera-1.0.0-1.x86_64.rpm` (~25-30 MB)

**Build Process (using FPM - fastest method):**

```bash
#!/bin/bash
set -e

VERSION="1.0.0"
BUILD_DIR="build/rpm"
OUTPUT_DIR="dist"

mkdir -p "$BUILD_DIR/usr/bin"
mkdir -p "$BUILD_DIR/usr/share/man/man1"

cp build/athera "$BUILD_DIR/usr/bin/"
cp installer/linux/athera.1.gz "$BUILD_DIR/usr/share/man/man1/"

# Use FPM to create RPM
fpm -s dir \
    -t rpm \
    -n athera \
    -v "$VERSION" \
    -p "$OUTPUT_DIR/athera-VERSION-RELEASE.ARCH.rpm" \
    --after-install installer/linux/postinstall.sh \
    --license MIT \
    --url https://athera.dev \
    --maintainer "Athera Project <team@athera.dev>" \
    --description "Athera Programming Language" \
    -C "$BUILD_DIR" \
    .

echo "✓ RPM package created"
```

### Installation (Fedora/RHEL)

```bash
# Download
wget https://athera.dev/releases/athera-1.0.0-1.x86_64.rpm

# Install
sudo dnf install ./athera-1.0.0-1.x86_64.rpm

# Verify
athera --version

# To uninstall
sudo dnf remove athera
```

### Format 3: Generic Tarball

**File:** `athera-1.0.0-linux-amd64.tar.gz` (~25-30 MB)

**Build Process:**

```bash
#!/bin/bash
set -e

VERSION="1.0.0"
BUILD_DIR="build/tarball"
OUTPUT_DIR="dist"

# Create structure
mkdir -p "$BUILD_DIR/athera-$VERSION/bin"
mkdir -p "$BUILD_DIR/athera-$VERSION/share/man/man1"
mkdir -p "$BUILD_DIR/athera-$VERSION/lib"

# Copy files
cp build/athera "$BUILD_DIR/athera-$VERSION/bin/"
cp installer/linux/athera.1.gz "$BUILD_DIR/athera-$VERSION/share/man/man1/"
cp LICENSE "$BUILD_DIR/athera-$VERSION/"
cp README.md "$BUILD_DIR/athera-$VERSION/"

# Create install script
cat > "$BUILD_DIR/athera-$VERSION/INSTALL.sh" <<'EOF'
#!/bin/bash
echo "Installing Athera..."
sudo cp bin/athera /usr/local/bin/
sudo chmod +x /usr/local/bin/athera
echo "✓ Installed to /usr/local/bin/athera"
EOF

chmod +x "$BUILD_DIR/athera-$VERSION/INSTALL.sh"

# Create tarball
cd "$BUILD_DIR"
tar -czf "$OUTPUT_DIR/athera-$VERSION-linux-amd64.tar.gz" "athera-$VERSION/"
cd -

echo "✓ Tarball created: $OUTPUT_DIR/athera-$VERSION-linux-amd64.tar.gz"
```

### Installation (Generic Tarball)

```bash
# Download
wget https://athera.dev/releases/athera-1.0.0-linux-amd64.tar.gz

# Extract
tar -xzf athera-1.0.0-linux-amd64.tar.gz
cd athera-1.0.0

# Install
sudo cp bin/athera /usr/local/bin/
sudo chmod +x /usr/local/bin/athera

# Verify
athera --version
```

### Man Page

**File:** `athera.1`

```
.TH ATHERA 1 "January 27, 2026" "athera 1.0.0" "Athera Programming Language"
.SH NAME
athera \- Execute Athera programming language scripts

.SH SYNOPSIS
.B athera
[\fI OPTIONS \fR] [\fI FILE \fR] [\fI ARGS \fR]

.SH DESCRIPTION
Athera is a professional programming language designed for clarity and safety.
Execute Athera programs (.ath files) with the athera command.

.SH COMMANDS
.TP
.B run \fIFILE\fR
Execute an Athera program (default)

.TP
.B compile \fIFILE\fR -o \fIOUTPUT\fR
Compile to standalone executable

.TP
.B repl
Start interactive REPL

.TP
.B install \fIPACKAGE\fR
Install a module from Athera Hub

.SH OPTIONS
.TP
.B -v, --verbose
Verbose output

.TP
.B --version
Show version information

.TP
.B -h, --help
Show help message

.SH EXAMPLES
.TP
.B athera hello.ath
Execute hello.ath program

.TP
.B athera compile hello.ath -o hello
Compile hello.ath to executable

.TP
.B athera repl
Start interactive shell

.SH SEE ALSO
Documentation at https://athera.dev

.SH AUTHOR
Athera Project <team@athera.dev>
```

---

## Verification & Testing

### Quality Assurance

**Post-Install Tests:**

```bash
# Test 1: Binary exists and is executable
test -x /usr/local/bin/athera

# Test 2: Version works
/usr/local/bin/athera --version

# Test 3: Help works
/usr/local/bin/athera --help

# Test 4: Run simple program
echo 'task main: greet "Test"' > test.ath
/usr/local/bin/athera test.ath | grep "Test"

# Test 5: Compile works
/usr/local/bin/athera compile test.ath -o test_bin
./test_bin | grep "Test"

# All tests
installer/tests/verify.sh
```

### Checksum Verification

**Generate checksums:**

```bash
sha256sum athera-1.0.0.pkg > athera-1.0.0.pkg.sha256
sha256sum athera-1.0.0-setup.msi > athera-1.0.0-setup.msi.sha256
sha256sum athera_1.0.0_amd64.deb > athera_1.0.0_amd64.deb.sha256
sha256sum athera-1.0.0-linux-amd64.tar.gz > athera-1.0.0-linux-amd64.tar.gz.sha256
```

**Users verify downloads:**

```bash
sha256sum -c athera-1.0.0.pkg.sha256
```

---

## Distribution Checklist

- [ ] macOS installer built and notarized
- [ ] Windows MSI signed
- [ ] Linux DEB, RPM, and tarball built
- [ ] All checksums generated
- [ ] Virus scan completed (VirusTotal)
- [ ] Installation tested on each platform
- [ ] Uninstall tested
- [ ] Documentation updated
- [ ] Release notes written
- [ ] Posted to athera.dev
- [ ] Announced on social media
- [ ] GitHub releases updated

---

This specification ensures professional-grade installers that users expect from established languages like Python or Node.js.
