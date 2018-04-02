package jumpdb

import (
	fmt "fmt"
	"regexp"
	"testing"
)

func setupDB() (*DB, error) {
	db := NewDB()

	e := DBEntry{Path: "/root", Weight: 23}
	if err := db.SetEntry(e.Path, e.Weight); err != nil {
		return nil, fmt.Errorf("Could not add entry: %v", e)
	}

	e = DBEntry{Path: "/rck", Weight: 42}
	if err := db.SetEntry(e.Path, e.Weight); err != nil {
		return nil, fmt.Errorf("Could not add entry: %v", e)
	}

	return db, nil
}

func TestAdd(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Errorf("Could not set up DB: %v", err)
	}

	l := len(db.db.GetPathWeight())
	if l != 2 {
		t.Errorf("Expected 2 matching entries, but got: %d", l)
	}

	pw := db.db.GetPathWeight()
	if w, ok := pw["/rck"]; !ok || w != 42 {
		t.Errorf("Expected '/rck' with weight 42, but got: %t/%d", ok, w)
	}
	if w, ok := pw["/root"]; !ok || w != 23 {
		t.Errorf("Expected '/rck' with weight 42, but got: %t/%d", ok, w)
	}

}

func TestNew(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Errorf("Could not set up DB: %v", err)
	}

	db.IncEntry("/didnotexist")
	pw := db.db.GetPathWeight()
	if w, ok := pw["/didnotexist"]; !ok || w != 1 {
		t.Errorf("Expected '/didnotexist' with weight 1, but got: %t/%d", ok, w)
	}

}

func TestCompletion(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Errorf("Could not set up DB: %v", err)
	}

	r := regexp.MustCompile(".*r.*")
	entries := db.Complete(r)
	l := len(entries)
	if l != 2 {
		t.Errorf("Expected 2 entries, but got: %d", l)
	}

	r = regexp.MustCompile(".*root.*")
	entries = db.Complete(r)
	l = len(entries)
	if l != 1 {
		t.Errorf("Expected 2 entries, but got: %d", l)
	}
}

func TestNormalization(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Errorf("Could not set up DB: %v", err)
	}

	max := int64(1<<63 - 1)
	db.db.PathWeight["/rck"] = max

	root := db.db.PathWeight["/root"]
	if err := db.IncEntry("/rck"); err != nil {
		t.Errorf("Could not normalize DB: %v", err)
	}

	pw := db.db.GetPathWeight()
	if w, ok := pw["/root"]; !ok || w != 1 {
		t.Errorf("Expected '/root' with weight 1, but got: %t/%d", ok, w)
	}

	exp := max - root + 2
	if w, ok := pw["/rck"]; !ok || w != exp {
		t.Errorf("Expected '/rck' with weight %d, but got: %t/%d", exp, ok, w)
	}
}

func TestNormalizationFails(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Errorf("Could not set up DB: %v", err)
	}

	max := int64(1<<63 - 1)
	db.db.PathWeight["/root"] = 1
	db.db.PathWeight["/rck"] = max
	if err := db.IncEntry("/rck"); err != nil {
		t.Errorf("Could not normalize DB: %v", err)
	}

	pw := db.db.GetPathWeight()
	if w, ok := pw["/root"]; !ok || w != 1 {
		t.Errorf("Expected '/root' with weight 1, but got: %t/%d", ok, w)
	}

	exp := max
	if w, ok := pw["/rck"]; !ok || w != exp {
		t.Errorf("Expected '/rck' with weight %d, but got: %t/%d", exp, ok, w)
	}
}
