# goguard

🔐 Keep your Go dependencies secure, up-to-date, and compliant.

Go Guard is a CLI tool + GitHub Action for scanning Go projects to detect: Vulnerabilities, Outdated dependencies, Unmaintained packages, License risks. It helps developers keep Go projects secure, up-to-date, and compliant with minimal effort.

> 🚀 Current version: **v0.2.0**


## ✨ Features

- 🔍 Scan `go.mod` / `go.sum` dependencies
- 🛡️ Detect vulnerable dependencies using [Go vulnerability database](https://pkg.go.dev/vuln)  
- 📦 Check for outdated dependencies via the Go proxy
- 🏚️ Detect unmaintained repos (archived or stale >2 years)
- ⚖️ Identify licenses (via GitHub `LICENSE` file)
- 📊 Output as **table** or **JSON**
- 🚪 Exit codes for CI pipelines (see below)



## 🚪 Exit Codes

- 0 → Scan successful, no issues
- 1 → Warnings (outdated or stale repos)
- 2 → Vulnerabilities found

Use these exit codes in CI/CD pipelines to fail builds on security issues.



## 📦 Installation

```bash
go install github.com/AumSahayata/goguard@latest
```


## 🚀 Usage

### Basic scan

```bash
goguard scan
```

Outputs a table of all modules, their versions, and issues.

### JSON output to console
```bash
goguard scan --json
```

### JSON output to file
```bash
goguard scan --json-file result.json
```

### Verbose exit reasons
```bash
goguard scan --verbose
```

## 📊 Example Output

### Table Output

| Package                  | Version | Latest  | Status         | Issues                     |
|---------------------------|---------|---------|----------------|----------------------------|
| github.com/gin-gonic/gin  | v1.7.0  | v1.9.1  | [WARN] Outdated | CVE-2023-1234 (High)       |
| github.com/pkg/errors     | v0.9.1  | v0.9.1  | [OK] Up-to-date | -                          |
| github.com/old/lib        | v1.0.0  | Unknown       | [FAIL] Unmaint. | Repo archived              |

### JSON Output

```json
[
  {
    "Name": "github.com/spf13/pflag",
    "Version": "v1.0.9",
    "Latest": "v1.0.10",
    "Vulnerable": false,
    "CVEs": [],
    "Status": "[WARN] Outdated",
    "Issues": "-"
  }
]

```

## 📜 License

MIT License. See [LICENSE](https://github.com/AumSahayata/goguard/blob/main/LICENSE)
 for details.


## 🤝 Contributing

PRs and issues are welcome and feel free to suggest features.
