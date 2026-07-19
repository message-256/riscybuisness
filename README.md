# riscybuisness
a small custom risc vm\
its theoretically turing complete \
but i havent tested it to hard\
usage: \
%instruction(8 bits)%registera(8 bit)%registerb(8 bit)\
instructions must be padded with 0s so fff becomes 0f0f0f\
in the assembler all instructions go\
instruction arg1,arg2\
as well to comment something you write\
#comment\
and all labels are declared as\
label:\
if you dont put it in this way it'll do something funny(most of the time error)\
instructions: \
&emsp; &emsp; arithmatic:\
&emsp; &emsp; all arithmatic functions place the result of arg1 by arg2 in arg1 unless other wise specified\
&emsp; &emsp; mul: multiply\
&emsp; &emsp; sub: subtract\
&emsp; &emsp; add: addition\
&emsp; &emsp; div: division\
&emsp; &emsp; and: bitwise and\
&emsp; &emsp;  or: bitwise or\
&emsp; &emsp;  not: butwise not arg2 and then store it in arg1\
&emsp; &emsp;  xor: bitwise xor\
&emsp; &emsp;  shl: shift left\
&emsp; &emsp;  shr: shift right\
\
&emsp;  special:\
&emsp; &emsp; cmp: compare arg1 to arg2 and {\
&emsp; &emsp; &emsp; set bit 1 of cmpr to 1 if arg1 == arg2\
&emsp; &emsp; &emsp;set bit 2 of cmpr to 1 if arg1 < arg2\
&emsp; &emsp; &emsp;and set bit 3 of cmpr to 1 if arg1 > arg2(this term might not be useful)\
&emsp; &emsp; }\
&emsp; &emsp; exit: exit printing arg1 and arg2\
\
&emsp; control:\
&emsp; &emsp; all control instructions move arg2 into arg1\
&emsp; &emsp; ld: move unconditionally treating arg1 as an address arg2 as number\
&emsp; &emsp; mov: move unconditionally treating both as address\
&emsp; &emsp; movne: mov if previous args to cmp were not equal\
&emsp; &emsp; move: mov if previous args to cmp were equal\
&emsp; &emsp; movl: mov if previous args to cmp were arg1 < arg2\
&emsp; &emsp; movg: mov if previous args to cmp were arg1 \> arg2\
\
&emsp; registers:\
&emsp; &emsp; null: do load into addr[1|2](based on the arguments position) instead use the previous assigned value\
&emsp; &emsp; cmpr: comparison register\
&emsp; &emsp; outputr: output register\
&emsp; &emsp; insr: instruction register(contains the current instruction)\
&emsp; &emsp; insp: instruction pointer(contains the address of the current instruction)\
&emsp; &emsp; addr1: address register (contains address of arg1)\
&emsp; &emsp; addr2: address register (contains address of arg2)\
&emsp; &emsp; intr: interupt register (n/a)\
&emsp; &emsp; r1,10: general purpose registers
\
assembler:\
&emsp; &emsp;takes in a file and then prints the assembled code to stdout
