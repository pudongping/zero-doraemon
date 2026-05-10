package httpRest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 避免请求网站图标出现 404
		if strings.HasPrefix(r.URL.Path, "/favicon.ico") {
			w.WriteHeader(http.StatusOK)
			return
		}

		if strings.Contains(r.Header.Get("Accept"), "text/html") {
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte("页面无法找到 ...(｡•ˇ‸ˇ•｡) ...")); err != nil {
				logx.WithContext(r.Context()).Errorf("NotFoundHandler 写响应失败: %v", err)
				panic(fmt.Sprintf("NotFoundHandler 写响应失败: %v", err))
			}
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)

		resp := map[string]interface{}{
			"code": http.StatusNotFound,
			"msg":  fmt.Sprintf("路由[%s]未定义", r.URL.Path),
		}
		bs, _ := jsonx.Marshal(resp)
		if _, err := w.Write(bs); err != nil {
			logx.WithContext(r.Context()).Errorf("NotFoundHandler 写响应失败: %v", err)
			panic(fmt.Sprintf("NotFoundHandler 写响应失败: %v", err))
		}
	})
}
