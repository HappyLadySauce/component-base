// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package scheme

import (
	"fmt"
	"strings"
)

// ParseResourceArg takes the common style of string which may be either `resource.group.com` or
// `resource.version.group.com`
// and parses it out into both possibilities.  This code takes no responsibility for knowing which representation was
// intended
// but with a knowledge of all GroupVersions, calling code can take a very good guess.  If there are only two segments,
// then
// `*GroupVersionResource` is nil.
// `resource.group.com` -> `group=com, version=group, resource=resource` and `group=group.com, resource=resource`.
// ParseResourceArg是将资源参数解析为GroupVersionResource和GroupResource。
func ParseResourceArg(arg string) (*GroupVersionResource, GroupResource) {
	var gvr *GroupVersionResource
	if strings.Count(arg, ".") >= 2 {
		s := strings.SplitN(arg, ".", 3)
		gvr = &GroupVersionResource{Group: s[2], Version: s[1], Resource: s[0]}
	}

	return gvr, ParseGroupResource(arg)
}

// ParseKindArg takes the common style of string which may be either `Kind.group.com` or `Kind.version.group.com`
// and parses it out into both possibilities. This code takes no responsibility for knowing which representation was
// intended
// but with a knowledge of all GroupKinds, calling code can take a very good guess. If there are only two segments, then
// `*GroupVersionResource` is nil.
// `Kind.group.com` -> `group=com, version=group, kind=Kind` and `group=group.com, kind=Kind`.
// ParseKindArg是将Kind参数解析为GroupVersionKind和GroupKind。
func ParseKindArg(arg string) (*GroupVersionKind, GroupKind) {
	var gvk *GroupVersionKind
	if strings.Count(arg, ".") >= 2 {
		s := strings.SplitN(arg, ".", 3)
		gvk = &GroupVersionKind{Group: s[2], Version: s[1], Kind: s[0]}
	}

	return gvk, ParseGroupKind(arg)
}

// GroupResource specifies a Group and a Resource, but does not force a version.  This is useful for identifying
// concepts during lookup stages without having partially valid types.
// GroupResource是用来表示组和资源的结构体。
type GroupResource struct {
	Group    string
	Resource string
}

// WithVersion add version to GroupVersionResource.
// WithVersion是在GroupVersionResource中添加版本。
func (gr GroupResource) WithVersion(version string) GroupVersionResource {
	return GroupVersionResource{Group: gr.Group, Version: version, Resource: gr.Resource}
}

// Empty return true if Group and Resource both 0.
// Empty是用来判断Group和Resource是否为空。
func (gr GroupResource) Empty() bool {
	return len(gr.Group) == 0 && len(gr.Resource) == 0
}

// String defines the string format of GroupResource.
// String是用来定义GroupResource的字符串格式。
func (gr GroupResource) String() string {
	if len(gr.Group) == 0 {
		return gr.Resource
	}
	return gr.Resource + "." + gr.Group
}

// ParseGroupKind parse a string to GroupKind.
// ParseGroupKind是将字符串解析为GroupKind。
func ParseGroupKind(gk string) GroupKind {
	i := strings.Index(gk, ".")
	if i == -1 {
		return GroupKind{Kind: gk}
	}

	return GroupKind{Group: gk[i+1:], Kind: gk[:i]}
}

// ParseGroupResource turns "resource.group" string into a GroupResource struct.  Empty strings are allowed
// for each field.
// ParseGroupResource是将字符串解析为GroupResource。
func ParseGroupResource(gr string) GroupResource {
	if i := strings.Index(gr, "."); i >= 0 {
		return GroupResource{Group: gr[i+1:], Resource: gr[:i]}
	}
	return GroupResource{Resource: gr}
}

// GroupVersionResource unambiguously identifies a resource.  It doesn't anonymously include GroupVersion
// to avoid automatic coercion.  It doesn't use a GroupVersion to avoid custom marshaling.
// GroupVersionResource是用来表示组和版本的资源。
type GroupVersionResource struct {
	Group    string
	Version  string
	Resource string
}

// Empty returns true if GroupVersionResource is 0.
// Empty是用来判断GroupVersionResource是否为空。
func (gvr GroupVersionResource) Empty() bool {
	return len(gvr.Group) == 0 && len(gvr.Version) == 0 && len(gvr.Resource) == 0
}

// GroupResource return the group resource of GroupVersionResource.
// GroupResource是用来返回GroupVersionResource的组资源。
func (gvr GroupVersionResource) GroupResource() GroupResource {
	return GroupResource{Group: gvr.Group, Resource: gvr.Resource}
}

