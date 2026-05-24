#!/usr/bin/env bash
set -g

REPO="InuzDev/InuSDK"
SHIMS_DIR="$HOME/.inusdk/shims"
BIN_PATH="$SHIMS_DIR/inusdk"

echo ""
echo "Installing InuSDK. . ."

# Detect the OS and the architecture of the machine
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "$ARCH" in
   x86_64) ARCH="amd64" ;;
   aarch64) ARCH="arm64" ;;
   arm64) ARCH="arm64" ;;
   *)
      echo "Unsupported architecture: $ARCH"
      exit 1
      ;;
esac

case "$OS" in
   linux) OS="linux" ;;
   darwin) OS="darwin" ;;
   *)
      echo "Unsupported OS: $OS"
      exit 1
      ;;
esac

# Fetch the latest release
echo "Fetching the latest releas. . ."
RELEASE=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest")
VERSION=$(echo "$RELEASE" | grep '"tag_name"' | cut -d'"' -f4)
DOWNLOAD_URL=$(echo "$RELEASE" | grep "browser_download_url" | grep "$OS" | grep "$ARCH" | cut -d'"' -f4)

if [ -z "$DOWNLOAD_URL" ]; then
   echo "No binary found for $OS/$ARCH"
   exit 1
fi

# Create the shims directory
mkdir -p "$SHIMS_DIR"

# Download and extract the files
echo "Downlaoding $VERSION for $OS/$ARCH"
TEMP_DIR=$(mktemp -d)
TEMP_FILE="$TEMP_DIR/inusdk.tar.gz"

curl -fsSL "$DOWNLOAD_URL" -o "$TEMP_FILE"

echo "Extracting. . ."
tar -xzf "$TEMP_FILE" -C "$TEMP_DIR"
cp "$TEMP_DIR/inusdk" "$BIN_PATH"

chmod +x "$BIN_PATH"

# Cleanup
rm -rf "$TEMP_DIR"

# Add shims to path
add_to_path() {
   local config_files="$1"
   local export_line="export PATH=\"$SHIMS_DIR:\$PATH\""

   if [ -f "$config_files" ] && grep -q "$SHIMS_DIR" "$config_files"; then
      echo "Already in PATH via $config_files"
      return
   fi

   echo "" >> "$config_files"
   echo "# InuSDK" >> "$config_files"
   echo "$export_line" >> "$config_files"
   echo "Added to PATH via $config_files"
}

SHELL_NAME="$(basename "$SHELL")"
case "$SHELL_NAME" in
   zsh) add_to_path "$HOME/.zshrc" ;;
   bash) add_to_path "$HOME/.bashrc" ;;
   fish) add_to_path "$HOME/.config/fish/config.fish" ;;
   *)
      add_to_path "$HOME/.bashrc" ;;
esac

echo ""
echo "InuSDK $VERSION installed"
echo "Restart your terminal and run: inusdk"
echo ""
