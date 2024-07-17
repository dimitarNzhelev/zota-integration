package helpers

type TesterStatusChecker struct{}

func (t TesterStatusChecker) IsFinalStatus(status string) bool {
	return status == "APPROVED" || status == "DECLINED" || status == "ERROR" || status == "CREATED"
}