// GroupVersion return the version of GroupVersionResource.
// GroupVersion是用来返回GroupVersionResource的版本。
func (gvr GroupVersionResource) GroupVersion() GroupVersion {
	return GroupVersion{Group: gvr.Group, Version: gvr.Version}
}

// String defines the string format of GroupVersionResource.
// String是用来定义GroupVersionResource的字符串格式。
func (gvr GroupVersionResource) String() string {
	return strings.Join([]string{gvr.Group, "/", gvr.Version, ", Resource=", gvr.Resource}, "")
}

// GroupKind specifies a Group and a Kind, but does not force a version.  This is useful for identifying
// concepts during lookup stages without having partially valid types.
// GroupKind是用来表示组和Kind的结构体。
type GroupKind struct {
	Group string
	Kind  string
}

// Empty return true if group and kine is zero value.
// Empty是用来判断Group和Kind是否为空。
func (gk GroupKind) Empty() bool {
	return len(gk.Group) == 0 && len(gk.Kind) == 0
}

// WithVersion fill GroupKind with version.
// WithVersion是在GroupKind中添加版本。
func (gk GroupKind) WithVersion(version string) GroupVersionKind {
	return GroupVersionKind{Group: gk.Group, Version: version, Kind: gk.Kind}
}

// String defines the string format of GroupKind.
// String是用来定义GroupKind的字符串格式。
func (gk GroupKind) String() string {
	if len(gk.Group) == 0 {
		return gk.Kind
	}
	return gk.Kind + "." + gk.Group
}

// GroupVersionKind unambiguously identifies a kind.  It doesn't anonymously include GroupVersion
// to avoid automatic coercion.  It doesn't use a GroupVersion to avoid custom marshaling.
// GroupVersionKind是用来表示组和版本的Kind。
type GroupVersionKind struct {
	Group   string
	Version string
	Kind    string
}

// Empty returns true if group, version, and kind are empty.
// Empty是用来判断Group、Version和Kind是否为空。
func (gvk GroupVersionKind) Empty() bool {
	return len(gvk.Group) == 0 && len(gvk.Version) == 0 && len(gvk.Kind) == 0
}

// GroupKind returns the group kind.
// GroupKind是用来返回GroupVersionKind的组和Kind。
func (gvk GroupVersionKind) GroupKind() GroupKind {
	return GroupKind{Group: gvk.Group, Kind: gvk.Kind}
}

// GroupVersion returns the group version.
// GroupVersion是用来返回GroupVersionKind的组和版本。
func (gvk GroupVersionKind) GroupVersion() GroupVersion {
	return GroupVersion{Group: gvk.Group, Version: gvk.Version}
}

// String returns the string format of GroupVersionKind.
// String是用来定义GroupVersionKind的字符串格式。
func (gvk GroupVersionKind) String() string {
	return gvk.Group + "/" + gvk.Version + ", Kind=" + gvk.Kind
}

// GroupVersion contains the "group" and the "version", which uniquely identifies the API.
// GroupVersion是用来表示组和版本的结构体。
type GroupVersion struct {
	Group   string
	Version string
}

// Empty returns true if group and version are empty.
// Empty是用来判断Group和Version是否为空。
func (gv GroupVersion) Empty() bool {
	return len(gv.Group) == 0 && len(gv.Version) == 0
}

// String puts "group" and "version" into a single "group/version" string. For the legacy v1
// it returns "v1".
// String是用来定义GroupVersion的字符串格式。
func (gv GroupVersion) String() string {
	if len(gv.Group) > 0 {
		return gv.Group + "/" + gv.Version
	}
	return gv.Version
}

// Identifier implements runtime.GroupVersioner interface.
// Identifier是用来实现GroupVersioner接口。
func (gv GroupVersion) Identifier() string {
	return gv.String()
}

// KindForGroupVersionKinds identifies the preferred GroupVersionKind out of a list. It returns ok false
// if none of the options match the group. It prefers a match to group and version over just group.
// TODO: Move GroupVersion to a package under pkg/runtime, since it's used by scheme.
// TODO: Introduce an adapter type between GroupVersion and runtime.GroupVersioner, and use LegacyCodec(GroupVersion)
//   in fewer places.
// KindForGroupVersionKinds是用来识别GroupVersionKind。
func (gv GroupVersion) KindForGroupVersionKinds(kinds []GroupVersionKind) (target GroupVersionKind, ok bool) {
	for _, gvk := range kinds {
		if gvk.Group == gv.Group && gvk.Version == gv.Version {
			return gvk, true
		}
	}
	for _, gvk := range kinds {
		if gvk.Group == gv.Group {
			return gv.WithKind(gvk.Kind), true
		}
	}
	return GroupVersionKind{}, false
}

