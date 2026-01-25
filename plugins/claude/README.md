# PRD Tool - Claude Code Plugin

Claude Code plugin for AI-assisted PRD creation and management.

## Installation

Copy this plugin to your Claude Code plugins directory:

```bash
cp -r plugins/claude ~/.claude/plugins/prdtool
```

Or symlink for development:

```bash
ln -s $(pwd)/plugins/claude ~/.claude/plugins/prdtool
```

## Prerequisites

Ensure the MCP server is installed:

```bash
# Install from source
go install github.com/agentplexus/agent-team-prd/cmd/prdtool-mcp@latest

# Or build locally
go build -o /usr/local/bin/prdtool-mcp ./cmd/prdtool-mcp
```

## Available Commands

| Command | Description |
|---------|-------------|
| `/prdtool:create` | Create a new PRD interactively |
| `/prdtool:review` | Review and score an existing PRD |
| `/prdtool:exec-summary` | Generate executive summary |

## Skills

The plugin includes skills that activate based on conversation context:

| Skill | Triggers |
|-------|----------|
| PRD Creation | "create prd", "new prd", "product requirements" |
| PRD Review | "review prd", "score prd", "improve prd" |
| Executive Summary | "exec summary", "leadership review" |

## Usage Examples

### Create a PRD

```
/prdtool:create User Authentication Feature
```

Guides you through:

1. Problem discovery
2. User persona definition
3. Goals and non-goals
4. Solution exploration
5. Requirements documentation
6. Success metrics
7. Risk identification

### Review a PRD

```
/prdtool:review
```

Provides:

- Quality score (0-10)
- Category breakdown
- Improvement recommendations
- Prioritized action items

### Generate Executive Summary

```
/prdtool:exec-summary
```

Creates a decision-ready summary for leadership review.

## Quality Scoring

PRDs are scored across 10 categories:

| Category | Weight |
|----------|--------|
| Problem Definition | 20% |
| Solution Fit | 15% |
| User Understanding | 10% |
| Market Awareness | 10% |
| Scope Discipline | 10% |
| Requirements Quality | 10% |
| Metrics Quality | 10% |
| UX Coverage | 5% |
| Technical Feasibility | 5% |
| Risk Management | 5% |

**Decision Thresholds:**

- >= 8.0: APPROVE
- >= 6.5: REVISE
- < 6.5: HUMAN REVIEW
- <= 3.0: REJECT

## Note

This plugin is auto-generated from the canonical spec. Edit `plugins/spec/` and regenerate with:

```bash
go run ./plugins/generate
```
