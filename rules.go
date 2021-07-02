package main

import (
	"fmt"
)

type Rule struct {
	x 				string
	y 				string
	relationship 	bool
	model 			string
}

type Set struct {
	set []string
}

type RuleSet struct {
	rulesSets Set
	rulesList []Rule
}

func (s *Set) add_element(element string)  {
	for _,k := range s.set {
		if k == element {
			return
		}
	}
	s.set = append(s.set, element)
}

func (s *Set) contains(element string) bool {
	for _, k := range s.set {
		if k == element {
			return true
		}
	}
	return false
}

func (s *Set) relationship(x string, y string) bool {
	if s.contains(x) == true && s.contains(y) == true {
		return true
	}
	return false
}

func (r *RuleSet) NewRuleSet() Set {
	return r.rulesSets
}

func (r *RuleSet) AddDep(x string, y string)  {
	if r.rulesSets.relationship(x, y) == false {
		r.rulesList = append(r.rulesList, Rule{x: x, y: y, relationship: true, model: "d"})
	}

	if r.rulesSets.relationship(x, y) == true {
		rule := Rule{x: x, y: y, relationship: true, model: "d"}
		for _, k := range r.rulesList {
			if (k.x == x || k.y == y) || (k.x == y || k.y == x) {
				if k.model != "d" {
					rule.relationship = false
				}
			}
		}

		r.rulesList = append(r.rulesList, rule)
	}

	if r.rulesSets.contains(x) == false {
		r.rulesSets.add_element(x)
	}

	if r.rulesSets.contains(y) == false {
		r.rulesSets.add_element(y)
	}

}

func (r *RuleSet) AddConflict(x string, y string)  {
	if r.rulesSets.relationship(x, y) == false {
		r.rulesList = append(r.rulesList, Rule{x: x, y: y, relationship: true, model: "c"})
	}

	if r.rulesSets.relationship(x, y) == true {
		rule := Rule{x: x, y: y, relationship: true, model: "c"}
		for _, k := range r.rulesList {
			if k.x == x || k.y == y || k.x == y || k.y == x {
				if k.model != "c" {
					rule.relationship = false
				}
			}
		}

		r.rulesList = append(r.rulesList, rule)
	}

	if r.rulesSets.contains(x) == false {
		r.rulesSets.add_element(x)
	}

	if r.rulesSets.contains(y) == false {
		r.rulesSets.add_element(y)
	}

}

func (r *RuleSet) IsCoherent() bool {
	for _, key := range r.rulesList {
		if key.relationship == false {
			return false
		}
	}
	return true
}

func TestDependsAA()  {
	r := RuleSet{}
	r.AddDep("A", "B")
	fmt.Println("TestDependsAA - Coherent", r.IsCoherent())
}

func TestDependsAA_BA()  {
	r := RuleSet{}
	r.AddDep("A", "B")
	r.AddDep("B", "A")
	fmt.Println("TestDependsAA_BA - Coherent", r.IsCoherent())
}

func TestExclusiveAB()  {
	r := RuleSet{}
	r.AddDep("A", "B")
	r.AddConflict("A", "B")
	fmt.Println("TestDependsAA_BA - Coherent", r.IsCoherent())
}

func TestExclusiveAB_BC()  {
	r := RuleSet{}
	r.AddDep("A", "B")
	r.AddDep("B", "C")
	r.AddConflict("A", "C")
	fmt.Println("TestDependsAA_BA - Coherent", r.IsCoherent())
}

func TestDeepDeps()  {
	r := RuleSet{}
	r.AddDep("A", "B")
	r.AddDep("B", "C")
	r.AddDep("C", "D")
	r.AddDep("D", "E")
	r.AddConflict("A", "F")
	r.AddConflict("E", "F")
	fmt.Println("TestDependsAA_BA - Coherent", r.IsCoherent())
}

func main() {
	TestDependsAA()
	TestDependsAA_BA()
	TestExclusiveAB()
	TestExclusiveAB_BC()
	TestDeepDeps()
}