// ParseGroupVersion turns "group/version" string into a GroupVersion struct. It reports error
// if it cannot parse the string.
// ParseGroupVersion是将字符串解析为GroupVersion。
func ParseGroupVersion(gv string) (GroupVersion, error) {
	// this can be the internal version for the legacy apimachinery types
	// TODO once we've cleared the last uses as strings, this special case should be removed.
	if (len(gv) == 0) || (gv == "/") {
		return GroupVersion{}, nil
	}

	switch strings.Count(gv, "/") {
	case 0:
		return GroupVersion{"", gv}, nil
	case 1:
		i := strings.Index(gv, "/")
		return GroupVersion{gv[:i], gv[i+1:]}, nil
	default:
		return GroupVersion{}, fmt.Errorf("unexpected GroupVersion string: %v", gv)
	}
}

// WithKind creates a GroupVersionKind based on the method receiver's GroupVersion and the passed Kind.
// WithKind是用来创建GroupVersionKind。
func (gv GroupVersion) WithKind(kind string) GroupVersionKind {
	return GroupVersionKind{Group: gv.Group, Version: gv.Version, Kind: kind}
}

// WithResource creates a GroupVersionResource based on the method receiver's GroupVersion and the passed Resource.
// WithResource是用来创建GroupVersionResource。
func (gv GroupVersion) WithResource(resource string) GroupVersionResource {
	return GroupVersionResource{Group: gv.Group, Version: gv.Version, Resource: resource}
}

// GroupVersions can be used to represent a set of desired group versions.
// TODO: Move GroupVersions to a package under pkg/runtime, since it's used by scheme.
// TODO: Introduce an adapter type between GroupVersions and runtime.GroupVersioner, and use LegacyCodec(GroupVersion)
//   in fewer places.
// GroupVersions是用来表示组和版本的结构体。
type GroupVersions []GroupVersion

// Identifier implements runtime.GroupVersioner interface.
// Identifier是用来实现GroupVersioner接口。
func (gvs GroupVersions) Identifier() string {
	groupVersions := make([]string, 0, len(gvs))
	for i := range gvs {
		groupVersions = append(groupVersions, gvs[i].String())
	}
	return fmt.Sprintf("[%s]", strings.Join(groupVersions, ","))
}

// KindForGroupVersionKinds identifies the preferred GroupVersionKind out of a list. It returns ok false
// if none of the options match the group.
// KindForGroupVersionKinds是用来识别GroupVersionKind。
func (gvs GroupVersions) KindForGroupVersionKinds(kinds []GroupVersionKind) (GroupVersionKind, bool) {
	// nolint:prealloc //no need
	var targets []GroupVersionKind
	for _, gv := range gvs {
		target, ok := gv.KindForGroupVersionKinds(kinds)
		if !ok {
			continue
		}
		targets = append(targets, target)
	}
	if len(targets) == 1 {
		return targets[0], true
	}
	if len(targets) > 1 {
		return bestMatch(kinds, targets), true
	}
	return GroupVersionKind{}, false
}

// bestMatch tries to pick best matching GroupVersionKind and falls back to the first
// found if no exact match exists.
// bestMatch是用来选择最佳的GroupVersionKind。
func bestMatch(kinds []GroupVersionKind, targets []GroupVersionKind) GroupVersionKind {
	for _, gvk := range targets {
		for _, k := range kinds {
			if k == gvk {
				return k
			}
		}
	}
	return targets[0]
}

// ToAPIVersionAndKind is a convenience method for satisfying runtime.Object on types that
// do not use TypeMeta.
// ToAPIVersionAndKind是用来转换GroupVersionKind为字符串。
func (gvk GroupVersionKind) ToAPIVersionAndKind() (string, string) {
	if gvk.Empty() {
		return "", ""
	}
	return gvk.GroupVersion().String(), gvk.Kind
}

// FromAPIVersionAndKind returns a GVK representing the provided fields for types that
// do not use TypeMeta. This method exists to support test types and legacy serializations
// that have a distinct group and kind.
// TODO: further reduce usage of this method.
// FromAPIVersionAndKind是将字符串转换为GroupVersionKind。
func FromAPIVersionAndKind(apiVersion, kind string) GroupVersionKind {
	if gv, err := ParseGroupVersion(apiVersion); err == nil {
		return GroupVersionKind{Group: gv.Group, Version: gv.Version, Kind: kind}
	}
	return GroupVersionKind{Kind: kind}
}
