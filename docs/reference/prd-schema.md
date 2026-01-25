# PRD Schema Reference

Complete reference for the PRD JSON schema.

## Overview

A PRD document is a JSON file with the following top-level structure:

```json
{
  "metadata": { ... },
  "context": { ... },
  "problem": { ... },
  "users": { ... },
  "market": { ... },
  "goals_and_non_goals": { ... },
  "solution": { ... },
  "requirements": { ... },
  "ux": { ... },
  "technical": { ... },
  "metrics": { ... },
  "risks_and_assumptions": { ... },
  "decisions": { ... },
  "reviews_and_scoring": { ... },
  "revision_history": [ ... ]
}
```

**Required sections:** `metadata`, `problem`, `goals_and_non_goals`

**Optional sections:** All others (marked with `omitempty`)

---

## Metadata

Control plane information for the PRD.

```json
{
  "metadata": {
    "prd_id": "PRD-2026-001",
    "title": "User Authentication System",
    "owner": "Jane Smith",
    "status": "draft",
    "created_at": "2026-01-15T10:00:00Z",
    "last_updated_at": "2026-01-20T14:30:00Z",
    "version": "1.0.0",
    "related_prds": ["PRD-2025-042"]
  }
}
```

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `prd_id` | string | Yes | Unique identifier (format: `PRD-YYYY-NNN`) |
| `title` | string | Yes | PRD title |
| `owner` | string | Yes | Owner name |
| `status` | Status | Yes | Lifecycle status |
| `created_at` | datetime | No | Creation timestamp |
| `last_updated_at` | datetime | No | Last update timestamp |
| `version` | string | No | Semantic version |
| `related_prds` | string[] | No | Related PRD IDs |

### Status Values

| Value | Description |
|-------|-------------|
| `draft` | Initial creation, work in progress |
| `in_review` | Submitted for review |
| `revised` | Updated after review feedback |
| `approved` | Ready for implementation |
| `deprecated` | No longer active |

---

## Context

Background and strategic context.

