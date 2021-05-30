package deepcopy

import (
	"reflect"
)

func Copy(object interface{}) (interface{}, error) {
	d := reflect.New(reflect.TypeOf(object)).Elem()
	err := copy(&d, reflect.ValueOf(object))
	return d.Interface(), err
}

func copy(d *reflect.Value, f reflect.Value) error {
	//fmt.Printf("start %v to %v\n", f.Interface(), d.Interface())
	switch f.Kind() {
	case reflect.Array:
		*d = reflect.New(f.Type()).Elem()
		for i := 0; i < f.Len(); i++ {
			n := reflect.New(f.Index(i).Type()).Elem()
			if err := copy(&n, f.Index(i)); err != nil {
				return err
			}
			d.Index(i).Set(n)
		}
	case reflect.Slice:
		*d = reflect.MakeSlice(f.Type(), f.Len(), f.Cap())
		for i := 0; i < f.Len(); i++ {
			n := reflect.New(f.Index(i).Type()).Elem()
			if err := copy(&n, f.Index(i)); err != nil {
				return err
			}
			d.Index(i).Set(n)
		}
	case reflect.Interface:
		fallthrough
	case reflect.Ptr:
		c := reflect.Value{}
		if err := copy(&c, f.Elem()); err != nil {
			return err
		}
		*d = reflect.New(c.Type())
		d.Elem().Set(c)
	case reflect.Struct:
		*d = reflect.New(f.Type()).Elem()
		for i := 0; i < f.NumField(); i++ {
			v := d.Field(i)
			if err := copy(&v, f.Field(i)); err != nil {
				return err
			}
			d.Field(i).Set(v)
		}
	case reflect.Chan:
		*d = reflect.MakeChan(f.Type(), f.Cap())
	case reflect.Map:
		iter := f.MapRange()
		*d = reflect.MakeMap(f.Type())
		for iter.Next() {
			d.SetMapIndex(iter.Key(), iter.Value())
		}
	default:
		*d = f
	}
	return nil
}
