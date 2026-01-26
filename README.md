# Product Requirements Agent Team

[![Build Status][build-status-svg]][build-status-url]
[![Lint Status][lint-status-svg]][lint-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

CLI and AI assistant integrations for creating, validating, scoring, and managing Product Requirements Documents.

## Features

- **Structured PRDs**: JSON-based schema for consistent, machine-readable documents
- **Quality Scoring**: Automated rubric-based scoring across 10 categories
- **Multiple Views**: PM-focused and Executive summary views
- **AI Integration**: Claude Code (MCP) and Kiro IDE (Power) support
- **Validation**: Schema validation, ID format checking, traceability verification

## Quick Start

```bash
# Install
go install github.com/agentplexus/agent-team-prd/cmd/prdtool@latest

# Create a PRD
prdtool init --title "User Authentication" --owner "Jane PM"
prdtool add problem --statement "Users cannot securely access accounts"
prdtool add goal --statement "Reduce password tickets by 50%"
prdtool add req --description "Support OAuth 2.0 login" --priority must

# Validate and score
prdtool validate
prdtool score
```

## AI Assistant Integration

```bash
# Claude Code
prdtool deploy --target claude

# Kiro IDE
prdtool deploy --target kiro-power --output ~/.kiro/powers/prdtool
```

## Documentation

Full documentation: https://agentplexus.github.io/agent-team-prd/

- [Installation](https://agentplexus.github.io/agent-team-prd/installation/)
- [CLI Reference](https://agentplexus.github.io/agent-team-prd/cli/commands/)
- [Claude Code Integration](https://agentplexus.github.io/agent-team-prd/integrations/claude-code/)
- [Kiro IDE Integration](https://agentplexus.github.io/agent-team-prd/integrations/kiro-ide/)
- [PRD Schema Reference](https://agentplexus.github.io/agent-team-prd/reference/prd-schema/)

## Build from Source

```bash
git clone https://github.com/agentplexus/agent-team-prd.git
cd agent-team-prd
go build -o bin/prdtool ./cmd/prdtool
go build -o bin/prdtool-mcp ./cmd/prdtool-mcp
```

## License

MIT

 [build-status-svg]: https://github.com/agentplexus/agent-team-prd/actions/workflows/ci.yaml/badge.svg?branch=main
 [build-status-url]: https://github.com/agentplexus/agent-team-prd/actions/workflows/ci.yaml
 [lint-status-svg]: https://github.com/agentplexus/agent-team-prd/actions/workflows/lint.yaml/badge.svg?branch=main
 [lint-status-url]: https://github.com/agentplexus/agent-team-prd/actions/workflows/lint.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/agentplexus/agent-team-prd
 [goreport-url]: https://goreportcard.com/report/github.com/agentplexus/agent-team-prd
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/agentplexus/agent-team-prd
 [docs-godoc-url]: https://pkg.go.dev/github.com/agentplexus/agent-team-prd
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/agentplexus/agent-team-prd/blob/master/LICENSE
 [used-by-svg]: https://sourcegraph.com/github.com/agentplexus/agent-team-prd/-/badge.svg
 [used-by-url]: https://sourcegraph.com/github.com/agentplexus/agent-team-prd?badge