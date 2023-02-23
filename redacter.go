package redacter

import (
	"strings"
)

const defaultCensor = "[REDACTED]"

type Redacter struct {
	fields []string
	censor any
	strict bool
	remove bool
}

func New() *Redacter {
	return &Redacter{
		censor: defaultCensor,
		remove: false,
		strict: true,
	}
}

// WithFields sets the name of the fields to be redacted.
func (r *Redacter) WithFields(v ...string) *Redacter {
	r.fields = v
	return r
}

// WithCensor sets a value to overwride redacted filds.
// The default value is [REDACTED].
func (r *Redacter) WithCensor(v any) *Redacter {
	r.censor = v
	return r
}

// WithRemove sets whether the redacted fields will be removed.
// By default is false.
func (r *Redacter) WithRemove(v bool) *Redacter {
	r.remove = v
	return r
}

// WithStrict sets whether the fields names will be case sensitive.
// By default is true.
func (r *Redacter) WithStrict(v bool) *Redacter {
	r.strict = v
	return r
}

func (r *Redacter) shouldRedact(v string) bool {
	if !r.strict {
		v = strings.ToLower(v)
	}

	for _, field := range r.fields {
		if !r.strict {
			field = strings.ToLower(field)
		}

		if field == v {
			return true
		}
	}

	return false
}

func (r *Redacter) redact(obj any) any {
	switch value := obj.(type) {
	case map[string]any:
		for k, v := range value {
			if r.shouldRedact(k) {
				if r.remove {
					delete(value, k)
				} else {
					value[k] = r.censor
				}
			} else {
				value[k] = r.redact(v)
			}
		}

		return value
	case []any:
		for i, v := range value {
			value[i] = r.redact(v)
		}

		return value
	default:
		return value
	}
}
