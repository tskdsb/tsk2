module github.com/tskdsb/tsk2

go 1.13

require (
	github.com/tskdsb/tsk2/pkg/aaa v0.0.0-00010101000000-000000000000
	github.com/tskdsb/tsk2/pkg/bbb v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.2.7 // indirect
	k8s.io/client-go v11.0.0+incompatible
)

replace github.com/tskdsb/tsk2/pkg/aaa => ./pkg/aaa

replace github.com/tskdsb/tsk2/pkg/bbb => ./pkg/bbb
