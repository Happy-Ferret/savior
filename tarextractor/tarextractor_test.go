package tarextractor_test

import (
	"bytes"
	"testing"

	"github.com/itchio/arkive/tar"
	"github.com/itchio/savior"
	"github.com/stretchr/testify/assert"

	"github.com/itchio/savior/seeksource"
	"github.com/itchio/savior/tarextractor"
)

func must(t *testing.T, err error) {
	assert.NoError(t, err)
	if err != nil {
		t.FailNow()
	}
}

func TestTar(t *testing.T) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	helloData := []byte("Hello, it's me!")

	err := tw.WriteHeader(&tar.Header{
		Typeflag: tar.TypeReg,
		Name:     "hello.txt",
		Mode:     0644,
		Size:     int64(len(helloData)),
	})
	must(t, err)

	_, err = tw.Write(helloData)
	must(t, err)

	err = tw.Close()
	must(t, err)

	source := seeksource.New(bytes.NewReader(buf.Bytes()))

	ex := tarextractor.New()

	err = ex.Configure(&savior.ExtractorParams{
		LastCheckpoint: nil,
		OnProgress:     nil,
		Source:         source,
		Sink: &savior.Sink{
			Directory: "./ignored",
		},
	})
	must(t, err)

	err = ex.Work()
	must(t, err)
}
