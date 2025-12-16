// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package scheme defines some useful function for group version.
package scheme

// ObjectKind is used by serialization to set type information from the Scheme onto the serialized version of an object.
// For objects that cannot be serialized or have unique requirements, this interface may be a no-op.
// ObjectKind是用来设置对象的类型信息。
type ObjectKind interface {
	// SetGroupVersionKind sets or clears the intended serialized kind of an object. Passing kind nil
	// should clear the current setting.
	SetGroupVersionKind(kind GroupVersionKind)
	// GroupVersionKind returns the stored group, version, and kind of an object, or nil if the object does
	// not expose or provide these fields.
	GroupVersionKind() GroupVersionKind
}

// EmptyObjectKind implements the ObjectKind interface as a noop.
// EmptyObjectKind是用来实现ObjectKind接口。
var EmptyObjectKind = emptyObjectKind{}

// emptyObjectKind是用来实现ObjectKind接口。
type emptyObjectKind struct{}


// SetGroupVersionKind implements the ObjectKind interface.
// SetGroupVersionKind是用来设置对象的类型信息。
func (emptyObjectKind) SetGroupVersionKind(gvk GroupVersionKind) {}

// GroupVersionKind implements the ObjectKind interface.
// GroupVersionKind是用来返回对象的类型信息。
func (emptyObjectKind) GroupVersionKind() GroupVersionKind { return GroupVersionKind{} }
