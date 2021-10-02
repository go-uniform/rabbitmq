package service

import (
	"service/service/_base"
	"service/service/actions"
	"service/service/commands"
	"service/service/events"
	"service/service/hooks"
	"service/service/info"
)

var MustAsset = info.MustAsset

// load all actions, commands, events and hooks
func init() {
	actions.Load(args, info.MustAsset)
	commands.Load(args, info.MustAsset)
	events.Load(args, info.MustAsset)
	hooks.Load(args, info.MustAsset)
	_base.AppClient = info.AppClient
	_base.AppProject = info.AppProject
	_base.AppService = info.AppService
	_base.Database = info.Database
}

// wrap base const to avoid circular reference
const (
	AppName        = info.AppName
	AppDescription = info.AppDescription
	AppVersion     = info.AppVersion
	AppCommit      = info.AppCommit
)

// wrap base type to avoid circular reference
type M _base.M
