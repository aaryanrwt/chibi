# Chibi Security Checklist (Pre-Release)

This checklist MUST be completed prior to every public release to ensure zero exposure of sensitive data and maintain the integrity of the Chibi repository.

## Repository Security
- [ ] No internal machine names or IP addresses are hardcoded in the codebase.
- [ ] No local development paths (e.g., `C:\Users\...` or `/Users/...`) exist in source files.
- [ ] Default configurations point to `localhost` or safe placeholder domains (`example.com`).

## Git Security
- [ ] The `.gitignore` is up to date and aggressively excludes secrets, binaries, and local configs.
- [ ] Git history has been checked using `git log -p` or tools like `trufflehog`/`gitleaks` to ensure no keys were ever committed.

## Documentation Security
- [ ] README.md contains no real API keys.
- [ ] Screenshots contain no PII, internal cluster names, or sensitive infrastructure details.
- [ ] Markdown files in `docs/` and `research/` are sanitized of internal network graphs.

## Dependency Security
- [ ] `go list -m all` reviewed for deprecated or unmaintained libraries.
- [ ] `govulncheck ./...` executed with zero high-severity findings.
- [ ] Dependabot (or Renovate) is configured in the repository.

## Secret Management
- [ ] `OPENAI_API_KEY`, `ANTHROPIC_API_KEY`, etc. are only loaded via OS environment variables.
- [ ] No fallback API keys exist in the Go source code.
- [ ] The repository contains no `.pem`, `.crt`, `.key`, or `kubeconfig` files.

## CI/CD Security
- [ ] GitHub Actions workflows use least-privilege `permissions`.
- [ ] No secrets are echoed into build logs.
- [ ] Third-party GitHub Actions are pinned to specific SHA hashes where applicable.

## Plugin Security (WASM)
- [ ] WASM plugin sandbox configuration explicitly denies host filesystem mounting.
- [ ] WASM plugin sandbox configuration explicitly denies host network access.

## Kubernetes Security
- [ ] RBAC tokens are never cached to disk in plaintext.
- [ ] Default `client-go` connection is read-only unless auto-remediation is explicitly enabled by the user.
