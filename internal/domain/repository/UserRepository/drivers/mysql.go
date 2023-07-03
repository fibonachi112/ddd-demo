package drivers

import (
	"context"
	"database/sql"
	"fmt"
	"repo/config"
	"repo/internal/DI"
	"repo/internal/domain/entity"
)

const STMT_GET_ONE = `insert * from users where id = ?`
const STMT_SAVE_ONE = `insert into users (age, name, surname, email, password) values(?,?,?,?,?)`
const STMT_UPDATE_ONE = `update repo.users set age = ? , name= ?, surname = ?, email = ?, password = ?, updated_at = ?  where id = ?`
const STMT_DELETE_ONE = `delete from users where id = ?`

type Statements struct {
	GetOneStmt    *sql.Stmt
	SaveOneStmt   *sql.Stmt
	UpdateOneStmt *sql.Stmt
	DeleteOneStmt *sql.Stmt
}

type MysqlUserRepository struct {
	Config     *config.Configuration
	Di         *DI.ServiceContainer
	Statements *Statements
	Ctx        *context.Context
}

func (m *MysqlUserRepository) initQueries() {
	m.Statements = &Statements{
		GetOneStmt:    prepareStatement(m.Di, STMT_GET_ONE),
		SaveOneStmt:   prepareStatement(m.Di, STMT_SAVE_ONE),
		UpdateOneStmt: prepareStatement(m.Di, STMT_UPDATE_ONE),
		DeleteOneStmt: prepareStatement(m.Di, STMT_DELETE_ONE),
	}
}

func (m *MysqlUserRepository) GetOne(id int) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MysqlUserRepository) SaveOne(user *entity.User) error {
	//TODO implement me

	result, err := m.Statements.SaveOneStmt.Exec(user.Age, user.Name, user.Surname, user.Email, user.Password)
	if err != nil {
		return err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = int(lastInsertId)
	return nil
}

func (m *MysqlUserRepository) UpdateOne(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (m *MysqlUserRepository) DeleteOne(user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (m *MysqlUserRepository) SetContext(ctx *context.Context) {
	m.Ctx = ctx
}

func prepareStatement(di *DI.ServiceContainer, sql string) *sql.Stmt {
	stmt, err := di.MysqlDatabase.Prepare(sql)
	if err != nil {
		panic(fmt.Sprintf("error occured while prepare statement : '%s'", sql))
	}
	return stmt
}

func MakeMysqlDriver(config *config.Configuration, di *DI.ServiceContainer) *MysqlUserRepository {
	return &MysqlUserRepository{
		Config: config,
		Di:     di,
	}
}
