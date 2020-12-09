package dan

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var day8Input = `acc +28
jmp +481
nop +87
acc -10
acc +21
jmp +548
jmp +250
acc -9
acc +19
acc -3
nop +85
jmp +84
acc +43
acc +34
acc -5
jmp +602
nop +416
acc +19
acc +46
acc +33
jmp +133
acc -3
acc +5
jmp +143
nop -18
acc +50
acc +0
acc -6
jmp +243
acc +11
acc -18
jmp +501
acc +45
jmp +449
acc +0
acc -10
acc +20
nop +516
jmp +544
jmp +551
nop +59
nop +247
acc -1
jmp +589
acc +41
jmp +457
jmp +102
acc +22
jmp +195
acc +30
jmp -38
acc +21
nop -3
jmp +466
acc +47
acc +4
jmp +226
jmp +404
acc -19
acc +13
acc +37
acc -6
jmp +240
nop +272
acc +41
acc -12
acc +41
jmp +159
acc +26
jmp +363
acc +33
jmp +1
jmp +76
acc -15
acc -6
jmp +1
jmp +259
nop +149
jmp +1
jmp +20
acc +31
jmp +131
acc +47
acc +18
jmp +355
acc +9
acc +42
jmp +28
jmp +138
acc +29
acc +7
nop +270
acc +6
jmp +445
jmp +282
jmp +266
acc +30
nop +283
jmp +134
acc -16
acc +50
acc +30
jmp +344
acc +25
acc +42
acc -8
jmp +274
acc +27
acc +33
acc -18
acc +21
jmp +448
acc -7
nop +342
jmp -65
acc +43
nop +407
nop +122
jmp +170
acc +34
jmp +1
acc +1
acc +11
jmp +436
acc -1
jmp +346
jmp +250
acc +17
acc +21
nop +482
jmp -90
acc +45
jmp +244
acc +29
acc +23
nop +241
acc +34
jmp +168
acc +28
acc -6
jmp +299
nop +441
jmp +304
nop +41
acc +49
acc -18
acc +7
jmp +244
acc +45
jmp -95
acc -9
acc -18
jmp +14
acc +35
acc +47
jmp +297
acc +34
acc +33
jmp +1
jmp +210
nop +128
jmp +305
acc +29
acc +17
acc +47
jmp -24
nop -86
jmp +97
acc -7
acc +14
nop +118
jmp +411
acc +42
acc -8
acc -15
jmp +232
jmp +1
acc -6
jmp +20
acc -15
acc +3
acc +50
jmp +1
jmp +430
acc -18
acc +35
acc +46
acc +1
jmp +51
acc +13
nop +209
jmp +114
nop +319
acc +35
jmp +194
acc +6
acc +42
jmp -44
jmp +1
nop +208
acc -13
jmp +412
acc -12
nop +327
nop -27
acc +34
jmp +324
acc +21
acc +16
jmp +273
acc -14
jmp +363
acc +4
jmp +140
acc +6
acc +28
jmp +110
jmp +432
nop -20
acc +4
acc +1
jmp +392
jmp -164
acc +23
acc +10
jmp +57
jmp +385
acc +5
acc +28
acc +25
acc +0
jmp +251
acc -17
jmp +38
nop -2
acc +35
acc -3
acc +40
jmp +123
acc +19
jmp +271
acc -7
jmp +40
acc -7
acc +0
acc +46
jmp -161
acc +33
acc +1
jmp +341
acc -12
acc +45
nop -145
jmp +116
acc +49
jmp +20
jmp -34
acc +41
acc +35
jmp -132
acc +45
jmp +119
acc +25
jmp +12
acc +8
jmp -15
acc +14
acc +43
acc +17
acc +35
jmp -98
jmp -182
nop +260
nop +338
jmp +31
acc +23
nop -64
acc +36
acc +47
jmp +318
acc +10
jmp +180
acc +37
acc -13
acc -6
nop -29
jmp -179
jmp -5
jmp +1
acc +43
jmp -34
nop -218
acc -10
jmp -5
nop +67
acc +9
acc +47
jmp -276
acc +28
nop +260
acc +38
jmp -280
jmp +303
acc +42
jmp -209
acc -2
acc +13
acc -1
acc +43
jmp +244
acc -15
acc +47
jmp +149
jmp +1
nop +233
acc +14
jmp -287
jmp -110
acc -12
jmp -159
nop -94
jmp -188
nop +20
jmp +53
nop -64
jmp -296
acc +20
jmp +134
acc +8
jmp +47
acc -15
jmp -37
acc -14
acc -15
jmp +163
nop -22
acc -3
jmp -210
acc +37
acc -11
nop +179
jmp -193
acc +27
jmp +81
acc +46
acc +21
jmp +93
acc -9
nop -131
nop -177
jmp -173
acc -19
jmp +55
nop +16
acc +45
acc +10
jmp +83
acc +15
acc +17
acc +28
jmp -121
nop -19
acc +47
jmp +184
acc +45
acc +15
acc +6
acc -3
jmp +164
acc +1
acc +30
nop +226
acc +20
jmp -339
nop +223
jmp +266
acc +0
jmp -253
acc -8
jmp -26
nop -326
nop -106
jmp +73
acc +2
acc -7
acc +4
acc -1
jmp -61
jmp +186
acc -14
jmp +221
nop +111
acc -7
acc +46
jmp -345
acc +2
nop -212
acc +23
jmp -159
nop -258
acc +3
acc +40
jmp -142
acc +49
acc +27
acc +45
jmp -214
acc +1
acc +5
acc +23
acc -7
jmp +119
nop +201
nop -215
nop -197
jmp +115
acc -9
acc +38
jmp -211
acc +40
acc -1
acc +16
acc +35
jmp -70
acc -18
acc -10
jmp +93
acc +14
nop -216
acc -12
nop -223
jmp -342
acc -19
acc +43
acc -6
acc +11
jmp -106
acc +2
jmp +73
acc +48
acc +20
acc +18
acc +44
jmp +162
acc -7
jmp -202
acc -2
acc +34
acc -17
jmp -443
acc +40
jmp -129
jmp -181
acc +30
jmp -400
acc +42
nop -185
acc +34
acc -5
jmp -275
jmp -406
acc -16
jmp -270
acc -4
acc +40
jmp -299
acc -7
jmp -76
acc +10
acc +11
jmp -4
acc +0
acc +33
acc -8
acc +8
jmp +166
nop +62
jmp +1
nop +13
jmp +23
jmp -167
jmp -469
acc +32
nop -403
acc +7
acc -3
jmp -345
acc -9
acc -2
acc +48
jmp +57
acc +9
acc +3
nop -32
acc +25
jmp -177
jmp -369
nop -167
acc -9
jmp -453
acc -6
jmp -360
acc +11
nop -457
acc +50
nop +121
jmp -231
nop -83
acc -4
nop -253
jmp +28
acc +29
acc +0
jmp -11
nop -373
acc -19
acc +26
acc +0
jmp -7
jmp +71
acc +18
acc +50
jmp -234
nop -399
acc -7
acc +32
jmp -294
nop -481
acc +8
jmp -113
jmp -183
jmp -333
acc -7
acc +10
acc +31
jmp -501
nop +75
jmp -219
acc +2
acc +43
acc +19
acc -2
jmp -443
acc +16
acc +6
acc +40
jmp -461
jmp +79
acc +32
acc +20
acc -7
nop -71
jmp -498
jmp -231
acc +37
nop -526
acc +38
acc +41
jmp -509
acc +26
acc +45
acc +33
acc +37
jmp -410
nop -289
acc +46
acc +39
jmp -44
jmp -169
jmp +1
acc +33
jmp +1
acc +39
jmp -446
acc -9
nop -187
acc +41
jmp -444
acc +20
acc +9
nop -500
jmp -43
acc +23
acc +28
acc +35
jmp -94
acc +47
nop -514
acc +38
jmp -577
nop -31
nop -348
acc -18
jmp -202
nop -162
nop -373
jmp +45
acc +45
acc -16
acc -15
jmp -532
acc +29
acc -14
acc -11
jmp -509
jmp -279
acc -18
acc +42
acc +12
jmp +1
jmp -450
jmp -88
jmp -505
acc +0
jmp -397
nop -438
jmp -86
acc +14
acc +27
nop -66
nop -278
jmp -548
acc +19
nop -305
nop -440
jmp -441
acc +35
acc +36
acc +45
jmp -505
acc -8
jmp -313
nop -517
acc +38
acc +12
nop -511
jmp -354
acc +29
acc -7
acc +29
jmp -201
acc -2
acc -3
jmp -75
acc +24
acc -2
jmp +1
acc +19
jmp -58
nop -432
acc +1
acc +33
acc +1
jmp +1`

