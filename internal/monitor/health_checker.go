package handler

type HealthTest struct {

}

type HealthMonitor interface {
	Ping() error
}

type Health struct {
	monitor map[string]HealthMonitor
}

type HealthResponse struct {
	Status [] Status `json:"status"`
}

type Status struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

func (h Health) check() (HealthResponse, bool ,error) {

	resp := HealthResponse{}
	isOk := true
	for name, impl := range h.monitor {
		if err := impl.Ping(); err != nil {
			isOk = false
			resp.Status = append(resp.Status, Status{Name: name, Value:"DOWN"})
		} else {
			resp.Status = append(resp.Status, Status{Name: name, Value:"UP"})
		}
	}
	return resp, isOk, nil
}