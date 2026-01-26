# Product Requirements Agent Team

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
