package conv

func (this *Map) IsDefault(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsDefault()
	}
	return this.Get(key[0]).IsDefault(key[1:]...)
}

func (this *Map) IsNumber(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsNumber()
	}
	return this.Get(key[0]).IsNumber(key[1:]...)
}

func (this *Map) IsString(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsString()
	}
	return this.Get(key[0]).IsString(key[1:]...)
}

func (this *Map) IsInt(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsInt()
	}
	return this.Get(key[0]).IsInt(key[1:]...)
}

func (this *Map) IsFloat(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsFloat()
	}
	return this.Get(key[0]).IsFloat(key[1:]...)
}

func (this *Map) IsBool(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsBool()
	}
	return this.Get(key[0]).IsBool(key[1:]...)
}

func (this *Map) IsTime(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsTime()
	}
	return this.Get(key[0]).IsTime(key[1:]...)
}

func (this *Map) IsArray(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsArray()
	}
	return this.Get(key[0]).IsArray(key[1:]...)
}

func (this *Map) IsPointer(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsPointer()
	}
	return this.Get(key[0]).IsPointer(key[1:]...)
}

func (this *Map) IsNil(key ...string) bool {
	if len(key) == 0 {
		return this.Var.IsNil()
	}
	return this.Get(key[0]).IsNil(key[1:]...)
}
