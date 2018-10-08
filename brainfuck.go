package brainfuck

import (
	"bufio"
	"io"
)

type BrainFuck struct {
	code        string
	codePointer int
	oldPoints   []int

	pointer int
	memory  map[int]byte

	input  io.Reader
	output io.Writer
}

func New(code string, input io.Reader, output io.Writer) *BrainFuck {
	return &BrainFuck{code: code, memory: make(map[int]byte), oldPoints: make([]int, 0), input: input, output: output}
}

func (b *BrainFuck) Run() {
	for b.codePointer < len(b.code) {
		b.interpret()
		b.codePointer++
	}
}

func (b *BrainFuck) interpret() {
	switch b.getCommand() {
	case '>':
		b.pointer++
		break
	case '<':
		b.pointer--
		break
	case '+':
		b.setCurrentByte(b.getCurrentByte() + 1)
		break
	case '-':
		b.setCurrentByte(b.getCurrentByte() - 1)
		break
	case '.':
		byteList := []byte{b.getCurrentByte(),}
		b.output.Write(byteList)
		break
	case ',':
		reader := bufio.NewReader(b.input)
		input, _ := reader.ReadByte()
		b.setCurrentByte(input)
		break
	case '[':
		for b.getCurrentByte() != 0 {
			codePointer := b.codePointer
			for b.codePointer++; b.getCommand() != ']'; b.codePointer++ {
				b.interpret()
			}
			b.codePointer = codePointer
		}

		numberOfJumps := 1
		for numberOfJumps > 0 {
			b.codePointer++
			switch b.getCommand() {
			case '[':
				numberOfJumps++
				break
			case ']':
				numberOfJumps--
				break
			}
		}

		break
	}
}

func (b *BrainFuck) getCommand() rune {
	return rune(b.code[b.codePointer])
}

func (b *BrainFuck) getCurrentByte() byte {
	currentByte, _ := b.memory[b.pointer]
	return currentByte
}

func (b *BrainFuck) setCurrentByte(val byte) {
	b.memory[b.pointer] = val
}
