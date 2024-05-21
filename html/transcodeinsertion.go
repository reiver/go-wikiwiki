package wikiwikihtml

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"sourcecode.social/reiver/go-htmlescape"
	"github.com/reiver/go-path"
	"github.com/reiver/go-skip"
	"github.com/reiver/go-unicode"
	"sourcecode.social/reiver/go-erorr"
)

func transcodeInsertion(writer io.Writer, runescanner io.RuneScanner) (err error) {
	const magic rune = '⁁'

	if nil == writer {
		return errNilWriter
	}
	if nil == runescanner {
		return errNilRuneScanner
	}

	{
		err := skip.AnyRunes(runescanner, unicode.CR)
		if io.EOF == err {
			return err
		}
		if nil != err && !errors.Is(err, io.EOF) {
			return erorr.Errorf("wikiwiki: problem skipping possible CR at the beginning of insertion block: %w", err)
		}
	}

	{
		err := skip.OneExpectedRune(runescanner, magic)
		if skip.ErrRuneNotFound == err {
			return erorr.Errorf("wikiwiki: insertion block does not begin with magic character %q (%U)", magic, magic)
		}
		if nil != err {
			return erorr.Errorf("wikiwiki: could you read one expected rune %q (%U): %w", magic, magic, err)
		}
	}

	{
		const str string = `<figure>`

		_, err = io.WriteString(writer, str)
		if nil != err {
			return erorr.Errorf("wikiwiki: problem writing %q: %w", str, err)
		}
	}
	defer func(){
		const str string = `</figure>`+"\n"

		_, err = io.WriteString(writer, str)
		if nil != err {
			err = erorr.Errorf("wikiwiki: problem writing %q: %w", str, err)
		}
	}()


	{
		err := skip.OneOrMoreRunes(runescanner, unicode.HT, unicode.SP)
		if skip.ErrRuneNotFound == err {
			return erorr.Errorf("wikiwiki: insertion block does not have proper spacing after magic character %q (%U)", magic, magic)
		}
		if nil != err {
			return erorr.Errorf("wikiwiki: problem reading spacing after the magic character %q (%U): %w", magic, magic, err)
		}
	}

	var target string
	{
		var targetBytes []byte

		readloop: for {
			r, size, err := runescanner.ReadRune()
			if size <= 0 {
				switch {
				case errors.Is(err, io.EOF):
		/////////////////////// BREAK
					break readloop
				default:
					return erorr.Errorf("wikiwiki: problem reading rune: %w", err)
				}
			}

			switch r {
			case unicode.HT, unicode.LF, unicode.CR, unicode.SP, unicode.NEL, unicode.LS, unicode.PS:
				err := runescanner.UnreadRune()
				if nil != err {
					return erorr.Errorf("wikiwiki: problem un-reading rune: %w", err)
				}

		/////////////// BREAK
				break readloop
			default:
				targetBytes = append(targetBytes, string(r)...)
			}
		}

		target = string(targetBytes)
	}

	{
		err := skip.AnyRunes(runescanner, unicode.HT, unicode.SP)
		if skip.ErrRuneNotFound == err {
			return erorr.Errorf("wikiwiki: insertion block does not have proper spacing after magic character %q (%U)", magic, magic)
		}
		if nil != err && !errors.Is(err, io.EOF) {
			return erorr.Errorf("wikiwiki: problem reading spacing after the magic character %q (%U): %w", magic, magic, err)
		}
	}

	{
		err := skip.AnyRunes(runescanner, unicode.CR)
		if nil != err && !errors.Is(err, io.EOF) {
			return erorr.Errorf("wikiwiki: problem skipping possible CR at the beginning of insertion block: %w", err)
		}
	}

	{
		err := skip.OneExpectedRunes(runescanner, unicode.LF, unicode.NEL, unicode.LS, unicode.PS)
		if nil != err && !errors.Is(err, io.EOF) {
			return erorr.Errorf("wikiwiki: problem skipping possible CR at the beginning of insertion block: %w", err)
		}
	}

	{
		var ext string = strings.ToLower(path.Ext(target))

		var str string
		switch ext {
		case ".apng",".gif",".png",".jpeg",".jpg",".svg",".webp":
			var escaped string = htmlescape.String(target)
			str = fmt.Sprintf(`<img src="%s" />`, escaped)
		case ".ogv",".mp4",".webm":
			var escaped string = htmlescape.String(target)
			str = fmt.Sprintf(`<video src="%s" controls></video>`, escaped)
		case ".flac", ".oga",".mp3",".wav":
			var escaped string = htmlescape.String(target)
			str = fmt.Sprintf(`<audio src="%s" controls></audio>`, escaped)
		default:
			var escaped string = htmlescape.String(target)
			str = fmt.Sprintf(`<a href="%s" />%s</a>`, escaped, escaped)
		}

		_, err = io.WriteString(writer, str)
		if nil != err {
			return erorr.Errorf("wikiwiki: problem writing %q: %w", str, err)
		}
	}

	return nil
}
