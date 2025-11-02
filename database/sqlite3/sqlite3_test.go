package sqlite3

import (
	"os"
	"testing"
	"time"

	"github.com/yodstar/goutil/database/sqlbuilder"
)

const TestCreateSQL = `CREATE TABLE IF NOT EXISTS article (
	id INTEGER PRIMARY KEY NOT NULL,
	title VARCHAR(255) NOT NULL,
	content TEXT NOT NULL,
	created_at DATETIME
)`

var (
	option = &sqlbuilder.Option{DataSourceName: "./sqlite3_test.db3"}
)

type Article struct {
	ID          int64  `db:"id" sqlite3:"-"`
	Title       string `db:"title"`
	Content     string `db:"content"`
	CreatedAt   string `db:"created_at"`
	CreatedDate string `db:"createdDate" sqlite3:"strftime('%Y-%m-%d', created_at)"`
	CreatedTime string `db:"createdTime" sqlite3:"strftime('%H:%M:%S', created_at)"`
}

func TestMain(m *testing.M) {
	MustOpen(sqlbuilder.Options{option})
	m.Run()
	os.Remove("./sqlite3_test.db3")
}

func TestExec(t *testing.T) {
	_, err := Exec(TestCreateSQL)
	if err != nil {
		t.Fail()
	}
	t.Log(err)
}

func TestInsert(t *testing.T) {
	a := Article{Title: "Title1", Content: "Content1", CreatedAt: time.Now().Format("2006-01-02 15:04:05")}
	_, err := Insert(&a)
	if err != nil {
		t.Fail()
	}
	t.Log(err)
}

func TestCount(t *testing.T) {
	count, err := Count(&Article{}, "")
	if err != nil {
		t.Fail()
	}
	t.Log(count, err)
}

func TestSelect(t *testing.T) {
	a := []Article{}
	err := Select(&a, "")
	if err != nil {
		t.Fail()
	}
	t.Log(a, err)
}

func TestUpdate(t *testing.T) {
	a := Article{}
	err := Select(&a, "id = ?", 1)
	if err != nil {
		t.Fail()
	}
	a.Title = "Title2"
	a.Content = "Content2"
	a.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	res, err := Update(&a, "id = :id")
	if err != nil {
		t.Fail()
	}
	t.Log(res, err)
}

func TestDelete(t *testing.T) {
	a := []Article{}
	err := Select(&a, "")
	if err != nil {
		t.Fail()
	}
	t.Log(a, err)
	res, err := Delete(&a, "id = ?", 1)
	if err != nil {
		t.Fail()
	}
	t.Log(res, err)
}
