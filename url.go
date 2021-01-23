package robinhood

// LOGIN URL
func login() string {
	return "https://api.robinhood.com/oauth2/token/"
}

//	PROFILE URL
func accountProfile() string {
	return "https://api.robinhood.com/accounts/"
}

func basicProfile() string {
	return "https://api.robinhood.com/user/basic_info/"
}

func investmentProfile() string {
	return "https://api.robinhood.com/user/investment_profile/"
}

func portfolioProfile() string {
	return "https://api.robinhood.com/portfolios/"
}

func securityProfile() string {
	return "https://api.robinhood.com/user/additional_info/"
}

func userProfile() string {
	return "https://api.robinhood.com/user/"
}

func historicalPortfolio(accountNumber string) string {
	return "https://api.robinhood.com/portfolios/historicals/" + accountNumber
}

// STOCK URL
func earnings() string {
	return "https://api.robinhood.com/marketdata/earnings/"
}

func events() string {
	return "https://api.robinhood.com/options/events/"
}

func fundamentals() string {
	return "https://api.robinhood.com/fundamentals/"
}

func historicals() string {
	return "https://api.robinhood.com/quotes/historicals/"
}

func instruments() string {
	return "https://api.robinhood.com/instruments/"
}

func news(symbol string) string {
	return "https://api.robinhood.com/midlands/news/" + symbol
}

func popularity(symbol string) string {
	return "https://api.robinhood.com/instruments/" + symbol + "/popularity/"
}

func quotes() string {
	return "https://api.robinhood.com/quotes/"
}

func ratings(symbol string) string {
	return "https://api.robinhood.com/midlands/ratings/" + symbol
}

func splits(symbol string) string {
	return "https://api.robinhood.com/instruments/" + symbol + "/splits/"
}
