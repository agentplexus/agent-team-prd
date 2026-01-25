# PRD Review Skill

You are reviewing a PRD for quality and completeness.

## Review Framework

### Step 1: Load and Understand

```
prd_load path=PRD.json
```

Read through the entire document to understand:
- What problem is being solved?
- Who is the target user?
- What solution is proposed?
- What are the success criteria?

### Step 2: Structural Validation

```
prd_validate path=PRD.json
```

Check for:
- Required fields present
- IDs properly formatted (PROB-1, PER-1, GOAL-1, etc.)
- Cross-references valid
- No orphaned elements

### Step 3: Quality Scoring

```
prd_score path=PRD.json
```

Analyze scores across 10 categories:

| Category | Weight | What to Check |
|----------|--------|---------------|
| Problem Definition | 20% | Evidence, clarity, impact |
| Solution Fit | 15% | Addresses problems, tradeoffs documented |
| User Understanding | 10% | Validated personas, real pain points |
| Market Awareness | 10% | Alternatives analyzed |
| Scope Discipline | 10% | Clear goals AND non-goals |
| Requirements Quality | 10% | Testable, prioritized, traced |
| Metrics Quality | 10% | North Star defined, measurable |
| UX Coverage | 5% | User journeys documented |
| Technical Feasibility | 5% | Constraints, dependencies |
| Risk Management | 5% | Risks with mitigations |

### Step 4: Gap Analysis

For each low-scoring category (< 7.0):

1. What's specifically missing?
2. What would improve the score?
3. What's the effort to fix?
4. Who should own the fix?

### Step 5: Prioritized Recommendations

Rank improvements by:
1. Score impact (higher weight categories first)
2. Effort required (quick wins first)
3. Blocker status (blockers must be fixed)

## Output Template

```markdown
## PRD Quality Review: [Title]

### Overall Score: X.X / 10.0
**Decision**: [APPROVE | REVISE | HUMAN REVIEW | REJECT]

### Category Breakdown

| Category | Score | Status |
|----------|-------|--------|
| Problem Definition | X.X | ✓/⚠/✗ |
| Solution Fit | X.X | ... |
| ... | | |

### Strengths
- [What's done well]

### Blockers (must fix)
- [Critical issues]

### Top 3 Recommendations

1. **[Category]**: [Specific action]
   - Impact: +X.X points
   - Effort: [Low/Medium/High]

2. **[Category]**: [Specific action]
   - Impact: +X.X points
   - Effort: [Low/Medium/High]

3. **[Category]**: [Specific action]
   - Impact: +X.X points
   - Effort: [Low/Medium/High]

### Next Steps
1. [Immediate action]
2. [Follow-up action]
```

## Decision Thresholds

- **>= 8.0**: APPROVE - Ready for implementation
- **>= 6.5**: REVISE - Minor issues, can proceed with fixes
- **< 6.5**: HUMAN REVIEW - Significant gaps need attention
- **<= 3.0**: REJECT - Critical blockers, not ready

## Common Issues

### Problem Definition
- No evidence cited
- Impact not quantified
- Problem too vague

### User Understanding
- Generic personas
- No validated pain points
- Missing user research

### Scope Discipline
- No non-goals defined
- Goals not measurable
- Scope too broad

### Requirements Quality
- Missing acceptance criteria
- No priority assigned
- Not traceable to goals

### Metrics Quality
- No North Star metric
- Metrics not measurable
- Missing baselines/targets