// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"time"

	"github.com/HappyLadySauce/component-base/pkg/scheme"
)

// ObjectMetaAccessor is an interface that provides access to the ObjectMeta of an object.
// 是用来获取对象的ObjectMeta的接口。
type ObjectMetaAccessor interface {
	// GetObjectMeta returns the ObjectMeta of the object.
	// 获取对象的ObjectMeta。
	GetObjectMeta() Object
}

// Object lets you work with object metadata from any of the versioned or
// internal API objects. Attempting to set or retrieve a field on an object that does
// not support that field (Name, UID, Namespace on lists) will be a no-op and return
// a default value.
// Object是用来操作对象的接口。
type Object interface {
	// GetID returns the ID of the object.
	// 获取对象的ID。
	GetID() uint64
	// SetID sets the ID of the object.
	// 设置对象的ID。
	SetID(id uint64)
	// GetName returns the name of the object.
	// 获取对象的名称。
	GetName() string
	// SetName sets the name of the object.
	// 设置对象的名称。
	SetName(name string)
	// GetCreatedAt returns the creation time of the object.
	// 获取对象的创建时间。
	GetCreatedAt() time.Time
	// SetCreatedAt sets the creation time of the object.
	// 设置对象的创建时间。
	SetCreatedAt(createdAt time.Time)
	// GetUpdatedAt returns the update time of the object.
	// 获取对象的更新时间。
	GetUpdatedAt() time.Time
	// SetUpdatedAt sets the update time of the object.
	// 设置对象的更新时间。
	SetUpdatedAt(updatedAt time.Time)
}

// ListInterface lets you work with list metadata from any of the versioned or
// internal API objects. Attempting to set or retrieve a field on an object that does
// not support that field will be a no-op and return a default value.
// ListInterface是用来操作列表的接口。
type ListInterface interface {
	// GetTotalCount returns the total count of the list.
	// 获取列表的总数。
	GetTotalCount() int64
	// SetTotalCount sets the total count of the list.
	// 设置列表的总数。
	SetTotalCount(count int64)
}

// Type exposes the type and APIVersion of versioned or internal API objects.
// Type是用来操作类型的接口。
type Type interface {
	// GetAPIVersion returns the API version of the object.
	// 获取对象的API版本。
	GetAPIVersion() string
	// SetAPIVersion sets the API version of the object.
	// 设置对象的API版本。
	SetAPIVersion(version string)
	// GetKind returns the kind of the object.
	// 获取对象的类型。
	GetKind() string
	// SetKind sets the kind of the object.
	// 设置对象的类型。
	SetKind(kind string)
}

// ListMeta是用来操作列表的元数据的接口。
var _ ListInterface = &ListMeta{}

// GetTotalCount returns the total count of the list.
// 获取列表的总数。
func (meta *ListMeta) GetTotalCount() int64 { return meta.TotalCount }

// SetTotalCount sets the total count of the list.
// 设置列表的总数。
func (meta *ListMeta) SetTotalCount(count int64) { meta.TotalCount = count }

// TypeMeta是用来操作类型的元数据的接口。
var _ Type = &TypeMeta{}

// GetObjectKind returns the ObjectKind of the object.
// 获取对象的ObjectKind。
func (obj *TypeMeta) GetObjectKind() scheme.ObjectKind { return obj }

// SetGroupVersionKind satisfies the ObjectKind interface for all objects that embed TypeMeta.
// 设置对象的GroupVersionKind。
func (obj *TypeMeta) SetGroupVersionKind(gvk scheme.GroupVersionKind) {
	obj.APIVersion, obj.Kind = gvk.ToAPIVersionAndKind()
}

// GroupVersionKind satisfies the ObjectKind interface for all objects that embed TypeMeta.
// 获取对象的GroupVersionKind。
func (obj *TypeMeta) GroupVersionKind() scheme.GroupVersionKind {
	return scheme.FromAPIVersionAndKind(obj.APIVersion, obj.Kind)
}

// GetAPIVersion returns the API version of the object.
// 获取对象的API版本。
func (meta *TypeMeta) GetAPIVersion() string { return meta.APIVersion }

// SetAPIVersion sets the API version of the object.
// 设置对象的API版本。
func (meta *TypeMeta) SetAPIVersion(version string) { meta.APIVersion = version }

// GetKind returns the kind of the object.
// 获取对象的类型。
func (meta *TypeMeta) GetKind() string { return meta.Kind }

// SetKind sets the kind of the object.
// 设置对象的类型。
func (meta *TypeMeta) SetKind(kind string) { meta.Kind = kind }

// GetListMeta returns the ListMeta of the object.
// 获取对象的ListMeta。
func (obj *ListMeta) GetListMeta() ListInterface { return obj }

// GetObjectMeta returns the ObjectMeta of the object.
// 获取对象的ObjectMeta。
func (obj *ObjectMeta) GetObjectMeta() Object { return obj }

// ObjectMeta is the interface that provides access to the ObjectMeta of an object.
// ObjectMeta是用来操作对象的元数据的接口。
var _ Object = &ObjectMeta{}

// GetID returns the ID of the object.
// 获取对象的ID。
func (meta *ObjectMeta) GetID() uint64 { return meta.ID }

// SetID sets the ID of the object.
// 设置对象的ID。
func (meta *ObjectMeta) SetID(id uint64) { meta.ID = id }

// GetName returns the name of the object.
// 获取对象的名称。
func (meta *ObjectMeta) GetName() string { return meta.Name }

// SetName sets the name of the object.
// 设置对象的名称。
func (meta *ObjectMeta) SetName(name string) { meta.Name = name }

// GetCreatedAt returns the creation time of the object.
// 获取对象的创建时间。
func (meta *ObjectMeta) GetCreatedAt() time.Time { return meta.CreatedAt }

// SetCreatedAt sets the creation time of the object.
// 设置对象的创建时间。
func (meta *ObjectMeta) SetCreatedAt(createdAt time.Time) { meta.CreatedAt = createdAt }

// GetUpdatedAt returns the update time of the object.
// 获取对象的更新时间。
func (meta *ObjectMeta) GetUpdatedAt() time.Time { return meta.UpdatedAt }

// SetUpdatedAt sets the update time of the object.
// 设置对象的更新时间。
func (meta *ObjectMeta) SetUpdatedAt(updatedAt time.Time) { meta.UpdatedAt = updatedAt }
