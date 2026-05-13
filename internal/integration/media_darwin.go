//go:build darwin

package integration

import tea "github.com/charmbracelet/bubbletea"

type Instance struct{}

func Init(p *tea.Program) *Instance {

	return nil
}

func (ins *Instance) GetStatus() string {
	return "Stopped"
}

func (ins *Instance) UpdateStatus(status string)    {}
func (ins *Instance) UpdatePosition(position int64) {}
func (ins *Instance) UpdateMetadata(meta Metadata)  {}
func (ins *Instance) ClearMetadata()                {}
func (ins *Instance) Close()                        {}
