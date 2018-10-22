// Code generated by "enumer -type=State"; DO NOT EDIT.

package pruner

import (
	"fmt"
)

const (
	_StateName_0 = "PlanPlanWait"
	_StateName_1 = "Exec"
	_StateName_2 = "ExecWait"
	_StateName_3 = "ErrPerm"
	_StateName_4 = "Done"
)

var (
	_StateIndex_0 = [...]uint8{0, 4, 12}
	_StateIndex_1 = [...]uint8{0, 4}
	_StateIndex_2 = [...]uint8{0, 8}
	_StateIndex_3 = [...]uint8{0, 7}
	_StateIndex_4 = [...]uint8{0, 4}
)

func (i State) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _StateName_0[_StateIndex_0[i]:_StateIndex_0[i+1]]
	case i == 4:
		return _StateName_1
	case i == 8:
		return _StateName_2
	case i == 16:
		return _StateName_3
	case i == 32:
		return _StateName_4
	default:
		return fmt.Sprintf("State(%d)", i)
	}
}

var _StateValues = []State{1, 2, 4, 8, 16, 32}

var _StateNameToValueMap = map[string]State{
	_StateName_0[0:4]:  1,
	_StateName_0[4:12]: 2,
	_StateName_1[0:4]:  4,
	_StateName_2[0:8]:  8,
	_StateName_3[0:7]:  16,
	_StateName_4[0:4]:  32,
}

// StateString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func StateString(s string) (State, error) {
	if val, ok := _StateNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to State values", s)
}

// StateValues returns all values of the enum
func StateValues() []State {
	return _StateValues
}

// IsAState returns "true" if the value is listed in the enum definition. "false" otherwise
func (i State) IsAState() bool {
	for _, v := range _StateValues {
		if i == v {
			return true
		}
	}
	return false
}
