package linter

import (
	"sync"
)

type RuleViolation struct {
	RuleName string
	Ref      string
	Failure  string
}

type Report struct {
	violations []RuleViolation
	lock       sync.Mutex
}

func NewReport() *Report {
	return &Report{
		violations: []RuleViolation{},
		lock:       sync.Mutex{},
	}
}

func (report *Report) AddViolation(violation RuleViolation) {
	report.lock.Lock()
	defer report.lock.Unlock()
	report.violations = append(report.violations, violation)
}

func (report *Report) GetViolations() []RuleViolation {
	return report.violations
}
