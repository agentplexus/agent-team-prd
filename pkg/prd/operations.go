package prd

import (
	"fmt"
	"regexp"
	"strconv"
)

// NextID generates the next ID for a given prefix based on existing IDs in the PRD.
func NextID(p *PRD, prefix string) string {
	max := 0
	pattern := regexp.MustCompile(fmt.Sprintf(`^%s-(\d+)$`, prefix))

	checkID := func(id string) {
		if matches := pattern.FindStringSubmatch(id); len(matches) == 2 {
			if num, err := strconv.Atoi(matches[1]); err == nil && num > max {
				max = num
			}
		}
	}

	// Check problem IDs
	if p.Problem != nil {
		checkID(p.Problem.ID)
		for _, prob := range p.Problem.SecondaryProblems {
			checkID(prob.ID)
		}
	}

	// Check persona IDs
	for _, persona := range p.Personas {
		checkID(persona.ID)
	}

	// Check alternative IDs (market)
	if p.Market != nil {
		for _, alt := range p.Market.Alternatives {
			checkID(alt.ID)
		}
	}

	// Check objective IDs
	for _, obj := range p.Objectives.BusinessObjectives {
		checkID(obj.ID)
	}
	for _, obj := range p.Objectives.ProductGoals {
		checkID(obj.ID)
	}
	for _, m := range p.Objectives.SuccessMetrics {
		checkID(m.ID)
	}

	// Check solution IDs
	if p.Solution != nil {
		for _, sol := range p.Solution.SolutionOptions {
			checkID(sol.ID)
		}
	}

	// Check requirement IDs
	for _, req := range p.Requirements.Functional {
		checkID(req.ID)
	}
	for _, nfr := range p.Requirements.NonFunctional {
		checkID(nfr.ID)
	}

	// Check user story IDs
	for _, story := range p.UserStories {
		checkID(story.ID)
	}

	// Check risk IDs
	for _, risk := range p.Risks {
		checkID(risk.ID)
	}

	// Check decision IDs
	if p.Decisions != nil {
		for _, dec := range p.Decisions.Records {
			checkID(dec.ID)
		}
	}

	// Check phase IDs (roadmap)
	for _, phase := range p.Roadmap.Phases {
		checkID(phase.ID)
	}

	return fmt.Sprintf("%s-%d", prefix, max+1)
}

// SetProblemStatement sets the problem statement in the executive summary
// and optionally creates a detailed ProblemDefinition.
func SetProblemStatement(p *PRD, statement, impact string, confidence float64) {
	// Set in executive summary
	p.ExecutiveSummary.ProblemStatement = statement

	// Create or update detailed problem definition
	if p.Problem == nil {
		p.Problem = &ProblemDefinition{}
	}

	if p.Problem.ID == "" {
		p.Problem.ID = NextID(p, "PROB")
	}
	p.Problem.Statement = statement
	p.Problem.UserImpact = impact
	p.Problem.Confidence = confidence
}

// AddPersona adds a user persona to the PRD.
// Returns the generated ID.
func AddPersona(p *PRD, name, role string, painPoints []string) string {
	id := NextID(p, "PER")
	persona := Persona{
		ID:         id,
		Name:       name,
		Role:       role,
		PainPoints: painPoints,
	}

	p.Personas = append(p.Personas, persona)

	// Set as primary if first persona
	if len(p.Personas) == 1 {
		persona.IsPrimary = true
		p.Personas[0] = persona
	}

	return id
}

// AddBusinessObjective adds a business objective to the PRD.
// Returns the generated ID.
func AddBusinessObjective(p *PRD, description, rationale string) string {
	id := NextID(p, "BO")
	obj := Objective{
		ID:          id,
		Description: description,
		Rationale:   rationale,
	}

	p.Objectives.BusinessObjectives = append(p.Objectives.BusinessObjectives, obj)
	return id
}

// AddProductGoal adds a product goal to the PRD.
// Returns the generated ID.
func AddProductGoal(p *PRD, description, rationale string) string {
	id := NextID(p, "PG")
	goal := Objective{
		ID:          id,
		Description: description,
		Rationale:   rationale,
	}

	p.Objectives.ProductGoals = append(p.Objectives.ProductGoals, goal)
	return id
}

// AddOutOfScope adds an out-of-scope item (non-goal) to the PRD.
func AddOutOfScope(p *PRD, item string) {
	p.OutOfScope = append(p.OutOfScope, item)
}

// AddSolution adds a solution option to the PRD.
// Returns the generated ID.
func AddSolution(p *PRD, name, description string, tradeoffs []string) string {
	if p.Solution == nil {
		p.Solution = &SolutionDefinition{}
	}

	id := NextID(p, "SOL")
	solution := SolutionOption{
		ID:          id,
		Name:        name,
		Description: description,
		Tradeoffs:   tradeoffs,
	}

	p.Solution.SolutionOptions = append(p.Solution.SolutionOptions, solution)
	return id
}

