module github.com/tskdsb/tsk2

go 1.13

require (
	github.com/Tnze/go-mc v1.15.1 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/tskdsb/tsk2/pkg/aaa v0.0.0-00010101000000-000000000000 // indirect
	github.com/tskdsb/tsk2/pkg/bbb v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
)

replace github.com/tskdsb/tsk2/pkg/aaa => ./pkg/aaa

replace github.com/tskdsb/tsk2/pkg/bbb => ./pkg/bbb
