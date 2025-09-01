import { useEffect, useState } from "react";
import User from '../components/user'
import  {type UserType} from '../types/usertype';

export default function UsersPage() {
  const [users, setUsers] = useState<UserType[]>([]);
  const [loading, setLoading] = useState(true);
  const [skip,setSkip]=useState(0);
  const [limit,setLimit]=useState(10);
 
    useEffect(() => {
    fetch("https://dummyjson.com/users?limit="+limit+"&+skip="+skip)
      .then((res) => res.json())
      .then((data) => {
        setUsers(data.users);
        setLoading(false);
      })
      .catch((err) => {
        console.error("Error fetching users:", err);
        setLoading(false);
      });
  }, []);
  if (loading) return <p className="text-center">Loading...</p>;

//"id":1,"firstName":"Emily","lastName":"Johnson","maidenName":"Smith","age":28,"gender":"female","email":"emily.johnson@x.dummyjson.com","phone":"+81 965-431-3024"
    
  return (
    <div className="p-6">
      <h2 className="text-2xl font-bold mb-4">Users List</h2>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
        {users.map((user) => (
           <User user={user}/>
          //  
        ))}
        <button>Next</button>
      </div>
    </div>
  );
}