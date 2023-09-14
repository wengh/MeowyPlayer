package cwidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"meowyplayer.com/source/player"
)

type MusicViewList struct {
	widget.BaseWidget
	box      *fyne.Container
	makeView func(*player.Music) fyne.CanvasObject
}

func NewMusicViewList(makeMusicView func(*player.Music) fyne.CanvasObject) *MusicViewList {
	list := &MusicViewList{widget.BaseWidget{}, container.NewVBox(), makeMusicView}
	list.ExtendBaseWidget(list)
	return list
}

func (m *MusicViewList) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewScroll(m.box))
}

func (m *MusicViewList) Notify(data []player.Music) {
	m.box.RemoveAll()
	views := m.makeViews(data)
	for i := range views {
		m.box.Add(views[i])
	}
}

func (m *MusicViewList) makeViews(data []player.Music) []fyne.CanvasObject {
	views := make([]fyne.CanvasObject, len(data))
	for i := range data {
		views[i] = m.makeView(&data[i])
	}
	return views
}
