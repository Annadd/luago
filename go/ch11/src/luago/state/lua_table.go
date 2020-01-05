package state

import "math"
import "luago/number"

type luaTbale struct {
	metatable *luaTbale
	arr       []luaValue
	_map      map[luaValue]luaValue
}

func newLuaTable(nArr, nRec int) *luaTbale {
	t := &luaTbale{}

	if nArr > 0 {
		t.arr = make([]luaValue, 0, nArr)
	}

	if nRec > 0 {
		t._map = make(map[luaValue]luaValue, nRec)
	}

	return t
}

func (self *luaTbale) get(key luaValue) luaValue {
	key = _floatToIntger(key)

	if idx, ok := key.(int64); ok {
		if idx >= 1 && idx <= int64(len(self.arr)) {
			return self.arr[idx-1]
		}
	}

	return self._map[key]
}

func _floatToIntger(key luaValue) luaValue {
	if f, ok := key.(float64); ok {
		if i, ok := number.FloatToInteger(f); ok {
			return i
		}
	}

	return key
}

func (self *luaTbale) put(key, val luaValue) {
	if key == nil {
		panic("table index is nil!")
	}

	if f, ok := key.(float64); ok && math.IsNaN(f) {
		panic("table index is NaN")
	}

	key = _floatToIntger(key)
	if idx, ok := key.(int64); ok && idx >= 1 {
		arrLen := int64(len(self.arr))
		if idx <= arrLen {
			self.arr[idx-1] = val
			if idx == arrLen && val == nil {
				self._shrinkArray()
			}
			return
		}

		if idx == arrLen+1 {
			delete(self._map, key)
			if val != nil {
				self.arr = append(self.arr, val)
				self._expandArray()
			}
			return
		}
	}

	if val != nil {
		if self._map == nil {
			self._map = make(map[luaValue]luaValue, 8)
		}
		self._map[key] = val
	} else {
		delete(self._map, key)
	}
}

func (self *luaTbale) _shrinkArray() {
	for i := len(self.arr) - 1; i >= 0; i-- {
		if self.arr[i] == nil {
			self.arr = self.arr[0:i]
		}
	}
}

func (self *luaTbale) _expandArray() {
	for idx := int64(len(self.arr)) + 1; true; idx++ {
		if val, found := self._map[idx]; found {
			delete(self._map, idx)
			self.arr = append(self.arr, val)
		} else {
			break
		}
	}
}

func (self *luaTbale) len() int {
	return len(self.arr)
}

func (self *luaTbale) hasMetafield(fieldName string) bool {
	return self.metatable != nil && self.metatable.get(fieldName) != nil
}