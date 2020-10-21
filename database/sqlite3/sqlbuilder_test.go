package sqlite3

import (
	"testing"
)

type TestArticle struct {
	ID        int64  `db:"id" sqlite3:"-"`
	Type      string `db:"type"`
	Title     string `db:"title"`
	Content   string `db:"content"`
	CreatedAt string `db:"createdAt" sqlite3:"DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s')"`
}

func TestNewSqlBuilder(t *testing.T) {
	a := TestArticle{}
	sb := NewSQLBuilder(&a)
	if sb.Error != nil {
		t.Fail()
	}
	t.Log(sb)
}

func TestSQLBuilderFields(t *testing.T) {
	query, args, err := Fields("id, title").Where("id = ?", 1).SelectSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderWhere(t *testing.T) {
	query, args, err := Where("id = ?", 1).Fields("id, title").SelectSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)

	query, args, err = WhereOr("id = ?", 2).DeleteSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)

	query, args, err = Fields("id, title").WhereNot("id = ?", 3).UpdateSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderGroupBy(t *testing.T) {
	query, args, err := Fields("id, title").Where("id = ?", 1).GroupBy("type").SelectSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderHaving(t *testing.T) {
	query, args, err := Fields("id, title").Where("id = ?", 1).GroupBy("type").Having("COUNT(1) > 0").SelectSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderOrderBy(t *testing.T) {
	query, args, err := Fields("id, title").Where("type = ?", 1).OrderBy("id DESC").SelectSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderLimit(t *testing.T) {
	query, args, err := Fields("id, title").Where("type = ?", 1).OrderBy("id DESC").Limit(1).SelectSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderInsert(t *testing.T) {
	query, args, err := Fields("id, title").InsertSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}
