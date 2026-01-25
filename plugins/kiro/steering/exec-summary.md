# Executive Summary Skill

You are preparing a PRD for executive review.

## Purpose

Executives need:
- Quick understanding of the proposal
- Clear recommendation (approve/revise/reject)
- Key risks and mitigations
- Resource requirements
- Decision points

## Process

### Step 1: Score the PRD

```
prd_score path=PRD.json
```

This provides:
- Overall quality score
- Decision recommendation
- Category breakdown
- Identified blockers

### Step 2: Generate Executive View

```
prd_view path=PRD.json format=exec
```

This extracts:
- Title and status
- Problem summary
- Proposed solution
- Success metrics
- Key risks

### Step 3: Frame the Decision

Based on the score:

**Score >= 8.0 (APPROVE)**
- Strong recommendation to proceed
- Highlight readiness
- Note any minor conditions

**Score 6.5-7.9 (CONDITIONAL)**
- Can proceed with specific conditions
- List required fixes before launch
- Identify owners for fixes

**Score < 6.5 (ESCALATE)**
- Not ready for approval
- Identify critical gaps
- Recommend next steps

## Output Format

```markdown
# Executive Summary: [PRD Title]

**PRD ID**: [PRD-YYYY-NNN]
**Status**: [Draft | In Review | Approved]
**Owner**: [Name]
**Quality Score**: X.X / 10.0

---

## Recommendation: [APPROVE | CONDITIONAL | ESCALATE]

[1-2 sentence summary of recommendation]

---

## Problem

[2-3 sentences describing the problem with evidence]

**Impact**: [Quantified user/business impact]

---

## Proposed Solution

[2-3 sentences describing selected solution]

**Key Tradeoffs**:
- [Tradeoff 1]
- [Tradeoff 2]

---

## Success Metrics

| Metric | Target | Current |
|--------|--------|---------||
| [North Star] | [Target] | [Baseline] |
| [Supporting 1] | [Target] | [Baseline] |
| [Guardrail 1] | [Threshold] | [Current] |

---

## Key Risks

| Risk | Impact | Likelihood | Mitigation |
|------|--------|------------|------------|
| [Risk 1] | High/Med/Low | High/Med/Low | [Strategy] |
| [Risk 2] | | | |

---

## Resource Requirements

- **Engineering**: [Estimate]
- **Design**: [Estimate]
- **Timeline**: [Estimate]

---

## Decision Required

[Specific ask: What decision do you need from leadership?]

**Options**:
1. [Option A] - [Implication]
2. [Option B] - [Implication]

---

## Appendix

### Quality Breakdown

| Category | Score |
|----------|-------|
| Problem Definition | X.X |
| Solution Fit | X.X |
| ... | |

### Blockers (if any)

- [Blocker 1]
- [Blocker 2]
```

## Presentation Tips

1. **Lead with the ask**: What decision do you need?
2. **Be concise**: Executives scan, don't read
3. **Quantify impact**: Numbers speak louder than words
4. **Surface risks early**: Don't bury bad news
5. **Provide options**: Give choices, not ultimatums

## Common Mistakes

- Too much detail (save for appendix)
- Burying the recommendation
- Vague success metrics
- Missing resource requirements
- No clear decision point