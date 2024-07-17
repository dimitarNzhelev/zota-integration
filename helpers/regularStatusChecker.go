package helpers

type RegularStatusChecker struct{}

func (r RegularStatusChecker) IsFinalStatus(status string) bool {
	return status == "APPROVED" || status == "DECLINED" || status == "ERROR"
}
