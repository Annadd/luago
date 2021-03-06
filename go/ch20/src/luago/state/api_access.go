package state

import (
	"fmt"
	. "luago/api"
)

// http://www.lua.org/manual/5.3/manual.html#lua_rawlen
func (self *luaState) RawLen(idx int) uint {
	val := self.stack.get(idx)
	switch x := val.(type) {
	case string:
		return uint(len(x))
	case *luaTable:
		return uint(x.len())
	default:
		return 0
	}
}

// http://www.lua.org/manual/5.3/manual.html#lua_typename
func (self *luaState) TypeName(tp LuaType) string {
	switch tp {
	case LUA_TNONE:
		return "no value"
	case LUA_TNIL:
		return "nil"
	case LUA_TBOOLEAN:
		return "boolean"
	case LUA_TNUMBER:
		return "number"
	case LUA_TSTRING:
		return "string"
	case LUA_TTABLE:
		return "table"
	case LUA_TFUNCTION:
		return "function"
	case LUA_TTHREAD:
		return "thread"
	default:
		return "userdata"
	}
}

// http://www.lua.org/manual/5.3/manual.html#lua_type
func (self *luaState) Type(idx int) LuaType {
	if self.stack.isValid(idx) {
		val := self.stack.get(idx)
		return typeOf(val)
	}
	return LUA_TNONE
}

// http://www.lua.org/manual/5.3/manual.html#lua_isnone
func (self *luaState) IsNone(idx int) bool {
	return self.Type(idx) == LUA_TNONE
}

// http://www.lua.org/manual/5.3/manual.html#lua_isnil
func (self *luaState) IsNil(idx int) bool {
	return self.Type(idx) == LUA_TNIL
}

// http://www.lua.org/manual/5.3/manual.html#lua_isnoneornil
func (self *luaState) IsNoneOrNil(idx int) bool {
	return self.Type(idx) <= LUA_TNIL
}

// http://www.lua.org/manual/5.3/manual.html#lua_isboolean
func (self *luaState) IsBoolean(idx int) bool {
	return self.Type(idx) == LUA_TBOOLEAN
}

// http://www.lua.org/manual/5.3/manual.html#lua_istable
func (self *luaState) IsTable(idx int) bool {
	return self.Type(idx) == LUA_TTABLE
}

// http://www.lua.org/manual/5.3/manual.html#lua_isfunction
func (self *luaState) IsFunction(idx int) bool {
	return self.Type(idx) == LUA_TFUNCTION
}

// http://www.lua.org/manual/5.3/manual.html#lua_isthread
func (self *luaState) IsThread(idx int) bool {
	return self.Type(idx) == LUA_TTHREAD
}

// http://www.lua.org/manual/5.3/manual.html#lua_isstring
func (self *luaState) IsString(idx int) bool {
	t := self.Type(idx)
	return t == LUA_TSTRING || t == LUA_TNUMBER
}

// http://www.lua.org/manual/5.3/manual.html#lua_isnumber
func (self *luaState) IsNumber(idx int) bool {
	_, ok := self.ToNumberX(idx)
	return ok
}

// http://www.lua.org/manual/5.3/manual.html#lua_isinteger
func (self *luaState) IsInteger(idx int) bool {
	val := self.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

// http://www.lua.org/manual/5.3/manual.html#lua_iscfunction
func (self *luaState) IsGoFunction(idx int) bool {
	val := self.stack.get(idx)
	if c, ok := val.(*closure); ok {
		return c.goFunc != nil
	}

	return false
}

// http://www.lua.org/manual/5.3/manual.html#lua_toboolean
func (self *luaState) ToBoolean(idx int) bool {
	val := self.stack.get(idx)
	return convertToBoolean(val)
}

// http://www.lua.org/manual/5.3/manual.html#lua_tointeger
func (self *luaState) ToInteger(idx int) int64 {
	i, _ := self.ToIntegerX(idx)
	return i
}

// http://www.lua.org/manual/5.3/manual.html#lua_tointegerx
func (self *luaState) ToIntegerX(idx int) (int64, bool) {
	val := self.stack.get(idx)
	return convertToInteger(val)
}

// http://www.lua.org/manual/5.3/manual.html#lua_tonumber
func (self *luaState) ToNumber(idx int) float64 {
	n, _ := self.ToNumberX(idx)
	return n
}

// http://www.lua.org/manual/5.3/manual.html#lua_tonumberx
func (self *luaState) ToNumberX(idx int) (float64, bool) {
	val := self.stack.get(idx)
	return convertToFloat(val)
}

// http://www.lua.org/manual/5.3/manual.html#lua_tostring
func (self *luaState) ToString(idx int) string {
	s, _ := self.ToStringX(idx)
	return s
}

func (self *luaState) ToStringX(idx int) (string, bool) {
	val := self.stack.get(idx)
	switch x := val.(type) {
	case string:
		return x, true
	case int64, float64:
		s := fmt.Sprintf("%v", x)
		self.stack.set(idx, s)
		return s, true
	default:
		return "", false
	}
}

// http://www.lua.org/manual/5.3/manual.html#lua_tocfunction
func (self *luaState) ToGoFunction(idx int) GoFunction {
	val := self.stack.get(idx)
	if c, ok := val.(*closure); ok {
		return c.goFunc
	}

	return nil
}

// http://www.lua.org/manual/5.3/manual.html#lua_topointer
func (self *luaState) ToPointer(idx int) interface{} {
	// todo
	return self.stack.get(idx)
}
