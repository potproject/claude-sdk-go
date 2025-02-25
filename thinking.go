package v1

func UseThinking(budgetTokens int) *RequestBodyMessagesThinking {
	t := RequestBodyMessagesThinking{
		Type:         "enabled",
		BudgetTokens: budgetTokens,
	}
	return &t
}
