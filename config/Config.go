package config

import (
	"fmt"
	"go_database_migration/helper"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var workdir, _ = os.Getwd()

var BASEDIR = filepath.Join(workdir, "..") //jika untuk test
// var BASEDIR = filepath.Join(workdir) //jika main.go
var DataPerPage int
var DBHOST, DBPORT, DBUSER, DBPASS, DBNAME, TBNAME, TBNAME2, DIALECT, CONNECT, HOST, PORT, ADDR string
var BooksColoumnOut, BooksColoumnIn []string

func init() {
	err := godotenv.Load(filepath.Join(BASEDIR, ".env"))
	helper.PanicIFError(err)

	DataPerPage, _ = strconv.Atoi(os.Getenv("dataperpage"))
	DBHOST = os.Getenv("dbhost")
	DBPORT = os.Getenv("dbport")
	DBUSER = os.Getenv("dbuser")
	DBPASS = os.Getenv("dbpass")
	DIALECT = os.Getenv("dialect")
	DBNAME = os.Getenv("dbname")
	TBNAME = os.Getenv("tbname1")
	TBNAME2 = os.Getenv("tbname2")
	HOST = os.Getenv("host")
	PORT = os.Getenv("port")
	ADDR = fmt.Sprintf("%s:%s", HOST, PORT)
	CONNECT = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", DBUSER, DBPASS, DBHOST, DBPORT, DBNAME)
	BooksColoumnOut = []string{"id", "isbn", "title", "author", "status_borrow", "publisher", "publication_years", "description", "createdAt", "updatedAt"}
	BooksColoumnIn = []string{"isbn", "title", "author", "status_borrow", "publisher", "publication_years", "description"}
}
