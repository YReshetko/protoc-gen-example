// Code generated by protoc-gen-example. DO NOT EDIT.
package pb


var rolesMap = map[string][]string{	
	"/pb.Service/SomeAnotherMethod": {"SuperAdmin"},	
	"/pb.Service/SomeDifferentMethod": {"User", "Unknown"},	
	"/pb.Service/SomeMethod": {"Admin", "User"},	
}

func CheckRole(method, role string)bool{
	v, ok := rolesMap[method]	
	if !ok {	
		return false	
	}	
	for _, roleToCheck := range v {	
		if role == roleToCheck {		
			return true		
		}		
	}	
	return false
}