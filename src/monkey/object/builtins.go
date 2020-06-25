package object

import (
	"fmt"
)

var Builtins = map[string]*Builtin{
	"len":   &Builtin{Fn: Len},
	"first": &Builtin{Fn: First},
	"last":  &Builtin{Fn: Last},
	"rest":  &Builtin{Fn: Rest},
	"push":  &Builtin{Fn: Push},
	"puts":  &Builtin{Fn: Puts},
}

func Len(args ...Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	switch arg := args[0].(type) {
	case *String:
		length := int64(len(arg.Value))
		return &Integer{Value: length}
	case *Array:
		return &Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

func First(args ...Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments, got=%d want=1", len(args))
	}

	if args[0].Type() != ARRAY_OBJ {
		return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	} else {
		return nil
	}
}

func Last(args ...Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments, got=%d want=1", len(args))
	}

	if args[0].Type() != ARRAY_OBJ {
		return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*Array)
	if len(arr.Elements) == 0 {
		return nil
	} else {
		return arr.Elements[len(arr.Elements)-1]
	}
}

func Rest(args ...Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments, got=%d want=1", len(args))
	}

	if args[0].Type() != ARRAY_OBJ {
		return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*Array)
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]Object, length-1, length-1)
		copy(newElements, arr.Elements[1:length])
		return &Array{Elements: newElements}
	} else {
		return nil
	}
}

func Push(args ...Object) Object {
	if len(args) != 2 {
		return newError("wrong number of arguments, got=%d want=2", len(args))
	}

	if args[0].Type() != ARRAY_OBJ {
		return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*Array)
	length := len(arr.Elements)
	newElements := make([]Object, length+1, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &Array{Elements: newElements}
}

func Puts(args ...Object) Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}

	return nil
}
