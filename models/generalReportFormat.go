package models

type existingReportFormat struct {
	ticketKey string
	tests     map[string]string
}

type newReportFormat struct {
	tests map[string]string
}
