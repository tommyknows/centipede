// Copyright 2018 Gabriel Boorse

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package centipede

import (
	"math"
	"time"
)

// Domain domain object
type Domain []interface{}

// Contains slice contains method for Domain
func (domain *Domain) Contains(value interface{}) bool {
	for _, item := range *domain {
		if item == value {
			return true
		}
	}
	return false
}

// Remove given a value and return the updated domain
func (domain Domain) Remove(value interface{}) Domain {
	newDomain := make(Domain, 0)
	for _, item := range domain {
		if item != value {
			newDomain = append(newDomain, item)
		}
	}
	if len(newDomain) == len(domain) { // domain unchanged
		return domain
	}
	return newDomain
}

// IntRange returns a slice of integers in the desired range with a step of 1
func IntRange(start int, end int) Domain {
	return IntRangeStep(start, end, 1)
}

// IntRangeStep returns a slice of integers in the desired range with the given step
func IntRangeStep(start int, end int, step int) Domain {
	rangeLength := (end - start) / step
	mod := (end - start) % step
	if mod > 0 {
		rangeLength++
	}
	domain := make(Domain, rangeLength)
	for i := int(0); i < rangeLength; i++ {
		domain[i] = i*step + start
	}
	return domain
}

// TimeRange get the range of days from the start to the end time
func TimeRange(start time.Time, end time.Time) Domain {
	return TimeRangeStep(start, end, time.Duration(1)*time.Hour*24)
}

// TimeRangeStep get the range of time between start to end with step
// as a Duration (in nanoseconds).
func TimeRangeStep(start time.Time, end time.Time, step time.Duration) Domain {
	// get the number of units in this range of time
	rangeLength := end.Sub(start) / step
	mod := end.Sub(start) % step
	if mod > 0 {
		rangeLength++
	}
	// populate domain with units from beginning to end
	domain := make(Domain, rangeLength)
	for i := time.Duration(0); i < rangeLength; i++ {
		domain[i] = start.Add(i * step)
	}
	return domain
}

// FloatRange returns a slice of integers in the desired range with a step of 1
func FloatRange(start float64, end float64) Domain {
	return FloatRangeStep(start, end, 1.0)
}

// FloatRangeStep returns a slice of integers in the desired range with the given step
func FloatRangeStep(start float64, end float64, step float64) Domain {
	rangeLength := int(math.Ceil((end - start) / step))
	domain := make(Domain, rangeLength)
	for i := int(0); i < rangeLength; i++ {
		domain[i] = float64(i)*step + start
	}
	return domain
}

// Generator generates a Domain from another input domain
// and a function f(x). For example:
func Generator(inputDomain Domain, fx func(interface{}) interface{}) Domain {
	outputDomain := make(Domain, 0)
	for _, input := range inputDomain {
		outputDomain = append(outputDomain, fx(input))
	}
	return outputDomain
}
