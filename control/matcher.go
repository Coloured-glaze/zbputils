package control

import (
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/extension/rate"
)

// Matcher 是 ZeroBot 匹配和处理事件的最小单元
type Matcher interface {
	// SetBlock 设置是否阻断后面的 Matcher 触发
	SetBlock(block bool) Matcher
	// Handle 直接处理事件
	Handle(handler zero.Handler)
	// Limit 限速器
	//    postfn 当请求被拒绝时的操作
	Limit(limiterfn func(*zero.Ctx) *rate.Limiter, postfn ...func(*zero.Ctx)) Matcher
}

type matcherinstance zero.Matcher

// SetBlock 设置是否阻断后面的 Matcher 触发
func (m *matcherinstance) SetBlock(block bool) Matcher {
	_ = (*zero.Matcher)(m).SetBlock(block)
	return m
}

// Handle 直接处理事件
func (m *matcherinstance) Handle(handler zero.Handler) {
	_ = (*zero.Matcher)(m).Handle(handler)
}
