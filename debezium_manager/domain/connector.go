package domain

type Connector struct {
	Name   string            `json:"name"`
	Config DebeziumConnector `json:"config"`
}

type DebeziumConnector struct {
	Source          DatabaseHistory `json:"debezium.source.database.history"`
	ConnectorClass  string          `json:"connector.class"`
	TasksMax        string          `json:"tasks.max"`
	DatabaseServer  string          `json:"database.server.name"`
	DatabaseHost    string          `json:"database.hostname"`
	DatabasePort    string          `json:"database.port"`
	DatabaseUser    string          `json:"database.user"`
	DatabasePass    string          `json:"database.password"`
	DatabaseName    string          `json:"database.dbname"`
	TopicPrefix     string          `json:"topic.prefix"`
	TableInclude    string          `json:"table.include.list"`
	BootstrapServer string          `json:"database.history.kafka.bootstrap.servers"`
	HistoryTopic    string          `json:"database.history.kafka.topic"`
}

type DatabaseHistory string
