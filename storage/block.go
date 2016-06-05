package storage

import (
	"io"
	"time"

	"code.uber.internal/infra/memtsdb"
	"code.uber.internal/infra/memtsdb/encoding"
	xtime "code.uber.internal/infra/memtsdb/x/time"
)

type dbBlock struct {
	start   time.Time
	opts    memtsdb.DatabaseOptions
	encoder encoding.Encoder
}

// NewDatabaseBlock creates a new DatabaseBlock instance.
func NewDatabaseBlock(start time.Time, data []byte, opts memtsdb.DatabaseOptions) memtsdb.DatabaseBlock {
	newEncoderFn := opts.GetNewEncoderFn()
	return &dbBlock{
		start:   start,
		opts:    opts,
		encoder: newEncoderFn(start, data),
	}
}

func (b *dbBlock) StartTime() time.Time {
	return b.start
}

func (b *dbBlock) Write(timestamp time.Time, value float64, unit xtime.Unit, annotation []byte) error {
	return b.encoder.Encode(encoding.Datapoint{Timestamp: timestamp, Value: value}, annotation, unit)
}

func (b *dbBlock) Stream() io.Reader {
	return b.encoder.Stream()
}

type databaseSeriesBlocks struct {
	elems  map[time.Time]memtsdb.DatabaseBlock
	min    time.Time
	max    time.Time
	dbOpts memtsdb.DatabaseOptions
}

// NewDatabaseSeriesBlocks creates a databaseSeriesBlocks instance.
func NewDatabaseSeriesBlocks(dbOpts memtsdb.DatabaseOptions) memtsdb.DatabaseSeriesBlocks {
	return &databaseSeriesBlocks{
		elems:  make(map[time.Time]memtsdb.DatabaseBlock),
		dbOpts: dbOpts,
	}
}

func (dbb *databaseSeriesBlocks) AddBlock(block memtsdb.DatabaseBlock) {
	start := block.StartTime()
	if dbb.min.Equal(timeNone) || start.Before(dbb.min) {
		dbb.min = start
	}
	if dbb.max.Equal(timeNone) || start.After(dbb.max) {
		dbb.max = start
	}
	dbb.elems[start] = block
}

func (dbb *databaseSeriesBlocks) AddSeries(other memtsdb.DatabaseSeriesBlocks) {
	if other == nil {
		return
	}
	blocks := other.GetAllBlocks()
	for _, b := range blocks {
		dbb.AddBlock(b)
	}
}

// GetMinTime returns the min time of the blocks contained.
func (dbb *databaseSeriesBlocks) GetMinTime() time.Time {
	return dbb.min
}

// GetMaxTime returns the max time of the blocks contained.
func (dbb *databaseSeriesBlocks) GetMaxTime() time.Time {
	return dbb.max
}

func (dbb *databaseSeriesBlocks) GetBlockAt(t time.Time) (memtsdb.DatabaseBlock, bool) {
	b, ok := dbb.elems[t]
	return b, ok
}

func (dbb *databaseSeriesBlocks) GetBlockOrAdd(t time.Time) memtsdb.DatabaseBlock {
	b, ok := dbb.elems[t]
	if ok {
		return b
	}
	newBlock := NewDatabaseBlock(t, nil, dbb.dbOpts)
	dbb.AddBlock(newBlock)
	return newBlock
}

func (dbb *databaseSeriesBlocks) GetAllBlocks() map[time.Time]memtsdb.DatabaseBlock {
	return dbb.elems
}
