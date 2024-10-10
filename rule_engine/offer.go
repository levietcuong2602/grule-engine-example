package rule_engine

type UserOfferContext struct {
	UserOfferInput  *UserOfferInput
	UserOfferOutput *UserOfferOutput
}

func (uoc *UserOfferContext) RuleName() string {
	return "user_offers"
}

func (uoc *UserOfferContext) RuleInput() RuleInput {
	return uoc.UserOfferInput
}

func (uoc *UserOfferContext) RuleOutput() RuleOutput {
	return uoc.UserOfferOutput
}

// User Data attributes
type UserOfferInput struct {
	Name              string  `json:"name"`
	Username          string  `json:"username"`
	Email             string  `json:"email"`
	Age               int     `json:"age"`
	Gender            string  `json:"gender"`
	TotalOrders       int     `json:"total_orders"`
	AverageOrderValue float64 `json:"average_order_value"`
}

func (uoi *UserOfferInput) DataKey() string {
	return "UserOfferInput"
}

// Offer output attributes
type UserOfferOutput struct {
	IsOfferApplicable bool `json:"is_offer_applicable"`
}

func (uoo *UserOfferOutput) DataKey() string {
	return "UserOfferOutput"
}

func NewUserOfferContext() *UserOfferContext {
	return &UserOfferContext{
		UserOfferInput:  &UserOfferInput{},
		UserOfferOutput: &UserOfferOutput{},
	}
}
