package repository_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/piigyy/dot-go/internal/model"
	"github.com/piigyy/dot-go/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryPGSQLTestSuite struct {
	suite.Suite
	db      *sql.DB
	sqlMock sqlmock.Sqlmock
}

func TestUserRepositoryPGSQL(t *testing.T) {
	suite.Run(t, new(UserRepositoryPGSQLTestSuite))
}

func (suite *UserRepositoryPGSQLTestSuite) TestCreateUser() {
	testCases := []struct {
		name           string
		arg            model.User
		doMock         func()
		expectedResult int
		expectedErr    error
	}{
		{
			name: "success create user to db",
			arg: model.User{
				Fullname:  "John Doe",
				Email:     "johndoe@mail.com",
				BirthDate: time.Date(1999, 12, 19, 0, 0, 0, 0, time.UTC),
			},
			doMock: func() {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1)
				suite.
					sqlMock.
					ExpectQuery(regexp.QuoteMeta(repository.QueryCreateUser)).
					WithArgs("John Doe", "johndoe@mail.com", time.Date(1999, 12, 19, 0, 0, 0, 0, time.UTC).Format("2006-01-02")).
					WillReturnRows(rows)
			},
			expectedResult: 1,
			expectedErr:    nil,
		},
		{
			name: "db connection done",
			arg: model.User{
				Fullname:  "John Doe",
				Email:     "johndoe@mail.com",
				BirthDate: time.Date(1999, 12, 19, 0, 0, 0, 0, time.UTC),
			},
			doMock: func() {
				suite.
					sqlMock.
					ExpectQuery(regexp.QuoteMeta(repository.QueryCreateUser)).
					WithArgs("John Doe", "johndoe@mail.com", time.Date(1999, 12, 19, 0, 0, 0, 0, time.UTC).Format("2006-01-02")).
					WillReturnError(sql.ErrConnDone)
			},
			expectedResult: 0,
			expectedErr:    sql.ErrConnDone,
		},
	}

	t := suite.T()
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.doMock()
			ctx := context.Background()
			assertion := assert.New(t)

			repository := repository.NewUserRepositoryPGSQL(suite.db)
			actualID, actualErr := repository.CreateUser(ctx, tt.arg)
			assertion.Equal(tt.expectedResult, actualID)
			assertion.Equal(tt.expectedErr, actualErr)
		})
	}
}

func (suite *UserRepositoryPGSQLTestSuite) SetupTest() {
	db, sqlMock, err := sqlmock.New()
	assert.NoError(suite.T(), err)

	suite.db = db
	suite.sqlMock = sqlMock
}
