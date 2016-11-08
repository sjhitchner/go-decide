
// generated by gocc; DO NOT EDIT.

package lexer



/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates] func(rune) int

var TransTab = TransitionTable{
	
		// S0
		func(r rune) int {
			switch {
			case r == 9 : // ['\t','\t']
				return 1
			case r == 10 : // ['\n','\n']
				return 1
			case r == 13 : // ['\r','\r']
				return 1
			case r == 32 : // [' ',' ']
				return 1
			case r == 33 : // ['!','!']
				return 2
			case r == 34 : // ['"','"']
				return 3
			case r == 39 : // [''',''']
				return 4
			case 48 <= r && r <= 57 : // ['0','9']
				return 5
			case r == 60 : // ['<','<']
				return 6
			case r == 61 : // ['=','=']
				return 7
			case r == 62 : // ['>','>']
				return 8
			case 65 <= r && r <= 90 : // ['A','Z']
				return 9
			case 97 <= r && r <= 98 : // ['a','b']
				return 9
			case r == 99 : // ['c','c']
				return 10
			case 100 <= r && r <= 101 : // ['d','e']
				return 9
			case r == 102 : // ['f','f']
				return 11
			case 103 <= r && r <= 104 : // ['g','h']
				return 9
			case r == 105 : // ['i','i']
				return 12
			case 106 <= r && r <= 108 : // ['j','l']
				return 9
			case r == 109 : // ['m','m']
				return 13
			case r == 110 : // ['n','n']
				return 14
			case r == 111 : // ['o','o']
				return 15
			case 112 <= r && r <= 115 : // ['p','s']
				return 9
			case r == 116 : // ['t','t']
				return 16
			case 117 <= r && r <= 122 : // ['u','z']
				return 9
			
			
			
			}
			return NoState
			
		},
	
		// S1
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S2
		func(r rune) int {
			switch {
			case r == 61 : // ['=','=']
				return 17
			
			
			
			}
			return NoState
			
		},
	
		// S3
		func(r rune) int {
			switch {
			case r == 34 : // ['"','"']
				return 18
			
			
			default:
				return 3
			}
			
		},
	
		// S4
		func(r rune) int {
			switch {
			case r == 39 : // [''',''']
				return 18
			
			
			default:
				return 4
			}
			
		},
	
		// S5
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 5
			
			
			
			}
			return NoState
			
		},
	
		// S6
		func(r rune) int {
			switch {
			case r == 61 : // ['=','=']
				return 20
			
			
			
			}
			return NoState
			
		},
	
		// S7
		func(r rune) int {
			switch {
			case r == 61 : // ['=','=']
				return 21
			
			
			
			}
			return NoState
			
		},
	
		// S8
		func(r rune) int {
			switch {
			case r == 61 : // ['=','=']
				return 22
			
			
			
			}
			return NoState
			
		},
	
		// S9
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 122 : // ['a','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S10
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 110 : // ['a','n']
				return 25
			case r == 111 : // ['o','o']
				return 26
			case 112 <= r && r <= 122 : // ['p','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S11
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case r == 97 : // ['a','a']
				return 27
			case 98 <= r && r <= 122 : // ['b','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S12
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 114 : // ['a','r']
				return 25
			case r == 115 : // ['s','s']
				return 28
			case 116 <= r && r <= 122 : // ['t','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S13
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case r == 97 : // ['a','a']
				return 29
			case 98 <= r && r <= 122 : // ['b','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S14
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 110 : // ['a','n']
				return 25
			case r == 111 : // ['o','o']
				return 30
			case 112 <= r && r <= 122 : // ['p','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S15
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 113 : // ['a','q']
				return 25
			case r == 114 : // ['r','r']
				return 31
			case 115 <= r && r <= 122 : // ['s','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S16
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 113 : // ['a','q']
				return 25
			case r == 114 : // ['r','r']
				return 32
			case 115 <= r && r <= 122 : // ['s','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S17
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S18
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S19
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 33
			
			
			
			}
			return NoState
			
		},
	
		// S20
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S21
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S22
		func(r rune) int {
			switch {
			
			
			
			}
			return NoState
			
		},
	
		// S23
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 122 : // ['a','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S24
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 122 : // ['a','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S25
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 122 : // ['a','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S26
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 109 : // ['a','m']
				return 25
			case r == 110 : // ['n','n']
				return 34
			case 111 <= r && r <= 122 : // ['o','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S27
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 107 : // ['a','k']
				return 25
			case r == 108 : // ['l','l']
				return 35
			case 109 <= r && r <= 122 : // ['m','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S28
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 122 : // ['a','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S29
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 115 : // ['a','s']
				return 25
			case r == 116 : // ['t','t']
				return 36
			case 117 <= r && r <= 122 : // ['u','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S30
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 115 : // ['a','s']
				return 25
			case r == 116 : // ['t','t']
				return 37
			case 117 <= r && r <= 122 : // ['u','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S31
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 122 : // ['a','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S32
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 116 : // ['a','t']
				return 25
			case r == 117 : // ['u','u']
				return 38
			case 118 <= r && r <= 122 : // ['v','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S33
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 19
			case 48 <= r && r <= 57 : // ['0','9']
				return 33
			
			
			
			}
			return NoState
			
		},
	
		// S34
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 115 : // ['a','s']
				return 25
			case r == 116 : // ['t','t']
				return 39
			case 117 <= r && r <= 122 : // ['u','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S35
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 114 : // ['a','r']
				return 25
			case r == 115 : // ['s','s']
				return 40
			case 116 <= r && r <= 122 : // ['t','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S36
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 98 : // ['a','b']
				return 25
			case r == 99 : // ['c','c']
				return 41
			case 100 <= r && r <= 122 : // ['d','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S37
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 122 : // ['a','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S38
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 100 : // ['a','d']
				return 25
			case r == 101 : // ['e','e']
				return 42
			case 102 <= r && r <= 122 : // ['f','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S39
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case r == 97 : // ['a','a']
				return 43
			case 98 <= r && r <= 122 : // ['b','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S40
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 100 : // ['a','d']
				return 25
			case r == 101 : // ['e','e']
				return 44
			case 102 <= r && r <= 122 : // ['f','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S41
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 103 : // ['a','g']
				return 25
			case r == 104 : // ['h','h']
				return 45
			case 105 <= r && r <= 122 : // ['i','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S42
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 122 : // ['a','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S43
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 104 : // ['a','h']
				return 25
			case r == 105 : // ['i','i']
				return 46
			case 106 <= r && r <= 122 : // ['j','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S44
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 122 : // ['a','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S45
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 100 : // ['a','d']
				return 25
			case r == 101 : // ['e','e']
				return 47
			case 102 <= r && r <= 122 : // ['f','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S46
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 109 : // ['a','m']
				return 25
			case r == 110 : // ['n','n']
				return 48
			case 111 <= r && r <= 122 : // ['o','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S47
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 114 : // ['a','r']
				return 25
			case r == 115 : // ['s','s']
				return 49
			case 116 <= r && r <= 122 : // ['t','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S48
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 114 : // ['a','r']
				return 25
			case r == 115 : // ['s','s']
				return 50
			case 116 <= r && r <= 122 : // ['t','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S49
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 122 : // ['a','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
		// S50
		func(r rune) int {
			switch {
			case r == 46 : // ['.','.']
				return 23
			case 48 <= r && r <= 57 : // ['0','9']
				return 24
			case 65 <= r && r <= 90 : // ['A','Z']
				return 25
			case r == 95 : // ['_','_']
				return 23
			case 97 <= r && r <= 122 : // ['a','z']
				return 25
			
			
			
			}
			return NoState
			
		},
	
}