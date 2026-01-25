---
description: Create a new PRD interactively with guided workflow
---

# Create

Create a new PRD interactively with guided workflow

## Usage

```
/create [title]
```

## Arguments

- **title**: PRD title (optional, will be prompted if not provided)

## Process

1. Initialize PRD with title and owner
2. Discover and document the primary problem
3. Define 1-3 user personas with pain points
4. Set SMART goals and explicit non-goals
5. Explore 2-3 solution options
6. Select solution with documented rationale
7. Document functional requirements with acceptance criteria
8. Define North Star and supporting metrics
9. Identify risks with mitigation strategies
10. Validate and score the final document

## Dependencies

- `prdtool-mcp`

## Instructions

Guide the user through creating a comprehensive PRD.

## Process

1. **Initialize**: Create PRD.json with `prd_init`
2. **Problem Discovery**: Ask probing questions about the problem
   - What evidence exists for this problem?
   - Who experiences this problem?
   - What's the impact if unsolved?
3. **User Definition**: Define target personas
   - Name, role, archetype
   - Key pain points
   - Current workarounds
4. **Scope Setting**: Establish goals and non-goals
   - SMART goals with metrics
   - Explicit non-goals to prevent scope creep
5. **Solution Exploration**: Explore 2-3 options
   - Pros, cons, tradeoffs
   - Cost and complexity estimates
   - Select with documented rationale
6. **Requirements**: Document functional requirements
   - Priority (P0-P2)
   - Acceptance criteria
   - Traceability to goals
7. **Metrics**: Define success criteria
   - One North Star metric
   - Supporting metrics
   - Guardrail metrics
8. **Risks**: Identify and mitigate risks
   - Impact and likelihood
   - Mitigation strategy
   - Owner assignment
9. **Validate**: Score the completed PRD
   - Aim for >= 8.0 score
   - Address any blockers
