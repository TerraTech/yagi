// generated by yagi. Don't modify this file!
// Any changes will be lost if this file is regenerated.

package wrapper

import "github.com/hneemann/yagi/example/wrapper/largecode"

type WrapperInt64 struct{ delegate largecode.List }

func (l *WrapperInt64) Add(item int64) {
	l.delegate.Add(item)
}

func (l *WrapperInt64) Get(index int) int64 {
	item, ok := l.delegate.Get(index).(int64)
	if ok {
		return item
	}
	panic("wrong type in list")
}

func (l *WrapperInt64) Remove(index int) {
	l.delegate.Remove(index)
}

func (l *WrapperInt64) Len() int {
	return l.delegate.Len()
}

type WrapperString struct{ delegate largecode.List }

func (l *WrapperString) Add(item string) {
	l.delegate.Add(item)
}

func (l *WrapperString) Get(index int) string {
	item, ok := l.delegate.Get(index).(string)
	if ok {
		return item
	}
	panic("wrong type in list")
}

func (l *WrapperString) Remove(index int) {
	l.delegate.Remove(index)
}

func (l *WrapperString) Len() int {
	return l.delegate.Len()
}
