package owl

type BaseSchema struct {
	name       string
	index      int
	operations []*Operation
}

func Base(name string, op *Operation) *BaseSchema {
	v := BaseSchema{
		name:       name,
		index:      0,
		operations: []*Operation{op},
	}

	return &v
}

func (self *BaseSchema) AddOperation(op *Operation) *BaseSchema {
	self.operations = append(self.operations, op)
	self.index++
	return self
}

func (self *BaseSchema) CurrentOperation() *Operation {
	if self.index == -1 {
		return nil
	}

	return self.operations[self.index]
}

func (self *BaseSchema) Required() *BaseSchema {
	return self
}

func (self *BaseSchema) Message(message string) *BaseSchema {
	if op := self.CurrentOperation(); op != nil {
		op.Message = message
	}

	return self
}

func (self *BaseSchema) Validate(v any) []*Error {
	errors := []*Error{}
	valid := false

	for _, op := range self.operations {
		if v, valid = op.Eval(v); !valid {
			errors = append(errors, NewError(self.name, op.Message))
		}
	}

	return errors
}