```json
{
  "context": {
    "product_area": "Authentication",
    "strategic_theme": "Security Modernization",
    "business_context": "Part of Q1 security initiative",
    "why_now": "Recent security audit findings",
    "constraints": ["Must integrate with existing SSO"],
    "out_of_scope_context": ["Mobile biometrics"]
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `product_area` | string | Product area or domain |
| `strategic_theme` | string | Strategic initiative this supports |
| `business_context` | string | Business background |
| `why_now` | string | Urgency justification |
| `constraints` | string[] | Known constraints |
| `out_of_scope_context` | string[] | Context for scope exclusions |

---

## Problem

Problem definition with evidence.

```json
{
  "problem": {
    "primary_problem": {
      "id": "PROB-1",
      "statement": "Users cannot securely access accounts",
      "user_impact": "30% of support tickets are password-related",
      "confidence": 0.9,
      "evidence": [
        {
          "type": "analytics",
          "description": "Password reset rate: 15% monthly",
          "strength": "high",
          "sample_size": 10000
        }
      ],
      "assumptions": ["Users prefer passwordless options"]
    },
    "secondary_problems": [],
    "root_causes": ["Complex password requirements"],
    "non_problems": ["Users wanting biometric auth"]
  }
}
```

### Problem Object

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique ID (format: `PROB-N`) |
| `statement` | string | Problem statement |
| `user_impact` | string | Impact on users |
| `confidence` | float | Confidence level (0-1) |
| `evidence` | Evidence[] | Supporting evidence |
| `assumptions` | string[] | Underlying assumptions |

### Evidence Object

| Field | Type | Description |
|-------|------|-------------|
| `type` | EvidenceType | Evidence category |
| `description` | string | Evidence description |
| `strength` | Strength | Evidence quality |
| `sample_size` | int | Sample size (if applicable) |
| `source_url` | string | Source reference |

**Evidence Types:** `interview`, `survey`, `analytics`, `support_ticket`, `market_research`, `assumption`

**Strength Levels:** `low`, `medium`, `high`

---

## Users

User personas and target audience.

```json
{
  "users": {
    "personas": [
      {
        "id": "PER-1",
        "name": "Enterprise Admin",
        "role": "IT Administrator",
        "goals": ["Manage user credentials securely"],
        "pain_points": ["Complex compliance requirements"],
        "behaviors": ["Uses SSO dashboard daily"],
        "constraints": ["Limited time for training"],
        "confidence": 0.8,
        "assumptions": ["Has basic security knowledge"]
      }
    ],
    "primary_persona_id": "PER-1",
    "excluded_users": ["Consumer users"]
  }
}
```

### Persona Object

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique ID (format: `PER-N`) |
| `name` | string | Persona name |
| `role` | string | Job role or function |
| `goals` | string[] | What they want to achieve |
| `pain_points` | string[] | Current frustrations |
| `behaviors` | string[] | Relevant behaviors |
| `constraints` | string[] | Limitations they face |
| `confidence` | float | Confidence in persona (0-1) |
| `assumptions` | string[] | Assumptions about persona |

---

## Market

Competitive landscape and alternatives.

```json
{
  "market": {
    "alternatives": [
      {
        "id": "ALT-1",
        "name": "Auth0",
        "type": "competitor",
        "strengths": ["Easy integration", "Good docs"],
        "weaknesses": ["Expensive at scale"]
      }
    ],
    "differentiation": ["Deeper enterprise integration"],
    "market_risks": ["Commoditization of auth"]
  }
}
```

### Alternative Object

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique ID (format: `ALT-N`) |
| `name` | string | Alternative name |
| `type` | AlternativeType | Category |
| `strengths` | string[] | Competitive strengths |
| `weaknesses` | string[] | Competitive weaknesses |

**Alternative Types:** `competitor`, `workaround`, `do_nothing`, `internal`

---

## Goals and Non-Goals

Explicit scope boundaries.

```json
{
  "goals_and_non_goals": {
    "goals": [
      {
        "id": "GOAL-1",
        "statement": "Reduce password reset tickets by 50%",
        "success_metric_ids": ["MET-1"]
      }
    ],
    "non_goals": [
      "Mobile biometric integration",
      "Custom auth protocols"
    ],
    "success_criteria": ["All must-have requirements implemented"]
  }
}
```

### Goal Object

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique ID (format: `GOAL-N`) |
| `statement` | string | Goal statement |
| `success_metric_ids` | string[] | Linked metric IDs |

---

## Solution

Solution options and selection.

```json
{
  "solution": {
    "solution_options": [
      {
        "id": "SOL-1",
        "name": "OAuth 2.0 + Magic Links",
        "description": "Social login with email fallback",
        "problems_addressed": ["PROB-1"],
        "benefits": ["Industry standard", "User friendly"],
        "tradeoffs": ["Email deliverability dependency"],
        "risks": ["Third-party outage"]
      }
    ],
    "selected_solution_id": "SOL-1",
    "solution_rationale": "Best balance of security and UX",
    "confidence": 0.85
  }
}
```

### SolutionOption Object

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique ID (format: `SOL-N`) |
| `name` | string | Solution name |
| `description` | string | Solution description |
| `problems_addressed` | string[] | Problem IDs addressed |
| `benefits` | string[] | Expected benefits |
| `tradeoffs` | string[] | Known tradeoffs |
| `risks` | string[] | Associated risks |

---

## Requirements

Functional and non-functional requirements.

```json
{
  "requirements": {
    "functional": [
      {
        "id": "REQ-1",
        "description": "Support Google OAuth login",
        "priority": "must",
        "derived_from": ["GOAL-1"],
        "acceptance_criteria": [
          "User can click 'Sign in with Google'",
          "New users auto-registered on first login"
        ],
        "assumptions": ["Google OAuth API remains stable"]
      }
    ],
    "non_functional": [
      {
        "id": "NFR-1",
        "category": "performance",
        "requirement": "Login completes in under 3 seconds"
      }
    ]
  }
}
```

### Requirement Object (Functional)

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique ID (format: `REQ-N`) |
| `description` | string | Requirement description |
| `priority` | Priority | MoSCoW priority |
| `derived_from` | string[] | Source goal/problem IDs |
| `acceptance_criteria` | string[] | Testable criteria |
| `assumptions` | string[] | Underlying assumptions |

**Priority Values:** `must`, `should`, `could`

### NFR Object (Non-Functional)

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique ID (format: `NFR-N`) |
| `category` | NFRCategory | Requirement category |
| `requirement` | string | Requirement description |

**NFR Categories:** `performance`, `security`, `reliability`, `scalability`, `usability`, `compliance`, `accessibility`

---

## UX

User experience flows.

```json
{
  "ux": {
    "user_journeys": [
      {
        "id": "UJ-1",
        "persona_id": "PER-1",
        "happy_path": [
          "User clicks login",
          "Selects Google OAuth",
          "Redirected to dashboard"
        ],
        "edge_cases": ["User cancels OAuth flow"],
        "failure_states": ["OAuth provider unavailable"]
      }
    ],
    "ux_risks": ["Learning curve for new auth flow"]
  }
}
```

### UserJourney Object

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique ID (format: `UJ-N`) |
| `persona_id` | string | Associated persona |
| `happy_path` | string[] | Expected flow steps |
| `edge_cases` | string[] | Edge case handling |
| `failure_states` | string[] | Error scenarios |

---

## Technical

Technical feasibility assessment.

```json
{
  "technical": {
    "constraints": ["Must use existing database"],
    "dependencies": ["Auth service v2", "User service"],
    "risks": ["Token rotation complexity"],
    "complexity_estimate": "medium",
    "confidence": 0.75
  }
}
```

| Field | Type | Description |
|-------|------|-------------|
| `constraints` | string[] | Technical constraints |
| `dependencies` | string[] | System dependencies |
| `risks` | string[] | Technical risks |
| `complexity_estimate` | Complexity | Effort estimate |
| `confidence` | float | Assessment confidence (0-1) |

**Complexity Values:** `low`, `medium`, `high`

---

## Metrics

Success metrics definition.

```json
{
  "metrics": {
    "north_star": {
      "id": "MET-1",
      "name": "Password Reset Tickets",
      "definition": "Weekly count of password-related support tickets",
      "target": "50% reduction",
      "baseline": "100 tickets/week"
    },
    "supporting_metrics": [
      {
        "id": "MET-2",
        "name": "Time to Login",
        "definition": "P95 time from landing to authenticated",
        "target": "<5 seconds"
      }
    ],
    "guardrail_metrics": [
      {
        "id": "MET-3",
        "name": "Login Success Rate",
        "definition": "Successful logins / Total attempts",
        "target": ">99%"
      }
    ],
    "instrumentation_gaps": ["OAuth error tracking"]
  }
}
```

### Metric Object

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique ID (format: `MET-N`) |
| `name` | string | Metric name |
| `definition` | string | How metric is calculated |
| `target` | string | Target value |
| `baseline` | string | Current baseline |

**Metric Types:**

- `north_star`: Primary success measure (one per PRD)
- `supporting_metrics`: Secondary success indicators
- `guardrail_metrics`: Metrics that should not regress

---

## Risks and Assumptions

Risk management and key assumptions.

```json
{
  "risks_and_assumptions": {
    "assumptions": [
      {
        "id": "ASM-1",
        "statement": "Users prefer passwordless options",
        "confidence": 0.7,
        "validation_plan": "A/B test with 1000 users"
      }
    ],
    "risks": [
      {
        "id": "RISK-1",
        "description": "OAuth provider outage",
        "impact": "high",
        "likelihood": "low",
        "mitigation": "Implement magic link fallback"
      }
    ],
    "open_questions": ["Token expiration policy?"]
  }
}
```

### Assumption Object

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique ID (format: `ASM-N`) |
| `statement` | string | Assumption statement |
| `confidence` | float | Confidence level (0-1) |
| `validation_plan` | string | How to validate |

### Risk Object

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique ID (format: `RISK-N`) |
| `description` | string | Risk description |
| `impact` | Impact | Severity if realized |
| `likelihood` | Impact | Probability |
| `mitigation` | string | Mitigation strategy |

**Impact/Likelihood Values:** `low`, `medium`, `high`

---

## Decisions

Immutable decision records.

```json
{
  "decisions": {
    "records": [
      {
        "id": "DEC-1",
        "decision": "Use JWT for session management",
        "rationale": "Stateless, scalable architecture",
        "alternatives_considered": ["Session cookies", "Server-side sessions"],
        "made_by": "Tech Lead",
        "date": "2026-01-15",
        "confidence": 0.9
      }
    ]
  }
}
```

### DecisionRecord Object

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique ID (format: `DEC-N`) |
| `decision` | string | Decision made |
| `rationale` | string | Reasoning |
| `alternatives_considered` | string[] | Options evaluated |
| `made_by` | string | Decision maker |
| `date` | string | Decision date |
| `confidence` | float | Confidence level (0-1) |

---

## ID Format Reference

All IDs follow a consistent format: `PREFIX-N`

| Prefix | Object Type |
|--------|-------------|
| `PRD` | PRD document |
| `PROB` | Problem |
| `PER` | Persona |
| `ALT` | Alternative |
| `GOAL` | Goal |
| `SOL` | Solution |
| `REQ` | Functional Requirement |
| `NFR` | Non-Functional Requirement |
| `UJ` | User Journey |
| `MET` | Metric |
| `ASM` | Assumption |
| `RISK` | Risk |
| `DEC` | Decision |

IDs are auto-generated by the CLI and MCP tools. When adding items, the next available ID is assigned automatically.

---

## Complete Example

See the [CLI Examples](../cli/examples.md) page for a complete PRD creation walkthrough.
