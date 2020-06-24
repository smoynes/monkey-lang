package evaluator

import (
	"fmt"

	"monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len":   &object.Builtin{Fn: len_builtin},
	"first": &object.Builtin{Fn: first_builtin},
	"last": &object.Builtin{Fn: last_builtin},
	"rest": &object.Builtin{Fn: rest_builtin},
	"push": &object.Builtin{Fn: push_builtin},
	"puts": &object.Builtin{Fn: puts_builtin},
}

func len_builtin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		length := int64(len(arg.Value))
		return &object.Integer{Value: length}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

func first_builtin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments, got=%d want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	} else {
		return NULL
	}
}

func last_builtin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments, got=%d want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) == 0 {
		return NULL
	} else {
		return arr.Elements[len(arr.Elements) - 1]
	}
}

func rest_builtin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments, got=%d want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]object.Object, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &object.Array{Elements: newElements}
	} else {
		return NULL
	}
}

func push_builtin(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments, got=%d want=2", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	newElements := make([]object.Object, length + 1, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &object.Array{Elements: newElements}
}

func puts_builtin(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}

	return NULL
}
