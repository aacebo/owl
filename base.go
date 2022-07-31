package owl

type Schema interface {
	Validate(v any) []*Error
}

type BaseSchema struct {
	name       string
	required   bool
	operations []*Operation
}

func (self *BaseSchema) AddOperation(op *Operation) *BaseSchema {
	self.operations = append(self.operations, op)
	return self
}

func (self *BaseSchema) Required() *BaseSchema {
	self.required = true
	return self
}

func (self *BaseSchema) Message(message string) *BaseSchema {
	if length := len(self.operations); length > -1 {
		self.operations[length-1].message = message
	}

	return self
}

func (self *BaseSchema) Validate(v any) []*Error {
	errors := []*Error{}
	valid := false

	for _, op := range self.operations {
		if v, valid = op.Eval(v); !valid {
			errors = append(errors, NewError(self.name, op.message))
		}
	}

	return errors
}
