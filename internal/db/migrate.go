package db

import (
	"embed"
	"fmt"
	"github.com/alireza-fa/ghofle/internal/db/models"
	"github.com/alireza-fa/ghofle/pkg/rdbms"
	"github.com/alireza-fa/ghofle/pkg/utils"
	"io/fs"
	"strings"
)

//go:embed migrations
var migrations embed.FS

func Migrate(direction models.Migrate, db rdbms.RDBMS) error {
	files, err := fs.ReadDir(migrations, "migrations")
	if err != nil {
		return fmt.Errorf("error reading migrations directory:\n%s", err)
	}

	result := make([]string, 0, len(files)/2)

	for _, file := range files {
		splits := strings.Split(file.Name(), ".")
		if splits[1] == string(direction) {
			result = append(result, file.Name())
		}
	}

	result = utils.Sort(result)

	for index := 0; index < len(result); index++ {
		file := "migrations/"

		if direction == models.Up {
			file += result[index]
		} else {
			file += result[len(result)-index-1]
		}

		data, err := fs.ReadFile(migrations, file)
		if err != nil {
			return fmt.Errorf("error reading migrations file: %s\n%s", file, err)
		}

		if err := db.Execute(string(data), []interface{}{}); err != nil {
			return fmt.Errorf("error migreating the file: %s\n%s", file, err)
		}
	}

	return nil
}
