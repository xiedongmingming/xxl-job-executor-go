package xxl

import (
	"time"

	"github.com/go-basic/ipv4"
)

type Options struct {
	ServerAddr   string        `json:"server_addr"`   // 调度中心地址
	AccessToken  string        `json:"access_token"`  // 请求令牌 -- 到调度中心的令牌
	Timeout      time.Duration `json:"timeout"`       // 接口超时时间--到调度中心的超时时间
	ExecutorIp   string        `json:"executor_ip"`   // 本地(执行器)IP(可自行获取)
	ExecutorPort string        `json:"executor_port"` // 本地(执行器)端口
	RegistryKey  string        `json:"registry_key"`  // 执行器名称
	LogDir       string        `json:"log_dir"`       // 日志目录
	l            Logger        // 日志处理
}

func newOptions(opts ...Option) Options {

	opt := Options{
		ExecutorIp:   ipv4.LocalIP(),      // 获取当前主机IP
		ExecutorPort: DefaultExecutorPort, // 默认端口号(9999)
		RegistryKey:  DefaultRegistryKey,  // 默认注册KEY
	}

	for _, o := range opts {
		o(&opt)
	}

	if opt.l == nil {
		opt.l = &logger{}
	}

	return opt

}

type Option func(o *Options)

var (
	DefaultExecutorPort = "9999"
	DefaultRegistryKey  = "golang-jobs"
)

// ********************************************************************
// 选项配置

// ServerAddr 设置调度中心地址
func ServerAddr(addr string) Option {
	return func(o *Options) {
		o.ServerAddr = addr
	}
}

// AccessToken 请求令牌
func AccessToken(token string) Option {
	return func(o *Options) {
		o.AccessToken = token
	}
}

// ExecutorIp 设置执行器IP
func ExecutorIp(ip string) Option {
	return func(o *Options) {
		o.ExecutorIp = ip
	}
}

// ExecutorPort 设置执行器端口
func ExecutorPort(port string) Option {
	return func(o *Options) {
		o.ExecutorPort = port
	}
}

// RegistryKey 设置执行器标识
func RegistryKey(registryKey string) Option {
	return func(o *Options) {
		o.RegistryKey = registryKey
	}
}

// SetLogger 设置日志处理器
func SetLogger(l Logger) Option {
	return func(o *Options) {
		o.l = l
	}
}

// ********************************************************************
