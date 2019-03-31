// Copyright 2019 Kos.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"github.com/manifoldco/promptui"
	"os/exec"
	"strings"
)

//控制台項目
type Control struct {
	Name     string
	Cmd      string
	CmdStyle int
}

func getControls() []Control {
	controls := []Control{
		{Name: "地區及語言選項", Cmd: "intl.cpl", CmdStyle: 1},
		{Name: "日期和時間", Cmd: "timedate.cpl", CmdStyle: 1},
		{Name: "排定的工作", Cmd: "schedtasks", CmdStyle: 1},
		{Name: "系統", Cmd: "sysdm.cpl", CmdStyle: 1},
		{Name: "裝置管理員", Cmd: "hdwwiz.cpl", CmdStyle: 1},
		{Name: "電源選項", Cmd: "powercfg.cpl", CmdStyle: 1},
		{Name: "新增/移除程式", Cmd: "appwiz.cpl", CmdStyle: 1},
		{Name: "顯示", Cmd: "desk.cpl", CmdStyle: 1},
		{Name: "資料夾選項", Cmd: "folders", CmdStyle: 1},
		{Name: "使用者帳戶", Cmd: "nusrmgr.cpl", CmdStyle: 1},
		{Name: "系統管理工具", Cmd: "admintools", CmdStyle: 1},
		{Name: "電腦管理", Cmd: "compmgmt.msc", CmdStyle: 2},
		{Name: "元件服務", Cmd: "dcomcnfg", CmdStyle: 1},
		{Name: "事件檢視器", Cmd: "eventvwr.msc", CmdStyle: 2},
		{Name: "服務", Cmd: "services.msc", CmdStyle: 2},
		{Name: "效能 - perfmon", Cmd: "perfmon.msc", CmdStyle: 2},
		{Name: "自動更新", Cmd: "wuaucpl.cpl", CmdStyle: 1},
		{Name: "Windows 防火牆", Cmd: "firewall.cpl", CmdStyle: 1},
		{Name: "網路連線", Cmd: "ncpa.cpl", CmdStyle: 1},
		{Name: "本機使用者和群組", Cmd: "lusrmgr.msc", CmdStyle: 2},
		{Name: "登錄編輯程式", Cmd: "regedit", CmdStyle: 2},
		{Name: "離開命令列", Cmd: "quite", CmdStyle: 3},
	}
	return controls
}

func main() {
	controls := getControls()
	funcMap := promptui.FuncMap
	funcMap["callCmdName"] = func(cmdStyle int, cmdName string) string {
		if cmdStyle == 1 {
			return "control " + cmdName
		} else if cmdStyle == 2 {
			return cmdName
		}
		return ""
	}
	templates := promptui.SelectTemplates{
		Active:   `🍕 {{ .Name | cyan | bold }}`,
		Inactive: `   {{ .Name | cyan }}`,
		Selected: `{{ "✔" | green | bold }} {{ .Name | cyan }}`,
		Details: `命令:
		{{ callCmdName .CmdStyle .Cmd }}`,
	}
	list := promptui.Select{
		Label:     "控制台",
		Items:     controls,
		Templates: &templates,
		Searcher: func(input string, idx int) bool {
			control := controls[idx]
			if strings.Contains(control.Name, input) {
				return true
			}
			if strings.Contains(control.Cmd, input) {
				return true
			}
			return false
		},
	}
	i, _, err := list.Run()
	if err != nil {
		//掛了
		fmt.Printf("錯誤選項 %v\n", err)
		return
	}
	if controls[i].CmdStyle == 3 {
		return
	}
	var out bytes.Buffer
	var stderr bytes.Buffer
	var c *exec.Cmd
	if controls[i].CmdStyle == 1 {
		c = exec.Command("cmd", "/C", "control", controls[i].Cmd)
	} else {
		c = exec.Command("cmd", "/C", controls[i].Cmd)
	}
	c.Stdout = &out
	c.Stderr = &stderr
	if err := c.Run(); err != nil && stderr.String() != "" {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
}
