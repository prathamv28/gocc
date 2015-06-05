package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* + */
			nil,      /* * */
			shift(5), /* ( */
			nil,      /* ) */
			shift(6), /* int64 */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* + */
			nil,          /* * */
			nil,          /* ( */
			nil,          /* ) */
			nil,          /* int64 */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Calc */
			shift(7),  /* + */
			nil,       /* * */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* int64 */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Expr */
			reduce(3), /* +, reduce: Expr */
			shift(8),  /* * */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* int64 */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: Term */
			reduce(5), /* +, reduce: Term */
			reduce(5), /* *, reduce: Term */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* int64 */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* + */
			nil,       /* * */
			shift(12), /* ( */
			nil,       /* ) */
			shift(13), /* int64 */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(7), /* $, reduce: Factor */
			reduce(7), /* +, reduce: Factor */
			reduce(7), /* *, reduce: Factor */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* int64 */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* + */
			nil,      /* * */
			shift(5), /* ( */
			nil,      /* ) */
			shift(6), /* int64 */

		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* + */
			nil,      /* * */
			shift(5), /* ( */
			nil,      /* ) */
			shift(6), /* int64 */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* + */
			nil,       /* * */
			nil,       /* ( */
			shift(17), /* ) */
			nil,       /* int64 */

		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(3), /* +, reduce: Expr */
			shift(18), /* * */
			nil,       /* ( */
			reduce(3), /* ), reduce: Expr */
			nil,       /* int64 */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(5), /* +, reduce: Term */
			reduce(5), /* *, reduce: Term */
			nil,       /* ( */
			reduce(5), /* ), reduce: Term */
			nil,       /* int64 */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* + */
			nil,       /* * */
			shift(12), /* ( */
			nil,       /* ) */
			shift(13), /* int64 */

		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(7), /* +, reduce: Factor */
			reduce(7), /* *, reduce: Factor */
			nil,       /* ( */
			reduce(7), /* ), reduce: Factor */
			nil,       /* int64 */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Expr */
			reduce(2), /* +, reduce: Expr */
			shift(8),  /* * */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* int64 */

		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: Term */
			reduce(4), /* +, reduce: Term */
			reduce(4), /* *, reduce: Term */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* int64 */

		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* + */
			nil,       /* * */
			shift(12), /* ( */
			nil,       /* ) */
			shift(13), /* int64 */

		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(6), /* $, reduce: Factor */
			reduce(6), /* +, reduce: Factor */
			reduce(6), /* *, reduce: Factor */
			nil,       /* ( */
			nil,       /* ) */
			nil,       /* int64 */

		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* + */
			nil,       /* * */
			shift(12), /* ( */
			nil,       /* ) */
			shift(13), /* int64 */

		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			shift(16), /* + */
			nil,       /* * */
			nil,       /* ( */
			shift(22), /* ) */
			nil,       /* int64 */

		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(2), /* +, reduce: Expr */
			shift(18), /* * */
			nil,       /* ( */
			reduce(2), /* ), reduce: Expr */
			nil,       /* int64 */

		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(4), /* +, reduce: Term */
			reduce(4), /* *, reduce: Term */
			nil,       /* ( */
			reduce(4), /* ), reduce: Term */
			nil,       /* int64 */

		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			reduce(6), /* +, reduce: Factor */
			reduce(6), /* *, reduce: Factor */
			nil,       /* ( */
			reduce(6), /* ), reduce: Factor */
			nil,       /* int64 */

		},
	},
}