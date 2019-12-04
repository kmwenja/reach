package reach

import (
	"fmt"

	"github.com/d5/tengo/objects"
	"github.com/d5/tengo/script"
	"github.com/d5/tengo/stdlib"
)

// Run takes a byte array containing the reach script
// and executes it.
func Run(sb []byte, args []string, data map[string]interface{}) error {
	s := script.New(sb)

	argsArray := &objects.Array{Value: make([]objects.Object, 0)}
	for _, i := range args {
		argsArray.Value = append(argsArray.Value, &objects.String{Value: i})
	}

	dataMap, err := objects.FromInterface(data)
	if err != nil {
		return fmt.Errorf("could not inject data into script: %v", err)
	}

	reachModuleMap := objects.NewModuleMap()
	reachModuleMap.AddBuiltinModule("reach", map[string]objects.Object{
		"args": argsArray,
		"data": dataMap,
		"test": &objects.Map{Value: map[string]objects.Object{
			"a": &objects.Int{Value: int64(1)},
			"b": &objects.Int{Value: int64(2)},
		}},
		// TODO connect
		"connect":  objects.UndefinedValue,
		"register": objects.UndefinedValue,
	})
	reachModuleMap.AddMap(stdlib.GetModuleMap("fmt"))

	s.SetImports(reachModuleMap)

	cs, err := s.Compile()
	if err != nil {
		return fmt.Errorf("could not compile script: %v", err)
	}

	if err := cs.Run(); err != nil {
		return fmt.Errorf("could not run script at: %v", err)
	}

	return nil
}
