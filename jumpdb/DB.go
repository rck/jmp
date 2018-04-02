package jumpdb

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"text/tabwriter"

	"github.com/golang/protobuf/proto"
)

const version = 1

// DBEntry contains entries in the database
type DBEntry struct {
	Path   string
	Weight int64
}

type byWeight []DBEntry

// sort Interface
func (e byWeight) Len() int           { return len(e) }
func (e byWeight) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e byWeight) Less(i, j int) bool { return e[i].Weight < e[j].Weight }

// DB represents the internal database
type DB struct {
	db     *Database
	loaded chan error
}

// NewDB returns a pointer to a DB
func NewDB() *DB {
	return &DB{
		db: &Database{
			Version:    version,
			PathWeight: make(map[string]int64),
		},
	}
}

func (d *DB) load(fileName string) {
	defer close(d.loaded)

	fi, err := os.Stat(fileName)
	if err != nil {
		// d.loaded <- fmt.Errorf("Could not stat %s: %v", d.path, err)
		// This is okay, maybe a fresh install, so start with an empty db
		return
	}
	if !fi.Mode().IsRegular() {
		d.loaded <- fmt.Errorf("Path %s is not a regular file", fileName)
		return
	}

	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		d.loaded <- err
		return
	}
	d.loaded <- proto.Unmarshal(in, d.db)
}

// Load loads the database from the file system synchronously
func (d *DB) Load(fileName string) error {
	d.loaded = make(chan error)
	d.load(fileName)
	return <-d.loaded
}

// LoadAsync loads the database from the flile system asynchronously
func (d *DB) LoadAsync(fileName string) chan error {
	d.loaded = make(chan error)
	go d.load(fileName)
	return d.loaded
}

// List lists entries and their weight
func (d *DB) List() error {
	var entries []DBEntry
	for p, w := range d.db.GetPathWeight() {
		entries = append(entries, DBEntry{Path: p, Weight: w})
	}
	sort.Sort(sort.Reverse(byWeight(entries)))

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "Weight\t | Path:")
	for _, e := range entries {
		fmt.Fprintf(w, "%d\t | %s\n", e.Weight, e.Path)
	}
	return w.Flush()
}

// Save stores the DB to the file system
func (d *DB) Save(fileName string) error {
	out, err := proto.Marshal(d.db)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, out, 0644)
}

func (d *DB) normalize() error {
	min := int64(1<<63 - 1)

	for _, w := range d.db.GetPathWeight() {
		if w < min {
			min = w
		}
	}

	if min == 1 {
		return fmt.Errorf("Can not normalize DB, minimal entry has weight of 1")
	}

	for p, w := range d.db.GetPathWeight() {
		d.db.PathWeight[p] = w - min + 1
	}

	return nil
}

func (d *DB) setEntry(entry DBEntry) error {
	if entry.Weight == 0 {
		delete(d.db.PathWeight, entry.Path)
		return nil
	}
	d.db.PathWeight[entry.Path] = entry.Weight
	return nil
}

// IncEntry increases the weight by one
func (d *DB) IncEntry(entry DBEntry) error {
	if entry.Weight == 1<<63-1 {
		if err := d.normalize(); err != nil {
			// Could not normalize, keep everything as is
			return nil
		}
	}
	curWeight, ok := d.db.PathWeight[entry.Path]
	if !ok {
		return fmt.Errorf("No existing entry matching %s", entry.Path)
	}
	entry.Weight = curWeight + 1
	return d.setEntry(entry)
}

// Complete returns sorted entries matching the given regex
func (d *DB) Complete(r *regexp.Regexp) []DBEntry {
	var entries []DBEntry

	for p, w := range d.db.GetPathWeight() {
		if r.MatchString(p) {
			entries = append(entries, DBEntry{Path: p, Weight: w})
		}
	}

	sort.Sort(sort.Reverse(byWeight(entries)))
	return entries
}

// AddCwd adds the current working directory with the given weight
func (d *DB) AddCwd(weight int64) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	if weight < 0 {
		weight = 0
	}
	return d.setEntry(DBEntry{Path: wd, Weight: weight})

}
