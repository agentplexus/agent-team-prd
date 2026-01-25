package prd

import (
	"testing"
)

func TestSetProblemStatement(t *testing.T) {
	p := New("PRD-2026-001", "Test PRD", Person{Name: "Owner"})

	SetProblemStatement(p, "Primary problem", "High impact", 0.9)

	if p.Problem == nil {
		t.Fatal("expected Problem to be set")
	}
	if p.Problem.Statement != "Primary problem" {
		t.Errorf("expected statement 'Primary problem', got %s", p.Problem.Statement)
	}
	if p.Problem.UserImpact != "High impact" {
		t.Errorf("expected impact 'High impact', got %s", p.Problem.UserImpact)
	}
	if p.Problem.Confidence != 0.9 {
		t.Errorf("expected confidence 0.9, got %f", p.Problem.Confidence)
	}
	if p.ExecutiveSummary.ProblemStatement != "Primary problem" {
		t.Errorf("expected executive summary problem statement to be set")
	}
}

func TestAddPersona(t *testing.T) {
	p := New("PRD-2026-001", "Test PRD", Person{Name: "Owner"})

	// First persona becomes primary
	id1 := AddPersona(p, "Developer Dan", "Backend Developer", []string{"Slow builds", "Complex configs"})
	if len(p.Personas) != 1 {
		t.Errorf("expected 1 persona, got %d", len(p.Personas))
	}
	if p.Personas[0].ID != id1 {
		t.Errorf("expected persona ID %s, got %s", id1, p.Personas[0].ID)
	}
	if p.Personas[0].Name != "Developer Dan" {
		t.Errorf("expected name 'Developer Dan', got %s", p.Personas[0].Name)
	}
	if len(p.Personas[0].PainPoints) != 2 {
		t.Errorf("expected 2 pain points, got %d", len(p.Personas[0].PainPoints))
	}
	if !p.Personas[0].IsPrimary {
		t.Error("expected first persona to be primary")
	}

	// Second persona doesn't become primary
	id2 := AddPersona(p, "Manager Mike", "Engineering Manager", nil)
	if len(p.Personas) != 2 {
		t.Errorf("expected 2 personas, got %d", len(p.Personas))
	}
	if p.Personas[1].IsPrimary {
		t.Error("expected second persona to not be primary")
	}
	_ = id2
}

func TestAddProductGoal(t *testing.T) {
	p := New("PRD-2026-001", "Test PRD", Person{Name: "Owner"})

	id := AddProductGoal(p, "Reduce latency by 50%", "Improve user experience")
	if len(p.Objectives.ProductGoals) != 1 {
		t.Errorf("expected 1 product goal, got %d", len(p.Objectives.ProductGoals))
	}
	if p.Objectives.ProductGoals[0].ID != id {
		t.Errorf("expected goal ID %s, got %s", id, p.Objectives.ProductGoals[0].ID)
	}
	if p.Objectives.ProductGoals[0].Description != "Reduce latency by 50%" {
		t.Errorf("expected description 'Reduce latency by 50%%', got %s", p.Objectives.ProductGoals[0].Description)
	}
}

func TestAddOutOfScope(t *testing.T) {
	p := New("PRD-2026-001", "Test PRD", Person{Name: "Owner"})

	AddOutOfScope(p, "Mobile support")
	AddOutOfScope(p, "Offline mode")

	if len(p.OutOfScope) != 2 {
		t.Errorf("expected 2 out-of-scope items, got %d", len(p.OutOfScope))
	}
	if p.OutOfScope[0] != "Mobile support" {
		t.Errorf("expected first item 'Mobile support', got %s", p.OutOfScope[0])
	}
}

func TestAddSolution(t *testing.T) {
	p := New("PRD-2026-001", "Test PRD", Person{Name: "Owner"})

	id := AddSolution(p, "OAuth 2.0", "Standard authentication", []string{"Third-party dependency"})
	if p.Solution == nil {
		t.Fatal("expected Solution to be initialized")
	}
	if len(p.Solution.SolutionOptions) != 1 {
		t.Errorf("expected 1 solution, got %d", len(p.Solution.SolutionOptions))
	}
	if p.Solution.SolutionOptions[0].ID != id {
		t.Errorf("expected solution ID %s, got %s", id, p.Solution.SolutionOptions[0].ID)
	}
	if p.Solution.SolutionOptions[0].Name != "OAuth 2.0" {
		t.Errorf("expected name 'OAuth 2.0', got %s", p.Solution.SolutionOptions[0].Name)
	}
}

