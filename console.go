package rpgio

import (
	"regexp"
	"strings"

	"github.com/g3n/engine/gui"
)

type Console struct {
	*gui.Panel
	output  *gui.List
	input   *gui.Edit
	lineMax int
}

func NewConsole(width int) *Console {
	c := new(Console)
	c.Panel = gui.NewPanel(float32(width), 280)
	vbox := gui.NewVBoxLayout()
	c.Panel.SetLayout(vbox)

	c.output = gui.NewVList(float32(width), 200)
	c.Panel.Add(c.output)

	c.input = gui.NewEdit(width, "")
	c.Panel.Add(c.input)

	return c
}

func (c *Console) Log(msg string) {
	if msg != "" {
		m := gui.NewImageLabel(msg)
		c.input.SetText("")
		if c.lineMax > 0 && c.output.ItemScroller.Len() > c.lineMax {
			c.output.RemoveAt(0)
		}
		c.output.Add(m)
		c.output.ItemScroller.ScrollDown()
	}
}

func (c *Console) Cmd() (string, []string) {
	cmd := ""
	in := c.deleteExtraSpace(c.input.Text())
	ins := strings.Split(in, " ")
	if len(ins) > 0 {
		cmd = ins[0]
	}
	parm := ins[1:]
	return cmd, parm
}

func (c *Console) Input() *gui.Edit {
	return c.input
}

func (c *Console) SetLineMax(max int) {
	c.lineMax = max
}

func (c *Console) deleteExtraSpace(s string) string {
	s1 := strings.Replace(s, "	", " ", -1)
	regstr := "\\s{2,}"
	reg, _ := regexp.Compile(regstr)
	s2 := make([]byte, len(s1))
	copy(s2, s1)
	spc_index := reg.FindStringIndex(string(s2))
	for len(spc_index) > 0 {
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...)
		spc_index = reg.FindStringIndex(string(s2))
	}
	return string(s2)
}
