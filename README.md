# brainfuck-go

A [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) interpreter written in Go.

#### How Brainfuck works

Brainfuck is an esoteric programming language designed to have the simplest compiler possible. Beginning with a byte array with a length of 30 thousand, and a pointer to the first element, the programmer can perform 8 actions:

`>` - increment the data pointer (to point to the next cell to the right).

`<` - decrement the data pointer (to point to the next cell to the left).

`+` - increment (increase by one) the byte at the data pointer.

`-` - decrement (decrease by one) the byte at the data pointer.

`.` - output the byte at the data pointer.

`,` - accept one byte of input, storing its value in the byte at the data pointer.

`[` - if the byte at the data pointer is zero, then instead of moving the instruction pointer forward to the next command, jump it forward to the command after the matching ] command.

`]` - if the byte at the data pointer is nonzero, then instead of moving the instruction pointer forward to the next command, jump it back to the command after the matching [ command.

### Running locally

Requires Go to build.

```bash
git clone https://github.com/samuel-pratt/brainfuck-go.git
cd brainfuck-go
go build
./brainfuck-go hello_world.bf
```

#### Thanks to

Fireship, who's [video](https://www.youtube.com/watch?v=hdHjjBS4cs8) introduced my to Brainfuck.

Steve Kemp, who's [article](https://blog.steve.fi/writing_a_brainfuck_compiler_.html) was a huge help in writing the interpreter (and where I got the steps above from).
