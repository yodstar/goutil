package sqlbuilder

import (
	"os"
	"testing"
	"time"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

const TestCreateSQL = `CREATE TABLE IF NOT EXISTS article (
	id INTEGER PRIMARY KEY NOT NULL,
	title VARCHAR(255) NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME
)`

var (
	options   *Options
	dbSqlite3 *Dao
)

type Article struct {
	ID          int64  `db:"id" sqlbuilder:"-"`
	Title       string `db:"title"`
	Content     string `db:"content"`
	CreatedAt   string `db:"created_at"`
	CreatedDate string `db:"createdDate" sqlbuilder:"strftime('%Y-%m-%d', created_at)"`
	CreatedTime string `db:"createdTime" sqlbuilder:"strftime('%H:%M:%S', created_at)"`
}

func TestMain(m *testing.M) {
	dbSqlite3 = MustOpen("sqlite3", Option{DataSourceName: "./database_test.db3"})
	m.Run()
	os.Remove("./database_test.db3")
}

func TestDatabaseWriter(t *testing.T) {
	_, err := dbSqlite3.Exec(TestCreateSQL)
	if err != nil {
		t.Fail()
	}
	t.Log(err)
}

func TestDatabaseInsert(t *testing.T) {
	a := Article{Title: "Title1", Content: "Content1", CreatedAt: time.Now().Format("2006-01-02 15:04:05")}
	_, err := dbSqlite3.Insert(&a)
	if err != nil {
		t.Fail()
	}
	t.Log(err)
}

func TestDatabaseCount(t *testing.T) {
	count, err := dbSqlite3.Count(&Article{}, "")
	if err != nil {
		t.Fail()
	}
	t.Log(count, err)
}

func TestDatabaseSelect(t *testing.T) {
	a := []Article{}
	err := dbSqlite3.Select(&a, "")
	if err != nil {
		t.Fail()
	}
	t.Log(a, err)
}

func TestDatabaseUpdate(t *testing.T) {
	a := Article{}
	err := dbSqlite3.Select(&a, "id = ?", 1)
	if err != nil {
		t.Fail()
	}
	a.Title = "Title2"
	a.Content = "Content2"
	a.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	res, err := dbSqlite3.Update(&a, "id = :id")
	if err != nil {
		t.Fail()
	}
	t.Log(res, err)
}

func TestDatabaseDelete(t *testing.T) {
	a := []Article{}
	err := dbSqlite3.Select(&a, "")
	if err != nil {
		t.Fail()
	}
	t.Log(a, err)
	res, err := dbSqlite3.Delete(&a, "id = ?", 1)
	if err != nil {
		t.Fail()
	}
	t.Log(res, err)
}
