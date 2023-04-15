package object

import (
	"encoding/json"
)

type Object map[string]any

func (o Object) Object(property string) Object {
	if o == nil {
		return nil
	}

	return AsObject(o[property])
}

func (o Object) Copy() Object {
	if o == nil {
		return nil
	}

	object := make(Object)
	for key, value := range o {
		if IsObject(value) {
			object[key] = AsObject(value).Copy()
		} else {
			object[key] = value
		}
	}

	return object
}

func (o Object) Merge(object Object) Object {
	if o == nil {
		return object
	}

	return o.Copy().ApplyFrom(object)
}

func (o Object) ApplyFrom(src Object) Object {
	if o == nil {
		return src
	}

	for key, value := range src {
		if _, ok := o[key]; ok {
			if IsObject(value) && IsObject(o[key]) {
				o[key] = AsObject(o[key]).ApplyFrom(AsObject(value))
			} else {
				o[key] = value
			}
		} else {
			o[key] = value
		}
	}

	return o
}

func (o Object) String() string {
	if o == nil {
		return ""
	}

	raw, err := json.Marshal(o)
	if err != nil {
		return ""
	}

	return string(raw)
}

func AsObject(object any) Object {
	if v, ok := object.(map[string]any); ok {
		return v
	}
	if v, ok := object.(Object); ok {
		return v
	}
	return nil
}

func IsObject(object any) bool {
	if object == nil {
		return false
	}

	_, ok := object.(map[string]any)
	if !ok {
		_, ok = object.(Object)
	}

	return ok
}

func CreateObjects(options ...func(object Object)) Object {
	object := make(Object)

	for _, opt := range options {
		opt(object)
	}

	return object
}

func Property(key string, value any) func(Object) {
	return func(object Object) {
		if object != nil {
			object[key] = value
		}
	}
}
