package logfs_test

import (
	"testing"

	logfs "github.com/jncornett/afero-logfs"
	"github.com/spf13/afero"
)

type testLogger struct {
	records []logfs.Record
}

func (l *testLogger) Log(r logfs.Record) {
	l.records = append(l.records, r)
}

func TestLogFs(t *testing.T) {
	logger := &testLogger{}
	fs := &logfs.Fs{
		Logger: logger,
		Fs:     afero.NewMemMapFs(),
	}
	f, err := fs.Create("foo")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	if 1 != len(logger.records) {
		t.Fatal("expected len(records) to be %v, got %v", 1, len(logger.records))
	}
	r := logger.records[0]
	if "create" != r.Op {
		t.Error("expected op to be %v, got %v", "create", r.Op)
	}
	if 1 != len(r.Args) {
		t.Fatal("expected len(Args) to be %v, got %v", 1, len(r.Args))
	}
	if "foo" != r.Args[0].(string) {
		t.Error("expected Args[0] to be %q, got %+v", "foo", r.Args[0].(string))
	}
	if nil != r.Error {
		t.Error("expected Err to be nil, got %v", r.Error)
	}
}
