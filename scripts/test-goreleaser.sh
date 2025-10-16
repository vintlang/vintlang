#!/bin/bash

# Test GoReleaser configuration locally
# This script will build the project using GoReleaser in snapshot mode

set -e

echo "🔧 Testing GoReleaser configuration..."

# Ensure ~/go/bin is in PATH
export PATH="$HOME/go/bin:$PATH"

# Function to install GoReleaser via Homebrew if Go install fails
install_goreleaser_fallback() {
    echo "🍺 Trying to install GoReleaser via Homebrew..."
    if command -v brew &> /dev/null; then
        # Use cask version as the formula is deprecated
        brew install --cask goreleaser
        return 0
    else
        echo "❌ Homebrew not found. Trying direct binary download..."
        # Try direct binary download for macOS ARM64
        GORELEASER_VERSION="v2.12.5"
        BINARY_URL="https://github.com/goreleaser/goreleaser/releases/download/${GORELEASER_VERSION}/goreleaser_Darwin_arm64.tar.gz"
        
        echo "📥 Downloading GoReleaser binary..."
        curl -L -o /tmp/goreleaser.tar.gz "${BINARY_URL}"
        
        echo "📦 Extracting to ~/go/bin..."
        mkdir -p ~/go/bin
        tar -xzf /tmp/goreleaser.tar.gz -C /tmp
        mv /tmp/goreleaser ~/go/bin/
        chmod +x ~/go/bin/goreleaser
        
        # Cleanup
        rm /tmp/goreleaser.tar.gz
        
        echo "✅ GoReleaser installed via direct download"
        return 0
    fi
}

# Check if goreleaser is installed
if ! command -v goreleaser &> /dev/null; then
    echo "❌ GoReleaser is not installed. Installing..."
    
    # Try Go install first, with fallback if it fails
    if ! go install github.com/goreleaser/goreleaser@v1.25.1; then
        echo "⚠️  Go install failed, trying alternative installation methods..."
        install_goreleaser_fallback
    fi
fi

# Verify installation
if ! command -v goreleaser &> /dev/null; then
    echo "❌ Failed to install GoReleaser. Please install manually:"
    echo "   brew install goreleaser/tap/goreleaser"
    echo "   or download from: https://github.com/goreleaser/goreleaser/releases"
    exit 1
fi

echo "✅ GoReleaser found: $(goreleaser --version)"

# Clean previous builds
echo "🧹 Cleaning previous builds..."
rm -rf dist/

# Run GoReleaser in snapshot mode (no git tag required)
echo "🚀 Running GoReleaser in snapshot mode..."
goreleaser release --snapshot --clean

echo "✅ GoReleaser test completed successfully!"
echo "📦 Built artifacts can be found in the 'dist/' directory"

# List the generated artifacts
echo ""
echo "Generated artifacts:"
ls -la dist/