func TestSelectSolution(t *testing.T) {
	p := New("PRD-2026-001", "Test PRD", Person{Name: "Owner"})

	// Can't select without solutions
	if SelectSolution(p, "SOL-1", "test") {
		t.Error("expected SelectSolution to fail with no solutions")
	}

	id := AddSolution(p, "OAuth 2.0", "Standard auth", nil)

	// Select non-existent ID
	if SelectSolution(p, "SOL-99", "test") {
		t.Error("expected SelectSolution to fail with non-existent ID")
	}

	// Select valid ID
	if !SelectSolution(p, id, "Best option for security") {
		t.Error("expected SelectSolution to succeed")
	}
	if p.Solution.SelectedSolutionID != id {
		t.Errorf("expected selected ID %s, got %s", id, p.Solution.SelectedSolutionID)
	}
	if p.Solution.SolutionRationale != "Best option for security" {
		t.Errorf("expected rationale 'Best option for security', got %s", p.Solution.SolutionRationale)
	}
}

func TestAddFunctionalRequirement(t *testing.T) {
	p := New("PRD-2026-001", "Test PRD", Person{Name: "Owner"})

	id := AddFunctionalRequirement(p, "OAuth Login", "Support OAuth login", MoSCoWMust)
	if len(p.Requirements.Functional) != 1 {
		t.Errorf("expected 1 requirement, got %d", len(p.Requirements.Functional))
	}
	if p.Requirements.Functional[0].ID != id {
		t.Errorf("expected requirement ID %s, got %s", id, p.Requirements.Functional[0].ID)
	}
	if p.Requirements.Functional[0].Priority != MoSCoWMust {
		t.Errorf("expected priority must, got %s", p.Requirements.Functional[0].Priority)
	}
}

func TestAddNonFunctionalRequirement(t *testing.T) {
	p := New("PRD-2026-001", "Test PRD", Person{Name: "Owner"})

	id := AddNonFunctionalRequirement(p, NFRSecurity, "Data Encryption", "All data encrypted at rest", "AES-256", MoSCoWMust)
	if len(p.Requirements.NonFunctional) != 1 {
		t.Errorf("expected 1 NFR, got %d", len(p.Requirements.NonFunctional))
	}
	if p.Requirements.NonFunctional[0].ID != id {
		t.Errorf("expected NFR ID %s, got %s", id, p.Requirements.NonFunctional[0].ID)
	}
	if p.Requirements.NonFunctional[0].Category != NFRSecurity {
		t.Errorf("expected category security, got %s", p.Requirements.NonFunctional[0].Category)
	}
}

func TestAddSuccessMetric(t *testing.T) {
	p := New("PRD-2026-001", "Test PRD", Person{Name: "Owner"})

	id := AddSuccessMetric(p, "Login Success Rate", "Successful logins / Total attempts", "99.5%")
	if len(p.Objectives.SuccessMetrics) != 1 {
		t.Errorf("expected 1 metric, got %d", len(p.Objectives.SuccessMetrics))
	}
	if p.Objectives.SuccessMetrics[0].ID != id {
		t.Errorf("expected metric ID %s, got %s", id, p.Objectives.SuccessMetrics[0].ID)
	}
	if p.Objectives.SuccessMetrics[0].Target != "99.5%" {
		t.Errorf("expected target '99.5%%', got %s", p.Objectives.SuccessMetrics[0].Target)
	}
}

func TestAddRisk(t *testing.T) {
	p := New("PRD-2026-001", "Test PRD", Person{Name: "Owner"})

	id := AddRisk(p, "OAuth provider outage", RiskProbabilityMedium, RiskImpactHigh, "Implement fallback auth")
	if len(p.Risks) != 1 {
		t.Errorf("expected 1 risk, got %d", len(p.Risks))
	}
	if p.Risks[0].ID != id {
		t.Errorf("expected risk ID %s, got %s", id, p.Risks[0].ID)
	}
	if p.Risks[0].Impact != RiskImpactHigh {
		t.Errorf("expected impact high, got %s", p.Risks[0].Impact)
	}
}

