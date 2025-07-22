package pkg

// ToolInfo holds the name and description of a tool
type ToolInfo struct {
	Name        string
	Description string
}

// ToolList is the list of all tools and their descriptions
var ToolList = []ToolInfo{
	{
		Name:        "fetch_net_worth",
		Description: "Calculate comprehensive net worth using ONLY actual data from accounts users connected on Fi Money including: Bank account balances, Mutual fund investment holdings, Indian Stocks investment holdings, Total US Stocks investment (If investing through Fi Money app), EPF account balances, Credit card debt and loan balances (if credit report connected), Any other assets/liabilities linked to Fi Money platform.",
	},
	{
		Name:        "fetch_credit_report",
		Description: "Retrieve comprehensive credit report including scores, active loans, credit card utilization, payment history, date of birth and recent inquiries from connected credit bureaus.",
	},
	{
		Name:        "fetch_epf_details",
		Description: "Retrieve detailed EPF (Employee Provident Fund) account information including: Account balance and contributions, Employer and employee contribution history, Interest earned and credited amounts.",
	},
	{
		Name:        "fetch_mf_transactions",
		Description: "Retrieve detailed transaction history from accounts connected to Fi Money platform including: Mutual fund transactions.",
	},
	{
		Name:        "fetch_bank_transactions",
		Description: "Retrieve detailed bank transactions for each bank account connected to Fi money platform.",
	},
	{
		Name:        "fetch_stock_transactions",
		Description: "Retrieve detailed indian stock transactions for all connected indian stock accounts to Fi money platform.",
	},
}
