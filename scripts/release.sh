#!/bin/bash

# Release script for VintLang
# Usage: ./scripts/release.sh <version>
# Example: ./scripts/release.sh v0.2.0

set -e

if [ $# -eq 0 ]; then
    echo "❌ Error: Please provide a version number"
    echo "Usage: $0 <version>"
    echo "Example: $0 v0.2.0"
    exit 1
fi

VERSION=$1

# Validate version format (should start with 'v')
if [[ ! $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+.*$ ]]; then
    echo "❌ Error: Version should be in format v1.2.3 (with optional suffix)"
    echo "Example: v0.2.0 or v0.2.0-beta.1"
    exit 1
fi

echo "🏷️  Creating release for version: $VERSION"

# Check if we're on the main branch
CURRENT_BRANCH=$(git branch --show-current)
if [ "$CURRENT_BRANCH" != "main" ]; then
    echo "⚠️  Warning: You're not on the main branch (current: $CURRENT_BRANCH)"
    read -p "Do you want to continue? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "❌ Aborted"
        exit 1
    fi
fi

# Check for uncommitted changes
if ! git diff-index --quiet HEAD --; then
    echo "❌ Error: You have uncommitted changes. Please commit or stash them first."
    exit 1
fi

# Version is injected at build time using Go's ldflags. No need to update main.go.
# Create and push tag
echo "🏷️  Creating and pushing tag..."
git tag -a "$VERSION" -m "Release $VERSION"
git push origin main
git push origin "$VERSION"

echo "✅ Release $VERSION created successfully!"
echo "🚀 GitHub Actions will now build and publish the release."
echo "📦 Check the progress at: https://github.com/vintlang/vintlang/actions"
