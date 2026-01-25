---
name: "prdtool"
displayName: "PRD Tool"
description: "Create, validate, score, and manage Product Requirements Documents with AI assistance. Provides MCP tools for the complete PRD lifecycle."
version: "1.0.0"
keywords:
  - "prd"
  - "product requirements"
  - "requirements document"
  - "product spec"
  - "feature spec"
  - "problem statement"
  - "user persona"
  - "success metrics"
  - "north star"
  - "acceptance criteria"
---

# PRD Tool

Create, validate, score, and manage Product Requirements Documents with AI assistance. Provides MCP tools for the complete PRD lifecycle.

## Onboarding

## Prerequisites

### prdtool

PRD lifecycle tools including init, validate, score, and content management

Verify the server is available:

```bash
which prdtool-mcp || echo "prdtool-mcp not found in PATH"
```



## Available Tools

This power provides the following MCP servers:

### prdtool

PRD lifecycle tools including init, validate, score, and content management

**Command:** `prdtool-mcp`

## Workflows

This power includes steering for the following workflows:

- **exec-summary**: Generates executive summaries from PRDs for leadership review. Produces concise decision-ready documents with key risks and recommendations. Activated when preparing PRDs for leadership or stakeholder review. (triggers: exec summary, executive summary, leadership review, stakeholder review, prepare for review)
- **prd-creation**: Guides through comprehensive PRD creation with problem discovery, persona definition, goal setting, solution exploration, and requirements documentation. Activated when creating new PRDs or discussing product requirements. (triggers: create prd, new prd, start prd, product requirements, requirements document, prd creation)
- **prd-review**: Reviews existing PRDs against quality rubric, identifies gaps, and provides prioritized improvement recommendations. Activated when reviewing, scoring, or improving PRDs. (triggers: review prd, score prd, prd review, improve prd, prd score, quality check)

## Instructions

# PRD Tool Power

# PRD Tool

This plugin provides AI-assisted Product Requirements Document (PRD) creation and management.

## Available Commands

- `/prdtool:create` - Create a new PRD interactively
- `/prdtool:review` - Review and score an existing PRD
- `/prdtool:exec-summary` - Generate executive summary

## MCP Server

The prdtool-mcp server provides these tools:

- `prd_init` - Initialize a new PRD
- `prd_load` - Load PRD contents
- `prd_validate` - Validate structure
- `prd_score` - Score quality (0-10)
- `prd_view` - Generate views
- `prd_add_*` - Add content sections

## Quality Scoring

PRDs are scored across 10 categories:

- Problem Definition (20%)
- Solution Fit (15%)
- User Understanding (10%)
- Market Awareness (10%)
- Scope Discipline (10%)
- Requirements Quality (10%)
- Metrics Quality (10%)
- UX Coverage (5%)
- Technical Feasibility (5%)
- Risk Management (5%)

## Dependencies

- `prdtool-mcp` - MCP server for PRD operations
- `prdtool` - CLI tool (optional)

## Workflows

### Exec Summary Workflow
Generates executive summaries from PRDs for leadership review. Produces concise decision-ready documents with key risks and recommendations. Activated when preparing PRDs for leadership or stakeholder review.

### Prd Creation Workflow
Guides through comprehensive PRD creation with problem discovery, persona definition, goal setting, solution exploration, and requirements documentation. Activated when creating new PRDs or discussing product requirements.

### Prd Review Workflow
Reviews existing PRDs against quality rubric, identifies gaps, and provides prioritized improvement recommendations. Activated when reviewing, scoring, or improving PRDs.


