# GoReleaser Quick Start Guide

This is a quick reference guide for using GoReleaser with VintLang. For comprehensive documentation, see [RELEASE_PROCESS.md](RELEASE_PROCESS.md).

## What is GoReleaser?

GoReleaser is a tool that automates the building and releasing of Go projects. It handles:
- Cross-platform compilation
- Archive creation
- Package generation (deb, rpm, apk)
- Checksum calculation
- GitHub release creation
- Release notes generation

## Quick Commands

### Test Locally

```bash
# Test the GoReleaser configuration
./scripts/test-goreleaser.sh

# Or manually:
goreleaser check                                    # Validate config
goreleaser build --snapshot --clean --single-target # Quick build
goreleaser release --snapshot --clean --skip=publish # Full build
```

### Create a Release

```bash
# Using the release script (recommended)
./scripts/release.sh v0.3.0

# Or manually
git tag -a v0.3.0 -m "Release v0.3.0"
git push origin main
git push origin v0.3.0
```

## Understanding the Build Output

When you run GoReleaser, it creates a `dist/` directory with:

```
dist/
├── vintLang_Linux_x86_64.tar.gz          # Linux amd64 archive
├── vintLang_Linux_arm64.tar.gz           # Linux arm64 archive
├── vintLang_Linux_i386.tar.gz            # Linux 32-bit archive
├── vintLang_Darwin_x86_64.tar.gz         # macOS Intel archive
├── vintLang_Darwin_arm64.tar.gz          # macOS Apple Silicon archive
├── vintLang_Windows_x86_64.zip           # Windows 64-bit archive
├── vintLang_Windows_i386.zip             # Windows 32-bit archive
├── vintlang_X.Y.Z_linux_amd64.deb        # Debian package
├── vintlang_X.Y.Z_linux_amd64.rpm        # Red Hat package
├── vintlang_X.Y.Z_linux_amd64.apk        # Alpine package
├── checksums.txt                          # SHA256 checksums
└── vintlang_<platform>_<arch>/           # Raw binaries
```

## How GitHub Actions Uses GoReleaser

When you push a tag (e.g., `v0.3.0`):

1. **Trigger**: GitHub Actions detects the tag push
2. **Test**: Runs the test job (validates the build environment)
3. **Build**: GoReleaser builds for all platforms
4. **Package**: Creates archives and Linux packages
5. **Release**: Uploads all artifacts to GitHub Releases
6. **Changelog**: Generates release notes from commits

## Configuration Files

### `.goreleaser.yml`

This is the main configuration file. Key sections:

- `before.hooks`: Commands to run before building (e.g., `go mod tidy`)
- `builds`: Platform/architecture matrix and build settings
- `archives`: Archive format and contents
- `nfpms`: Linux package configuration (deb, rpm, apk)
- `checksum`: Checksum algorithm
- `changelog`: Changelog generation rules
- `release`: GitHub release settings

### `.github/workflows/build.yml`

The GitHub Actions workflow that:
- Triggers on tag pushes matching `v*`
- Sets up the build environment
- Runs GoReleaser
- Publishes the release

## Common Issues and Solutions

### Issue: "only configurations files on version: 1 are supported"

**Solution**: Ensure `.goreleaser.yml` has `version: 1` (not `version: 2`)

### Issue: "goreleaser: command not found"

**Solution**: Install goreleaser:
```bash
go install github.com/goreleaser/goreleaser@latest
export PATH="$HOME/go/bin:$PATH"
```

### Issue: Build fails with "no git tag found"

**Solution**: You're trying to do a real release without a tag. Use snapshot mode:
```bash
goreleaser release --snapshot --clean --skip=publish
```

### Issue: Binary is too large

**Solution**: The binaries are stripped (`-s -w` ldflags). For even smaller binaries, consider using UPX compression (see `upx_build` directory).

## Customization

To modify what gets built or how:

1. Edit `.goreleaser.yml`
2. Test with `goreleaser check`
3. Build locally with `./scripts/test-goreleaser.sh`
4. Commit and push changes

Common customizations:
- Add/remove platforms in `builds.goos` and `builds.goarch`
- Change archive contents in `archives.files`
- Modify Linux package metadata in `nfpms`
- Adjust changelog rules in `changelog.filters`

## Resources

- [Full Release Documentation](RELEASE_PROCESS.md)
- [GoReleaser Official Docs](https://goreleaser.com)
- [VintLang Releases](https://github.com/vintlang/vintlang/releases)

## Support

For questions or issues:
1. Check the [Release Process Documentation](RELEASE_PROCESS.md)
2. Review GoReleaser logs in GitHub Actions
3. Test locally with `./scripts/test-goreleaser.sh`
4. Open an issue on GitHub
