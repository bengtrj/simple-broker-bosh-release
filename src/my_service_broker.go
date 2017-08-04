package main

import (
	"context"

	"github.com/pivotal-cf/brokerapi"
)

type FakeServiceBroker struct {
	ProvisionDetails   brokerapi.ProvisionDetails
	UpdateDetails      brokerapi.UpdateDetails
	DeprovisionDetails brokerapi.DeprovisionDetails

	ProvisionedInstanceIDs   []string
	DeprovisionedInstanceIDs []string
	UpdatedInstanceIDs       []string

	BoundInstanceIDs    []string
	BoundBindingIDs     []string
	BoundBindingDetails brokerapi.BindDetails
	VolumeMounts        []brokerapi.VolumeMount

	UnbindingDetails brokerapi.UnbindDetails

	InstanceLimit int

	ProvisionError     error
	BindError          error
	UnbindError        error
	DeprovisionError   error
	LastOperationError error
	UpdateError        error

	ShouldReturnAsync     bool
	DashboardURL          string
	OperationDataToReturn string

	LastOperationInstanceID string
	LastOperationData       string

	ReceivedContext bool
}

type FakeAsyncServiceBroker struct {
	FakeServiceBroker
	ShouldProvisionAsync bool
}

type FakeAsyncOnlyServiceBroker struct {
	FakeServiceBroker
}

func (fakeBroker *FakeServiceBroker) Services(context context.Context) []brokerapi.Service {
	return []brokerapi.Service{
		{
			ID:            "0000001111",
			Name:          "morengt_broker",
			Description:   "Nah",
			Bindable:      true,
			PlanUpdatable: true,
			Plans: []brokerapi.ServicePlan{
				{
					ID:          "111111",
					Name:        "morengt_broker_plan",
					Description: "The default plan",
					Metadata: &brokerapi.ServicePlanMetadata{
						Bullets:     []string{},
						DisplayName: "morengt",
					},
				},
			},
			Metadata: &brokerapi.ServiceMetadata{
				DisplayName:      "morengt",
				LongDescription:  "Long description",
				DocumentationUrl: "http://thedocs.com",
				SupportUrl:       "http://helpme.no",
			},
			Tags: []string{
				"pivotal",
				"morengt",
			},
		},
	}
}

func (fakeBroker *FakeServiceBroker) Provision(context context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	return brokerapi.ProvisionedServiceSpec{DashboardURL: "http://NO-OP"}, nil
}

func (fakeBroker *FakeServiceBroker) Update(context context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{IsAsync: false, OperationData: "Operation Data"}, nil
}

func (fakeBroker *FakeServiceBroker) Deprovision(context context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	return brokerapi.DeprovisionServiceSpec{OperationData: "Operation Data", IsAsync: false}, nil
}

func (fakeBroker *FakeServiceBroker) Bind(context context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	result := brokerapi.Binding{
		Credentials: Credentials{
			Host:     "127.0.0.1",
			Port:     3000,
			Username: "batman",
			Password: "robin",
		},
		SyslogDrainURL:  "fakeBroker.SyslogDrainURL",
		RouteServiceURL: "fakeBroker.RouteServiceURL",
		VolumeMounts:    []brokerapi.VolumeMount{{}},
	}

	// res2B, _ := json.Marshal(result)
	// fmt.Println(string(res2B))

	return result, nil
}

func (fakeBroker *FakeServiceBroker) Unbind(context context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	return nil
}

func (fakeBroker *FakeServiceBroker) LastOperation(context context.Context, instanceID, operationData string) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{State: "succeeded", Description: "Nah"}, nil
}

type Credentials struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}
