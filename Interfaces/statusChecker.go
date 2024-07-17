package Interfaces

type StatusChecker interface {
	IsFinalStatus(status string) bool
}
