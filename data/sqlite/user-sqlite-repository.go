package sqlite

import (
	_ "github.com/mattn/go-sqlite3"
	"zeego/data/entities"
	"zeego/data"
	"database/sql"
	"fmt"
)

type UserSQLiteRepository struct {
	db *sql.DB
}

func NewUserSQLiteRepository(dataSource string) (*UserSQLiteRepository, data.Error) {
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		return nil, &data.RepoError{"failed to connect to database. " + err.Error(), data.Default}
	}

	if _, err = db.Exec(sqlUserTable); err != nil {
		return nil, &data.RepoError{"Failed to migrate database. " + err.Error(), data.Default}
	}

	return &UserSQLiteRepository{db}, nil
}

func (repo *UserSQLiteRepository) Close() error {
	return repo.db.Close()
}

func (repo *UserSQLiteRepository) GetByName(name string) (*entities.User, data.Error) {
	var user entities.User

	query := `SELECT id, name, passphrase, first_name, last_name, joined_at, modified_at, deleted_at
			  FROM user WHERE UPPER(name) = UPPER(?)`
	err := repo.db.QueryRow(query, name).Scan(
		&user.Id,
		&user.Name,
		&user.PassphraseHash,
		&user.FirstName,
		&user.LastName,
		&user.JoinedAt,
		&user.ModifiedAt,
		&user.DeletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, &data.RepoError{
			fmt.Sprintf(`Failed to find entity with ID "%s"`, name),
			data.NotFound,
		}
	} else if err != nil {
		return nil, data.NewRepoError(err, data.Default)
	}

	return &user, nil
}

func (repo *UserSQLiteRepository) Insert(u *entities.User) data.Error {
	if u.Name == "" || u.FirstName == "" || u.PassphraseHash == "" {
		return &data.RepoError{"Invalid entity", data.InvalidValue}
	}

	var err error
	tx, err := repo.db.Begin()
	if err != nil {
		return data.NewRepoError(err, data.Default)
	}

	stmt := `INSERT INTO user(name, passphrase, first_name) VALUES (?,?,?)`
	args := []interface{}{u.Name, u.PassphraseHash, u.FirstName}

	//hasJoinedAt := !u.JoinedAt.IsZero()
	//if hasJoinedAt {
	//	stmt = strings.Replace(stmt, ``, `joined_at, last_name`, 1)
	//	stmt = strings.Replace(stmt, `,?)`, `,?,?)`, 1)
	//}
	//
	//hasLastName := u.LastName != ""
	//if hasLastName {
	//	stmt = strings.Replace(stmt, `joined_at`, `joined_at, last_name`, 1)
	//	stmt = strings.Replace(stmt, `,?)`, `,?,?)`, 1)
	//}

	result, err := tx.Exec(stmt, args...)

	if err != nil {
		return data.NewRepoError(err, data.Default)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return data.NewRepoError(err, data.Default)
	}
	u.Id = string(id)

	if err := tx.Commit(); err != nil {
		return data.NewRepoError(err, data.Default)
	}

	return nil
}
