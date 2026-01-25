---
description: Review and score an existing PRD with improvement recommendations
---

# Review

Review and score an existing PRD with improvement recommendations

## Usage

```
/review [path]
```

## Arguments

- **path**: Path to PRD file (defaults to PRD.json)

## Process

1. Load the PRD document
2. Validate structure and references
3. Score against quality rubric
4. Analyze each category for gaps
5. Generate prioritized recommendations

## Dependencies

- `prdtool-mcp`

## Instructions

Review an existing PRD for quality and completeness.

## Process

1. **Load**: Use `prd_load` to examine the document
2. **Validate**: Use `prd_validate` to check structure
   - Required fields present
   - IDs properly formatted
   - References valid
3. **Score**: Use `prd_score` to evaluate quality
   - 10 weighted categories
   - Overall 0-10 score
   - Decision recommendation
4. **Analyze**: Examine each category
   - Identify gaps and blockers
   - Find missing evidence
   - Check traceability
5. **Recommend**: Provide specific improvements
   - Prioritize by score impact
   - Include concrete examples
   - Suggest specific additions

## Output Format

### Quality Score: X.X/10
**Decision**: APPROVE | REVISE | HUMAN REVIEW

### Category Breakdown
| Category | Score | Issues |
|----------|-------|--------|
| Problem Definition | X.X | ... |
| ... | | |

### Top Recommendations
1. [Highest impact fix]
2. [Second priority]
3. [Third priority]

### Blockers (if any)
- Critical issues that must be resolved