// SelectSolution selects a solution by ID and records the rationale.
// Returns true if the solution was found and selected.
func SelectSolution(p *PRD, solutionID, rationale string) bool {
	if p.Solution == nil {
		return false
	}

	// Verify solution exists
	found := false
	for _, opt := range p.Solution.SolutionOptions {
		if opt.ID == solutionID {
			found = true
			break
		}
	}

	if !found {
		return false
	}

	p.Solution.SelectedSolutionID = solutionID
	p.Solution.SolutionRationale = rationale

	// Also set in executive summary
	for _, opt := range p.Solution.SolutionOptions {
		if opt.ID == solutionID {
			p.ExecutiveSummary.ProposedSolution = opt.Description
			break
		}
	}

	return true
}

// AddFunctionalRequirement adds a functional requirement to the PRD.
// Returns the generated ID.
func AddFunctionalRequirement(p *PRD, title, description string, priority MoSCoW) string {
	id := NextID(p, "FR")
	req := FunctionalRequirement{
		ID:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
	}

	p.Requirements.Functional = append(p.Requirements.Functional, req)
	return id
}

// AddNonFunctionalRequirement adds a non-functional requirement to the PRD.
// Returns the generated ID.
func AddNonFunctionalRequirement(p *PRD, category NFRCategory, title, description, target string, priority MoSCoW) string {
	id := NextID(p, "NFR")
	nfr := NonFunctionalRequirement{
		ID:          id,
		Category:    category,
		Title:       title,
		Description: description,
		Target:      target,
		Priority:    priority,
	}

	p.Requirements.NonFunctional = append(p.Requirements.NonFunctional, nfr)
	return id
}

// AddSuccessMetric adds a success metric to the PRD.
// Returns the generated ID.
func AddSuccessMetric(p *PRD, name, description, target string) string {
	id := NextID(p, "SM")
	metric := SuccessMetric{
		ID:          id,
		Name:        name,
		Description: description,
		Target:      target,
	}

	p.Objectives.SuccessMetrics = append(p.Objectives.SuccessMetrics, metric)
	return id
}

// AddRisk adds a risk to the PRD.
// Returns the generated ID.
func AddRisk(p *PRD, description string, probability RiskProbability, impact RiskImpact, mitigation string) string {
	id := NextID(p, "RISK")
	risk := Risk{
		ID:          id,
		Description: description,
		Probability: probability,
		Impact:      impact,
		Mitigation:  mitigation,
	}

	p.Risks = append(p.Risks, risk)
	return id
}

// AddDecision adds a decision record to the PRD.
// Returns the generated ID.
func AddDecision(p *PRD, decision, rationale, madeBy string) string {
	if p.Decisions == nil {
		p.Decisions = &DecisionsDefinition{}
	}

	id := NextID(p, "DEC")
	record := DecisionRecord{
		ID:        id,
		Decision:  decision,
		Rationale: rationale,
		MadeBy:    madeBy,
	}

	p.Decisions.Records = append(p.Decisions.Records, record)
	return id
}

// UpdateStatus changes the PRD status.
func UpdateStatus(p *PRD, status Status) {
	p.Metadata.Status = status
}

// ParseMoSCoW converts a string to MoSCoW priority type.
func ParseMoSCoW(s string) MoSCoW {
	switch s {
	case "must":
		return MoSCoWMust
	case "should":
		return MoSCoWShould
	case "could":
		return MoSCoWCould
	case "wont", "won't":
		return MoSCoWWont
	default:
		return MoSCoWShould
	}
}

// ParseRiskImpact converts a string to RiskImpact type.
func ParseRiskImpact(s string) RiskImpact {
	switch s {
	case "low":
		return RiskImpactLow
	case "medium":
		return RiskImpactMedium
	case "high":
		return RiskImpactHigh
	case "critical":
		return RiskImpactCritical
	default:
		return RiskImpactMedium
	}
}

// ParseRiskProbability converts a string to RiskProbability type.
func ParseRiskProbability(s string) RiskProbability {
	switch s {
	case "low":
		return RiskProbabilityLow
	case "medium":
		return RiskProbabilityMedium
	case "high":
		return RiskProbabilityHigh
	default:
		return RiskProbabilityMedium
	}
}

// ParseNFRCategory converts a string to NFRCategory type.
func ParseNFRCategory(s string) NFRCategory {
	switch s {
	case "performance":
		return NFRPerformance
	case "security":
		return NFRSecurity
	case "reliability":
		return NFRReliability
	case "scalability":
		return NFRScalability
	case "usability":
		return NFRUsability
	case "compliance":
		return NFRCompliance
	case "maintainability":
		return NFRMaintainability
	case "availability":
		return NFRAvailability
	case "observability":
		return NFRObservability
	default:
		return NFRPerformance
	}
}

// ParseStatus converts a string to Status type.
func ParseStatus(s string) (Status, bool) {
	switch s {
	case "draft":
		return StatusDraft, true
	case "in_review", "review":
		return StatusInReview, true
	case "approved":
		return StatusApproved, true
	case "deprecated":
		return StatusDeprecated, true
	default:
		return "", false
	}
}
