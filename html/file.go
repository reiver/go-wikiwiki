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
			style+
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
