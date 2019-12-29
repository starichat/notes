package credit

import (
	"github.com/starichat/notes/DesignPatter/credit_sys/config"
	"github.com/starichat/notes/DesignPatter/credit_sys/dao"
	"testing"
)

func Test_addCredit(t *testing.T) {
	config.Init()
	db := dao.InitDB()
	db.AutoMigrate(
		&Credit{})
	defer db.Close()

	addCredit()
}
