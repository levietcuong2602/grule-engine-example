package rule_engine

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

var knowledgeLibary = *ast.NewKnowledgeLibrary()

// Rule input object
type RuleInput interface {
	DataKey() string
}

// Rule output object
type RuleOutput interface {
	DataKey() string
}

// configs associated with the rule engine
type RuleEngineConfig interface {
	RuleName() string
	RuleInput() RuleInput
	RuleOutput() RuleOutput
}

type RuleEngineSvc struct {
}

func NewRuleEngineSvc() *RuleEngineSvc {
	// you should add your cloud provider here instead of keeping rule file in your code.
	buildRuleEngine()
	return &RuleEngineSvc{}
}

func buildRuleEngine() {
	ruleBuilder := builder.NewRuleBuilder(&knowledgeLibary)

	// Read rule from file and build the rule engine
	ruleFile := pkg.NewFileResource("rules.grl")
	err := ruleBuilder.BuildRuleFromResource("Rules", "0.0.1", ruleFile)
	if err != nil {
		panic(err)
	}
}

func (svc *RuleEngineSvc) ExecuteRuleEngine(config RuleEngineConfig) error {
	// get KnowlegeBase instance to execute particular rulek
	knowledgeBase, _ := knowledgeLibary.NewKnowledgeBaseInstance("Rules", "0.0.1")

	dataCtx := ast.NewDataContext()
	//  add input data context
	err := dataCtx.Add(config.RuleInput().DataKey(), config.RuleInput())
	if err != nil {
		return err
	}

	// add output data context
	err = dataCtx.Add(config.RuleOutput().DataKey(), config.RuleOutput())
	if err != nil {
		return err
	}

	// create rule engine and execute on provided data and knowledge base
	ruleEngine := engine.NewGruleEngine()
	err = ruleEngine.Execute(dataCtx, knowledgeBase)
	if err != nil {
		return err
	}

	return nil
}
