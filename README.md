# riscybuisness
a small custom risc vm\
its theoretically turing complete \
but i havent tested it to hard\
usage: \
%instruction(8 bits)%registera(8 bit)%registerb(8 bit)
instructions must be padded\
in the assembler all instructions go i arg1,arg2\
if you dont put it in this way it'll do something funny(im not really sure what)\
instructions: \
  arithmatic:
  all arithmatic functions place the result of arg1 by arg2 in arg1 unless other wise specified\
  mul: multiply\
  sub: subtract\
  add: addition\
  div: division\
  and: bitwise and\
  or: bitwise or\
  not: butwise not arg2 and then store it in arg1\
  xor: bitwise xor\
  shl: shift left\
  shr: shift right\
  special:\
  cmp: compare arg1 to arg2 and {\
    set bit 1 of cmpr to 1 if arg1 == arg2\
    set bit 2 of cmpr to 1 if arg1 < arg2\
    and set but 3 of cmpr to 1 if arg1 > arg2(this term might not be useful)\
    }\
  exit: exit printing arg1 and arg2\
  control:\
  all control instructions move arg2 into arg1:\
  ld: move unconditionally treating arg1 as an address arg2 as number\
  mov: move unconditionally treating both as address\
  movne: mov if previous args to cmp were not equal\
  move: mov if previous args to cmp were equal\
  movl: mov if previous args to cmp were arg1 < arg2\
  movg: mov if previous args to cmp were arg1 \> arg2\
registers:\
  cmpr: comparison register\
  outputr: output register\
  insr: instruction register(contains the current instruction)\
  insp: instruction pointer(contains the address of the current instruction)\
  addr1: address register (contains address of arg1)\
  addr2: address register (contains address of arg2)\
  intr: interupt register (n/a)
  r1,10: general purpose registers
  
  
