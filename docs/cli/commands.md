# Command Reference

Complete reference for all `prdtool` commands.

## init

Initialize a new PRD with required metadata.

```bash
prdtool init --title <title> --owner <owner> [--id <id>] [-f <file>]
```

| Flag | Required | Description | Default |
|------|----------|-------------|---------|
| `--title` | Yes | PRD title | |
| `--owner` | Yes | PRD owner name | |
| `--id` | No | PRD ID | Auto-generated |
| `-f, --file` | No | Output file path | `PRD.json` |

**Examples:**

```bash
prdtool init --title "User Authentication" --owner "Jane PM"
prdtool init --title "Search Feature" --owner "John PM" --id PRD-2026-042
prdtool init --title "Mobile App" --owner "Team Lead" -f mobile-prd.json
```

---

## show

Display PRD contents as formatted JSON.

```bash
prdtool show [file] [-f <file>]
```

**Examples:**

```bash
prdtool show
prdtool show my-prd.json
prdtool show -f feature.json
```

---

## validate

Validate a PRD file against the schema.

```bash
prdtool validate [file] [-f <file>]
```

Performs:

- JSON schema validation
- ID format checking (e.g., `REQ-1`, `RISK-2`)
- Traceability verification
- Required field validation

**Examples:**

```bash
prdtool validate
prdtool validate PRD.json
```

**Output:**

```json
{
  "valid": true,
  "errors": [],
  "warnings": ["No metrics defined"]
}
```

---

## score

Score a PRD's quality against the rubric.

```bash
prdtool score [file] [-f <file>] [--verbose] [--json]
```

| Flag | Description |
|------|-------------|
| `-v, --verbose` | Show detailed scoring breakdown |
| `--json` | Output as JSON |

**Scoring Categories:**

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

**Thresholds:**

- ≥8.0 → Approve
- ≥6.5 → Revise
- <6.5 → Human Review
- ≤3.0 → Blocker

**Examples:**

```bash
prdtool score
prdtool score --verbose
prdtool score --json | jq '.overall_score'
```

---

## view

Generate human-readable PRD views.

```bash
prdtool view [file] [-f <file>] [-t <type>] [-o <format>]
```

| Flag | Description | Options | Default |
|------|-------------|---------|---------|
| `-t, --type` | View type | `pm`, `exec` | `pm` |
| `-o, --format` | Output format | `markdown`, `json` | `markdown` |

**View Types:**

- **pm**: Product Manager view - detailed operational information
- **exec**: Executive view - high-level decision summary with scores

**Examples:**

```bash
prdtool view                          # PM view, markdown
prdtool view --type exec              # Executive summary
prdtool view --type pm --format json  # PM view as JSON
prdtool view -t exec -o markdown > summary.md
```

---

## deploy

Generate AI assistant configurations.

```bash
prdtool deploy [-t <target>] [-o <output>]
```

| Flag | Description | Options | Default |
|------|-------------|---------|---------|
| `-t, --target` | Deployment target | `kiro-cli`, `kiro-power`, `claude`, `all` | `kiro-cli` |
| `-o, --output` | Output directory | | Platform-specific |

**Targets:**

- **kiro-cli**: AWS Kiro CLI (MCP config + agent configs)
- **kiro-power**: AWS Kiro IDE Power (POWER.md + mcp.json + steering/)
- **claude**: Claude Code (MCP config)
- **all**: All supported platforms

**Examples:**

```bash
prdtool deploy --target claude
prdtool deploy --target kiro-power --output ~/.kiro/powers/prdtool
prdtool deploy --target all
```

---

## add

Add items to various sections of a PRD. All `add` commands support the `-f, --file` flag.

### add problem

Add a problem statement.

```bash
prdtool add problem --statement <text> [--impact <text>] [--confidence <0-1>]
```

| Flag | Required | Description | Default |
|------|----------|-------------|---------|
| `--statement` | Yes | Problem statement | |
| `--impact` | No | User impact description | |
| `--confidence` | No | Confidence level (0-1) | 0.5 |

