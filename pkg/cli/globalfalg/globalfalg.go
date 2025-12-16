// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package globalflag

import (
	"flag"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

// AddGlobalFlags explicitly registers flags that libraries (log, verflag, etc.) register
// against the global flagsets from "flag".
// We do this in order to prevent unwanted flags from leaking into the component's flagset.
// AddGlobalFlags 显式地注册库（log、verflag 等）注册到 "flag" 的全局标志集中的标志。
// 我们这样做是为了防止不需要的标志泄漏到组件的标志集中。
func AddGlobalFlags(fs *pflag.FlagSet, name string) {
	fs.BoolP("help", "h", false, fmt.Sprintf("帮助 for %s", name))
	fs.BoolP("help", "h", false, fmt.Sprintf("help for %s", name))
}

// normalize replaces underscores with hyphens
// we should always use hyphens instead of underscores when registering component flags.
// normalize 将下划线替换为连字符
// 我们应该始终使用连字符而不是下划线来注册组件标志。
func normalize(s string) string {
	return strings.ReplaceAll(s, "_", "-")
}

// Register adds a flag to local that targets the Value associated with the Flag named globalName in flag.CommandLine.
// Register 将一个标志添加到 local，其目标值与 flag.CommandLine 中的名为 globalName 的 Flag 相关联。
func Register(local *pflag.FlagSet, globalName string) {
	if f := flag.CommandLine.Lookup(globalName); f != nil {
		pflagFlag := pflag.PFlagFromGoFlag(f)
		pflagFlag.Name = normalize(pflagFlag.Name)
		local.AddFlag(pflagFlag)
	} else {
		panic(fmt.Sprintf("failed to find flag in global flagset (flag): %s", globalName))
	}
}