type ComputeState struct {
	NextLine    int
	Accumulator int
}

type Instruction interface {
	Source() string
	ExecutionCount() int
	ResetExecutionCount()
	Execute(cs *ComputeState)
}

type instruction struct {
	source         string
	executionCount int
}

func (i *instruction) Source() string {
	return i.source
}

func (i *instruction) ExecutionCount() int {
	return i.executionCount
}

func (i *instruction) ResetExecutionCount() {
	i.executionCount = 0
}

func (i *instruction) Execute(cs *ComputeState) {
	i.executionCount++
	cs.NextLine++
}

type noopInstruction struct {
	instruction
	value int
}

func (i *noopInstruction) Execute(cs *ComputeState) {
	i.instruction.Execute(cs)
}

func newNoopInstruction(value int) *noopInstruction {
	return &noopInstruction{
		value: value,
	}
}

type jumpInstruction struct {
	instruction
	offset int
}

func (i *jumpInstruction) Execute(cs *ComputeState) {
	i.executionCount++
	cs.NextLine += i.offset
}

func newJumpInstruction(offset int) *jumpInstruction {
	return &jumpInstruction{
		offset: offset,
	}
}

type accInstruction struct {
	instruction
	increment int
}

func (i *accInstruction) Execute(cs *ComputeState) {
	i.instruction.Execute(cs)
	cs.Accumulator += i.increment
}

