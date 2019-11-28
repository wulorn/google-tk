# google-tk
Google TK 参数生成及demo使用
# Usage

```
import "github.com/wulorn/google-tk"
```
谷歌翻译API的tk值通常不会改变，可以通过请求解析并缓存tk值，然后根据tk值调用API

```
# 获取谷歌翻译tk值
tkk, _ := tk.GetTKK()
# 根据tk值解析出请求需要携带的tk值
tk := tk.GetTK("queryString", tkk)

# 直接调用api接口 返回 []byte, err
buf, err := tk.Translate("queryString")
```