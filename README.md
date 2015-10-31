# 说明

### 使用说明

* 下载安装
```
git clone https://github.com/victorcoder/dkron.git

go get github.com/mattn/goreman

goreman start
```

* 配置
```
vim Procfile
vim config/dkron.json
```

* Example
```
import "github.com/gunsluo/dkron-api-go/job_schedule"

url := "http://127.0.0.1:8080/"
client := job_schedule.NewClient(url)

// status
statusResp, err := client.Status()

//...
```
