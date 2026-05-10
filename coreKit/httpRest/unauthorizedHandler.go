package httpRest

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

func UnauthorizedHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	resp := map[string]interface{}{
		"code": http.StatusUnauthorized,
		"msg":  "授权认证失败",
	}
	logx.WithContext(r.Context()).Errorf("UnauthorizedHandler 授权认证失败[%+v]", err)
	bs, _ := jsonx.Marshal(resp)
	if _, err := w.Write(bs); err != nil {
		logx.WithContext(r.Context()).Errorf("UnauthorizedHandler 写响应失败: %v", err)
		panic(fmt.Sprintf("UnauthorizedHandler 写响应失败: %v", err))
	}
}
