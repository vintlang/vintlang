#!/bin/bash

# Test GoReleaser configuration locally
# This script will build the project using GoReleaser in snapshot mode

set -e

echo "🔧 Testing GoReleaser configuration..."

# Check if goreleaser is installed
if ! command -v goreleaser &> /dev/null; then
    echo "❌ GoReleaser is not installed. Installing..."
    go install github.com/goreleaser/goreleaser@latest
fi

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
