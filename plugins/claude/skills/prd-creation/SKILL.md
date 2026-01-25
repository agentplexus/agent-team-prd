---
name: prd-creation
description: Guides through comprehensive PRD creation with problem discovery, persona definition, goal setting, solution exploration, and requirements documentation. Activated when creating new PRDs or discussing product requirements.
triggers: [create prd, new prd, start prd, product requirements, requirements document, prd creation]
dependencies: [prdtool-mcp]
---

# Prd Creation

Guides through comprehensive PRD creation with problem discovery, persona definition, goal setting, solution exploration, and requirements documentation. Activated when creating new PRDs or discussing product requirements.

## Instructions

# PRD Creation Skill

You are guiding the user through creating a comprehensive Product Requirements Document.

## Philosophy

- **Evidence over opinion**: Every claim needs supporting data
- **Users first**: Start with who we're building for
- **Scope discipline**: Non-goals are as important as goals
- **Multiple options**: Never commit to the first solution

## Creation Flow

### Phase 1: Problem Discovery

Before any solution, deeply understand the problem:

1. What evidence shows this is a real problem?
2. How many users are affected?
3. What's the cost of not solving it?
4. Is this the root cause or a symptom?

Use `prd_add_problem` with:
- Clear problem statement
- User impact description
- Confidence level (0-1)

### Phase 2: User Definition

Define 1-3 target personas:

1. Who experiences this problem most acutely?
2. What are their top 3 pain points?
3. What workarounds do they use today?
4. What would success look like for them?

Use `prd_add_persona` with:
- Descriptive name and archetype
- Specific pain points (not generic)

### Phase 3: Scope Setting

Establish clear boundaries:

**Goals** (use `prd_add_goal`):
- SMART: Specific, Measurable, Achievable, Relevant, Time-bound
- Link to success metrics

**Non-Goals** (use `prd_add_nongoal`):
- Explicitly state what you won't do
- Prevents scope creep later
- Equally important as goals

### Phase 4: Solution Exploration

Explore 2-3 options before committing:

For each option (use `prd_add_solution`):
- Brief description
- Pros and cons
- Cost estimate (T-shirt: S/M/L/XL)
- Technical complexity
- Problems it addresses

Then select one (use `prd_select_solution`):
- Clear rationale for choice
- Acknowledge tradeoffs accepted

### Phase 5: Requirements

Document functional requirements:

For each (use `prd_add_requirement`):
- Clear, testable statement
- Priority: P0 (launch blocker), P1 (important), P2 (nice-to-have)
- Acceptance criteria
- Trace to goals/problems

### Phase 6: Success Metrics

Define measurable outcomes:

1. **North Star** (one primary metric)
   - What single number indicates success?
   - Use `prd_add_metric --type=northstar`

2. **Supporting Metrics**
   - 2-4 metrics that ladder to North Star
   - Use `prd_add_metric --type=supporting`

3. **Guardrail Metrics**
   - What shouldn't regress?
   - Use `prd_add_metric --type=guardrail`

### Phase 7: Risk Management

Identify and mitigate risks:

For each (use `prd_add_risk`):
- Impact: high/medium/low
- Likelihood: high/medium/low
- Mitigation strategy
- Owner for mitigation

## Quality Checks

Before finalizing:

1. Run `prd_validate` - fix any structural issues
2. Run `prd_score` - aim for >= 8.0
3. Review category breakdowns
4. Address any blockers

## Common Pitfalls

- Vague problems without evidence
- Generic personas (not real users)
- Missing non-goals
- Jumping to solutions
- Unmeasurable metrics
- Risks without mitigations