func newAccInstruction(increment int) *accInstruction {
	return &accInstruction{
		increment: increment,
	}
}

func parseInstruction(s string) (Instruction, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid instruction `%s` - invalid number of parts", s)
	}
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("Invalid instruction `%s` - could not parse value", s)
	}
	switch parts[0] {
	case "nop":
		return newNoopInstruction(value), nil
	case "jmp":
		return newJumpInstruction(value), nil
	case "acc":
		return newAccInstruction(value), nil
	}
	return nil, fmt.Errorf("Invalid instruction `%s` - unknown type", s)
}

func TestParseInstruction(t *testing.T) {
	type td struct {
		s string
		i Instruction
	}
	tests := []td{
		{"nop -481", newNoopInstruction(-481)},
		{"acc +8", newAccInstruction(8)},
		{"jmp -113", newJumpInstruction(-113)},
	}
	for n, test := range tests {
		i, err := parseInstruction(test.s)
		assert.NoErrorf(t, err, "Error parsing instruction %d - %#v", n, test)
		assert.EqualValues(t, test.i, i, "Incorrect instruction for test %d - %#v", n, test)
	}
}

type Instructions []Instruction

func (instructions Instructions) ResetExecutionCount() {
	for _, i := range instructions {
		i.ResetExecutionCount()
	}
}

func parseInstructions(s string) (Instructions, error) {
	var instructions Instructions
	for _, line := range strings.Split(s, "\n") {
		i, err := parseInstruction(line)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, i)
	}
	return instructions, nil
}

func TestParseInstructions(t *testing.T) {
	expected := []Instruction{newNoopInstruction(-481), newAccInstruction(8), newJumpInstruction(-113)}
	value, err := parseInstructions("nop -481\nacc +8\njmp -113")
	require.NoError(t, err, "Error parsing instructions")
	assert.EqualValues(t, expected, value, "Instructions parsed incorrectly")
}

type Stopper func(instructions Instructions, cs *ComputeState) bool

func reExecuteStopper(instructions Instructions, cs *ComputeState) bool {
	return instructions[cs.NextLine].ExecutionCount() > 0
}

func runInstructions(instructions Instructions, stopper Stopper) (ComputeState, bool) {
	var cs ComputeState
	for cs.NextLine >= 0 && cs.NextLine < len(instructions) {
		if stopper != nil && stopper(instructions, &cs) {
			return cs, true
		}
		i := instructions[cs.NextLine]
		i.Execute(&cs)
	}
	return cs, false
}

func TestRunInstructions(t *testing.T) {
	expected := ComputeState{5, 16}
	instructions, err := parseInstructions("nop -481\nacc +8\njmp +1\nnop +5\nacc +8")
	require.NoError(t, err, "Error parsing instructions")
	cs, stopped := runInstructions(instructions, nil)
	assert.EqualValues(t, expected, cs, "Incorrect compute state")
	assert.False(t, stopped, "Execution was stopped")
}

func TestRunInstructionsWithStopper(t *testing.T) {
	expected := ComputeState{1, 16}
	instructions, err := parseInstructions("nop -481\nacc +8\njmp +1\nnop +5\nacc +8\njmp -4")
	require.NoError(t, err, "Error parsing instructions")
	cs, stopped := runInstructions(instructions, reExecuteStopper)
	assert.EqualValues(t, expected, cs, "Incorrect compute state")
	assert.True(t, stopped, "Execution was not stopped")
}

func TestDay8a(t *testing.T) {
	instructions, err := parseInstructions(day8Input)
	require.NoError(t, err, "Error parsing instructions")
	cs, stopped := runInstructions(instructions, reExecuteStopper)
	assert.True(t, stopped, "Execution was not stopped, loop not detected")
	t.Logf("Compute state when loop detected: %#v", cs)
}

func TestDay8b(t *testing.T) {
	instructions, err := parseInstructions(day8Input)
	require.NoError(t, err, "Error parsing instructions")
	for n, i := range instructions {
		var newInstruction Instruction
		switch v := i.(type) {
		case *noopInstruction:
			newInstruction = newJumpInstruction(v.value)
		case *jumpInstruction:
			newInstruction = newNoopInstruction(v.offset)
		}
		if newInstruction == nil {
			continue
		}
		instructions[n] = newInstruction
		cs, stopped := runInstructions(instructions, reExecuteStopper)
		if !stopped {
			t.Logf("Swapped instruction %d and no loop found\n\tOriginal: %#v\n\tNew: %#v\n\tState: %#v", n, i, newInstruction, cs)
			return
		}
		instructions[n] = i
		instructions.ResetExecutionCount()
	}
	t.Error("Did not find an instruction to swap to prevent a loop")
}