package core

import (
	"errors"
	"sync"

	"github.com/Yuzuki616/V2bX/api/panel"
	"github.com/Yuzuki616/V2bX/conf"
	"github.com/hashicorp/go-multierror"
)

type Selecter struct {
	cores []Core
	nodes sync.Map
}

func (s *Selecter) Start() error {
	for i := range s.cores {
		err := s.cores[i].Start()
		return err
	}
	return nil
}

func (s *Selecter) Close() error {
	var errs error
	for i := range s.cores {
		errs = multierror.Append(errs, s.cores[i].Close())
	}
	return errs
}

func isSupported(protocol string, protocols []string) bool {
	for i := range protocols {
		if protocol == protocols[i] {
			return true
		}
	}
	return false
}

func (s *Selecter) AddNode(tag string, info *panel.NodeInfo, config *conf.ControllerConfig) error {
	for i := range s.cores {
		if !isSupported(info.Type, s.cores[i].Protocols()) {
			continue
		}
		err := s.cores[i].AddNode(tag, info, config)
		if err != nil {
			return err
		}
		s.nodes.Store(tag, i)
	}
	return errors.New("the node type is not support")
}

func (s *Selecter) DelNode(tag string) error {
	if t, e := s.nodes.Load(tag); e {
		err := s.cores[t.(int)].DelNode(tag)
		if err != nil {
			return err
		}
		s.nodes.Delete(tag)
		return nil
	}
	return errors.New("the node is not have")
}

func (s *Selecter) AddUsers(p *AddUsersParams) (added int, err error) {
	t, e := s.nodes.Load(p.Tag)
	if !e {
		return 0, errors.New("the node is not have")
	}
	return s.cores[t.(int)].AddUsers(p)
}

func (s *Selecter) GetUserTraffic(tag, uuid string, reset bool) (up int64, down int64) {
	t, e := s.nodes.Load(tag)
	if !e {
		return 0, 0
	}
	return s.cores[t.(int)].GetUserTraffic(tag, uuid, reset)
}

func (s *Selecter) DelUsers(users []string, tag string) error {
	t, e := s.nodes.Load(tag)
	if !e {
		return errors.New("the node is not have")
	}
	return s.cores[t.(int)].DelUsers(users, tag)
}

func (s *Selecter) Protocols() []string {
	protocols := make([]string, 0)
	for i := range s.cores {
		protocols = append(protocols, s.cores[i].Protocols()...)
	}
	return protocols
}
