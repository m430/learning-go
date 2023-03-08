package observer

// 可被观察的接口
type Observable interface {
	// 添加观察者
	addObserver(o Observer)
	// 移除观察者
	removeObserver(o Observer)
	// 通知所有观察者
	notifyAll()
}

// 被观察主体
type Subject struct {
	observers []Observer
}

func (s *Subject) addObserver(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) removeObserver(o Observer) {
	s.observers = removeFromSlice(s.observers, o)
}

func (s *Subject) notifyAll() {
	for _, o := range s.observers {
		o.update(s)
	}
}

func removeFromSlice(obs []Observer, removeObserver Observer) []Observer {
	len := len(obs)
	for i, o := range obs {
		if removeObserver.getID() == o.getID() {
			obs[len-1], obs[i] = obs[i], obs[len-1]
			return obs[:len-1]
		}
	}
	return obs
}

// 观察者接口
type Observer interface {
	update(data any)
	getID() string
}
