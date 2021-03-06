package sq

import (
	"log"
	"os"
)

// LogFlag is a flag that affects the verbosity of the Logger output.
type LogFlag int

// LogFlags
const (
	Linterpolate LogFlag = 1 << iota
	Lstats
	Lresults
	Lparse
	Lverbose = Lstats | Lresults
)

// ExecFlag is a flag that affects the behavior of Exec.
type ExecFlag int

// ExecFlags
const (
	ElastInsertID ExecFlag = 1 << iota
	ErowsAffected
)

var defaultLogger = log.New(os.Stdout, "[sq] ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)

// BaseQuery is a common query builder that can transform into a SelectQuery,
// InsertQuery, UpdateQuery or DeleteQuery depending on the method that you
// call on it.
type BaseQuery struct {
	DB      DB
	Log     Logger
	LogFlag LogFlag
	CTEs    CTEs
}

// WithLog creates a new BaseQuery with a custom logger and the LogFlag.
func WithLog(logger Logger, flag LogFlag) BaseQuery {
	return BaseQuery{
		Log:     logger,
		LogFlag: flag,
	}
}

// WithDefaultLog creates a new BaseQuery with the default logger and the LogFlag
func WithDefaultLog(flag LogFlag) BaseQuery {
	return BaseQuery{
		Log:     defaultLogger,
		LogFlag: flag,
	}
}

// WithDB creates a new BaseQuery with the DB.
func WithDB(db DB) BaseQuery {
	return BaseQuery{
		DB: db,
	}
}

// With creates a new BaseQuery with the CTEs.
func With(CTEs ...CTE) BaseQuery {
	return BaseQuery{
		CTEs: CTEs,
	}
}

// WithDefaultLog adds the default logger and the LogFlag to the BaseQuery.
func (q BaseQuery) WithDefaultLog(flag LogFlag) BaseQuery {
	q.Log = defaultLogger
	q.LogFlag = flag
	return q
}

// WithLog adds a custom logger and the LogFlag to the BaseQuery.
func (q BaseQuery) WithLog(logger Logger, flag LogFlag) BaseQuery {
	q.Log = logger
	q.LogFlag = flag
	return q
}

// WithDB adds the DB to the BaseQuery.
func (q BaseQuery) WithDB(db DB) BaseQuery {
	q.DB = db
	return q
}

// With adds the CTEs to the BaseQuery.
func (q BaseQuery) With(CTEs ...CTE) BaseQuery {
	q.CTEs = append(q.CTEs, CTEs...)
	return q
}

// From transforms the BaseQuery into a SelectQuery.
func (q BaseQuery) From(table Table) SelectQuery {
	return SelectQuery{
		FromTable: table,
		Alias:     RandomString(8),
		CTEs:      q.CTEs,
		DB:        q.DB,
		Log:       q.Log,
		LogFlag:   q.LogFlag,
	}
}

// Select transforms the BaseQuery into a SelectQuery.
func (q BaseQuery) Select(fields ...Field) SelectQuery {
	return SelectQuery{
		SelectFields: fields,
		Alias:        RandomString(8),
		CTEs:         q.CTEs,
		DB:           q.DB,
		Log:          q.Log,
		LogFlag:      q.LogFlag,
	}
}

// SelectOne transforms the BaseQuery into a SelectQuery.
func (q BaseQuery) SelectOne() SelectQuery {
	return SelectQuery{
		SelectFields: Fields{FieldLiteral("1")},
		Alias:        RandomString(8),
		CTEs:         q.CTEs,
		DB:           q.DB,
		Log:          q.Log,
		LogFlag:      q.LogFlag,
	}
}

// SelectAll transforms the BaseQuery into a SelectQuery.
func (q BaseQuery) SelectAll() SelectQuery {
	return SelectQuery{
		SelectFields: Fields{FieldLiteral("*")},
		Alias:        RandomString(8),
		CTEs:         q.CTEs,
		DB:           q.DB,
		Log:          q.Log,
		LogFlag:      q.LogFlag,
	}
}

// SelectCount transforms the BaseQuery into a SelectQuery.
func (q BaseQuery) SelectCount() SelectQuery {
	return SelectQuery{
		SelectFields: Fields{FieldLiteral("COUNT(*)")},
		Alias:        RandomString(8),
		CTEs:         q.CTEs,
		DB:           q.DB,
		Log:          q.Log,
		LogFlag:      q.LogFlag,
	}
}

// SelectDistinct transforms the BaseQuery into a SelectQuery.
func (q BaseQuery) SelectDistinct(fields ...Field) SelectQuery {
	return SelectQuery{
		SelectType:   SelectTypeDistinct,
		SelectFields: fields,
		Alias:        RandomString(8),
		CTEs:         q.CTEs,
		DB:           q.DB,
		Log:          q.Log,
		LogFlag:      q.LogFlag,
	}
}

// Selectx transforms the BaseQuery into a SelectQuery.
func (q BaseQuery) Selectx(mapper func(*Row), accumulator func()) SelectQuery {
	return SelectQuery{
		Mapper:      mapper,
		Accumulator: accumulator,
		Alias:       RandomString(8),
		CTEs:        q.CTEs,
		DB:          q.DB,
		Log:         q.Log,
		LogFlag:     q.LogFlag,
	}
}

// SelectRowx transforms the BaseQuery into a SelectQuery.
func (q BaseQuery) SelectRowx(mapper func(*Row)) SelectQuery {
	return SelectQuery{
		Mapper:  mapper,
		Alias:   RandomString(8),
		CTEs:    q.CTEs,
		DB:      q.DB,
		Log:     q.Log,
		LogFlag: q.LogFlag,
	}
}

// InsertInto transforms the BaseQuery into an InsertQuery.
func (q BaseQuery) InsertInto(table BaseTable) InsertQuery {
	return InsertQuery{
		IntoTable: table,
		Alias:     RandomString(8),
		DB:        q.DB,
		Log:       q.Log,
		LogFlag:   q.LogFlag,
	}
}

// InsertIgnoreInto transforms the BaseQuery into an InsertQuery.
func (q BaseQuery) InsertIgnoreInto(table BaseTable) InsertQuery {
	return InsertQuery{
		Ignore:    true,
		IntoTable: table,
		Alias:     RandomString(8),
		DB:        q.DB,
		Log:       q.Log,
		LogFlag:   q.LogFlag,
	}
}

// Update transforms the BaseQuery into an UpdateQuery.
func (q BaseQuery) Update(table BaseTable) UpdateQuery {
	return UpdateQuery{
		UpdateTable: table,
		Alias:       RandomString(8),
		CTEs:        q.CTEs,
		DB:          q.DB,
		Log:         q.Log,
		LogFlag:     q.LogFlag,
	}
}

// DeleteFrom transforms the BaseQuery into a DeleteQuery.
func (q BaseQuery) DeleteFrom(tables ...BaseTable) DeleteQuery {
	return DeleteQuery{
		FromTables: tables,
		Alias:      RandomString(8),
		CTEs:       q.CTEs,
		DB:         q.DB,
		Log:        q.Log,
		LogFlag:    q.LogFlag,
	}
}
