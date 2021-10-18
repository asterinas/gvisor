// automatically generated by stateify.

package pipe

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (i *inodeOperations) StateTypeName() string {
	return "pkg/sentry/kernel/pipe.inodeOperations"
}

func (i *inodeOperations) StateFields() []string {
	return []string{
		"InodeSimpleAttributes",
		"p",
	}
}

func (i *inodeOperations) beforeSave() {}

// +checklocksignore
func (i *inodeOperations) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.InodeSimpleAttributes)
	stateSinkObject.Save(1, &i.p)
}

func (i *inodeOperations) afterLoad() {}

// +checklocksignore
func (i *inodeOperations) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.InodeSimpleAttributes)
	stateSourceObject.Load(1, &i.p)
}

func (p *Pipe) StateTypeName() string {
	return "pkg/sentry/kernel/pipe.Pipe"
}

func (p *Pipe) StateFields() []string {
	return []string{
		"Queue",
		"isNamed",
		"readers",
		"writers",
		"buf",
		"off",
		"size",
		"max",
		"hadWriter",
	}
}

func (p *Pipe) beforeSave() {}

// +checklocksignore
func (p *Pipe) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.Queue)
	stateSinkObject.Save(1, &p.isNamed)
	stateSinkObject.Save(2, &p.readers)
	stateSinkObject.Save(3, &p.writers)
	stateSinkObject.Save(4, &p.buf)
	stateSinkObject.Save(5, &p.off)
	stateSinkObject.Save(6, &p.size)
	stateSinkObject.Save(7, &p.max)
	stateSinkObject.Save(8, &p.hadWriter)
}

// +checklocksignore
func (p *Pipe) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.Queue)
	stateSourceObject.Load(1, &p.isNamed)
	stateSourceObject.Load(2, &p.readers)
	stateSourceObject.Load(3, &p.writers)
	stateSourceObject.Load(4, &p.buf)
	stateSourceObject.Load(5, &p.off)
	stateSourceObject.Load(6, &p.size)
	stateSourceObject.Load(7, &p.max)
	stateSourceObject.Load(8, &p.hadWriter)
	stateSourceObject.AfterLoad(p.afterLoad)
}

func (r *Reader) StateTypeName() string {
	return "pkg/sentry/kernel/pipe.Reader"
}

func (r *Reader) StateFields() []string {
	return []string{
		"ReaderWriter",
	}
}

func (r *Reader) beforeSave() {}

// +checklocksignore
func (r *Reader) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.ReaderWriter)
}

func (r *Reader) afterLoad() {}

// +checklocksignore
func (r *Reader) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.ReaderWriter)
}

func (rw *ReaderWriter) StateTypeName() string {
	return "pkg/sentry/kernel/pipe.ReaderWriter"
}

func (rw *ReaderWriter) StateFields() []string {
	return []string{
		"Pipe",
	}
}

func (rw *ReaderWriter) beforeSave() {}

// +checklocksignore
func (rw *ReaderWriter) StateSave(stateSinkObject state.Sink) {
	rw.beforeSave()
	stateSinkObject.Save(0, &rw.Pipe)
}

func (rw *ReaderWriter) afterLoad() {}

// +checklocksignore
func (rw *ReaderWriter) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &rw.Pipe)
}

func (vp *VFSPipe) StateTypeName() string {
	return "pkg/sentry/kernel/pipe.VFSPipe"
}

func (vp *VFSPipe) StateFields() []string {
	return []string{
		"pipe",
	}
}

func (vp *VFSPipe) beforeSave() {}

// +checklocksignore
func (vp *VFSPipe) StateSave(stateSinkObject state.Sink) {
	vp.beforeSave()
	stateSinkObject.Save(0, &vp.pipe)
}

func (vp *VFSPipe) afterLoad() {}

// +checklocksignore
func (vp *VFSPipe) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &vp.pipe)
}

func (fd *VFSPipeFD) StateTypeName() string {
	return "pkg/sentry/kernel/pipe.VFSPipeFD"
}

func (fd *VFSPipeFD) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"LockFD",
		"pipe",
	}
}

func (fd *VFSPipeFD) beforeSave() {}

// +checklocksignore
func (fd *VFSPipeFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &fd.LockFD)
	stateSinkObject.Save(4, &fd.pipe)
}

func (fd *VFSPipeFD) afterLoad() {}

// +checklocksignore
func (fd *VFSPipeFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &fd.LockFD)
	stateSourceObject.Load(4, &fd.pipe)
}

func (w *Writer) StateTypeName() string {
	return "pkg/sentry/kernel/pipe.Writer"
}

func (w *Writer) StateFields() []string {
	return []string{
		"ReaderWriter",
	}
}

func (w *Writer) beforeSave() {}

// +checklocksignore
func (w *Writer) StateSave(stateSinkObject state.Sink) {
	w.beforeSave()
	stateSinkObject.Save(0, &w.ReaderWriter)
}

func (w *Writer) afterLoad() {}

// +checklocksignore
func (w *Writer) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &w.ReaderWriter)
}

func init() {
	state.Register((*inodeOperations)(nil))
	state.Register((*Pipe)(nil))
	state.Register((*Reader)(nil))
	state.Register((*ReaderWriter)(nil))
	state.Register((*VFSPipe)(nil))
	state.Register((*VFSPipeFD)(nil))
	state.Register((*Writer)(nil))
}