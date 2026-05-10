package env

type Environment string

const (
	DevMode  Environment = "dev"  // 开发环境
	TestMode Environment = "test" // 测试环境
	PreMode  Environment = "pre"  // 预发环境
	ProMode  Environment = "pro"  // 正式环境，生产环境
)

var currentMode Environment

// Setup 初始化当前环境模式，通常在入口文件中加载完配置文件后调用
func Setup(mode string) {
	currentMode = Environment(mode)
}

// CurrentMode 获取当前运行环境
func CurrentMode() Environment {
	return currentMode
}

// String 获取当前运行环境的字符串表示
func (e Environment) String() string {
	return string(e)
}

// IsDev 判断是否为开发环境 (dev)
func IsDev() bool {
	return CurrentMode() == DevMode
}

// IsTest 判断是否为测试环境 (test)
func IsTest() bool {
	return CurrentMode() == TestMode
}

// IsPre 判断是否为预发环境 (pre)
func IsPre() bool {
	return CurrentMode() == PreMode
}

// IsPro 判断是否为正式环境 (pro)
func IsPro() bool {
	return CurrentMode() == ProMode
}
