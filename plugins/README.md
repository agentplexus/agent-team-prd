# PRD Tool Plugins

This directory contains plugins for various AI coding assistants.

## Structure

```
plugins/
├── spec/              # Canonical JSON specifications (source of truth)
│   ├── plugin.json    # Plugin metadata
│   ├── commands/      # Command definitions
│   └── skills/        # Skill definitions
├── claude/            # Claude Code plugin (generated)
│   ├── .claude-plugin/
│   ├── commands/
│   └── skills/
└── kiro/              # Kiro IDE Power (generated)
    ├── POWER.md
    ├── mcp.json
    └── steering/
```

## Generating Plugins

Plugins are auto-generated from the canonical spec using the `assistantkit` CLI:

```bash
# Install assistantkit (if not already installed)
go install github.com/agentplexus/assistantkit/cmd/assistantkit@latest

# Generate plugins from spec
assistantkit generate plugins
```

Or run directly without installing:

```bash
go run github.com/agentplexus/assistantkit/cmd/assistantkit@latest generate plugins
```

This reads from `plugins/spec/` and generates:

- `plugins/claude/` - Claude Code plugin
- `plugins/kiro/` - Kiro IDE Power

## Canonical Spec Format

### plugin.json

```json
{
  "name": "prdtool",
  "displayName": "PRD Tool",
  "version": "1.0.0",
  "description": "...",
  "keywords": ["prd", "product requirements", ...],
  "mcpServers": {
    "prdtool": {
      "command": "prdtool-mcp",
      "args": []
    }
  }
}
```

### commands/*.json

```json
{
  "name": "create",
  "description": "Create a new PRD interactively",
  "arguments": [...],
  "instructions": "...",
  "process": [...],
  "examples": [...]
}
```

### skills/*.json

```json
{
  "name": "prd-creation",
  "description": "Guides through PRD creation...",
  "instructions": "...",
  "triggers": ["create prd", "new prd", ...]
}
```

## Installation

### Claude Code

Copy the plugin to your Claude Code plugins directory:

```bash
cp -r plugins/claude ~/.claude/plugins/prdtool
```

Or symlink for development:

```bash
ln -s $(pwd)/plugins/claude ~/.claude/plugins/prdtool
```

### Kiro IDE

Import the Power from the local folder:

1. Open **Powers** panel in Kiro IDE
2. Click **Import**
3. Select `plugins/kiro/`

Or copy to Kiro powers directory:

```bash
cp -r plugins/kiro ~/.kiro/powers/prdtool
```

## Development Workflow

1. Edit canonical specs in `plugins/spec/`
2. Run generator: `go run ./plugins/generate`
3. Test the generated plugins
4. Commit both spec and generated files

The generator ensures consistency across all platforms.
