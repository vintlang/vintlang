# VintLang Installation Guide

## Installing vintLang on Linux

You can install **vintLang** on your Linux computer using the following steps. This guide will walk you through downloading, extracting, and confirming the installation.

### Step 1: Download the vintLang Binary

First, download the binary release of vintLang for Linux using the following `curl` command:

```bash
curl -O -L https://github.com/ekilie/vint-lang/releases/download/0.1.0/vintLang_linux_amd64_v0.1.2.tar.gz
```

### Step 2: Extract the Binary to a Global Location

Once the download is complete, extract the file and place the binary in a directory that is globally available (`/usr/local/bin` is typically used for this purpose):

```bash
sudo tar -C /usr/local/bin -xzvf vintLang_Linux_amd64.tar.gz
```

This will unpack the binary and make the `vintLang` command available to all users on your system.

### Step 3: Confirm the Installation

To verify that **vintLang** has been installed correctly, run the following command to check its version:

```bash
vintLang -v
```

If the installation was successful, this command will output the version of **vintLang** that was installed.

---



### How to Install `vintLang`:
1. Open your terminal.
2. Run the `curl` command to download the `vintLang` binary.
3. Extract the downloaded archive to a globally accessible directory (`/usr/local/bin`).
4. Confirm the installation by checking the version with `vintLang -v`.

This guide should be easy to follow for installing `vintLang` on Linux!

Now you can start using **vintLang** on your Linux system!