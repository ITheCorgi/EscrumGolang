package config

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name    string
		cfgPath string
		want    *ConfigMap
	}{
		{
			name:    "test",
			cfgPath: "../../configs/serverconfig.yml",
			want: &ConfigMap{
				Mongo: MongoConfig{
					URI:      "mongodb://localhost",
					User:     "",
					Password: "",
					Name:     "testDatabase",
				},
				HTTPData: HTTPConfig{
					Host: "127.0.0.1",
					Port: "27017",
				},
				UserAuth: AuthConfig{
					//FIXME: issue #3
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got, err := GetConfig(test.cfgPath); got != test.want && err != nil {
				t.Log(got)
				t.Errorf("The expexted value %q is not equal %q", got, test.want)
			}
		})
	}
}
