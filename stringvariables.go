package main

import "fmt"

// StringVariable indicates a CSP variable of string type
type StringVariable struct {
	Name   VariableName
	Value  string
	Domain StringDomain
	Empty  bool
}

// NewStringVariable constructor for StringVariable type
func NewStringVariable(name VariableName, domain StringDomain) StringVariable {
	return StringVariable{name, "", domain, true}
}

// SetValue setter for StringVariable value field
func (variable *StringVariable) SetValue(value string) {
	variable.Value = value
	variable.Empty = false
}

// Unset the variable
func (variable *StringVariable) Unset() {
	variable.Empty = true
	var s string
	variable.Value = s
}

// StringVariables collection type for string type variables
type StringVariables []StringVariable

// SetValue setter for StringVariables collection
func (variables *StringVariables) SetValue(name VariableName, value string) {
	foundIndex := -1

	for index, variable := range *variables {
		if variable.Name == name {
			foundIndex = index
		}
	}
	if !(foundIndex >= 0) {
		panic(fmt.Sprintf("StringVariable not found by name %v in variables %v", name, variables))
	} else {
		(*variables)[foundIndex].Value = value
		(*variables)[foundIndex].Empty = false

	}
}

// Find find a StringVariable by name in a StringVariables collection
func (variables *StringVariables) Find(name VariableName) StringVariable {
	for _, variable := range *variables {
		if variable.Name == name {
			return variable
		}
	}
	panic(fmt.Sprintf("StringVariable not found by name %v in variables %v", name, variables))
}

// Contains slice contains method for StringVariables
func (variables *StringVariables) Contains(name VariableName) bool {
	for _, variable := range *variables {
		if variable.Name == name {
			return true
		}
	}
	return false
}

// Unassigned return all unassigned variables
func (variables *StringVariables) Unassigned() StringVariables {
	unassigned := make(StringVariables, 0)
	for _, variable := range *variables {
		if variable.Empty {
			unassigned = append(unassigned, variable)
		}
	}
	return unassigned
}

// Complete indicates if all variables have been assigned to
func (variables *StringVariables) Complete() bool {
	return len(variables.Unassigned()) == 0
}
