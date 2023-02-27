package dag

import "github.com/zeromicro/go-zero/core/threading"

type (
	Starter interface {
		Start(dag *DAG)
	}
	Run interface {
		Run()
	}
	// Service is the interface that groups Start and Stop methods.
	Service interface {
		Starter
		Run
	}

	// A ServiceGroup is a group of services.
	// Attention: the starting order of the added services is not guaranteed.
	ServiceGroup struct {
		services []Service
		graph    *DAG
		//stopOnce func()
	}
)

// NewServiceGroup returns a ServiceGroup.
func NewServiceGroup() *ServiceGroup {
	sg := new(ServiceGroup)
	//sg.stopOnce = syncx.Once(sg.doStop)
	sg.graph = NewDAG()
	return sg
}

func (sg *ServiceGroup) Start() {
	sg.doStart()
}

// Register adds service into sg.
func (sg *ServiceGroup) Register(id string, service Service) {
	// push front, stop with reverse order.
	sg.services = append([]Service{service}, sg.services...)
}

func (sg *ServiceGroup) doStart() {
	routineGroup := threading.NewRoutineGroup()
	for i := range sg.services {
		service := sg.services[i]
		routineGroup.RunSafe(func() {
			service.Start(sg.graph)
		})
	}
	routineGroup.Wait()
}
