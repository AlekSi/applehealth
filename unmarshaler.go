package applehealth

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"io"
	"os"
	"path"
)

// Unmarshaler is responsible for parsing Apple Health exported ZIP or XML files.
//
// It embeds Parser as a public field, so Parser's exported methods like Next() and Meta()
// can be called directly.
type Unmarshaler struct {
	*Parser
	closers []io.Closer
	size    int64
}

// NewUnmarshaler creates a new Unmarshaler for a given ZIP or XML file.
func NewUnmarshaler(file string) (*Unmarshaler, error) {
	z, err := zip.OpenReader(file)
	if err == nil {
		return newZipUnmarshaler(z)
	}

	return newXmlUnmarshaler(file)
}

// newZipUnmarshaler creates a new Unmarshaler for a given ZIP file.
func newZipUnmarshaler(z *zip.ReadCloser) (*Unmarshaler, error) {
	// main file name is localized (e.g. it is "экспорт.xml" in Russian),
	// so we have to check all .xml files
	for _, f := range z.File {
		if path.Ext(f.Name) != ".xml" {
			continue
		}
		if path.Base(f.Name) == "export_cda.xml" { // that one is not localized
			continue
		}

		rc, err := f.Open()
		if err != nil {
			_ = z.Close()
			return nil, err
		}

		return &Unmarshaler{
			Parser:  NewParser(xml.NewDecoder(rc)),
			closers: []io.Closer{rc, z},
			size:    f.FileInfo().Size(),
		}, nil
	}

	return nil, errors.New("main export file not found")
}

// newXmlUnmarshaler creates a new Unmarshaler for a given XML file.
func newXmlUnmarshaler(file string) (*Unmarshaler, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	fi, err := f.Stat()
	if err != nil {
		_ = f.Close()
		return nil, err
	}

	return &Unmarshaler{
		Parser:  NewParser(xml.NewDecoder(f)),
		closers: []io.Closer{f},
		size:    fi.Size(),
	}, nil
}

// Close closes the file.
func (p *Unmarshaler) Close() error {
	p.Parser.d = nil
	p.Parser = nil

	var err error
	for _, c := range p.closers {
		if e := c.Close(); err == nil {
			err = e
		}
	}
	return err
}

// Stats represent Unmarshaler's statistics.
type Stats struct {
	Pos  int64
	Size int64
}

// Stats return the current Unmarshaler's statistics.
func (p *Unmarshaler) Stats() Stats {
	return Stats{
		Pos:  p.InputOffset(),
		Size: p.size,
	}
}
