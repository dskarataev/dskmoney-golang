package migrations

import (
	"fmt"
	"sort"
	"strconv"
)

var (
	theMigrations []Migration
)

type Migration struct {
	Version int64
	Comment string
	Up      func(DB) error
	Down    func(DB) error
}

func (m *Migration) String() string {
	return strconv.FormatInt(m.Version, 10)
}

// Register registers new database migration
// that is already versioned properly
func Register(version int64, comment string, up, down func(DB) error) {
	theMigrations = append(theMigrations, Migration{
		Version: version,
		Comment: comment,
		Up:      up,
		Down:    down,
	})
}

// MigrateApp runs "init" and "up" commands on the app migration table
func MigrateApp(db DB, name string) (appName string, oldVersion, newVersion int64, err error) {
	// to add info about which app we migrated
	appName = name

	SetAppTableName(name)

	_, _, err = Run(db, "init")
	if err != nil {
		return
	}

	oldVersion , newVersion, err = Run(db, "up")

	return
}

// Run runs command on the db. Supported commands are:
// - init - creates gopg_migrations table.
// - up - runs all available migrations.
// - down - reverts last migration.
// - version - prints current db version.
// - set_version - sets db version without running migrations.
func Run(db DB, a ...string) (oldVersion, newVersion int64, err error) {
	// Make a copy so there are no side effects of sorting.
	migrations := make([]Migration, len(theMigrations))
	copy(migrations, theMigrations)
	theMigrations = theMigrations[:0]
	return RunMigrations(db, migrations, a...)
}

// RunMigrations is like Run, but accepts list of migrations.
func RunMigrations(db DB, migrations []Migration, a ...string) (oldVersion, newVersion int64, err error) {
	sortMigrations(migrations)

	var cmd string
	if len(a) > 0 {
		cmd = a[0]
	}

	if cmd == "init" {
		err = createTables(db)
		cmd = "version"
	}

	oldVersion, err = Version(db)
	if err != nil {
		return
	}
	newVersion = oldVersion

	switch cmd {
	case "version":
		return
	case "up", "":
		for i := range migrations {
			m := &migrations[i]
			if m.Version <= oldVersion {
				continue
			}
			err = m.Up(db)
			if err != nil {
				fmt.Printf("cannot migrate UP: %s", m.Comment)
				return
			}
			newVersion = m.Version

			err = SetVersion(db, newVersion, m.Comment + " UP")
			if err != nil {
				return
			}
		}
		return
	case "down":
		if oldVersion == 0 {
			return
		}

		var m *Migration
		for i := len(migrations) - 1; i >= 0; i-- {
			mm := &migrations[i]
			if mm.Version <= oldVersion {
				m = mm
				break
			}
		}
		if m == nil {
			err = fmt.Errorf("migration %d not found\n", oldVersion)
			return
		}

		if m.Down != nil {
			err = m.Down(db)
			if err != nil {
				fmt.Printf("cannot migrate DOWN: %s", m.Comment)
				return
			}
		}

		newVersion = m.Version - 1
		err = SetVersion(db, newVersion, m.Comment + " DOWN")
		if err != nil {
			return
		}
		return
	case "set_version":
		if len(a) < 2 {
			err = fmt.Errorf("set_version requires version as 2nd arg, e.g. set_version 42")
			return
		}

		newVersion, err = strconv.ParseInt(a[1], 10, 64)
		if err != nil {
			return
		}
		err = SetVersion(db, newVersion, "version is set manually")
		return
	default:
		err = fmt.Errorf("unsupported command: %q", cmd)
		return
	}
}

type migrationSorter []Migration

func (ms migrationSorter) Len() int {
	return len(ms)
}

func (ms migrationSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func (ms migrationSorter) Less(i, j int) bool {
	return ms[i].Version < ms[j].Version
}

func sortMigrations(migrations []Migration) {
	ms := migrationSorter(migrations)
	sort.Sort(ms)
}
