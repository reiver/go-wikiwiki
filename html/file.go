package wikiwikihtml

import (
	"io"
	"io/fs"
	"time"
)

// File takes a fs.File with wiki content, and returns a new fs.File with HTML content.
//
// Note that the file-size of this will always return -1 regardless of what the file-size of the HTML content actually ends up being.
func File(wikifile fs.File) fs.File {
	if nil == wikifile {
		return nil
	}

	pipereader, pipewriter := io.Pipe()

	var f *internalFile = &internalFile{
		file:wikifile,
		pipereader:pipereader,
		pipewriter:pipewriter,
	}

	go f.transcode()

	return f
}

type internalFile struct {
	file fs.File
	pipereader *io.PipeReader
	pipewriter *io.PipeWriter
}

func (receiver internalFile) Close() error {
	var file fs.File = receiver.file
	if nil == file {
		return errNilFile
	}

	return file.Close()
}

func (receiver internalFile) IsDir() bool {
	var file fs.File = receiver.file
	if nil == file {
		return false
	}

	var fileinfo fs.FileInfo
	var err error

	fileinfo, err = file.Stat()
	if nil != err {
		return false
	}

	return fileinfo.IsDir()
}

func (receiver internalFile) Mode() fs.FileMode {
	var defaultFileMode fs.FileMode

	var file fs.File = receiver.file
	if nil == file {
		return defaultFileMode
	}

	var fileinfo fs.FileInfo
	var err error

	fileinfo, err = file.Stat()
	if nil != err {
		return defaultFileMode
	}

	return fileinfo.Mode()
}

func (receiver internalFile) ModTime() time.Time {
	var defaultTime time.Time

	var file fs.File = receiver.file
	if nil == file {
		return defaultTime
	}

	var fileinfo fs.FileInfo
	var err error

	fileinfo, err = file.Stat()
	if nil != err {
		return defaultTime
	}

	return fileinfo.ModTime()
}

func (receiver internalFile) Name() string {
	var file fs.File = receiver.file
	if nil == file {
		return ""
	}

	var fileinfo fs.FileInfo
	var err error

	fileinfo, err = file.Stat()
	if nil != err {
		return ""
	}

	return fileinfo.Name()+".html"
}

func (receiver internalFile) Read(p []byte) (int, error) {
	var reader io.Reader = receiver.pipereader
	if nil == reader {
		return 0, errNilPipeReader
	}

	return reader.Read(p)
}

// Size returns -1.
//
// We don't know the size ahead of time, so we are intentionally returning a non-sensical size.
func (receiver internalFile) Size() int64 {
	return -1
}

func (receiver *internalFile) Stat() (fs.FileInfo, error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	return receiver, nil
}

func (receiver internalFile) Sys() any {
	return receiver.file
}

func (receiver internalFile) transcode() {
	var file fs.File = receiver.file
	if nil == file {
		panic(errNilFile)
	}

	var writer *io.PipeWriter = receiver.pipewriter
	if nil == writer {
		panic(errNilPipeWriter)
	}
	defer func(){
		err := writer.Close()
		if nil != err {
			panic(err)
		}
	}()

	{
		_, err := io.WriteString(writer,
			`<html>`+"\n"+
			`<head>`+"\n"+
			`<meta charset="utf-8" />`+"\n"+
			`<style type="text/css">`+"\n"+
`
html, body {
	height: 100%;
}

body {
	font-family: serif;
	display: flex;
	margin: auto;
	justify-content: center;
}
body {
	color: #333333;
	background-color:white;
}

main {
	margin: 75pt 11.25pt;
	max-width: 502.5pt;

	@media screen and (min-width: 450pt){
		padding: 0 22.5pt;
	}
}

h1,h2,h3,h4,h5,h6 {
	color: #111111;
	font-weight: 700;
}
h1 {
	text-align:center;
}
h2 {
	border-bottom:1px solid #d0d0d0;
	margin:21pt 0 0 0;
	padding-bottom:3pt;
}
h3 {
	text-transform:uppercase;
}

p {
	font-size: 12pt;
	line-height: 1.625;
	word-wrap: break-word;
	overflow-wrap: break-word;
	hyphens: auto;

	@media screen and (min-width: 450pt){
		font-size: 13.5pt;
		line-height: 1.667;
	}
}
p, li {
	text-align: justify;
}

img[align="right"] {
	margin-left:1.5em;
}

body {
	counter-reset: citation-number-counter;
}
sup.citation a:before {
	counter-increment: citation-number-counter +1;
	content: "[" counter(citation-number-counter) "]";
}
`+

			`a.wiki-link::before{`+"\n"+
			`content:attr(href);`+"\n"+
			`}`+"\n"+

			`</style>`+"\n"+
			`</head>`+"\n"+
			`<body>`+"\n"+
			`<main>`+"\n",
		)
		if nil != err {
			writer.CloseWithError(err)
			return
		}
	}

	err := Transcode(writer, file)
	if nil != err {
		writer.CloseWithError(err)
		return
	}

	{
		_, err := io.WriteString(writer,
			`</main>`+"\n"+
			`</body>`+"\n"+
			`</html>`+"\n",
		)
		if nil != err {
			writer.CloseWithError(err)
			return
		}
	}
}
