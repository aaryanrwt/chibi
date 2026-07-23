<div align="center">
  <h1>Chibi</h1>
  <p><b>The AI-Native, Predictive SRE Assistant for Kubernetes.</b></p>

  [![Go Version](https://img.shields.io/github/go-mod/go-version/chibi/chibi)](https://golang.org/doc/devel/release.html)
  [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
  [![CI](https://github.com/chibi/chibi/actions/workflows/ci.yml/badge.svg)](https://github.com/chibi/chibi/actions/workflows/ci.yml)
  [![Release](https://img.shields.io/github/v/release/chibi/chibi?label=Stable)](https://github.com/chibi/chibi/releases)
  [![Stars](https://img.shields.io/github/stars/chibi/chibi?style=social)](https://github.com/chibi/chibi/stargazers)
  [![Open Source](https://img.shields.io/badge/Open%20Source-Yes-red)](https://github.com/chibi/chibi)
  [![2026 Edition](https://img.shields.io/badge/Edition-2026-black)](#)
</div>

<br/>

> **Chibi** bridges the gap between powerful but complex tools like `kubectl` and heavy visual dashboards like Lens. It lives in your terminal, understands your cluster via natural language, and predicts failures before they happen.

---

## Why Chibi?

Kubernetes troubleshooting is traditionally painful. Engineers bounce between `kubectl get events`, `kubectl logs`, and Prometheus dashboards, manually correlating data to find the root cause of a crashing pod. 

Chibi solves this by acting as a **Principal SRE in your terminal**. It automatically fetches cluster state, correlates it with live Prometheus telemetry, and feeds it into an LLM (OpenAI, Anthropic, Ollama) to give you an instant, actionable diagnosis. 

---

## Key Features

- **AI-Powered Diagnostics:** Ask "Why is the payments pod crashing?" and get a deterministic, context-aware answer.
- **Predictive SRE Daemon:** Chibi runs in the background, analyzing event streams via a local SQLite state database to predict failures before they trigger PagerDuty.
- **Zero-Trust WASM Plugins:** Extend Chibi securely. Third-party plugins execute in a strict WebAssembly (`wazero`) sandbox with zero host filesystem/network access.
- **Live Terminal Dashboards:** Beautiful, interactive TUI built with Bubble Tea, featuring live CPU/Memory sparklines natively in your terminal.
- **Autonomous Remediation:** Configure strict RBAC policies allowing Chibi to automatically heal known issues (e.g., auto-restarting degraded deployments).

---

## Three Unique Innovations

1. **Telemetry Correlator:** Competitors just feed raw YAML to LLMs. Chibi actively queries Prometheus, merging live time-series data with static YAML for profound context.
2. **WASM Sandboxing:** CLI plugins are a massive security risk. Chibi's `wazero` integration means you can run community plugins with mathematical certainty they aren't stealing your `KUBECONFIG`.
3. **True Terminal Native:** We didn't build a web app. Chibi boots in <200ms and renders 60fps interactive dashboards directly in your favorite terminal emulator.

---

## Installation

### Go Install
```bash
go install github.com/chibi/chibi/cmd/chibi@latest
```

### Pre-compiled Binaries (Windows/Linux/macOS)
Download the latest release from the [Releases page](https://github.com/chibi/chibi/releases).

---

## Quick Start

1. **Configure your AI Provider:**
   ```bash
   export OPENAI_API_KEY="sk-..."
   ```
2. **Launch Chibi:**
   ```bash
   chibi
   ```
3. **Diagnose a Pod:**
   Inside the TUI, type:
   ```text
   diagnose pod frontend-deployment-6f9b8c
   ```
   *Chibi will fetch the pod manifest, correlate CPU metrics, and stream the AI's diagnosis.*

---

## CLI Overview

- `chibi` - Launches the interactive Bubble Tea TUI.
- `chibi diagnose <resource> <name>` - One-shot AI diagnosis.
- `chibi watch` - Starts the predictive daemon in the foreground.
- `chibi plugins list` - Lists installed WASM plugins.
- `chibi config` - Manages provider credentials and sandbox policies.

---

## Architecture

Chibi is built in modular Go:
- **CLI Layer:** Cobra & Viper.
- **TUI Engine:** Bubble Tea, Lip Gloss, Bubbles (Sparklines).
- **K8s Layer:** `client-go` with `ristretto` caching.
- **AI Layer:** Provider router (OpenAI default) + Telemetry Correlator.
- **Plugin Runtime:** `wazero` WASM sandbox.
- **State Engine:** Pure-Go SQLite (`modernc.org/sqlite`) for event anomalies.

---

## Configuration

Configuration is stored in `~/.chibi.yaml`.
```yaml
ai:
  provider: openai
  model: gpt-4o
observability:
  prometheus_url: http://prometheus-operated.monitoring:9090
remediation:
  auto_apply_deployments: true
```

---

## AI Providers

Chibi natively supports:
- **OpenAI** (`GPT-4o`) - Recommended for deep logic.
- **Anthropic** (`Claude 3.5 Sonnet`) - Excellent for YAML analysis.
- **Ollama** - For highly secure, local, air-gapped clusters.

---

## Security & Privacy

Security is paramount. 
- **RBAC:** Chibi uses your local `kubeconfig` and inherits your exact RBAC permissions. It cannot bypass cluster security.
- **Zero-Trust Plugins:** External plugins cannot access the network.
- **Read-Only Default:** Chibi acts in read-only mode by default. Autonomous remediation must be explicitly enabled per resource type.

---

## Comparison

| Feature | Chibi | kubectl | Lens |
|---------|-------|---------|------|
| **Speed** | <200ms | <100ms | >2000ms |
| **Interface** | TUI | CLI | GUI |
| **AI Native** | Yes | No | Paid |
| **Predictive**| Yes | No | No |
| **Plugins** | WASM (Secure) | Binary (Unsafe) | Node (Heavy) |

---

## Contributing

We welcome contributions! 
Please read our [CONTRIBUTING.md](CONTRIBUTING.md) and [SECURITY.md](SECURITY.md) before submitting PRs. 

---

## Roadmap

- [x] **Version 1:** Reactive Diagnostics & Terminal UI
- [x] **Version 2:** Predictive Daemon, Prometheus Integration, WASM Sandbox

*(Note: V1 and V2 are fully implemented and stable).*

---

## License

Apache 2.0. See [LICENSE](LICENSE).

---

<div align="center">
  <p><b>Built for the Modern Kubernetes Era (2026)</b></p>
  <p>Research & Product Vision: Aaryan Rawat</p>
</div>
