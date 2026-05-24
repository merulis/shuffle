# shuffle

`shuffle` is a CLI tool for browsing, viewing, and downloading artifacts from configured remote sources.

The project is focused on creating a unified workflow for working with different artifact providers such as GitHub, GitLab, Nexus, and other remote systems.

---

# Features

Current functionality:

* source configuration;
* GitHub provider support;
* artifact listing;
* artifact viewing;
* artifact downloading;
* credential references via environment variables;
* Cobra-based CLI.

---

# Quick Start

## Export GitHub token

```bash
export GITHUB_TOKEN=your_token_here
```

## Add source

```bash
shuffle source add gh-cli \
  --type github \
  --owner cli \
  --repo cli \
  --ref trunk \
  --credential-ref env:GITHUB_TOKEN
```

## List configured sources

```bash
shuffle source list
```

---

# Usage

## List artifacts

```bash
shuffle list gh-cli:/cmd
```

## View file

```bash
shuffle view gh-cli:/go.mod
```

## Download file

```bash
shuffle download gh-cli:/go.mod
```

---

# Configuration

Configuration is stored in:

```text
~/.shuffle/config.json
```

Example:

```json
{
  "sources": [
    {
      "name": "gh-cli",
      "type": "github",
      "owner": "cli",
      "repo": "cli",
      "ref": "trunk",
      "credential_ref": "env:GITHUB_TOKEN"
    }
  ]
}
```

---

# Roadmap

Planned improvements:

* encrypted credentials;
* system keyring integration;
* GitLab provider;
* Nexus provider;
* artifact search;
* caching;
* TUI interface.
