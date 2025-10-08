# VintLang Release Process

This document explains how to create and publish releases for VintLang using GoReleaser and GitHub Actions.

## Overview

VintLang uses [GoReleaser](https://goreleaser.com) to automate the build and release process. When a new version tag is pushed to the repository, GitHub Actions automatically:

1. Builds binaries for multiple platforms (Linux, macOS, Windows)
2. Creates archives (tar.gz for Linux/macOS, zip for Windows)
3. Generates Linux packages (deb, rpm, apk)
4. Calculates checksums
5. Creates a GitHub release with all artifacts
6. Generates release notes from the changelog

## Prerequisites

- Maintainer access to the repository
- Git configured with proper credentials
- Go 1.21 or later installed (for local testing)

## Release Process

### 1. Prepare for Release

Before creating a release, ensure:

- All changes are merged to the `main` branch
- All tests are passing
- Documentation is up to date
- CHANGELOG.md is updated with the new version

### 2. Create a Release (Using the Release Script)

The easiest way to create a release is using the provided release script:

```bash
./scripts/release.sh v0.3.0
```

This script will:
- Validate the version format
- Check you're on the main branch
- Check for uncommitted changes
- Create and push the version tag
- Trigger the GitHub Actions workflow

#### Version Format

Versions must follow the format: `vX.Y.Z` (e.g., `v0.3.0`, `v1.0.0`)

Optional suffixes are supported: `v0.3.0-beta.1`, `v0.3.0-rc.1`

### 3. Manual Release Process

If you prefer to create a release manually:

```bash
# Ensure you're on main branch
git checkout main

# Pull latest changes
git pull origin main

# Create and push the tag
git tag -a v0.3.0 -m "Release v0.3.0"
git push origin main
git push origin v0.3.0
```

### 4. Monitor the Release

After pushing the tag:

1. Go to the [Actions tab](https://github.com/vintlang/vintlang/actions)
2. Find the "Release" workflow for your tag
3. Monitor the build progress
4. Once complete, verify the [release page](https://github.com/vintlang/vintlang/releases)

## Testing Releases Locally

Before creating an official release, you can test the build process locally using GoReleaser:

### Using the Test Script

```bash
./scripts/test-goreleaser.sh
```

This script will:
- Install GoReleaser if not already installed
- Build binaries for all platforms in snapshot mode
- Create archives and packages
- Generate checksums

All artifacts will be in the `dist/` directory.

### Manual Testing

```bash
# Install goreleaser (if not installed)
go install github.com/goreleaser/goreleaser@latest

# Ensure it's in your PATH
export PATH="$HOME/go/bin:$PATH"

# Test the configuration
goreleaser check

# Build for a single platform (fast)
goreleaser build --snapshot --clean --single-target

# Build for all platforms (slower)
goreleaser release --snapshot --clean --skip=publish
```

## Configuration

### GoReleaser Configuration

The GoReleaser configuration is in `.goreleaser.yml`. Key settings:

- **Builds**: Configured for Linux, macOS, and Windows (amd64, arm64, 386)
- **Archives**: tar.gz for Unix-like systems, zip for Windows
- **Packages**: Generates deb, rpm, and apk packages for Linux
- **Checksums**: SHA256 checksums for all artifacts

### GitHub Actions Workflow

The release workflow is in `.github/workflows/build.yml`. It:

- Triggers on tags matching `v*`
- Runs tests before building
- Uses GoReleaser to build and publish
- Requires `GITHUB_TOKEN` (automatically provided)

## Troubleshooting

### Build Fails

If the build fails:

1. Check the Actions logs for detailed error messages
2. Test locally with `./scripts/test-goreleaser.sh`
3. Ensure `.goreleaser.yml` is valid with `goreleaser check`

### Missing Artifacts

If some artifacts are missing from the release:

1. Check the GoReleaser configuration for the platform
2. Verify the build succeeded for that platform in the Actions logs
3. Test locally with `goreleaser release --snapshot --clean --skip=publish`

### Wrong Version Number

If the version number is incorrect:

1. Check that the tag name is correct (should start with `v`)
2. Verify the ldflags in `.goreleaser.yml` are set correctly
3. The version is injected at build time from the Git tag

## Platform Support

VintLang is built for the following platforms:

| OS      | Architectures        | Formats          |
|---------|---------------------|------------------|
| Linux   | amd64, arm64, 386   | tar.gz, deb, rpm, apk |
| macOS   | amd64, arm64        | tar.gz           |
| Windows | amd64, 386          | zip              |

Note: Windows arm64 and macOS 386 are excluded due to lack of support.

## Release Artifacts

Each release includes:

1. **Binaries**: Pre-built executables for each platform
2. **Archives**: Compressed archives containing the binary and documentation
3. **Linux Packages**: Native packages for Debian/Ubuntu (deb), Red Hat/Fedora (rpm), and Alpine (apk)
4. **Checksums**: SHA256 checksums for verifying downloads
5. **Release Notes**: Auto-generated from commit messages

## Best Practices

1. **Always test locally first**: Run `./scripts/test-goreleaser.sh` before pushing a tag
2. **Follow semantic versioning**: Use major.minor.patch (e.g., v1.2.3)
3. **Update documentation**: Ensure README.md and CHANGELOG.md are current
4. **Test the binaries**: Download and test artifacts from the release page
5. **Announce the release**: Update the community about new releases

## Additional Resources

- [GoReleaser Documentation](https://goreleaser.com)
- [Semantic Versioning](https://semver.org)
- [VintLang Releases](https://github.com/vintlang/vintlang/releases)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
