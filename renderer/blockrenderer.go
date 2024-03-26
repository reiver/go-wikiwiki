package wikiwikirenderer

import (
	"io"
)

type BlockRenderer interface {
	BeginRenderingBlock(io.Writer)error
	EndRenderingBlock(io.Writer)error
	BeginRenderingLine(io.Writer)error
	EndRenderingLine(io.Writer)error
	RenderLine(io.Writer,io.RuneScanner)error
}
