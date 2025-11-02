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
	query, args, err := NewSqlBuilder(&a).buildCountSQL("id = ?", 1)
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestBuildSelectSQL(t *testing.T) {
	a := TestArticle{}
	query, args, err := NewSqlBuilder(&a).buildSelectSQL("id IN (?) ORDER BY id DESC LIMIT 5", []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestBuildDeleteSQL(t *testing.T) {
	a := TestArticle{}
	query, args, err := NewSqlBuilder(&a).buildDeleteSQL("id IN (?)", []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestBuildUpdateSQL(t *testing.T) {
	a := TestArticle{ID: 1, Title: "Title", Content: "Content"}
	query, args, err := NewSqlBuilder(&a).buildUpdateSQL("id = :id")
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestBuildInsertSQL(t *testing.T) {
	a := TestArticle{Title: "Title", Content: "Content"}
	query, args, err := NewSqlBuilder(&a).buildInsertSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestNewSqlBuilder(t *testing.T) {
	a := TestArticle{}
	sb := NewSqlBuilder(&a)
	if sb.Error != nil {
		t.Fail()
	}
	t.Log(sb)
}

func TestSQLBuilderFields(t *testing.T) {
	a := TestArticle{}
	query, args, err := NewSqlBuilder(&a).Fields("id, title").buildSelectSQL("id = ?", 1)
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderWhere(t *testing.T) {
	a := TestArticle{}
	sb := NewSqlBuilder(&a)
	query, args, err := sb.Where("id = ?", 1).Fields("id, title").buildSelectSQL(sb.sqlWhere)
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)

	query, args, err = sb.WhereOr("id = ?", 2).buildDeleteSQL(sb.sqlWhere)
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)

	query, args, err = sb.Fields("id, title").WhereNot("id = ?", 3).buildUpdateSQL(sb.sqlWhere)
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderGroupBy(t *testing.T) {
	a := TestArticle{}
	sb := NewSqlBuilder(&a)
	query, args, err := sb.Fields("id, title").Where("id = ?", 1).GroupBy("type").buildSelectSQL(sb.sqlWhere)
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderHaving(t *testing.T) {
	a := TestArticle{}
	sb := NewSqlBuilder(&a)
	query, args, err := sb.Fields("id, title").Where("id = ?", 1).GroupBy("type").Having("COUNT(1) > 0").buildSelectSQL(sb.sqlWhere)
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderOrderBy(t *testing.T) {
	a := TestArticle{}
	sb := NewSqlBuilder(&a)
	query, args, err := sb.Fields("id, title").Where("type = ?", 1).OrderBy("id DESC").buildSelectSQL(sb.sqlWhere)
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderLimit(t *testing.T) {
	a := TestArticle{}
	sb := NewSqlBuilder(&a)
	query, args, err := sb.Fields("id, title").Where("type = ?", 1).OrderBy("id DESC").Limit(1).buildSelectSQL(sb.sqlWhere)
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}

func TestSQLBuilderInsert(t *testing.T) {
	a := TestArticle{}
	query, args, err := NewSqlBuilder(&a).Fields("id, title").buildInsertSQL()
	if err != nil {
		t.Fail()
	}
	t.Log(query, args, err)
}
