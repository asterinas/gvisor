// automatically generated by stateify.

//go:build !context_decoupling && context_decoupling
// +build !context_decoupling,context_decoupling

package systrap

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (r *subprocessRefs) StateTypeName() string {
	return "pkg/sentry/platform/systrap.subprocessRefs"
}

func (r *subprocessRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *subprocessRefs) beforeSave() {}

// +checklocksignore
func (r *subprocessRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *subprocessRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func init() {
	state.Register((*subprocessRefs)(nil))
}