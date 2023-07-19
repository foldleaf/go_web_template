package snowflake

// 雪花算法生成 id
import (
	"fmt"
	"time"

	sf "github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
)

var node *sf.Node

func Init() (err error) {
	var st time.Time
	// 格式化 1月2号下午3时4分5秒  2006年
	st, err = time.Parse("2006-01-02", viper.GetString("app.start_time"))
	if err != nil {
		fmt.Println(err)
		return
	}

	sf.Epoch = st.UnixNano() / 1e6
	node, err = sf.NewNode(viper.GetInt64("app.machine_id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

// 生成 64 位的 雪花 ID
func GenID() int64 {
	return node.Generate().Int64()
}