func TestAddDecision(t *testing.T) {
	p := New("PRD-2026-001", "Test PRD", Person{Name: "Owner"})

	id := AddDecision(p, "Use JWT for sessions", "Stateless and scalable", "Tech Lead")
	if p.Decisions == nil {
		t.Fatal("expected Decisions to be initialized")
	}
	if len(p.Decisions.Records) != 1 {
		t.Errorf("expected 1 decision, got %d", len(p.Decisions.Records))
	}
	if p.Decisions.Records[0].ID != id {
		t.Errorf("expected decision ID %s, got %s", id, p.Decisions.Records[0].ID)
	}
	if p.Decisions.Records[0].MadeBy != "Tech Lead" {
		t.Errorf("expected made by 'Tech Lead', got %s", p.Decisions.Records[0].MadeBy)
	}
}

func TestUpdateStatus(t *testing.T) {
	p := New("PRD-2026-001", "Test PRD", Person{Name: "Owner"})

	if p.Metadata.Status != StatusDraft {
		t.Errorf("expected initial status draft, got %s", p.Metadata.Status)
	}

	UpdateStatus(p, StatusInReview)
	if p.Metadata.Status != StatusInReview {
		t.Errorf("expected status in_review, got %s", p.Metadata.Status)
	}

	UpdateStatus(p, StatusApproved)
	if p.Metadata.Status != StatusApproved {
		t.Errorf("expected status approved, got %s", p.Metadata.Status)
	}
}

func TestParseMoSCoW(t *testing.T) {
	tests := []struct {
		input    string
		expected MoSCoW
	}{
		{"must", MoSCoWMust},
		{"should", MoSCoWShould},
		{"could", MoSCoWCould},
		{"wont", MoSCoWWont},
		{"invalid", MoSCoWShould}, // default
		{"", MoSCoWShould},        // default
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ParseMoSCoW(tt.input); got != tt.expected {
				t.Errorf("ParseMoSCoW(%s) = %s, want %s", tt.input, got, tt.expected)
			}
		})
	}
}

func TestParseRiskImpact(t *testing.T) {
	tests := []struct {
		input    string
		expected RiskImpact
	}{
		{"low", RiskImpactLow},
		{"medium", RiskImpactMedium},
		{"high", RiskImpactHigh},
		{"critical", RiskImpactCritical},
		{"invalid", RiskImpactMedium}, // default
		{"", RiskImpactMedium},        // default
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ParseRiskImpact(tt.input); got != tt.expected {
				t.Errorf("ParseRiskImpact(%s) = %s, want %s", tt.input, got, tt.expected)
			}
		})
	}
}

func TestParseNFRCategory(t *testing.T) {
	tests := []struct {
		input    string
		expected NFRCategory
	}{
		{"performance", NFRPerformance},
		{"security", NFRSecurity},
		{"reliability", NFRReliability},
		{"scalability", NFRScalability},
		{"usability", NFRUsability},
		{"compliance", NFRCompliance},
		{"invalid", NFRPerformance}, // default
		{"", NFRPerformance},        // default
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ParseNFRCategory(tt.input); got != tt.expected {
				t.Errorf("ParseNFRCategory(%s) = %s, want %s", tt.input, got, tt.expected)
			}
		})
	}
}

func TestParseStatus(t *testing.T) {
	tests := []struct {
		input    string
		expected Status
		ok       bool
	}{
		{"draft", StatusDraft, true},
		{"in_review", StatusInReview, true},
		{"review", StatusInReview, true},
		{"approved", StatusApproved, true},
		{"deprecated", StatusDeprecated, true},
		{"invalid", "", false},
		{"", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, ok := ParseStatus(tt.input)
			if ok != tt.ok {
				t.Errorf("ParseStatus(%s) ok = %v, want %v", tt.input, ok, tt.ok)
			}
			if ok && got != tt.expected {
				t.Errorf("ParseStatus(%s) = %s, want %s", tt.input, got, tt.expected)
			}
		})
	}
}

func TestNextID(t *testing.T) {
	p := New("PRD-2026-001", "Test PRD", Person{Name: "Owner"})

	// First persona should get PER-1
	id1 := NextID(p, "PER")
	if id1 != "PER-1" {
		t.Errorf("expected PER-1, got %s", id1)
	}

	// After adding, next should be PER-2
	AddPersona(p, "Test", "Role", nil)
	id2 := NextID(p, "PER")
	if id2 != "PER-2" {
		t.Errorf("expected PER-2, got %s", id2)
	}
}
