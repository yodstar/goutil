package sqlbuilder

import (
	"testing"
)

type TestArticle struct {
	ID        int64  `db:"id" sqlbuilder:"-"`
	Type      string `db:"type"`
	Title     string `db:"title"`
	Content   string `db:"content"`
	CreatedAt string `db:"createdAt" sqlbuilder:"DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s')"`
}

func TestBuildCountSQL(t *testing.T) {
	a := TestArticle{}
	query, args, err := NewSQLBuilder(&a).BuildCountSQL("id = ?", 1)
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestBuildSelectSQL(t *testing.T) {
	a := TestArticle{}
	query, args, err := NewSQLBuilder(&a).BuildSelectSQL("id IN (?) ORDER BY id DESC LIMIT 5", []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestBuildDeleteSQL(t *testing.T) {
	a := TestArticle{}
	query, args, err := NewSQLBuilder(&a).BuildDeleteSQL("id IN (?)", []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestBuildUpdateSQL(t *testing.T) {
	a := TestArticle{ID: 1, Title: "Title", Content: "Content"}
	query, args, err := NewSQLBuilder(&a).BuildUpdateSQL("id = :id")
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestBuildInsertSQL(t *testing.T) {
	a := TestArticle{Title: "Title", Content: "Content"}
	query, args, err := NewSQLBuilder(&a).BuildInsertSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
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
	a := TestArticle{}
	query, args, err := NewSQLBuilder(&a).Fields("id, title").BuildSelectSQL("id = ?", 1)
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderWhere(t *testing.T) {
	a := TestArticle{}
	sb := NewSQLBuilder(&a)
	query, args, err := sb.Where("id = ?", 1).Fields("id, title").SelectSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)

	query, args, err = sb.WhereOr("id = ?", 2).DeleteSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)

	query, args, err = sb.Fields("id, title").WhereNot("id = ?", 3).UpdateSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderGroupBy(t *testing.T) {
	a := TestArticle{}
	sb := NewSQLBuilder(&a)
	query, args, err := sb.Fields("id, title").Where("id = ?", 1).GroupBy("type").SelectSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderHaving(t *testing.T) {
	a := TestArticle{}
	sb := NewSQLBuilder(&a)
	query, args, err := sb.Fields("id, title").Where("id = ?", 1).GroupBy("type").Having("COUNT(1) > 0").SelectSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderOrderBy(t *testing.T) {
	a := TestArticle{}
	sb := NewSQLBuilder(&a)
	query, args, err := sb.Fields("id, title").Where("type = ?", 1).OrderBy("id DESC").SelectSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderLimit(t *testing.T) {
	a := TestArticle{}
	sb := NewSQLBuilder(&a)
	query, args, err := sb.Fields("id, title").Where("type = ?", 1).OrderBy("id DESC").Limit(1).SelectSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderInsert(t *testing.T) {
	a := TestArticle{}
	query, args, err := NewSQLBuilder(&a).Fields("id, title").InsertSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}
