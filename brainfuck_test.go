package brainfuck

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	code := "+[-->-[>>+>-----<<]<--<---]>-.>>>+.>>..+++[.>]<<<<.+++.------.<<-.>>>>+."
	writer := bytes.NewBufferString("")
	b := New(code, nil, writer)
	b.Run()
	assert.Equal(t, "Hello, World!", writer.String())
}

func TestHelloWorld2(t *testing.T) {
	code := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>->+>>+[<]<-]>>.>>---.+++++++..+++.>.<<-.>.+++.------.--------.>+."
	writer := bytes.NewBufferString("")
	b := New(code, nil, writer)
	b.Run()
	assert.Equal(t, "Hello World!", writer.String())
}

func TestHelloWorld3(t *testing.T) {
	code := "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."
	writer := bytes.NewBufferString("")
	b := New(code, nil, writer)
	b.Run()
	assert.Equal(t, "Hello World!\n", writer.String())
}

func TestAlphabet(t *testing.T) {
	code := "++++++++++++++++++++++++++>+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++<[>.+<-]"
	writer := bytes.NewBufferString("")
	b := New(code, nil, writer)
	b.Run()
	assert.Equal(t, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", writer.String())
}

func TestAlphabet2(t *testing.T) {
	code := "[-]>[-]<++++[>++++++<-]>+>>++++++++[<++++++++>-]<+.>>++++[<++++++++>-]<.<<[->+.>.<<]<++++[>++++++<-]>+>+++++++.>.<<[->+.>.<<]++++++++.++."
	writer := bytes.NewBufferString("")
	b := New(code, nil, writer)
	b.Run()
	assert.Equal(t, "A B C D E F G H I J K L M N O P Q R S T U V W X Y Z a b c d e f g h i j k l m n o p q r s t u v w x y z \b\n", writer.String())
}


func TestFibonacci(t *testing.T) {
	code := "+++++++++++>+>>>>++++++++++++++++++++++++++++++++++++++++++++>++++++++++++++++++++++++++++++++<<<<<<[>[>>>>>>+>+<<<<<<<-]>>>>>>>[<<<<<<<+>>>>>>>-]<[>++++++++++[-<-[>>+>+<<<-]>>>[<<<+>>>-]+<[>[-]<[-]]>[<<[>>>+<<<-]>>[-]]<<]>>>[>>+>+<<<-]>>>[<<<+>>>-]+<[>[-]<[-]]>[<<+>>[-]]<<<<<<<]>>>>>[++++++++++++++++++++++++++++++++++++++++++++++++.[-]]++++++++++<[->-<]>++++++++++++++++++++++++++++++++++++++++++++++++.[-]<<<<<<<<<<<<[>>>+>+<<<<-]>>>>[<<<<+>>>>-]<-[>>.>.<<<[-]]<<[>>+>+<<<-]>>>[<<<+>>>-]<<[<+>-]>[<+>-]<<<-]"
	writer := bytes.NewBufferString("")
	b := New(code, nil, writer)
	b.Run()
	assert.Equal(t, "1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89", writer.String())
}
