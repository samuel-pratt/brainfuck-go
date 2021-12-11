# brainfuck-go

A [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) interpreter written in Go.

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