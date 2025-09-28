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
- 📊 Output as **table** or **JSON** or **HTML**
- 🚪 Exit codes for CI pipelines (see below)
- ⏲ Each module gets a RiskScore (numeric) and RiskLevel (Low / Medium / High).



## 🚪 Exit Codes

- 0 → All checks passed (OK)
- 1 → Warnings detected (e.g. outdated or risky license or stale repo)
- 2 → Failures detected (vulnerabilities, archived repos)

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

Outputs a table of all modules, their versions, and issues on the console.

### JSON output to console
```bash
goguard scan --json
```

### JSON output to file
```bash
goguard scan --json-file result.json
```

### HTML output to file
```bash
goguard scan --html-file result.html
```

### Verbose exit reasons
```bash
goguard scan --verbose
```

### Strict check
```bash
goguard scan --strict
```
Fails even for warnings

## GitHub Actions Workflow Example

```yml
name: GoGuard Scan

on:
  pull-request:
    branches: [main]

jobs:
  goguard:
    name: GoGuard
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.24'

      - name: Install goguard
        run: go install github.com/AumSahayata/goguard@latest
      
      - name: Run gogurad
        run: goguard scan --json-file report.json --verbose
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
    "Name": "golang.org/x/mod",
    "Version": "v0.28.0",
    "Latest": "v0.28.0",
    "Vulnerable": false,
    "CVEs": [],
    "Status": "[WARN] License",
    "Issues": "License: Unknown",
    "RiskScore": 2,
    "RiskLevel": "Low"
  },
]

```

## 📜 License

MIT License. See [LICENSE](https://github.com/AumSahayata/goguard/blob/main/LICENSE)
 for details.


## 🤝 Contributing

PRs and issues are welcome and feel free to suggest features.
