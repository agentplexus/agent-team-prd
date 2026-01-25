# CLI Overview

The `prdtool` CLI provides commands for creating, managing, validating, and scoring Product Requirements Documents.

## Command Structure

```
prdtool [command] [subcommand] [flags]
```

## Global Flags

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--file` | `-f` | PRD file path | `PRD.json` |
| `--help` | `-h` | Help for any command | |
| `--version` | `-v` | Version information | |

## Commands

### Document Management

| Command | Description |
|---------|-------------|
| `init` | Initialize a new PRD with required metadata |
| `show` | Display PRD contents as JSON |
| `validate` | Validate PRD against schema |
| `score` | Score PRD quality against rubric |
| `view` | Generate human-readable views |

### Content Addition

| Command | Description |
|---------|-------------|
| `add problem` | Add a problem statement |
| `add persona` | Add a user persona |
| `add goal` | Add a goal |
| `add nongoal` | Add a non-goal |
| `add solution` | Add a solution option |
| `add req` | Add a functional requirement |
| `add nfr` | Add a non-functional requirement |
| `add metric` | Add a success metric |
| `add risk` | Add a risk |
| `add decision` | Add a decision record |

### Deployment

| Command | Description |
|---------|-------------|
| `deploy` | Generate AI assistant configurations |

## Typical Workflow

1. **Initialize** a new PRD:
   ```bash
   prdtool init --title "Feature Name" --owner "PM Name"
   ```

2. **Add content** incrementally:
   ```bash
   prdtool add problem --statement "..."
   prdtool add persona --name "..." --role "..."
   prdtool add goal --statement "..."
   prdtool add req --description "..." --priority must
   ```

3. **Validate** the structure:
   ```bash
   prdtool validate
   ```

4. **Score** for quality:
   ```bash
   prdtool score
   ```

5. **Generate views** for stakeholders:
   ```bash
   prdtool view --type pm      # For product managers
   prdtool view --type exec    # For executives
   ```

See [Command Reference](commands.md) for detailed documentation of each command, or [Examples](examples.md) for complete workflows.
