# goguard

🔐 Keep your Go dependencies secure, up-to-date, and compliant.

Go Guard is a CLI tool + GitHub Action for scanning Go projects to detect: Vulnerabilities, Outdated dependencies, Unmaintained packages, License risks. It helps developers keep Go projects secure, up-to-date, and compliant with minimal effort.

**Go Guard** is a lightweight CLI tool that scans your Go projects to detect:

- 🚨 Vulnerable packages (via Go vulnerability database & CVEs)  
- ⬆️ Outdated dependencies (compares with the latest versions)  
- 🛠 Clear reports for developers & CI/CD pipelines  

> **Status:** Early stage (v0.1.0). Core functionality is ready — expect improvements and new features soon.  

---

## ✨ Features
- Parse `go.mod` and `go.sum` files  
- Detect vulnerable dependencies using [Go vulnerability database](https://pkg.go.dev/vuln)  
- Check for outdated modules (semver comparison)  
- Generate a clean CLI report  

---

## 📦 Installation

```bash
go install github.com/AumSahayata/goguard@latest
```