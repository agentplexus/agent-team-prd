---
description: Generate an executive summary for leadership review
---

# Exec Summary

Generate an executive summary for leadership review

## Usage

```
/exec-summary [path]
```

## Arguments

- **path**: Path to PRD file (defaults to PRD.json)

## Process

1. Load and score the PRD
2. Generate executive view
3. Extract decision recommendation
4. Highlight risks and blockers
5. Format for leadership presentation

## Dependencies

- `prdtool-mcp`

## Instructions

Generate an executive summary suitable for leadership review.

## Process

1. **Score**: Use `prd_score` to get quality assessment
2. **Generate**: Use `prd_view --format=exec` for executive view
3. **Summarize**: Present decision recommendation
   - GO / NO-GO / CONDITIONAL
   - Key factors for decision
4. **Highlight**: Surface critical information
   - Top risks and blockers
   - Required resources
   - Timeline dependencies

## Output Format

### Executive Summary: [PRD Title]

**Status**: [Draft | In Review | Approved]
**Quality Score**: X.X/10
**Recommendation**: [APPROVE | REVISE | ESCALATE]

#### Problem
[1-2 sentence problem statement with evidence]

#### Proposed Solution
[Selected solution with key tradeoffs]

#### Success Metrics
- **North Star**: [Primary metric with target]
- **Supporting**: [Key supporting metrics]

#### Key Risks
1. [Risk 1 with mitigation]
2. [Risk 2 with mitigation]

#### Decision Required
[Specific ask of leadership]