```bash
prdtool add problem --statement "Users cannot securely access accounts" --impact "Lost revenue" --confidence 0.8
```

### add persona

Add a user persona.

```bash
prdtool add persona --name <name> [--role <role>] [--pain-point <text>...]
```

| Flag | Required | Description |
|------|----------|-------------|
| `--name` | Yes | Persona name |
| `--role` | No | Persona role |
| `--pain-point` | No | Pain points (repeatable) |

```bash
prdtool add persona --name "Developer Dan" --role "Backend Developer" --pain-point "Slow builds" --pain-point "Complex configs"
```

### add goal

Add a goal.

```bash
prdtool add goal --statement <text>
```

```bash
prdtool add goal --statement "Reduce authentication time by 50%"
```

### add nongoal

Add a non-goal (explicit scope exclusion).

```bash
prdtool add nongoal --statement <text>
```

```bash
prdtool add nongoal --statement "Mobile app support is out of scope"
```

### add solution

Add a solution option.

```bash
prdtool add solution --name <name> [--description <text>] [--tradeoff <text>...]
```

| Flag | Required | Description |
|------|----------|-------------|
| `--name` | Yes | Solution name |
| `--description` | No | Solution description |
| `--tradeoff` | No | Tradeoffs (repeatable) |

```bash
prdtool add solution --name "OAuth 2.0" --description "Industry standard auth" --tradeoff "Complex setup" --tradeoff "Third-party dependency"
```

### add req

Add a functional requirement.

```bash
prdtool add req --description <text> [--priority <level>] [--ac <text>...]
```

| Flag | Required | Description | Default |
|------|----------|-------------|---------|
| `--description` | Yes | Requirement description | |
| `--priority` | No | Priority: `must`, `should`, `could` | `should` |
| `--ac` | No | Acceptance criteria (repeatable) | |

```bash
prdtool add req --description "Support OAuth 2.0 login" --priority must --ac "User can login with Google" --ac "User can login with GitHub"
```

### add nfr

Add a non-functional requirement.

```bash
prdtool add nfr --requirement <text> [--category <cat>]
```

| Flag | Required | Description | Default |
|------|----------|-------------|---------|
| `--requirement` | Yes | NFR description | |
| `--category` | No | Category (see below) | `performance` |

**Categories:** `performance`, `security`, `reliability`, `scalability`, `usability`, `compliance`, `accessibility`

```bash
prdtool add nfr --requirement "API response time < 200ms" --category performance
prdtool add nfr --requirement "All data encrypted at rest" --category security
```

### add metric

Add a success metric.

```bash
prdtool add metric --name <name> [--definition <text>] [--target <text>] [--type <type>]
```

| Flag | Required | Description | Default |
|------|----------|-------------|---------|
| `--name` | Yes | Metric name | |
| `--definition` | No | How metric is calculated | |
| `--target` | No | Target value | |
| `--type` | No | Type: `northstar`, `supporting`, `guardrail` | `supporting` |

```bash
prdtool add metric --name "Login Success Rate" --definition "Successful logins / Total attempts" --target "99.5%" --type northstar
```

### add risk

Add a risk.

```bash
prdtool add risk --description <text> [--impact <level>] [--mitigation <text>]
```

| Flag | Required | Description | Default |
|------|----------|-------------|---------|
| `--description` | Yes | Risk description | |
| `--impact` | No | Impact: `low`, `medium`, `high` | `medium` |
| `--mitigation` | No | Mitigation strategy | |

```bash
prdtool add risk --description "Third-party OAuth provider outage" --impact high --mitigation "Implement fallback local auth"
```

### add decision

Add a decision record.

```bash
prdtool add decision --decision <text> [--rationale <text>] [--by <name>]
```

| Flag | Required | Description |
|------|----------|-------------|
| `--decision` | Yes | Decision made |
| `--rationale` | No | Rationale for decision |
| `--by` | No | Who made the decision |

```bash
prdtool add decision --decision "Use JWT for session management" --rationale "Stateless, scalable" --by "Tech Lead"
```
